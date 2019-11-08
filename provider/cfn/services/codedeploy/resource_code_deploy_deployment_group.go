// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package codedeploy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodeDeployDeploymentGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCodeDeployDeploymentGroupCreate,
		Read:   resourceCodeDeployDeploymentGroupRead,
		Update: resourceCodeDeployDeploymentGroupUpdate,
		Delete: resourceCodeDeployDeploymentGroupDelete,

		Schema: map[string]*schema.Schema{
			"alarm_configuration": {
				Type: schema.TypeList,
				Elem: propertyAlarmConfiguration(),
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
				Elem: propertyAutoRollbackConfiguration(),
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
				Elem: propertyDeployment(),
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
				Elem: propertyDeploymentStyle(),
				Required: false,
				MaxItems: 1,
			},
			"ec2_tag_filters": {
				Type: schema.TypeSet,
				Elem: propertyEC2TagFilter(),
				Required: false,
			},
			"ec2_tag_set": {
				Type: schema.TypeList,
				Elem: propertyEC2TagSet(),
				Required: false,
				MaxItems: 1,
			},
			"load_balancer_info": {
				Type: schema.TypeList,
				Elem: propertyLoadBalancerInfo(),
				Required: false,
				MaxItems: 1,
			},
			"on_premises_instance_tag_filters": {
				Type: schema.TypeSet,
				Elem: propertyTagFilter(),
				Required: false,
			},
			"on_premises_tag_set": {
				Type: schema.TypeList,
				Elem: propertyOnPremisesTagSet(),
				Required: false,
				MaxItems: 1,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"trigger_configurations": {
				Type: schema.TypeSet,
				Elem: propertyTriggerConfig(),
				Required: false,
			},
		},
	}
}

func resourceCodeDeployDeploymentGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeDeploy::DeploymentGroup", data, meta)
}

func resourceCodeDeployDeploymentGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodeDeploy::DeploymentGroup", data, meta)
}

func resourceCodeDeployDeploymentGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeDeploy::DeploymentGroup", data, meta)
}

func resourceCodeDeployDeploymentGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeDeploy::DeploymentGroup", data, meta)
}