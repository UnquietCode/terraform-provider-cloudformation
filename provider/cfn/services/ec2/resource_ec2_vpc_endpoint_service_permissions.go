// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservicepermissions.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPCEndpointServicePermissionsType string = "AWS::EC2::VPCEndpointServicePermissions"

func ResourceEC2VPCEndpointServicePermissions() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2VPCEndpointServicePermissionsRead,
		Create: resourceEC2VPCEndpointServicePermissionsCreate,
		Update: resourceEC2VPCEndpointServicePermissionsUpdate,
		Delete: resourceEC2VPCEndpointServicePermissionsDelete,
		CustomizeDiff: resourceEC2VPCEndpointServicePermissionsCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"allowed_principals": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"service_id": {
				Type: schema.TypeString,
				Required: true,
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

func resourceEC2VPCEndpointServicePermissionsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPCEndpointServicePermissionsType, ResourceEC2VPCEndpointServicePermissions(), data, meta)
}

func resourceEC2VPCEndpointServicePermissionsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPCEndpointServicePermissionsType, ResourceEC2VPCEndpointServicePermissions(), data, meta)
}

func resourceEC2VPCEndpointServicePermissionsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPCEndpointServicePermissionsType, ResourceEC2VPCEndpointServicePermissions(), data, meta)
}

func resourceEC2VPCEndpointServicePermissionsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPCEndpointServicePermissionsType, data, meta)
}

func resourceEC2VPCEndpointServicePermissionsCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPCEndpointServicePermissionsType, data, meta)
}
