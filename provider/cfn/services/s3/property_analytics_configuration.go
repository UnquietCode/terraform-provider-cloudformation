// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAnalyticsConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type: schema.TypeString,
				Required: true,
			},
			"prefix": {
				Type: schema.TypeString,
				Required: false,
			},
			"storage_class_analysis": {
				Type: schema.TypeList,
				Elem: propertyStorageClassAnalysis(),
				Required: true,
				MaxItems: 1,
			},
			"tag_filters": {
				Type: schema.TypeSet,
				Elem: propertyTagFilter(),
				Required: false,
			},
		},
	}
}