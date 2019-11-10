// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cloudfront

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceStreamingDistribution() *schema.Resource {
	return &schema.Resource{
		Create: resourceStreamingDistributionCreate,
		Read:   resourceStreamingDistributionRead,
		Update: resourceStreamingDistributionUpdate,
		Delete: resourceStreamingDistributionDelete,

		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourceStreamingDistributionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFront::StreamingDistribution", data, meta)
}

func resourceStreamingDistributionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFront::StreamingDistribution", data, meta)
}

func resourceStreamingDistributionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFront::StreamingDistribution", data, meta)
}

func resourceStreamingDistributionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFront::StreamingDistribution", data, meta)
}