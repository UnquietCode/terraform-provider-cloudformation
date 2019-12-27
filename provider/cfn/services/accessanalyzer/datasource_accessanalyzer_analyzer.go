// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html

package accessanalyzer

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const accessAnalyzerAnalyzerType string = "AWS::AccessAnalyzer::Analyzer"

func DatasourceAccessAnalyzerAnalyzer() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAccessAnalyzerAnalyzerRead,
		
		Schema: map[string]*schema.Schema{
			"analyzer_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"archive_rules": {
				Type: schema.TypeList,
				Elem: propertyAnalyzerArchiveRule(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"type": {
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

func datasourceAccessAnalyzerAnalyzerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(accessAnalyzerAnalyzerType, DatasourceAccessAnalyzerAnalyzer(), data, meta)
}
