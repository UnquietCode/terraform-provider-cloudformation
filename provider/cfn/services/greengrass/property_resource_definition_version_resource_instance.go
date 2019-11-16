// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.html

package greengrass

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var resourceDefinitionVersionResourceInstanceProperties map[string]string = map[string]string{
	"resource_data_container": "ResourceDataContainer",
	"id": "Id",
	"name": "Name",
}

func propertyResourceDefinitionVersionResourceInstance(extras...string) *schema.Resource {
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
			"resource_data_container": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionVersionResourceDataContainer(),
				Required: true,
				MaxItems: 1,
			},
			"id": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
