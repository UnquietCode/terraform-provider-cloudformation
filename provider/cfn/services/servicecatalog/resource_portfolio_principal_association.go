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

func ResourcePortfolioPrincipalAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourcePortfolioPrincipalAssociationCreate,
		Read:   resourcePortfolioPrincipalAssociationRead,
		Update: resourcePortfolioPrincipalAssociationUpdate,
		Delete: resourcePortfolioPrincipalAssociationDelete,

		Schema: map[string]*schema.Schema{
			"principal_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"portfolio_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"principal_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePortfolioPrincipalAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::PortfolioPrincipalAssociation", data, meta)
}

func resourcePortfolioPrincipalAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::PortfolioPrincipalAssociation", data, meta)
}

func resourcePortfolioPrincipalAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::PortfolioPrincipalAssociation", data, meta)
}

func resourcePortfolioPrincipalAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::PortfolioPrincipalAssociation", data, meta)
}