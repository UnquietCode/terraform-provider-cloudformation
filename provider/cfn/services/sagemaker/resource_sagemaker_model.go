// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-model.html

package sagemaker

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sageMakerModelType string = "AWS::SageMaker::Model"

func ResourceSageMakerModel() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceSageMakerModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerModelType, ResourceSageMakerModel(), data, meta)
}

func resourceSageMakerModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sageMakerModelType, ResourceSageMakerModel(), data, meta)
}

func resourceSageMakerModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sageMakerModelType, ResourceSageMakerModel(), data, meta)
}

func resourceSageMakerModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sageMakerModelType, data, meta)
}

func resourceSageMakerModelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sageMakerModelType, data, meta)
}
