// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package dynamodb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDynamoDBTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceDynamoDBTableCreate,
		Read:   resourceDynamoDBTableRead,
		Update: resourceDynamoDBTableUpdate,
		Delete: resourceDynamoDBTableDelete,

		Schema: map[string]*schema.Schema{
			"attribute_definitions": {
				Type: schema.TypeList,
				Elem: propertyAttributeDefinition(),
				Required: false,
			},
			"billing_mode": {
				Type: schema.TypeString,
				Required: false,
			},
			"global_secondary_indexes": {
				Type: schema.TypeList,
				Elem: propertyGlobalSecondaryIndex(),
				Required: false,
			},
			"key_schema": {
				Type: schema.TypeSet,
				Elem: propertyKeySchema(),
				Required: true,
				ForceNew: true,
			},
			"local_secondary_indexes": {
				Type: schema.TypeList,
				Elem: propertyLocalSecondaryIndex(),
				Required: false,
				ForceNew: true,
			},
			"point_in_time_recovery_specification": {
				Type: schema.TypeList,
				Elem: propertyPointInTimeRecoverySpecification(),
				Required: false,
				MaxItems: 1,
			},
			"provisioned_throughput": {
				Type: schema.TypeList,
				Elem: propertyProvisionedThroughput(),
				Required: false,
				MaxItems: 1,
			},
			"sse_specification": {
				Type: schema.TypeList,
				Elem: propertySSESpecification(),
				Required: false,
				MaxItems: 1,
			},
			"stream_specification": {
				Type: schema.TypeList,
				Elem: propertyStreamSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"table_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"time_to_live_specification": {
				Type: schema.TypeList,
				Elem: propertyTimeToLiveSpecification(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceDynamoDBTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DynamoDB::Table", data, meta)
}

func resourceDynamoDBTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DynamoDB::Table", data, meta)
}

func resourceDynamoDBTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DynamoDB::Table", data, meta)
}

func resourceDynamoDBTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DynamoDB::Table", data, meta)
}