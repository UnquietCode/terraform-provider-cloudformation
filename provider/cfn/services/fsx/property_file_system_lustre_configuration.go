// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html

package fsx

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyFileSystemLustreConfiguration(extras...string) *schema.Resource {
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
			"import_path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"weekly_maintenance_start_time": {
				Type: schema.TypeString,
				Optional: true,
			},
			"imported_file_chunk_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"export_path": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
