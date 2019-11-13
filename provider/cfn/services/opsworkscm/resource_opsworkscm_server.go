// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html

package opsworkscm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceOpsWorksCMServer() *schema.Resource {
	return &schema.Resource{
		Exists: resourceOpsWorksCMServerExists,
		Read:   resourceOpsWorksCMServerRead,
		Create: resourceOpsWorksCMServerCreate,
		Update: resourceOpsWorksCMServerUpdate,
		Delete: resourceOpsWorksCMServerDelete,
		CustomizeDiff: resourceOpsWorksCMServerCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"key_pair": {
				Type: schema.TypeString,
				Optional: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"disable_automated_backup": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"backup_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"engine_model": {
				Type: schema.TypeString,
				Optional: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"associate_public_ip_address": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"instance_profile_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"preferred_backup_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"server_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"engine_attributes": {
				Type: schema.TypeList,
				Elem: propertyServerEngineAttribute(),
				Optional: true,
			},
			"backup_retention_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"engine": {
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

func resourceOpsWorksCMServerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceOpsWorksCMServerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorksCM::Server", ResourceOpsWorksCMServer(), data, meta)
}

func resourceOpsWorksCMServerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorksCM::Server", ResourceOpsWorksCMServer(), data, meta)
}

func resourceOpsWorksCMServerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorksCM::Server", ResourceOpsWorksCMServer(), data, meta)
}

func resourceOpsWorksCMServerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorksCM::Server", data, meta)
}

func resourceOpsWorksCMServerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}