// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-model.html

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sageMakerModelType string = "AWS::SageMaker::Model"

var sageMakerModelProperties map[string]string = map[string]string{
	"execution_role_arn": "ExecutionRoleArn",
	"primary_container": "PrimaryContainer",
	"model_name": "ModelName",
	"vpc_config": "VpcConfig",
	"containers": "Containers",
	"tags": "Tags",
}

func ResourceSageMakerModel() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSageMakerModelExists,
		Read: resourceSageMakerModelRead,
		Create: resourceSageMakerModelCreate,
		Update: resourceSageMakerModelUpdate,
		Delete: resourceSageMakerModelDelete,
		CustomizeDiff: resourceSageMakerModelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"execution_role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"primary_container": {
				Type: schema.TypeSet,
				Elem: propertyModelContainerDefinition(),
				Optional: true,
				MaxItems: 1,
			},
			"model_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc_config": {
				Type: schema.TypeSet,
				Elem: propertyModelVpcConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"containers": {
				Type: schema.TypeList,
				Elem: propertyModelContainerDefinition(),
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

func resourceSageMakerModelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSageMakerModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerModelType, ResourceSageMakerModel(), data, meta)
}

func resourceSageMakerModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sageMakerModelType, ResourceSageMakerModel(), data, sageMakerModelProperties, meta)
}

func resourceSageMakerModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sageMakerModelType, ResourceSageMakerModel(), data, sageMakerModelProperties, meta)
}

func resourceSageMakerModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sageMakerModelType, data, meta)
}

func resourceSageMakerModelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sageMakerModelType, data, meta)
}
