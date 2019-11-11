// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
				Optional: true,
				MaxItems: 1,
			},
			"comment": {
				Type: schema.TypeString,
				Optional: true,
			},
			"failover": {
				Type: schema.TypeString,
				Optional: true,
			},
			"geo_location": {
				Type: schema.TypeList,
				Elem: propertyRecordSetGeoLocation(),
				Optional: true,
				MaxItems: 1,
			},
			"health_check_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"hosted_zone_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"hosted_zone_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"multi_value_answer": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region": {
				Type: schema.TypeString,
				Optional: true,
			},
			"resource_records": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"set_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ttl": {
				Type: schema.TypeString,
				Optional: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"weight": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRoute53RecordSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53::RecordSet", ResourceRoute53RecordSet(), data, meta)
}

func resourceRoute53RecordSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53::RecordSet", ResourceRoute53RecordSet(), data, meta)
}

func resourceRoute53RecordSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53::RecordSet", ResourceRoute53RecordSet(), data, meta)
}

func resourceRoute53RecordSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53::RecordSet", data, meta)
}