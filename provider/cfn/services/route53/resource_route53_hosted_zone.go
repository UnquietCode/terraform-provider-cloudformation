// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html

package route53

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const route53HostedZoneType string = "AWS::Route53::HostedZone"

func ResourceRoute53HostedZone() *schema.Resource {
	return &schema.Resource{
		Read: resourceRoute53HostedZoneRead,
		Create: resourceRoute53HostedZoneCreate,
		Update: resourceRoute53HostedZoneUpdate,
		Delete: resourceRoute53HostedZoneDelete,
		CustomizeDiff: resourceRoute53HostedZoneCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"hosted_zone_config": {
				Type: schema.TypeSet,
				Elem: propertyHostedZoneHostedZoneConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"hosted_zone_tags": {
				Type: schema.TypeList,
				Elem: propertyHostedZoneHostedZoneTag(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"query_logging_config": {
				Type: schema.TypeSet,
				Elem: propertyHostedZoneQueryLoggingConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"vp_cs": {
				Type: schema.TypeList,
				Elem: propertyHostedZoneVPC(),
				Optional: true,
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

func resourceRoute53HostedZoneRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53HostedZoneType, ResourceRoute53HostedZone(), data, meta)
}

func resourceRoute53HostedZoneCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(route53HostedZoneType, ResourceRoute53HostedZone(), data, meta)
}

func resourceRoute53HostedZoneUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(route53HostedZoneType, ResourceRoute53HostedZone(), data, meta)
}

func resourceRoute53HostedZoneDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(route53HostedZoneType, data, meta)
}

func resourceRoute53HostedZoneCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(route53HostedZoneType, data, meta)
}
