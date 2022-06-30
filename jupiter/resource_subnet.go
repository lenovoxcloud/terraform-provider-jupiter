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

func resourceSubnet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSubnetCreate,
		ReadContext:   resourceSubnetRead,
		UpdateContext: resourceSubnetUpdate,
		DeleteContext: resourceSubnetDelete,
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
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"network_type": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"segmentation_id": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"subnet_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"vpc_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"subnets": &schema.Schema{
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allocation_pools": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"cidr": &schema.Schema{
													Type:     schema.TypeString,
													Required: true,
												},
												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"dns_nameservers": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"enable_dhcp": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"gateway_ip": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"host_routes": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ip_version": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
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

func resourceSubnetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*apiClient)
	var diags diag.Diagnostics
	jupiter_url := c.jupiter_url
	auth_token := c.authToken
	jupiter_route := "/jupiter/api/v1/smf/network"
	items := d.Get("items").([]interface{})

	for _, item := range items {
		i := item.(map[string]interface{})

		vm := i["vm"].([]interface{})[0]
		vm_interface := vm.(map[string]interface{})
		cookie := "authToken" + auth_token
		name := vm_interface["name"].(string)
		segmentation_id := vm_interface["segmentation_id"].(string)
		network_type := vm_interface["network_type"].(string)
		description := vm_interface["description"].(string)
		subnet_uuid := vm_interface["subnet_uuid"].(string)
		subnets := vm_interface["subnets"].(string)
		vpc_uuid := vm_interface["vpc_uuid"].(string)

		payload := strings.NewReader(
			fmt.Sprintf(`{"name":"%s","segmentation_id":"%s","network_type":%s,"description":"%s","subnet_uuid":"%s","vpc_uuid":"%s","subnets":%s}`,
				name, segmentation_id, network_type, description, subnet_uuid, vpc_uuid, subnets))

		req, _ := http.NewRequest("POST", jupiter_url+jupiter_route, payload)

		req.Header.Add("Content-Type", "application/json")

		req.Header.Add("Cookie", cookie)

		response, err := http.DefaultClient.Do(req)
		defer response.Body.Close()
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(name)
	}
	resourceSubnetRead(ctx, d, m)
	return diags

}

//diag.Diagnostics
func resourceSubnetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Printf(err.(string))
	// 	}
	// }()
	// c := m.(*apiClient)
	var diags diag.Diagnostics

	return diags
}

func resourceSubnetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceSubnetRead(ctx, d, m)
	return diags
}

func resourceSubnetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
