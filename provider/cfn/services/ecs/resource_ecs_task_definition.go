// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
				Optional: true,
				ForceNew: true,
			},
			"cpu": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"execution_role_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"family": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ipc_mode": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"memory": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"network_mode": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"pid_mode": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"placement_constraints": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionTaskDefinitionPlacementConstraint(),
				Optional: true,
				ForceNew: true,
			},
			"proxy_configuration": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionProxyConfiguration(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"requires_compatibilities": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
				Set: schema.HashString,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"task_role_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"volumes": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionVolume(),
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceECSTaskDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ECS::TaskDefinition", ResourceECSTaskDefinition(), data, meta)
}

func resourceECSTaskDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ECS::TaskDefinition", ResourceECSTaskDefinition(), data, meta)
}

func resourceECSTaskDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ECS::TaskDefinition", ResourceECSTaskDefinition(), data, meta)
}

func resourceECSTaskDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ECS::TaskDefinition", data, meta)
}