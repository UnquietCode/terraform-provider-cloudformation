// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-classifier.html

package glue

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueClassifierType string = "AWS::Glue::Classifier"

func ResourceGlueClassifier() *schema.Resource {
	return &schema.Resource{
		Read: resourceGlueClassifierRead,
		Create: resourceGlueClassifierCreate,
		Update: resourceGlueClassifierUpdate,
		Delete: resourceGlueClassifierDelete,
		CustomizeDiff: resourceGlueClassifierCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"xml_classifier": {
				Type: schema.TypeSet,
				Elem: propertyClassifierXMLClassifier(),
				Optional: true,
				MaxItems: 1,
			},
			"json_classifier": {
				Type: schema.TypeSet,
				Elem: propertyClassifierJsonClassifier(),
				Optional: true,
				MaxItems: 1,
			},
			"csv_classifier": {
				Type: schema.TypeSet,
				Elem: propertyClassifierCsvClassifier(),
				Optional: true,
				MaxItems: 1,
			},
			"grok_classifier": {
				Type: schema.TypeSet,
				Elem: propertyClassifierGrokClassifier(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceGlueClassifierRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueClassifierType, ResourceGlueClassifier(), data, meta)
}

func resourceGlueClassifierCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueClassifierType, ResourceGlueClassifier(), data, meta)
}

func resourceGlueClassifierUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueClassifierType, ResourceGlueClassifier(), data, meta)
}

func resourceGlueClassifierDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueClassifierType, data, meta)
}

func resourceGlueClassifierCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueClassifierType, data, meta)
}
