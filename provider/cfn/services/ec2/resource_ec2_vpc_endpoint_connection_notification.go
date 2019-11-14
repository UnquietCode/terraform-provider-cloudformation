// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2VPCEndpointConnectionNotification() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2VPCEndpointConnectionNotificationExists,
		Read: resourceEC2VPCEndpointConnectionNotificationRead,
		Create: resourceEC2VPCEndpointConnectionNotificationCreate,
		Update: resourceEC2VPCEndpointConnectionNotificationUpdate,
		Delete: resourceEC2VPCEndpointConnectionNotificationDelete,
		CustomizeDiff: resourceEC2VPCEndpointConnectionNotificationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"connection_events": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"vpc_endpoint_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"connection_notification_arn": {
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

func resourceEC2VPCEndpointConnectionNotificationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2VPCEndpointConnectionNotificationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCEndpointConnectionNotification", ResourceEC2VPCEndpointConnectionNotification(), data, meta)
}

func resourceEC2VPCEndpointConnectionNotificationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCEndpointConnectionNotification", ResourceEC2VPCEndpointConnectionNotification(), data, meta)
}

func resourceEC2VPCEndpointConnectionNotificationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCEndpointConnectionNotification", ResourceEC2VPCEndpointConnectionNotification(), data, meta)
}

func resourceEC2VPCEndpointConnectionNotificationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCEndpointConnectionNotification", data, meta)
}

func resourceEC2VPCEndpointConnectionNotificationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

