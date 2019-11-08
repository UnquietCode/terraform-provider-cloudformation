// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMResourceDataSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSMResourceDataSyncCreate,
		Read:   resourceSSMResourceDataSyncRead,
		Update: resourceSSMResourceDataSyncUpdate,
		Delete: resourceSSMResourceDataSyncDelete,

		Schema: map[string]*schema.Schema{
			"kms_key_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"bucket_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bucket_region": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sync_format": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sync_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bucket_prefix": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceSSMResourceDataSyncCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::ResourceDataSync", data, meta)
}

func resourceSSMResourceDataSyncRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::ResourceDataSync", data, meta)
}

func resourceSSMResourceDataSyncUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::ResourceDataSync", data, meta)
}

func resourceSSMResourceDataSyncDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::ResourceDataSync", data, meta)
}