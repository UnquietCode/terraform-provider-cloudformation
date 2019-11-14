// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceOpsWorksInstance() *schema.Resource {
	return &schema.Resource{
		Exists: resourceOpsWorksInstanceExists,
		Read: resourceOpsWorksInstanceRead,
		Create: resourceOpsWorksInstanceCreate,
		Update: resourceOpsWorksInstanceUpdate,
		Delete: resourceOpsWorksInstanceDelete,
		CustomizeDiff: resourceOpsWorksInstanceCustomizeDiff,
		
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
			},
		},
	}
}

func resourceOpsWorksInstanceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceOpsWorksInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorks::Instance", ResourceOpsWorksInstance(), data, meta)
}

func resourceOpsWorksInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorks::Instance", ResourceOpsWorksInstance(), data, meta)
}

func resourceOpsWorksInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorks::Instance", ResourceOpsWorksInstance(), data, meta)
}

func resourceOpsWorksInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorks::Instance", data, meta)
}

func resourceOpsWorksInstanceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
