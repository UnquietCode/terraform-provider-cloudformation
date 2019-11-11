// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html

package directoryservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDirectoryServiceMicrosoftAD() *schema.Resource {
	return &schema.Resource{
		Create: resourceDirectoryServiceMicrosoftADCreate,
		Read:   resourceDirectoryServiceMicrosoftADRead,
		Update: resourceDirectoryServiceMicrosoftADUpdate,
		Delete: resourceDirectoryServiceMicrosoftADDelete,

		Schema: map[string]*schema.Schema{
			"alias": {
				Type: schema.TypeString,
				Computed: true,
			},
			"dns_ip_addresses": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"create_alias": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"edition": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"enable_sso": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"short_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vpc_settings": {
				Type: schema.TypeList,
				Elem: propertyMicrosoftADVpcSettings(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDirectoryServiceMicrosoftADCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DirectoryService::MicrosoftAD", ResourceDirectoryServiceMicrosoftAD(), data, meta)
}

func resourceDirectoryServiceMicrosoftADRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DirectoryService::MicrosoftAD", ResourceDirectoryServiceMicrosoftAD(), data, meta)
}

func resourceDirectoryServiceMicrosoftADUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DirectoryService::MicrosoftAD", ResourceDirectoryServiceMicrosoftAD(), data, meta)
}

func resourceDirectoryServiceMicrosoftADDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DirectoryService::MicrosoftAD", data, meta)
}