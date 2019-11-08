// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package transfer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTransferServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceTransferServerCreate,
		Read:   resourceTransferServerRead,
		Update: resourceTransferServerUpdate,
		Delete: resourceTransferServerDelete,

		Schema: map[string]*schema.Schema{
			"logging_role": {
				Type: schema.TypeString,
				Required: false,
			},
			"identity_provider_details": {
				Type: schema.TypeList,
				Elem: propertyIdentityProviderDetails(),
				Required: false,
				MaxItems: 1,
			},
			"endpoint_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"endpoint_details": {
				Type: schema.TypeList,
				Elem: propertyEndpointDetails(),
				Required: false,
				MaxItems: 1,
			},
			"identity_provider_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceTransferServerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Transfer::Server", data, meta)
}

func resourceTransferServerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Transfer::Server", data, meta)
}

func resourceTransferServerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Transfer::Server", data, meta)
}

func resourceTransferServerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Transfer::Server", data, meta)
}