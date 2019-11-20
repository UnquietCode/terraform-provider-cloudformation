// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html

package greengrass

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyResourceDefinitionResourceDataContainer(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secrets_manager_secret_resource_data": {
				Type: schema.TypeSet,
				Elem: propertyResourceDefinitionSecretsManagerSecretResourceData(),
				Optional: true,
				MaxItems: 1,
			},
			"sage_maker_machine_learning_model_resource_data": {
				Type: schema.TypeSet,
				Elem: propertyResourceDefinitionSageMakerMachineLearningModelResourceData(),
				Optional: true,
				MaxItems: 1,
			},
			"local_volume_resource_data": {
				Type: schema.TypeSet,
				Elem: propertyResourceDefinitionLocalVolumeResourceData(),
				Optional: true,
				MaxItems: 1,
			},
			"local_device_resource_data": {
				Type: schema.TypeSet,
				Elem: propertyResourceDefinitionLocalDeviceResourceData(),
				Optional: true,
				MaxItems: 1,
			},
			"s3_machine_learning_model_resource_data": {
				Type: schema.TypeSet,
				Elem: propertyResourceDefinitionS3MachineLearningModelResourceData(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
