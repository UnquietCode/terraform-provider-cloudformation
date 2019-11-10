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

func ResourceModel() *schema.Resource {
	return &schema.Resource{
		Create: resourceModelCreate,
		Read:   resourceModelRead,
		Update: resourceModelUpdate,
		Delete: resourceModelDelete,

		Schema: map[string]*schema.Schema{
			"execution_role_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"primary_container": {
				Type: schema.TypeList,
				Elem: propertyModelContainerDefinition(),
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
				Elem: propertyModelVpcConfig(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"containers": {
				Type: schema.TypeList,
				Elem: propertyModelContainerDefinition(),
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

func resourceModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::Model", data, meta)
}

func resourceModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::Model", data, meta)
}

func resourceModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::Model", data, meta)
}

func resourceModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::Model", data, meta)
}