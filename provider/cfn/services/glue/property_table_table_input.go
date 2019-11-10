// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTableTableInput() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"owner": {
				Type: schema.TypeString,
				Required: false,
			},
			"view_original_text": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"table_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"parameters": {
				Type: schema.TypeMap,
				Required: false,
			},
			"view_expanded_text": {
				Type: schema.TypeString,
				Required: false,
			},
			"storage_descriptor": {
				Type: schema.TypeList,
				Elem: propertyTableStorageDescriptor(),
				Required: false,
				MaxItems: 1,
			},
			"partition_keys": {
				Type: schema.TypeList,
				Elem: propertyTableColumn(),
				Required: false,
			},
			"retention": {
				Type: schema.TypeInt,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}