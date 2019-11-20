// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueMLTransformType string = "AWS::Glue::MLTransform"

func ResourceGlueMLTransform() *schema.Resource {
	return &schema.Resource{
		Read: resourceGlueMLTransformRead,
		Create: resourceGlueMLTransformCreate,
		Update: resourceGlueMLTransformUpdate,
		Delete: resourceGlueMLTransformDelete,
		CustomizeDiff: resourceGlueMLTransformCustomizeDiff,
		
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
			"transform_parameters": {
				Type: schema.TypeSet,
				Elem: propertyMLTransformTransformParameters(),
				Required: true,
				MaxItems: 1,
			},
			"input_record_tables": {
				Type: schema.TypeSet,
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
			},
		},
	}
}

func resourceGlueMLTransformRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueMLTransformType, ResourceGlueMLTransform(), data, meta)
}

func resourceGlueMLTransformCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueMLTransformType, ResourceGlueMLTransform(), data, meta)
}

func resourceGlueMLTransformUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueMLTransformType, ResourceGlueMLTransform(), data, meta)
}

func resourceGlueMLTransformDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueMLTransformType, data, meta)
}

func resourceGlueMLTransformCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueMLTransformType, data, meta)
}
