// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityGroupIngressCreate,
		Read:   resourceSecurityGroupIngressRead,
		Update: resourceSecurityGroupIngressUpdate,
		Delete: resourceSecurityGroupIngressDelete,

		Schema: map[string]*schema.Schema{
			"cache_security_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ec2_security_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ec2_security_group_owner_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceSecurityGroupIngressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElastiCache::SecurityGroupIngress", data, meta)
}

func resourceSecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElastiCache::SecurityGroupIngress", data, meta)
}

func resourceSecurityGroupIngressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElastiCache::SecurityGroupIngress", data, meta)
}

func resourceSecurityGroupIngressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElastiCache::SecurityGroupIngress", data, meta)
}