// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html

package opsworks

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const opsWorksElasticLoadBalancerAttachmentType string = "AWS::OpsWorks::ElasticLoadBalancerAttachment"

func ResourceOpsWorksElasticLoadBalancerAttachment() *schema.Resource {
	return &schema.Resource{
		Read: resourceOpsWorksElasticLoadBalancerAttachmentRead,
		Create: resourceOpsWorksElasticLoadBalancerAttachmentCreate,
		Update: resourceOpsWorksElasticLoadBalancerAttachmentUpdate,
		Delete: resourceOpsWorksElasticLoadBalancerAttachmentDelete,
		CustomizeDiff: resourceOpsWorksElasticLoadBalancerAttachmentCustomizeDiff,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceOpsWorksElasticLoadBalancerAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(opsWorksElasticLoadBalancerAttachmentType, ResourceOpsWorksElasticLoadBalancerAttachment(), data, meta)
}

func resourceOpsWorksElasticLoadBalancerAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(opsWorksElasticLoadBalancerAttachmentType, ResourceOpsWorksElasticLoadBalancerAttachment(), data, meta)
}

func resourceOpsWorksElasticLoadBalancerAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(opsWorksElasticLoadBalancerAttachmentType, ResourceOpsWorksElasticLoadBalancerAttachment(), data, meta)
}

func resourceOpsWorksElasticLoadBalancerAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(opsWorksElasticLoadBalancerAttachmentType, data, meta)
}

func resourceOpsWorksElasticLoadBalancerAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(opsWorksElasticLoadBalancerAttachmentType, data, meta)
}
