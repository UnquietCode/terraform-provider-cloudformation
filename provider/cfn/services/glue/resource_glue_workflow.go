// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html

package glue

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueWorkflowType string = "AWS::Glue::Workflow"

func ResourceGlueWorkflow() *schema.Resource {
	return &schema.Resource{
		Read: resourceGlueWorkflowRead,
		Create: resourceGlueWorkflowCreate,
		Update: resourceGlueWorkflowUpdate,
		Delete: resourceGlueWorkflowDelete,
		CustomizeDiff: resourceGlueWorkflowCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_run_properties": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceGlueWorkflowRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueWorkflowType, ResourceGlueWorkflow(), data, meta)
}

func resourceGlueWorkflowCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueWorkflowType, ResourceGlueWorkflow(), data, meta)
}

func resourceGlueWorkflowUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueWorkflowType, ResourceGlueWorkflow(), data, meta)
}

func resourceGlueWorkflowDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueWorkflowType, data, meta)
}

func resourceGlueWorkflowCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueWorkflowType, data, meta)
}
