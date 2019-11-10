// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-partition.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGluePartition() *schema.Resource {
	return &schema.Resource{
		Create: resourceGluePartitionCreate,
		Read:   resourceGluePartitionRead,
		Update: resourceGluePartitionUpdate,
		Delete: resourceGluePartitionDelete,

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

func resourceGluePartitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Partition", data, meta)
}

func resourceGluePartitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Partition", data, meta)
}

func resourceGluePartitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Partition", data, meta)
}

func resourceGluePartitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Partition", data, meta)
}