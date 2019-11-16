// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html

package gamelift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const gameLiftAliasType string = "AWS::GameLift::Alias"

func ResourceGameLiftAlias() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGameLiftAliasExists,
		Read: resourceGameLiftAliasRead,
		Create: resourceGameLiftAliasCreate,
		Update: resourceGameLiftAliasUpdate,
		Delete: resourceGameLiftAliasDelete,
		CustomizeDiff: resourceGameLiftAliasCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"routing_strategy": {
				Type: schema.TypeSet,
				Elem: propertyAliasRoutingStrategy(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGameLiftAliasExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGameLiftAliasRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(gameLiftAliasType, ResourceGameLiftAlias(), data, meta)
}

func resourceGameLiftAliasCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(gameLiftAliasType, ResourceGameLiftAlias(), data, meta)
}

func resourceGameLiftAliasUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(gameLiftAliasType, ResourceGameLiftAlias(), data, meta)
}

func resourceGameLiftAliasDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(gameLiftAliasType, data, meta)
}

func resourceGameLiftAliasCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(gameLiftAliasType, data, meta)
}
