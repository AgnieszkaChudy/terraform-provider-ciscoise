package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntSessionDisconnect() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMntSessionDisconnectRead,
		Schema: map[string]*schema.Schema{
			"dis_con_nec_tty_pe": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"end_poi_nti_p": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"nas_ipv4": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"psn_nam_e": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceMntSessionDisconnectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vENDPOINTIP := d.Get("end_poi_nti_p")
	vPSNNAME := d.Get("psn_nam_e")
	vMAC := d.Get("mac")
	vDISCONNECTTYPE := d.Get("dis_con_nec_tty_pe")
	vNASIPV4 := d.Get("nas_ipv4")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SessionDisconnect")
		vvENDPOINTIP := vENDPOINTIP.(string)
		vvPSNNAME := vPSNNAME.(string)
		vvMAC := vMAC.(string)
		vvDISCONNECTTYPE := vDISCONNECTTYPE.(string)
		vvNASIPV4 := vNASIPV4.(string)

		response1, err := client.Misc.SessionDisconnect(vvENDPOINTIP, vvPSNNAME, vvMAC, vvDISCONNECTTYPE, vvNASIPV4)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SessionDisconnect", err,
				"Failure at SessionDisconnect, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SessionDisconnect response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
