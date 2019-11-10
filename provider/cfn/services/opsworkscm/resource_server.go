// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package opsworkscm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"key_pair": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"disable_automated_backup": {
				Type: schema.TypeBool,
				Required: false,
			},
			"backup_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"engine_model": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"associate_public_ip_address": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"instance_profile_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"preferred_backup_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"server_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"engine_attributes": {
				Type: schema.TypeList,
				Elem: propertyServerEngineAttribute(),
				Required: false,
			},
			"backup_retention_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"engine": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceServerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorksCM::Server", data, meta)
}

func resourceServerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorksCM::Server", data, meta)
}

func resourceServerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorksCM::Server", data, meta)
}

func resourceServerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorksCM::Server", data, meta)
}