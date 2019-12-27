// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html

package opsworkscm

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const opsWorksCMServerType string = "AWS::OpsWorksCM::Server"

func DatasourceOpsWorksCMServer() *schema.Resource {
	return &schema.Resource{
		Read: datasourceOpsWorksCMServerRead,
		
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
			"custom_certificate": {
				Type: schema.TypeString,
				Optional: true,
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
			"custom_domain": {
				Type: schema.TypeString,
				Optional: true,
			},
			"custom_private_key": {
				Type: schema.TypeString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceOpsWorksCMServerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(opsWorksCMServerType, DatasourceOpsWorksCMServer(), data, meta)
}
