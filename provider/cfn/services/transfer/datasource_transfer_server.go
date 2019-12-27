// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceTransferServer() *schema.Resource {
	return &schema.Resource{
		Read: datasourceTransferServerRead,
		
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
			"tags": misc.PropertyTags(),
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

func datasourceTransferServerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(transferServerType, DatasourceTransferServer(), data, meta)
}
