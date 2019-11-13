// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html

package directoryservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDirectoryServiceSimpleAD() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDirectoryServiceSimpleADExists,
		Read:   resourceDirectoryServiceSimpleADRead,
		Create: resourceDirectoryServiceSimpleADCreate,
		Update: resourceDirectoryServiceSimpleADUpdate,
		Delete: resourceDirectoryServiceSimpleADDelete,
		
		Schema: map[string]*schema.Schema{
			"create_alias": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enable_sso": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"password": {
				Type: schema.TypeString,
				Required: true,
			},
			"short_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"size": {
				Type: schema.TypeString,
				Required: true,
			},
			"vpc_settings": {
				Type: schema.TypeList,
				Elem: propertySimpleADVpcSettings(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDirectoryServiceSimpleADExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDirectoryServiceSimpleADRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DirectoryService::SimpleAD", ResourceDirectoryServiceSimpleAD(), data, meta)
}

func resourceDirectoryServiceSimpleADCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DirectoryService::SimpleAD", ResourceDirectoryServiceSimpleAD(), data, meta)
}

func resourceDirectoryServiceSimpleADUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DirectoryService::SimpleAD", ResourceDirectoryServiceSimpleAD(), data, meta)
}

func resourceDirectoryServiceSimpleADDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DirectoryService::SimpleAD", data, meta)
}