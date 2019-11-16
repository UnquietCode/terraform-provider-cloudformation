// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html

package ecs

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var taskDefinitionContainerDefinitionProperties map[string]string = map[string]string{
	"command": "Command",
	"cpu": "Cpu",
	"depends_on": "DependsOn",
	"disable_networking": "DisableNetworking",
	"dns_search_domains": "DnsSearchDomains",
	"dns_servers": "DnsServers",
	"docker_labels": "DockerLabels",
	"docker_security_options": "DockerSecurityOptions",
	"entry_point": "EntryPoint",
	"environment": "Environment",
	"essential": "Essential",
	"extra_hosts": "ExtraHosts",
	"health_check": "HealthCheck",
	"hostname": "Hostname",
	"image": "Image",
	"interactive": "Interactive",
	"links": "Links",
	"linux_parameters": "LinuxParameters",
	"log_configuration": "LogConfiguration",
	"memory": "Memory",
	"memory_reservation": "MemoryReservation",
	"mount_points": "MountPoints",
	"name": "Name",
	"port_mappings": "PortMappings",
	"privileged": "Privileged",
	"pseudo_terminal": "PseudoTerminal",
	"readonly_root_filesystem": "ReadonlyRootFilesystem",
	"repository_credentials": "RepositoryCredentials",
	"resource_requirements": "ResourceRequirements",
	"secrets": "Secrets",
	"start_timeout": "StartTimeout",
	"stop_timeout": "StopTimeout",
	"system_controls": "SystemControls",
	"ulimits": "Ulimits",
	"user": "User",
	"volumes_from": "VolumesFrom",
	"working_directory": "WorkingDirectory",
}

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
				Type: schema.TypeList,
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
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionLinuxParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"log_configuration": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
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
