// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogTagOptionType string = "AWS::ServiceCatalog::TagOption"

var serviceCatalogTagOptionProperties map[string]string = map[string]string{
	"active": "Active",
	"value": "Value",
	"key": "Key",
}

func ResourceServiceCatalogTagOption() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogTagOptionExists,
		Read: resourceServiceCatalogTagOptionRead,
		Create: resourceServiceCatalogTagOptionCreate,
		Update: resourceServiceCatalogTagOptionUpdate,
		Delete: resourceServiceCatalogTagOptionDelete,
		CustomizeDiff: resourceServiceCatalogTagOptionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"active": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"value": {
				Type: schema.TypeString,
				Required: true,
			},
			"key": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceCatalogTagOptionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogTagOptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogTagOptionType, ResourceServiceCatalogTagOption(), data, meta)
}

func resourceServiceCatalogTagOptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogTagOptionType, ResourceServiceCatalogTagOption(), data, serviceCatalogTagOptionProperties, meta)
}

func resourceServiceCatalogTagOptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogTagOptionType, ResourceServiceCatalogTagOption(), data, serviceCatalogTagOptionProperties, meta)
}

func resourceServiceCatalogTagOptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogTagOptionType, data, meta)
}

func resourceServiceCatalogTagOptionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogTagOptionType, data, meta)
}
