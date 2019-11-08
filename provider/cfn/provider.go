// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package cfn

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/amazonmq"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/amplify"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/apigateway"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/apigatewayv2"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/applicationautoscaling"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/appmesh"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/appstream"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/appsync"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/athena"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/autoscaling"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/autoscalingplans"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/backup"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/batch"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/budgets"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/certificatemanager"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/cloud9"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/cloudformation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/cloudfront"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/cloudtrail"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/cloudwatch"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/codebuild"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/codecommit"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/codedeploy"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/codepipeline"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/codestar"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/cognito"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/config"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/datapipeline"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/dax"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/directoryservice"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/dlm"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/dms"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/docdb"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/dynamodb"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/ec2"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/ecr"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/ecs"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/efs"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/eks"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/elasticache"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/elasticbeanstalk"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/elasticloadbalancing"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/elasticloadbalancingv2"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/elasticsearch"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/emr"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/events"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/fsx"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/gamelift"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/glue"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/greengrass"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/guardduty"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/iam"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/inspector"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/iot"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/iot1click"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/iotanalytics"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/iotevents"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/iotthingsgraph"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/kinesis"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/kinesisanalytics"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/kinesisanalyticsv2"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/kinesisfirehose"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/kms"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/lakeformation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/lambda"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/logs"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/managedblockchain"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/medialive"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/mediastore"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/msk"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/neptune"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/opsworks"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/opsworkscm"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/pinpoint"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/pinpointemail"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/qldb"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/ram"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/rds"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/redshift"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/robomaker"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/route53"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/route53resolver"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/s3"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/sagemaker"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/sdb"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/secretsmanager"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/securityhub"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/servicecatalog"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/servicediscovery"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/ses"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/sns"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/sqs"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/ssm"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/stepfunctions"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/transfer"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/waf"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/wafregional"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/workspaces"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
		},
		ResourcesMap: map[string]*schema.Resource{
			"amazon_mq_broker_resource": amazonmq.ResourceAmazonMQBroker(),
			"amazon_mq_configuration_resource": amazonmq.ResourceAmazonMQConfiguration(),
			"amazon_mq_configuration_association_resource": amazonmq.ResourceAmazonMQConfigurationAssociation(),
			"amplify_app_resource": amplify.ResourceAmplifyApp(),
			"amplify_branch_resource": amplify.ResourceAmplifyBranch(),
			"amplify_domain_resource": amplify.ResourceAmplifyDomain(),
			"api_gateway_account_resource": apigateway.ResourceApiGatewayAccount(),
			"api_gateway_api_key_resource": apigateway.ResourceApiGatewayApiKey(),
			"api_gateway_authorizer_resource": apigateway.ResourceApiGatewayAuthorizer(),
			"api_gateway_base_path_mapping_resource": apigateway.ResourceApiGatewayBasePathMapping(),
			"api_gateway_client_certificate_resource": apigateway.ResourceApiGatewayClientCertificate(),
			"api_gateway_deployment_resource": apigateway.ResourceApiGatewayDeployment(),
			"api_gateway_documentation_part_resource": apigateway.ResourceApiGatewayDocumentationPart(),
			"api_gateway_documentation_version_resource": apigateway.ResourceApiGatewayDocumentationVersion(),
			"api_gateway_domain_name_resource": apigateway.ResourceApiGatewayDomainName(),
			"api_gateway_gateway_response_resource": apigateway.ResourceApiGatewayGatewayResponse(),
			"api_gateway_method_resource": apigateway.ResourceApiGatewayMethod(),
			"api_gateway_model_resource": apigateway.ResourceApiGatewayModel(),
			"api_gateway_request_validator_resource": apigateway.ResourceApiGatewayRequestValidator(),
			"api_gateway_resource_resource": apigateway.ResourceApiGatewayResource(),
			"api_gateway_rest_api_resource": apigateway.ResourceApiGatewayRestApi(),
			"api_gateway_stage_resource": apigateway.ResourceApiGatewayStage(),
			"api_gateway_usage_plan_resource": apigateway.ResourceApiGatewayUsagePlan(),
			"api_gateway_usage_plan_key_resource": apigateway.ResourceApiGatewayUsagePlanKey(),
			"api_gateway_vpc_link_resource": apigateway.ResourceApiGatewayVpcLink(),
			"api_gateway_v2_api_resource": apigatewayv2.ResourceApiGatewayV2Api(),
			"api_gateway_v2_api_mapping_resource": apigatewayv2.ResourceApiGatewayV2ApiMapping(),
			"api_gateway_v2_authorizer_resource": apigatewayv2.ResourceApiGatewayV2Authorizer(),
			"api_gateway_v2_deployment_resource": apigatewayv2.ResourceApiGatewayV2Deployment(),
			"api_gateway_v2_domain_name_resource": apigatewayv2.ResourceApiGatewayV2DomainName(),
			"api_gateway_v2_integration_resource": apigatewayv2.ResourceApiGatewayV2Integration(),
			"api_gateway_v2_integration_response_resource": apigatewayv2.ResourceApiGatewayV2IntegrationResponse(),
			"api_gateway_v2_model_resource": apigatewayv2.ResourceApiGatewayV2Model(),
			"api_gateway_v2_route_resource": apigatewayv2.ResourceApiGatewayV2Route(),
			"api_gateway_v2_route_response_resource": apigatewayv2.ResourceApiGatewayV2RouteResponse(),
			"api_gateway_v2_stage_resource": apigatewayv2.ResourceApiGatewayV2Stage(),
			"application_auto_scaling_scalable_target_resource": applicationautoscaling.ResourceApplicationAutoScalingScalableTarget(),
			"application_auto_scaling_scaling_policy_resource": applicationautoscaling.ResourceApplicationAutoScalingScalingPolicy(),
			"app_mesh_mesh_resource": appmesh.ResourceAppMeshMesh(),
			"app_mesh_route_resource": appmesh.ResourceAppMeshRoute(),
			"app_mesh_virtual_node_resource": appmesh.ResourceAppMeshVirtualNode(),
			"app_mesh_virtual_router_resource": appmesh.ResourceAppMeshVirtualRouter(),
			"app_mesh_virtual_service_resource": appmesh.ResourceAppMeshVirtualService(),
			"app_stream_directory_config_resource": appstream.ResourceAppStreamDirectoryConfig(),
			"app_stream_fleet_resource": appstream.ResourceAppStreamFleet(),
			"app_stream_image_builder_resource": appstream.ResourceAppStreamImageBuilder(),
			"app_stream_stack_resource": appstream.ResourceAppStreamStack(),
			"app_stream_stack_fleet_association_resource": appstream.ResourceAppStreamStackFleetAssociation(),
			"app_stream_stack_user_association_resource": appstream.ResourceAppStreamStackUserAssociation(),
			"app_stream_user_resource": appstream.ResourceAppStreamUser(),
			"app_sync_api_key_resource": appsync.ResourceAppSyncApiKey(),
			"app_sync_data_source_resource": appsync.ResourceAppSyncDataSource(),
			"app_sync_function_configuration_resource": appsync.ResourceAppSyncFunctionConfiguration(),
			"app_sync_graph_ql_api_resource": appsync.ResourceAppSyncGraphQLApi(),
			"app_sync_graph_ql_schema_resource": appsync.ResourceAppSyncGraphQLSchema(),
			"app_sync_resolver_resource": appsync.ResourceAppSyncResolver(),
			"athena_named_query_resource": athena.ResourceAthenaNamedQuery(),
			"auto_scaling_auto_scaling_group_resource": autoscaling.ResourceAutoScalingAutoScalingGroup(),
			"auto_scaling_launch_configuration_resource": autoscaling.ResourceAutoScalingLaunchConfiguration(),
			"auto_scaling_lifecycle_hook_resource": autoscaling.ResourceAutoScalingLifecycleHook(),
			"auto_scaling_scaling_policy_resource": autoscaling.ResourceAutoScalingScalingPolicy(),
			"auto_scaling_scheduled_action_resource": autoscaling.ResourceAutoScalingScheduledAction(),
			"auto_scaling_plans_scaling_plan_resource": autoscalingplans.ResourceAutoScalingPlansScalingPlan(),
			"backup_backup_plan_resource": backup.ResourceBackupBackupPlan(),
			"backup_backup_selection_resource": backup.ResourceBackupBackupSelection(),
			"backup_backup_vault_resource": backup.ResourceBackupBackupVault(),
			"batch_compute_environment_resource": batch.ResourceBatchComputeEnvironment(),
			"batch_job_definition_resource": batch.ResourceBatchJobDefinition(),
			"batch_job_queue_resource": batch.ResourceBatchJobQueue(),
			"budgets_budget_resource": budgets.ResourceBudgetsBudget(),
			"certificate_manager_certificate_resource": certificatemanager.ResourceCertificateManagerCertificate(),
			"cloud9_environment_ec2_resource": cloud9.ResourceCloud9EnvironmentEC2(),
			"cloud_formation_custom_resource_resource": cloudformation.ResourceCloudFormationCustomResource(),
			"cloud_formation_macro_resource": cloudformation.ResourceCloudFormationMacro(),
			"cloud_formation_stack_resource": cloudformation.ResourceCloudFormationStack(),
			"cloud_formation_wait_condition_resource": cloudformation.ResourceCloudFormationWaitCondition(),
			"cloud_formation_wait_condition_handle_resource": cloudformation.ResourceCloudFormationWaitConditionHandle(),
			"cloud_front_cloud_front_origin_access_identity_resource": cloudfront.ResourceCloudFrontCloudFrontOriginAccessIdentity(),
			"cloud_front_distribution_resource": cloudfront.ResourceCloudFrontDistribution(),
			"cloud_front_streaming_distribution_resource": cloudfront.ResourceCloudFrontStreamingDistribution(),
			"cloud_trail_trail_resource": cloudtrail.ResourceCloudTrailTrail(),
			"cloud_watch_alarm_resource": cloudwatch.ResourceCloudWatchAlarm(),
			"cloud_watch_anomaly_detector_resource": cloudwatch.ResourceCloudWatchAnomalyDetector(),
			"cloud_watch_dashboard_resource": cloudwatch.ResourceCloudWatchDashboard(),
			"code_build_project_resource": codebuild.ResourceCodeBuildProject(),
			"code_build_source_credential_resource": codebuild.ResourceCodeBuildSourceCredential(),
			"code_commit_repository_resource": codecommit.ResourceCodeCommitRepository(),
			"code_deploy_application_resource": codedeploy.ResourceCodeDeployApplication(),
			"code_deploy_deployment_config_resource": codedeploy.ResourceCodeDeployDeploymentConfig(),
			"code_deploy_deployment_group_resource": codedeploy.ResourceCodeDeployDeploymentGroup(),
			"code_pipeline_custom_action_type_resource": codepipeline.ResourceCodePipelineCustomActionType(),
			"code_pipeline_pipeline_resource": codepipeline.ResourceCodePipelinePipeline(),
			"code_pipeline_webhook_resource": codepipeline.ResourceCodePipelineWebhook(),
			"code_star_git_hub_repository_resource": codestar.ResourceCodeStarGitHubRepository(),
			"cognito_identity_pool_resource": cognito.ResourceCognitoIdentityPool(),
			"cognito_identity_pool_role_attachment_resource": cognito.ResourceCognitoIdentityPoolRoleAttachment(),
			"cognito_user_pool_resource": cognito.ResourceCognitoUserPool(),
			"cognito_user_pool_client_resource": cognito.ResourceCognitoUserPoolClient(),
			"cognito_user_pool_domain_resource": cognito.ResourceCognitoUserPoolDomain(),
			"cognito_user_pool_group_resource": cognito.ResourceCognitoUserPoolGroup(),
			"cognito_user_pool_identity_provider_resource": cognito.ResourceCognitoUserPoolIdentityProvider(),
			"cognito_user_pool_resource_server_resource": cognito.ResourceCognitoUserPoolResourceServer(),
			"cognito_user_pool_risk_configuration_attachment_resource": cognito.ResourceCognitoUserPoolRiskConfigurationAttachment(),
			"cognito_user_pool_ui_customization_attachment_resource": cognito.ResourceCognitoUserPoolUICustomizationAttachment(),
			"cognito_user_pool_user_resource": cognito.ResourceCognitoUserPoolUser(),
			"cognito_user_pool_user_to_group_attachment_resource": cognito.ResourceCognitoUserPoolUserToGroupAttachment(),
			"config_aggregation_authorization_resource": config.ResourceConfigAggregationAuthorization(),
			"config_config_rule_resource": config.ResourceConfigConfigRule(),
			"config_configuration_aggregator_resource": config.ResourceConfigConfigurationAggregator(),
			"config_configuration_recorder_resource": config.ResourceConfigConfigurationRecorder(),
			"config_delivery_channel_resource": config.ResourceConfigDeliveryChannel(),
			"config_organization_config_rule_resource": config.ResourceConfigOrganizationConfigRule(),
			"config_remediation_configuration_resource": config.ResourceConfigRemediationConfiguration(),
			"data_pipeline_pipeline_resource": datapipeline.ResourceDataPipelinePipeline(),
			"dax_cluster_resource": dax.ResourceDAXCluster(),
			"dax_parameter_group_resource": dax.ResourceDAXParameterGroup(),
			"dax_subnet_group_resource": dax.ResourceDAXSubnetGroup(),
			"directory_service_microsoft_ad_resource": directoryservice.ResourceDirectoryServiceMicrosoftAD(),
			"directory_service_simple_ad_resource": directoryservice.ResourceDirectoryServiceSimpleAD(),
			"dlm_lifecycle_policy_resource": dlm.ResourceDLMLifecyclePolicy(),
			"dms_certificate_resource": dms.ResourceDMSCertificate(),
			"dms_endpoint_resource": dms.ResourceDMSEndpoint(),
			"dms_event_subscription_resource": dms.ResourceDMSEventSubscription(),
			"dms_replication_instance_resource": dms.ResourceDMSReplicationInstance(),
			"dms_replication_subnet_group_resource": dms.ResourceDMSReplicationSubnetGroup(),
			"dms_replication_task_resource": dms.ResourceDMSReplicationTask(),
			"doc_dbdb_cluster_resource": docdb.ResourceDocDBDBCluster(),
			"doc_dbdb_cluster_parameter_group_resource": docdb.ResourceDocDBDBClusterParameterGroup(),
			"doc_dbdb_instance_resource": docdb.ResourceDocDBDBInstance(),
			"doc_dbdb_subnet_group_resource": docdb.ResourceDocDBDBSubnetGroup(),
			"dynamo_db_table_resource": dynamodb.ResourceDynamoDBTable(),
			"ec2_capacity_reservation_resource": ec2.ResourceEC2CapacityReservation(),
			"ec2_client_vpn_authorization_rule_resource": ec2.ResourceEC2ClientVpnAuthorizationRule(),
			"ec2_client_vpn_endpoint_resource": ec2.ResourceEC2ClientVpnEndpoint(),
			"ec2_client_vpn_route_resource": ec2.ResourceEC2ClientVpnRoute(),
			"ec2_client_vpn_target_network_association_resource": ec2.ResourceEC2ClientVpnTargetNetworkAssociation(),
			"ec2_customer_gateway_resource": ec2.ResourceEC2CustomerGateway(),
			"ec2_dhcp_options_resource": ec2.ResourceEC2DHCPOptions(),
			"ec2_ec2_fleet_resource": ec2.ResourceEC2EC2Fleet(),
			"ec2_eip_resource": ec2.ResourceEC2EIP(),
			"ec2_eip_association_resource": ec2.ResourceEC2EIPAssociation(),
			"ec2_egress_only_internet_gateway_resource": ec2.ResourceEC2EgressOnlyInternetGateway(),
			"ec2_flow_log_resource": ec2.ResourceEC2FlowLog(),
			"ec2_host_resource": ec2.ResourceEC2Host(),
			"ec2_instance_resource": ec2.ResourceEC2Instance(),
			"ec2_internet_gateway_resource": ec2.ResourceEC2InternetGateway(),
			"ec2_launch_template_resource": ec2.ResourceEC2LaunchTemplate(),
			"ec2_nat_gateway_resource": ec2.ResourceEC2NatGateway(),
			"ec2_network_acl_resource": ec2.ResourceEC2NetworkAcl(),
			"ec2_network_acl_entry_resource": ec2.ResourceEC2NetworkAclEntry(),
			"ec2_network_interface_resource": ec2.ResourceEC2NetworkInterface(),
			"ec2_network_interface_attachment_resource": ec2.ResourceEC2NetworkInterfaceAttachment(),
			"ec2_network_interface_permission_resource": ec2.ResourceEC2NetworkInterfacePermission(),
			"ec2_placement_group_resource": ec2.ResourceEC2PlacementGroup(),
			"ec2_route_resource": ec2.ResourceEC2Route(),
			"ec2_route_table_resource": ec2.ResourceEC2RouteTable(),
			"ec2_security_group_resource": ec2.ResourceEC2SecurityGroup(),
			"ec2_security_group_egress_resource": ec2.ResourceEC2SecurityGroupEgress(),
			"ec2_security_group_ingress_resource": ec2.ResourceEC2SecurityGroupIngress(),
			"ec2_spot_fleet_resource": ec2.ResourceEC2SpotFleet(),
			"ec2_subnet_resource": ec2.ResourceEC2Subnet(),
			"ec2_subnet_cidr_block_resource": ec2.ResourceEC2SubnetCidrBlock(),
			"ec2_subnet_network_acl_association_resource": ec2.ResourceEC2SubnetNetworkAclAssociation(),
			"ec2_subnet_route_table_association_resource": ec2.ResourceEC2SubnetRouteTableAssociation(),
			"ec2_traffic_mirror_filter_resource": ec2.ResourceEC2TrafficMirrorFilter(),
			"ec2_traffic_mirror_filter_rule_resource": ec2.ResourceEC2TrafficMirrorFilterRule(),
			"ec2_traffic_mirror_session_resource": ec2.ResourceEC2TrafficMirrorSession(),
			"ec2_traffic_mirror_target_resource": ec2.ResourceEC2TrafficMirrorTarget(),
			"ec2_transit_gateway_resource": ec2.ResourceEC2TransitGateway(),
			"ec2_transit_gateway_attachment_resource": ec2.ResourceEC2TransitGatewayAttachment(),
			"ec2_transit_gateway_route_resource": ec2.ResourceEC2TransitGatewayRoute(),
			"ec2_transit_gateway_route_table_resource": ec2.ResourceEC2TransitGatewayRouteTable(),
			"ec2_transit_gateway_route_table_association_resource": ec2.ResourceEC2TransitGatewayRouteTableAssociation(),
			"ec2_transit_gateway_route_table_propagation_resource": ec2.ResourceEC2TransitGatewayRouteTablePropagation(),
			"ec2_vpc_resource": ec2.ResourceEC2VPC(),
			"ec2_vpc_cidr_block_resource": ec2.ResourceEC2VPCCidrBlock(),
			"ec2_vpcdhcp_options_association_resource": ec2.ResourceEC2VPCDHCPOptionsAssociation(),
			"ec2_vpc_endpoint_resource": ec2.ResourceEC2VPCEndpoint(),
			"ec2_vpc_endpoint_connection_notification_resource": ec2.ResourceEC2VPCEndpointConnectionNotification(),
			"ec2_vpc_endpoint_service_resource": ec2.ResourceEC2VPCEndpointService(),
			"ec2_vpc_endpoint_service_permissions_resource": ec2.ResourceEC2VPCEndpointServicePermissions(),
			"ec2_vpc_gateway_attachment_resource": ec2.ResourceEC2VPCGatewayAttachment(),
			"ec2_vpc_peering_connection_resource": ec2.ResourceEC2VPCPeeringConnection(),
			"ec2_vpn_connection_resource": ec2.ResourceEC2VPNConnection(),
			"ec2_vpn_connection_route_resource": ec2.ResourceEC2VPNConnectionRoute(),
			"ec2_vpn_gateway_resource": ec2.ResourceEC2VPNGateway(),
			"ec2_vpn_gateway_route_propagation_resource": ec2.ResourceEC2VPNGatewayRoutePropagation(),
			"ec2_volume_resource": ec2.ResourceEC2Volume(),
			"ec2_volume_attachment_resource": ec2.ResourceEC2VolumeAttachment(),
			"ecr_repository_resource": ecr.ResourceECRRepository(),
			"ecs_cluster_resource": ecs.ResourceECSCluster(),
			"ecs_service_resource": ecs.ResourceECSService(),
			"ecs_task_definition_resource": ecs.ResourceECSTaskDefinition(),
			"efs_file_system_resource": efs.ResourceEFSFileSystem(),
			"efs_mount_target_resource": efs.ResourceEFSMountTarget(),
			"eks_cluster_resource": eks.ResourceEKSCluster(),
			"elasti_cache_cache_cluster_resource": elasticache.ResourceElastiCacheCacheCluster(),
			"elasti_cache_parameter_group_resource": elasticache.ResourceElastiCacheParameterGroup(),
			"elasti_cache_replication_group_resource": elasticache.ResourceElastiCacheReplicationGroup(),
			"elasti_cache_security_group_resource": elasticache.ResourceElastiCacheSecurityGroup(),
			"elasti_cache_security_group_ingress_resource": elasticache.ResourceElastiCacheSecurityGroupIngress(),
			"elasti_cache_subnet_group_resource": elasticache.ResourceElastiCacheSubnetGroup(),
			"elastic_beanstalk_application_resource": elasticbeanstalk.ResourceElasticBeanstalkApplication(),
			"elastic_beanstalk_application_version_resource": elasticbeanstalk.ResourceElasticBeanstalkApplicationVersion(),
			"elastic_beanstalk_configuration_template_resource": elasticbeanstalk.ResourceElasticBeanstalkConfigurationTemplate(),
			"elastic_beanstalk_environment_resource": elasticbeanstalk.ResourceElasticBeanstalkEnvironment(),
			"elastic_load_balancing_load_balancer_resource": elasticloadbalancing.ResourceElasticLoadBalancingLoadBalancer(),
			"elastic_load_balancing_v2_listener_resource": elasticloadbalancingv2.ResourceElasticLoadBalancingV2Listener(),
			"elastic_load_balancing_v2_listener_certificate_resource": elasticloadbalancingv2.ResourceElasticLoadBalancingV2ListenerCertificate(),
			"elastic_load_balancing_v2_listener_rule_resource": elasticloadbalancingv2.ResourceElasticLoadBalancingV2ListenerRule(),
			"elastic_load_balancing_v2_load_balancer_resource": elasticloadbalancingv2.ResourceElasticLoadBalancingV2LoadBalancer(),
			"elastic_load_balancing_v2_target_group_resource": elasticloadbalancingv2.ResourceElasticLoadBalancingV2TargetGroup(),
			"elasticsearch_domain_resource": elasticsearch.ResourceElasticsearchDomain(),
			"emr_cluster_resource": emr.ResourceEMRCluster(),
			"emr_instance_fleet_config_resource": emr.ResourceEMRInstanceFleetConfig(),
			"emr_instance_group_config_resource": emr.ResourceEMRInstanceGroupConfig(),
			"emr_security_configuration_resource": emr.ResourceEMRSecurityConfiguration(),
			"emr_step_resource": emr.ResourceEMRStep(),
			"events_event_bus_resource": events.ResourceEventsEventBus(),
			"events_event_bus_policy_resource": events.ResourceEventsEventBusPolicy(),
			"events_rule_resource": events.ResourceEventsRule(),
			"fsx_file_system_resource": fsx.ResourceFSxFileSystem(),
			"game_lift_alias_resource": gamelift.ResourceGameLiftAlias(),
			"game_lift_build_resource": gamelift.ResourceGameLiftBuild(),
			"game_lift_fleet_resource": gamelift.ResourceGameLiftFleet(),
			"glue_classifier_resource": glue.ResourceGlueClassifier(),
			"glue_connection_resource": glue.ResourceGlueConnection(),
			"glue_crawler_resource": glue.ResourceGlueCrawler(),
			"glue_data_catalog_encryption_settings_resource": glue.ResourceGlueDataCatalogEncryptionSettings(),
			"glue_database_resource": glue.ResourceGlueDatabase(),
			"glue_dev_endpoint_resource": glue.ResourceGlueDevEndpoint(),
			"glue_job_resource": glue.ResourceGlueJob(),
			"glue_ml_transform_resource": glue.ResourceGlueMLTransform(),
			"glue_partition_resource": glue.ResourceGluePartition(),
			"glue_security_configuration_resource": glue.ResourceGlueSecurityConfiguration(),
			"glue_table_resource": glue.ResourceGlueTable(),
			"glue_trigger_resource": glue.ResourceGlueTrigger(),
			"glue_workflow_resource": glue.ResourceGlueWorkflow(),
			"greengrass_connector_definition_resource": greengrass.ResourceGreengrassConnectorDefinition(),
			"greengrass_connector_definition_version_resource": greengrass.ResourceGreengrassConnectorDefinitionVersion(),
			"greengrass_core_definition_resource": greengrass.ResourceGreengrassCoreDefinition(),
			"greengrass_core_definition_version_resource": greengrass.ResourceGreengrassCoreDefinitionVersion(),
			"greengrass_device_definition_resource": greengrass.ResourceGreengrassDeviceDefinition(),
			"greengrass_device_definition_version_resource": greengrass.ResourceGreengrassDeviceDefinitionVersion(),
			"greengrass_function_definition_resource": greengrass.ResourceGreengrassFunctionDefinition(),
			"greengrass_function_definition_version_resource": greengrass.ResourceGreengrassFunctionDefinitionVersion(),
			"greengrass_group_resource": greengrass.ResourceGreengrassGroup(),
			"greengrass_group_version_resource": greengrass.ResourceGreengrassGroupVersion(),
			"greengrass_logger_definition_resource": greengrass.ResourceGreengrassLoggerDefinition(),
			"greengrass_logger_definition_version_resource": greengrass.ResourceGreengrassLoggerDefinitionVersion(),
			"greengrass_resource_definition_resource": greengrass.ResourceGreengrassResourceDefinition(),
			"greengrass_resource_definition_version_resource": greengrass.ResourceGreengrassResourceDefinitionVersion(),
			"greengrass_subscription_definition_resource": greengrass.ResourceGreengrassSubscriptionDefinition(),
			"greengrass_subscription_definition_version_resource": greengrass.ResourceGreengrassSubscriptionDefinitionVersion(),
			"guard_duty_detector_resource": guardduty.ResourceGuardDutyDetector(),
			"guard_duty_filter_resource": guardduty.ResourceGuardDutyFilter(),
			"guard_duty_ip_set_resource": guardduty.ResourceGuardDutyIPSet(),
			"guard_duty_master_resource": guardduty.ResourceGuardDutyMaster(),
			"guard_duty_member_resource": guardduty.ResourceGuardDutyMember(),
			"guard_duty_threat_intel_set_resource": guardduty.ResourceGuardDutyThreatIntelSet(),
			"iam_access_key_resource": iam.ResourceIAMAccessKey(),
			"iam_group_resource": iam.ResourceIAMGroup(),
			"iam_instance_profile_resource": iam.ResourceIAMInstanceProfile(),
			"iam_managed_policy_resource": iam.ResourceIAMManagedPolicy(),
			"iam_policy_resource": iam.ResourceIAMPolicy(),
			"iam_role_resource": iam.ResourceIAMRole(),
			"iam_service_linked_role_resource": iam.ResourceIAMServiceLinkedRole(),
			"iam_user_resource": iam.ResourceIAMUser(),
			"iam_user_to_group_addition_resource": iam.ResourceIAMUserToGroupAddition(),
			"inspector_assessment_target_resource": inspector.ResourceInspectorAssessmentTarget(),
			"inspector_assessment_template_resource": inspector.ResourceInspectorAssessmentTemplate(),
			"inspector_resource_group_resource": inspector.ResourceInspectorResourceGroup(),
			"iot_certificate_resource": iot.ResourceIoTCertificate(),
			"iot_policy_resource": iot.ResourceIoTPolicy(),
			"iot_policy_principal_attachment_resource": iot.ResourceIoTPolicyPrincipalAttachment(),
			"iot_thing_resource": iot.ResourceIoTThing(),
			"iot_thing_principal_attachment_resource": iot.ResourceIoTThingPrincipalAttachment(),
			"iot_topic_rule_resource": iot.ResourceIoTTopicRule(),
			"iot1_click_device_resource": iot1click.ResourceIoT1ClickDevice(),
			"iot1_click_placement_resource": iot1click.ResourceIoT1ClickPlacement(),
			"iot1_click_project_resource": iot1click.ResourceIoT1ClickProject(),
			"iot_analytics_channel_resource": iotanalytics.ResourceIoTAnalyticsChannel(),
			"iot_analytics_dataset_resource": iotanalytics.ResourceIoTAnalyticsDataset(),
			"iot_analytics_datastore_resource": iotanalytics.ResourceIoTAnalyticsDatastore(),
			"iot_analytics_pipeline_resource": iotanalytics.ResourceIoTAnalyticsPipeline(),
			"iot_events_detector_model_resource": iotevents.ResourceIoTEventsDetectorModel(),
			"iot_events_input_resource": iotevents.ResourceIoTEventsInput(),
			"iot_things_graph_flow_template_resource": iotthingsgraph.ResourceIoTThingsGraphFlowTemplate(),
			"kinesis_stream_resource": kinesis.ResourceKinesisStream(),
			"kinesis_stream_consumer_resource": kinesis.ResourceKinesisStreamConsumer(),
			"kinesis_analytics_application_resource": kinesisanalytics.ResourceKinesisAnalyticsApplication(),
			"kinesis_analytics_application_output_resource": kinesisanalytics.ResourceKinesisAnalyticsApplicationOutput(),
			"kinesis_analytics_application_reference_data_source_resource": kinesisanalytics.ResourceKinesisAnalyticsApplicationReferenceDataSource(),
			"kinesis_analytics_v2_application_resource": kinesisanalyticsv2.ResourceKinesisAnalyticsV2Application(),
			"kinesis_analytics_v2_application_cloud_watch_logging_option_resource": kinesisanalyticsv2.ResourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption(),
			"kinesis_analytics_v2_application_output_resource": kinesisanalyticsv2.ResourceKinesisAnalyticsV2ApplicationOutput(),
			"kinesis_analytics_v2_application_reference_data_source_resource": kinesisanalyticsv2.ResourceKinesisAnalyticsV2ApplicationReferenceDataSource(),
			"kinesis_firehose_delivery_stream_resource": kinesisfirehose.ResourceKinesisFirehoseDeliveryStream(),
			"kms_alias_resource": kms.ResourceKMSAlias(),
			"kms_key_resource": kms.ResourceKMSKey(),
			"lake_formation_data_lake_settings_resource": lakeformation.ResourceLakeFormationDataLakeSettings(),
			"lake_formation_permissions_resource": lakeformation.ResourceLakeFormationPermissions(),
			"lake_formation_resource_resource": lakeformation.ResourceLakeFormationResource(),
			"lambda_alias_resource": lambda.ResourceLambdaAlias(),
			"lambda_event_source_mapping_resource": lambda.ResourceLambdaEventSourceMapping(),
			"lambda_function_resource": lambda.ResourceLambdaFunction(),
			"lambda_layer_version_resource": lambda.ResourceLambdaLayerVersion(),
			"lambda_layer_version_permission_resource": lambda.ResourceLambdaLayerVersionPermission(),
			"lambda_permission_resource": lambda.ResourceLambdaPermission(),
			"lambda_version_resource": lambda.ResourceLambdaVersion(),
			"logs_destination_resource": logs.ResourceLogsDestination(),
			"logs_log_group_resource": logs.ResourceLogsLogGroup(),
			"logs_log_stream_resource": logs.ResourceLogsLogStream(),
			"logs_metric_filter_resource": logs.ResourceLogsMetricFilter(),
			"logs_subscription_filter_resource": logs.ResourceLogsSubscriptionFilter(),
			"managed_blockchain_member_resource": managedblockchain.ResourceManagedBlockchainMember(),
			"managed_blockchain_node_resource": managedblockchain.ResourceManagedBlockchainNode(),
			"media_live_channel_resource": medialive.ResourceMediaLiveChannel(),
			"media_live_input_resource": medialive.ResourceMediaLiveInput(),
			"media_live_input_security_group_resource": medialive.ResourceMediaLiveInputSecurityGroup(),
			"media_store_container_resource": mediastore.ResourceMediaStoreContainer(),
			"msk_cluster_resource": msk.ResourceMSKCluster(),
			"neptune_db_cluster_resource": neptune.ResourceNeptuneDBCluster(),
			"neptune_db_cluster_parameter_group_resource": neptune.ResourceNeptuneDBClusterParameterGroup(),
			"neptune_db_instance_resource": neptune.ResourceNeptuneDBInstance(),
			"neptune_db_parameter_group_resource": neptune.ResourceNeptuneDBParameterGroup(),
			"neptune_db_subnet_group_resource": neptune.ResourceNeptuneDBSubnetGroup(),
			"ops_works_app_resource": opsworks.ResourceOpsWorksApp(),
			"ops_works_elastic_load_balancer_attachment_resource": opsworks.ResourceOpsWorksElasticLoadBalancerAttachment(),
			"ops_works_instance_resource": opsworks.ResourceOpsWorksInstance(),
			"ops_works_layer_resource": opsworks.ResourceOpsWorksLayer(),
			"ops_works_stack_resource": opsworks.ResourceOpsWorksStack(),
			"ops_works_user_profile_resource": opsworks.ResourceOpsWorksUserProfile(),
			"ops_works_volume_resource": opsworks.ResourceOpsWorksVolume(),
			"ops_works_cm_server_resource": opsworkscm.ResourceOpsWorksCMServer(),
			"pinpoint_adm_channel_resource": pinpoint.ResourcePinpointADMChannel(),
			"pinpoint_apns_channel_resource": pinpoint.ResourcePinpointAPNSChannel(),
			"pinpoint_apns_sandbox_channel_resource": pinpoint.ResourcePinpointAPNSSandboxChannel(),
			"pinpoint_apns_voip_channel_resource": pinpoint.ResourcePinpointAPNSVoipChannel(),
			"pinpoint_apns_voip_sandbox_channel_resource": pinpoint.ResourcePinpointAPNSVoipSandboxChannel(),
			"pinpoint_app_resource": pinpoint.ResourcePinpointApp(),
			"pinpoint_application_settings_resource": pinpoint.ResourcePinpointApplicationSettings(),
			"pinpoint_baidu_channel_resource": pinpoint.ResourcePinpointBaiduChannel(),
			"pinpoint_campaign_resource": pinpoint.ResourcePinpointCampaign(),
			"pinpoint_email_channel_resource": pinpoint.ResourcePinpointEmailChannel(),
			"pinpoint_email_template_resource": pinpoint.ResourcePinpointEmailTemplate(),
			"pinpoint_event_stream_resource": pinpoint.ResourcePinpointEventStream(),
			"pinpoint_gcm_channel_resource": pinpoint.ResourcePinpointGCMChannel(),
			"pinpoint_push_template_resource": pinpoint.ResourcePinpointPushTemplate(),
			"pinpoint_sms_channel_resource": pinpoint.ResourcePinpointSMSChannel(),
			"pinpoint_segment_resource": pinpoint.ResourcePinpointSegment(),
			"pinpoint_sms_template_resource": pinpoint.ResourcePinpointSmsTemplate(),
			"pinpoint_voice_channel_resource": pinpoint.ResourcePinpointVoiceChannel(),
			"pinpoint_email_configuration_set_resource": pinpointemail.ResourcePinpointEmailConfigurationSet(),
			"pinpoint_email_configuration_set_event_destination_resource": pinpointemail.ResourcePinpointEmailConfigurationSetEventDestination(),
			"pinpoint_email_dedicated_ip_pool_resource": pinpointemail.ResourcePinpointEmailDedicatedIpPool(),
			"pinpoint_email_identity_resource": pinpointemail.ResourcePinpointEmailIdentity(),
			"qldb_ledger_resource": qldb.ResourceQLDBLedger(),
			"ram_resource_share_resource": ram.ResourceRAMResourceShare(),
			"rdsdb_cluster_resource": rds.ResourceRDSDBCluster(),
			"rdsdb_cluster_parameter_group_resource": rds.ResourceRDSDBClusterParameterGroup(),
			"rdsdb_instance_resource": rds.ResourceRDSDBInstance(),
			"rdsdb_parameter_group_resource": rds.ResourceRDSDBParameterGroup(),
			"rdsdb_security_group_resource": rds.ResourceRDSDBSecurityGroup(),
			"rdsdb_security_group_ingress_resource": rds.ResourceRDSDBSecurityGroupIngress(),
			"rdsdb_subnet_group_resource": rds.ResourceRDSDBSubnetGroup(),
			"rds_event_subscription_resource": rds.ResourceRDSEventSubscription(),
			"rds_option_group_resource": rds.ResourceRDSOptionGroup(),
			"redshift_cluster_resource": redshift.ResourceRedshiftCluster(),
			"redshift_cluster_parameter_group_resource": redshift.ResourceRedshiftClusterParameterGroup(),
			"redshift_cluster_security_group_resource": redshift.ResourceRedshiftClusterSecurityGroup(),
			"redshift_cluster_security_group_ingress_resource": redshift.ResourceRedshiftClusterSecurityGroupIngress(),
			"redshift_cluster_subnet_group_resource": redshift.ResourceRedshiftClusterSubnetGroup(),
			"robo_maker_fleet_resource": robomaker.ResourceRoboMakerFleet(),
			"robo_maker_robot_resource": robomaker.ResourceRoboMakerRobot(),
			"robo_maker_robot_application_resource": robomaker.ResourceRoboMakerRobotApplication(),
			"robo_maker_robot_application_version_resource": robomaker.ResourceRoboMakerRobotApplicationVersion(),
			"robo_maker_simulation_application_resource": robomaker.ResourceRoboMakerSimulationApplication(),
			"robo_maker_simulation_application_version_resource": robomaker.ResourceRoboMakerSimulationApplicationVersion(),
			"route53_health_check_resource": route53.ResourceRoute53HealthCheck(),
			"route53_hosted_zone_resource": route53.ResourceRoute53HostedZone(),
			"route53_record_set_resource": route53.ResourceRoute53RecordSet(),
			"route53_record_set_group_resource": route53.ResourceRoute53RecordSetGroup(),
			"route53_resolver_resolver_endpoint_resource": route53resolver.ResourceRoute53ResolverResolverEndpoint(),
			"route53_resolver_resolver_rule_resource": route53resolver.ResourceRoute53ResolverResolverRule(),
			"route53_resolver_resolver_rule_association_resource": route53resolver.ResourceRoute53ResolverResolverRuleAssociation(),
			"s3_bucket_resource": s3.ResourceS3Bucket(),
			"s3_bucket_policy_resource": s3.ResourceS3BucketPolicy(),
			"sage_maker_code_repository_resource": sagemaker.ResourceSageMakerCodeRepository(),
			"sage_maker_endpoint_resource": sagemaker.ResourceSageMakerEndpoint(),
			"sage_maker_endpoint_config_resource": sagemaker.ResourceSageMakerEndpointConfig(),
			"sage_maker_model_resource": sagemaker.ResourceSageMakerModel(),
			"sage_maker_notebook_instance_resource": sagemaker.ResourceSageMakerNotebookInstance(),
			"sage_maker_notebook_instance_lifecycle_config_resource": sagemaker.ResourceSageMakerNotebookInstanceLifecycleConfig(),
			"sage_maker_workteam_resource": sagemaker.ResourceSageMakerWorkteam(),
			"sdb_domain_resource": sdb.ResourceSDBDomain(),
			"secrets_manager_resource_policy_resource": secretsmanager.ResourceSecretsManagerResourcePolicy(),
			"secrets_manager_rotation_schedule_resource": secretsmanager.ResourceSecretsManagerRotationSchedule(),
			"secrets_manager_secret_resource": secretsmanager.ResourceSecretsManagerSecret(),
			"secrets_manager_secret_target_attachment_resource": secretsmanager.ResourceSecretsManagerSecretTargetAttachment(),
			"security_hub_hub_resource": securityhub.ResourceSecurityHubHub(),
			"service_catalog_accepted_portfolio_share_resource": servicecatalog.ResourceServiceCatalogAcceptedPortfolioShare(),
			"service_catalog_cloud_formation_product_resource": servicecatalog.ResourceServiceCatalogCloudFormationProduct(),
			"service_catalog_cloud_formation_provisioned_product_resource": servicecatalog.ResourceServiceCatalogCloudFormationProvisionedProduct(),
			"service_catalog_launch_notification_constraint_resource": servicecatalog.ResourceServiceCatalogLaunchNotificationConstraint(),
			"service_catalog_launch_role_constraint_resource": servicecatalog.ResourceServiceCatalogLaunchRoleConstraint(),
			"service_catalog_launch_template_constraint_resource": servicecatalog.ResourceServiceCatalogLaunchTemplateConstraint(),
			"service_catalog_portfolio_resource": servicecatalog.ResourceServiceCatalogPortfolio(),
			"service_catalog_portfolio_principal_association_resource": servicecatalog.ResourceServiceCatalogPortfolioPrincipalAssociation(),
			"service_catalog_portfolio_product_association_resource": servicecatalog.ResourceServiceCatalogPortfolioProductAssociation(),
			"service_catalog_portfolio_share_resource": servicecatalog.ResourceServiceCatalogPortfolioShare(),
			"service_catalog_resource_update_constraint_resource": servicecatalog.ResourceServiceCatalogResourceUpdateConstraint(),
			"service_catalog_stack_set_constraint_resource": servicecatalog.ResourceServiceCatalogStackSetConstraint(),
			"service_catalog_tag_option_resource": servicecatalog.ResourceServiceCatalogTagOption(),
			"service_catalog_tag_option_association_resource": servicecatalog.ResourceServiceCatalogTagOptionAssociation(),
			"service_discovery_http_namespace_resource": servicediscovery.ResourceServiceDiscoveryHttpNamespace(),
			"service_discovery_instance_resource": servicediscovery.ResourceServiceDiscoveryInstance(),
			"service_discovery_private_dns_namespace_resource": servicediscovery.ResourceServiceDiscoveryPrivateDnsNamespace(),
			"service_discovery_public_dns_namespace_resource": servicediscovery.ResourceServiceDiscoveryPublicDnsNamespace(),
			"service_discovery_service_resource": servicediscovery.ResourceServiceDiscoveryService(),
			"ses_configuration_set_resource": ses.ResourceSESConfigurationSet(),
			"ses_configuration_set_event_destination_resource": ses.ResourceSESConfigurationSetEventDestination(),
			"ses_receipt_filter_resource": ses.ResourceSESReceiptFilter(),
			"ses_receipt_rule_resource": ses.ResourceSESReceiptRule(),
			"ses_receipt_rule_set_resource": ses.ResourceSESReceiptRuleSet(),
			"ses_template_resource": ses.ResourceSESTemplate(),
			"sns_subscription_resource": sns.ResourceSNSSubscription(),
			"sns_topic_resource": sns.ResourceSNSTopic(),
			"sns_topic_policy_resource": sns.ResourceSNSTopicPolicy(),
			"sqs_queue_resource": sqs.ResourceSQSQueue(),
			"sqs_queue_policy_resource": sqs.ResourceSQSQueuePolicy(),
			"ssm_association_resource": ssm.ResourceSSMAssociation(),
			"ssm_document_resource": ssm.ResourceSSMDocument(),
			"ssm_maintenance_window_resource": ssm.ResourceSSMMaintenanceWindow(),
			"ssm_maintenance_window_target_resource": ssm.ResourceSSMMaintenanceWindowTarget(),
			"ssm_maintenance_window_task_resource": ssm.ResourceSSMMaintenanceWindowTask(),
			"ssm_parameter_resource": ssm.ResourceSSMParameter(),
			"ssm_patch_baseline_resource": ssm.ResourceSSMPatchBaseline(),
			"ssm_resource_data_sync_resource": ssm.ResourceSSMResourceDataSync(),
			"step_functions_activity_resource": stepfunctions.ResourceStepFunctionsActivity(),
			"step_functions_state_machine_resource": stepfunctions.ResourceStepFunctionsStateMachine(),
			"transfer_server_resource": transfer.ResourceTransferServer(),
			"transfer_user_resource": transfer.ResourceTransferUser(),
			"waf_byte_match_set_resource": waf.ResourceWAFByteMatchSet(),
			"wafip_set_resource": waf.ResourceWAFIPSet(),
			"waf_rule_resource": waf.ResourceWAFRule(),
			"waf_size_constraint_set_resource": waf.ResourceWAFSizeConstraintSet(),
			"waf_sql_injection_match_set_resource": waf.ResourceWAFSqlInjectionMatchSet(),
			"waf_web_acl_resource": waf.ResourceWAFWebACL(),
			"waf_xss_match_set_resource": waf.ResourceWAFXssMatchSet(),
			"waf_regional_byte_match_set_resource": wafregional.ResourceWAFRegionalByteMatchSet(),
			"waf_regional_geo_match_set_resource": wafregional.ResourceWAFRegionalGeoMatchSet(),
			"waf_regional_ip_set_resource": wafregional.ResourceWAFRegionalIPSet(),
			"waf_regional_rate_based_rule_resource": wafregional.ResourceWAFRegionalRateBasedRule(),
			"waf_regional_regex_pattern_set_resource": wafregional.ResourceWAFRegionalRegexPatternSet(),
			"waf_regional_rule_resource": wafregional.ResourceWAFRegionalRule(),
			"waf_regional_size_constraint_set_resource": wafregional.ResourceWAFRegionalSizeConstraintSet(),
			"waf_regional_sql_injection_match_set_resource": wafregional.ResourceWAFRegionalSqlInjectionMatchSet(),
			"waf_regional_web_acl_resource": wafregional.ResourceWAFRegionalWebACL(),
			"waf_regional_web_acl_association_resource": wafregional.ResourceWAFRegionalWebACLAssociation(),
			"waf_regional_xss_match_set_resource": wafregional.ResourceWAFRegionalXssMatchSet(),
			"work_spaces_workspace_resource": workspaces.ResourceWorkSpacesWorkspace(),
		},
	}
}