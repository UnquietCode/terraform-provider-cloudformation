// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package batch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyContainerProperties() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"user": {
				Type: schema.TypeString,
				Required: false,
			},
			"memory": {
				Type: schema.TypeInt,
				Required: true,
			},
			"privileged": {
				Type: schema.TypeBool,
				Required: false,
			},
			"linux_parameters": {
				Type: schema.TypeList,
				Elem: propertyLinuxParameters(),
				Required: false,
				MaxItems: 1,
			},
			"job_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"readonly_root_filesystem": {
				Type: schema.TypeBool,
				Required: false,
			},
			"vcpus": {
				Type: schema.TypeInt,
				Required: true,
			},
			"image": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_requirements": {
				Type: schema.TypeList,
				Elem: propertyResourceRequirement(),
				Required: false,
			},
			"mount_points": {
				Type: schema.TypeList,
				Elem: propertyMountPoints(),
				Required: false,
			},
			"volumes": {
				Type: schema.TypeList,
				Elem: propertyVolumes(),
				Required: false,
			},
			"command": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"environment": {
				Type: schema.TypeList,
				Elem: propertyEnvironment(),
				Required: false,
			},
			"ulimits": {
				Type: schema.TypeList,
				Elem: propertyUlimit(),
				Required: false,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}