// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-script.html

package gamelift

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const gameLiftScriptType string = "AWS::GameLift::Script"

func DatasourceGameLiftScript() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGameLiftScriptRead,
		
		Schema: map[string]*schema.Schema{
			"version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"storage_location": {
				Type: schema.TypeList,
				Elem: propertyScriptS3Location(),
				Required: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceGameLiftScriptRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(gameLiftScriptType, DatasourceGameLiftScript(), data, meta)
}
