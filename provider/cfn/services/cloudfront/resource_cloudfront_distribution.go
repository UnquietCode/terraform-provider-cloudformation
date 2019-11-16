// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html

package cloudfront

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudFrontDistributionType string = "AWS::CloudFront::Distribution"

var cloudFrontDistributionProperties map[string]string = map[string]string{
	"distribution_config": "DistributionConfig",
	"tags": "Tags",
}

func ResourceCloudFrontDistribution() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCloudFrontDistributionExists,
		Read: resourceCloudFrontDistributionRead,
		Create: resourceCloudFrontDistributionCreate,
		Update: resourceCloudFrontDistributionUpdate,
		Delete: resourceCloudFrontDistributionDelete,
		CustomizeDiff: resourceCloudFrontDistributionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"distribution_config": {
				Type: schema.TypeSet,
				Elem: propertyDistributionDistributionConfig(),
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

func resourceCloudFrontDistributionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCloudFrontDistributionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudFrontDistributionType, ResourceCloudFrontDistribution(), data, meta)
}

func resourceCloudFrontDistributionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloudFrontDistributionType, ResourceCloudFrontDistribution(), data, cloudFrontDistributionProperties, meta)
}

func resourceCloudFrontDistributionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloudFrontDistributionType, ResourceCloudFrontDistribution(), data, cloudFrontDistributionProperties, meta)
}

func resourceCloudFrontDistributionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloudFrontDistributionType, data, meta)
}

func resourceCloudFrontDistributionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloudFrontDistributionType, data, meta)
}
