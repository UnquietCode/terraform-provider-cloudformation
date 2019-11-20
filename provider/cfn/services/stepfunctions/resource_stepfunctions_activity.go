// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html

package stepfunctions

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const stepFunctionsActivityType string = "AWS::StepFunctions::Activity"

func ResourceStepFunctionsActivity() *schema.Resource {
	return &schema.Resource{
		Read: resourceStepFunctionsActivityRead,
		Create: resourceStepFunctionsActivityCreate,
		Update: resourceStepFunctionsActivityUpdate,
		Delete: resourceStepFunctionsActivityDelete,
		CustomizeDiff: resourceStepFunctionsActivityCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"tags": {
				Type: schema.TypeList,
				Elem: propertyActivityTagsEntry(),
				Optional: true,
			},
			"name": {
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

func resourceStepFunctionsActivityRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(stepFunctionsActivityType, ResourceStepFunctionsActivity(), data, meta)
}

func resourceStepFunctionsActivityCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(stepFunctionsActivityType, ResourceStepFunctionsActivity(), data, meta)
}

func resourceStepFunctionsActivityUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(stepFunctionsActivityType, ResourceStepFunctionsActivity(), data, meta)
}

func resourceStepFunctionsActivityDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(stepFunctionsActivityType, data, meta)
}

func resourceStepFunctionsActivityCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(stepFunctionsActivityType, data, meta)
}
