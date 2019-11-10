// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceOpsWorksStack() *schema.Resource {
	return &schema.Resource{
		Create: resourceOpsWorksStackCreate,
		Read:   resourceOpsWorksStackRead,
		Update: resourceOpsWorksStackUpdate,
		Delete: resourceOpsWorksStackDelete,

		Schema: map[string]*schema.Schema{
			"agent_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"attributes": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"chef_configuration": {
				Type: schema.TypeList,
				Elem: propertyStackChefConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"clone_app_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"clone_permissions": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"configuration_manager": {
				Type: schema.TypeList,
				Elem: propertyStackStackConfigurationManager(),
				Required: false,
				MaxItems: 1,
			},
			"custom_cookbooks_source": {
				Type: schema.TypeList,
				Elem: propertyStackSource(),
				Required: false,
				MaxItems: 1,
			},
			"custom_json": {
				Type: schema.TypeMap,
				Required: false,
			},
			"default_availability_zone": {
				Type: schema.TypeString,
				Required: false,
			},
			"default_instance_profile_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"default_os": {
				Type: schema.TypeString,
				Required: false,
			},
			"default_root_device_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"default_ssh_key_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"default_subnet_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"ecs_cluster_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"elastic_ips": {
				Type: schema.TypeSet,
				Elem: propertyStackElasticIp(),
				Required: false,
			},
			"hostname_theme": {
				Type: schema.TypeString,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"rds_db_instances": {
				Type: schema.TypeSet,
				Elem: propertyStackRdsDbInstance(),
				Required: false,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_stack_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"use_custom_cookbooks": {
				Type: schema.TypeBool,
				Required: false,
			},
			"use_opsworks_security_groups": {
				Type: schema.TypeBool,
				Required: false,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceOpsWorksStackCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorks::Stack", data, meta)
}

func resourceOpsWorksStackRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorks::Stack", data, meta)
}

func resourceOpsWorksStackUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorks::Stack", data, meta)
}

func resourceOpsWorksStackDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorks::Stack", data, meta)
}