// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-additionalauthenticationprovider.html

package appsync

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyGraphQLApiAdditionalAuthenticationProvider(extras...string) *schema.Resource {
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
			"open_id_connect_config": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiOpenIDConnectConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"user_pool_config": {
				Type: schema.TypeList,
				Elem: propertyGraphQLApiCognitoUserPoolConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"authentication_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
