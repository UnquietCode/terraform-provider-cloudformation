// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceHostedZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceHostedZoneCreate,
		Read:   resourceHostedZoneRead,
		Update: resourceHostedZoneUpdate,
		Delete: resourceHostedZoneDelete,

		Schema: map[string]*schema.Schema{
			"hosted_zone_config": {
				Type: schema.TypeList,
				Elem: propertyHostedZoneHostedZoneConfig(),
				Required: false,
				MaxItems: 1,
			},
			"hosted_zone_tags": {
				Type: schema.TypeList,
				Elem: propertyHostedZoneHostedZoneTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"query_logging_config": {
				Type: schema.TypeList,
				Elem: propertyHostedZoneQueryLoggingConfig(),
				Required: false,
				MaxItems: 1,
			},
			"vp_cs": {
				Type: schema.TypeList,
				Elem: propertyHostedZoneVPC(),
				Required: false,
			},
		},
	}
}

func resourceHostedZoneCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53::HostedZone", data, meta)
}

func resourceHostedZoneRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53::HostedZone", data, meta)
}

func resourceHostedZoneUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53::HostedZone", data, meta)
}

func resourceHostedZoneDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53::HostedZone", data, meta)
}