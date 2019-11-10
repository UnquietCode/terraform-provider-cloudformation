// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceListenerCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceListenerCertificateCreate,
		Read:   resourceListenerCertificateRead,
		Update: resourceListenerCertificateUpdate,
		Delete: resourceListenerCertificateDelete,

		Schema: map[string]*schema.Schema{
			"certificates": {
				Type: schema.TypeSet,
				Elem: propertyListenerCertificateCertificate(),
				Required: true,
				ForceNew: true,
			},
			"listener_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceListenerCertificateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticLoadBalancingV2::ListenerCertificate", data, meta)
}

func resourceListenerCertificateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticLoadBalancingV2::ListenerCertificate", data, meta)
}

func resourceListenerCertificateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticLoadBalancingV2::ListenerCertificate", data, meta)
}

func resourceListenerCertificateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticLoadBalancingV2::ListenerCertificate", data, meta)
}