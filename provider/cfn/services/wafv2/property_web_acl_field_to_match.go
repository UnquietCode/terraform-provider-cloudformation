// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html

package wafv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyWebACLFieldToMatch(extras...string) *schema.Resource {
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
			"single_header": {
				Type: schema.TypeList,
				Elem: propertyWebACLSingleHeader(),
				Optional: true,
				MaxItems: 1,
			},
			"single_query_argument": {
				Type: schema.TypeList,
				Elem: propertyWebACLSingleQueryArgument(),
				Optional: true,
				MaxItems: 1,
			},
			"all_query_arguments": {
				Type: schema.TypeList,
				Elem: propertyWebACLAllQueryArguments(),
				Optional: true,
				MaxItems: 1,
			},
			"uri_path": {
				Type: schema.TypeList,
				Elem: propertyWebACLUriPath(),
				Optional: true,
				MaxItems: 1,
			},
			"query_string": {
				Type: schema.TypeList,
				Elem: propertyWebACLQueryString(),
				Optional: true,
				MaxItems: 1,
			},
			"body": {
				Type: schema.TypeList,
				Elem: propertyWebACLBody(),
				Optional: true,
				MaxItems: 1,
			},
			"method": {
				Type: schema.TypeList,
				Elem: propertyWebACLMethod(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
