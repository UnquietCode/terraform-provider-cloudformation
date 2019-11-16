// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const eCSTaskDefinitionType string = "AWS::ECS::TaskDefinition"

var eCSTaskDefinitionProperties map[string]string = map[string]string{
	"container_definitions": "ContainerDefinitions",
	"cpu": "Cpu",
	"execution_role_arn": "ExecutionRoleArn",
	"family": "Family",
	"ipc_mode": "IpcMode",
	"memory": "Memory",
	"network_mode": "NetworkMode",
	"pid_mode": "PidMode",
	"placement_constraints": "PlacementConstraints",
	"proxy_configuration": "ProxyConfiguration",
	"requires_compatibilities": "RequiresCompatibilities",
	"tags": "Tags",
	"task_role_arn": "TaskRoleArn",
	"volumes": "Volumes",
}

func ResourceECSTaskDefinition() *schema.Resource {
	return &schema.Resource{
		Exists: resourceECSTaskDefinitionExists,
		Read: resourceECSTaskDefinitionRead,
		Create: resourceECSTaskDefinitionCreate,
		Update: resourceECSTaskDefinitionUpdate,
		Delete: resourceECSTaskDefinitionDelete,
		CustomizeDiff: resourceECSTaskDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"container_definitions": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionContainerDefinition(),
				Optional: true,
			},
			"cpu": {
				Type: schema.TypeString,
				Optional: true,
			},
			"execution_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"family": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ipc_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"memory": {
				Type: schema.TypeString,
				Optional: true,
			},
			"network_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"pid_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"placement_constraints": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionTaskDefinitionPlacementConstraint(),
				Optional: true,
			},
			"proxy_configuration": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionProxyConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"requires_compatibilities": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
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
			},
			"volumes": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionVolume(),
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

func resourceECSTaskDefinitionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceECSTaskDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eCSTaskDefinitionType, ResourceECSTaskDefinition(), data, meta)
}

func resourceECSTaskDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eCSTaskDefinitionType, ResourceECSTaskDefinition(), data, eCSTaskDefinitionProperties, meta)
}

func resourceECSTaskDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eCSTaskDefinitionType, ResourceECSTaskDefinition(), data, eCSTaskDefinitionProperties, meta)
}

func resourceECSTaskDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eCSTaskDefinitionType, data, meta)
}

func resourceECSTaskDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eCSTaskDefinitionType, data, meta)
}
