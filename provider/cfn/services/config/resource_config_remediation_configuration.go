// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-remediationconfiguration.html

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigRemediationConfiguration() *schema.Resource {
	return &schema.Resource{
		Exists: resourceConfigRemediationConfigurationExists,
		Read:   resourceConfigRemediationConfigurationRead,
		Create: resourceConfigRemediationConfigurationCreate,
		Update: resourceConfigRemediationConfigurationUpdate,
		Delete: resourceConfigRemediationConfigurationDelete,
		CustomizeDiff: resourceConfigRemediationConfigurationCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"target_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"execution_controls": {
				Type: schema.TypeList,
				Elem: propertyRemediationConfigurationExecutionControls(),
				Optional: true,
				MaxItems: 1,
			},
			"parameters": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"target_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"config_rule_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"retry_attempt_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"maximum_automatic_attempts": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"target_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"automatic": {
				Type: schema.TypeBool,
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

func resourceConfigRemediationConfigurationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceConfigRemediationConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::RemediationConfiguration", ResourceConfigRemediationConfiguration(), data, meta)
}

func resourceConfigRemediationConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::RemediationConfiguration", ResourceConfigRemediationConfiguration(), data, meta)
}

func resourceConfigRemediationConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::RemediationConfiguration", ResourceConfigRemediationConfiguration(), data, meta)
}

func resourceConfigRemediationConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::RemediationConfiguration", data, meta)
}

func resourceConfigRemediationConfigurationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}