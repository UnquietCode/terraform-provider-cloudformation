// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-remediationconfiguration.html

package config

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const configRemediationConfigurationType string = "AWS::Config::RemediationConfiguration"

func DatasourceConfigRemediationConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConfigRemediationConfigurationRead,
		
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
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceConfigRemediationConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(configRemediationConfigurationType, DatasourceConfigRemediationConfiguration(), data, meta)
}
