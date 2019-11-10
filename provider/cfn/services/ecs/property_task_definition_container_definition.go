// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"cpu": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"depends_on": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionContainerDependency(),
				Required: false,
				ForceNew: true,
			},
			"disable_networking": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"dns_search_domains": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"dns_servers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"docker_labels": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"docker_security_options": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"entry_point": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"environment": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionKeyValuePair(),
				Required: false,
				ForceNew: true,
			},
			"essential": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"extra_hosts": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionHostEntry(),
				Required: false,
				ForceNew: true,
			},
			"health_check": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionHealthCheck(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"hostname": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"image": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"interactive": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"links": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"linux_parameters": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionLinuxParameters(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"log_configuration": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionLogConfiguration(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"memory": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"memory_reservation": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"mount_points": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionMountPoint(),
				Required: false,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"port_mappings": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionPortMapping(),
				Required: false,
				ForceNew: true,
			},
			"privileged": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"pseudo_terminal": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"readonly_root_filesystem": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"repository_credentials": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionRepositoryCredentials(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"resource_requirements": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionResourceRequirement(),
				Required: false,
				ForceNew: true,
			},
			"secrets": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionSecret(),
				Required: false,
				ForceNew: true,
			},
			"start_timeout": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"stop_timeout": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"system_controls": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionSystemControl(),
				Required: false,
				ForceNew: true,
			},
			"ulimits": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionUlimit(),
				Required: false,
				ForceNew: true,
			},
			"user": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"volumes_from": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionVolumeFrom(),
				Required: false,
				ForceNew: true,
			},
			"working_directory": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}