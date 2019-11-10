// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ecr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceRepositoryCreate,
		Read:   resourceRepositoryRead,
		Update: resourceRepositoryUpdate,
		Delete: resourceRepositoryDelete,

		Schema: map[string]*schema.Schema{
			"lifecycle_policy": {
				Type: schema.TypeList,
				Elem: propertyRepositoryLifecyclePolicy(),
				Required: false,
				MaxItems: 1,
			},
			"repository_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"repository_policy_text": {
				Type: schema.TypeMap,
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

func resourceRepositoryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ECR::Repository", data, meta)
}

func resourceRepositoryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ECR::Repository", data, meta)
}

func resourceRepositoryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ECR::Repository", data, meta)
}

func resourceRepositoryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ECR::Repository", data, meta)
}