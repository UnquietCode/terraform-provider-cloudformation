// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package batch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceComputeEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeEnvironmentCreate,
		Read:   resourceComputeEnvironmentRead,
		Update: resourceComputeEnvironmentUpdate,
		Delete: resourceComputeEnvironmentDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"compute_environment_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"compute_resources": {
				Type: schema.TypeList,
				Elem: propertyComputeEnvironmentComputeResources(),
				Required: false,
				MaxItems: 1,
			},
			"state": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceComputeEnvironmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Batch::ComputeEnvironment", data, meta)
}

func resourceComputeEnvironmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Batch::ComputeEnvironment", data, meta)
}

func resourceComputeEnvironmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Batch::ComputeEnvironment", data, meta)
}

func resourceComputeEnvironmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Batch::ComputeEnvironment", data, meta)
}