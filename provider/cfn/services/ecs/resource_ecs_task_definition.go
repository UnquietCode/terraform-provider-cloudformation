// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html

package ecs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceECSTaskDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceECSTaskDefinitionCreate,
		Read:   resourceECSTaskDefinitionRead,
		Update: resourceECSTaskDefinitionUpdate,
		Delete: resourceECSTaskDefinitionDelete,

		Schema: map[string]*schema.Schema{
			"container_definitions": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionContainerDefinition(),
				Required: false,
				ForceNew: true,
			},
			"cpu": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"execution_role_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"family": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ipc_mode": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"memory": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"network_mode": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"pid_mode": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"placement_constraints": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionTaskDefinitionPlacementConstraint(),
				Required: false,
				ForceNew: true,
			},
			"proxy_configuration": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionProxyConfiguration(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"requires_compatibilities": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"task_role_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"volumes": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionVolume(),
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceECSTaskDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ECS::TaskDefinition", data, meta)
}

func resourceECSTaskDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ECS::TaskDefinition", data, meta)
}

func resourceECSTaskDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ECS::TaskDefinition", data, meta)
}

func resourceECSTaskDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ECS::TaskDefinition", data, meta)
}