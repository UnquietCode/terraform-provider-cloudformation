// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html

package ecs

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eCSTaskSetType string = "AWS::ECS::TaskSet"

func DatasourceECSTaskSet() *schema.Resource {
	return &schema.Resource{
		Read: datasourceECSTaskSetRead,
		
		Schema: map[string]*schema.Schema{
			"cluster": {
				Type: schema.TypeString,
				Required: true,
			},
			"external_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"launch_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"load_balancers": {
				Type: schema.TypeList,
				Elem: propertyTaskSetLoadBalancer(),
				Optional: true,
			},
			"network_configuration": {
				Type: schema.TypeList,
				Elem: propertyTaskSetNetworkConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"platform_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scale": {
				Type: schema.TypeList,
				Elem: propertyTaskSetScale(),
				Optional: true,
				MaxItems: 1,
			},
			"service": {
				Type: schema.TypeString,
				Required: true,
			},
			"service_registries": {
				Type: schema.TypeList,
				Elem: propertyTaskSetServiceRegistry(),
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceECSTaskSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eCSTaskSetType, DatasourceECSTaskSet(), data, meta)
}
