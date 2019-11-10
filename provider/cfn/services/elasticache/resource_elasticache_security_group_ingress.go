// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElastiCacheSecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceElastiCacheSecurityGroupIngressCreate,
		Read:   resourceElastiCacheSecurityGroupIngressRead,
		Update: resourceElastiCacheSecurityGroupIngressUpdate,
		Delete: resourceElastiCacheSecurityGroupIngressDelete,

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
				Optional: true,
			},
		},
	}
}

func resourceElastiCacheSecurityGroupIngressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElastiCache::SecurityGroupIngress", data, meta)
}

func resourceElastiCacheSecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElastiCache::SecurityGroupIngress", data, meta)
}

func resourceElastiCacheSecurityGroupIngressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElastiCache::SecurityGroupIngress", data, meta)
}

func resourceElastiCacheSecurityGroupIngressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElastiCache::SecurityGroupIngress", data, meta)
}