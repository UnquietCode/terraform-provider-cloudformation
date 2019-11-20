// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-partition.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const gluePartitionType string = "AWS::Glue::Partition"

func ResourceGluePartition() *schema.Resource {
	return &schema.Resource{
		Read: resourceGluePartitionRead,
		Create: resourceGluePartitionCreate,
		Update: resourceGluePartitionUpdate,
		Delete: resourceGluePartitionDelete,
		CustomizeDiff: resourceGluePartitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"table_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"catalog_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"partition_input": {
				Type: schema.TypeSet,
				Elem: propertyPartitionPartitionInput(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGluePartitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(gluePartitionType, ResourceGluePartition(), data, meta)
}

func resourceGluePartitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(gluePartitionType, ResourceGluePartition(), data, meta)
}

func resourceGluePartitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(gluePartitionType, ResourceGluePartition(), data, meta)
}

func resourceGluePartitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(gluePartitionType, data, meta)
}

func resourceGluePartitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(gluePartitionType, data, meta)
}
