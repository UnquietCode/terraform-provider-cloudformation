// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functionconfiguration.html

package greengrass

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyFunctionDefinitionFunctionConfiguration(extras...string) *schema.Resource {
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
			"memory_size": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"pinned": {
				Type: schema.TypeBool,
				Required: false,
			},
			"exec_args": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"timeout": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"encoding_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"environment": {
				Type: schema.TypeList,
				Elem: propertyFunctionDefinitionEnvironment(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"executable": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}