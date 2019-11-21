// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html

package servicecatalog

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogTagOptionAssociationType string = "AWS::ServiceCatalog::TagOptionAssociation"

func ResourceServiceCatalogTagOptionAssociation() *schema.Resource {
	return &schema.Resource{
		Read: resourceServiceCatalogTagOptionAssociationRead,
		Create: resourceServiceCatalogTagOptionAssociationCreate,
		Update: resourceServiceCatalogTagOptionAssociationUpdate,
		Delete: resourceServiceCatalogTagOptionAssociationDelete,
		CustomizeDiff: resourceServiceCatalogTagOptionAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"tag_option_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_id": {
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

func resourceServiceCatalogTagOptionAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogTagOptionAssociationType, ResourceServiceCatalogTagOptionAssociation(), data, meta)
}

func resourceServiceCatalogTagOptionAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogTagOptionAssociationType, ResourceServiceCatalogTagOptionAssociation(), data, meta)
}

func resourceServiceCatalogTagOptionAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogTagOptionAssociationType, ResourceServiceCatalogTagOptionAssociation(), data, meta)
}

func resourceServiceCatalogTagOptionAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogTagOptionAssociationType, data, meta)
}

func resourceServiceCatalogTagOptionAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogTagOptionAssociationType, data, meta)
}
