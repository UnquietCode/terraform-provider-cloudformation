// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html

package config

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var configurationRecorderRecordingGroupProperties map[string]string = map[string]string{
	"all_supported": "AllSupported",
	"include_global_resource_types": "IncludeGlobalResourceTypes",
	"resource_types": "ResourceTypes",
}

func propertyConfigurationRecorderRecordingGroup(extras...string) *schema.Resource {
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
			"all_supported": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"include_global_resource_types": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"resource_types": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
		},
	}
}
