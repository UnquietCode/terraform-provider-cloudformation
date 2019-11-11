// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html

package ecs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceECSService() *schema.Resource {
	return &schema.Resource{
		Exists: resourceECSServiceExists,
		Read: resourceECSServiceRead,
		Create: resourceECSServiceCreate,
		Update: resourceECSServiceUpdate,
		Delete: resourceECSServiceDelete,
		CustomizeDiff: resourceECSServiceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"cluster": {
				Type: schema.TypeString,
				Optional: true,
			},
			"deployment_configuration": {
				Type: schema.TypeList,
				Elem: propertyServiceDeploymentConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"desired_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"enable_ecs_managed_tags": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"health_check_grace_period_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"launch_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"load_balancers": {
				Type: schema.TypeSet,
				Elem: propertyServiceLoadBalancer(),
				Optional: true,
			},
			"network_configuration": {
				Type: schema.TypeList,
				Elem: propertyServiceNetworkConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"placement_constraints": {
				Type: schema.TypeSet,
				Elem: propertyServicePlacementConstraint(),
				Optional: true,
			},
			"placement_strategies": {
				Type: schema.TypeSet,
				Elem: propertyServicePlacementStrategy(),
				Optional: true,
			},
			"platform_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"propagate_tags": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scheduling_strategy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_registries": {
				Type: schema.TypeSet,
				Elem: propertyServiceServiceRegistry(),
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"task_definition": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceECSServiceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceECSServiceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ECS::Service", ResourceECSService(), data, meta)
}

func resourceECSServiceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ECS::Service", ResourceECSService(), data, meta)
}

func resourceECSServiceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ECS::Service", ResourceECSService(), data, meta)
}

func resourceECSServiceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ECS::Service", data, meta)
}

func resourceECSServiceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::ECS::Service", data, meta)
}

