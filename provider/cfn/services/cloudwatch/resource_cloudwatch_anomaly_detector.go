// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html

package cloudwatch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCloudWatchAnomalyDetector() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudWatchAnomalyDetectorCreate,
		Read:   resourceCloudWatchAnomalyDetectorRead,
		Update: resourceCloudWatchAnomalyDetectorUpdate,
		Delete: resourceCloudWatchAnomalyDetectorDelete,

		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stat": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"configuration": {
				Type: schema.TypeList,
				Elem: propertyAnomalyDetectorConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"dimensions": {
				Type: schema.TypeList,
				Elem: propertyAnomalyDetectorDimension(),
				Required: false,
				ForceNew: true,
			},
			"namespace": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudWatchAnomalyDetectorCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudWatch::AnomalyDetector", data, meta)
}

func resourceCloudWatchAnomalyDetectorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudWatch::AnomalyDetector", data, meta)
}

func resourceCloudWatchAnomalyDetectorUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudWatch::AnomalyDetector", data, meta)
}

func resourceCloudWatchAnomalyDetectorDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudWatch::AnomalyDetector", data, meta)
}