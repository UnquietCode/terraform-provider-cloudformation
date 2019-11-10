// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceOpsWorksInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceOpsWorksInstanceCreate,
		Read:   resourceOpsWorksInstanceRead,
		Update: resourceOpsWorksInstanceUpdate,
		Delete: resourceOpsWorksInstanceDelete,

		Schema: map[string]*schema.Schema{
			"agent_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"ami_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"architecture": {
				Type: schema.TypeString,
				Required: false,
			},
			"auto_scaling_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"block_device_mappings": {
				Type: schema.TypeSet,
				Elem: propertyInstanceBlockDeviceMapping(),
				Required: false,
				ForceNew: true,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"elastic_ips": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"hostname": {
				Type: schema.TypeString,
				Required: false,
			},
			"install_updates_on_boot": {
				Type: schema.TypeBool,
				Required: false,
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
				Required: false,
			},
			"root_device_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ssh_key_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"stack_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tenancy": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"time_based_auto_scaling": {
				Type: schema.TypeList,
				Elem: propertyInstanceTimeBasedAutoScaling(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"virtualization_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"volumes": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
		},
	}
}

func resourceOpsWorksInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorks::Instance", data, meta)
}

func resourceOpsWorksInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorks::Instance", data, meta)
}

func resourceOpsWorksInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorks::Instance", data, meta)
}

func resourceOpsWorksInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorks::Instance", data, meta)
}