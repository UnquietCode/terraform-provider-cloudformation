// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-sourcecredential.html

package codebuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeBuildSourceCredentialType string = "AWS::CodeBuild::SourceCredential"

var codeBuildSourceCredentialProperties map[string]string = map[string]string{
	"server_type": "ServerType",
	"username": "Username",
	"token": "Token",
	"auth_type": "AuthType",
}

func ResourceCodeBuildSourceCredential() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCodeBuildSourceCredentialExists,
		Read: resourceCodeBuildSourceCredentialRead,
		Create: resourceCodeBuildSourceCredentialCreate,
		Update: resourceCodeBuildSourceCredentialUpdate,
		Delete: resourceCodeBuildSourceCredentialDelete,
		CustomizeDiff: resourceCodeBuildSourceCredentialCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"server_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"username": {
				Type: schema.TypeString,
				Optional: true,
			},
			"token": {
				Type: schema.TypeString,
				Required: true,
			},
			"auth_type": {
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

func resourceCodeBuildSourceCredentialExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCodeBuildSourceCredentialRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeBuildSourceCredentialType, ResourceCodeBuildSourceCredential(), data, meta)
}

func resourceCodeBuildSourceCredentialCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codeBuildSourceCredentialType, ResourceCodeBuildSourceCredential(), data, codeBuildSourceCredentialProperties, meta)
}

func resourceCodeBuildSourceCredentialUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codeBuildSourceCredentialType, ResourceCodeBuildSourceCredential(), data, codeBuildSourceCredentialProperties, meta)
}

func resourceCodeBuildSourceCredentialDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codeBuildSourceCredentialType, data, meta)
}

func resourceCodeBuildSourceCredentialCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codeBuildSourceCredentialType, data, meta)
}
