// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyResourceDefinitionVersionResourceDataContainer(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secrets_manager_secret_resource_data": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionVersionSecretsManagerSecretResourceData(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"sage_maker_machine_learning_model_resource_data": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionVersionSageMakerMachineLearningModelResourceData(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"local_volume_resource_data": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionVersionLocalVolumeResourceData(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"local_device_resource_data": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionVersionLocalDeviceResourceData(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"s3_machine_learning_model_resource_data": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionVersionS3MachineLearningModelResourceData(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}