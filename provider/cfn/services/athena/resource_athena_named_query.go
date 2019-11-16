// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html

package athena

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const athenaNamedQueryType string = "AWS::Athena::NamedQuery"

var athenaNamedQueryProperties map[string]string = map[string]string{
	"description": "Description",
	"query_string": "QueryString",
	"database": "Database",
	"name": "Name",
}

func ResourceAthenaNamedQuery() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAthenaNamedQueryExists,
		Read: resourceAthenaNamedQueryRead,
		Create: resourceAthenaNamedQueryCreate,
		Update: resourceAthenaNamedQueryUpdate,
		Delete: resourceAthenaNamedQueryDelete,
		CustomizeDiff: resourceAthenaNamedQueryCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"query_string": {
				Type: schema.TypeString,
				Required: true,
			},
			"database": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
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

func resourceAthenaNamedQueryExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAthenaNamedQueryRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(athenaNamedQueryType, ResourceAthenaNamedQuery(), data, meta)
}

func resourceAthenaNamedQueryCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(athenaNamedQueryType, ResourceAthenaNamedQuery(), data, athenaNamedQueryProperties, meta)
}

func resourceAthenaNamedQueryUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(athenaNamedQueryType, ResourceAthenaNamedQuery(), data, athenaNamedQueryProperties, meta)
}

func resourceAthenaNamedQueryDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(athenaNamedQueryType, data, meta)
}

func resourceAthenaNamedQueryCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(athenaNamedQueryType, data, meta)
}
