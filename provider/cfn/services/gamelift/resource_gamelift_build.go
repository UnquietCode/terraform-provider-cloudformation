// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceGameLiftBuildCreate,
		Read:   resourceGameLiftBuildRead,
		Update: resourceGameLiftBuildUpdate,
		Delete: resourceGameLiftBuildDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"storage_location": {
				Type: schema.TypeList,
				Elem: propertyBuildS3Location(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGameLiftBuildCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GameLift::Build", ResourceGameLiftBuild(), data, meta)
}

func resourceGameLiftBuildRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GameLift::Build", ResourceGameLiftBuild(), data, meta)
}

func resourceGameLiftBuildUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GameLift::Build", ResourceGameLiftBuild(), data, meta)
}

func resourceGameLiftBuildDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GameLift::Build", data, meta)
}