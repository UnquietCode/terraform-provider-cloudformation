// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html

package sagemaker

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sageMakerEndpointConfigType string = "AWS::SageMaker::EndpointConfig"

func DatasourceSageMakerEndpointConfig() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSageMakerEndpointConfigRead,
		
		Schema: map[string]*schema.Schema{
			"production_variants": {
				Type: schema.TypeList,
				Elem: propertyEndpointConfigProductionVariant(),
				Required: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"endpoint_config_name": {
				Type: schema.TypeString,
				Optional: true,
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

func datasourceSageMakerEndpointConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerEndpointConfigType, DatasourceSageMakerEndpointConfig(), data, meta)
}
