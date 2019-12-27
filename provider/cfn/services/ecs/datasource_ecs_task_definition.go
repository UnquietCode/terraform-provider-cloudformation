// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html

package ecs

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eCSTaskDefinitionType string = "AWS::ECS::TaskDefinition"

func DatasourceECSTaskDefinition() *schema.Resource {
	return &schema.Resource{
		Read: datasourceECSTaskDefinitionRead,
		
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
			"inference_accelerators": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionInferenceAccelerator(),
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
			"tags": misc.PropertyTags(),
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceECSTaskDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eCSTaskDefinitionType, DatasourceECSTaskDefinition(), data, meta)
}
