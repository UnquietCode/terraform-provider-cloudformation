// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-streamingdistribution.html

package cloudfront

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudFrontStreamingDistributionType string = "AWS::CloudFront::StreamingDistribution"

func ResourceCloudFrontStreamingDistribution() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCloudFrontStreamingDistributionExists,
		Read: resourceCloudFrontStreamingDistributionRead,
		Create: resourceCloudFrontStreamingDistributionCreate,
		Update: resourceCloudFrontStreamingDistributionUpdate,
		Delete: resourceCloudFrontStreamingDistributionDelete,
		CustomizeDiff: resourceCloudFrontStreamingDistributionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"streaming_distribution_config": {
				Type: schema.TypeSet,
				Elem: propertyStreamingDistributionStreamingDistributionConfig(),
				Required: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudFrontStreamingDistributionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCloudFrontStreamingDistributionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudFrontStreamingDistributionType, ResourceCloudFrontStreamingDistribution(), data, meta)
}

func resourceCloudFrontStreamingDistributionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloudFrontStreamingDistributionType, ResourceCloudFrontStreamingDistribution(), data, meta)
}

func resourceCloudFrontStreamingDistributionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloudFrontStreamingDistributionType, ResourceCloudFrontStreamingDistribution(), data, meta)
}

func resourceCloudFrontStreamingDistributionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloudFrontStreamingDistributionType, data, meta)
}

func resourceCloudFrontStreamingDistributionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloudFrontStreamingDistributionType, data, meta)
}
