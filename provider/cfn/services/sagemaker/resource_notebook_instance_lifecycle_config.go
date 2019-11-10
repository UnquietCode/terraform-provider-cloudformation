// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceNotebookInstanceLifecycleConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotebookInstanceLifecycleConfigCreate,
		Read:   resourceNotebookInstanceLifecycleConfigRead,
		Update: resourceNotebookInstanceLifecycleConfigUpdate,
		Delete: resourceNotebookInstanceLifecycleConfigDelete,

		Schema: map[string]*schema.Schema{
			"on_start": {
				Type: schema.TypeList,
				Elem: propertyNotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook(),
				Required: false,
			},
			"notebook_instance_lifecycle_config_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"on_create": {
				Type: schema.TypeList,
				Elem: propertyNotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook(),
				Required: false,
			},
		},
	}
}

func resourceNotebookInstanceLifecycleConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::NotebookInstanceLifecycleConfig", data, meta)
}

func resourceNotebookInstanceLifecycleConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::NotebookInstanceLifecycleConfig", data, meta)
}

func resourceNotebookInstanceLifecycleConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::NotebookInstanceLifecycleConfig", data, meta)
}

func resourceNotebookInstanceLifecycleConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::NotebookInstanceLifecycleConfig", data, meta)
}