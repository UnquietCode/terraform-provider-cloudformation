// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-resourceupdateconstraint.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogResourceUpdateConstraint() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogResourceUpdateConstraintExists,
		Read: resourceServiceCatalogResourceUpdateConstraintRead,
		Create: resourceServiceCatalogResourceUpdateConstraintCreate,
		Update: resourceServiceCatalogResourceUpdateConstraintUpdate,
		Delete: resourceServiceCatalogResourceUpdateConstraintDelete,
		CustomizeDiff: resourceServiceCatalogResourceUpdateConstraintCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tag_update_on_provisioned_product": {
				Type: schema.TypeString,
				Required: true,
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

func resourceServiceCatalogResourceUpdateConstraintExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::ResourceUpdateConstraint", ResourceServiceCatalogResourceUpdateConstraint(), data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::ResourceUpdateConstraint", ResourceServiceCatalogResourceUpdateConstraint(), data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::ResourceUpdateConstraint", ResourceServiceCatalogResourceUpdateConstraint(), data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::ResourceUpdateConstraint", data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

