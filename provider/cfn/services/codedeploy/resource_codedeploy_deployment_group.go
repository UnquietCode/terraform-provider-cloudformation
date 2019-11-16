// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html

package codedeploy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeDeployDeploymentGroupType string = "AWS::CodeDeploy::DeploymentGroup"

var codeDeployDeploymentGroupProperties map[string]string = map[string]string{
	"alarm_configuration": "AlarmConfiguration",
	"application_name": "ApplicationName",
	"auto_rollback_configuration": "AutoRollbackConfiguration",
	"auto_scaling_groups": "AutoScalingGroups",
	"deployment": "Deployment",
	"deployment_config_name": "DeploymentConfigName",
	"deployment_group_name": "DeploymentGroupName",
	"deployment_style": "DeploymentStyle",
	"ec2_tag_filters": "Ec2TagFilters",
	"ec2_tag_set": "Ec2TagSet",
	"load_balancer_info": "LoadBalancerInfo",
	"on_premises_instance_tag_filters": "OnPremisesInstanceTagFilters",
	"on_premises_tag_set": "OnPremisesTagSet",
	"service_role_arn": "ServiceRoleArn",
	"trigger_configurations": "TriggerConfigurations",
}

func ResourceCodeDeployDeploymentGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCodeDeployDeploymentGroupExists,
		Read: resourceCodeDeployDeploymentGroupRead,
		Create: resourceCodeDeployDeploymentGroupCreate,
		Update: resourceCodeDeployDeploymentGroupUpdate,
		Delete: resourceCodeDeployDeploymentGroupDelete,
		CustomizeDiff: resourceCodeDeployDeploymentGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"alarm_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupAlarmConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"auto_rollback_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupAutoRollbackConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"auto_scaling_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"deployment": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupDeployment(),
				Optional: true,
				MaxItems: 1,
			},
			"deployment_config_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"deployment_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"deployment_style": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupDeploymentStyle(),
				Optional: true,
				MaxItems: 1,
			},
			"ec2_tag_filters": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentGroupEC2TagFilter(),
				Optional: true,
			},
			"ec2_tag_set": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupEC2TagSet(),
				Optional: true,
				MaxItems: 1,
			},
			"load_balancer_info": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupLoadBalancerInfo(),
				Optional: true,
				MaxItems: 1,
			},
			"on_premises_instance_tag_filters": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentGroupTagFilter(),
				Optional: true,
			},
			"on_premises_tag_set": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupOnPremisesTagSet(),
				Optional: true,
				MaxItems: 1,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"trigger_configurations": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentGroupTriggerConfig(),
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

func resourceCodeDeployDeploymentGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCodeDeployDeploymentGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeDeployDeploymentGroupType, ResourceCodeDeployDeploymentGroup(), data, meta)
}

func resourceCodeDeployDeploymentGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codeDeployDeploymentGroupType, ResourceCodeDeployDeploymentGroup(), data, codeDeployDeploymentGroupProperties, meta)
}

func resourceCodeDeployDeploymentGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codeDeployDeploymentGroupType, ResourceCodeDeployDeploymentGroup(), data, codeDeployDeploymentGroupProperties, meta)
}

func resourceCodeDeployDeploymentGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codeDeployDeploymentGroupType, data, meta)
}

func resourceCodeDeployDeploymentGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codeDeployDeploymentGroupType, data, meta)
}
