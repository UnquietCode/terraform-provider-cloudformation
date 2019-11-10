// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyChannelChannelStorage() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"customer_managed_s3": {
				Type: schema.TypeList,
				Elem: propertyChannelCustomerManagedS3(),
				Required: false,
				MaxItems: 1,
			},
			"service_managed_s3": {
				Type: schema.TypeList,
				Elem: propertyChannelServiceManagedS3(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}