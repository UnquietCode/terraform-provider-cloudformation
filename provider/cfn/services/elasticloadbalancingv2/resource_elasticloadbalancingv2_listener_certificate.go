// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticLoadBalancingV2ListenerCertificateType string = "AWS::ElasticLoadBalancingV2::ListenerCertificate"

func ResourceElasticLoadBalancingV2ListenerCertificate() *schema.Resource {
	return &schema.Resource{
		Read: resourceElasticLoadBalancingV2ListenerCertificateRead,
		Create: resourceElasticLoadBalancingV2ListenerCertificateCreate,
		Update: resourceElasticLoadBalancingV2ListenerCertificateUpdate,
		Delete: resourceElasticLoadBalancingV2ListenerCertificateDelete,
		CustomizeDiff: resourceElasticLoadBalancingV2ListenerCertificateCustomizeDiff,
		
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
			},
		},
	}
}

func resourceElasticLoadBalancingV2ListenerCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticLoadBalancingV2ListenerCertificateType, ResourceElasticLoadBalancingV2ListenerCertificate(), data, meta)
}

func resourceElasticLoadBalancingV2ListenerCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elasticLoadBalancingV2ListenerCertificateType, ResourceElasticLoadBalancingV2ListenerCertificate(), data, meta)
}

func resourceElasticLoadBalancingV2ListenerCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elasticLoadBalancingV2ListenerCertificateType, ResourceElasticLoadBalancingV2ListenerCertificate(), data, meta)
}

func resourceElasticLoadBalancingV2ListenerCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elasticLoadBalancingV2ListenerCertificateType, data, meta)
}

func resourceElasticLoadBalancingV2ListenerCertificateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elasticLoadBalancingV2ListenerCertificateType, data, meta)
}
