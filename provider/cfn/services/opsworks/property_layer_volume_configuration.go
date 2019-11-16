// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html

package opsworks

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var layerVolumeConfigurationProperties map[string]string = map[string]string{
	"encrypted": "Encrypted",
	"iops": "Iops",
	"mount_point": "MountPoint",
	"number_of_disks": "NumberOfDisks",
	"raid_level": "RaidLevel",
	"size": "Size",
	"volume_type": "VolumeType",
}

func propertyLayerVolumeConfiguration(extras...string) *schema.Resource {
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
			"encrypted": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"iops": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"mount_point": {
				Type: schema.TypeString,
				Optional: true,
			},
			"number_of_disks": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"raid_level": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"volume_type": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
