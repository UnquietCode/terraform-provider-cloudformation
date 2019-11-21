// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html

package dynamodb

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dynamoDBTableType string = "AWS::DynamoDB::Table"

func ResourceDynamoDBTable() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeSet,
				Elem: propertyTablePointInTimeRecoverySpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"provisioned_throughput": {
				Type: schema.TypeSet,
				Elem: propertyTableProvisionedThroughput(),
				Optional: true,
				MaxItems: 1,
			},
			"sse_specification": {
				Type: schema.TypeSet,
				Elem: propertyTableSSESpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"stream_specification": {
				Type: schema.TypeSet,
				Elem: propertyTableStreamSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"table_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"time_to_live_specification": {
				Type: schema.TypeSet,
				Elem: propertyTableTimeToLiveSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceDynamoDBTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dynamoDBTableType, ResourceDynamoDBTable(), data, meta)
}

func resourceDynamoDBTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dynamoDBTableType, ResourceDynamoDBTable(), data, meta)
}

func resourceDynamoDBTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dynamoDBTableType, ResourceDynamoDBTable(), data, meta)
}

func resourceDynamoDBTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dynamoDBTableType, data, meta)
}

func resourceDynamoDBTableCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dynamoDBTableType, data, meta)
}
