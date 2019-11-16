// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html

package opsworks

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var appSourceProperties map[string]string = map[string]string{
	"password": "Password",
	"revision": "Revision",
	"ssh_key": "SshKey",
	"type": "Type",
	"url": "Url",
	"username": "Username",
}

func propertyAppSource(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"password": {
				Type: schema.TypeString,
				Optional: true,
			},
			"revision": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ssh_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"username": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
