// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediastore-container.html

package mediastore

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const mediaStoreContainerType string = "AWS::MediaStore::Container"

func ResourceMediaStoreContainer() *schema.Resource {
	return &schema.Resource{
		Read: resourceMediaStoreContainerRead,
		Create: resourceMediaStoreContainerCreate,
		Update: resourceMediaStoreContainerUpdate,
		Delete: resourceMediaStoreContainerDelete,
		CustomizeDiff: resourceMediaStoreContainerCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"policy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"container_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"cors_policy": {
				Type: schema.TypeList,
				Elem: propertyContainerCorsRule(),
				Optional: true,
			},
			"lifecycle_policy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"access_logging_enabled": {
				Type: schema.TypeBool,
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

func resourceMediaStoreContainerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(mediaStoreContainerType, ResourceMediaStoreContainer(), data, meta)
}

func resourceMediaStoreContainerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(mediaStoreContainerType, ResourceMediaStoreContainer(), data, meta)
}

func resourceMediaStoreContainerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(mediaStoreContainerType, ResourceMediaStoreContainer(), data, meta)
}

func resourceMediaStoreContainerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(mediaStoreContainerType, data, meta)
}

func resourceMediaStoreContainerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(mediaStoreContainerType, data, meta)
}
