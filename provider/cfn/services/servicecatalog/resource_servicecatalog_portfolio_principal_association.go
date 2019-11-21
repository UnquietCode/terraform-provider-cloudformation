// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html

package servicecatalog

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogPortfolioPrincipalAssociationType string = "AWS::ServiceCatalog::PortfolioPrincipalAssociation"

func ResourceServiceCatalogPortfolioPrincipalAssociation() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceServiceCatalogPortfolioPrincipalAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogPortfolioPrincipalAssociationType, ResourceServiceCatalogPortfolioPrincipalAssociation(), data, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogPortfolioPrincipalAssociationType, ResourceServiceCatalogPortfolioPrincipalAssociation(), data, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogPortfolioPrincipalAssociationType, ResourceServiceCatalogPortfolioPrincipalAssociation(), data, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogPortfolioPrincipalAssociationType, data, meta)
}

func resourceServiceCatalogPortfolioPrincipalAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogPortfolioPrincipalAssociationType, data, meta)
}
