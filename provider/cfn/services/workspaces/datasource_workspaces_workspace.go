// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html

package workspaces

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const workSpacesWorkspaceType string = "AWS::WorkSpaces::Workspace"

func DatasourceWorkSpacesWorkspace() *schema.Resource {
	return &schema.Resource{
		Read: datasourceWorkSpacesWorkspaceRead,
		
		Schema: map[string]*schema.Schema{
			"bundle_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"directory_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"root_volume_encryption_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"user_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"user_volume_encryption_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"volume_encryption_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"workspace_properties": {
				Type: schema.TypeList,
				Elem: propertyWorkspaceWorkspaceProperties(),
				Optional: true,
				MaxItems: 1,
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

func datasourceWorkSpacesWorkspaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(workSpacesWorkspaceType, DatasourceWorkSpacesWorkspace(), data, meta)
}
