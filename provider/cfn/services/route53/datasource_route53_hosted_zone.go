// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceRoute53HostedZone() *schema.Resource {
	return &schema.Resource{
		Read: datasourceRoute53HostedZoneRead,
		
		Schema: map[string]*schema.Schema{
			"hosted_zone_config": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceRoute53HostedZoneRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53HostedZoneType, DatasourceRoute53HostedZone(), data, meta)
}
