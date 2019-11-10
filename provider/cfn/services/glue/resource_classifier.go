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

func ResourceClassifier() *schema.Resource {
	return &schema.Resource{
		Create: resourceClassifierCreate,
		Read:   resourceClassifierRead,
		Update: resourceClassifierUpdate,
		Delete: resourceClassifierDelete,

		Schema: map[string]*schema.Schema{
			"xml_classifier": {
				Type: schema.TypeList,
				Elem: propertyClassifierXMLClassifier(),
				Required: false,
				MaxItems: 1,
			},
			"json_classifier": {
				Type: schema.TypeList,
				Elem: propertyClassifierJsonClassifier(),
				Required: false,
				MaxItems: 1,
			},
			"csv_classifier": {
				Type: schema.TypeList,
				Elem: propertyClassifierCsvClassifier(),
				Required: false,
				MaxItems: 1,
			},
			"grok_classifier": {
				Type: schema.TypeList,
				Elem: propertyClassifierGrokClassifier(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceClassifierCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Classifier", data, meta)
}

func resourceClassifierRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Classifier", data, meta)
}

func resourceClassifierUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Classifier", data, meta)
}

func resourceClassifierDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Classifier", data, meta)
}