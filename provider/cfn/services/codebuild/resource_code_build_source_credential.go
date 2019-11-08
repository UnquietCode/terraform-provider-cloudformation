// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package codebuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodeBuildSourceCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceCodeBuildSourceCredentialCreate,
		Read:   resourceCodeBuildSourceCredentialRead,
		Update: resourceCodeBuildSourceCredentialUpdate,
		Delete: resourceCodeBuildSourceCredentialDelete,

		Schema: map[string]*schema.Schema{
			"server_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": {
				Type: schema.TypeString,
				Required: false,
			},
			"token": {
				Type: schema.TypeString,
				Required: true,
			},
			"auth_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCodeBuildSourceCredentialCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodeBuild::SourceCredential", data, meta)
}

func resourceCodeBuildSourceCredentialRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodeBuild::SourceCredential", data, meta)
}

func resourceCodeBuildSourceCredentialUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodeBuild::SourceCredential", data, meta)
}

func resourceCodeBuildSourceCredentialDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodeBuild::SourceCredential", data, meta)
}