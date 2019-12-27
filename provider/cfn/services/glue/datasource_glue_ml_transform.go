// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html

package glue

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueMLTransformType string = "AWS::Glue::MLTransform"

func DatasourceGlueMLTransform() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGlueMLTransformRead,
		
		Schema: map[string]*schema.Schema{
			"role": {
				Type: schema.TypeString,
				Required: true,
			},
			"max_retries": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"worker_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"glue_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"transform_parameters": {
				Type: schema.TypeList,
				Elem: propertyMLTransformTransformParameters(),
				Required: true,
				MaxItems: 1,
			},
			"input_record_tables": {
				Type: schema.TypeList,
				Elem: propertyMLTransformInputRecordTables(),
				Required: true,
				MaxItems: 1,
			},
			"number_of_workers": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"max_capacity": {
				Type: schema.TypeFloat,
				Optional: true,
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

func datasourceGlueMLTransformRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueMLTransformType, DatasourceGlueMLTransform(), data, meta)
}
