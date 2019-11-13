// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogPortfolioPrincipalAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogPortfolioPrincipalAssociationExists,
		Read:   resourceServiceCatalogPortfolioPrincipalAssociationRead,
		Create: resourceServiceCatalogPortfolioPrincipalAssociationCreate,
		Update: resourceServiceCatalogPortfolioPrincipalAssociationUpdate,
		Delete: resourceServiceCatalogPortfolioPrincipalAssociationDelete,
		
		Schema: map[string]*schema.Schema{
			"principal_arn": {
				Type: schema.TypeString,
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
			"principal_type": {
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

func resourceServiceCatalogPortfolioPrincipalAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::PortfolioPrincipalAssociation", ResourceServiceCatalogPortfolioPrincipalAssociation(), data, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::PortfolioPrincipalAssociation", ResourceServiceCatalogPortfolioPrincipalAssociation(), data, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::PortfolioPrincipalAssociation", ResourceServiceCatalogPortfolioPrincipalAssociation(), data, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::PortfolioPrincipalAssociation", data, meta)
}