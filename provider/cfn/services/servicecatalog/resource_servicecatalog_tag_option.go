// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogTagOption() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCatalogTagOptionCreate,
		Read:   resourceServiceCatalogTagOptionRead,
		Update: resourceServiceCatalogTagOptionUpdate,
		Delete: resourceServiceCatalogTagOptionDelete,

		Schema: map[string]*schema.Schema{
			"active": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"value": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceCatalogTagOptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::TagOption", ResourceServiceCatalogTagOption(), data, meta)
}

func resourceServiceCatalogTagOptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::TagOption", ResourceServiceCatalogTagOption(), data, meta)
}

func resourceServiceCatalogTagOptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::TagOption", ResourceServiceCatalogTagOption(), data, meta)
}

func resourceServiceCatalogTagOptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::TagOption", data, meta)
}