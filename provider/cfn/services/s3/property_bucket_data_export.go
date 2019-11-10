// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBucketDataExport() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"destination": {
				Type: schema.TypeList,
				Elem: propertyBucketDestination(),
				Required: true,
				MaxItems: 1,
			},
			"output_schema_version": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}