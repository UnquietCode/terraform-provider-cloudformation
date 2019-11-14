// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceSageMakerEndpointConfigExists,
		Read: resourceSageMakerEndpointConfigRead,
		Create: resourceSageMakerEndpointConfigCreate,
		Update: resourceSageMakerEndpointConfigUpdate,
		Delete: resourceSageMakerEndpointConfigDelete,
		CustomizeDiff: resourceSageMakerEndpointConfigCustomizeDiff,
		
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
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSageMakerEndpointConfigExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSageMakerEndpointConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::EndpointConfig", ResourceSageMakerEndpointConfig(), data, meta)
}

func resourceSageMakerEndpointConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::EndpointConfig", ResourceSageMakerEndpointConfig(), data, meta)
}

func resourceSageMakerEndpointConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::EndpointConfig", ResourceSageMakerEndpointConfig(), data, meta)
}

func resourceSageMakerEndpointConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::EndpointConfig", data, meta)
}

func resourceSageMakerEndpointConfigCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
