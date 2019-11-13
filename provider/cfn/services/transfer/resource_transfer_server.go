// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html

package transfer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTransferServer() *schema.Resource {
	return &schema.Resource{
		Exists: resourceTransferServerExists,
		Read:   resourceTransferServerRead,
		Create: resourceTransferServerCreate,
		Update: resourceTransferServerUpdate,
		Delete: resourceTransferServerDelete,
		CustomizeDiff: resourceTransferServerCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"logging_role": {
				Type: schema.TypeString,
				Optional: true,
			},
			"identity_provider_details": {
				Type: schema.TypeList,
				Elem: propertyServerIdentityProviderDetails(),
				Optional: true,
				MaxItems: 1,
			},
			"endpoint_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"endpoint_details": {
				Type: schema.TypeList,
				Elem: propertyServerEndpointDetails(),
				Optional: true,
				MaxItems: 1,
			},
			"identity_provider_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceTransferServerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceTransferServerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Transfer::Server", ResourceTransferServer(), data, meta)
}

func resourceTransferServerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Transfer::Server", ResourceTransferServer(), data, meta)
}

func resourceTransferServerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Transfer::Server", ResourceTransferServer(), data, meta)
}

func resourceTransferServerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Transfer::Server", data, meta)
}

func resourceTransferServerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}