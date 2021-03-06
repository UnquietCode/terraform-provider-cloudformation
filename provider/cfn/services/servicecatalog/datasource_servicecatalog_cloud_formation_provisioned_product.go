// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceServiceCatalogCloudFormationProvisionedProduct() *schema.Resource {
	return &schema.Resource{
		Read: datasourceServiceCatalogCloudFormationProvisionedProductRead,
		
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
				Type: schema.TypeList,
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceServiceCatalogCloudFormationProvisionedProductRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogCloudFormationProvisionedProductType, DatasourceServiceCatalogCloudFormationProvisionedProduct(), data, meta)
}
