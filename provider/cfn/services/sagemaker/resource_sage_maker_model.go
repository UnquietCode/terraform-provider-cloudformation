// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSageMakerModel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSageMakerModelCreate,
		Read:   resourceSageMakerModelRead,
		Update: resourceSageMakerModelUpdate,
		Delete: resourceSageMakerModelDelete,

		Schema: map[string]*schema.Schema{
			"execution_role_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"primary_container": {
				Type: schema.TypeList,
				Elem: propertyContainerDefinition(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"model_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"vpc_config": {
				Type: schema.TypeList,
				Elem: propertyVpcConfig(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"containers": {
				Type: schema.TypeList,
				Elem: propertyContainerDefinition(),
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

func resourceSageMakerModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::Model", data, meta)
}

func resourceSageMakerModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::Model", data, meta)
}

func resourceSageMakerModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::Model", data, meta)
}

func resourceSageMakerModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::Model", data, meta)
}