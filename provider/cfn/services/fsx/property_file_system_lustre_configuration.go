// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package fsx

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyFileSystemLustreConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"import_path": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"weekly_maintenance_start_time": {
				Type: schema.TypeString,
				Required: false,
			},
			"imported_file_chunk_size": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"export_path": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}