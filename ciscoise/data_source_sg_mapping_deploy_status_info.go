package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSgMappingDeployStatusInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMapping.

- This data source action allows the client to get the IP to SGT mapping deployment status.
Deploy Status will show last Deploy command output. The information will be saved until the next Deploy command
`,

		ReadContext: dataSourceSgMappingDeployStatusInfoRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"result_value": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceSgMappingDeployStatusInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeployStatusIPToSgtMapping")

		response1, restyResp1, err := client.IPToSgtMapping.GetDeployStatusIPToSgtMapping()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeployStatusIPToSgtMapping", err,
				"Failure at GetDeployStatusIPToSgtMapping, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenIPToSgtMappingGetDeployStatusIPToSgtMappingItem(response1.OperationResult)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeployStatusIPToSgtMapping response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIPToSgtMappingGetDeployStatusIPToSgtMappingItem(item *isegosdk.ResponseIPToSgtMappingGetDeployStatusIPToSgtMappingOperationResult) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["result_value"] = flattenIPToSgtMappingGetDeployStatusIPToSgtMappingItemResultValue(item.ResultValue)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIPToSgtMappingGetDeployStatusIPToSgtMappingItemResultValue(items *[]isegosdk.ResponseIPToSgtMappingGetDeployStatusIPToSgtMappingOperationResultResultValue) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["value"] = item.Value
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
