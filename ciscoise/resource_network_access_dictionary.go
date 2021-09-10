package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkAccessDictionary() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceNetworkAccessDictionaryCreate,
		ReadContext:   resourceNetworkAccessDictionaryRead,
		UpdateContext: resourceNetworkAccessDictionaryUpdate,
		DeleteContext: resourceNetworkAccessDictionaryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dictionary_attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"href": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:             schema.TypeList,
							DiffSuppressFunc: diffSuppressAlways(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
									"rel": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
									"type": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkAccessDictionaryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestNetworkAccessDictionaryCreateNetworkAccessDictionaries(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	vID, _ := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okName && vvName != "" {
		getResponse2, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["name"] = vvName
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaries()
		if response2 != nil && err == nil {
			items2 := getAllItemsNetworkAccessDictionaryGetNetworkAccessDictionaries(m, response2)
			item2, err := searchNetworkAccessDictionaryGetNetworkAccessDictionaries(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["name"] = vvName
				resourceMap["id"] = vvID
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	resp1, restyResp1, err := client.NetworkAccessDictionary.CreateNetworkAccessDictionaries(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkAccessDictionaries", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkAccessDictionaries", err))
		return diags
	}
	if vvID != resp1.Response.ID {
		vvID = resp1.Response.ID
	}
	if vvName != resp1.Response.Name {
		vvName = resp1.Response.Name
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNetworkAccessDictionaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		vvName := vName
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetNetworkAccessDictionaries")

		response1, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaries()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionaries", err,
				"Failure at GetNetworkAccessDictionaries, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsNetworkAccessDictionaryGetNetworkAccessDictionaries(m, response1)
		item1, err := searchNetworkAccessDictionaryGetNetworkAccessDictionaries(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetNetworkAccessDictionaries response", err,
				"Failure when searching item from GetNetworkAccessDictionaries, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaries search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessDictionaryByName")
		vvName := vName

		response2, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionaryByName", err,
				"Failure at GetNetworkAccessDictionaryByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkAccessDictionaryGetNetworkAccessDictionaryByNameItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaryByName response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkAccessDictionaryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		getResp1, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaries()
		if err == nil && getResp1 != nil {
			items1 := getAllItemsNetworkAccessDictionaryGetNetworkAccessDictionaries(m, getResp1)
			item1, err := searchNetworkAccessDictionaryGetNetworkAccessDictionaries(m, items1, vName, vID)
			if err == nil && item1 != nil {
				if vName != item1.Name {
					vvName = item1.Name
				} else {
					vvName = vName
				}
			}
		}
	}
	if selectedMethod == 1 {
		vvName = vName
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvName %s", vvName)
		request1 := expandRequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkAccessDictionary.UpdateNetworkAccessDictionaryByName(vvName, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkAccessDictionaryByName", err, restyResp1.String(),
					"Failure at UpdateNetworkAccessDictionaryByName, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkAccessDictionaryByName", err,
				"Failure at UpdateNetworkAccessDictionaryByName, unexpected response", ""))
			return diags
		}
	}

	return resourceNetworkAccessDictionaryRead(ctx, d, m)
}

func resourceNetworkAccessDictionaryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {

		getResp1, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaries()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkAccessDictionaryGetNetworkAccessDictionaries(m, getResp1)
		item1, err := searchNetworkAccessDictionaryGetNetworkAccessDictionaries(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vName != item1.Name {
			vvName = item1.Name
		} else {
			vvName = vName
		}
	}
	if selectedMethod == 1 {
		vvName = vName
		getResp, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NetworkAccessDictionary.DeleteNetworkAccessDictionaryByName(vvName)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkAccessDictionaryByName", err, restyResp1.String(),
				"Failure at DeleteNetworkAccessDictionaryByName, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkAccessDictionaryByName", err,
			"Failure at DeleteNetworkAccessDictionaryByName, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkAccessDictionaryCreateNetworkAccessDictionaries(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryCreateNetworkAccessDictionaries {
	request := isegosdk.RequestNetworkAccessDictionaryCreateNetworkAccessDictionaries{}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_attr_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_attr_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_attr_type"))) {
		request.DictionaryAttrType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessDictionaryCreateNetworkAccessDictionariesLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".version"); !isEmptyValue(reflect.ValueOf(d.Get(key+".version"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".version"))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryCreateNetworkAccessDictionariesLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryCreateNetworkAccessDictionariesLink {
	request := isegosdk.RequestNetworkAccessDictionaryCreateNetworkAccessDictionariesLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName {
	request := isegosdk.RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName{}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_attr_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_attr_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_attr_type"))) {
		request.DictionaryAttrType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".version"); !isEmptyValue(reflect.ValueOf(d.Get(key+".version"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".version"))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameLink {
	request := isegosdk.RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNetworkAccessDictionaryGetNetworkAccessDictionaries(m interface{}, response *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaries) []isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponse {
	var respItems []isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNetworkAccessDictionaryGetNetworkAccessDictionaries(m interface{}, items []isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponse, name string, id string) (*isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByNameResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByNameResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByName
			getItem, _, err = client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(item.Name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessDictionaryByName")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByName
			getItem, _, err = client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessDictionaryByName")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
