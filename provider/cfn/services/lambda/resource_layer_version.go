// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLayerVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceLayerVersionCreate,
		Read:   resourceLayerVersionRead,
		Update: resourceLayerVersionUpdate,
		Delete: resourceLayerVersionDelete,

		Schema: map[string]*schema.Schema{
			"compatible_runtimes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"license_info": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"layer_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"content": {
				Type: schema.TypeList,
				Elem: propertyLayerVersionContent(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceLayerVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::LayerVersion", data, meta)
}

func resourceLayerVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::LayerVersion", data, meta)
}

func resourceLayerVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::LayerVersion", data, meta)
}

func resourceLayerVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::LayerVersion", data, meta)
}