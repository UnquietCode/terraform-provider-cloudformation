// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigRemediationConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigRemediationConfigurationCreate,
		Read:   resourceConfigRemediationConfigurationRead,
		Update: resourceConfigRemediationConfigurationUpdate,
		Delete: resourceConfigRemediationConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"target_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"execution_controls": {
				Type: schema.TypeList,
				Elem: propertyExecutionControls(),
				Required: false,
				MaxItems: 1,
			},
			"parameters": {
				Type: schema.TypeMap,
				Required: false,
			},
			"target_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"config_rule_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"retry_attempt_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"maximum_automatic_attempts": {
				Type: schema.TypeInt,
				Required: false,
			},
			"target_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"automatic": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}

func resourceConfigRemediationConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::RemediationConfiguration", data, meta)
}

func resourceConfigRemediationConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::RemediationConfiguration", data, meta)
}

func resourceConfigRemediationConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::RemediationConfiguration", data, meta)
}

func resourceConfigRemediationConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::RemediationConfiguration", data, meta)
}