// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceECSServiceCreate,
		Read:   resourceECSServiceRead,
		Update: resourceECSServiceUpdate,
		Delete: resourceECSServiceDelete,

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"health_check_grace_period_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"launch_type": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"load_balancers": {
				Type: schema.TypeSet,
				Elem: propertyServiceLoadBalancer(),
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"placement_strategies": {
				Type: schema.TypeSet,
				Elem: propertyServicePlacementStrategy(),
				Optional: true,
				ForceNew: true,
			},
			"platform_version": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"propagate_tags": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"role": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scheduling_strategy": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"service_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"service_registries": {
				Type: schema.TypeSet,
				Elem: propertyServiceServiceRegistry(),
				Optional: true,
				ForceNew: true,
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
		},
	}
}

func resourceECSServiceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ECS::Service", data, meta)
}

func resourceECSServiceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ECS::Service", data, meta)
}

func resourceECSServiceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ECS::Service", data, meta)
}

func resourceECSServiceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ECS::Service", data, meta)
}