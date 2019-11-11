// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMDocument() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSMDocumentCreate,
		Read:   resourceSSMDocumentRead,
		Update: resourceSSMDocumentUpdate,
		Delete: resourceSSMDocumentDelete,

		Schema: map[string]*schema.Schema{
			"content": {
				Type: schema.TypeMap,
				Required: true,
				ForceNew: true,
			},
			"document_type": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSSMDocumentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::Document", ResourceSSMDocument(), data, meta)
}

func resourceSSMDocumentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::Document", ResourceSSMDocument(), data, meta)
}

func resourceSSMDocumentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::Document", ResourceSSMDocument(), data, meta)
}

func resourceSSMDocumentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::Document", data, meta)
}