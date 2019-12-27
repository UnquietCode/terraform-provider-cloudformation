// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html

package appstream

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appStreamUserType string = "AWS::AppStream::User"

func DatasourceAppStreamUser() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAppStreamUserRead,
		
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"first_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"message_action": {
				Type: schema.TypeString,
				Optional: true,
			},
			"last_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authentication_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceAppStreamUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appStreamUserType, DatasourceAppStreamUser(), data, meta)
}
