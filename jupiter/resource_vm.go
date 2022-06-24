package jupiter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVM() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVMCreate,
		ReadContext:   resourceVMRead,
		UpdateContext: resourceVMUpdate,
		DeleteContext: resourceVMDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vm_lookup_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"project_global_id": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"cloud_name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"flavor_id": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"image_id": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"vpc_id": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"network_id": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"password_type": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"password": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"vm_status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"power_state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vm_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vm_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"quantity": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*apiClient)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	auth_token := c.authToken
	jupiter_route := "/jupiter/api/v1/application"
	items := d.Get("items").([]interface{})

	for _, item := range items {
		i := item.(map[string]interface{})

		vm := i["vm"].([]interface{})[0]
		vm_interface := vm.(map[string]interface{})
		cookie := "authToken" + auth_token
		instance_name := vm_interface["instance_name"].(string)
		project_global_id := vm_interface["project_global_id"].(string)
		cloud_name := vm_interface["cloud_name"].(string)
		flavor_id := vm_interface["flavor_id"].(string)
		image_id := vm_interface["image_id"].(string)
		vpc_id := vm_interface["vpc_id"].(string)
		network_id := vm_interface["network_id"].(string)
		password_type := vm_interface["password_type"].(string)
		password := vm_interface["password"].(string)

		payload := strings.NewReader(
			fmt.Sprintf(`{"instance_name":"%s","project_global_id":"%s","cloud_name":"%s","flavor_id":"%s","image_id":"%s", "volumes":[],"vpc_id":"%s","network_id":"%s","count":1,"password_type":"%s","security_groups":[],"password":"%s"}`,
				instance_name, project_global_id, cloud_name, flavor_id, image_id, vpc_id, network_id, password_type, password))

		req, _ := http.NewRequest("PUT", jupiter_url+jupiter_route, payload)

		req.Header.Add("Content-Type", "application/json")

		req.Header.Add("Cookie", cookie)

		response, err := http.DefaultClient.Do(req)
		defer response.Body.Close()
		if err != nil {
			return diag.FromErr(err)
		} else {
			body, _ := ioutil.ReadAll(response.Body)
			bodystr := string(body)
			var dataAttr map[string]interface{}

			if err := json.Unmarshal([]byte(bodystr), &dataAttr); err == nil {
				for key, value := range dataAttr {
					if key == "data" {
						for vkey, vvalue := range value.(map[string]interface{}) {
							if vkey == "request_id" {
								// vmmap := vvalue.(map[string]interface{})
								d.Set("vm_lookup_key", vvalue.(string))
							}
						}
					}
				}
			} else {
				return diag.FromErr(err)
			}
		}
		d.SetId(instance_name)
	}
	time.Sleep(time.Second * 1)
	resourceVMRead(ctx, d, m)
	return diags

}

//diag.Diagnostics
func resourceVMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Printf(err.(string))
	// 	}
	// }()
	c := m.(*apiClient)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	auth_token := c.authToken
	// jupiter_route := "/jupiter/api/v1/delivery/"
	jupiter_route := "/jupiter/api/v1/application/detail/"
	vm_lookup_key := d.Get("vm_lookup_key").(string)
	cookie := "authToken" + auth_token
	req, _ := http.NewRequest("GET", jupiter_url+jupiter_route+vm_lookup_key, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Authorization", auth_token)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	bodystr := string(body)

	application := &ApplicationData{}
	if err := json.Unmarshal([]byte(bodystr), &application); err == nil {
		flavor_id := application.AData.Request.FlavorID
		image_id := application.AData.Request.ImageID
		vpc_id := application.AData.Request.VPCID
		network_id := application.AData.Request.FlavorID
		items := make([]interface{}, len(application.AData.Servers), len(application.AData.Servers))
		i := 0
		for key, value := range application.AData.Servers {
			fmt.Println(key)
			vm := map[string]interface{}{
				"instance_name":     value.VMName,
				"project_global_id": value.ProjectGlobalID,
				"cloud_name":        value.CloudName,
				"flavor_id":         flavor_id,
				"image_id":          image_id,
				"vpc_id":            vpc_id,
				"network_id":        network_id,
				"vm_status":         value.Status,
				"power_state":       value.Status,
				"vm_uuid":           value.VMUUid,
				"vm_id":             value.ID,
			}
			oi := make(map[string]interface{})
			oi["vm"] = vm
			oi["quantity"] = value.ServerNo
			items[i] = oi
			i += 1
		}
		d.Set("items", items)

	} else {
		return diag.FromErr(err)
	}

	return diags
}

func resourceVMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*apiClient)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	auth_token := c.authToken
	cookie := "authToken" + auth_token

	if d.HasChange("items") {
		items := d.Get("items").([]interface{})
		for _, item := range items {
			i := item.(map[string]interface{})
			vm := i["vm"].([]interface{})[0]
			vm_interface := vm.(map[string]interface{})
			power_state := vm_interface["power_state"].(string)
			log.Println("power_state" + power_state)
			if power_state == "shutoff" {
				log.Println("shutoff start")
				jupiter_route := "/jupiter/api/v1/vm/power"
				// powerStateOldRaw, powerStateNewRaw := d.GetChange("power_state")
				// powerStateOld := powerStateOldRaw.(string)
				// fmt.Println(powerStateOld)
				// powerStateNew := powerStateNewRaw.(string)
				// if strings.ToLower(powerStateNew) == "shutoff" {
				vm_uuid := d.Get("vm_lookup_key").(string)
				reason := "terraform operate"
				operation := "stop"

				payload := strings.NewReader(
					fmt.Sprintf(`{"vm_uuids":["%s"],"reason":"%s","operation":"%s"}`,
						vm_uuid, reason, operation))

				req, _ := http.NewRequest("POST", jupiter_url+jupiter_route, payload)

				req.Header.Add("Content-Type", "application/json")

				req.Header.Add("Cookie", cookie)

				response, err := http.DefaultClient.Do(req)
				log.Println(err)
				log.Println(response)
				if err != nil {
					return diag.FromErr(err)
				}

			} else if power_state == "active" {
				jupiter_route := "/jupiter/api/v1/vm/power"
				vm_uuid := d.Get("vm_lookup_key").(string)
				reason := "terraform operate"
				operation := "start"

				payload := strings.NewReader(
					fmt.Sprintf(`{"vm_uuids":["%s"],"reason":"%s","operation":"%s"}`,
						vm_uuid, reason, operation))

				req, _ := http.NewRequest("POST", jupiter_url+jupiter_route, payload)

				req.Header.Add("Content-Type", "application/json")

				req.Header.Add("Cookie", cookie)

				_, err := http.DefaultClient.Do(req)
				if err != nil {
					return diag.FromErr(err)
				}
			}
		}

	}

	resourceVMRead(ctx, d, m)
	return diags
}

func resourceVMDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*apiClient)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	auth_token := c.authToken
	cookie := "authToken" + auth_token
	jupiter_route := "/jupiter/api/v1/vm/delete"

	items := d.Get("items").([]interface{})

	for _, item := range items {
		i := item.(map[string]interface{})

		vm := i["vm"].([]interface{})[0]
		vm_interface := vm.(map[string]interface{})
		vm_uuid := vm_interface["vm_uuid"].(string)
		payload := strings.NewReader(fmt.Sprintf(`{"vm_uuid":"%s"}`, vm_uuid))

		req, _ := http.NewRequest("DELETE", jupiter_url+jupiter_route, payload)

		req.Header.Add("Content-Type", "application/json")

		req.Header.Add("Cookie", cookie)

		_, err := http.DefaultClient.Do(req)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	resourceVMRead(ctx, d, m)
	return diags
}
