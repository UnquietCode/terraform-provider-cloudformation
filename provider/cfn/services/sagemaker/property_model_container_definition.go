// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-containerdefinition.html

package sagemaker

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var modelContainerDefinitionProperties map[string]string = map[string]string{
	"container_hostname": "ContainerHostname",
	"environment": "Environment",
	"model_data_url": "ModelDataUrl",
	"image": "Image",
}

func propertyModelContainerDefinition(extras...string) *schema.Resource {
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
			"container_hostname": {
				Type: schema.TypeString,
				Optional: true,
			},
			"environment": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"model_data_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
