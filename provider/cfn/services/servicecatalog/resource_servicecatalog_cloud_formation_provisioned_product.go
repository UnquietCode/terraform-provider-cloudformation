// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html

package servicecatalog

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogCloudFormationProvisionedProductType string = "AWS::ServiceCatalog::CloudFormationProvisionedProduct"

func ResourceServiceCatalogCloudFormationProvisionedProduct() *schema.Resource {
	return &schema.Resource{
		Read: resourceServiceCatalogCloudFormationProvisionedProductRead,
		Create: resourceServiceCatalogCloudFormationProvisionedProductCreate,
		Update: resourceServiceCatalogCloudFormationProvisionedProductUpdate,
		Delete: resourceServiceCatalogCloudFormationProvisionedProductDelete,
		CustomizeDiff: resourceServiceCatalogCloudFormationProvisionedProductCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"path_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"provisioning_parameters": {
				Type: schema.TypeList,
				Elem: propertyCloudFormationProvisionedProductProvisioningParameter(),
				Optional: true,
			},
			"provisioning_preferences": {
				Type: schema.TypeSet,
				Elem: propertyCloudFormationProvisionedProductProvisioningPreferences(),
				Optional: true,
				MaxItems: 1,
			},
			"product_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"provisioning_artifact_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notification_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Optional: true,
			},
			"product_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"provisioned_product_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"provisioning_artifact_id": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceServiceCatalogCloudFormationProvisionedProductRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogCloudFormationProvisionedProductType, ResourceServiceCatalogCloudFormationProvisionedProduct(), data, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogCloudFormationProvisionedProductType, ResourceServiceCatalogCloudFormationProvisionedProduct(), data, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogCloudFormationProvisionedProductType, ResourceServiceCatalogCloudFormationProvisionedProduct(), data, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogCloudFormationProvisionedProductType, data, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogCloudFormationProvisionedProductType, data, meta)
}
