// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLambdaVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLambdaVersionExists,
		Read:   resourceLambdaVersionRead,
		Create: resourceLambdaVersionCreate,
		Update: resourceLambdaVersionUpdate,
		Delete: resourceLambdaVersionDelete,
		CustomizeDiff: resourceLambdaVersionCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"code_sha256": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"function_name": {
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

func resourceLambdaVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLambdaVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::Version", ResourceLambdaVersion(), data, meta)
}

func resourceLambdaVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::Version", ResourceLambdaVersion(), data, meta)
}

func resourceLambdaVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::Version", ResourceLambdaVersion(), data, meta)
}

func resourceLambdaVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::Version", data, meta)
}

func resourceLambdaVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}