// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group.html

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElastiCacheSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceElastiCacheSecurityGroupCreate,
		Read:   resourceElastiCacheSecurityGroupRead,
		Update: resourceElastiCacheSecurityGroupUpdate,
		Delete: resourceElastiCacheSecurityGroupDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceElastiCacheSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElastiCache::SecurityGroup", ResourceElastiCacheSecurityGroup(), data, meta)
}

func resourceElastiCacheSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElastiCache::SecurityGroup", ResourceElastiCacheSecurityGroup(), data, meta)
}

func resourceElastiCacheSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElastiCache::SecurityGroup", ResourceElastiCacheSecurityGroup(), data, meta)
}

func resourceElastiCacheSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElastiCache::SecurityGroup", data, meta)
}