// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSecurityConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityConfigurationCreate,
		Read:   resourceSecurityConfigurationRead,
		Update: resourceSecurityConfigurationUpdate,
		Delete: resourceSecurityConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"encryption_configuration": {
				Type: schema.TypeList,
				Elem: propertySecurityConfigurationEncryptionConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSecurityConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::SecurityConfiguration", data, meta)
}

func resourceSecurityConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::SecurityConfiguration", data, meta)
}

func resourceSecurityConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::SecurityConfiguration", data, meta)
}

func resourceSecurityConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::SecurityConfiguration", data, meta)
}