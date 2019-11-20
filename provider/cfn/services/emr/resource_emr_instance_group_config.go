// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eMRInstanceGroupConfigType string = "AWS::EMR::InstanceGroupConfig"

func ResourceEMRInstanceGroupConfig() *schema.Resource {
	return &schema.Resource{
		Read: resourceEMRInstanceGroupConfigRead,
		Create: resourceEMRInstanceGroupConfigCreate,
		Update: resourceEMRInstanceGroupConfigUpdate,
		Delete: resourceEMRInstanceGroupConfigDelete,
		CustomizeDiff: resourceEMRInstanceGroupConfigCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"auto_scaling_policy": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
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

func resourceEMRInstanceGroupConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eMRInstanceGroupConfigType, ResourceEMRInstanceGroupConfig(), data, meta)
}

func resourceEMRInstanceGroupConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eMRInstanceGroupConfigType, ResourceEMRInstanceGroupConfig(), data, meta)
}

func resourceEMRInstanceGroupConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eMRInstanceGroupConfigType, ResourceEMRInstanceGroupConfig(), data, meta)
}

func resourceEMRInstanceGroupConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eMRInstanceGroupConfigType, data, meta)
}

func resourceEMRInstanceGroupConfigCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eMRInstanceGroupConfigType, data, meta)
}
