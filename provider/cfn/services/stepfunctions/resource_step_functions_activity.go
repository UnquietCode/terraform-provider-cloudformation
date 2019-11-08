// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package stepfunctions

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceStepFunctionsActivity() *schema.Resource {
	return &schema.Resource{
		Create: resourceStepFunctionsActivityCreate,
		Read:   resourceStepFunctionsActivityRead,
		Update: resourceStepFunctionsActivityUpdate,
		Delete: resourceStepFunctionsActivityDelete,

		Schema: map[string]*schema.Schema{
			"tags": {
				Type: schema.TypeList,
				Elem: propertyTagsEntry(),
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

func resourceStepFunctionsActivityCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::StepFunctions::Activity", data, meta)
}

func resourceStepFunctionsActivityRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::StepFunctions::Activity", data, meta)
}

func resourceStepFunctionsActivityUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::StepFunctions::Activity", data, meta)
}

func resourceStepFunctionsActivityDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::StepFunctions::Activity", data, meta)
}