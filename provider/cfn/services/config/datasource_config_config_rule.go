// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html

package config

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const configConfigRuleType string = "AWS::Config::ConfigRule"

func DatasourceConfigConfigRule() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConfigConfigRuleRead,
		
		Schema: map[string]*schema.Schema{
			"config_rule_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"input_parameters": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"maximum_execution_frequency": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scope": {
				Type: schema.TypeList,
				Elem: propertyConfigRuleScope(),
				Optional: true,
				MaxItems: 1,
			},
			"source": {
				Type: schema.TypeList,
				Elem: propertyConfigRuleSource(),
				Required: true,
				MaxItems: 1,
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

func datasourceConfigConfigRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(configConfigRuleType, DatasourceConfigConfigRule(), data, meta)
}
