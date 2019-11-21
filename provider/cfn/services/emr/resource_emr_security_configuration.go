// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html

package emr

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eMRSecurityConfigurationType string = "AWS::EMR::SecurityConfiguration"

func ResourceEMRSecurityConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: resourceEMRSecurityConfigurationRead,
		Create: resourceEMRSecurityConfigurationCreate,
		Update: resourceEMRSecurityConfigurationUpdate,
		Delete: resourceEMRSecurityConfigurationDelete,
		CustomizeDiff: resourceEMRSecurityConfigurationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_configuration": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEMRSecurityConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eMRSecurityConfigurationType, ResourceEMRSecurityConfiguration(), data, meta)
}

func resourceEMRSecurityConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eMRSecurityConfigurationType, ResourceEMRSecurityConfiguration(), data, meta)
}

func resourceEMRSecurityConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eMRSecurityConfigurationType, ResourceEMRSecurityConfiguration(), data, meta)
}

func resourceEMRSecurityConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eMRSecurityConfigurationType, data, meta)
}

func resourceEMRSecurityConfigurationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eMRSecurityConfigurationType, data, meta)
}
