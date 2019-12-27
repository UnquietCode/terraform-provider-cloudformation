// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html

package sagemaker

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sageMakerEndpointType string = "AWS::SageMaker::Endpoint"

func DatasourceSageMakerEndpoint() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSageMakerEndpointRead,
		
		Schema: map[string]*schema.Schema{
			"retain_all_variant_properties": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"endpoint_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"exclude_retained_variant_properties": {
				Type: schema.TypeList,
				Elem: propertyEndpointVariantProperty(),
				Optional: true,
			},
			"endpoint_config_name": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceSageMakerEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerEndpointType, DatasourceSageMakerEndpoint(), data, meta)
}
