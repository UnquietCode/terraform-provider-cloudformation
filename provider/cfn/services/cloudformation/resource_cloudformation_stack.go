// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cloudformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCloudFormationStack() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFormationStackCreate,
		Read:   resourceCloudFormationStackRead,
		Update: resourceCloudFormationStackUpdate,
		Delete: resourceCloudFormationStackDelete,

		Schema: map[string]*schema.Schema{
			"notification_ar_ns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"template_url": {
				Type: schema.TypeString,
				Required: true,
			},
			"timeout_in_minutes": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}

func resourceCloudFormationStackCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFormation::Stack", data, meta)
}

func resourceCloudFormationStackRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFormation::Stack", data, meta)
}

func resourceCloudFormationStackUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFormation::Stack", data, meta)
}

func resourceCloudFormationStackDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFormation::Stack", data, meta)
}