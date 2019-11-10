// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRecordSetGroupRecordSet() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alias_target": {
				Type: schema.TypeList,
				Elem: propertyRecordSetGroupAliasTarget(),
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
				Elem: propertyRecordSetGroupGeoLocation(),
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
			},
			"hosted_zone_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"multi_value_answer": {
				Type: schema.TypeBool,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"region": {
				Type: schema.TypeString,
				Required: false,
			},
			"resource_records": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
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