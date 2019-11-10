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

func ResourceCloudFormationProduct() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFormationProductCreate,
		Read:   resourceCloudFormationProductRead,
		Update: resourceCloudFormationProductUpdate,
		Delete: resourceCloudFormationProductDelete,

		Schema: map[string]*schema.Schema{
			"owner": {
				Type: schema.TypeString,
				Required: true,
			},
			"support_description": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"distributor": {
				Type: schema.TypeString,
				Required: false,
			},
			"support_email": {
				Type: schema.TypeString,
				Required: false,
			},
			"accept_language": {
				Type: schema.TypeString,
				Required: false,
			},
			"support_url": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
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

func resourceCloudFormationProductCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::CloudFormationProduct", data, meta)
}

func resourceCloudFormationProductRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::CloudFormationProduct", data, meta)
}

func resourceCloudFormationProductUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::CloudFormationProduct", data, meta)
}

func resourceCloudFormationProductDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::CloudFormationProduct", data, meta)
}