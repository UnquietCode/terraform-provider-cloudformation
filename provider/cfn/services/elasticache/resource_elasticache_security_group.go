// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group.html

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elastiCacheSecurityGroupType string = "AWS::ElastiCache::SecurityGroup"

var elastiCacheSecurityGroupProperties map[string]string = map[string]string{
	"description": "Description",
}

func ResourceElastiCacheSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElastiCacheSecurityGroupExists,
		Read: resourceElastiCacheSecurityGroupRead,
		Create: resourceElastiCacheSecurityGroupCreate,
		Update: resourceElastiCacheSecurityGroupUpdate,
		Delete: resourceElastiCacheSecurityGroupDelete,
		CustomizeDiff: resourceElastiCacheSecurityGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
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

func resourceElastiCacheSecurityGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElastiCacheSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elastiCacheSecurityGroupType, ResourceElastiCacheSecurityGroup(), data, meta)
}

func resourceElastiCacheSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elastiCacheSecurityGroupType, ResourceElastiCacheSecurityGroup(), data, elastiCacheSecurityGroupProperties, meta)
}

func resourceElastiCacheSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elastiCacheSecurityGroupType, ResourceElastiCacheSecurityGroup(), data, elastiCacheSecurityGroupProperties, meta)
}

func resourceElastiCacheSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elastiCacheSecurityGroupType, data, meta)
}

func resourceElastiCacheSecurityGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elastiCacheSecurityGroupType, data, meta)
}
