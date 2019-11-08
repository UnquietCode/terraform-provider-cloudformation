// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyActivity() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"select_attributes": {
				Type: schema.TypeList,
				Elem: propertySelectAttributes(),
				Required: false,
				MaxItems: 1,
			},
			"datastore": {
				Type: schema.TypeList,
				Elem: propertyDatastore(),
				Required: false,
				MaxItems: 1,
			},
			"filter": {
				Type: schema.TypeList,
				Elem: propertyFilter(),
				Required: false,
				MaxItems: 1,
			},
			"add_attributes": {
				Type: schema.TypeList,
				Elem: propertyAddAttributes(),
				Required: false,
				MaxItems: 1,
			},
			"channel": {
				Type: schema.TypeList,
				Elem: propertyChannel(),
				Required: false,
				MaxItems: 1,
			},
			"device_shadow_enrich": {
				Type: schema.TypeList,
				Elem: propertyDeviceShadowEnrich(),
				Required: false,
				MaxItems: 1,
			},
			"math": {
				Type: schema.TypeList,
				Elem: propertyMath(),
				Required: false,
				MaxItems: 1,
			},
			"lambda": {
				Type: schema.TypeList,
				Elem: propertyLambda(),
				Required: false,
				MaxItems: 1,
			},
			"device_registry_enrich": {
				Type: schema.TypeList,
				Elem: propertyDeviceRegistryEnrich(),
				Required: false,
				MaxItems: 1,
			},
			"remove_attributes": {
				Type: schema.TypeList,
				Elem: propertyRemoveAttributes(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}