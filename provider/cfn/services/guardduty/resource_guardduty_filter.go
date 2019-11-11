// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGuardDutyFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceGuardDutyFilterCreate,
		Read:   resourceGuardDutyFilterRead,
		Update: resourceGuardDutyFilterUpdate,
		Delete: resourceGuardDutyFilterDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"detector_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"finding_criteria": {
				Type: schema.TypeList,
				Elem: propertyFilterFindingCriteria(),
				Required: true,
				MaxItems: 1,
			},
			"rank": {
				Type: schema.TypeInt,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGuardDutyFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GuardDuty::Filter", ResourceGuardDutyFilter(), data, meta)
}

func resourceGuardDutyFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GuardDuty::Filter", ResourceGuardDutyFilter(), data, meta)
}

func resourceGuardDutyFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GuardDuty::Filter", ResourceGuardDutyFilter(), data, meta)
}

func resourceGuardDutyFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GuardDuty::Filter", data, meta)
}