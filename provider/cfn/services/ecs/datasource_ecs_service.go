// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html

package ecs

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eCSServiceType string = "AWS::ECS::Service"

func DatasourceECSService() *schema.Resource {
	return &schema.Resource{
		Read: datasourceECSServiceRead,
		
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
			"deployment_controller": {
				Type: schema.TypeList,
				Elem: propertyServiceDeploymentController(),
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
			"tags": misc.PropertyTags(),
			"task_definition": {
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

func datasourceECSServiceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eCSServiceType, DatasourceECSService(), data, meta)
}
