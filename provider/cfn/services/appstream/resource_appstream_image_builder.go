// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceAppStreamImageBuilder() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppStreamImageBuilderCreate,
		Read:   resourceAppStreamImageBuilderRead,
		Update: resourceAppStreamImageBuilderUpdate,
		Delete: resourceAppStreamImageBuilderDelete,

		Schema: map[string]*schema.Schema{
			"image_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"vpc_config": {
				Type: schema.TypeList,
				Elem: propertyImageBuilderVpcConfig(),
				Required: false,
				MaxItems: 1,
			},
			"enable_default_internet_access": {
				Type: schema.TypeBool,
				Required: false,
			},
			"display_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"domain_join_info": {
				Type: schema.TypeList,
				Elem: propertyImageBuilderDomainJoinInfo(),
				Required: false,
				MaxItems: 1,
			},
			"appstream_agent_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
			"image_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceAppStreamImageBuilderCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::ImageBuilder", data, meta)
}

func resourceAppStreamImageBuilderRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::ImageBuilder", data, meta)
}

func resourceAppStreamImageBuilderUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::ImageBuilder", data, meta)
}

func resourceAppStreamImageBuilderDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::ImageBuilder", data, meta)
}