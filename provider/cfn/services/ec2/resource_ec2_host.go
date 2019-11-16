// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2HostType string = "AWS::EC2::Host"

var eC2HostProperties map[string]string = map[string]string{
	"auto_placement": "AutoPlacement",
	"availability_zone": "AvailabilityZone",
	"host_recovery": "HostRecovery",
	"instance_type": "InstanceType",
}

func ResourceEC2Host() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2HostExists,
		Read: resourceEC2HostRead,
		Create: resourceEC2HostCreate,
		Update: resourceEC2HostUpdate,
		Delete: resourceEC2HostDelete,
		CustomizeDiff: resourceEC2HostCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"auto_placement": {
				Type: schema.TypeString,
				Optional: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: true,
			},
			"host_recovery": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
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

func resourceEC2HostExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2HostRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2HostType, ResourceEC2Host(), data, meta)
}

func resourceEC2HostCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2HostType, ResourceEC2Host(), data, eC2HostProperties, meta)
}

func resourceEC2HostUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2HostType, ResourceEC2Host(), data, eC2HostProperties, meta)
}

func resourceEC2HostDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2HostType, data, meta)
}

func resourceEC2HostCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2HostType, data, meta)
}
