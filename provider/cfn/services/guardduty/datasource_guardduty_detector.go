// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html

package guardduty

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const guardDutyDetectorType string = "AWS::GuardDuty::Detector"

func DatasourceGuardDutyDetector() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGuardDutyDetectorRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceGuardDutyDetectorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(guardDutyDetectorType, DatasourceGuardDutyDetector(), data, meta)
}
