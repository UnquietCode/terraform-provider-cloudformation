// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceServiceCatalogCloudFormationProvisionedProduct() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCatalogCloudFormationProvisionedProductCreate,
		Read:   resourceServiceCatalogCloudFormationProvisionedProductRead,
		Update: resourceServiceCatalogCloudFormationProvisionedProductUpdate,
		Delete: resourceServiceCatalogCloudFormationProvisionedProductDelete,

		Schema: map[string]*schema.Schema{
			"path_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"provisioning_parameters": {
				Type: schema.TypeList,
				Elem: propertyCloudFormationProvisionedProductProvisioningParameter(),
				Required: false,
			},
			"provisioning_preferences": {
				Type: schema.TypeList,
				Elem: propertyCloudFormationProvisionedProductProvisioningPreferences(),
				Required: false,
				MaxItems: 1,
			},
			"product_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"provisioning_artifact_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"notification_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Required: false,
			},
			"product_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"provisioned_product_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"provisioning_artifact_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceServiceCatalogCloudFormationProvisionedProductCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::CloudFormationProvisionedProduct", data, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::CloudFormationProvisionedProduct", data, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::CloudFormationProvisionedProduct", data, meta)
}

func resourceServiceCatalogCloudFormationProvisionedProductDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::CloudFormationProvisionedProduct", data, meta)
}