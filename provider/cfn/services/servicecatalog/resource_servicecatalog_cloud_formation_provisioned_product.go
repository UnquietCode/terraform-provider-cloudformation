// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogCloudFormationProvisionedProductType string = "AWS::ServiceCatalog::CloudFormationProvisionedProduct"

var serviceCatalogCloudFormationProvisionedProductProperties map[string]string = map[string]string{
	"path_id": "PathId",
	"provisioning_parameters": "ProvisioningParameters",
	"provisioning_preferences": "ProvisioningPreferences",
	"product_name": "ProductName",
	"provisioning_artifact_name": "ProvisioningArtifactName",
	"notification_arns": "NotificationArns",
	"accept_language": "AcceptLanguage",
	"product_id": "ProductId",
	"tags": "Tags",
	"provisioned_product_name": "ProvisionedProductName",
	"provisioning_artifact_id": "ProvisioningArtifactId",
}

func ResourceServiceCatalogCloudFormationProvisionedProduct() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogCloudFormationProvisionedProductExists,
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
			},
		},
	}
}

func resourceServiceCatalogCloudFormationProvisionedProductExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogCloudFormationProvisionedProductType, ResourceServiceCatalogCloudFormationProvisionedProduct(), data, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogCloudFormationProvisionedProductType, ResourceServiceCatalogCloudFormationProvisionedProduct(), data, serviceCatalogCloudFormationProvisionedProductProperties, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogCloudFormationProvisionedProductType, ResourceServiceCatalogCloudFormationProvisionedProduct(), data, serviceCatalogCloudFormationProvisionedProductProperties, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogCloudFormationProvisionedProductType, data, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogCloudFormationProvisionedProductType, data, meta)
}
