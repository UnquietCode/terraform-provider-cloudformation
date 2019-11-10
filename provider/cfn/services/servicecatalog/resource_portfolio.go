// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePortfolio() *schema.Resource {
	return &schema.Resource{
		Create: resourcePortfolioCreate,
		Read:   resourcePortfolioRead,
		Update: resourcePortfolioUpdate,
		Delete: resourcePortfolioDelete,

		Schema: map[string]*schema.Schema{
			"provider_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"display_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourcePortfolioCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::Portfolio", data, meta)
}

func resourcePortfolioRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::Portfolio", data, meta)
}

func resourcePortfolioUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::Portfolio", data, meta)
}

func resourcePortfolioDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::Portfolio", data, meta)
}