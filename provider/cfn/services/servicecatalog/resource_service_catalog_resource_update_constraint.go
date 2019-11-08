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

func ResourceServiceCatalogResourceUpdateConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCatalogResourceUpdateConstraintCreate,
		Read:   resourceServiceCatalogResourceUpdateConstraintRead,
		Update: resourceServiceCatalogResourceUpdateConstraintUpdate,
		Delete: resourceServiceCatalogResourceUpdateConstraintDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"accept_language": {
				Type: schema.TypeString,
				Required: false,
			},
			"tag_update_on_provisioned_product": {
				Type: schema.TypeString,
				Required: true,
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

func resourceServiceCatalogResourceUpdateConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::ResourceUpdateConstraint", data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::ResourceUpdateConstraint", data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::ResourceUpdateConstraint", data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::ResourceUpdateConstraint", data, meta)
}