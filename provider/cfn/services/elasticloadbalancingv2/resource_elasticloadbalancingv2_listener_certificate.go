// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElasticLoadBalancingV2ListenerCertificate() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElasticLoadBalancingV2ListenerCertificateExists,
		Read:   resourceElasticLoadBalancingV2ListenerCertificateRead,
		Create: resourceElasticLoadBalancingV2ListenerCertificateCreate,
		Update: resourceElasticLoadBalancingV2ListenerCertificateUpdate,
		Delete: resourceElasticLoadBalancingV2ListenerCertificateDelete,
		
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

func resourceElasticLoadBalancingV2ListenerCertificateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElasticLoadBalancingV2ListenerCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticLoadBalancingV2::ListenerCertificate", ResourceElasticLoadBalancingV2ListenerCertificate(), data, meta)
}

func resourceElasticLoadBalancingV2ListenerCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticLoadBalancingV2::ListenerCertificate", ResourceElasticLoadBalancingV2ListenerCertificate(), data, meta)
}

func resourceElasticLoadBalancingV2ListenerCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticLoadBalancingV2::ListenerCertificate", ResourceElasticLoadBalancingV2ListenerCertificate(), data, meta)
}

func resourceElasticLoadBalancingV2ListenerCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticLoadBalancingV2::ListenerCertificate", data, meta)
}