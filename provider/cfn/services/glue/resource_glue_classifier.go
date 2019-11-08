// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGlueClassifier() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlueClassifierCreate,
		Read:   resourceGlueClassifierRead,
		Update: resourceGlueClassifierUpdate,
		Delete: resourceGlueClassifierDelete,

		Schema: map[string]*schema.Schema{
			"xml_classifier": {
				Type: schema.TypeList,
				Elem: propertyXMLClassifier(),
				Required: false,
				MaxItems: 1,
			},
			"json_classifier": {
				Type: schema.TypeList,
				Elem: propertyJsonClassifier(),
				Required: false,
				MaxItems: 1,
			},
			"csv_classifier": {
				Type: schema.TypeList,
				Elem: propertyCsvClassifier(),
				Required: false,
				MaxItems: 1,
			},
			"grok_classifier": {
				Type: schema.TypeList,
				Elem: propertyGrokClassifier(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceGlueClassifierCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Classifier", data, meta)
}

func resourceGlueClassifierRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Classifier", data, meta)
}

func resourceGlueClassifierUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Classifier", data, meta)
}

func resourceGlueClassifierDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Classifier", data, meta)
}