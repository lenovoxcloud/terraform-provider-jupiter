package jupiter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	// "log"
	// "time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVolume() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVolumeCreate,
		ReadContext:   resourceVolumeRead,
		UpdateContext: resourceVolumeUpdate,
		DeleteContext: resourceVolumeDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"volume_lookup_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
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

									"is_thin_provisioning": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"size": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"user": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"volume_feature": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"volume_type": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"vm_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
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

func resourceVolumeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*AuthStruct)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	auth_token := c.authToken
	jupiter_route := "/jupiter/api/v1/volume/create/"
	items := d.Get("items").([]interface{})

	for _, item := range items {
		i := item.(map[string]interface{})

		volume := i["volume"].([]interface{})[0]
		volume_interface := volume.(map[string]interface{})
		cookie := "authToken" + auth_token
		name := volume_interface["name"].(string)
		project_global_id := volume_interface["project_global_id"].(string)
		cloud_name := volume_interface["cloud_name"].(string)
		description := volume_interface["description"].(string)
		is_thin_provisioning := volume_interface["is_thin_provisioning"].(string)
		size := volume_interface["size"].(string)
		user := volume_interface["user"].(string)
		volume_feature := volume_interface["volume_feature"].(string)
		volume_type := volume_interface["volume_type"].(string)

		payload := strings.NewReader(
			fmt.Sprintf(`{"name":"%s","cloud_name":"%s","description":"%s","is_thin_provisioning":"%s", "size":%s,"user":"%s","volume_feature":"%s","volume_type":"%s"}`,
				name, cloud_name, description, is_thin_provisioning, size, user, volume_feature, volume_type))

		req, _ := http.NewRequest("PUT", jupiter_url+jupiter_route+project_global_id, payload)

		req.Header.Add("Content-Type", "application/json")

		req.Header.Add("Cookie", cookie)

		response, err := http.DefaultClient.Do(req)
		defer response.Body.Close()
		if err != nil {
			return diag.FromErr(err)
		} else {
			d.Set("volume_lookup_key", name)
			// body, _ := ioutil.ReadAll(response.Body)
			// bodystr := string(body)
			// var dataAttr map[string]interface{}

			// if err := json.Unmarshal([]byte(bodystr), &dataAttr); err == nil {
			// 	for key, value := range dataAttr {
			// 		if key == "data" {
			// 			for vkey, vvalue := range value.(map[string]interface{}) {
			// 				if vkey == "request_id" {
			// 					// vmmap := vvalue.(map[string]interface{})

			// 				}
			// 			}
			// 		}
			// 	}
			// } else {
			// 	return diag.FromErr(err)
			// }
		}
		d.SetId(name)
	}
	// time.Sleep(time.Second * 60)
	// resourceVolumeRead(ctx, d, m)
	return diags

}

//diag.Diagnostics
func resourceVolumeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Printf(err.(string))
	// 	}
	// }()
	c := m.(*AuthStruct)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	auth_token := c.authToken
	// jupiter_route := "/jupiter/api/v1/delivery/"
	jupiter_route := "/jupiter/api/v1/volume/detail/"
	volume_lookup_key := d.Get("volume_lookup_key").(string)
	cookie := "authToken" + auth_token
	req, _ := http.NewRequest("GET", jupiter_url+jupiter_route+volume_lookup_key, nil)
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

	volumeinfo := &VolumeData{}
	if err := json.Unmarshal([]byte(bodystr), &volumeinfo); err == nil {
		items := make([]interface{}, 1, 1)
		volume := map[string]interface{}{
			"name":              volumeinfo.VData.Detail.Name,
			"project_global_id": volumeinfo.VData.Detail.ProjectGlobalID,
			"cloud_name":        volumeinfo.VData.Detail.CloudName,
			"description":       volumeinfo.VData.Detail.Description,
			"size":              volumeinfo.VData.Detail.Size,
			"user":              volumeinfo.VData.Detail.UserId,
			"volume_feature":    volumeinfo.VData.Detail.VolumeFeature,
			"volume_type":       volumeinfo.VData.Detail.VolumeType,
			"volume_uuid":       volumeinfo.VData.Detail.Name,
		}
		oi := make(map[string]interface{})
		oi["volume"] = volume
		oi["quantity"] = 1
		items[0] = oi

		d.Set("items", items)

	} else {
		return diag.FromErr(err)
	}

	return diags
}

func resourceVolumeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*AuthStruct)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	auth_token := c.authToken
	cookie := "authToken" + auth_token
	if d.HasChange("size") {
		volume_lookup_key := d.Get("volume_lookup_key").(string)
		jupiter_route := "/jupiter/api/v1/vm/" + volume_lookup_key + "/size"
		sizeOldRaw, sizeNewRaw := d.GetChange("size")
		sizeOld := sizeOldRaw.(string)
		fmt.Println(sizeOld)
		newsize := sizeNewRaw.(string)
		payload := strings.NewReader(
			fmt.Sprintf(`{"size":%s}`, newsize))

		req, _ := http.NewRequest("PUT", jupiter_url+jupiter_route, payload)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Cookie", cookie)

		_, err := http.DefaultClient.Do(req)
		if err != nil {
			return diag.FromErr(err)
		}

	}

	resourceVolumeRead(ctx, d, m)
	return diags
}

func resourceVolumeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*AuthStruct)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	auth_token := c.authToken
	cookie := "authToken" + auth_token
	jupiter_route := "/jupiter/api/v1/volume/delete"

	items := d.Get("items").([]interface{})

	for _, item := range items {
		i := item.(map[string]interface{})

		vm := i["vm"].([]interface{})[0]
		vm_interface := vm.(map[string]interface{})
		vm_uuid := vm_interface["vm_uuid"].(string)
		payload := strings.NewReader(fmt.Sprintf(`{"volume_uuids":["%s"]}`, vm_uuid))

		req, _ := http.NewRequest("DELETE", jupiter_url+jupiter_route, payload)

		req.Header.Add("Content-Type", "application/json")

		req.Header.Add("Cookie", cookie)

		_, err := http.DefaultClient.Do(req)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	resourceVolumeRead(ctx, d, m)
	return diags
}
