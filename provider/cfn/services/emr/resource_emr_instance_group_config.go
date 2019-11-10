// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEMRInstanceGroupConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceEMRInstanceGroupConfigCreate,
		Read:   resourceEMRInstanceGroupConfigRead,
		Update: resourceEMRInstanceGroupConfigUpdate,
		Delete: resourceEMRInstanceGroupConfigDelete,

		Schema: map[string]*schema.Schema{
			"auto_scaling_policy": {
				Type: schema.TypeList,
				Elem: propertyInstanceGroupConfigAutoScalingPolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"bid_price": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"configurations": {
				Type: schema.TypeSet,
				Elem: propertyInstanceGroupConfigConfiguration(),
				Optional: true,
				ForceNew: true,
			},
			"ebs_configuration": {
				Type: schema.TypeList,
				Elem: propertyInstanceGroupConfigEbsConfiguration(),
				Optional: true,
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
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEMRInstanceGroupConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::InstanceGroupConfig", data, meta)
}

func resourceEMRInstanceGroupConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::InstanceGroupConfig", data, meta)
}

func resourceEMRInstanceGroupConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::InstanceGroupConfig", data, meta)
}

func resourceEMRInstanceGroupConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::InstanceGroupConfig", data, meta)
}