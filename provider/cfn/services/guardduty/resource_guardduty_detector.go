// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const guardDutyDetectorType string = "AWS::GuardDuty::Detector"

var guardDutyDetectorProperties map[string]string = map[string]string{
	"finding_publishing_frequency": "FindingPublishingFrequency",
	"enable": "Enable",
}

func ResourceGuardDutyDetector() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGuardDutyDetectorExists,
		Read: resourceGuardDutyDetectorRead,
		Create: resourceGuardDutyDetectorCreate,
		Update: resourceGuardDutyDetectorUpdate,
		Delete: resourceGuardDutyDetectorDelete,
		CustomizeDiff: resourceGuardDutyDetectorCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"finding_publishing_frequency": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enable": {
				Type: schema.TypeBool,
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

func resourceGuardDutyDetectorExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGuardDutyDetectorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(guardDutyDetectorType, ResourceGuardDutyDetector(), data, meta)
}

func resourceGuardDutyDetectorCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(guardDutyDetectorType, ResourceGuardDutyDetector(), data, guardDutyDetectorProperties, meta)
}

func resourceGuardDutyDetectorUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(guardDutyDetectorType, ResourceGuardDutyDetector(), data, guardDutyDetectorProperties, meta)
}

func resourceGuardDutyDetectorDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(guardDutyDetectorType, data, meta)
}

func resourceGuardDutyDetectorCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(guardDutyDetectorType, data, meta)
}
