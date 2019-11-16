// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogPortfolioPrincipalAssociationType string = "AWS::ServiceCatalog::PortfolioPrincipalAssociation"

var serviceCatalogPortfolioPrincipalAssociationProperties map[string]string = map[string]string{
	"principal_arn": "PrincipalARN",
	"accept_language": "AcceptLanguage",
	"portfolio_id": "PortfolioId",
	"principal_type": "PrincipalType",
}

func ResourceServiceCatalogPortfolioPrincipalAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogPortfolioPrincipalAssociationExists,
		Read: resourceServiceCatalogPortfolioPrincipalAssociationRead,
		Create: resourceServiceCatalogPortfolioPrincipalAssociationCreate,
		Update: resourceServiceCatalogPortfolioPrincipalAssociationUpdate,
		Delete: resourceServiceCatalogPortfolioPrincipalAssociationDelete,
		CustomizeDiff: resourceServiceCatalogPortfolioPrincipalAssociationCustomizeDiff,
		
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
	return plugin.ResourceRead(serviceCatalogPortfolioPrincipalAssociationType, ResourceServiceCatalogPortfolioPrincipalAssociation(), data, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogPortfolioPrincipalAssociationType, ResourceServiceCatalogPortfolioPrincipalAssociation(), data, serviceCatalogPortfolioPrincipalAssociationProperties, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogPortfolioPrincipalAssociationType, ResourceServiceCatalogPortfolioPrincipalAssociation(), data, serviceCatalogPortfolioPrincipalAssociationProperties, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogPortfolioPrincipalAssociationType, data, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogPortfolioPrincipalAssociationType, data, meta)
}
