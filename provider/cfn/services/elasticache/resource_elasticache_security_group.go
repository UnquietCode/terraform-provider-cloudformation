// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group.html

package elasticache

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elastiCacheSecurityGroupType string = "AWS::ElastiCache::SecurityGroup"

func ResourceElastiCacheSecurityGroup() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceElastiCacheSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elastiCacheSecurityGroupType, ResourceElastiCacheSecurityGroup(), data, meta)
}

func resourceElastiCacheSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elastiCacheSecurityGroupType, ResourceElastiCacheSecurityGroup(), data, meta)
}

func resourceElastiCacheSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elastiCacheSecurityGroupType, ResourceElastiCacheSecurityGroup(), data, meta)
}

func resourceElastiCacheSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elastiCacheSecurityGroupType, data, meta)
}

func resourceElastiCacheSecurityGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elastiCacheSecurityGroupType, data, meta)
}
