// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbclusterparametergroup.html

package docdb

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const docDBDBClusterParameterGroupType string = "AWS::DocDB::DBClusterParameterGroup"

func ResourceDocDBDBClusterParameterGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceDocDBDBClusterParameterGroupRead,
		Create: resourceDocDBDBClusterParameterGroupCreate,
		Update: resourceDocDBDBClusterParameterGroupUpdate,
		Delete: resourceDocDBDBClusterParameterGroupDelete,
		CustomizeDiff: resourceDocDBDBClusterParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"family": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"name": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceDocDBDBClusterParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(docDBDBClusterParameterGroupType, ResourceDocDBDBClusterParameterGroup(), data, meta)
}

func resourceDocDBDBClusterParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(docDBDBClusterParameterGroupType, ResourceDocDBDBClusterParameterGroup(), data, meta)
}

func resourceDocDBDBClusterParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(docDBDBClusterParameterGroupType, ResourceDocDBDBClusterParameterGroup(), data, meta)
}

func resourceDocDBDBClusterParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(docDBDBClusterParameterGroupType, data, meta)
}

func resourceDocDBDBClusterParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(docDBDBClusterParameterGroupType, data, meta)
}
