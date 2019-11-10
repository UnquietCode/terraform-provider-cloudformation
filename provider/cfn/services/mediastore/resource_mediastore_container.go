// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediastore-container.html

package mediastore

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMediaStoreContainer() *schema.Resource {
	return &schema.Resource{
		Create: resourceMediaStoreContainerCreate,
		Read:   resourceMediaStoreContainerRead,
		Update: resourceMediaStoreContainerUpdate,
		Delete: resourceMediaStoreContainerDelete,

		Schema: map[string]*schema.Schema{
			"policy": {
				Type: schema.TypeString,
				Required: false,
			},
			"container_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cors_policy": {
				Type: schema.TypeList,
				Elem: propertyContainerCorsRule(),
				Required: false,
			},
			"lifecycle_policy": {
				Type: schema.TypeString,
				Required: false,
			},
			"access_logging_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}

func resourceMediaStoreContainerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::MediaStore::Container", data, meta)
}

func resourceMediaStoreContainerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::MediaStore::Container", data, meta)
}

func resourceMediaStoreContainerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::MediaStore::Container", data, meta)
}

func resourceMediaStoreContainerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::MediaStore::Container", data, meta)
}