// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html

package gamelift

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const gameLiftMatchmakingConfigurationType string = "AWS::GameLift::MatchmakingConfiguration"

func DatasourceGameLiftMatchmakingConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGameLiftMatchmakingConfigurationRead,
		
		Schema: map[string]*schema.Schema{
			"game_properties": {
				Type: schema.TypeList,
				Elem: propertyMatchmakingConfigurationGameProperty(),
				Optional: true,
			},
			"game_session_data": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"acceptance_timeout_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"notification_target": {
				Type: schema.TypeString,
				Optional: true,
			},
			"custom_event_data": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"additional_player_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"backfill_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"request_timeout_seconds": {
				Type: schema.TypeInt,
				Required: true,
			},
			"acceptance_required": {
				Type: schema.TypeBool,
				Required: true,
			},
			"rule_set_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"game_session_queue_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
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

func datasourceGameLiftMatchmakingConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(gameLiftMatchmakingConfigurationType, DatasourceGameLiftMatchmakingConfiguration(), data, meta)
}
