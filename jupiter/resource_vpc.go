package jupiter

import (
	"context"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "log"
	"net/http"
	"strings"
	// "time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVPC() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVPCCreate,
		ReadContext:   resourceVPCRead,
		UpdateContext: resourceVPCUpdate,
		DeleteContext: resourceVPCDelete,
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
									"vpc_name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"owner": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"platform_name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"project_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"vpc_uuid": &schema.Schema{
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

func resourceVPCCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*apiClient)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	auth_token := c.authToken
	jupiter_route := "/jupiter/api/v1/smf/vpc"
	items := d.Get("items").([]interface{})

	for _, item := range items {
		i := item.(map[string]interface{})

		vm := i["vm"].([]interface{})[0]
		vm_interface := vm.(map[string]interface{})
		cookie := "authToken" + auth_token
		vpc_name := vm_interface["vpc_name"].(string)
		vtype := vm_interface["type"].(string)
		owner := vm_interface["owner"].(string)
		description := vm_interface["description"].(string)
		platform_name := vm_interface["platform_name"].(string)
		project_uuid := vm_interface["project_uuid"].(string)

		payload := strings.NewReader(
			fmt.Sprintf(`{"vpc_name":"%s","type":"%s","owner":%s,"description":"%s","platform_info":{"platform_name":"%s", "project_uuid":[]}}`,
				vpc_name, vtype, owner, description, platform_name, project_uuid))

		req, _ := http.NewRequest("POST", jupiter_url+jupiter_route, payload)

		req.Header.Add("Content-Type", "application/json")

		req.Header.Add("Cookie", cookie)

		response, err := http.DefaultClient.Do(req)
		defer response.Body.Close()
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(vpc_name)
	}
	resourceVPCRead(ctx, d, m)
	return diags

}

//diag.Diagnostics
func resourceVPCRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Printf(err.(string))
	// 	}
	// }()
	// c := m.(*apiClient)
	var diags diag.Diagnostics

	return diags
}

func resourceVPCUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceVPCRead(ctx, d, m)
	return diags
}

func resourceVPCDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
