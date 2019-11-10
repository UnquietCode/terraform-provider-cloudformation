// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceVPC() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPCCreate,
		Read:   resourceVPCRead,
		Update: resourceVPCUpdate,
		Delete: resourceVPCDelete,

		Schema: map[string]*schema.Schema{
			"cidr_block": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_dns_hostnames": {
				Type: schema.TypeBool,
				Required: false,
			},
			"enable_dns_support": {
				Type: schema.TypeBool,
				Required: false,
			},
			"instance_tenancy": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceVPCCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPC", data, meta)
}

func resourceVPCRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPC", data, meta)
}

func resourceVPCUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPC", data, meta)
}

func resourceVPCDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPC", data, meta)
}