// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html

package dms

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEndpointElasticsearchSettings(extras...string) *schema.Resource {
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
			"endpoint_uri": {
				Type: schema.TypeString,
				Optional: true,
			},
			"full_load_error_percentage": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"error_retry_duration": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"service_access_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
