// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cloudwatch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDashboard() *schema.Resource {
	return &schema.Resource{
		Create: resourceDashboardCreate,
		Read:   resourceDashboardRead,
		Update: resourceDashboardUpdate,
		Delete: resourceDashboardDelete,

		Schema: map[string]*schema.Schema{
			"dashboard_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"dashboard_body": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDashboardCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudWatch::Dashboard", data, meta)
}

func resourceDashboardRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudWatch::Dashboard", data, meta)
}

func resourceDashboardUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudWatch::Dashboard", data, meta)
}

func resourceDashboardDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudWatch::Dashboard", data, meta)
}