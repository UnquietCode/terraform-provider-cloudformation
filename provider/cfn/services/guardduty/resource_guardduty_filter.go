// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html

package guardduty

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const guardDutyFilterType string = "AWS::GuardDuty::Filter"

func ResourceGuardDutyFilter() *schema.Resource {
	return &schema.Resource{
		Read: resourceGuardDutyFilterRead,
		Create: resourceGuardDutyFilterCreate,
		Update: resourceGuardDutyFilterUpdate,
		Delete: resourceGuardDutyFilterDelete,
		CustomizeDiff: resourceGuardDutyFilterCustomizeDiff,
		
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
			},
			"finding_criteria": {
				Type: schema.TypeSet,
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
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceGuardDutyFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(guardDutyFilterType, ResourceGuardDutyFilter(), data, meta)
}

func resourceGuardDutyFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(guardDutyFilterType, ResourceGuardDutyFilter(), data, meta)
}

func resourceGuardDutyFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(guardDutyFilterType, ResourceGuardDutyFilter(), data, meta)
}

func resourceGuardDutyFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(guardDutyFilterType, data, meta)
}

func resourceGuardDutyFilterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(guardDutyFilterType, data, meta)
}
