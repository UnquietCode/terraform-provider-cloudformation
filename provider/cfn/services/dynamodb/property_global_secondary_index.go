// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package dynamodb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyGlobalSecondaryIndex() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"key_schema": {
				Type: schema.TypeSet,
				Elem: propertyKeySchema(),
				Required: true,
			},
			"projection": {
				Type: schema.TypeList,
				Elem: propertyProjection(),
				Required: true,
				MaxItems: 1,
			},
			"provisioned_throughput": {
				Type: schema.TypeList,
				Elem: propertyProvisionedThroughput(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}