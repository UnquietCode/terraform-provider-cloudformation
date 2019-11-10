// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iotanalytics

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyPipelineActivity(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"select_attributes": {
				Type: schema.TypeList,
				Elem: propertyPipelineSelectAttributes(),
				Required: false,
				MaxItems: 1,
			},
			"datastore": {
				Type: schema.TypeList,
				Elem: propertyPipelineDatastore(),
				Required: false,
				MaxItems: 1,
			},
			"filter": {
				Type: schema.TypeList,
				Elem: propertyPipelineFilter(),
				Required: false,
				MaxItems: 1,
			},
			"add_attributes": {
				Type: schema.TypeList,
				Elem: propertyPipelineAddAttributes(),
				Required: false,
				MaxItems: 1,
			},
			"channel": {
				Type: schema.TypeList,
				Elem: propertyPipelineChannel(),
				Required: false,
				MaxItems: 1,
			},
			"device_shadow_enrich": {
				Type: schema.TypeList,
				Elem: propertyPipelineDeviceShadowEnrich(),
				Required: false,
				MaxItems: 1,
			},
			"math": {
				Type: schema.TypeList,
				Elem: propertyPipelineMath(),
				Required: false,
				MaxItems: 1,
			},
			"lambda": {
				Type: schema.TypeList,
				Elem: propertyPipelineLambda(),
				Required: false,
				MaxItems: 1,
			},
			"device_registry_enrich": {
				Type: schema.TypeList,
				Elem: propertyPipelineDeviceRegistryEnrich(),
				Required: false,
				MaxItems: 1,
			},
			"remove_attributes": {
				Type: schema.TypeList,
				Elem: propertyPipelineRemoveAttributes(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}