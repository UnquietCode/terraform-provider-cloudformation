// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const opsWorksLayerType string = "AWS::OpsWorks::Layer"

func ResourceOpsWorksLayer() *schema.Resource {
	return &schema.Resource{
		Exists: resourceOpsWorksLayerExists,
		Read: resourceOpsWorksLayerRead,
		Create: resourceOpsWorksLayerCreate,
		Update: resourceOpsWorksLayerUpdate,
		Delete: resourceOpsWorksLayerDelete,
		CustomizeDiff: resourceOpsWorksLayerCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"attributes": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"auto_assign_elastic_ips": {
				Type: schema.TypeBool,
				Required: true,
			},
			"auto_assign_public_ips": {
				Type: schema.TypeBool,
				Required: true,
			},
			"custom_instance_profile_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"custom_json": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"custom_recipes": {
				Type: schema.TypeSet,
				Elem: propertyLayerRecipes(),
				Optional: true,
				MaxItems: 1,
			},
			"custom_security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"enable_auto_healing": {
				Type: schema.TypeBool,
				Required: true,
			},
			"install_updates_on_boot": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"lifecycle_event_configuration": {
				Type: schema.TypeSet,
				Elem: propertyLayerLifecycleEventConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"load_based_auto_scaling": {
				Type: schema.TypeSet,
				Elem: propertyLayerLoadBasedAutoScaling(),
				Optional: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"packages": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"shortname": {
				Type: schema.TypeString,
				Required: true,
			},
			"stack_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"use_ebs_optimized_instances": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"volume_configurations": {
				Type: schema.TypeList,
				Elem: propertyLayerVolumeConfiguration(),
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

func resourceOpsWorksLayerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceOpsWorksLayerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(opsWorksLayerType, ResourceOpsWorksLayer(), data, meta)
}

func resourceOpsWorksLayerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(opsWorksLayerType, ResourceOpsWorksLayer(), data, meta)
}

func resourceOpsWorksLayerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(opsWorksLayerType, ResourceOpsWorksLayer(), data, meta)
}

func resourceOpsWorksLayerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(opsWorksLayerType, data, meta)
}

func resourceOpsWorksLayerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(opsWorksLayerType, data, meta)
}
