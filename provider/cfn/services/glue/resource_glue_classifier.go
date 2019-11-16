// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-classifier.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueClassifierType string = "AWS::Glue::Classifier"

var glueClassifierProperties map[string]string = map[string]string{
	"xml_classifier": "XMLClassifier",
	"json_classifier": "JsonClassifier",
	"csv_classifier": "CsvClassifier",
	"grok_classifier": "GrokClassifier",
}

func ResourceGlueClassifier() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGlueClassifierExists,
		Read: resourceGlueClassifierRead,
		Create: resourceGlueClassifierCreate,
		Update: resourceGlueClassifierUpdate,
		Delete: resourceGlueClassifierDelete,
		CustomizeDiff: resourceGlueClassifierCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"xml_classifier": {
				Type: schema.TypeList,
				Elem: propertyClassifierXMLClassifier(),
				Optional: true,
				MaxItems: 1,
			},
			"json_classifier": {
				Type: schema.TypeList,
				Elem: propertyClassifierJsonClassifier(),
				Optional: true,
				MaxItems: 1,
			},
			"csv_classifier": {
				Type: schema.TypeList,
				Elem: propertyClassifierCsvClassifier(),
				Optional: true,
				MaxItems: 1,
			},
			"grok_classifier": {
				Type: schema.TypeList,
				Elem: propertyClassifierGrokClassifier(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGlueClassifierExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGlueClassifierRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueClassifierType, ResourceGlueClassifier(), data, meta)
}

func resourceGlueClassifierCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueClassifierType, ResourceGlueClassifier(), data, glueClassifierProperties, meta)
}

func resourceGlueClassifierUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueClassifierType, ResourceGlueClassifier(), data, glueClassifierProperties, meta)
}

func resourceGlueClassifierDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueClassifierType, data, meta)
}

func resourceGlueClassifierCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueClassifierType, data, meta)
}
