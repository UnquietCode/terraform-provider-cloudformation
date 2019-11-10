// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codedeploy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDeploymentGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentGroupCreate,
		Read:   resourceDeploymentGroupRead,
		Update: resourceDeploymentGroupUpdate,
		Delete: resourceDeploymentGroupDelete,

		Schema: map[string]*schema.Schema{
			"alarm_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupAlarmConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"application_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auto_rollback_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupAutoRollbackConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"auto_scaling_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"deployment": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupDeployment(),
				Required: false,
				MaxItems: 1,
			},
			"deployment_config_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"deployment_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"deployment_style": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupDeploymentStyle(),
				Required: false,
				MaxItems: 1,
			},
			"ec2_tag_filters": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentGroupEC2TagFilter(),
				Required: false,
			},
			"ec2_tag_set": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupEC2TagSet(),
				Required: false,
				MaxItems: 1,
			},
			"load_balancer_info": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupLoadBalancerInfo(),
				Required: false,
				MaxItems: 1,
			},
			"on_premises_instance_tag_filters": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentGroupTagFilter(),
				Required: false,
			},
			"on_premises_tag_set": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupOnPremisesTagSet(),
				Required: false,
				MaxItems: 1,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"trigger_configurations": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentGroupTriggerConfig(),
				Required: false,
			},
		},
	}
}

func resourceDeploymentGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeDeploy::DeploymentGroup", data, meta)
}

func resourceDeploymentGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodeDeploy::DeploymentGroup", data, meta)
}

func resourceDeploymentGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeDeploy::DeploymentGroup", data, meta)
}

func resourceDeploymentGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeDeploy::DeploymentGroup", data, meta)
}