// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDocument() *schema.Resource {
	return &schema.Resource{
		Create: resourceDocumentCreate,
		Read:   resourceDocumentRead,
		Update: resourceDocumentUpdate,
		Delete: resourceDocumentDelete,

		Schema: map[string]*schema.Schema{
			"content": {
				Type: schema.TypeMap,
				Required: true,
				ForceNew: true,
			},
			"document_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceDocumentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::Document", data, meta)
}

func resourceDocumentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::Document", data, meta)
}

func resourceDocumentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::Document", data, meta)
}

func resourceDocumentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::Document", data, meta)
}