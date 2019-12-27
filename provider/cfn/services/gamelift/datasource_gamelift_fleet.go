// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html

package gamelift

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const gameLiftFleetType string = "AWS::GameLift::Fleet"

func DatasourceGameLiftFleet() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGameLiftFleetRead,
		
		Schema: map[string]*schema.Schema{
			"build_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"certificate_configuration": {
				Type: schema.TypeList,
				Elem: propertyFleetCertificateConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"desired_ec2_instances": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"ec2_inbound_permissions": {
				Type: schema.TypeSet,
				Elem: propertyFleetIpPermission(),
				Optional: true,
			},
			"ec2_instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"fleet_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"log_paths": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"max_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"metric_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"min_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"new_game_session_protection_policy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"peer_vpc_aws_account_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"peer_vpc_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"resource_creation_limit_policy": {
				Type: schema.TypeList,
				Elem: propertyFleetResourceCreationLimitPolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"runtime_configuration": {
				Type: schema.TypeList,
				Elem: propertyFleetRuntimeConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"script_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"server_launch_parameters": {
				Type: schema.TypeString,
				Optional: true,
			},
			"server_launch_path": {
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

func datasourceGameLiftFleetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(gameLiftFleetType, DatasourceGameLiftFleet(), data, meta)
}
