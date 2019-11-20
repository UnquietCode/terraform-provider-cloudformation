// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html

package ecs

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTaskDefinitionContainerDefinition(extras...string) *schema.Resource {
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
			"command": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"cpu": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"depends_on": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionContainerDependency(),
				Optional: true,
			},
			"disable_networking": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"dns_search_domains": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"dns_servers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"docker_labels": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"docker_security_options": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"entry_point": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"environment": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionKeyValuePair(),
				Optional: true,
			},
			"essential": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"extra_hosts": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionHostEntry(),
				Optional: true,
			},
			"health_check": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionHealthCheck(),
				Optional: true,
				MaxItems: 1,
			},
			"hostname": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image": {
				Type: schema.TypeString,
				Optional: true,
			},
			"interactive": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"links": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"linux_parameters": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionLinuxParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"log_configuration": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionLogConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"memory": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"memory_reservation": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"mount_points": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionMountPoint(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"port_mappings": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionPortMapping(),
				Optional: true,
			},
			"privileged": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"pseudo_terminal": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"readonly_root_filesystem": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"repository_credentials": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionRepositoryCredentials(),
				Optional: true,
				MaxItems: 1,
			},
			"resource_requirements": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionResourceRequirement(),
				Optional: true,
			},
			"secrets": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionSecret(),
				Optional: true,
			},
			"start_timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"stop_timeout": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"system_controls": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionSystemControl(),
				Optional: true,
			},
			"ulimits": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionUlimit(),
				Optional: true,
			},
			"user": {
				Type: schema.TypeString,
				Optional: true,
			},
			"volumes_from": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionVolumeFrom(),
				Optional: true,
			},
			"working_directory": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
