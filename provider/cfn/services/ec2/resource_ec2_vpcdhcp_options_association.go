// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VPCDHCPOptionsAssociationType string = "AWS::EC2::VPCDHCPOptionsAssociation"

func ResourceEC2VPCDHCPOptionsAssociation() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2VPCDHCPOptionsAssociationRead,
		Create: resourceEC2VPCDHCPOptionsAssociationCreate,
		Update: resourceEC2VPCDHCPOptionsAssociationUpdate,
		Delete: resourceEC2VPCDHCPOptionsAssociationDelete,
		CustomizeDiff: resourceEC2VPCDHCPOptionsAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"dhcp_options_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"vpc_id": {
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

func resourceEC2VPCDHCPOptionsAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VPCDHCPOptionsAssociationType, ResourceEC2VPCDHCPOptionsAssociation(), data, meta)
}

func resourceEC2VPCDHCPOptionsAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VPCDHCPOptionsAssociationType, ResourceEC2VPCDHCPOptionsAssociation(), data, meta)
}

func resourceEC2VPCDHCPOptionsAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VPCDHCPOptionsAssociationType, ResourceEC2VPCDHCPOptionsAssociation(), data, meta)
}

func resourceEC2VPCDHCPOptionsAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VPCDHCPOptionsAssociationType, data, meta)
}

func resourceEC2VPCDHCPOptionsAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VPCDHCPOptionsAssociationType, data, meta)
}
