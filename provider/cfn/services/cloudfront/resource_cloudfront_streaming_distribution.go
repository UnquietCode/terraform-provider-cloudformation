// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceCloudFrontStreamingDistribution() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFrontStreamingDistributionCreate,
		Read:   resourceCloudFrontStreamingDistributionRead,
		Update: resourceCloudFrontStreamingDistributionUpdate,
		Delete: resourceCloudFrontStreamingDistributionDelete,

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"streaming_distribution_config": {
				Type: schema.TypeList,
				Elem: propertyStreamingDistributionStreamingDistributionConfig(),
				Required: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudFrontStreamingDistributionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFront::StreamingDistribution", ResourceCloudFrontStreamingDistribution(), data, meta)
}

func resourceCloudFrontStreamingDistributionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFront::StreamingDistribution", ResourceCloudFrontStreamingDistribution(), data, meta)
}

func resourceCloudFrontStreamingDistributionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFront::StreamingDistribution", ResourceCloudFrontStreamingDistribution(), data, meta)
}

func resourceCloudFrontStreamingDistributionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFront::StreamingDistribution", data, meta)
}