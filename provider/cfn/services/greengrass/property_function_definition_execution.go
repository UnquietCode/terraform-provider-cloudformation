// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-execution.html

package greengrass

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var functionDefinitionExecutionProperties map[string]string = map[string]string{
	"isolation_mode": "IsolationMode",
	"run_as": "RunAs",
}

func propertyFunctionDefinitionExecution(extras...string) *schema.Resource {
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
			"isolation_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"run_as": {
				Type: schema.TypeList,
				Elem: propertyFunctionDefinitionRunAs(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
