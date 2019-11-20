// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchnotificationconstraint.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogLaunchNotificationConstraintType string = "AWS::ServiceCatalog::LaunchNotificationConstraint"

func ResourceServiceCatalogLaunchNotificationConstraint() *schema.Resource {
	return &schema.Resource{
		Read: resourceServiceCatalogLaunchNotificationConstraintRead,
		Create: resourceServiceCatalogLaunchNotificationConstraintCreate,
		Update: resourceServiceCatalogLaunchNotificationConstraintUpdate,
		Delete: resourceServiceCatalogLaunchNotificationConstraintDelete,
		CustomizeDiff: resourceServiceCatalogLaunchNotificationConstraintCustomizeDiff,
		
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
			},
			"product_id": {
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

func resourceServiceCatalogLaunchNotificationConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogLaunchNotificationConstraintType, ResourceServiceCatalogLaunchNotificationConstraint(), data, meta)
}

func resourceServiceCatalogLaunchNotificationConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogLaunchNotificationConstraintType, ResourceServiceCatalogLaunchNotificationConstraint(), data, meta)
}

func resourceServiceCatalogLaunchNotificationConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogLaunchNotificationConstraintType, ResourceServiceCatalogLaunchNotificationConstraint(), data, meta)
}

func resourceServiceCatalogLaunchNotificationConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogLaunchNotificationConstraintType, data, meta)
}

func resourceServiceCatalogLaunchNotificationConstraintCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogLaunchNotificationConstraintType, data, meta)
}
