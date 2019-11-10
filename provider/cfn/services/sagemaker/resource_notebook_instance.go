// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceNotebookInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotebookInstanceCreate,
		Read:   resourceNotebookInstanceRead,
		Update: resourceNotebookInstanceUpdate,
		Delete: resourceNotebookInstanceDelete,

		Schema: map[string]*schema.Schema{
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"volume_size_in_gb": {
				Type: schema.TypeInt,
				Required: false,
			},
			"additional_code_repositories": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"default_code_repository": {
				Type: schema.TypeString,
				Required: false,
			},
			"direct_internet_access": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"accelerator_types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"root_access": {
				Type: schema.TypeString,
				Required: false,
			},
			"notebook_instance_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"lifecycle_config_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceNotebookInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SageMaker::NotebookInstance", data, meta)
}

func resourceNotebookInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SageMaker::NotebookInstance", data, meta)
}

func resourceNotebookInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SageMaker::NotebookInstance", data, meta)
}

func resourceNotebookInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SageMaker::NotebookInstance", data, meta)
}