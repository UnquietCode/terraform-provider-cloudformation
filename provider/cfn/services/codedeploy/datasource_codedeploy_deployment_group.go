// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html

package codedeploy

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeDeployDeploymentGroupType string = "AWS::CodeDeploy::DeploymentGroup"

func DatasourceCodeDeployDeploymentGroup() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCodeDeployDeploymentGroupRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceCodeDeployDeploymentGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeDeployDeploymentGroupType, DatasourceCodeDeployDeploymentGroup(), data, meta)
}
