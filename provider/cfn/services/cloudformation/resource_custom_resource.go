// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cloudformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCustomResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCustomResourceCreate,
		Read:   resourceCustomResourceRead,
		Update: resourceCustomResourceUpdate,
		Delete: resourceCustomResourceDelete,

		Schema: map[string]*schema.Schema{
			"service_token": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCustomResourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFormation::CustomResource", data, meta)
}

func resourceCustomResourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFormation::CustomResource", data, meta)
}

func resourceCustomResourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFormation::CustomResource", data, meta)
}

func resourceCustomResourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFormation::CustomResource", data, meta)
}