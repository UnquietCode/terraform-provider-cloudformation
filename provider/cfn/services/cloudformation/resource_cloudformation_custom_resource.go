// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html

package cloudformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCloudFormationCustomResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFormationCustomResourceCreate,
		Read:   resourceCloudFormationCustomResourceRead,
		Update: resourceCloudFormationCustomResourceUpdate,
		Delete: resourceCloudFormationCustomResourceDelete,

		Schema: map[string]*schema.Schema{
			"service_token": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudFormationCustomResourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFormation::CustomResource", data, meta)
}

func resourceCloudFormationCustomResourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFormation::CustomResource", data, meta)
}

func resourceCloudFormationCustomResourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFormation::CustomResource", data, meta)
}

func resourceCloudFormationCustomResourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFormation::CustomResource", data, meta)
}