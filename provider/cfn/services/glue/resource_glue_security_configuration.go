// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-securityconfiguration.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueSecurityConfigurationType string = "AWS::Glue::SecurityConfiguration"

var glueSecurityConfigurationProperties map[string]string = map[string]string{
	"encryption_configuration": "EncryptionConfiguration",
	"name": "Name",
}

func ResourceGlueSecurityConfiguration() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGlueSecurityConfigurationExists,
		Read: resourceGlueSecurityConfigurationRead,
		Create: resourceGlueSecurityConfigurationCreate,
		Update: resourceGlueSecurityConfigurationUpdate,
		Delete: resourceGlueSecurityConfigurationDelete,
		CustomizeDiff: resourceGlueSecurityConfigurationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"encryption_configuration": {
				Type: schema.TypeSet,
				Elem: propertySecurityConfigurationEncryptionConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"name": {
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

func resourceGlueSecurityConfigurationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGlueSecurityConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueSecurityConfigurationType, ResourceGlueSecurityConfiguration(), data, meta)
}

func resourceGlueSecurityConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueSecurityConfigurationType, ResourceGlueSecurityConfiguration(), data, glueSecurityConfigurationProperties, meta)
}

func resourceGlueSecurityConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueSecurityConfigurationType, ResourceGlueSecurityConfiguration(), data, glueSecurityConfigurationProperties, meta)
}

func resourceGlueSecurityConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueSecurityConfigurationType, data, meta)
}

func resourceGlueSecurityConfigurationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueSecurityConfigurationType, data, meta)
}
