// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-streamingdistribution.html

package cloudfront

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudFrontStreamingDistributionType string = "AWS::CloudFront::StreamingDistribution"

func DatasourceCloudFrontStreamingDistribution() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCloudFrontStreamingDistributionRead,
		
		Schema: map[string]*schema.Schema{
			"streaming_distribution_config": {
				Type: schema.TypeList,
				Elem: propertyStreamingDistributionStreamingDistributionConfig(),
				Required: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
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

func datasourceCloudFrontStreamingDistributionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudFrontStreamingDistributionType, DatasourceCloudFrontStreamingDistribution(), data, meta)
}
