// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceEMRInstanceGroupConfigExists,
		Read:   resourceEMRInstanceGroupConfigRead,
		Create: resourceEMRInstanceGroupConfigCreate,
		Update: resourceEMRInstanceGroupConfigUpdate,
		Delete: resourceEMRInstanceGroupConfigDelete,
		CustomizeDiff: resourceEMRInstanceGroupConfigCustomizeDiff,

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
			},
			"configurations": {
				Type: schema.TypeSet,
				Elem: propertyInstanceGroupConfigConfiguration(),
				Optional: true,
			},
			"ebs_configuration": {
				Type: schema.TypeList,
				Elem: propertyInstanceGroupConfigEbsConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"instance_count": {
				Type: schema.TypeInt,
				Required: true,
			},
			"instance_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"job_flow_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"market": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
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

func resourceEMRInstanceGroupConfigExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEMRInstanceGroupConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::InstanceGroupConfig", ResourceEMRInstanceGroupConfig(), data, meta)
}

func resourceEMRInstanceGroupConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::InstanceGroupConfig", ResourceEMRInstanceGroupConfig(), data, meta)
}

func resourceEMRInstanceGroupConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::InstanceGroupConfig", ResourceEMRInstanceGroupConfig(), data, meta)
}

func resourceEMRInstanceGroupConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::InstanceGroupConfig", data, meta)
}

func resourceEMRInstanceGroupConfigCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}