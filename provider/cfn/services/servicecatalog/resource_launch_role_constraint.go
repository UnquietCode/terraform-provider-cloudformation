// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLaunchRoleConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceLaunchRoleConstraintCreate,
		Read:   resourceLaunchRoleConstraintRead,
		Update: resourceLaunchRoleConstraintUpdate,
		Delete: resourceLaunchRoleConstraintDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"accept_language": {
				Type: schema.TypeString,
				Required: false,
			},
			"portfolio_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"product_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceLaunchRoleConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::LaunchRoleConstraint", data, meta)
}

func resourceLaunchRoleConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::LaunchRoleConstraint", data, meta)
}

func resourceLaunchRoleConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::LaunchRoleConstraint", data, meta)
}

func resourceLaunchRoleConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::LaunchRoleConstraint", data, meta)
}