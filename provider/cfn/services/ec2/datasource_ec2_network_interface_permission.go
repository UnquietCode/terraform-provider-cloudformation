// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfacepermission.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2NetworkInterfacePermissionType string = "AWS::EC2::NetworkInterfacePermission"

func DatasourceEC2NetworkInterfacePermission() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEC2NetworkInterfacePermissionRead,
		
		Schema: map[string]*schema.Schema{
			"aws_account_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"permission": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceEC2NetworkInterfacePermissionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2NetworkInterfacePermissionType, DatasourceEC2NetworkInterfacePermission(), data, meta)
}
