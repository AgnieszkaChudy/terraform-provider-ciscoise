package ciscoise

import (
	"context"

	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceCsrGenerate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCsrGenerateRead,
		Schema: map[string]*schema.Schema{
			"allow_wild_card_cert": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"certificate_policies": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"digest_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostnames": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"key_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"portal_group_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"san_dns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"san_dir": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"san_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"san_uri": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subject_city": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_common_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_org": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_org_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"used_for": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceCsrGenerateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 2: GenerateCsr")
		request2 := expandRequestCsrGenerateGenerateCsr(ctx, "", d)

		response2, _, err := client.Certificates.GenerateCsr(request2)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GenerateCsr", err,
				"Failure at GenerateCsr, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItems2 := flattenCertificatesGenerateCsrItems(response2.Response)
		if err := d.Set("items", vItems2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GenerateCsr response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestCsrGenerateGenerateCsr(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesGenerateCsr {
	request := isegosdk.RequestCertificatesGenerateCsr{}
	if v, ok := d.GetOkExists(key + ".allow_wild_card_cert"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_wild_card_cert"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_wild_card_cert"))) {
		request.AllowWildCardCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".certificate_policies"); !isEmptyValue(reflect.ValueOf(d.Get(key+".certificate_policies"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".certificate_policies"))) {
		request.CertificatePolicies = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".digest_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".digest_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".digest_type"))) {
		request.DigestType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".hostnames"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hostnames"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hostnames"))) {
		request.Hostnames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".key_length"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key_length"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key_length"))) {
		request.KeyLength = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".key_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key_type"))) {
		request.KeyType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_group_tag"))) {
		request.PortalGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".san_dns"); !isEmptyValue(reflect.ValueOf(d.Get(key+".san_dns"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".san_dns"))) {
		request.SanDNS = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".san_dir"); !isEmptyValue(reflect.ValueOf(d.Get(key+".san_dir"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".san_dir"))) {
		request.SanDir = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".san_ip"); !isEmptyValue(reflect.ValueOf(d.Get(key+".san_ip"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".san_ip"))) {
		request.SanIP = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".san_uri"); !isEmptyValue(reflect.ValueOf(d.Get(key+".san_uri"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".san_uri"))) {
		request.SanURI = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".subject_city"); !isEmptyValue(reflect.ValueOf(d.Get(key+".subject_city"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".subject_city"))) {
		request.SubjectCity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".subject_common_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".subject_common_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".subject_common_name"))) {
		request.SubjectCommonName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".subject_country"); !isEmptyValue(reflect.ValueOf(d.Get(key+".subject_country"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".subject_country"))) {
		request.SubjectCountry = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".subject_org"); !isEmptyValue(reflect.ValueOf(d.Get(key+".subject_org"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".subject_org"))) {
		request.SubjectOrg = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".subject_org_unit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".subject_org_unit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".subject_org_unit"))) {
		request.SubjectOrgUnit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".subject_state"); !isEmptyValue(reflect.ValueOf(d.Get(key+".subject_state"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".subject_state"))) {
		request.SubjectState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".used_for"); !isEmptyValue(reflect.ValueOf(d.Get(key+".used_for"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".used_for"))) {
		request.UsedFor = interfaceToString(v)
	}
	return &request
}

func flattenCertificatesGenerateCsrItems(items *[]isegosdk.ResponseCertificatesGenerateCsrResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["link"] = flattenCertificatesGenerateCsrItemsLink(item.Link)
		respItem["message"] = item.Message
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenCertificatesGenerateCsrItemsLink(item *isegosdk.ResponseCertificatesGenerateCsrResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
