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

func ResourceDistribution() *schema.Resource {
	return &schema.Resource{
		Create: resourceDistributionCreate,
		Read:   resourceDistributionRead,
		Update: resourceDistributionUpdate,
		Delete: resourceDistributionDelete,

		Schema: map[string]*schema.Schema{
			"distribution_config": {
				Type: schema.TypeList,
				Elem: propertyDistributionDistributionConfig(),
				Required: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceDistributionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFront::Distribution", data, meta)
}

func resourceDistributionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFront::Distribution", data, meta)
}

func resourceDistributionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFront::Distribution", data, meta)
}

func resourceDistributionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFront::Distribution", data, meta)
}