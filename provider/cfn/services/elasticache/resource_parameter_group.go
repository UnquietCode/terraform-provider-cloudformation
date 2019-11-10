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

func ResourceParameterGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceParameterGroupCreate,
		Read:   resourceParameterGroupRead,
		Update: resourceParameterGroupUpdate,
		Delete: resourceParameterGroupDelete,

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
				Required: false,
			},
		},
	}
}

func resourceParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElastiCache::ParameterGroup", data, meta)
}

func resourceParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElastiCache::ParameterGroup", data, meta)
}

func resourceParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElastiCache::ParameterGroup", data, meta)
}

func resourceParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElastiCache::ParameterGroup", data, meta)
}