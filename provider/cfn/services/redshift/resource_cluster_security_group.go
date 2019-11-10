// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package redshift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceClusterSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterSecurityGroupCreate,
		Read:   resourceClusterSecurityGroupRead,
		Update: resourceClusterSecurityGroupUpdate,
		Delete: resourceClusterSecurityGroupDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceClusterSecurityGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Redshift::ClusterSecurityGroup", data, meta)
}

func resourceClusterSecurityGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Redshift::ClusterSecurityGroup", data, meta)
}

func resourceClusterSecurityGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Redshift::ClusterSecurityGroup", data, meta)
}

func resourceClusterSecurityGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Redshift::ClusterSecurityGroup", data, meta)
}