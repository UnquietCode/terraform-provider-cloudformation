// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPNConnectionRouteType string = "AWS::EC2::VPNConnectionRoute"

func ResourceEC2VPNConnectionRoute() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2VPNConnectionRouteExists,
		Read: resourceEC2VPNConnectionRouteRead,
		Create: resourceEC2VPNConnectionRouteCreate,
		Update: resourceEC2VPNConnectionRouteUpdate,
		Delete: resourceEC2VPNConnectionRouteDelete,
		CustomizeDiff: resourceEC2VPNConnectionRouteCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"destination_cidr_block": {
				Type: schema.TypeString,
				Required: true,
			},
			"vpn_connection_id": {
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

func resourceEC2VPNConnectionRouteExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2VPNConnectionRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPNConnectionRouteType, ResourceEC2VPNConnectionRoute(), data, meta)
}

func resourceEC2VPNConnectionRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPNConnectionRouteType, ResourceEC2VPNConnectionRoute(), data, meta)
}

func resourceEC2VPNConnectionRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPNConnectionRouteType, ResourceEC2VPNConnectionRoute(), data, meta)
}

func resourceEC2VPNConnectionRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPNConnectionRouteType, data, meta)
}

func resourceEC2VPNConnectionRouteCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPNConnectionRouteType, data, meta)
}
