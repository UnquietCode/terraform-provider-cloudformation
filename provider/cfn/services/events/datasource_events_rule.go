// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html

package events

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eventsRuleType string = "AWS::Events::Rule"

func DatasourceEventsRule() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEventsRuleRead,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"event_bus_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"event_pattern": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"schedule_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"state": {
				Type: schema.TypeString,
				Optional: true,
			},
			"targets": {
				Type: schema.TypeSet,
				Elem: propertyRuleTarget(),
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

func datasourceEventsRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eventsRuleType, DatasourceEventsRule(), data, meta)
}
