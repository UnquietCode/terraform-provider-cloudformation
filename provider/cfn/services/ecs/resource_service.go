// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ecs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCreate,
		Read:   resourceServiceRead,
		Update: resourceServiceUpdate,
		Delete: resourceServiceDelete,

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"deployment_configuration": {
				Type: schema.TypeList,
				Elem: propertyServiceDeploymentConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"desired_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"enable_ecs_managed_tags": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"health_check_grace_period_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"launch_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"load_balancers": {
				Type: schema.TypeSet,
				Elem: propertyServiceLoadBalancer(),
				Required: false,
				ForceNew: true,
			},
			"network_configuration": {
				Type: schema.TypeList,
				Elem: propertyServiceNetworkConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"placement_constraints": {
				Type: schema.TypeSet,
				Elem: propertyServicePlacementConstraint(),
				Required: false,
				ForceNew: true,
			},
			"placement_strategies": {
				Type: schema.TypeSet,
				Elem: propertyServicePlacementStrategy(),
				Required: false,
				ForceNew: true,
			},
			"platform_version": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"propagate_tags": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"role": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"scheduling_strategy": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"service_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"service_registries": {
				Type: schema.TypeSet,
				Elem: propertyServiceServiceRegistry(),
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"task_definition": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ECS::Service", data, meta)
}

func resourceServiceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ECS::Service", data, meta)
}

func resourceServiceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ECS::Service", data, meta)
}

func resourceServiceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ECS::Service", data, meta)
}