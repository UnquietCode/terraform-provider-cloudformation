// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueConnectionType string = "AWS::Glue::Connection"

var glueConnectionProperties map[string]string = map[string]string{
	"connection_input": "ConnectionInput",
	"catalog_id": "CatalogId",
}

func ResourceGlueConnection() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGlueConnectionExists,
		Read: resourceGlueConnectionRead,
		Create: resourceGlueConnectionCreate,
		Update: resourceGlueConnectionUpdate,
		Delete: resourceGlueConnectionDelete,
		CustomizeDiff: resourceGlueConnectionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"connection_input": {
				Type: schema.TypeSet,
				Elem: propertyConnectionConnectionInput(),
				Required: true,
				MaxItems: 1,
			},
			"catalog_id": {
				Type: schema.TypeString,
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

func resourceGlueConnectionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGlueConnectionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueConnectionType, ResourceGlueConnection(), data, meta)
}

func resourceGlueConnectionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueConnectionType, ResourceGlueConnection(), data, glueConnectionProperties, meta)
}

func resourceGlueConnectionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueConnectionType, ResourceGlueConnection(), data, glueConnectionProperties, meta)
}

func resourceGlueConnectionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueConnectionType, data, meta)
}

func resourceGlueConnectionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueConnectionType, data, meta)
}
