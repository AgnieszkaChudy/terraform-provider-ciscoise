package ciscoise

import (
	"context"

	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSupportBundleDownload() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSupportBundleDownloadRead,
		Schema: map[string]*schema.Schema{
			"dirpath": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"file_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSupportBundleDownloadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DownloadSupportBundle")
		request1 := expandRequestSupportBundleDownloadDownloadSupportBundle(ctx, "", d)

		response1, _, err := client.SupportBundleDownload.DownloadSupportBundle(request1)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing DownloadSupportBundle", err))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response")

		vvDirpath := d.Get("dirpath").(string)
		err = response1.SaveDownload(vvDirpath)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when downloading file", err))
			return diags
		}
		log.Printf("[DEBUG] Downloaded file %s", vvDirpath)

	}
	return diags
}

func expandRequestSupportBundleDownloadDownloadSupportBundle(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSupportBundleDownloadDownloadSupportBundle {
	request := isegosdk.RequestSupportBundleDownloadDownloadSupportBundle{}
	request.ErsSupportBundleDownload = expandRequestSupportBundleDownloadDownloadSupportBundleErsSupportBundleDownload(ctx, key, d)
	return &request
}

func expandRequestSupportBundleDownloadDownloadSupportBundleErsSupportBundleDownload(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSupportBundleDownloadDownloadSupportBundleErsSupportBundleDownload {
	request := isegosdk.RequestSupportBundleDownloadDownloadSupportBundleErsSupportBundleDownload{}
	if v, ok := d.GetOkExists("file_name"); !isEmptyValue(reflect.ValueOf(d.Get("file_name"))) && (ok || !reflect.DeepEqual(v, d.Get("file_name"))) {
		request.FileName = interfaceToString(v)
	}
	return &request
}
