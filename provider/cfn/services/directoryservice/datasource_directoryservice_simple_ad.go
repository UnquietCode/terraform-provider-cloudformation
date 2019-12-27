// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html

package directoryservice

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const directoryServiceSimpleADType string = "AWS::DirectoryService::SimpleAD"

func DatasourceDirectoryServiceSimpleAD() *schema.Resource {
	return &schema.Resource{
		Read: datasourceDirectoryServiceSimpleADRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceDirectoryServiceSimpleADRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(directoryServiceSimpleADType, DatasourceDirectoryServiceSimpleAD(), data, meta)
}
