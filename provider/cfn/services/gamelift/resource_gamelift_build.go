// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
				Required: false,
			},
			"storage_location": {
				Type: schema.TypeList,
				Elem: propertyBuildS3Location(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"version": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceGameLiftBuildCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GameLift::Build", data, meta)
}

func resourceGameLiftBuildRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GameLift::Build", data, meta)
}

func resourceGameLiftBuildUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GameLift::Build", data, meta)
}

func resourceGameLiftBuildDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GameLift::Build", data, meta)
}