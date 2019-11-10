// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyResourceDefinitionResourceDataContainer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secrets_manager_secret_resource_data": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionSecretsManagerSecretResourceData(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"sage_maker_machine_learning_model_resource_data": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionSageMakerMachineLearningModelResourceData(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"local_volume_resource_data": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionLocalVolumeResourceData(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"local_device_resource_data": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionLocalDeviceResourceData(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"s3_machine_learning_model_resource_data": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionS3MachineLearningModelResourceData(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}