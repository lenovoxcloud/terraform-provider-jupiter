package jupiter

import (
	"context"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVM() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVMRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
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
						"last_updated": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
					},
				},
			},
		},
	}
}

func dataSourceVMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*AuthStruct)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	jupiter_route := "/jupiter/api/v1/delivery/"
	vm_lookup_key := d.Get("id").(string)
	// vm_create_route := ""

	auth_token := c.authToken
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
	var dataAttr map[string]interface{}

	if err := json.Unmarshal([]byte(bodystr), &dataAttr); err == nil {
		for key, value := range dataAttr {
			if key == "data" {
				for vkey, vvalue := range value.(map[string]interface{}) {
					if vkey == "vm" {
						vmmap := vvalue.(map[string]interface{})
						vm := make(map[string]string)
						vm["instance_name"] = vmmap["vm_name"].(string)
						vm["project_global_id"] = vmmap["project_global_id"].(string)
						vm["cloud_name"] = vmmap["cloud_name"].(string)
						vm["flavor_id"] = vmmap["flavor_id"].(string)
						vm["image_id"] = vmmap["image_id"].(string)
						vm["vpc_id"] = vmmap["vpc_id"].(string)
						vm["network_id"] = vmmap["network_id"].(string)
						oi := make(map[string]interface{})

						oi["vm"] = vm
						oi["quantity"] = 1
						items := make([]interface{}, 1, 1)
						items[0] = oi
						d.Set("items", items)
					}
				}
			}
		}
	} else {
		return diag.FromErr(err)
	}

	return diags
}
