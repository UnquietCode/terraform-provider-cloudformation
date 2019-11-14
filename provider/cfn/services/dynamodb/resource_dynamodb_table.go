// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html

package dynamodb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDynamoDBTable() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDynamoDBTableExists,
		Read: resourceDynamoDBTableRead,
		Create: resourceDynamoDBTableCreate,
		Update: resourceDynamoDBTableUpdate,
		Delete: resourceDynamoDBTableDelete,
		CustomizeDiff: resourceDynamoDBTableCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"attribute_definitions": {
				Type: schema.TypeList,
				Elem: propertyTableAttributeDefinition(),
				Optional: true,
			},
			"billing_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"global_secondary_indexes": {
				Type: schema.TypeList,
				Elem: propertyTableGlobalSecondaryIndex(),
				Optional: true,
			},
			"key_schema": {
				Type: schema.TypeSet,
				Elem: propertyTableKeySchema(),
				Required: true,
			},
			"local_secondary_indexes": {
				Type: schema.TypeList,
				Elem: propertyTableLocalSecondaryIndex(),
				Optional: true,
			},
			"point_in_time_recovery_specification": {
				Type: schema.TypeList,
				Elem: propertyTablePointInTimeRecoverySpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"provisioned_throughput": {
				Type: schema.TypeList,
				Elem: propertyTableProvisionedThroughput(),
				Optional: true,
				MaxItems: 1,
			},
			"sse_specification": {
				Type: schema.TypeList,
				Elem: propertyTableSSESpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"stream_specification": {
				Type: schema.TypeList,
				Elem: propertyTableStreamSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"table_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"time_to_live_specification": {
				Type: schema.TypeList,
				Elem: propertyTableTimeToLiveSpecification(),
				Optional: true,
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

func resourceDynamoDBTableExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDynamoDBTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DynamoDB::Table", ResourceDynamoDBTable(), data, meta)
}

func resourceDynamoDBTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DynamoDB::Table", ResourceDynamoDBTable(), data, meta)
}

func resourceDynamoDBTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DynamoDB::Table", ResourceDynamoDBTable(), data, meta)
}

func resourceDynamoDBTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DynamoDB::Table", data, meta)
}

func resourceDynamoDBTableCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

