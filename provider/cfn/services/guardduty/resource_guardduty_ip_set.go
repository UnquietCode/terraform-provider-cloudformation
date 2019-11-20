// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const guardDutyIPSetType string = "AWS::GuardDuty::IPSet"

func ResourceGuardDutyIPSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceGuardDutyIPSetRead,
		Create: resourceGuardDutyIPSetCreate,
		Update: resourceGuardDutyIPSetUpdate,
		Delete: resourceGuardDutyIPSetDelete,
		CustomizeDiff: resourceGuardDutyIPSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"format": {
				Type: schema.TypeString,
				Required: true,
			},
			"activate": {
				Type: schema.TypeBool,
				Required: true,
			},
			"detector_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"location": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGuardDutyIPSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(guardDutyIPSetType, ResourceGuardDutyIPSet(), data, meta)
}

func resourceGuardDutyIPSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(guardDutyIPSetType, ResourceGuardDutyIPSet(), data, meta)
}

func resourceGuardDutyIPSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(guardDutyIPSetType, ResourceGuardDutyIPSet(), data, meta)
}

func resourceGuardDutyIPSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(guardDutyIPSetType, data, meta)
}

func resourceGuardDutyIPSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(guardDutyIPSetType, data, meta)
}
