// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEMRSecurityConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceEMRSecurityConfigurationCreate,
		Read:   resourceEMRSecurityConfigurationRead,
		Update: resourceEMRSecurityConfigurationUpdate,
		Delete: resourceEMRSecurityConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"security_configuration": {
				Type: schema.TypeMap,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEMRSecurityConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::SecurityConfiguration", data, meta)
}

func resourceEMRSecurityConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::SecurityConfiguration", data, meta)
}

func resourceEMRSecurityConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::SecurityConfiguration", data, meta)
}

func resourceEMRSecurityConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::SecurityConfiguration", data, meta)
}