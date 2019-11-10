// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dynamodb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTableLocalSecondaryIndex() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"key_schema": {
				Type: schema.TypeSet,
				Elem: propertyTableKeySchema(),
				Required: true,
			},
			"projection": {
				Type: schema.TypeList,
				Elem: propertyTableProjection(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}