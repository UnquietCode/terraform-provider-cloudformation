// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html

package emr

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyInstanceGroupConfigEbsBlockDeviceConfig(extras...string) *schema.Resource {
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
			"volume_specification": {
				Type: schema.TypeList,
				Elem: propertyInstanceGroupConfigVolumeSpecification(),
				Required: true,
				MaxItems: 1,
			},
			"volumes_per_instance": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}