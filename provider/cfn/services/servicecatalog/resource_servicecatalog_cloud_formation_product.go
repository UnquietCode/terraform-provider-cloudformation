// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogCloudFormationProduct() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogCloudFormationProductExists,
		Read: resourceServiceCatalogCloudFormationProductRead,
		Create: resourceServiceCatalogCloudFormationProductCreate,
		Update: resourceServiceCatalogCloudFormationProductUpdate,
		Delete: resourceServiceCatalogCloudFormationProductDelete,
		CustomizeDiff: resourceServiceCatalogCloudFormationProductCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"owner": {
				Type: schema.TypeString,
				Required: true,
			},
			"support_description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"distributor": {
				Type: schema.TypeString,
				Optional: true,
			},
			"support_email": {
				Type: schema.TypeString,
				Optional: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Optional: true,
			},
			"support_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"provisioning_artifact_parameters": {
				Type: schema.TypeList,
				Elem: propertyCloudFormationProductProvisioningArtifactProperties(),
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceCatalogCloudFormationProductExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogCloudFormationProductRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::CloudFormationProduct", ResourceServiceCatalogCloudFormationProduct(), data, meta)
}

func resourceServiceCatalogCloudFormationProductCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::CloudFormationProduct", ResourceServiceCatalogCloudFormationProduct(), data, meta)
}

func resourceServiceCatalogCloudFormationProductUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::CloudFormationProduct", ResourceServiceCatalogCloudFormationProduct(), data, meta)
}

func resourceServiceCatalogCloudFormationProductDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::CloudFormationProduct", data, meta)
}

func resourceServiceCatalogCloudFormationProductCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
