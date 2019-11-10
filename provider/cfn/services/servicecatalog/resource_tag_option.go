// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTagOption() *schema.Resource {
	return &schema.Resource{
		Create: resourceTagOptionCreate,
		Read:   resourceTagOptionRead,
		Update: resourceTagOptionUpdate,
		Delete: resourceTagOptionDelete,

		Schema: map[string]*schema.Schema{
			"active": {
				Type: schema.TypeBool,
				Required: false,
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

func resourceTagOptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::TagOption", data, meta)
}

func resourceTagOptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::TagOption", data, meta)
}

func resourceTagOptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::TagOption", data, meta)
}

func resourceTagOptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::TagOption", data, meta)
}