// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePartition() *schema.Resource {
	return &schema.Resource{
		Create: resourcePartitionCreate,
		Read:   resourcePartitionRead,
		Update: resourcePartitionUpdate,
		Delete: resourcePartitionDelete,

		Schema: map[string]*schema.Schema{
			"table_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"catalog_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"partition_input": {
				Type: schema.TypeList,
				Elem: propertyPartitionPartitionInput(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourcePartitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Partition", data, meta)
}

func resourcePartitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Partition", data, meta)
}

func resourcePartitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Partition", data, meta)
}

func resourcePartitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Partition", data, meta)
}