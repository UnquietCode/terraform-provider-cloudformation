// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2SecurityGroupIngressType string = "AWS::EC2::SecurityGroupIngress"

func ResourceEC2SecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2SecurityGroupIngressRead,
		Create: resourceEC2SecurityGroupIngressCreate,
		Update: resourceEC2SecurityGroupIngressUpdate,
		Delete: resourceEC2SecurityGroupIngressDelete,
		CustomizeDiff: resourceEC2SecurityGroupIngressCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"cidr_ip": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cidr_ipv6": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"from_port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"group_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ip_protocol": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_prefix_list_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_security_group_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_security_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_security_group_owner_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"to_port": {
				Type: schema.TypeInt,
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

func resourceEC2SecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2SecurityGroupIngressType, ResourceEC2SecurityGroupIngress(), data, meta)
}

func resourceEC2SecurityGroupIngressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2SecurityGroupIngressType, ResourceEC2SecurityGroupIngress(), data, meta)
}

func resourceEC2SecurityGroupIngressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2SecurityGroupIngressType, ResourceEC2SecurityGroupIngress(), data, meta)
}

func resourceEC2SecurityGroupIngressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2SecurityGroupIngressType, data, meta)
}

func resourceEC2SecurityGroupIngressCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2SecurityGroupIngressType, data, meta)
}
