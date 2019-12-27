// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html

package kms

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kMSKeyType string = "AWS::KMS::Key"

func DatasourceKMSKey() *schema.Resource {
	return &schema.Resource{
		Read: datasourceKMSKeyRead,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enable_key_rotation": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"key_policy": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"key_usage": {
				Type: schema.TypeString,
				Optional: true,
			},
			"pending_window_in_days": {
				Type: schema.TypeInt,
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

func datasourceKMSKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kMSKeyType, DatasourceKMSKey(), data, meta)
}
