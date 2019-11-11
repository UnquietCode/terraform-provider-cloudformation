// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceOpsWorksElasticLoadBalancerAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceOpsWorksElasticLoadBalancerAttachmentCreate,
		Read:   resourceOpsWorksElasticLoadBalancerAttachmentRead,
		Update: resourceOpsWorksElasticLoadBalancerAttachmentUpdate,
		Delete: resourceOpsWorksElasticLoadBalancerAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"elastic_load_balancer_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"layer_id": {
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

func resourceOpsWorksElasticLoadBalancerAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorks::ElasticLoadBalancerAttachment", ResourceOpsWorksElasticLoadBalancerAttachment(), data, meta)
}

func resourceOpsWorksElasticLoadBalancerAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorks::ElasticLoadBalancerAttachment", ResourceOpsWorksElasticLoadBalancerAttachment(), data, meta)
}

func resourceOpsWorksElasticLoadBalancerAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorks::ElasticLoadBalancerAttachment", ResourceOpsWorksElasticLoadBalancerAttachment(), data, meta)
}

func resourceOpsWorksElasticLoadBalancerAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorks::ElasticLoadBalancerAttachment", data, meta)
}