// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEndpointConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointConfigCreate,
		Read:   resourceEndpointConfigRead,
		Update: resourceEndpointConfigUpdate,
		Delete: resourceEndpointConfigDelete,

		Schema: map[string]*schema.Schema{
			"production_variants": {
				Type: schema.TypeList,
				Elem: propertyEndpointConfigProductionVariant(),
				Required: true,
				ForceNew: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"endpoint_config_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceEndpointConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::EndpointConfig", data, meta)
}

func resourceEndpointConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::EndpointConfig", data, meta)
}

func resourceEndpointConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::EndpointConfig", data, meta)
}

func resourceEndpointConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::EndpointConfig", data, meta)
}