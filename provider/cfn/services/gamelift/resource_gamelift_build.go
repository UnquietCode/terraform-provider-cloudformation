// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html

package gamelift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGameLiftBuild() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGameLiftBuildExists,
		Read: resourceGameLiftBuildRead,
		Create: resourceGameLiftBuildCreate,
		Update: resourceGameLiftBuildUpdate,
		Delete: resourceGameLiftBuildDelete,
		CustomizeDiff: resourceGameLiftBuildCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"name": {
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
			},
		},
	}
}

func resourceGameLiftBuildExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGameLiftBuildRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GameLift::Build", ResourceGameLiftBuild(), data, meta)
}

func resourceGameLiftBuildCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GameLift::Build", ResourceGameLiftBuild(), data, meta)
}

func resourceGameLiftBuildUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GameLift::Build", ResourceGameLiftBuild(), data, meta)
}

func resourceGameLiftBuildDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GameLift::Build", data, meta)
}

func resourceGameLiftBuildCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

