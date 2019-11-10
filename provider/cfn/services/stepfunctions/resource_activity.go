// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package stepfunctions

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceActivity() *schema.Resource {
	return &schema.Resource{
		Create: resourceActivityCreate,
		Read:   resourceActivityRead,
		Update: resourceActivityUpdate,
		Delete: resourceActivityDelete,

		Schema: map[string]*schema.Schema{
			"tags": {
				Type: schema.TypeList,
				Elem: propertyActivityTagsEntry(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceActivityCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::StepFunctions::Activity", data, meta)
}

func resourceActivityRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::StepFunctions::Activity", data, meta)
}

func resourceActivityUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::StepFunctions::Activity", data, meta)
}

func resourceActivityDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::StepFunctions::Activity", data, meta)
}