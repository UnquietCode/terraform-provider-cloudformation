// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html

package elasticbeanstalk

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationApplicationVersionLifecycleConfig(extras...string) *schema.Resource {
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
			"max_age_rule": {
				Type: schema.TypeList,
				Elem: propertyApplicationMaxAgeRule(),
				Optional: true,
				MaxItems: 1,
			},
			"max_count_rule": {
				Type: schema.TypeList,
				Elem: propertyApplicationMaxCountRule(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
