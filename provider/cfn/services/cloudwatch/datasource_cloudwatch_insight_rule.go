// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-insightrule.html

package cloudwatch

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudWatchInsightRuleType string = "AWS::CloudWatch::InsightRule"

func DatasourceCloudWatchInsightRule() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCloudWatchInsightRuleRead,
		
		Schema: map[string]*schema.Schema{
			"rule_state": {
				Type: schema.TypeString,
				Required: true,
			},
			"rule_body": {
				Type: schema.TypeString,
				Required: true,
			},
			"rule_name": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceCloudWatchInsightRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudWatchInsightRuleType, DatasourceCloudWatchInsightRule(), data, meta)
}
