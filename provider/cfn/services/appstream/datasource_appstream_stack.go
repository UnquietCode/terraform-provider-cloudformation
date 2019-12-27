// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html

package appstream

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appStreamStackType string = "AWS::AppStream::Stack"

func DatasourceAppStreamStack() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAppStreamStackRead,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"storage_connectors": {
				Type: schema.TypeList,
				Elem: propertyStackStorageConnector(),
				Optional: true,
			},
			"delete_storage_connectors": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"embed_host_domains": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"user_settings": {
				Type: schema.TypeList,
				Elem: propertyStackUserSetting(),
				Optional: true,
			},
			"attributes_to_delete": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"redirect_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"feedback_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"application_settings": {
				Type: schema.TypeList,
				Elem: propertyStackApplicationSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"display_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"access_endpoints": {
				Type: schema.TypeList,
				Elem: propertyStackAccessEndpoint(),
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

func datasourceAppStreamStackRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appStreamStackType, DatasourceAppStreamStack(), data, meta)
}
