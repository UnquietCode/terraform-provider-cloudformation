// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html

package route53

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const route53RecordSetType string = "AWS::Route53::RecordSet"

func DatasourceRoute53RecordSet() *schema.Resource {
	return &schema.Resource{
		Read: datasourceRoute53RecordSetRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceRoute53RecordSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53RecordSetType, DatasourceRoute53RecordSet(), data, meta)
}
