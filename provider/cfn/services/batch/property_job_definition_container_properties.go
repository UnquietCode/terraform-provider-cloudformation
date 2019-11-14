// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html

package batch

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyJobDefinitionContainerProperties(extras...string) *schema.Resource {
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
			"user": {
				Type: schema.TypeString,
				Optional: true,
			},
			"memory": {
				Type: schema.TypeInt,
				Required: true,
			},
			"privileged": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"linux_parameters": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionLinuxParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"job_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"readonly_root_filesystem": {
				Type: schema.TypeBool,
				Optional: true,
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
				Elem: propertyJobDefinitionResourceRequirement(),
				Optional: true,
			},
			"mount_points": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionMountPoints(),
				Optional: true,
			},
			"volumes": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionVolumes(),
				Optional: true,
			},
			"command": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"environment": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionEnvironment(),
				Optional: true,
			},
			"ulimits": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionUlimit(),
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

