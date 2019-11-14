// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMResourceDataSync() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSSMResourceDataSyncExists,
		Read: resourceSSMResourceDataSyncRead,
		Create: resourceSSMResourceDataSyncCreate,
		Update: resourceSSMResourceDataSyncUpdate,
		Delete: resourceSSMResourceDataSyncDelete,
		CustomizeDiff: resourceSSMResourceDataSyncCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"kms_key_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"bucket_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"bucket_region": {
				Type: schema.TypeString,
				Required: true,
			},
			"sync_format": {
				Type: schema.TypeString,
				Required: true,
			},
			"sync_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"bucket_prefix": {
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

func resourceSSMResourceDataSyncExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSSMResourceDataSyncRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::ResourceDataSync", ResourceSSMResourceDataSync(), data, meta)
}

func resourceSSMResourceDataSyncCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::ResourceDataSync", ResourceSSMResourceDataSync(), data, meta)
}

func resourceSSMResourceDataSyncUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::ResourceDataSync", ResourceSSMResourceDataSync(), data, meta)
}

func resourceSSMResourceDataSyncDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::ResourceDataSync", data, meta)
}

func resourceSSMResourceDataSyncCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

