// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html

package elasticloadbalancingv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticLoadBalancingV2ListenerCertificateType string = "AWS::ElasticLoadBalancingV2::ListenerCertificate"

func DatasourceElasticLoadBalancingV2ListenerCertificate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceElasticLoadBalancingV2ListenerCertificateRead,
		
		Schema: map[string]*schema.Schema{
			"certificates": {
				Type: schema.TypeSet,
				Elem: propertyListenerCertificateCertificate(),
				Required: true,
			},
			"listener_arn": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceElasticLoadBalancingV2ListenerCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticLoadBalancingV2ListenerCertificateType, DatasourceElasticLoadBalancingV2ListenerCertificate(), data, meta)
}
