// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceSageMakerNotebookInstanceLifecycleConfigCreate,
		Read:   resourceSageMakerNotebookInstanceLifecycleConfigRead,
		Update: resourceSageMakerNotebookInstanceLifecycleConfigUpdate,
		Delete: resourceSageMakerNotebookInstanceLifecycleConfigDelete,

		Schema: map[string]*schema.Schema{
			"notebook_instance_lifecycle_config_name": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"on_start": {
				Type: schema.TypeList,
				Elem: propertyNotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook(),
				Optional: true,
			},
			"on_create": {
				Type: schema.TypeList,
				Elem: propertyNotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook(),
				Optional: true,
			},
		},
	}
}

func resourceSageMakerNotebookInstanceLifecycleConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::NotebookInstanceLifecycleConfig", ResourceSageMakerNotebookInstanceLifecycleConfig(), data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::NotebookInstanceLifecycleConfig", ResourceSageMakerNotebookInstanceLifecycleConfig(), data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::NotebookInstanceLifecycleConfig", ResourceSageMakerNotebookInstanceLifecycleConfig(), data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::NotebookInstanceLifecycleConfig", data, meta)
}