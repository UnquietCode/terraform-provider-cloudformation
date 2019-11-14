// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSageMakerEndpoint() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSageMakerEndpointExists,
		Read: resourceSageMakerEndpointRead,
		Create: resourceSageMakerEndpointCreate,
		Update: resourceSageMakerEndpointUpdate,
		Delete: resourceSageMakerEndpointDelete,
		CustomizeDiff: resourceSageMakerEndpointCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"endpoint_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"endpoint_config_name": {
				Type: schema.TypeString,
				Required: true,
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

func resourceSageMakerEndpointExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSageMakerEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::Endpoint", ResourceSageMakerEndpoint(), data, meta)
}

func resourceSageMakerEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::Endpoint", ResourceSageMakerEndpoint(), data, meta)
}

func resourceSageMakerEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::Endpoint", ResourceSageMakerEndpoint(), data, meta)
}

func resourceSageMakerEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::Endpoint", data, meta)
}

func resourceSageMakerEndpointCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
