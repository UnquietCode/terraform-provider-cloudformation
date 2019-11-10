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

func ResourceVPCEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPCEndpointCreate,
		Read:   resourceVPCEndpointRead,
		Update: resourceVPCEndpointUpdate,
		Delete: resourceVPCEndpointDelete,

		Schema: map[string]*schema.Schema{
			"policy_document": {
				Type: schema.TypeMap,
				Required: false,
			},
			"private_dns_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"route_table_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"security_group_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"service_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"vpc_endpoint_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVPCEndpointCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCEndpoint", data, meta)
}

func resourceVPCEndpointRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCEndpoint", data, meta)
}

func resourceVPCEndpointUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCEndpoint", data, meta)
}

func resourceVPCEndpointDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCEndpoint", data, meta)
}