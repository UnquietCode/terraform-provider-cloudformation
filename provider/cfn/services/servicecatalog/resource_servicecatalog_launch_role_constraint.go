// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogLaunchRoleConstraint() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogLaunchRoleConstraintExists,
		Read: resourceServiceCatalogLaunchRoleConstraintRead,
		Create: resourceServiceCatalogLaunchRoleConstraintCreate,
		Update: resourceServiceCatalogLaunchRoleConstraintUpdate,
		Delete: resourceServiceCatalogLaunchRoleConstraintDelete,
		CustomizeDiff: resourceServiceCatalogLaunchRoleConstraintCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Optional: true,
			},
			"portfolio_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"product_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
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

func resourceServiceCatalogLaunchRoleConstraintExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogLaunchRoleConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::LaunchRoleConstraint", ResourceServiceCatalogLaunchRoleConstraint(), data, meta)
}

func resourceServiceCatalogLaunchRoleConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::LaunchRoleConstraint", ResourceServiceCatalogLaunchRoleConstraint(), data, meta)
}

func resourceServiceCatalogLaunchRoleConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::LaunchRoleConstraint", ResourceServiceCatalogLaunchRoleConstraint(), data, meta)
}

func resourceServiceCatalogLaunchRoleConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::LaunchRoleConstraint", data, meta)
}

func resourceServiceCatalogLaunchRoleConstraintCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::ServiceCatalog::LaunchRoleConstraint", data, meta)
}

