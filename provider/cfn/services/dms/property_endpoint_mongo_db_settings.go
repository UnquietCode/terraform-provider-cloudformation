// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-mongodbsettings.html

package dms

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEndpointMongoDbSettings(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_source": {
				Type: schema.TypeString,
				Required: false,
			},
			"auth_mechanism": {
				Type: schema.TypeString,
				Required: false,
			},
			"username": {
				Type: schema.TypeString,
				Required: false,
			},
			"docs_to_investigate": {
				Type: schema.TypeString,
				Required: false,
			},
			"server_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"extract_doc_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"auth_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"password": {
				Type: schema.TypeString,
				Required: false,
			},
			"nesting_level": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}