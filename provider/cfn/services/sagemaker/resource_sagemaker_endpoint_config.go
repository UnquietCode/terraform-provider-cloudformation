// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

const sageMakerEndpointConfigType string = "AWS::SageMaker::EndpointConfig"

func ResourceSageMakerEndpointConfig() *schema.Resource {
	return &schema.Resource{
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
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSageMakerEndpointConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerEndpointConfigType, ResourceSageMakerEndpointConfig(), data, meta)
}

func resourceSageMakerEndpointConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sageMakerEndpointConfigType, ResourceSageMakerEndpointConfig(), data, meta)
}

func resourceSageMakerEndpointConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sageMakerEndpointConfigType, ResourceSageMakerEndpointConfig(), data, meta)
}

func resourceSageMakerEndpointConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sageMakerEndpointConfigType, data, meta)
}

func resourceSageMakerEndpointConfigCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sageMakerEndpointConfigType, data, meta)
}
