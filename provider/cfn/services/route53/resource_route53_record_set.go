// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoute53RecordSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoute53RecordSetCreate,
		Read:   resourceRoute53RecordSetRead,
		Update: resourceRoute53RecordSetUpdate,
		Delete: resourceRoute53RecordSetDelete,

		Schema: map[string]*schema.Schema{
			"alias_target": {
				Type: schema.TypeList,
				Elem: propertyRecordSetAliasTarget(),
				Required: false,
				MaxItems: 1,
			},
			"comment": {
				Type: schema.TypeString,
				Required: false,
			},
			"failover": {
				Type: schema.TypeString,
				Required: false,
			},
			"geo_location": {
				Type: schema.TypeList,
				Elem: propertyRecordSetGeoLocation(),
				Required: false,
				MaxItems: 1,
			},
			"health_check_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"hosted_zone_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"hosted_zone_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"multi_value_answer": {
				Type: schema.TypeBool,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region": {
				Type: schema.TypeString,
				Required: false,
			},
			"resource_records": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"set_identifier": {
				Type: schema.TypeString,
				Required: false,
			},
			"ttl": {
				Type: schema.TypeString,
				Required: false,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"weight": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}

func resourceRoute53RecordSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53::RecordSet", data, meta)
}

func resourceRoute53RecordSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53::RecordSet", data, meta)
}

func resourceRoute53RecordSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53::RecordSet", data, meta)
}

func resourceRoute53RecordSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53::RecordSet", data, meta)
}