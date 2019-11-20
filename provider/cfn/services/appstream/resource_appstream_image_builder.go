// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appStreamImageBuilderType string = "AWS::AppStream::ImageBuilder"

func ResourceAppStreamImageBuilder() *schema.Resource {
	return &schema.Resource{
		Read: resourceAppStreamImageBuilderRead,
		Create: resourceAppStreamImageBuilderCreate,
		Update: resourceAppStreamImageBuilderUpdate,
		Delete: resourceAppStreamImageBuilderDelete,
		CustomizeDiff: resourceAppStreamImageBuilderCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"image_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc_config": {
				Type: schema.TypeSet,
				Elem: propertyImageBuilderVpcConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"enable_default_internet_access": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"domain_join_info": {
				Type: schema.TypeSet,
				Elem: propertyImageBuilderDomainJoinInfo(),
				Optional: true,
				MaxItems: 1,
			},
			"appstream_agent_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image_arn": {
				Type: schema.TypeString,
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

func resourceAppStreamImageBuilderRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appStreamImageBuilderType, ResourceAppStreamImageBuilder(), data, meta)
}

func resourceAppStreamImageBuilderCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appStreamImageBuilderType, ResourceAppStreamImageBuilder(), data, meta)
}

func resourceAppStreamImageBuilderUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appStreamImageBuilderType, ResourceAppStreamImageBuilder(), data, meta)
}

func resourceAppStreamImageBuilderDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appStreamImageBuilderType, data, meta)
}

func resourceAppStreamImageBuilderCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appStreamImageBuilderType, data, meta)
}
