// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

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
			"on_start": {
				Type: schema.TypeList,
				Elem: propertyNotebookInstanceLifecycleHook(),
				Required: false,
			},
			"notebook_instance_lifecycle_config_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"on_create": {
				Type: schema.TypeList,
				Elem: propertyNotebookInstanceLifecycleHook(),
				Required: false,
			},
		},
	}
}

func resourceSageMakerNotebookInstanceLifecycleConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::NotebookInstanceLifecycleConfig", data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::NotebookInstanceLifecycleConfig", data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::NotebookInstanceLifecycleConfig", data, meta)
}

func resourceSageMakerNotebookInstanceLifecycleConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::NotebookInstanceLifecycleConfig", data, meta)
}