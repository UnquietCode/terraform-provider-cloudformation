// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html

package cloudwatch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudWatchAnomalyDetectorType string = "AWS::CloudWatch::AnomalyDetector"

func ResourceCloudWatchAnomalyDetector() *schema.Resource {
	return &schema.Resource{
		Read: resourceCloudWatchAnomalyDetectorRead,
		Create: resourceCloudWatchAnomalyDetectorCreate,
		Update: resourceCloudWatchAnomalyDetectorUpdate,
		Delete: resourceCloudWatchAnomalyDetectorDelete,
		CustomizeDiff: resourceCloudWatchAnomalyDetectorCustomizeDiff,
		
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
				Type: schema.TypeSet,
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
			},
		},
	}
}

func resourceCloudWatchAnomalyDetectorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudWatchAnomalyDetectorType, ResourceCloudWatchAnomalyDetector(), data, meta)
}

func resourceCloudWatchAnomalyDetectorCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloudWatchAnomalyDetectorType, ResourceCloudWatchAnomalyDetector(), data, meta)
}

func resourceCloudWatchAnomalyDetectorUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloudWatchAnomalyDetectorType, ResourceCloudWatchAnomalyDetector(), data, meta)
}

func resourceCloudWatchAnomalyDetectorDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloudWatchAnomalyDetectorType, data, meta)
}

func resourceCloudWatchAnomalyDetectorCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloudWatchAnomalyDetectorType, data, meta)
}
