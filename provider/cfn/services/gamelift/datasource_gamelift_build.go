// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html

package gamelift

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const gameLiftBuildType string = "AWS::GameLift::Build"

func DatasourceGameLiftBuild() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGameLiftBuildRead,
		
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"operating_system": {
				Type: schema.TypeString,
				Optional: true,
			},
			"storage_location": {
				Type: schema.TypeList,
				Elem: propertyBuildS3Location(),
				Optional: true,
				MaxItems: 1,
			},
			"version": {
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

func datasourceGameLiftBuildRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(gameLiftBuildType, DatasourceGameLiftBuild(), data, meta)
}
