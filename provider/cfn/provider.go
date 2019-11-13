// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   https://github.com/UnquietCode/terraform-provider-cloudformation

package cfn

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
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
		ConfigureFunc: plugin.ProviderConfigure,
		Schema: map[string]*schema.Schema{
			"workdir": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "working directory on the filesystem",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cfn_template_data": misc.ResourceTemplateData(),
			"cfn_amazonmq_broker": amazonmq.ResourceAmazonMQBroker(),
			"cfn_amazonmq_configuration": amazonmq.ResourceAmazonMQConfiguration(),
			"cfn_amazonmq_configuration_association": amazonmq.ResourceAmazonMQConfigurationAssociation(),
			"cfn_amplify_app": amplify.ResourceAmplifyApp(),
			"cfn_amplify_branch": amplify.ResourceAmplifyBranch(),
			"cfn_amplify_domain": amplify.ResourceAmplifyDomain(),
			"cfn_apigateway_account": apigateway.ResourceApiGatewayAccount(),
			"cfn_apigateway_api_key": apigateway.ResourceApiGatewayApiKey(),
			"cfn_apigateway_authorizer": apigateway.ResourceApiGatewayAuthorizer(),
			"cfn_apigateway_base_path_mapping": apigateway.ResourceApiGatewayBasePathMapping(),
			"cfn_apigateway_client_certificate": apigateway.ResourceApiGatewayClientCertificate(),
			"cfn_apigateway_deployment": apigateway.ResourceApiGatewayDeployment(),
			"cfn_apigateway_documentation_part": apigateway.ResourceApiGatewayDocumentationPart(),
			"cfn_apigateway_documentation_version": apigateway.ResourceApiGatewayDocumentationVersion(),
			"cfn_apigateway_domain_name": apigateway.ResourceApiGatewayDomainName(),
			"cfn_apigateway_gateway_response": apigateway.ResourceApiGatewayGatewayResponse(),
			"cfn_apigateway_method": apigateway.ResourceApiGatewayMethod(),
			"cfn_apigateway_model": apigateway.ResourceApiGatewayModel(),
			"cfn_apigateway_request_validator": apigateway.ResourceApiGatewayRequestValidator(),
			"cfn_apigateway_resource": apigateway.ResourceApiGatewayResource(),
			"cfn_apigateway_rest_api": apigateway.ResourceApiGatewayRestApi(),
			"cfn_apigateway_stage": apigateway.ResourceApiGatewayStage(),
			"cfn_apigateway_usage_plan": apigateway.ResourceApiGatewayUsagePlan(),
			"cfn_apigateway_usage_plan_key": apigateway.ResourceApiGatewayUsagePlanKey(),
			"cfn_apigateway_vpc_link": apigateway.ResourceApiGatewayVpcLink(),
			"cfn_apigatewayv2_api": apigatewayv2.ResourceApiGatewayV2Api(),
			"cfn_apigatewayv2_api_mapping": apigatewayv2.ResourceApiGatewayV2ApiMapping(),
			"cfn_apigatewayv2_authorizer": apigatewayv2.ResourceApiGatewayV2Authorizer(),
			"cfn_apigatewayv2_deployment": apigatewayv2.ResourceApiGatewayV2Deployment(),
			"cfn_apigatewayv2_domain_name": apigatewayv2.ResourceApiGatewayV2DomainName(),
			"cfn_apigatewayv2_integration": apigatewayv2.ResourceApiGatewayV2Integration(),
			"cfn_apigatewayv2_integration_response": apigatewayv2.ResourceApiGatewayV2IntegrationResponse(),
			"cfn_apigatewayv2_model": apigatewayv2.ResourceApiGatewayV2Model(),
			"cfn_apigatewayv2_route": apigatewayv2.ResourceApiGatewayV2Route(),
			"cfn_apigatewayv2_route_response": apigatewayv2.ResourceApiGatewayV2RouteResponse(),
			"cfn_apigatewayv2_stage": apigatewayv2.ResourceApiGatewayV2Stage(),
			"cfn_applicationautoscaling_scalable_target": applicationautoscaling.ResourceApplicationAutoScalingScalableTarget(),
			"cfn_applicationautoscaling_scaling_policy": applicationautoscaling.ResourceApplicationAutoScalingScalingPolicy(),
			"cfn_appmesh_mesh": appmesh.ResourceAppMeshMesh(),
			"cfn_appmesh_route": appmesh.ResourceAppMeshRoute(),
			"cfn_appmesh_virtual_node": appmesh.ResourceAppMeshVirtualNode(),
			"cfn_appmesh_virtual_router": appmesh.ResourceAppMeshVirtualRouter(),
			"cfn_appmesh_virtual_service": appmesh.ResourceAppMeshVirtualService(),
			"cfn_appstream_directory_config": appstream.ResourceAppStreamDirectoryConfig(),
			"cfn_appstream_fleet": appstream.ResourceAppStreamFleet(),
			"cfn_appstream_image_builder": appstream.ResourceAppStreamImageBuilder(),
			"cfn_appstream_stack": appstream.ResourceAppStreamStack(),
			"cfn_appstream_stack_fleet_association": appstream.ResourceAppStreamStackFleetAssociation(),
			"cfn_appstream_stack_user_association": appstream.ResourceAppStreamStackUserAssociation(),
			"cfn_appstream_user": appstream.ResourceAppStreamUser(),
			"cfn_appsync_api_key": appsync.ResourceAppSyncApiKey(),
			"cfn_appsync_data_source": appsync.ResourceAppSyncDataSource(),
			"cfn_appsync_function_configuration": appsync.ResourceAppSyncFunctionConfiguration(),
			"cfn_appsync_graph_ql_api": appsync.ResourceAppSyncGraphQLApi(),
			"cfn_appsync_graph_ql_schema": appsync.ResourceAppSyncGraphQLSchema(),
			"cfn_appsync_resolver": appsync.ResourceAppSyncResolver(),
			"cfn_athena_named_query": athena.ResourceAthenaNamedQuery(),
			"cfn_autoscaling_auto_scaling_group": autoscaling.ResourceAutoScalingAutoScalingGroup(),
			"cfn_autoscaling_launch_configuration": autoscaling.ResourceAutoScalingLaunchConfiguration(),
			"cfn_autoscaling_lifecycle_hook": autoscaling.ResourceAutoScalingLifecycleHook(),
			"cfn_autoscaling_scaling_policy": autoscaling.ResourceAutoScalingScalingPolicy(),
			"cfn_autoscaling_scheduled_action": autoscaling.ResourceAutoScalingScheduledAction(),
			"cfn_autoscalingplans_scaling_plan": autoscalingplans.ResourceAutoScalingPlansScalingPlan(),
			"cfn_backup_backup_plan": backup.ResourceBackupBackupPlan(),
			"cfn_backup_backup_selection": backup.ResourceBackupBackupSelection(),
			"cfn_backup_backup_vault": backup.ResourceBackupBackupVault(),
			"cfn_batch_compute_environment": batch.ResourceBatchComputeEnvironment(),
			"cfn_batch_job_definition": batch.ResourceBatchJobDefinition(),
			"cfn_batch_job_queue": batch.ResourceBatchJobQueue(),
			"cfn_budgets_budget": budgets.ResourceBudgetsBudget(),
			"cfn_certificatemanager_certificate": certificatemanager.ResourceCertificateManagerCertificate(),
			"cfn_cloud9_environment_ec2": cloud9.ResourceCloud9EnvironmentEC2(),
			"cfn_cloudformation_custom_resource": cloudformation.ResourceCloudFormationCustomResource(),
			"cfn_cloudformation_macro": cloudformation.ResourceCloudFormationMacro(),
			"cfn_cloudformation_stack": cloudformation.ResourceCloudFormationStack(),
			"cfn_cloudformation_wait_condition": cloudformation.ResourceCloudFormationWaitCondition(),
			"cfn_cloudformation_wait_condition_handle": cloudformation.ResourceCloudFormationWaitConditionHandle(),
			"cfn_cloudfront_cloud_front_origin_access_identity": cloudfront.ResourceCloudFrontCloudFrontOriginAccessIdentity(),
			"cfn_cloudfront_distribution": cloudfront.ResourceCloudFrontDistribution(),
			"cfn_cloudfront_streaming_distribution": cloudfront.ResourceCloudFrontStreamingDistribution(),
			"cfn_cloudtrail_trail": cloudtrail.ResourceCloudTrailTrail(),
			"cfn_cloudwatch_alarm": cloudwatch.ResourceCloudWatchAlarm(),
			"cfn_cloudwatch_anomaly_detector": cloudwatch.ResourceCloudWatchAnomalyDetector(),
			"cfn_cloudwatch_dashboard": cloudwatch.ResourceCloudWatchDashboard(),
			"cfn_codebuild_project": codebuild.ResourceCodeBuildProject(),
			"cfn_codebuild_source_credential": codebuild.ResourceCodeBuildSourceCredential(),
			"cfn_codecommit_repository": codecommit.ResourceCodeCommitRepository(),
			"cfn_codedeploy_application": codedeploy.ResourceCodeDeployApplication(),
			"cfn_codedeploy_deployment_config": codedeploy.ResourceCodeDeployDeploymentConfig(),
			"cfn_codedeploy_deployment_group": codedeploy.ResourceCodeDeployDeploymentGroup(),
			"cfn_codepipeline_custom_action_type": codepipeline.ResourceCodePipelineCustomActionType(),
			"cfn_codepipeline_pipeline": codepipeline.ResourceCodePipelinePipeline(),
			"cfn_codepipeline_webhook": codepipeline.ResourceCodePipelineWebhook(),
			"cfn_codestar_git_hub_repository": codestar.ResourceCodeStarGitHubRepository(),
			"cfn_cognito_identity_pool": cognito.ResourceCognitoIdentityPool(),
			"cfn_cognito_identity_pool_role_attachment": cognito.ResourceCognitoIdentityPoolRoleAttachment(),
			"cfn_cognito_user_pool": cognito.ResourceCognitoUserPool(),
			"cfn_cognito_user_pool_client": cognito.ResourceCognitoUserPoolClient(),
			"cfn_cognito_user_pool_domain": cognito.ResourceCognitoUserPoolDomain(),
			"cfn_cognito_user_pool_group": cognito.ResourceCognitoUserPoolGroup(),
			"cfn_cognito_user_pool_identity_provider": cognito.ResourceCognitoUserPoolIdentityProvider(),
			"cfn_cognito_user_pool_resource_server": cognito.ResourceCognitoUserPoolResourceServer(),
			"cfn_cognito_user_pool_risk_configuration_attachment": cognito.ResourceCognitoUserPoolRiskConfigurationAttachment(),
			"cfn_cognito_user_pool_ui_customization_attachment": cognito.ResourceCognitoUserPoolUICustomizationAttachment(),
			"cfn_cognito_user_pool_user": cognito.ResourceCognitoUserPoolUser(),
			"cfn_cognito_user_pool_user_to_group_attachment": cognito.ResourceCognitoUserPoolUserToGroupAttachment(),
			"cfn_config_aggregation_authorization": config.ResourceConfigAggregationAuthorization(),
			"cfn_config_config_rule": config.ResourceConfigConfigRule(),
			"cfn_config_configuration_aggregator": config.ResourceConfigConfigurationAggregator(),
			"cfn_config_configuration_recorder": config.ResourceConfigConfigurationRecorder(),
			"cfn_config_delivery_channel": config.ResourceConfigDeliveryChannel(),
			"cfn_config_organization_config_rule": config.ResourceConfigOrganizationConfigRule(),
			"cfn_config_remediation_configuration": config.ResourceConfigRemediationConfiguration(),
			"cfn_datapipeline_pipeline": datapipeline.ResourceDataPipelinePipeline(),
			"cfn_dax_cluster": dax.ResourceDAXCluster(),
			"cfn_dax_parameter_group": dax.ResourceDAXParameterGroup(),
			"cfn_dax_subnet_group": dax.ResourceDAXSubnetGroup(),
			"cfn_directoryservice_microsoft_ad": directoryservice.ResourceDirectoryServiceMicrosoftAD(),
			"cfn_directoryservice_simple_ad": directoryservice.ResourceDirectoryServiceSimpleAD(),
			"cfn_dlm_lifecycle_policy": dlm.ResourceDLMLifecyclePolicy(),
			"cfn_dms_certificate": dms.ResourceDMSCertificate(),
			"cfn_dms_endpoint": dms.ResourceDMSEndpoint(),
			"cfn_dms_event_subscription": dms.ResourceDMSEventSubscription(),
			"cfn_dms_replication_instance": dms.ResourceDMSReplicationInstance(),
			"cfn_dms_replication_subnet_group": dms.ResourceDMSReplicationSubnetGroup(),
			"cfn_dms_replication_task": dms.ResourceDMSReplicationTask(),
			"cfn_docdb_db_cluster": docdb.ResourceDocDBDBCluster(),
			"cfn_docdb_db_cluster_parameter_group": docdb.ResourceDocDBDBClusterParameterGroup(),
			"cfn_docdb_db_instance": docdb.ResourceDocDBDBInstance(),
			"cfn_docdb_db_subnet_group": docdb.ResourceDocDBDBSubnetGroup(),
			"cfn_dynamodb_table": dynamodb.ResourceDynamoDBTable(),
			"cfn_ec2_capacity_reservation": ec2.ResourceEC2CapacityReservation(),
			"cfn_ec2_client_vpn_authorization_rule": ec2.ResourceEC2ClientVpnAuthorizationRule(),
			"cfn_ec2_client_vpn_endpoint": ec2.ResourceEC2ClientVpnEndpoint(),
			"cfn_ec2_client_vpn_route": ec2.ResourceEC2ClientVpnRoute(),
			"cfn_ec2_client_vpn_target_network_association": ec2.ResourceEC2ClientVpnTargetNetworkAssociation(),
			"cfn_ec2_customer_gateway": ec2.ResourceEC2CustomerGateway(),
			"cfn_ec2_dhcp_options": ec2.ResourceEC2DHCPOptions(),
			"cfn_ec2_ec2_fleet": ec2.ResourceEC2EC2Fleet(),
			"cfn_ec2_eip": ec2.ResourceEC2EIP(),
			"cfn_ec2_eip_association": ec2.ResourceEC2EIPAssociation(),
			"cfn_ec2_egress_only_internet_gateway": ec2.ResourceEC2EgressOnlyInternetGateway(),
			"cfn_ec2_flow_log": ec2.ResourceEC2FlowLog(),
			"cfn_ec2_host": ec2.ResourceEC2Host(),
			"cfn_ec2_instance": ec2.ResourceEC2Instance(),
			"cfn_ec2_internet_gateway": ec2.ResourceEC2InternetGateway(),
			"cfn_ec2_launch_template": ec2.ResourceEC2LaunchTemplate(),
			"cfn_ec2_nat_gateway": ec2.ResourceEC2NatGateway(),
			"cfn_ec2_network_acl": ec2.ResourceEC2NetworkAcl(),
			"cfn_ec2_network_acl_entry": ec2.ResourceEC2NetworkAclEntry(),
			"cfn_ec2_network_interface": ec2.ResourceEC2NetworkInterface(),
			"cfn_ec2_network_interface_attachment": ec2.ResourceEC2NetworkInterfaceAttachment(),
			"cfn_ec2_network_interface_permission": ec2.ResourceEC2NetworkInterfacePermission(),
			"cfn_ec2_placement_group": ec2.ResourceEC2PlacementGroup(),
			"cfn_ec2_route": ec2.ResourceEC2Route(),
			"cfn_ec2_route_table": ec2.ResourceEC2RouteTable(),
			"cfn_ec2_security_group": ec2.ResourceEC2SecurityGroup(),
			"cfn_ec2_security_group_egress": ec2.ResourceEC2SecurityGroupEgress(),
			"cfn_ec2_security_group_ingress": ec2.ResourceEC2SecurityGroupIngress(),
			"cfn_ec2_spot_fleet": ec2.ResourceEC2SpotFleet(),
			"cfn_ec2_subnet": ec2.ResourceEC2Subnet(),
			"cfn_ec2_subnet_cidr_block": ec2.ResourceEC2SubnetCidrBlock(),
			"cfn_ec2_subnet_network_acl_association": ec2.ResourceEC2SubnetNetworkAclAssociation(),
			"cfn_ec2_subnet_route_table_association": ec2.ResourceEC2SubnetRouteTableAssociation(),
			"cfn_ec2_traffic_mirror_filter": ec2.ResourceEC2TrafficMirrorFilter(),
			"cfn_ec2_traffic_mirror_filter_rule": ec2.ResourceEC2TrafficMirrorFilterRule(),
			"cfn_ec2_traffic_mirror_session": ec2.ResourceEC2TrafficMirrorSession(),
			"cfn_ec2_traffic_mirror_target": ec2.ResourceEC2TrafficMirrorTarget(),
			"cfn_ec2_transit_gateway": ec2.ResourceEC2TransitGateway(),
			"cfn_ec2_transit_gateway_attachment": ec2.ResourceEC2TransitGatewayAttachment(),
			"cfn_ec2_transit_gateway_route": ec2.ResourceEC2TransitGatewayRoute(),
			"cfn_ec2_transit_gateway_route_table": ec2.ResourceEC2TransitGatewayRouteTable(),
			"cfn_ec2_transit_gateway_route_table_association": ec2.ResourceEC2TransitGatewayRouteTableAssociation(),
			"cfn_ec2_transit_gateway_route_table_propagation": ec2.ResourceEC2TransitGatewayRouteTablePropagation(),
			"cfn_ec2_vpc": ec2.ResourceEC2VPC(),
			"cfn_ec2_vpc_cidr_block": ec2.ResourceEC2VPCCidrBlock(),
			"cfn_ec2_vpcdhcp_options_association": ec2.ResourceEC2VPCDHCPOptionsAssociation(),
			"cfn_ec2_vpc_endpoint": ec2.ResourceEC2VPCEndpoint(),
			"cfn_ec2_vpc_endpoint_connection_notification": ec2.ResourceEC2VPCEndpointConnectionNotification(),
			"cfn_ec2_vpc_endpoint_service": ec2.ResourceEC2VPCEndpointService(),
			"cfn_ec2_vpc_endpoint_service_permissions": ec2.ResourceEC2VPCEndpointServicePermissions(),
			"cfn_ec2_vpc_gateway_attachment": ec2.ResourceEC2VPCGatewayAttachment(),
			"cfn_ec2_vpc_peering_connection": ec2.ResourceEC2VPCPeeringConnection(),
			"cfn_ec2_vpn_connection": ec2.ResourceEC2VPNConnection(),
			"cfn_ec2_vpn_connection_route": ec2.ResourceEC2VPNConnectionRoute(),
			"cfn_ec2_vpn_gateway": ec2.ResourceEC2VPNGateway(),
			"cfn_ec2_vpn_gateway_route_propagation": ec2.ResourceEC2VPNGatewayRoutePropagation(),
			"cfn_ec2_volume": ec2.ResourceEC2Volume(),
			"cfn_ec2_volume_attachment": ec2.ResourceEC2VolumeAttachment(),
			"cfn_ecr_repository": ecr.ResourceECRRepository(),
			"cfn_ecs_cluster": ecs.ResourceECSCluster(),
			"cfn_ecs_service": ecs.ResourceECSService(),
			"cfn_ecs_task_definition": ecs.ResourceECSTaskDefinition(),
			"cfn_efs_file_system": efs.ResourceEFSFileSystem(),
			"cfn_efs_mount_target": efs.ResourceEFSMountTarget(),
			"cfn_eks_cluster": eks.ResourceEKSCluster(),
			"cfn_elasticache_cache_cluster": elasticache.ResourceElastiCacheCacheCluster(),
			"cfn_elasticache_parameter_group": elasticache.ResourceElastiCacheParameterGroup(),
			"cfn_elasticache_replication_group": elasticache.ResourceElastiCacheReplicationGroup(),
			"cfn_elasticache_security_group": elasticache.ResourceElastiCacheSecurityGroup(),
			"cfn_elasticache_security_group_ingress": elasticache.ResourceElastiCacheSecurityGroupIngress(),
			"cfn_elasticache_subnet_group": elasticache.ResourceElastiCacheSubnetGroup(),
			"cfn_elasticbeanstalk_application": elasticbeanstalk.ResourceElasticBeanstalkApplication(),
			"cfn_elasticbeanstalk_application_version": elasticbeanstalk.ResourceElasticBeanstalkApplicationVersion(),
			"cfn_elasticbeanstalk_configuration_template": elasticbeanstalk.ResourceElasticBeanstalkConfigurationTemplate(),
			"cfn_elasticbeanstalk_environment": elasticbeanstalk.ResourceElasticBeanstalkEnvironment(),
			"cfn_elasticloadbalancing_load_balancer": elasticloadbalancing.ResourceElasticLoadBalancingLoadBalancer(),
			"cfn_elasticloadbalancingv2_listener": elasticloadbalancingv2.ResourceElasticLoadBalancingV2Listener(),
			"cfn_elasticloadbalancingv2_listener_certificate": elasticloadbalancingv2.ResourceElasticLoadBalancingV2ListenerCertificate(),
			"cfn_elasticloadbalancingv2_listener_rule": elasticloadbalancingv2.ResourceElasticLoadBalancingV2ListenerRule(),
			"cfn_elasticloadbalancingv2_load_balancer": elasticloadbalancingv2.ResourceElasticLoadBalancingV2LoadBalancer(),
			"cfn_elasticloadbalancingv2_target_group": elasticloadbalancingv2.ResourceElasticLoadBalancingV2TargetGroup(),
			"cfn_elasticsearch_domain": elasticsearch.ResourceElasticsearchDomain(),
			"cfn_emr_cluster": emr.ResourceEMRCluster(),
			"cfn_emr_instance_fleet_config": emr.ResourceEMRInstanceFleetConfig(),
			"cfn_emr_instance_group_config": emr.ResourceEMRInstanceGroupConfig(),
			"cfn_emr_security_configuration": emr.ResourceEMRSecurityConfiguration(),
			"cfn_emr_step": emr.ResourceEMRStep(),
			"cfn_events_event_bus": events.ResourceEventsEventBus(),
			"cfn_events_event_bus_policy": events.ResourceEventsEventBusPolicy(),
			"cfn_events_rule": events.ResourceEventsRule(),
			"cfn_fsx_file_system": fsx.ResourceFSxFileSystem(),
			"cfn_gamelift_alias": gamelift.ResourceGameLiftAlias(),
			"cfn_gamelift_build": gamelift.ResourceGameLiftBuild(),
			"cfn_gamelift_fleet": gamelift.ResourceGameLiftFleet(),
			"cfn_glue_classifier": glue.ResourceGlueClassifier(),
			"cfn_glue_connection": glue.ResourceGlueConnection(),
			"cfn_glue_crawler": glue.ResourceGlueCrawler(),
			"cfn_glue_data_catalog_encryption_settings": glue.ResourceGlueDataCatalogEncryptionSettings(),
			"cfn_glue_database": glue.ResourceGlueDatabase(),
			"cfn_glue_dev_endpoint": glue.ResourceGlueDevEndpoint(),
			"cfn_glue_job": glue.ResourceGlueJob(),
			"cfn_glue_ml_transform": glue.ResourceGlueMLTransform(),
			"cfn_glue_partition": glue.ResourceGluePartition(),
			"cfn_glue_security_configuration": glue.ResourceGlueSecurityConfiguration(),
			"cfn_glue_table": glue.ResourceGlueTable(),
			"cfn_glue_trigger": glue.ResourceGlueTrigger(),
			"cfn_glue_workflow": glue.ResourceGlueWorkflow(),
			"cfn_greengrass_connector_definition": greengrass.ResourceGreengrassConnectorDefinition(),
			"cfn_greengrass_connector_definition_version": greengrass.ResourceGreengrassConnectorDefinitionVersion(),
			"cfn_greengrass_core_definition": greengrass.ResourceGreengrassCoreDefinition(),
			"cfn_greengrass_core_definition_version": greengrass.ResourceGreengrassCoreDefinitionVersion(),
			"cfn_greengrass_device_definition": greengrass.ResourceGreengrassDeviceDefinition(),
			"cfn_greengrass_device_definition_version": greengrass.ResourceGreengrassDeviceDefinitionVersion(),
			"cfn_greengrass_function_definition": greengrass.ResourceGreengrassFunctionDefinition(),
			"cfn_greengrass_function_definition_version": greengrass.ResourceGreengrassFunctionDefinitionVersion(),
			"cfn_greengrass_group": greengrass.ResourceGreengrassGroup(),
			"cfn_greengrass_group_version": greengrass.ResourceGreengrassGroupVersion(),
			"cfn_greengrass_logger_definition": greengrass.ResourceGreengrassLoggerDefinition(),
			"cfn_greengrass_logger_definition_version": greengrass.ResourceGreengrassLoggerDefinitionVersion(),
			"cfn_greengrass_resource_definition": greengrass.ResourceGreengrassResourceDefinition(),
			"cfn_greengrass_resource_definition_version": greengrass.ResourceGreengrassResourceDefinitionVersion(),
			"cfn_greengrass_subscription_definition": greengrass.ResourceGreengrassSubscriptionDefinition(),
			"cfn_greengrass_subscription_definition_version": greengrass.ResourceGreengrassSubscriptionDefinitionVersion(),
			"cfn_guardduty_detector": guardduty.ResourceGuardDutyDetector(),
			"cfn_guardduty_filter": guardduty.ResourceGuardDutyFilter(),
			"cfn_guardduty_ip_set": guardduty.ResourceGuardDutyIPSet(),
			"cfn_guardduty_master": guardduty.ResourceGuardDutyMaster(),
			"cfn_guardduty_member": guardduty.ResourceGuardDutyMember(),
			"cfn_guardduty_threat_intel_set": guardduty.ResourceGuardDutyThreatIntelSet(),
			"cfn_iam_access_key": iam.ResourceIAMAccessKey(),
			"cfn_iam_group": iam.ResourceIAMGroup(),
			"cfn_iam_instance_profile": iam.ResourceIAMInstanceProfile(),
			"cfn_iam_managed_policy": iam.ResourceIAMManagedPolicy(),
			"cfn_iam_policy": iam.ResourceIAMPolicy(),
			"cfn_iam_role": iam.ResourceIAMRole(),
			"cfn_iam_service_linked_role": iam.ResourceIAMServiceLinkedRole(),
			"cfn_iam_user": iam.ResourceIAMUser(),
			"cfn_iam_user_to_group_addition": iam.ResourceIAMUserToGroupAddition(),
			"cfn_inspector_assessment_target": inspector.ResourceInspectorAssessmentTarget(),
			"cfn_inspector_assessment_template": inspector.ResourceInspectorAssessmentTemplate(),
			"cfn_inspector_resource_group": inspector.ResourceInspectorResourceGroup(),
			"cfn_iot_certificate": iot.ResourceIoTCertificate(),
			"cfn_iot_policy": iot.ResourceIoTPolicy(),
			"cfn_iot_policy_principal_attachment": iot.ResourceIoTPolicyPrincipalAttachment(),
			"cfn_iot_thing": iot.ResourceIoTThing(),
			"cfn_iot_thing_principal_attachment": iot.ResourceIoTThingPrincipalAttachment(),
			"cfn_iot_topic_rule": iot.ResourceIoTTopicRule(),
			"cfn_iot1click_device": iot1click.ResourceIoT1ClickDevice(),
			"cfn_iot1click_placement": iot1click.ResourceIoT1ClickPlacement(),
			"cfn_iot1click_project": iot1click.ResourceIoT1ClickProject(),
			"cfn_iotanalytics_channel": iotanalytics.ResourceIoTAnalyticsChannel(),
			"cfn_iotanalytics_dataset": iotanalytics.ResourceIoTAnalyticsDataset(),
			"cfn_iotanalytics_datastore": iotanalytics.ResourceIoTAnalyticsDatastore(),
			"cfn_iotanalytics_pipeline": iotanalytics.ResourceIoTAnalyticsPipeline(),
			"cfn_iotevents_detector_model": iotevents.ResourceIoTEventsDetectorModel(),
			"cfn_iotevents_input": iotevents.ResourceIoTEventsInput(),
			"cfn_iotthingsgraph_flow_template": iotthingsgraph.ResourceIoTThingsGraphFlowTemplate(),
			"cfn_kinesis_stream": kinesis.ResourceKinesisStream(),
			"cfn_kinesis_stream_consumer": kinesis.ResourceKinesisStreamConsumer(),
			"cfn_kinesisanalytics_application": kinesisanalytics.ResourceKinesisAnalyticsApplication(),
			"cfn_kinesisanalytics_application_output": kinesisanalytics.ResourceKinesisAnalyticsApplicationOutput(),
			"cfn_kinesisanalytics_application_reference_data_source": kinesisanalytics.ResourceKinesisAnalyticsApplicationReferenceDataSource(),
			"cfn_kinesisanalyticsv2_application": kinesisanalyticsv2.ResourceKinesisAnalyticsV2Application(),
			"cfn_kinesisanalyticsv2_application_cloud_watch_logging_option": kinesisanalyticsv2.ResourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption(),
			"cfn_kinesisanalyticsv2_application_output": kinesisanalyticsv2.ResourceKinesisAnalyticsV2ApplicationOutput(),
			"cfn_kinesisanalyticsv2_application_reference_data_source": kinesisanalyticsv2.ResourceKinesisAnalyticsV2ApplicationReferenceDataSource(),
			"cfn_kinesisfirehose_delivery_stream": kinesisfirehose.ResourceKinesisFirehoseDeliveryStream(),
			"cfn_kms_alias": kms.ResourceKMSAlias(),
			"cfn_kms_key": kms.ResourceKMSKey(),
			"cfn_lakeformation_data_lake_settings": lakeformation.ResourceLakeFormationDataLakeSettings(),
			"cfn_lakeformation_permissions": lakeformation.ResourceLakeFormationPermissions(),
			"cfn_lakeformation_resource": lakeformation.ResourceLakeFormationResource(),
			"cfn_lambda_alias": lambda.ResourceLambdaAlias(),
			"cfn_lambda_event_source_mapping": lambda.ResourceLambdaEventSourceMapping(),
			"cfn_lambda_function": lambda.ResourceLambdaFunction(),
			"cfn_lambda_layer_version": lambda.ResourceLambdaLayerVersion(),
			"cfn_lambda_layer_version_permission": lambda.ResourceLambdaLayerVersionPermission(),
			"cfn_lambda_permission": lambda.ResourceLambdaPermission(),
			"cfn_lambda_version": lambda.ResourceLambdaVersion(),
			"cfn_logs_destination": logs.ResourceLogsDestination(),
			"cfn_logs_log_group": logs.ResourceLogsLogGroup(),
			"cfn_logs_log_stream": logs.ResourceLogsLogStream(),
			"cfn_logs_metric_filter": logs.ResourceLogsMetricFilter(),
			"cfn_logs_subscription_filter": logs.ResourceLogsSubscriptionFilter(),
			"cfn_managedblockchain_member": managedblockchain.ResourceManagedBlockchainMember(),
			"cfn_managedblockchain_node": managedblockchain.ResourceManagedBlockchainNode(),
			"cfn_medialive_channel": medialive.ResourceMediaLiveChannel(),
			"cfn_medialive_input": medialive.ResourceMediaLiveInput(),
			"cfn_medialive_input_security_group": medialive.ResourceMediaLiveInputSecurityGroup(),
			"cfn_mediastore_container": mediastore.ResourceMediaStoreContainer(),
			"cfn_msk_cluster": msk.ResourceMSKCluster(),
			"cfn_neptune_db_cluster": neptune.ResourceNeptuneDBCluster(),
			"cfn_neptune_db_cluster_parameter_group": neptune.ResourceNeptuneDBClusterParameterGroup(),
			"cfn_neptune_db_instance": neptune.ResourceNeptuneDBInstance(),
			"cfn_neptune_db_parameter_group": neptune.ResourceNeptuneDBParameterGroup(),
			"cfn_neptune_db_subnet_group": neptune.ResourceNeptuneDBSubnetGroup(),
			"cfn_opsworks_app": opsworks.ResourceOpsWorksApp(),
			"cfn_opsworks_elastic_load_balancer_attachment": opsworks.ResourceOpsWorksElasticLoadBalancerAttachment(),
			"cfn_opsworks_instance": opsworks.ResourceOpsWorksInstance(),
			"cfn_opsworks_layer": opsworks.ResourceOpsWorksLayer(),
			"cfn_opsworks_stack": opsworks.ResourceOpsWorksStack(),
			"cfn_opsworks_user_profile": opsworks.ResourceOpsWorksUserProfile(),
			"cfn_opsworks_volume": opsworks.ResourceOpsWorksVolume(),
			"cfn_opsworkscm_server": opsworkscm.ResourceOpsWorksCMServer(),
			"cfn_pinpoint_adm_channel": pinpoint.ResourcePinpointADMChannel(),
			"cfn_pinpoint_apns_channel": pinpoint.ResourcePinpointAPNSChannel(),
			"cfn_pinpoint_apns_sandbox_channel": pinpoint.ResourcePinpointAPNSSandboxChannel(),
			"cfn_pinpoint_apns_voip_channel": pinpoint.ResourcePinpointAPNSVoipChannel(),
			"cfn_pinpoint_apns_voip_sandbox_channel": pinpoint.ResourcePinpointAPNSVoipSandboxChannel(),
			"cfn_pinpoint_app": pinpoint.ResourcePinpointApp(),
			"cfn_pinpoint_application_settings": pinpoint.ResourcePinpointApplicationSettings(),
			"cfn_pinpoint_baidu_channel": pinpoint.ResourcePinpointBaiduChannel(),
			"cfn_pinpoint_campaign": pinpoint.ResourcePinpointCampaign(),
			"cfn_pinpoint_email_channel": pinpoint.ResourcePinpointEmailChannel(),
			"cfn_pinpoint_email_template": pinpoint.ResourcePinpointEmailTemplate(),
			"cfn_pinpoint_event_stream": pinpoint.ResourcePinpointEventStream(),
			"cfn_pinpoint_gcm_channel": pinpoint.ResourcePinpointGCMChannel(),
			"cfn_pinpoint_push_template": pinpoint.ResourcePinpointPushTemplate(),
			"cfn_pinpoint_sms_channel": pinpoint.ResourcePinpointSMSChannel(),
			"cfn_pinpoint_segment": pinpoint.ResourcePinpointSegment(),
			"cfn_pinpoint_sms_template": pinpoint.ResourcePinpointSmsTemplate(),
			"cfn_pinpoint_voice_channel": pinpoint.ResourcePinpointVoiceChannel(),
			"cfn_pinpointemail_configuration_set": pinpointemail.ResourcePinpointEmailConfigurationSet(),
			"cfn_pinpointemail_configuration_set_event_destination": pinpointemail.ResourcePinpointEmailConfigurationSetEventDestination(),
			"cfn_pinpointemail_dedicated_ip_pool": pinpointemail.ResourcePinpointEmailDedicatedIpPool(),
			"cfn_pinpointemail_identity": pinpointemail.ResourcePinpointEmailIdentity(),
			"cfn_qldb_ledger": qldb.ResourceQLDBLedger(),
			"cfn_ram_resource_share": ram.ResourceRAMResourceShare(),
			"cfn_rds_db_cluster": rds.ResourceRDSDBCluster(),
			"cfn_rds_db_cluster_parameter_group": rds.ResourceRDSDBClusterParameterGroup(),
			"cfn_rds_db_instance": rds.ResourceRDSDBInstance(),
			"cfn_rds_db_parameter_group": rds.ResourceRDSDBParameterGroup(),
			"cfn_rds_db_security_group": rds.ResourceRDSDBSecurityGroup(),
			"cfn_rds_db_security_group_ingress": rds.ResourceRDSDBSecurityGroupIngress(),
			"cfn_rds_db_subnet_group": rds.ResourceRDSDBSubnetGroup(),
			"cfn_rds_event_subscription": rds.ResourceRDSEventSubscription(),
			"cfn_rds_option_group": rds.ResourceRDSOptionGroup(),
			"cfn_redshift_cluster": redshift.ResourceRedshiftCluster(),
			"cfn_redshift_cluster_parameter_group": redshift.ResourceRedshiftClusterParameterGroup(),
			"cfn_redshift_cluster_security_group": redshift.ResourceRedshiftClusterSecurityGroup(),
			"cfn_redshift_cluster_security_group_ingress": redshift.ResourceRedshiftClusterSecurityGroupIngress(),
			"cfn_redshift_cluster_subnet_group": redshift.ResourceRedshiftClusterSubnetGroup(),
			"cfn_robomaker_fleet": robomaker.ResourceRoboMakerFleet(),
			"cfn_robomaker_robot": robomaker.ResourceRoboMakerRobot(),
			"cfn_robomaker_robot_application": robomaker.ResourceRoboMakerRobotApplication(),
			"cfn_robomaker_robot_application_version": robomaker.ResourceRoboMakerRobotApplicationVersion(),
			"cfn_robomaker_simulation_application": robomaker.ResourceRoboMakerSimulationApplication(),
			"cfn_robomaker_simulation_application_version": robomaker.ResourceRoboMakerSimulationApplicationVersion(),
			"cfn_route53_health_check": route53.ResourceRoute53HealthCheck(),
			"cfn_route53_hosted_zone": route53.ResourceRoute53HostedZone(),
			"cfn_route53_record_set": route53.ResourceRoute53RecordSet(),
			"cfn_route53_record_set_group": route53.ResourceRoute53RecordSetGroup(),
			"cfn_route53resolver_resolver_endpoint": route53resolver.ResourceRoute53ResolverResolverEndpoint(),
			"cfn_route53resolver_resolver_rule": route53resolver.ResourceRoute53ResolverResolverRule(),
			"cfn_route53resolver_resolver_rule_association": route53resolver.ResourceRoute53ResolverResolverRuleAssociation(),
			"cfn_s3_bucket": s3.ResourceS3Bucket(),
			"cfn_s3_bucket_policy": s3.ResourceS3BucketPolicy(),
			"cfn_sagemaker_code_repository": sagemaker.ResourceSageMakerCodeRepository(),
			"cfn_sagemaker_endpoint": sagemaker.ResourceSageMakerEndpoint(),
			"cfn_sagemaker_endpoint_config": sagemaker.ResourceSageMakerEndpointConfig(),
			"cfn_sagemaker_model": sagemaker.ResourceSageMakerModel(),
			"cfn_sagemaker_notebook_instance": sagemaker.ResourceSageMakerNotebookInstance(),
			"cfn_sagemaker_notebook_instance_lifecycle_config": sagemaker.ResourceSageMakerNotebookInstanceLifecycleConfig(),
			"cfn_sagemaker_workteam": sagemaker.ResourceSageMakerWorkteam(),
			"cfn_sdb_domain": sdb.ResourceSDBDomain(),
			"cfn_secretsmanager_resource_policy": secretsmanager.ResourceSecretsManagerResourcePolicy(),
			"cfn_secretsmanager_rotation_schedule": secretsmanager.ResourceSecretsManagerRotationSchedule(),
			"cfn_secretsmanager_secret": secretsmanager.ResourceSecretsManagerSecret(),
			"cfn_secretsmanager_secret_target_attachment": secretsmanager.ResourceSecretsManagerSecretTargetAttachment(),
			"cfn_securityhub_hub": securityhub.ResourceSecurityHubHub(),
			"cfn_servicecatalog_accepted_portfolio_share": servicecatalog.ResourceServiceCatalogAcceptedPortfolioShare(),
			"cfn_servicecatalog_cloud_formation_product": servicecatalog.ResourceServiceCatalogCloudFormationProduct(),
			"cfn_servicecatalog_cloud_formation_provisioned_product": servicecatalog.ResourceServiceCatalogCloudFormationProvisionedProduct(),
			"cfn_servicecatalog_launch_notification_constraint": servicecatalog.ResourceServiceCatalogLaunchNotificationConstraint(),
			"cfn_servicecatalog_launch_role_constraint": servicecatalog.ResourceServiceCatalogLaunchRoleConstraint(),
			"cfn_servicecatalog_launch_template_constraint": servicecatalog.ResourceServiceCatalogLaunchTemplateConstraint(),
			"cfn_servicecatalog_portfolio": servicecatalog.ResourceServiceCatalogPortfolio(),
			"cfn_servicecatalog_portfolio_principal_association": servicecatalog.ResourceServiceCatalogPortfolioPrincipalAssociation(),
			"cfn_servicecatalog_portfolio_product_association": servicecatalog.ResourceServiceCatalogPortfolioProductAssociation(),
			"cfn_servicecatalog_portfolio_share": servicecatalog.ResourceServiceCatalogPortfolioShare(),
			"cfn_servicecatalog_resource_update_constraint": servicecatalog.ResourceServiceCatalogResourceUpdateConstraint(),
			"cfn_servicecatalog_stack_set_constraint": servicecatalog.ResourceServiceCatalogStackSetConstraint(),
			"cfn_servicecatalog_tag_option": servicecatalog.ResourceServiceCatalogTagOption(),
			"cfn_servicecatalog_tag_option_association": servicecatalog.ResourceServiceCatalogTagOptionAssociation(),
			"cfn_servicediscovery_http_namespace": servicediscovery.ResourceServiceDiscoveryHttpNamespace(),
			"cfn_servicediscovery_instance": servicediscovery.ResourceServiceDiscoveryInstance(),
			"cfn_servicediscovery_private_dns_namespace": servicediscovery.ResourceServiceDiscoveryPrivateDnsNamespace(),
			"cfn_servicediscovery_public_dns_namespace": servicediscovery.ResourceServiceDiscoveryPublicDnsNamespace(),
			"cfn_servicediscovery_service": servicediscovery.ResourceServiceDiscoveryService(),
			"cfn_ses_configuration_set": ses.ResourceSESConfigurationSet(),
			"cfn_ses_configuration_set_event_destination": ses.ResourceSESConfigurationSetEventDestination(),
			"cfn_ses_receipt_filter": ses.ResourceSESReceiptFilter(),
			"cfn_ses_receipt_rule": ses.ResourceSESReceiptRule(),
			"cfn_ses_receipt_rule_set": ses.ResourceSESReceiptRuleSet(),
			"cfn_ses_template": ses.ResourceSESTemplate(),
			"cfn_sns_subscription": sns.ResourceSNSSubscription(),
			"cfn_sns_topic": sns.ResourceSNSTopic(),
			"cfn_sns_topic_policy": sns.ResourceSNSTopicPolicy(),
			"cfn_sqs_queue": sqs.ResourceSQSQueue(),
			"cfn_sqs_queue_policy": sqs.ResourceSQSQueuePolicy(),
			"cfn_ssm_association": ssm.ResourceSSMAssociation(),
			"cfn_ssm_document": ssm.ResourceSSMDocument(),
			"cfn_ssm_maintenance_window": ssm.ResourceSSMMaintenanceWindow(),
			"cfn_ssm_maintenance_window_target": ssm.ResourceSSMMaintenanceWindowTarget(),
			"cfn_ssm_maintenance_window_task": ssm.ResourceSSMMaintenanceWindowTask(),
			"cfn_ssm_parameter": ssm.ResourceSSMParameter(),
			"cfn_ssm_patch_baseline": ssm.ResourceSSMPatchBaseline(),
			"cfn_ssm_resource_data_sync": ssm.ResourceSSMResourceDataSync(),
			"cfn_stepfunctions_activity": stepfunctions.ResourceStepFunctionsActivity(),
			"cfn_stepfunctions_state_machine": stepfunctions.ResourceStepFunctionsStateMachine(),
			"cfn_transfer_server": transfer.ResourceTransferServer(),
			"cfn_transfer_user": transfer.ResourceTransferUser(),
			"cfn_waf_byte_match_set": waf.ResourceWAFByteMatchSet(),
			"cfn_waf_ip_set": waf.ResourceWAFIPSet(),
			"cfn_waf_rule": waf.ResourceWAFRule(),
			"cfn_waf_size_constraint_set": waf.ResourceWAFSizeConstraintSet(),
			"cfn_waf_sql_injection_match_set": waf.ResourceWAFSqlInjectionMatchSet(),
			"cfn_waf_web_acl": waf.ResourceWAFWebACL(),
			"cfn_waf_xss_match_set": waf.ResourceWAFXssMatchSet(),
			"cfn_wafregional_byte_match_set": wafregional.ResourceWAFRegionalByteMatchSet(),
			"cfn_wafregional_geo_match_set": wafregional.ResourceWAFRegionalGeoMatchSet(),
			"cfn_wafregional_ip_set": wafregional.ResourceWAFRegionalIPSet(),
			"cfn_wafregional_rate_based_rule": wafregional.ResourceWAFRegionalRateBasedRule(),
			"cfn_wafregional_regex_pattern_set": wafregional.ResourceWAFRegionalRegexPatternSet(),
			"cfn_wafregional_rule": wafregional.ResourceWAFRegionalRule(),
			"cfn_wafregional_size_constraint_set": wafregional.ResourceWAFRegionalSizeConstraintSet(),
			"cfn_wafregional_sql_injection_match_set": wafregional.ResourceWAFRegionalSqlInjectionMatchSet(),
			"cfn_wafregional_web_acl": wafregional.ResourceWAFRegionalWebACL(),
			"cfn_wafregional_web_acl_association": wafregional.ResourceWAFRegionalWebACLAssociation(),
			"cfn_wafregional_xss_match_set": wafregional.ResourceWAFRegionalXssMatchSet(),
			"cfn_workspaces_workspace": workspaces.ResourceWorkSpacesWorkspace(),
		},
		DataSourcesMap: map[string]*schema.Resource{
		},
	}
}