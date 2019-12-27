// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html

package opsworks

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const opsWorksInstanceType string = "AWS::OpsWorks::Instance"

func DatasourceOpsWorksInstance() *schema.Resource {
	return &schema.Resource{
		Read: datasourceOpsWorksInstanceRead,
		
		Schema: map[string]*schema.Schema{
			"agent_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ami_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"architecture": {
				Type: schema.TypeString,
				Optional: true,
			},
			"auto_scaling_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"block_device_mappings": {
				Type: schema.TypeSet,
				Elem: propertyInstanceBlockDeviceMapping(),
				Optional: true,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"elastic_ips": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"hostname": {
				Type: schema.TypeString,
				Optional: true,
			},
			"install_updates_on_boot": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"layer_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"os": {
				Type: schema.TypeString,
				Optional: true,
			},
			"root_device_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ssh_key_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"stack_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tenancy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"time_based_auto_scaling": {
				Type: schema.TypeList,
				Elem: propertyInstanceTimeBasedAutoScaling(),
				Optional: true,
				MaxItems: 1,
			},
			"virtualization_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"volumes": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
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

func datasourceOpsWorksInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(opsWorksInstanceType, DatasourceOpsWorksInstance(), data, meta)
}
