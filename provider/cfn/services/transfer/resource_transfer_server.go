// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html

package transfer

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const transferServerType string = "AWS::Transfer::Server"

func ResourceTransferServer() *schema.Resource {
	return &schema.Resource{
		Read: resourceTransferServerRead,
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
				Type: schema.TypeSet,
				Elem: propertyServerIdentityProviderDetails(),
				Optional: true,
				MaxItems: 1,
			},
			"endpoint_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"endpoint_details": {
				Type: schema.TypeSet,
				Elem: propertyServerEndpointDetails(),
				Optional: true,
				MaxItems: 1,
			},
			"identity_provider_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceTransferServerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(transferServerType, ResourceTransferServer(), data, meta)
}

func resourceTransferServerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(transferServerType, ResourceTransferServer(), data, meta)
}

func resourceTransferServerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(transferServerType, ResourceTransferServer(), data, meta)
}

func resourceTransferServerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(transferServerType, data, meta)
}

func resourceTransferServerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(transferServerType, data, meta)
}
