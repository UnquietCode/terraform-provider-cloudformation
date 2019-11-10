// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package redshift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceClusterSecurityGroupIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterSecurityGroupIngressCreate,
		Read:   resourceClusterSecurityGroupIngressRead,
		Update: resourceClusterSecurityGroupIngressUpdate,
		Delete: resourceClusterSecurityGroupIngressDelete,

		Schema: map[string]*schema.Schema{
			"cidrip": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"cluster_security_group_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ec2_security_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ec2_security_group_owner_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceClusterSecurityGroupIngressCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Redshift::ClusterSecurityGroupIngress", data, meta)
}

func resourceClusterSecurityGroupIngressRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Redshift::ClusterSecurityGroupIngress", data, meta)
}

func resourceClusterSecurityGroupIngressUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Redshift::ClusterSecurityGroupIngress", data, meta)
}

func resourceClusterSecurityGroupIngressDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Redshift::ClusterSecurityGroupIngress", data, meta)
}