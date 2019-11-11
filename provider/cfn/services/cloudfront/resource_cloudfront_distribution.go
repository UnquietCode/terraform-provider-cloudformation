// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceCloudFrontDistribution() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFrontDistributionCreate,
		Read:   resourceCloudFrontDistributionRead,
		Update: resourceCloudFrontDistributionUpdate,
		Delete: resourceCloudFrontDistributionDelete,

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"distribution_config": {
				Type: schema.TypeList,
				Elem: propertyDistributionDistributionConfig(),
				Required: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudFrontDistributionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFront::Distribution", ResourceCloudFrontDistribution(), data, meta)
}

func resourceCloudFrontDistributionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFront::Distribution", ResourceCloudFrontDistribution(), data, meta)
}

func resourceCloudFrontDistributionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFront::Distribution", ResourceCloudFrontDistribution(), data, meta)
}

func resourceCloudFrontDistributionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFront::Distribution", data, meta)
}