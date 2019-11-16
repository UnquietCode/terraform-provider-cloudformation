// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const opsWorksStackType string = "AWS::OpsWorks::Stack"

var opsWorksStackProperties map[string]string = map[string]string{
	"agent_version": "AgentVersion",
	"attributes": "Attributes",
	"chef_configuration": "ChefConfiguration",
	"clone_app_ids": "CloneAppIds",
	"clone_permissions": "ClonePermissions",
	"configuration_manager": "ConfigurationManager",
	"custom_cookbooks_source": "CustomCookbooksSource",
	"custom_json": "CustomJson",
	"default_availability_zone": "DefaultAvailabilityZone",
	"default_instance_profile_arn": "DefaultInstanceProfileArn",
	"default_os": "DefaultOs",
	"default_root_device_type": "DefaultRootDeviceType",
	"default_ssh_key_name": "DefaultSshKeyName",
	"default_subnet_id": "DefaultSubnetId",
	"ecs_cluster_arn": "EcsClusterArn",
	"elastic_ips": "ElasticIps",
	"hostname_theme": "HostnameTheme",
	"name": "Name",
	"rds_db_instances": "RdsDbInstances",
	"service_role_arn": "ServiceRoleArn",
	"source_stack_id": "SourceStackId",
	"tags": "Tags",
	"use_custom_cookbooks": "UseCustomCookbooks",
	"use_opsworks_security_groups": "UseOpsworksSecurityGroups",
	"vpc_id": "VpcId",
}

func ResourceOpsWorksStack() *schema.Resource {
	return &schema.Resource{
		Exists: resourceOpsWorksStackExists,
		Read: resourceOpsWorksStackRead,
		Create: resourceOpsWorksStackCreate,
		Update: resourceOpsWorksStackUpdate,
		Delete: resourceOpsWorksStackDelete,
		CustomizeDiff: resourceOpsWorksStackCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"agent_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"attributes": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"chef_configuration": {
				Type: schema.TypeList,
				Elem: propertyStackChefConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"clone_app_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"clone_permissions": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"configuration_manager": {
				Type: schema.TypeList,
				Elem: propertyStackStackConfigurationManager(),
				Optional: true,
				MaxItems: 1,
			},
			"custom_cookbooks_source": {
				Type: schema.TypeList,
				Elem: propertyStackSource(),
				Optional: true,
				MaxItems: 1,
			},
			"custom_json": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"default_availability_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_instance_profile_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"default_os": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_root_device_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_ssh_key_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_subnet_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ecs_cluster_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"elastic_ips": {
				Type: schema.TypeSet,
				Elem: propertyStackElasticIp(),
				Optional: true,
			},
			"hostname_theme": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"rds_db_instances": {
				Type: schema.TypeSet,
				Elem: propertyStackRdsDbInstance(),
				Optional: true,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_stack_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"use_custom_cookbooks": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"use_opsworks_security_groups": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"vpc_id": {
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

func resourceOpsWorksStackExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceOpsWorksStackRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(opsWorksStackType, ResourceOpsWorksStack(), data, meta)
}

func resourceOpsWorksStackCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(opsWorksStackType, ResourceOpsWorksStack(), data, opsWorksStackProperties, meta)
}

func resourceOpsWorksStackUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(opsWorksStackType, ResourceOpsWorksStack(), data, opsWorksStackProperties, meta)
}

func resourceOpsWorksStackDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(opsWorksStackType, data, meta)
}

func resourceOpsWorksStackCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(opsWorksStackType, data, meta)
}
