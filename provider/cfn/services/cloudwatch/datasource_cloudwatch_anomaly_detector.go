// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html

package cloudwatch

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudWatchAnomalyDetectorType string = "AWS::CloudWatch::AnomalyDetector"

func DatasourceCloudWatchAnomalyDetector() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCloudWatchAnomalyDetectorRead,
		
		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"stat": {
				Type: schema.TypeString,
				Required: true,
			},
			"configuration": {
				Type: schema.TypeList,
				Elem: propertyAnomalyDetectorConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"dimensions": {
				Type: schema.TypeList,
				Elem: propertyAnomalyDetectorDimension(),
				Optional: true,
			},
			"namespace": {
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

func datasourceCloudWatchAnomalyDetectorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudWatchAnomalyDetectorType, DatasourceCloudWatchAnomalyDetector(), data, meta)
}
