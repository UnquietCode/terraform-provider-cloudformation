// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceVPCEndpointConnectionNotification() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPCEndpointConnectionNotificationCreate,
		Read:   resourceVPCEndpointConnectionNotificationRead,
		Update: resourceVPCEndpointConnectionNotificationUpdate,
		Delete: resourceVPCEndpointConnectionNotificationDelete,

		Schema: map[string]*schema.Schema{
			"connection_events": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"vpc_endpoint_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"service_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"connection_notification_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVPCEndpointConnectionNotificationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCEndpointConnectionNotification", data, meta)
}

func resourceVPCEndpointConnectionNotificationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCEndpointConnectionNotification", data, meta)
}

func resourceVPCEndpointConnectionNotificationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCEndpointConnectionNotification", data, meta)
}

func resourceVPCEndpointConnectionNotificationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCEndpointConnectionNotification", data, meta)
}