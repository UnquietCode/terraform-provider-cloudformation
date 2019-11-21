// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-resource.html

package lakeformation

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lakeFormationResourceType string = "AWS::LakeFormation::Resource"

func ResourceLakeFormationResource() *schema.Resource {
	return &schema.Resource{
		Read: resourceLakeFormationResourceRead,
		Create: resourceLakeFormationResourceCreate,
		Update: resourceLakeFormationResourceUpdate,
		Delete: resourceLakeFormationResourceDelete,
		CustomizeDiff: resourceLakeFormationResourceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"resource_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"use_service_linked_role": {
				Type: schema.TypeBool,
				Required: true,
			},
			"role_arn": {
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

func resourceLakeFormationResourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lakeFormationResourceType, ResourceLakeFormationResource(), data, meta)
}

func resourceLakeFormationResourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(lakeFormationResourceType, ResourceLakeFormationResource(), data, meta)
}

func resourceLakeFormationResourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(lakeFormationResourceType, ResourceLakeFormationResource(), data, meta)
}

func resourceLakeFormationResourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(lakeFormationResourceType, data, meta)
}

func resourceLakeFormationResourceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(lakeFormationResourceType, data, meta)
}
