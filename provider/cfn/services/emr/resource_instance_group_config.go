// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceInstanceGroupConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceGroupConfigCreate,
		Read:   resourceInstanceGroupConfigRead,
		Update: resourceInstanceGroupConfigUpdate,
		Delete: resourceInstanceGroupConfigDelete,

		Schema: map[string]*schema.Schema{
			"auto_scaling_policy": {
				Type: schema.TypeList,
				Elem: propertyInstanceGroupConfigAutoScalingPolicy(),
				Required: false,
				MaxItems: 1,
			},
			"bid_price": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"configurations": {
				Type: schema.TypeSet,
				Elem: propertyInstanceGroupConfigConfiguration(),
				Required: false,
				ForceNew: true,
			},
			"ebs_configuration": {
				Type: schema.TypeList,
				Elem: propertyInstanceGroupConfigEbsConfiguration(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"instance_count": {
				Type: schema.TypeInt,
				Required: true,
			},
			"instance_role": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"job_flow_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"market": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceInstanceGroupConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::InstanceGroupConfig", data, meta)
}

func resourceInstanceGroupConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::InstanceGroupConfig", data, meta)
}

func resourceInstanceGroupConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::InstanceGroupConfig", data, meta)
}

func resourceInstanceGroupConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::InstanceGroupConfig", data, meta)
}