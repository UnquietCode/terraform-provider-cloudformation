// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTriggerAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notification_property": {
				Type: schema.TypeList,
				Elem: propertyTriggerNotificationProperty(),
				Required: false,
				MaxItems: 1,
			},
			"crawler_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"timeout": {
				Type: schema.TypeInt,
				Required: false,
			},
			"job_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"arguments": {
				Type: schema.TypeMap,
				Required: false,
			},
			"security_configuration": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}