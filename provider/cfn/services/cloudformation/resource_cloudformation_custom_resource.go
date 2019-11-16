// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html

package cloudformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudFormationCustomResourceType string = "AWS::CloudFormation::CustomResource"

func ResourceCloudFormationCustomResource() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCloudFormationCustomResourceExists,
		Read: resourceCloudFormationCustomResourceRead,
		Create: resourceCloudFormationCustomResourceCreate,
		Update: resourceCloudFormationCustomResourceUpdate,
		Delete: resourceCloudFormationCustomResourceDelete,
		CustomizeDiff: resourceCloudFormationCustomResourceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"service_token": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudFormationCustomResourceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCloudFormationCustomResourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudFormationCustomResourceType, ResourceCloudFormationCustomResource(), data, meta)
}

func resourceCloudFormationCustomResourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloudFormationCustomResourceType, ResourceCloudFormationCustomResource(), data, meta)
}

func resourceCloudFormationCustomResourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloudFormationCustomResourceType, ResourceCloudFormationCustomResource(), data, meta)
}

func resourceCloudFormationCustomResourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloudFormationCustomResourceType, data, meta)
}

func resourceCloudFormationCustomResourceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloudFormationCustomResourceType, data, meta)
}
