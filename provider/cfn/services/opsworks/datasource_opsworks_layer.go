// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html

package opsworks

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const opsWorksLayerType string = "AWS::OpsWorks::Layer"

func DatasourceOpsWorksLayer() *schema.Resource {
	return &schema.Resource{
		Read: datasourceOpsWorksLayerRead,
		
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
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"custom_recipes": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
				Elem: propertyLayerLifecycleEventConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"load_based_auto_scaling": {
				Type: schema.TypeList,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceOpsWorksLayerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(opsWorksLayerType, DatasourceOpsWorksLayer(), data, meta)
}
