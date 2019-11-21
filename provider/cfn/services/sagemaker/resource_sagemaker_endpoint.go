// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
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

func ResourceSageMakerEndpoint() *schema.Resource {
	return &schema.Resource{
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
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceSageMakerEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerEndpointType, ResourceSageMakerEndpoint(), data, meta)
}

func resourceSageMakerEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sageMakerEndpointType, ResourceSageMakerEndpoint(), data, meta)
}

func resourceSageMakerEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sageMakerEndpointType, ResourceSageMakerEndpoint(), data, meta)
}

func resourceSageMakerEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sageMakerEndpointType, data, meta)
}

func resourceSageMakerEndpointCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sageMakerEndpointType, data, meta)
}
