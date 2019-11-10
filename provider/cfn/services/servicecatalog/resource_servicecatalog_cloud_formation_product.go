// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceServiceCatalogCloudFormationProductCreate,
		Read:   resourceServiceCatalogCloudFormationProductRead,
		Update: resourceServiceCatalogCloudFormationProductUpdate,
		Delete: resourceServiceCatalogCloudFormationProductDelete,

		Schema: map[string]*schema.Schema{
			"product_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"provisioning_artifact_ids": {
				Type: schema.TypeString,
				Computed: true,
			},
			"provisioning_artifact_names": {
				Type: schema.TypeString,
				Computed: true,
			},
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
		},
	}
}

func resourceServiceCatalogCloudFormationProductCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::CloudFormationProduct", data, meta)
}

func resourceServiceCatalogCloudFormationProductRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::CloudFormationProduct", data, meta)
}

func resourceServiceCatalogCloudFormationProductUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::CloudFormationProduct", data, meta)
}

func resourceServiceCatalogCloudFormationProductDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::CloudFormationProduct", data, meta)
}