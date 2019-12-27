// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html

package appstream

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appStreamImageBuilderType string = "AWS::AppStream::ImageBuilder"

func DatasourceAppStreamImageBuilder() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAppStreamImageBuilderRead,
		
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
				Type: schema.TypeList,
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
				Type: schema.TypeList,
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
			"access_endpoints": {
				Type: schema.TypeList,
				Elem: propertyImageBuilderAccessEndpoint(),
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

func datasourceAppStreamImageBuilderRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appStreamImageBuilderType, DatasourceAppStreamImageBuilder(), data, meta)
}
