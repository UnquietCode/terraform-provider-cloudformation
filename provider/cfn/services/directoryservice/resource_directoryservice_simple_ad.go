// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html

package directoryservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const directoryServiceSimpleADType string = "AWS::DirectoryService::SimpleAD"

func ResourceDirectoryServiceSimpleAD() *schema.Resource {
	return &schema.Resource{
		Read: resourceDirectoryServiceSimpleADRead,
		Create: resourceDirectoryServiceSimpleADCreate,
		Update: resourceDirectoryServiceSimpleADUpdate,
		Delete: resourceDirectoryServiceSimpleADDelete,
		CustomizeDiff: resourceDirectoryServiceSimpleADCustomizeDiff,
		
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
				Type: schema.TypeSet,
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

func resourceDirectoryServiceSimpleADRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(directoryServiceSimpleADType, ResourceDirectoryServiceSimpleAD(), data, meta)
}

func resourceDirectoryServiceSimpleADCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(directoryServiceSimpleADType, ResourceDirectoryServiceSimpleAD(), data, meta)
}

func resourceDirectoryServiceSimpleADUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(directoryServiceSimpleADType, ResourceDirectoryServiceSimpleAD(), data, meta)
}

func resourceDirectoryServiceSimpleADDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(directoryServiceSimpleADType, data, meta)
}

func resourceDirectoryServiceSimpleADCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(directoryServiceSimpleADType, data, meta)
}
