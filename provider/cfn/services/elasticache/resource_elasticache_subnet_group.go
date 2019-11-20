// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elastiCacheSubnetGroupType string = "AWS::ElastiCache::SubnetGroup"

func ResourceElastiCacheSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceElastiCacheSubnetGroupRead,
		Create: resourceElastiCacheSubnetGroupCreate,
		Update: resourceElastiCacheSubnetGroupUpdate,
		Delete: resourceElastiCacheSubnetGroupDelete,
		CustomizeDiff: resourceElastiCacheSubnetGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"cache_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceElastiCacheSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elastiCacheSubnetGroupType, ResourceElastiCacheSubnetGroup(), data, meta)
}

func resourceElastiCacheSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elastiCacheSubnetGroupType, ResourceElastiCacheSubnetGroup(), data, meta)
}

func resourceElastiCacheSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elastiCacheSubnetGroupType, ResourceElastiCacheSubnetGroup(), data, meta)
}

func resourceElastiCacheSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elastiCacheSubnetGroupType, data, meta)
}

func resourceElastiCacheSubnetGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elastiCacheSubnetGroupType, data, meta)
}
