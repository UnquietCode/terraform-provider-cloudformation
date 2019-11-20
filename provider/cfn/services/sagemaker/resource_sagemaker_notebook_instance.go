// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sageMakerNotebookInstanceType string = "AWS::SageMaker::NotebookInstance"

func ResourceSageMakerNotebookInstance() *schema.Resource {
	return &schema.Resource{
		Read: resourceSageMakerNotebookInstanceRead,
		Create: resourceSageMakerNotebookInstanceCreate,
		Update: resourceSageMakerNotebookInstanceUpdate,
		Delete: resourceSageMakerNotebookInstanceDelete,
		CustomizeDiff: resourceSageMakerNotebookInstanceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"volume_size_in_gb": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"additional_code_repositories": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"default_code_repository": {
				Type: schema.TypeString,
				Optional: true,
			},
			"direct_internet_access": {
				Type: schema.TypeString,
				Optional: true,
			},
			"accelerator_types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"root_access": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notebook_instance_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"lifecycle_config_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSageMakerNotebookInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sageMakerNotebookInstanceType, ResourceSageMakerNotebookInstance(), data, meta)
}

func resourceSageMakerNotebookInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sageMakerNotebookInstanceType, ResourceSageMakerNotebookInstance(), data, meta)
}

func resourceSageMakerNotebookInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sageMakerNotebookInstanceType, ResourceSageMakerNotebookInstance(), data, meta)
}

func resourceSageMakerNotebookInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sageMakerNotebookInstanceType, data, meta)
}

func resourceSageMakerNotebookInstanceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sageMakerNotebookInstanceType, data, meta)
}
