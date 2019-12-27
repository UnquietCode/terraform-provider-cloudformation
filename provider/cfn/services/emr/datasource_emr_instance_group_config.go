// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html

package emr

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eMRInstanceGroupConfigType string = "AWS::EMR::InstanceGroupConfig"

func DatasourceEMRInstanceGroupConfig() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEMRInstanceGroupConfigRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceEMRInstanceGroupConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eMRInstanceGroupConfigType, DatasourceEMRInstanceGroupConfig(), data, meta)
}
