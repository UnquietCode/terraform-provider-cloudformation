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

func ResourcePortfolioShare() *schema.Resource {
	return &schema.Resource{
		Create: resourcePortfolioShareCreate,
		Read:   resourcePortfolioShareRead,
		Update: resourcePortfolioShareUpdate,
		Delete: resourcePortfolioShareDelete,

		Schema: map[string]*schema.Schema{
			"account_id": {
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
		},
	}
}

func resourcePortfolioShareCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::PortfolioShare", data, meta)
}

func resourcePortfolioShareRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::PortfolioShare", data, meta)
}

func resourcePortfolioShareUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::PortfolioShare", data, meta)
}

func resourcePortfolioShareDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::PortfolioShare", data, meta)
}