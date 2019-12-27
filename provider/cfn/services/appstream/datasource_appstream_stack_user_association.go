// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html

package appstream

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appStreamStackUserAssociationType string = "AWS::AppStream::StackUserAssociation"

func DatasourceAppStreamStackUserAssociation() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAppStreamStackUserAssociationRead,
		
		Schema: map[string]*schema.Schema{
			"send_email_notification": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"user_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"stack_name": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceAppStreamStackUserAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appStreamStackUserAssociationType, DatasourceAppStreamStackUserAssociation(), data, meta)
}
