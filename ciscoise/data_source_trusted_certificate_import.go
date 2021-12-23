package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceTrustedCertificateImport() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Certificates.

- Import an X509 certificate as a trust certificate.

NOTE:
Request parameters accepting True and False as input can be replaced by 1 and 0 respectively.

`,

		ReadContext: dataSourceTrustedCertificateImportRead,
		Schema: map[string]*schema.Schema{
			"allow_basic_constraint_cafalse": &schema.Schema{
				Description:  `Allow certificates with Basic Constraints CA Field as False (required)`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"allow_out_of_date_cert": &schema.Schema{
				Description:  `Allow out of date certificates (required)`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"allow_sha1_certificates": &schema.Schema{
				Description:  `Allow SHA1 based certificates (required)`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"data": &schema.Schema{
				Description: `Certificate content (required)`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": &schema.Schema{
				Description: `Description of the certificate`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `ID of the imported trust certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Response message on import of system or trust certificate`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status": &schema.Schema{
							Description: `HTTP response status after import`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Description: `Name of the certificate`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"trust_for_certificate_based_admin_auth": &schema.Schema{
				Description:  `Trust for Certificate based Admin authentication`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"trust_for_cisco_services_auth": &schema.Schema{
				Description:  `Trust for authentication of Cisco Services`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"trust_for_client_auth": &schema.Schema{
				Description:  `Trust for client authentication and Syslog`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"trust_for_ise_auth": &schema.Schema{
				Description:  `Trust for authentication within Cisco ISE`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"validate_certificate_extensions": &schema.Schema{
				Description:  `Validate trust certificate extension`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
		},
	}
}

func dataSourceTrustedCertificateImportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ImportTrustCert")
		request1 := expandRequestTrustedCertificateImportImportTrustCert(ctx, "", d)

		response1, restyResp1, err := client.Certificates.ImportTrustCert(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportTrustCert", err,
				"Failure at ImportTrustCert, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenCertificatesImportTrustCertItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportTrustCert response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestTrustedCertificateImportImportTrustCert(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesImportTrustCert {
	request := isegosdk.RequestCertificatesImportTrustCert{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_basic_constraint_cafalse")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_basic_constraint_cafalse")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_basic_constraint_cafalse")))) {
		request.AllowBasicConstraintCaFalse = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_out_of_date_cert")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_out_of_date_cert")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_out_of_date_cert")))) {
		request.AllowOutOfDateCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_sha1_certificates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_sha1_certificates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_sha1_certificates")))) {
		request.AllowSHA1Certificates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trust_for_certificate_based_admin_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trust_for_certificate_based_admin_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trust_for_certificate_based_admin_auth")))) {
		request.TrustForCertificateBasedAdminAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trust_for_cisco_services_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trust_for_cisco_services_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trust_for_cisco_services_auth")))) {
		request.TrustForCiscoServicesAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trust_for_client_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trust_for_client_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trust_for_client_auth")))) {
		request.TrustForClientAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trust_for_ise_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trust_for_ise_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trust_for_ise_auth")))) {
		request.TrustForIseAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".validate_certificate_extensions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".validate_certificate_extensions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".validate_certificate_extensions")))) {
		request.ValidateCertificateExtensions = interfaceToBoolPtr(v)
	}
	return &request
}

func flattenCertificatesImportTrustCertItem(item *isegosdk.ResponseCertificatesImportTrustCertResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["message"] = item.Message
	respItem["status"] = item.Status
	return []map[string]interface{}{
		respItem,
	}
}
