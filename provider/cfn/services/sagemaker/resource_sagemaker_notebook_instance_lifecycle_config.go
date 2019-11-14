// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSageMakerNotebookInstanceLifecycleConfig() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSageMakerNotebookInstanceLifecycleConfigExists,
		Read: resourceSageMakerNotebookInstanceLifecycleConfigRead,
		Create: resourceSageMakerNotebookInstanceLifecycleConfigCreate,
		Update: resourceSageMakerNotebookInstanceLifecycleConfigUpdate,
		Delete: resourceSageMakerNotebookInstanceLifecycleConfigDelete,
		CustomizeDiff: resourceSageMakerNotebookInstanceLifecycleConfigCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"on_start": {
				Type: schema.TypeList,
				Elem: propertyNotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook(),
				Optional: true,
			},
			"notebook_instance_lifecycle_config_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"on_create": {
				Type: schema.TypeList,
				Elem: propertyNotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook(),
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

func resourceSageMakerNotebookInstanceLifecycleConfigExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::NotebookInstanceLifecycleConfig", ResourceSageMakerNotebookInstanceLifecycleConfig(), data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::NotebookInstanceLifecycleConfig", ResourceSageMakerNotebookInstanceLifecycleConfig(), data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::NotebookInstanceLifecycleConfig", ResourceSageMakerNotebookInstanceLifecycleConfig(), data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::NotebookInstanceLifecycleConfig", data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

