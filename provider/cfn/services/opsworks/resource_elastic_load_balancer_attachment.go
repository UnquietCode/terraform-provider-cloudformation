// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElasticLoadBalancerAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticLoadBalancerAttachmentCreate,
		Read:   resourceElasticLoadBalancerAttachmentRead,
		Update: resourceElasticLoadBalancerAttachmentUpdate,
		Delete: resourceElasticLoadBalancerAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"elastic_load_balancer_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"layer_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceElasticLoadBalancerAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorks::ElasticLoadBalancerAttachment", data, meta)
}

func resourceElasticLoadBalancerAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorks::ElasticLoadBalancerAttachment", data, meta)
}

func resourceElasticLoadBalancerAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorks::ElasticLoadBalancerAttachment", data, meta)
}

func resourceElasticLoadBalancerAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorks::ElasticLoadBalancerAttachment", data, meta)
}