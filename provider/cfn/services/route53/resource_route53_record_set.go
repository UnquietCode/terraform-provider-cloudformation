// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const route53RecordSetType string = "AWS::Route53::RecordSet"

var route53RecordSetProperties map[string]string = map[string]string{
	"alias_target": "AliasTarget",
	"comment": "Comment",
	"failover": "Failover",
	"geo_location": "GeoLocation",
	"health_check_id": "HealthCheckId",
	"hosted_zone_id": "HostedZoneId",
	"hosted_zone_name": "HostedZoneName",
	"multi_value_answer": "MultiValueAnswer",
	"name": "Name",
	"region": "Region",
	"resource_records": "ResourceRecords",
	"set_identifier": "SetIdentifier",
	"ttl": "TTL",
	"type": "Type",
	"weight": "Weight",
}

func ResourceRoute53RecordSet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRoute53RecordSetExists,
		Read: resourceRoute53RecordSetRead,
		Create: resourceRoute53RecordSetCreate,
		Update: resourceRoute53RecordSetUpdate,
		Delete: resourceRoute53RecordSetDelete,
		CustomizeDiff: resourceRoute53RecordSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"alias_target": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
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
			},
			"hosted_zone_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"multi_value_answer": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
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

func resourceRoute53RecordSetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRoute53RecordSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53RecordSetType, ResourceRoute53RecordSet(), data, meta)
}

func resourceRoute53RecordSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(route53RecordSetType, ResourceRoute53RecordSet(), data, route53RecordSetProperties, meta)
}

func resourceRoute53RecordSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(route53RecordSetType, ResourceRoute53RecordSet(), data, route53RecordSetProperties, meta)
}

func resourceRoute53RecordSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(route53RecordSetType, data, meta)
}

func resourceRoute53RecordSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(route53RecordSetType, data, meta)
}
