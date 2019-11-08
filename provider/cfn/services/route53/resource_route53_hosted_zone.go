// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoute53HostedZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoute53HostedZoneCreate,
		Read:   resourceRoute53HostedZoneRead,
		Update: resourceRoute53HostedZoneUpdate,
		Delete: resourceRoute53HostedZoneDelete,

		Schema: map[string]*schema.Schema{
			"hosted_zone_config": {
				Type: schema.TypeList,
				Elem: propertyHostedZoneConfig(),
				Required: false,
				MaxItems: 1,
			},
			"hosted_zone_tags": {
				Type: schema.TypeList,
				Elem: propertyHostedZoneTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"query_logging_config": {
				Type: schema.TypeList,
				Elem: propertyQueryLoggingConfig(),
				Required: false,
				MaxItems: 1,
			},
			"vp_cs": {
				Type: schema.TypeList,
				Elem: propertyVPC(),
				Required: false,
			},
		},
	}
}

func resourceRoute53HostedZoneCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53::HostedZone", data, meta)
}

func resourceRoute53HostedZoneRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53::HostedZone", data, meta)
}

func resourceRoute53HostedZoneUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53::HostedZone", data, meta)
}

func resourceRoute53HostedZoneDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53::HostedZone", data, meta)
}