// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dynamodb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceTableCreate,
		Read:   resourceTableRead,
		Update: resourceTableUpdate,
		Delete: resourceTableDelete,

		Schema: map[string]*schema.Schema{
			"attribute_definitions": {
				Type: schema.TypeList,
				Elem: propertyTableAttributeDefinition(),
				Required: false,
			},
			"billing_mode": {
				Type: schema.TypeString,
				Required: false,
			},
			"global_secondary_indexes": {
				Type: schema.TypeList,
				Elem: propertyTableGlobalSecondaryIndex(),
				Required: false,
			},
			"key_schema": {
				Type: schema.TypeSet,
				Elem: propertyTableKeySchema(),
				Required: true,
				ForceNew: true,
			},
			"local_secondary_indexes": {
				Type: schema.TypeList,
				Elem: propertyTableLocalSecondaryIndex(),
				Required: false,
				ForceNew: true,
			},
			"point_in_time_recovery_specification": {
				Type: schema.TypeList,
				Elem: propertyTablePointInTimeRecoverySpecification(),
				Required: false,
				MaxItems: 1,
			},
			"provisioned_throughput": {
				Type: schema.TypeList,
				Elem: propertyTableProvisionedThroughput(),
				Required: false,
				MaxItems: 1,
			},
			"sse_specification": {
				Type: schema.TypeList,
				Elem: propertyTableSSESpecification(),
				Required: false,
				MaxItems: 1,
			},
			"stream_specification": {
				Type: schema.TypeList,
				Elem: propertyTableStreamSpecification(),
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
				Elem: propertyTableTimeToLiveSpecification(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceTableCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DynamoDB::Table", data, meta)
}

func resourceTableRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DynamoDB::Table", data, meta)
}

func resourceTableUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DynamoDB::Table", data, meta)
}

func resourceTableDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DynamoDB::Table", data, meta)
}