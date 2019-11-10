// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogTagOptionAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCatalogTagOptionAssociationCreate,
		Read:   resourceServiceCatalogTagOptionAssociationRead,
		Update: resourceServiceCatalogTagOptionAssociationUpdate,
		Delete: resourceServiceCatalogTagOptionAssociationDelete,

		Schema: map[string]*schema.Schema{
			"tag_option_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceCatalogTagOptionAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::TagOptionAssociation", data, meta)
}

func resourceServiceCatalogTagOptionAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::TagOptionAssociation", data, meta)
}

func resourceServiceCatalogTagOptionAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::TagOptionAssociation", data, meta)
}

func resourceServiceCatalogTagOptionAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::TagOptionAssociation", data, meta)
}