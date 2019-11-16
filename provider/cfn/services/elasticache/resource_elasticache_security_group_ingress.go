// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elastiCacheSecurityGroupIngressType string = "AWS::ElastiCache::SecurityGroupIngress"

func ResourceElastiCacheSecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElastiCacheSecurityGroupIngressExists,
		Read: resourceElastiCacheSecurityGroupIngressRead,
		Create: resourceElastiCacheSecurityGroupIngressCreate,
		Update: resourceElastiCacheSecurityGroupIngressUpdate,
		Delete: resourceElastiCacheSecurityGroupIngressDelete,
		CustomizeDiff: resourceElastiCacheSecurityGroupIngressCustomizeDiff,
		
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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceElastiCacheSecurityGroupIngressExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElastiCacheSecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elastiCacheSecurityGroupIngressType, ResourceElastiCacheSecurityGroupIngress(), data, meta)
}

func resourceElastiCacheSecurityGroupIngressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elastiCacheSecurityGroupIngressType, ResourceElastiCacheSecurityGroupIngress(), data, meta)
}

func resourceElastiCacheSecurityGroupIngressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elastiCacheSecurityGroupIngressType, ResourceElastiCacheSecurityGroupIngress(), data, meta)
}

func resourceElastiCacheSecurityGroupIngressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elastiCacheSecurityGroupIngressType, data, meta)
}

func resourceElastiCacheSecurityGroupIngressCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elastiCacheSecurityGroupIngressType, data, meta)
}
