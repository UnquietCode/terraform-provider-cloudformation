// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const sSMDocumentType string = "AWS::SSM::Document"

var sSMDocumentProperties map[string]string = map[string]string{
	"content": "Content",
	"document_type": "DocumentType",
	"tags": "Tags",
}

func ResourceSSMDocument() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSSMDocumentExists,
		Read: resourceSSMDocumentRead,
		Create: resourceSSMDocumentCreate,
		Update: resourceSSMDocumentUpdate,
		Delete: resourceSSMDocumentDelete,
		CustomizeDiff: resourceSSMDocumentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"content": {
				Type: schema.TypeMap,
				Required: true,
			},
			"document_type": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceSSMDocumentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSSMDocumentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sSMDocumentType, ResourceSSMDocument(), data, meta)
}

func resourceSSMDocumentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sSMDocumentType, ResourceSSMDocument(), data, sSMDocumentProperties, meta)
}

func resourceSSMDocumentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sSMDocumentType, ResourceSSMDocument(), data, sSMDocumentProperties, meta)
}

func resourceSSMDocumentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sSMDocumentType, data, meta)
}

func resourceSSMDocumentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sSMDocumentType, data, meta)
}
