// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html

package sagemaker

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sageMakerNotebookInstanceLifecycleConfigType string = "AWS::SageMaker::NotebookInstanceLifecycleConfig"

func DatasourceSageMakerNotebookInstanceLifecycleConfig() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSageMakerNotebookInstanceLifecycleConfigRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceSageMakerNotebookInstanceLifecycleConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerNotebookInstanceLifecycleConfigType, DatasourceSageMakerNotebookInstanceLifecycleConfig(), data, meta)
}
