// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchnotificationconstraint.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogLaunchNotificationConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCatalogLaunchNotificationConstraintCreate,
		Read:   resourceServiceCatalogLaunchNotificationConstraintRead,
		Update: resourceServiceCatalogLaunchNotificationConstraintUpdate,
		Delete: resourceServiceCatalogLaunchNotificationConstraintDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notification_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Optional: true,
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
		},
	}
}

func resourceServiceCatalogLaunchNotificationConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::LaunchNotificationConstraint", ResourceServiceCatalogLaunchNotificationConstraint(), data, meta)
}

func resourceServiceCatalogLaunchNotificationConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::LaunchNotificationConstraint", ResourceServiceCatalogLaunchNotificationConstraint(), data, meta)
}

func resourceServiceCatalogLaunchNotificationConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::LaunchNotificationConstraint", ResourceServiceCatalogLaunchNotificationConstraint(), data, meta)
}

func resourceServiceCatalogLaunchNotificationConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::LaunchNotificationConstraint", data, meta)
}