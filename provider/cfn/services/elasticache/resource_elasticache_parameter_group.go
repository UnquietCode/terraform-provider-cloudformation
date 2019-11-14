// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElastiCacheParameterGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElastiCacheParameterGroupExists,
		Read: resourceElastiCacheParameterGroupRead,
		Create: resourceElastiCacheParameterGroupCreate,
		Update: resourceElastiCacheParameterGroupUpdate,
		Delete: resourceElastiCacheParameterGroupDelete,
		CustomizeDiff: resourceElastiCacheParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"cache_parameter_group_family": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"properties": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceElastiCacheParameterGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElastiCacheParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElastiCache::ParameterGroup", ResourceElastiCacheParameterGroup(), data, meta)
}

func resourceElastiCacheParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElastiCache::ParameterGroup", ResourceElastiCacheParameterGroup(), data, meta)
}

func resourceElastiCacheParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElastiCache::ParameterGroup", ResourceElastiCacheParameterGroup(), data, meta)
}

func resourceElastiCacheParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElastiCache::ParameterGroup", data, meta)
}

func resourceElastiCacheParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

