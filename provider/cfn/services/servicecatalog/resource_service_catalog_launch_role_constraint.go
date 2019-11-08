// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogLaunchRoleConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCatalogLaunchRoleConstraintCreate,
		Read:   resourceServiceCatalogLaunchRoleConstraintRead,
		Update: resourceServiceCatalogLaunchRoleConstraintUpdate,
		Delete: resourceServiceCatalogLaunchRoleConstraintDelete,

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

func resourceServiceCatalogLaunchRoleConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::LaunchRoleConstraint", data, meta)
}

func resourceServiceCatalogLaunchRoleConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::LaunchRoleConstraint", data, meta)
}

func resourceServiceCatalogLaunchRoleConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::LaunchRoleConstraint", data, meta)
}

func resourceServiceCatalogLaunchRoleConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::LaunchRoleConstraint", data, meta)
}