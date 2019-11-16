// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration-selfmanagedactivedirectoryconfiguration.html

package fsx

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyFileSystemSelfManagedActiveDirectoryConfiguration(extras...string) *schema.Resource {
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
			"file_system_administrators_group": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"organizational_unit_distinguished_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"dns_ips": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"password": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
