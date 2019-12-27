// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   https://github.com/UnquietCode/terraform-provider-cloudformation

package cfn

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/accessanalyzer"
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
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/codestarnotifications"
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
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/eventschemas"
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
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/mediaconvert"
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
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/wafv2"
	"github.com/unquietcode/terraform-cfn-provider/cfn/services/workspaces"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ConfigureFunc: plugin.ProviderConfigure,
		
		Schema: map[string]*schema.Schema{
			"workdir": {
				Type: schema.TypeString,
				Required: true,
				Description: "working directory on the filesystem",
			},
		},
		
		ResourcesMap: map[string]*schema.Resource{
		},
		
		DataSourcesMap: map[string]*schema.Resource{
			"cfn_accessanalyzer_analyzer": accessanalyzer.DatasourceAccessAnalyzerAnalyzer(),
			"cfn_amazonmq_broker": amazonmq.DatasourceAmazonMQBroker(),
			"cfn_amazonmq_configuration": amazonmq.DatasourceAmazonMQConfiguration(),
			"cfn_amazonmq_configuration_association": amazonmq.DatasourceAmazonMQConfigurationAssociation(),
			"cfn_amplify_app": amplify.DatasourceAmplifyApp(),
			"cfn_amplify_branch": amplify.DatasourceAmplifyBranch(),
			"cfn_amplify_domain": amplify.DatasourceAmplifyDomain(),
			"cfn_apigateway_account": apigateway.DatasourceApiGatewayAccount(),
			"cfn_apigateway_api_key": apigateway.DatasourceApiGatewayApiKey(),
			"cfn_apigateway_authorizer": apigateway.DatasourceApiGatewayAuthorizer(),
			"cfn_apigateway_base_path_mapping": apigateway.DatasourceApiGatewayBasePathMapping(),
			"cfn_apigateway_client_certificate": apigateway.DatasourceApiGatewayClientCertificate(),
			"cfn_apigateway_deployment": apigateway.DatasourceApiGatewayDeployment(),
			"cfn_apigateway_documentation_part": apigateway.DatasourceApiGatewayDocumentationPart(),
			"cfn_apigateway_documentation_version": apigateway.DatasourceApiGatewayDocumentationVersion(),
			"cfn_apigateway_domain_name": apigateway.DatasourceApiGatewayDomainName(),
			"cfn_apigateway_gateway_response": apigateway.DatasourceApiGatewayGatewayResponse(),
			"cfn_apigateway_method": apigateway.DatasourceApiGatewayMethod(),
			"cfn_apigateway_model": apigateway.DatasourceApiGatewayModel(),
			"cfn_apigateway_request_validator": apigateway.DatasourceApiGatewayRequestValidator(),
			"cfn_apigateway_resource": apigateway.DatasourceApiGatewayResource(),
			"cfn_apigateway_rest_api": apigateway.DatasourceApiGatewayRestApi(),
			"cfn_apigateway_stage": apigateway.DatasourceApiGatewayStage(),
			"cfn_apigateway_usage_plan": apigateway.DatasourceApiGatewayUsagePlan(),
			"cfn_apigateway_usage_plan_key": apigateway.DatasourceApiGatewayUsagePlanKey(),
			"cfn_apigateway_vpc_link": apigateway.DatasourceApiGatewayVpcLink(),
			"cfn_apigatewayv2_api": apigatewayv2.DatasourceApiGatewayV2Api(),
			"cfn_apigatewayv2_api_mapping": apigatewayv2.DatasourceApiGatewayV2ApiMapping(),
			"cfn_apigatewayv2_authorizer": apigatewayv2.DatasourceApiGatewayV2Authorizer(),
			"cfn_apigatewayv2_deployment": apigatewayv2.DatasourceApiGatewayV2Deployment(),
			"cfn_apigatewayv2_domain_name": apigatewayv2.DatasourceApiGatewayV2DomainName(),
			"cfn_apigatewayv2_integration": apigatewayv2.DatasourceApiGatewayV2Integration(),
			"cfn_apigatewayv2_integration_response": apigatewayv2.DatasourceApiGatewayV2IntegrationResponse(),
			"cfn_apigatewayv2_model": apigatewayv2.DatasourceApiGatewayV2Model(),
			"cfn_apigatewayv2_route": apigatewayv2.DatasourceApiGatewayV2Route(),
			"cfn_apigatewayv2_route_response": apigatewayv2.DatasourceApiGatewayV2RouteResponse(),
			"cfn_apigatewayv2_stage": apigatewayv2.DatasourceApiGatewayV2Stage(),
			"cfn_applicationautoscaling_scalable_target": applicationautoscaling.DatasourceApplicationAutoScalingScalableTarget(),
			"cfn_applicationautoscaling_scaling_policy": applicationautoscaling.DatasourceApplicationAutoScalingScalingPolicy(),
			"cfn_appmesh_mesh": appmesh.DatasourceAppMeshMesh(),
			"cfn_appmesh_route": appmesh.DatasourceAppMeshRoute(),
			"cfn_appmesh_virtual_node": appmesh.DatasourceAppMeshVirtualNode(),
			"cfn_appmesh_virtual_router": appmesh.DatasourceAppMeshVirtualRouter(),
			"cfn_appmesh_virtual_service": appmesh.DatasourceAppMeshVirtualService(),
			"cfn_appstream_directory_config": appstream.DatasourceAppStreamDirectoryConfig(),
			"cfn_appstream_fleet": appstream.DatasourceAppStreamFleet(),
			"cfn_appstream_image_builder": appstream.DatasourceAppStreamImageBuilder(),
			"cfn_appstream_stack": appstream.DatasourceAppStreamStack(),
			"cfn_appstream_stack_fleet_association": appstream.DatasourceAppStreamStackFleetAssociation(),
			"cfn_appstream_stack_user_association": appstream.DatasourceAppStreamStackUserAssociation(),
			"cfn_appstream_user": appstream.DatasourceAppStreamUser(),
			"cfn_appsync_api_cache": appsync.DatasourceAppSyncApiCache(),
			"cfn_appsync_api_key": appsync.DatasourceAppSyncApiKey(),
			"cfn_appsync_data_source": appsync.DatasourceAppSyncDataSource(),
			"cfn_appsync_function_configuration": appsync.DatasourceAppSyncFunctionConfiguration(),
			"cfn_appsync_graph_ql_api": appsync.DatasourceAppSyncGraphQLApi(),
			"cfn_appsync_graph_ql_schema": appsync.DatasourceAppSyncGraphQLSchema(),
			"cfn_appsync_resolver": appsync.DatasourceAppSyncResolver(),
			"cfn_athena_named_query": athena.DatasourceAthenaNamedQuery(),
			"cfn_autoscaling_auto_scaling_group": autoscaling.DatasourceAutoScalingAutoScalingGroup(),
			"cfn_autoscaling_launch_configuration": autoscaling.DatasourceAutoScalingLaunchConfiguration(),
			"cfn_autoscaling_lifecycle_hook": autoscaling.DatasourceAutoScalingLifecycleHook(),
			"cfn_autoscaling_scaling_policy": autoscaling.DatasourceAutoScalingScalingPolicy(),
			"cfn_autoscaling_scheduled_action": autoscaling.DatasourceAutoScalingScheduledAction(),
			"cfn_autoscalingplans_scaling_plan": autoscalingplans.DatasourceAutoScalingPlansScalingPlan(),
			"cfn_backup_backup_plan": backup.DatasourceBackupBackupPlan(),
			"cfn_backup_backup_selection": backup.DatasourceBackupBackupSelection(),
			"cfn_backup_backup_vault": backup.DatasourceBackupBackupVault(),
			"cfn_batch_compute_environment": batch.DatasourceBatchComputeEnvironment(),
			"cfn_batch_job_definition": batch.DatasourceBatchJobDefinition(),
			"cfn_batch_job_queue": batch.DatasourceBatchJobQueue(),
			"cfn_budgets_budget": budgets.DatasourceBudgetsBudget(),
			"cfn_certificatemanager_certificate": certificatemanager.DatasourceCertificateManagerCertificate(),
			"cfn_cloud9_environment_ec2": cloud9.DatasourceCloud9EnvironmentEC2(),
			"cfn_cloudformation_custom_resource": cloudformation.DatasourceCloudFormationCustomResource(),
			"cfn_cloudformation_macro": cloudformation.DatasourceCloudFormationMacro(),
			"cfn_cloudformation_stack": cloudformation.DatasourceCloudFormationStack(),
			"cfn_cloudformation_wait_condition": cloudformation.DatasourceCloudFormationWaitCondition(),
			"cfn_cloudformation_wait_condition_handle": cloudformation.DatasourceCloudFormationWaitConditionHandle(),
			"cfn_cloudfront_cloud_front_origin_access_identity": cloudfront.DatasourceCloudFrontCloudFrontOriginAccessIdentity(),
			"cfn_cloudfront_distribution": cloudfront.DatasourceCloudFrontDistribution(),
			"cfn_cloudfront_streaming_distribution": cloudfront.DatasourceCloudFrontStreamingDistribution(),
			"cfn_cloudtrail_trail": cloudtrail.DatasourceCloudTrailTrail(),
			"cfn_cloudwatch_alarm": cloudwatch.DatasourceCloudWatchAlarm(),
			"cfn_cloudwatch_anomaly_detector": cloudwatch.DatasourceCloudWatchAnomalyDetector(),
			"cfn_cloudwatch_dashboard": cloudwatch.DatasourceCloudWatchDashboard(),
			"cfn_cloudwatch_insight_rule": cloudwatch.DatasourceCloudWatchInsightRule(),
			"cfn_codebuild_project": codebuild.DatasourceCodeBuildProject(),
			"cfn_codebuild_source_credential": codebuild.DatasourceCodeBuildSourceCredential(),
			"cfn_codecommit_repository": codecommit.DatasourceCodeCommitRepository(),
			"cfn_codedeploy_application": codedeploy.DatasourceCodeDeployApplication(),
			"cfn_codedeploy_deployment_config": codedeploy.DatasourceCodeDeployDeploymentConfig(),
			"cfn_codedeploy_deployment_group": codedeploy.DatasourceCodeDeployDeploymentGroup(),
			"cfn_codepipeline_custom_action_type": codepipeline.DatasourceCodePipelineCustomActionType(),
			"cfn_codepipeline_pipeline": codepipeline.DatasourceCodePipelinePipeline(),
			"cfn_codepipeline_webhook": codepipeline.DatasourceCodePipelineWebhook(),
			"cfn_codestar_git_hub_repository": codestar.DatasourceCodeStarGitHubRepository(),
			"cfn_codestarnotifications_notification_rule": codestarnotifications.DatasourceCodeStarNotificationsNotificationRule(),
			"cfn_cognito_identity_pool": cognito.DatasourceCognitoIdentityPool(),
			"cfn_cognito_identity_pool_role_attachment": cognito.DatasourceCognitoIdentityPoolRoleAttachment(),
			"cfn_cognito_user_pool": cognito.DatasourceCognitoUserPool(),
			"cfn_cognito_user_pool_client": cognito.DatasourceCognitoUserPoolClient(),
			"cfn_cognito_user_pool_domain": cognito.DatasourceCognitoUserPoolDomain(),
			"cfn_cognito_user_pool_group": cognito.DatasourceCognitoUserPoolGroup(),
			"cfn_cognito_user_pool_identity_provider": cognito.DatasourceCognitoUserPoolIdentityProvider(),
			"cfn_cognito_user_pool_resource_server": cognito.DatasourceCognitoUserPoolResourceServer(),
			"cfn_cognito_user_pool_risk_configuration_attachment": cognito.DatasourceCognitoUserPoolRiskConfigurationAttachment(),
			"cfn_cognito_user_pool_ui_customization_attachment": cognito.DatasourceCognitoUserPoolUICustomizationAttachment(),
			"cfn_cognito_user_pool_user": cognito.DatasourceCognitoUserPoolUser(),
			"cfn_cognito_user_pool_user_to_group_attachment": cognito.DatasourceCognitoUserPoolUserToGroupAttachment(),
			"cfn_config_aggregation_authorization": config.DatasourceConfigAggregationAuthorization(),
			"cfn_config_config_rule": config.DatasourceConfigConfigRule(),
			"cfn_config_configuration_aggregator": config.DatasourceConfigConfigurationAggregator(),
			"cfn_config_configuration_recorder": config.DatasourceConfigConfigurationRecorder(),
			"cfn_config_delivery_channel": config.DatasourceConfigDeliveryChannel(),
			"cfn_config_organization_config_rule": config.DatasourceConfigOrganizationConfigRule(),
			"cfn_config_remediation_configuration": config.DatasourceConfigRemediationConfiguration(),
			"cfn_datapipeline_pipeline": datapipeline.DatasourceDataPipelinePipeline(),
			"cfn_dax_cluster": dax.DatasourceDAXCluster(),
			"cfn_dax_parameter_group": dax.DatasourceDAXParameterGroup(),
			"cfn_dax_subnet_group": dax.DatasourceDAXSubnetGroup(),
			"cfn_directoryservice_microsoft_ad": directoryservice.DatasourceDirectoryServiceMicrosoftAD(),
			"cfn_directoryservice_simple_ad": directoryservice.DatasourceDirectoryServiceSimpleAD(),
			"cfn_dlm_lifecycle_policy": dlm.DatasourceDLMLifecyclePolicy(),
			"cfn_dms_certificate": dms.DatasourceDMSCertificate(),
			"cfn_dms_endpoint": dms.DatasourceDMSEndpoint(),
			"cfn_dms_event_subscription": dms.DatasourceDMSEventSubscription(),
			"cfn_dms_replication_instance": dms.DatasourceDMSReplicationInstance(),
			"cfn_dms_replication_subnet_group": dms.DatasourceDMSReplicationSubnetGroup(),
			"cfn_dms_replication_task": dms.DatasourceDMSReplicationTask(),
			"cfn_docdb_db_cluster": docdb.DatasourceDocDBDBCluster(),
			"cfn_docdb_db_cluster_parameter_group": docdb.DatasourceDocDBDBClusterParameterGroup(),
			"cfn_docdb_db_instance": docdb.DatasourceDocDBDBInstance(),
			"cfn_docdb_db_subnet_group": docdb.DatasourceDocDBDBSubnetGroup(),
			"cfn_dynamodb_table": dynamodb.DatasourceDynamoDBTable(),
			"cfn_ec2_capacity_reservation": ec2.DatasourceEC2CapacityReservation(),
			"cfn_ec2_client_vpn_authorization_rule": ec2.DatasourceEC2ClientVpnAuthorizationRule(),
			"cfn_ec2_client_vpn_endpoint": ec2.DatasourceEC2ClientVpnEndpoint(),
			"cfn_ec2_client_vpn_route": ec2.DatasourceEC2ClientVpnRoute(),
			"cfn_ec2_client_vpn_target_network_association": ec2.DatasourceEC2ClientVpnTargetNetworkAssociation(),
			"cfn_ec2_customer_gateway": ec2.DatasourceEC2CustomerGateway(),
			"cfn_ec2_dhcp_options": ec2.DatasourceEC2DHCPOptions(),
			"cfn_ec2_ec2_fleet": ec2.DatasourceEC2EC2Fleet(),
			"cfn_ec2_eip": ec2.DatasourceEC2EIP(),
			"cfn_ec2_eip_association": ec2.DatasourceEC2EIPAssociation(),
			"cfn_ec2_egress_only_internet_gateway": ec2.DatasourceEC2EgressOnlyInternetGateway(),
			"cfn_ec2_flow_log": ec2.DatasourceEC2FlowLog(),
			"cfn_ec2_host": ec2.DatasourceEC2Host(),
			"cfn_ec2_instance": ec2.DatasourceEC2Instance(),
			"cfn_ec2_internet_gateway": ec2.DatasourceEC2InternetGateway(),
			"cfn_ec2_launch_template": ec2.DatasourceEC2LaunchTemplate(),
			"cfn_ec2_nat_gateway": ec2.DatasourceEC2NatGateway(),
			"cfn_ec2_network_acl": ec2.DatasourceEC2NetworkAcl(),
			"cfn_ec2_network_acl_entry": ec2.DatasourceEC2NetworkAclEntry(),
			"cfn_ec2_network_interface": ec2.DatasourceEC2NetworkInterface(),
			"cfn_ec2_network_interface_attachment": ec2.DatasourceEC2NetworkInterfaceAttachment(),
			"cfn_ec2_network_interface_permission": ec2.DatasourceEC2NetworkInterfacePermission(),
			"cfn_ec2_placement_group": ec2.DatasourceEC2PlacementGroup(),
			"cfn_ec2_route": ec2.DatasourceEC2Route(),
			"cfn_ec2_route_table": ec2.DatasourceEC2RouteTable(),
			"cfn_ec2_security_group": ec2.DatasourceEC2SecurityGroup(),
			"cfn_ec2_security_group_egress": ec2.DatasourceEC2SecurityGroupEgress(),
			"cfn_ec2_security_group_ingress": ec2.DatasourceEC2SecurityGroupIngress(),
			"cfn_ec2_spot_fleet": ec2.DatasourceEC2SpotFleet(),
			"cfn_ec2_subnet": ec2.DatasourceEC2Subnet(),
			"cfn_ec2_subnet_cidr_block": ec2.DatasourceEC2SubnetCidrBlock(),
			"cfn_ec2_subnet_network_acl_association": ec2.DatasourceEC2SubnetNetworkAclAssociation(),
			"cfn_ec2_subnet_route_table_association": ec2.DatasourceEC2SubnetRouteTableAssociation(),
			"cfn_ec2_traffic_mirror_filter": ec2.DatasourceEC2TrafficMirrorFilter(),
			"cfn_ec2_traffic_mirror_filter_rule": ec2.DatasourceEC2TrafficMirrorFilterRule(),
			"cfn_ec2_traffic_mirror_session": ec2.DatasourceEC2TrafficMirrorSession(),
			"cfn_ec2_traffic_mirror_target": ec2.DatasourceEC2TrafficMirrorTarget(),
			"cfn_ec2_transit_gateway": ec2.DatasourceEC2TransitGateway(),
			"cfn_ec2_transit_gateway_attachment": ec2.DatasourceEC2TransitGatewayAttachment(),
			"cfn_ec2_transit_gateway_route": ec2.DatasourceEC2TransitGatewayRoute(),
			"cfn_ec2_transit_gateway_route_table": ec2.DatasourceEC2TransitGatewayRouteTable(),
			"cfn_ec2_transit_gateway_route_table_association": ec2.DatasourceEC2TransitGatewayRouteTableAssociation(),
			"cfn_ec2_transit_gateway_route_table_propagation": ec2.DatasourceEC2TransitGatewayRouteTablePropagation(),
			"cfn_ec2_vpc": ec2.DatasourceEC2VPC(),
			"cfn_ec2_vpc_cidr_block": ec2.DatasourceEC2VPCCidrBlock(),
			"cfn_ec2_vpcdhcp_options_association": ec2.DatasourceEC2VPCDHCPOptionsAssociation(),
			"cfn_ec2_vpc_endpoint": ec2.DatasourceEC2VPCEndpoint(),
			"cfn_ec2_vpc_endpoint_connection_notification": ec2.DatasourceEC2VPCEndpointConnectionNotification(),
			"cfn_ec2_vpc_endpoint_service": ec2.DatasourceEC2VPCEndpointService(),
			"cfn_ec2_vpc_endpoint_service_permissions": ec2.DatasourceEC2VPCEndpointServicePermissions(),
			"cfn_ec2_vpc_gateway_attachment": ec2.DatasourceEC2VPCGatewayAttachment(),
			"cfn_ec2_vpc_peering_connection": ec2.DatasourceEC2VPCPeeringConnection(),
			"cfn_ec2_vpn_connection": ec2.DatasourceEC2VPNConnection(),
			"cfn_ec2_vpn_connection_route": ec2.DatasourceEC2VPNConnectionRoute(),
			"cfn_ec2_vpn_gateway": ec2.DatasourceEC2VPNGateway(),
			"cfn_ec2_vpn_gateway_route_propagation": ec2.DatasourceEC2VPNGatewayRoutePropagation(),
			"cfn_ec2_volume": ec2.DatasourceEC2Volume(),
			"cfn_ec2_volume_attachment": ec2.DatasourceEC2VolumeAttachment(),
			"cfn_ecr_repository": ecr.DatasourceECRRepository(),
			"cfn_ecs_cluster": ecs.DatasourceECSCluster(),
			"cfn_ecs_primary_task_set": ecs.DatasourceECSPrimaryTaskSet(),
			"cfn_ecs_service": ecs.DatasourceECSService(),
			"cfn_ecs_task_definition": ecs.DatasourceECSTaskDefinition(),
			"cfn_ecs_task_set": ecs.DatasourceECSTaskSet(),
			"cfn_efs_file_system": efs.DatasourceEFSFileSystem(),
			"cfn_efs_mount_target": efs.DatasourceEFSMountTarget(),
			"cfn_eks_cluster": eks.DatasourceEKSCluster(),
			"cfn_eks_nodegroup": eks.DatasourceEKSNodegroup(),
			"cfn_elasticache_cache_cluster": elasticache.DatasourceElastiCacheCacheCluster(),
			"cfn_elasticache_parameter_group": elasticache.DatasourceElastiCacheParameterGroup(),
			"cfn_elasticache_replication_group": elasticache.DatasourceElastiCacheReplicationGroup(),
			"cfn_elasticache_security_group": elasticache.DatasourceElastiCacheSecurityGroup(),
			"cfn_elasticache_security_group_ingress": elasticache.DatasourceElastiCacheSecurityGroupIngress(),
			"cfn_elasticache_subnet_group": elasticache.DatasourceElastiCacheSubnetGroup(),
			"cfn_elasticbeanstalk_application": elasticbeanstalk.DatasourceElasticBeanstalkApplication(),
			"cfn_elasticbeanstalk_application_version": elasticbeanstalk.DatasourceElasticBeanstalkApplicationVersion(),
			"cfn_elasticbeanstalk_configuration_template": elasticbeanstalk.DatasourceElasticBeanstalkConfigurationTemplate(),
			"cfn_elasticbeanstalk_environment": elasticbeanstalk.DatasourceElasticBeanstalkEnvironment(),
			"cfn_elasticloadbalancing_load_balancer": elasticloadbalancing.DatasourceElasticLoadBalancingLoadBalancer(),
			"cfn_elasticloadbalancingv2_listener": elasticloadbalancingv2.DatasourceElasticLoadBalancingV2Listener(),
			"cfn_elasticloadbalancingv2_listener_certificate": elasticloadbalancingv2.DatasourceElasticLoadBalancingV2ListenerCertificate(),
			"cfn_elasticloadbalancingv2_listener_rule": elasticloadbalancingv2.DatasourceElasticLoadBalancingV2ListenerRule(),
			"cfn_elasticloadbalancingv2_load_balancer": elasticloadbalancingv2.DatasourceElasticLoadBalancingV2LoadBalancer(),
			"cfn_elasticloadbalancingv2_target_group": elasticloadbalancingv2.DatasourceElasticLoadBalancingV2TargetGroup(),
			"cfn_elasticsearch_domain": elasticsearch.DatasourceElasticsearchDomain(),
			"cfn_emr_cluster": emr.DatasourceEMRCluster(),
			"cfn_emr_instance_fleet_config": emr.DatasourceEMRInstanceFleetConfig(),
			"cfn_emr_instance_group_config": emr.DatasourceEMRInstanceGroupConfig(),
			"cfn_emr_security_configuration": emr.DatasourceEMRSecurityConfiguration(),
			"cfn_emr_step": emr.DatasourceEMRStep(),
			"cfn_events_event_bus": events.DatasourceEventsEventBus(),
			"cfn_events_event_bus_policy": events.DatasourceEventsEventBusPolicy(),
			"cfn_events_rule": events.DatasourceEventsRule(),
			"cfn_eventschemas_discoverer": eventschemas.DatasourceEventSchemasDiscoverer(),
			"cfn_eventschemas_registry": eventschemas.DatasourceEventSchemasRegistry(),
			"cfn_eventschemas_schema": eventschemas.DatasourceEventSchemasSchema(),
			"cfn_fsx_file_system": fsx.DatasourceFSxFileSystem(),
			"cfn_gamelift_alias": gamelift.DatasourceGameLiftAlias(),
			"cfn_gamelift_build": gamelift.DatasourceGameLiftBuild(),
			"cfn_gamelift_fleet": gamelift.DatasourceGameLiftFleet(),
			"cfn_gamelift_game_session_queue": gamelift.DatasourceGameLiftGameSessionQueue(),
			"cfn_gamelift_matchmaking_configuration": gamelift.DatasourceGameLiftMatchmakingConfiguration(),
			"cfn_gamelift_matchmaking_rule_set": gamelift.DatasourceGameLiftMatchmakingRuleSet(),
			"cfn_gamelift_script": gamelift.DatasourceGameLiftScript(),
			"cfn_glue_classifier": glue.DatasourceGlueClassifier(),
			"cfn_glue_connection": glue.DatasourceGlueConnection(),
			"cfn_glue_crawler": glue.DatasourceGlueCrawler(),
			"cfn_glue_data_catalog_encryption_settings": glue.DatasourceGlueDataCatalogEncryptionSettings(),
			"cfn_glue_database": glue.DatasourceGlueDatabase(),
			"cfn_glue_dev_endpoint": glue.DatasourceGlueDevEndpoint(),
			"cfn_glue_job": glue.DatasourceGlueJob(),
			"cfn_glue_ml_transform": glue.DatasourceGlueMLTransform(),
			"cfn_glue_partition": glue.DatasourceGluePartition(),
			"cfn_glue_security_configuration": glue.DatasourceGlueSecurityConfiguration(),
			"cfn_glue_table": glue.DatasourceGlueTable(),
			"cfn_glue_trigger": glue.DatasourceGlueTrigger(),
			"cfn_glue_workflow": glue.DatasourceGlueWorkflow(),
			"cfn_greengrass_connector_definition": greengrass.DatasourceGreengrassConnectorDefinition(),
			"cfn_greengrass_connector_definition_version": greengrass.DatasourceGreengrassConnectorDefinitionVersion(),
			"cfn_greengrass_core_definition": greengrass.DatasourceGreengrassCoreDefinition(),
			"cfn_greengrass_core_definition_version": greengrass.DatasourceGreengrassCoreDefinitionVersion(),
			"cfn_greengrass_device_definition": greengrass.DatasourceGreengrassDeviceDefinition(),
			"cfn_greengrass_device_definition_version": greengrass.DatasourceGreengrassDeviceDefinitionVersion(),
			"cfn_greengrass_function_definition": greengrass.DatasourceGreengrassFunctionDefinition(),
			"cfn_greengrass_function_definition_version": greengrass.DatasourceGreengrassFunctionDefinitionVersion(),
			"cfn_greengrass_group": greengrass.DatasourceGreengrassGroup(),
			"cfn_greengrass_group_version": greengrass.DatasourceGreengrassGroupVersion(),
			"cfn_greengrass_logger_definition": greengrass.DatasourceGreengrassLoggerDefinition(),
			"cfn_greengrass_logger_definition_version": greengrass.DatasourceGreengrassLoggerDefinitionVersion(),
			"cfn_greengrass_resource_definition": greengrass.DatasourceGreengrassResourceDefinition(),
			"cfn_greengrass_resource_definition_version": greengrass.DatasourceGreengrassResourceDefinitionVersion(),
			"cfn_greengrass_subscription_definition": greengrass.DatasourceGreengrassSubscriptionDefinition(),
			"cfn_greengrass_subscription_definition_version": greengrass.DatasourceGreengrassSubscriptionDefinitionVersion(),
			"cfn_guardduty_detector": guardduty.DatasourceGuardDutyDetector(),
			"cfn_guardduty_filter": guardduty.DatasourceGuardDutyFilter(),
			"cfn_guardduty_ip_set": guardduty.DatasourceGuardDutyIPSet(),
			"cfn_guardduty_master": guardduty.DatasourceGuardDutyMaster(),
			"cfn_guardduty_member": guardduty.DatasourceGuardDutyMember(),
			"cfn_guardduty_threat_intel_set": guardduty.DatasourceGuardDutyThreatIntelSet(),
			"cfn_iam_access_key": iam.DatasourceIAMAccessKey(),
			"cfn_iam_group": iam.DatasourceIAMGroup(),
			"cfn_iam_instance_profile": iam.DatasourceIAMInstanceProfile(),
			"cfn_iam_managed_policy": iam.DatasourceIAMManagedPolicy(),
			"cfn_iam_policy": iam.DatasourceIAMPolicy(),
			"cfn_iam_role": iam.DatasourceIAMRole(),
			"cfn_iam_service_linked_role": iam.DatasourceIAMServiceLinkedRole(),
			"cfn_iam_user": iam.DatasourceIAMUser(),
			"cfn_iam_user_to_group_addition": iam.DatasourceIAMUserToGroupAddition(),
			"cfn_inspector_assessment_target": inspector.DatasourceInspectorAssessmentTarget(),
			"cfn_inspector_assessment_template": inspector.DatasourceInspectorAssessmentTemplate(),
			"cfn_inspector_resource_group": inspector.DatasourceInspectorResourceGroup(),
			"cfn_iot_certificate": iot.DatasourceIoTCertificate(),
			"cfn_iot_policy": iot.DatasourceIoTPolicy(),
			"cfn_iot_policy_principal_attachment": iot.DatasourceIoTPolicyPrincipalAttachment(),
			"cfn_iot_thing": iot.DatasourceIoTThing(),
			"cfn_iot_thing_principal_attachment": iot.DatasourceIoTThingPrincipalAttachment(),
			"cfn_iot_topic_rule": iot.DatasourceIoTTopicRule(),
			"cfn_iot1click_device": iot1click.DatasourceIoT1ClickDevice(),
			"cfn_iot1click_placement": iot1click.DatasourceIoT1ClickPlacement(),
			"cfn_iot1click_project": iot1click.DatasourceIoT1ClickProject(),
			"cfn_iotanalytics_channel": iotanalytics.DatasourceIoTAnalyticsChannel(),
			"cfn_iotanalytics_dataset": iotanalytics.DatasourceIoTAnalyticsDataset(),
			"cfn_iotanalytics_datastore": iotanalytics.DatasourceIoTAnalyticsDatastore(),
			"cfn_iotanalytics_pipeline": iotanalytics.DatasourceIoTAnalyticsPipeline(),
			"cfn_iotevents_detector_model": iotevents.DatasourceIoTEventsDetectorModel(),
			"cfn_iotevents_input": iotevents.DatasourceIoTEventsInput(),
			"cfn_iotthingsgraph_flow_template": iotthingsgraph.DatasourceIoTThingsGraphFlowTemplate(),
			"cfn_kinesis_stream": kinesis.DatasourceKinesisStream(),
			"cfn_kinesis_stream_consumer": kinesis.DatasourceKinesisStreamConsumer(),
			"cfn_kinesisanalytics_application": kinesisanalytics.DatasourceKinesisAnalyticsApplication(),
			"cfn_kinesisanalytics_application_output": kinesisanalytics.DatasourceKinesisAnalyticsApplicationOutput(),
			"cfn_kinesisanalytics_application_reference_data_source": kinesisanalytics.DatasourceKinesisAnalyticsApplicationReferenceDataSource(),
			"cfn_kinesisanalyticsv2_application": kinesisanalyticsv2.DatasourceKinesisAnalyticsV2Application(),
			"cfn_kinesisanalyticsv2_application_cloud_watch_logging_option": kinesisanalyticsv2.DatasourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption(),
			"cfn_kinesisanalyticsv2_application_output": kinesisanalyticsv2.DatasourceKinesisAnalyticsV2ApplicationOutput(),
			"cfn_kinesisanalyticsv2_application_reference_data_source": kinesisanalyticsv2.DatasourceKinesisAnalyticsV2ApplicationReferenceDataSource(),
			"cfn_kinesisfirehose_delivery_stream": kinesisfirehose.DatasourceKinesisFirehoseDeliveryStream(),
			"cfn_kms_alias": kms.DatasourceKMSAlias(),
			"cfn_kms_key": kms.DatasourceKMSKey(),
			"cfn_lakeformation_data_lake_settings": lakeformation.DatasourceLakeFormationDataLakeSettings(),
			"cfn_lakeformation_permissions": lakeformation.DatasourceLakeFormationPermissions(),
			"cfn_lakeformation_resource": lakeformation.DatasourceLakeFormationResource(),
			"cfn_lambda_alias": lambda.DatasourceLambdaAlias(),
			"cfn_lambda_event_invoke_config": lambda.DatasourceLambdaEventInvokeConfig(),
			"cfn_lambda_event_source_mapping": lambda.DatasourceLambdaEventSourceMapping(),
			"cfn_lambda_function": lambda.DatasourceLambdaFunction(),
			"cfn_lambda_layer_version": lambda.DatasourceLambdaLayerVersion(),
			"cfn_lambda_layer_version_permission": lambda.DatasourceLambdaLayerVersionPermission(),
			"cfn_lambda_permission": lambda.DatasourceLambdaPermission(),
			"cfn_lambda_version": lambda.DatasourceLambdaVersion(),
			"cfn_logs_destination": logs.DatasourceLogsDestination(),
			"cfn_logs_log_group": logs.DatasourceLogsLogGroup(),
			"cfn_logs_log_stream": logs.DatasourceLogsLogStream(),
			"cfn_logs_metric_filter": logs.DatasourceLogsMetricFilter(),
			"cfn_logs_subscription_filter": logs.DatasourceLogsSubscriptionFilter(),
			"cfn_managedblockchain_member": managedblockchain.DatasourceManagedBlockchainMember(),
			"cfn_managedblockchain_node": managedblockchain.DatasourceManagedBlockchainNode(),
			"cfn_mediaconvert_job_template": mediaconvert.DatasourceMediaConvertJobTemplate(),
			"cfn_mediaconvert_preset": mediaconvert.DatasourceMediaConvertPreset(),
			"cfn_mediaconvert_queue": mediaconvert.DatasourceMediaConvertQueue(),
			"cfn_medialive_channel": medialive.DatasourceMediaLiveChannel(),
			"cfn_medialive_input": medialive.DatasourceMediaLiveInput(),
			"cfn_medialive_input_security_group": medialive.DatasourceMediaLiveInputSecurityGroup(),
			"cfn_mediastore_container": mediastore.DatasourceMediaStoreContainer(),
			"cfn_msk_cluster": msk.DatasourceMSKCluster(),
			"cfn_neptune_db_cluster": neptune.DatasourceNeptuneDBCluster(),
			"cfn_neptune_db_cluster_parameter_group": neptune.DatasourceNeptuneDBClusterParameterGroup(),
			"cfn_neptune_db_instance": neptune.DatasourceNeptuneDBInstance(),
			"cfn_neptune_db_parameter_group": neptune.DatasourceNeptuneDBParameterGroup(),
			"cfn_neptune_db_subnet_group": neptune.DatasourceNeptuneDBSubnetGroup(),
			"cfn_opsworks_app": opsworks.DatasourceOpsWorksApp(),
			"cfn_opsworks_elastic_load_balancer_attachment": opsworks.DatasourceOpsWorksElasticLoadBalancerAttachment(),
			"cfn_opsworks_instance": opsworks.DatasourceOpsWorksInstance(),
			"cfn_opsworks_layer": opsworks.DatasourceOpsWorksLayer(),
			"cfn_opsworks_stack": opsworks.DatasourceOpsWorksStack(),
			"cfn_opsworks_user_profile": opsworks.DatasourceOpsWorksUserProfile(),
			"cfn_opsworks_volume": opsworks.DatasourceOpsWorksVolume(),
			"cfn_opsworkscm_server": opsworkscm.DatasourceOpsWorksCMServer(),
			"cfn_pinpoint_adm_channel": pinpoint.DatasourcePinpointADMChannel(),
			"cfn_pinpoint_apns_channel": pinpoint.DatasourcePinpointAPNSChannel(),
			"cfn_pinpoint_apns_sandbox_channel": pinpoint.DatasourcePinpointAPNSSandboxChannel(),
			"cfn_pinpoint_apns_voip_channel": pinpoint.DatasourcePinpointAPNSVoipChannel(),
			"cfn_pinpoint_apns_voip_sandbox_channel": pinpoint.DatasourcePinpointAPNSVoipSandboxChannel(),
			"cfn_pinpoint_app": pinpoint.DatasourcePinpointApp(),
			"cfn_pinpoint_application_settings": pinpoint.DatasourcePinpointApplicationSettings(),
			"cfn_pinpoint_baidu_channel": pinpoint.DatasourcePinpointBaiduChannel(),
			"cfn_pinpoint_campaign": pinpoint.DatasourcePinpointCampaign(),
			"cfn_pinpoint_email_channel": pinpoint.DatasourcePinpointEmailChannel(),
			"cfn_pinpoint_email_template": pinpoint.DatasourcePinpointEmailTemplate(),
			"cfn_pinpoint_event_stream": pinpoint.DatasourcePinpointEventStream(),
			"cfn_pinpoint_gcm_channel": pinpoint.DatasourcePinpointGCMChannel(),
			"cfn_pinpoint_push_template": pinpoint.DatasourcePinpointPushTemplate(),
			"cfn_pinpoint_sms_channel": pinpoint.DatasourcePinpointSMSChannel(),
			"cfn_pinpoint_segment": pinpoint.DatasourcePinpointSegment(),
			"cfn_pinpoint_sms_template": pinpoint.DatasourcePinpointSmsTemplate(),
			"cfn_pinpoint_voice_channel": pinpoint.DatasourcePinpointVoiceChannel(),
			"cfn_pinpointemail_configuration_set": pinpointemail.DatasourcePinpointEmailConfigurationSet(),
			"cfn_pinpointemail_configuration_set_event_destination": pinpointemail.DatasourcePinpointEmailConfigurationSetEventDestination(),
			"cfn_pinpointemail_dedicated_ip_pool": pinpointemail.DatasourcePinpointEmailDedicatedIpPool(),
			"cfn_pinpointemail_identity": pinpointemail.DatasourcePinpointEmailIdentity(),
			"cfn_qldb_ledger": qldb.DatasourceQLDBLedger(),
			"cfn_ram_resource_share": ram.DatasourceRAMResourceShare(),
			"cfn_rds_db_cluster": rds.DatasourceRDSDBCluster(),
			"cfn_rds_db_cluster_parameter_group": rds.DatasourceRDSDBClusterParameterGroup(),
			"cfn_rds_db_instance": rds.DatasourceRDSDBInstance(),
			"cfn_rds_db_parameter_group": rds.DatasourceRDSDBParameterGroup(),
			"cfn_rds_db_security_group": rds.DatasourceRDSDBSecurityGroup(),
			"cfn_rds_db_security_group_ingress": rds.DatasourceRDSDBSecurityGroupIngress(),
			"cfn_rds_db_subnet_group": rds.DatasourceRDSDBSubnetGroup(),
			"cfn_rds_event_subscription": rds.DatasourceRDSEventSubscription(),
			"cfn_rds_option_group": rds.DatasourceRDSOptionGroup(),
			"cfn_redshift_cluster": redshift.DatasourceRedshiftCluster(),
			"cfn_redshift_cluster_parameter_group": redshift.DatasourceRedshiftClusterParameterGroup(),
			"cfn_redshift_cluster_security_group": redshift.DatasourceRedshiftClusterSecurityGroup(),
			"cfn_redshift_cluster_security_group_ingress": redshift.DatasourceRedshiftClusterSecurityGroupIngress(),
			"cfn_redshift_cluster_subnet_group": redshift.DatasourceRedshiftClusterSubnetGroup(),
			"cfn_robomaker_fleet": robomaker.DatasourceRoboMakerFleet(),
			"cfn_robomaker_robot": robomaker.DatasourceRoboMakerRobot(),
			"cfn_robomaker_robot_application": robomaker.DatasourceRoboMakerRobotApplication(),
			"cfn_robomaker_robot_application_version": robomaker.DatasourceRoboMakerRobotApplicationVersion(),
			"cfn_robomaker_simulation_application": robomaker.DatasourceRoboMakerSimulationApplication(),
			"cfn_robomaker_simulation_application_version": robomaker.DatasourceRoboMakerSimulationApplicationVersion(),
			"cfn_route53_health_check": route53.DatasourceRoute53HealthCheck(),
			"cfn_route53_hosted_zone": route53.DatasourceRoute53HostedZone(),
			"cfn_route53_record_set": route53.DatasourceRoute53RecordSet(),
			"cfn_route53_record_set_group": route53.DatasourceRoute53RecordSetGroup(),
			"cfn_route53resolver_resolver_endpoint": route53resolver.DatasourceRoute53ResolverResolverEndpoint(),
			"cfn_route53resolver_resolver_rule": route53resolver.DatasourceRoute53ResolverResolverRule(),
			"cfn_route53resolver_resolver_rule_association": route53resolver.DatasourceRoute53ResolverResolverRuleAssociation(),
			"cfn_s3_access_point": s3.DatasourceS3AccessPoint(),
			"cfn_s3_bucket": s3.DatasourceS3Bucket(),
			"cfn_s3_bucket_policy": s3.DatasourceS3BucketPolicy(),
			"cfn_sagemaker_code_repository": sagemaker.DatasourceSageMakerCodeRepository(),
			"cfn_sagemaker_endpoint": sagemaker.DatasourceSageMakerEndpoint(),
			"cfn_sagemaker_endpoint_config": sagemaker.DatasourceSageMakerEndpointConfig(),
			"cfn_sagemaker_model": sagemaker.DatasourceSageMakerModel(),
			"cfn_sagemaker_notebook_instance": sagemaker.DatasourceSageMakerNotebookInstance(),
			"cfn_sagemaker_notebook_instance_lifecycle_config": sagemaker.DatasourceSageMakerNotebookInstanceLifecycleConfig(),
			"cfn_sagemaker_workteam": sagemaker.DatasourceSageMakerWorkteam(),
			"cfn_sdb_domain": sdb.DatasourceSDBDomain(),
			"cfn_secretsmanager_resource_policy": secretsmanager.DatasourceSecretsManagerResourcePolicy(),
			"cfn_secretsmanager_rotation_schedule": secretsmanager.DatasourceSecretsManagerRotationSchedule(),
			"cfn_secretsmanager_secret": secretsmanager.DatasourceSecretsManagerSecret(),
			"cfn_secretsmanager_secret_target_attachment": secretsmanager.DatasourceSecretsManagerSecretTargetAttachment(),
			"cfn_securityhub_hub": securityhub.DatasourceSecurityHubHub(),
			"cfn_servicecatalog_accepted_portfolio_share": servicecatalog.DatasourceServiceCatalogAcceptedPortfolioShare(),
			"cfn_servicecatalog_cloud_formation_product": servicecatalog.DatasourceServiceCatalogCloudFormationProduct(),
			"cfn_servicecatalog_cloud_formation_provisioned_product": servicecatalog.DatasourceServiceCatalogCloudFormationProvisionedProduct(),
			"cfn_servicecatalog_launch_notification_constraint": servicecatalog.DatasourceServiceCatalogLaunchNotificationConstraint(),
			"cfn_servicecatalog_launch_role_constraint": servicecatalog.DatasourceServiceCatalogLaunchRoleConstraint(),
			"cfn_servicecatalog_launch_template_constraint": servicecatalog.DatasourceServiceCatalogLaunchTemplateConstraint(),
			"cfn_servicecatalog_portfolio": servicecatalog.DatasourceServiceCatalogPortfolio(),
			"cfn_servicecatalog_portfolio_principal_association": servicecatalog.DatasourceServiceCatalogPortfolioPrincipalAssociation(),
			"cfn_servicecatalog_portfolio_product_association": servicecatalog.DatasourceServiceCatalogPortfolioProductAssociation(),
			"cfn_servicecatalog_portfolio_share": servicecatalog.DatasourceServiceCatalogPortfolioShare(),
			"cfn_servicecatalog_resource_update_constraint": servicecatalog.DatasourceServiceCatalogResourceUpdateConstraint(),
			"cfn_servicecatalog_stack_set_constraint": servicecatalog.DatasourceServiceCatalogStackSetConstraint(),
			"cfn_servicecatalog_tag_option": servicecatalog.DatasourceServiceCatalogTagOption(),
			"cfn_servicecatalog_tag_option_association": servicecatalog.DatasourceServiceCatalogTagOptionAssociation(),
			"cfn_servicediscovery_http_namespace": servicediscovery.DatasourceServiceDiscoveryHttpNamespace(),
			"cfn_servicediscovery_instance": servicediscovery.DatasourceServiceDiscoveryInstance(),
			"cfn_servicediscovery_private_dns_namespace": servicediscovery.DatasourceServiceDiscoveryPrivateDnsNamespace(),
			"cfn_servicediscovery_public_dns_namespace": servicediscovery.DatasourceServiceDiscoveryPublicDnsNamespace(),
			"cfn_servicediscovery_service": servicediscovery.DatasourceServiceDiscoveryService(),
			"cfn_ses_configuration_set": ses.DatasourceSESConfigurationSet(),
			"cfn_ses_configuration_set_event_destination": ses.DatasourceSESConfigurationSetEventDestination(),
			"cfn_ses_receipt_filter": ses.DatasourceSESReceiptFilter(),
			"cfn_ses_receipt_rule": ses.DatasourceSESReceiptRule(),
			"cfn_ses_receipt_rule_set": ses.DatasourceSESReceiptRuleSet(),
			"cfn_ses_template": ses.DatasourceSESTemplate(),
			"cfn_sns_subscription": sns.DatasourceSNSSubscription(),
			"cfn_sns_topic": sns.DatasourceSNSTopic(),
			"cfn_sns_topic_policy": sns.DatasourceSNSTopicPolicy(),
			"cfn_sqs_queue": sqs.DatasourceSQSQueue(),
			"cfn_sqs_queue_policy": sqs.DatasourceSQSQueuePolicy(),
			"cfn_ssm_association": ssm.DatasourceSSMAssociation(),
			"cfn_ssm_document": ssm.DatasourceSSMDocument(),
			"cfn_ssm_maintenance_window": ssm.DatasourceSSMMaintenanceWindow(),
			"cfn_ssm_maintenance_window_target": ssm.DatasourceSSMMaintenanceWindowTarget(),
			"cfn_ssm_maintenance_window_task": ssm.DatasourceSSMMaintenanceWindowTask(),
			"cfn_ssm_parameter": ssm.DatasourceSSMParameter(),
			"cfn_ssm_patch_baseline": ssm.DatasourceSSMPatchBaseline(),
			"cfn_ssm_resource_data_sync": ssm.DatasourceSSMResourceDataSync(),
			"cfn_stepfunctions_activity": stepfunctions.DatasourceStepFunctionsActivity(),
			"cfn_stepfunctions_state_machine": stepfunctions.DatasourceStepFunctionsStateMachine(),
			"cfn_transfer_server": transfer.DatasourceTransferServer(),
			"cfn_transfer_user": transfer.DatasourceTransferUser(),
			"cfn_waf_byte_match_set": waf.DatasourceWAFByteMatchSet(),
			"cfn_waf_ip_set": waf.DatasourceWAFIPSet(),
			"cfn_waf_rule": waf.DatasourceWAFRule(),
			"cfn_waf_size_constraint_set": waf.DatasourceWAFSizeConstraintSet(),
			"cfn_waf_sql_injection_match_set": waf.DatasourceWAFSqlInjectionMatchSet(),
			"cfn_waf_web_acl": waf.DatasourceWAFWebACL(),
			"cfn_waf_xss_match_set": waf.DatasourceWAFXssMatchSet(),
			"cfn_wafregional_byte_match_set": wafregional.DatasourceWAFRegionalByteMatchSet(),
			"cfn_wafregional_geo_match_set": wafregional.DatasourceWAFRegionalGeoMatchSet(),
			"cfn_wafregional_ip_set": wafregional.DatasourceWAFRegionalIPSet(),
			"cfn_wafregional_rate_based_rule": wafregional.DatasourceWAFRegionalRateBasedRule(),
			"cfn_wafregional_regex_pattern_set": wafregional.DatasourceWAFRegionalRegexPatternSet(),
			"cfn_wafregional_rule": wafregional.DatasourceWAFRegionalRule(),
			"cfn_wafregional_size_constraint_set": wafregional.DatasourceWAFRegionalSizeConstraintSet(),
			"cfn_wafregional_sql_injection_match_set": wafregional.DatasourceWAFRegionalSqlInjectionMatchSet(),
			"cfn_wafregional_web_acl": wafregional.DatasourceWAFRegionalWebACL(),
			"cfn_wafregional_web_acl_association": wafregional.DatasourceWAFRegionalWebACLAssociation(),
			"cfn_wafregional_xss_match_set": wafregional.DatasourceWAFRegionalXssMatchSet(),
			"cfn_wafv2_ip_set": wafv2.DatasourceWAFv2IPSet(),
			"cfn_wafv2_regex_pattern_set": wafv2.DatasourceWAFv2RegexPatternSet(),
			"cfn_wafv2_rule_group": wafv2.DatasourceWAFv2RuleGroup(),
			"cfn_wafv2_web_acl": wafv2.DatasourceWAFv2WebACL(),
			"cfn_workspaces_workspace": workspaces.DatasourceWorkSpacesWorkspace(),
			"cfn_template_data": misc.DatasourceTemplateData(),
		},
	}
}
