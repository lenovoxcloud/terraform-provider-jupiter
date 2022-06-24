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

func dataSourceVolume() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVolumeRead,
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
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"is_thin_provisioning": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"size": &schema.Schema{
							Type:     schema.TypeInt,
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
					},
				},
			},
		},
	}
}

func dataSourceVolumeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*apiClient)
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
			"vm_uuid":           volumeinfo.VData.Detail.Vmuuid,
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
