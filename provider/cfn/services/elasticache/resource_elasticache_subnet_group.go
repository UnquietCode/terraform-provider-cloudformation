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

func ResourceElastiCacheSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceElastiCacheSubnetGroupCreate,
		Read:   resourceElastiCacheSubnetGroupRead,
		Update: resourceElastiCacheSubnetGroupUpdate,
		Delete: resourceElastiCacheSubnetGroupDelete,

		Schema: map[string]*schema.Schema{
			"cache_subnet_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
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
		},
	}
}

func resourceElastiCacheSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElastiCache::SubnetGroup", data, meta)
}

func resourceElastiCacheSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElastiCache::SubnetGroup", data, meta)
}

func resourceElastiCacheSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElastiCache::SubnetGroup", data, meta)
}

func resourceElastiCacheSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElastiCache::SubnetGroup", data, meta)
}