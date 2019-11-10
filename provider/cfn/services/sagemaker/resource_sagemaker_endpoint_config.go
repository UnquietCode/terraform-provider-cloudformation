// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSageMakerEndpointConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceSageMakerEndpointConfigCreate,
		Read:   resourceSageMakerEndpointConfigRead,
		Update: resourceSageMakerEndpointConfigUpdate,
		Delete: resourceSageMakerEndpointConfigDelete,

		Schema: map[string]*schema.Schema{
			"production_variants": {
				Type: schema.TypeList,
				Elem: propertyEndpointConfigProductionVariant(),
				Required: true,
				ForceNew: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"endpoint_config_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
		},
	}
}

func resourceSageMakerEndpointConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::EndpointConfig", data, meta)
}

func resourceSageMakerEndpointConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::EndpointConfig", data, meta)
}

func resourceSageMakerEndpointConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::EndpointConfig", data, meta)
}

func resourceSageMakerEndpointConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::EndpointConfig", data, meta)
}