// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceOpsWorksCMServerCreate,
		Read:   resourceOpsWorksCMServerRead,
		Update: resourceOpsWorksCMServerUpdate,
		Delete: resourceOpsWorksCMServerDelete,

		Schema: map[string]*schema.Schema{
			"key_pair": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"disable_automated_backup": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"backup_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"engine_model": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"associate_public_ip_address": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"instance_profile_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"preferred_backup_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
			},
			"server_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"engine": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceOpsWorksCMServerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorksCM::Server", data, meta)
}

func resourceOpsWorksCMServerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorksCM::Server", data, meta)
}

func resourceOpsWorksCMServerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorksCM::Server", data, meta)
}

func resourceOpsWorksCMServerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorksCM::Server", data, meta)
}