// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html

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
	    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"select_attributes": {
				Type: schema.TypeSet,
				Elem: propertyPipelineSelectAttributes(),
				Optional: true,
				MaxItems: 1,
			},
			"datastore": {
				Type: schema.TypeSet,
				Elem: propertyPipelineDatastore(),
				Optional: true,
				MaxItems: 1,
			},
			"filter": {
				Type: schema.TypeSet,
				Elem: propertyPipelineFilter(),
				Optional: true,
				MaxItems: 1,
			},
			"add_attributes": {
				Type: schema.TypeSet,
				Elem: propertyPipelineAddAttributes(),
				Optional: true,
				MaxItems: 1,
			},
			"channel": {
				Type: schema.TypeSet,
				Elem: propertyPipelineChannel(),
				Optional: true,
				MaxItems: 1,
			},
			"device_shadow_enrich": {
				Type: schema.TypeSet,
				Elem: propertyPipelineDeviceShadowEnrich(),
				Optional: true,
				MaxItems: 1,
			},
			"math": {
				Type: schema.TypeSet,
				Elem: propertyPipelineMath(),
				Optional: true,
				MaxItems: 1,
			},
			"lambda": {
				Type: schema.TypeSet,
				Elem: propertyPipelineLambda(),
				Optional: true,
				MaxItems: 1,
			},
			"device_registry_enrich": {
				Type: schema.TypeSet,
				Elem: propertyPipelineDeviceRegistryEnrich(),
				Optional: true,
				MaxItems: 1,
			},
			"remove_attributes": {
				Type: schema.TypeSet,
				Elem: propertyPipelineRemoveAttributes(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
