// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package provider

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/version"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/accessapproval"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/accesscontextmanager"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/activedirectory"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/alloydb"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/apigateway"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/apigee"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/appengine"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/artifactregistry"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/backupdr"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/beyondcorp"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/biglake"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigquery"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigqueryanalyticshub"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigqueryconnection"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigquerydatapolicy"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigquerydatatransfer"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigqueryreservation"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigtable"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/billing"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/binaryauthorization"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/certificatemanager"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudasset"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudbuild"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudbuildv2"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudfunctions"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudfunctions2"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudidentity"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudids"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudiot"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudrun"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudrunv2"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudscheduler"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudtasks"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/compute"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containeranalysis"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containerattached"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/corebilling"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/databasemigrationservice"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/datacatalog"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/datafusion"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/datalossprevention"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataplex"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataproc"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataprocmetastore"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/datastore"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/datastream"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/deploymentmanager"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dialogflow"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dialogflowcx"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dns"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/documentai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/documentaiwarehouse"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/essentialcontacts"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/filestore"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebase"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebasedatabase"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebaseextensions"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebasehosting"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firebasestorage"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/firestore"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gameservices"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkebackup"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkehub"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkehub2"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkeonprem"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/healthcare"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/iam2"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/iambeta"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/iamworkforcepool"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/iap"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/identityplatform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/kms"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/logging"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/looker"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/memcache"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/mlengine"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/monitoring"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/networkconnectivity"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/networkmanagement"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/networksecurity"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/networkservices"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/notebooks"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/orgpolicy"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/osconfig"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/oslogin"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/privateca"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/publicca"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/pubsub"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/pubsublite"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/redis"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/resourcemanager"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/runtimeconfig"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/secretmanager"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/securitycenter"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/securityscanner"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/servicedirectory"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/servicemanagement"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/serviceusage"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/sourcerepo"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/spanner"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/sql"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/storage"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/storagetransfer"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/tags"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/tpu"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/vertexai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/vmwareengine"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/vpcaccess"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/workflows"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/workstations"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/composer"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/container"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containeraws"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/containerazure"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataflow"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/servicenetworking"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"

	googleoauth "golang.org/x/oauth2/google"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {

	// The mtls service client gives the type of endpoint (mtls/regular)
	// at client creation. Since we use a shared client for requests we must
	// rewrite the endpoints to be mtls endpoints for the scenario where
	// mtls is enabled.
	if isMtls() {
		// if mtls is enabled switch all default endpoints to use the mtls endpoint
		for key, bp := range transport_tpg.DefaultBasePaths {
			transport_tpg.DefaultBasePaths[key] = getMtlsEndpoint(bp)
		}
	}

	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"credentials": {
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  ValidateCredentials,
				ConflictsWith: []string{"access_token"},
			},

			"access_token": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"credentials"},
			},

			"impersonate_service_account": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"impersonate_service_account_delegates": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"project": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"billing_project": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"zone": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"batching": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"send_after": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateNonNegativeDuration(),
						},
						"enable_batching": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"user_project_override": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"request_timeout": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"request_reason": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Generated Products
			"access_approval_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"access_context_manager_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"active_directory_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"alloydb_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"api_gateway_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"apigee_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"app_engine_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"artifact_registry_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"backup_dr_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"beyondcorp_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"biglake_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"big_query_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"bigquery_analytics_hub_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"bigquery_connection_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"bigquery_datapolicy_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"bigquery_data_transfer_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"bigquery_reservation_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"bigtable_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"billing_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"binary_authorization_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"certificate_manager_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloud_asset_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloud_build_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloudbuildv2_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloud_functions_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloudfunctions2_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloud_identity_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloud_ids_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloud_iot_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloud_run_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloud_run_v2_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloud_scheduler_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"cloud_tasks_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"compute_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"container_analysis_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"container_attached_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"core_billing_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"database_migration_service_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"data_catalog_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"dataform_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"data_fusion_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"data_loss_prevention_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"dataplex_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"dataproc_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"dataproc_metastore_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"datastore_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"datastream_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"deployment_manager_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"dialogflow_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"dialogflow_cx_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"dns_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"document_ai_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"document_ai_warehouse_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"essential_contacts_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"filestore_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"firebase_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"firebase_database_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"firebase_extensions_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"firebase_hosting_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"firebase_storage_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"firestore_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"game_services_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"gke_backup_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"gke_hub_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"gke_hub2_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"gkeonprem_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"healthcare_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"iam2_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"iam_beta_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"iam_workforce_pool_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"iap_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"identity_platform_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"kms_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"logging_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"looker_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"memcache_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"ml_engine_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"monitoring_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"network_connectivity_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"network_management_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"network_security_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"network_services_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"notebooks_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"org_policy_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"os_config_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"os_login_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"privateca_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"public_ca_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"pubsub_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"pubsub_lite_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"redis_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"resource_manager_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"runtime_config_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"secret_manager_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"security_center_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"security_scanner_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"service_directory_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"service_management_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"service_usage_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"source_repo_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"spanner_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"sql_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"storage_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"storage_transfer_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"tags_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"tpu_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"vertex_ai_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"vmwareengine_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"vpc_access_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"workflows_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},
			"workstations_custom_endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: transport_tpg.ValidateCustomEndpoint,
			},

			// Handwritten Products / Versioned / Atypical Entries
			transport_tpg.CloudBillingCustomEndpointEntryKey:      transport_tpg.CloudBillingCustomEndpointEntry,
			transport_tpg.ComposerCustomEndpointEntryKey:          transport_tpg.ComposerCustomEndpointEntry,
			transport_tpg.ContainerCustomEndpointEntryKey:         transport_tpg.ContainerCustomEndpointEntry,
			transport_tpg.DataflowCustomEndpointEntryKey:          transport_tpg.DataflowCustomEndpointEntry,
			transport_tpg.IamCredentialsCustomEndpointEntryKey:    transport_tpg.IamCredentialsCustomEndpointEntry,
			transport_tpg.ResourceManagerV3CustomEndpointEntryKey: transport_tpg.ResourceManagerV3CustomEndpointEntry,
			transport_tpg.RuntimeConfigCustomEndpointEntryKey:     transport_tpg.RuntimeConfigCustomEndpointEntry,
			transport_tpg.IAMCustomEndpointEntryKey:               transport_tpg.IAMCustomEndpointEntry,
			transport_tpg.ServiceNetworkingCustomEndpointEntryKey: transport_tpg.ServiceNetworkingCustomEndpointEntry,
			transport_tpg.TagsLocationCustomEndpointEntryKey:      transport_tpg.TagsLocationCustomEndpointEntry,

			// dcl
			transport_tpg.ContainerAwsCustomEndpointEntryKey:   transport_tpg.ContainerAwsCustomEndpointEntry,
			transport_tpg.ContainerAzureCustomEndpointEntryKey: transport_tpg.ContainerAzureCustomEndpointEntry,
		},

		ProviderMetaSchema: map[string]*schema.Schema{
			"module_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},

		DataSourcesMap: DatasourceMap(),
		ResourcesMap:   ResourceMap(),
	}

	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return ProviderConfigure(ctx, d, provider)
	}

	transport_tpg.ConfigureDCLProvider(provider)

	return provider
}

func DatasourceMap() map[string]*schema.Resource {
	datasourceMap, _ := DatasourceMapWithErrors()
	return datasourceMap
}

func DatasourceMapWithErrors() (map[string]*schema.Resource, error) {
	return mergeResourceMaps(map[string]*schema.Resource{
		// ####### START handwritten datasources ###########
		// ####### START datasources ###########
		"google_access_approval_folder_service_account":       accessapproval.DataSourceAccessApprovalFolderServiceAccount(),
		"google_access_approval_organization_service_account": accessapproval.DataSourceAccessApprovalOrganizationServiceAccount(),
		"google_access_approval_project_service_account":      accessapproval.DataSourceAccessApprovalProjectServiceAccount(),
		"google_active_folder":                                resourcemanager.DataSourceGoogleActiveFolder(),
		"google_alloydb_locations":                            alloydb.DataSourceAlloydbLocations(),
		"google_alloydb_supported_database_flags":             alloydb.DataSourceAlloydbSupportedDatabaseFlags(),
		"google_artifact_registry_repository":                 artifactregistry.DataSourceArtifactRegistryRepository(),
		"google_app_engine_default_service_account":           appengine.DataSourceGoogleAppEngineDefaultServiceAccount(),
		"google_beyondcorp_app_connection":                    beyondcorp.DataSourceGoogleBeyondcorpAppConnection(),
		"google_beyondcorp_app_connector":                     beyondcorp.DataSourceGoogleBeyondcorpAppConnector(),
		"google_beyondcorp_app_gateway":                       beyondcorp.DataSourceGoogleBeyondcorpAppGateway(),
		"google_billing_account":                              billing.DataSourceGoogleBillingAccount(),
		"google_bigquery_default_service_account":             bigquery.DataSourceGoogleBigqueryDefaultServiceAccount(),
		"google_cloudbuild_trigger":                           cloudbuild.DataSourceGoogleCloudBuildTrigger(),
		"google_cloudfunctions_function":                      cloudfunctions.DataSourceGoogleCloudFunctionsFunction(),
		"google_cloudfunctions2_function":                     cloudfunctions2.DataSourceGoogleCloudFunctions2Function(),
		"google_cloud_asset_resources_search_all":             cloudasset.DataSourceGoogleCloudAssetResourcesSearchAll(),
		"google_cloud_identity_groups":                        cloudidentity.DataSourceGoogleCloudIdentityGroups(),
		"google_cloud_identity_group_memberships":             cloudidentity.DataSourceGoogleCloudIdentityGroupMemberships(),
		"google_cloud_run_locations":                          cloudrun.DataSourceGoogleCloudRunLocations(),
		"google_cloud_run_service":                            cloudrun.DataSourceGoogleCloudRunService(),
		"google_composer_environment":                         composer.DataSourceGoogleComposerEnvironment(),
		"google_composer_image_versions":                      composer.DataSourceGoogleComposerImageVersions(),
		"google_compute_address":                              compute.DataSourceGoogleComputeAddress(),
		"google_compute_addresses":                            compute.DataSourceGoogleComputeAddresses(),
		"google_compute_backend_service":                      compute.DataSourceGoogleComputeBackendService(),
		"google_compute_backend_bucket":                       compute.DataSourceGoogleComputeBackendBucket(),
		"google_compute_default_service_account":              compute.DataSourceGoogleComputeDefaultServiceAccount(),
		"google_compute_disk":                                 compute.DataSourceGoogleComputeDisk(),
		"google_compute_forwarding_rule":                      compute.DataSourceGoogleComputeForwardingRule(),
		"google_compute_global_address":                       compute.DataSourceGoogleComputeGlobalAddress(),
		"google_compute_global_forwarding_rule":               compute.DataSourceGoogleComputeGlobalForwardingRule(),
		"google_compute_ha_vpn_gateway":                       compute.DataSourceGoogleComputeHaVpnGateway(),
		"google_compute_health_check":                         compute.DataSourceGoogleComputeHealthCheck(),
		"google_compute_image":                                compute.DataSourceGoogleComputeImage(),
		"google_compute_instance":                             compute.DataSourceGoogleComputeInstance(),
		"google_compute_instance_group":                       compute.DataSourceGoogleComputeInstanceGroup(),
		"google_compute_instance_group_manager":               compute.DataSourceGoogleComputeInstanceGroupManager(),
		"google_compute_instance_serial_port":                 compute.DataSourceGoogleComputeInstanceSerialPort(),
		"google_compute_instance_template":                    compute.DataSourceGoogleComputeInstanceTemplate(),
		"google_compute_lb_ip_ranges":                         compute.DataSourceGoogleComputeLbIpRanges(),
		"google_compute_network":                              compute.DataSourceGoogleComputeNetwork(),
		"google_compute_network_endpoint_group":               compute.DataSourceGoogleComputeNetworkEndpointGroup(),
		"google_compute_network_peering":                      compute.DataSourceComputeNetworkPeering(),
		"google_compute_node_types":                           compute.DataSourceGoogleComputeNodeTypes(),
		"google_compute_regions":                              compute.DataSourceGoogleComputeRegions(),
		"google_compute_region_network_endpoint_group":        compute.DataSourceGoogleComputeRegionNetworkEndpointGroup(),
		"google_compute_region_instance_group":                compute.DataSourceGoogleComputeRegionInstanceGroup(),
		"google_compute_region_instance_template":             compute.DataSourceGoogleComputeRegionInstanceTemplate(),
		"google_compute_region_ssl_certificate":               compute.DataSourceGoogleRegionComputeSslCertificate(),
		"google_compute_resource_policy":                      compute.DataSourceGoogleComputeResourcePolicy(),
		"google_compute_router":                               compute.DataSourceGoogleComputeRouter(),
		"google_compute_router_nat":                           compute.DataSourceGoogleComputeRouterNat(),
		"google_compute_router_status":                        compute.DataSourceGoogleComputeRouterStatus(),
		"google_compute_snapshot":                             compute.DataSourceGoogleComputeSnapshot(),
		"google_compute_ssl_certificate":                      compute.DataSourceGoogleComputeSslCertificate(),
		"google_compute_ssl_policy":                           compute.DataSourceGoogleComputeSslPolicy(),
		"google_compute_subnetwork":                           compute.DataSourceGoogleComputeSubnetwork(),
		"google_compute_vpn_gateway":                          compute.DataSourceGoogleComputeVpnGateway(),
		"google_compute_zones":                                compute.DataSourceGoogleComputeZones(),
		"google_container_azure_versions":                     containerazure.DataSourceGoogleContainerAzureVersions(),
		"google_container_aws_versions":                       containeraws.DataSourceGoogleContainerAwsVersions(),
		"google_container_attached_versions":                  containerattached.DataSourceGoogleContainerAttachedVersions(),
		"google_container_attached_install_manifest":          containerattached.DataSourceGoogleContainerAttachedInstallManifest(),
		"google_container_cluster":                            container.DataSourceGoogleContainerCluster(),
		"google_container_engine_versions":                    container.DataSourceGoogleContainerEngineVersions(),
		"google_container_registry_image":                     containeranalysis.DataSourceGoogleContainerImage(),
		"google_container_registry_repository":                containeranalysis.DataSourceGoogleContainerRepo(),
		"google_dataproc_metastore_service":                   dataprocmetastore.DataSourceDataprocMetastoreService(),
		"google_datastream_static_ips":                        datastream.DataSourceGoogleDatastreamStaticIps(),
		"google_game_services_game_server_deployment_rollout": gameservices.DataSourceGameServicesGameServerDeploymentRollout(),
		"google_iam_policy":                                   resourcemanager.DataSourceGoogleIamPolicy(),
		"google_iam_role":                                     resourcemanager.DataSourceGoogleIamRole(),
		"google_iam_testable_permissions":                     resourcemanager.DataSourceGoogleIamTestablePermissions(),
		"google_iam_workload_identity_pool":                   iambeta.DataSourceIAMBetaWorkloadIdentityPool(),
		"google_iam_workload_identity_pool_provider":          iambeta.DataSourceIAMBetaWorkloadIdentityPoolProvider(),
		"google_iap_client":                                   iap.DataSourceGoogleIapClient(),
		"google_kms_crypto_key":                               kms.DataSourceGoogleKmsCryptoKey(),
		"google_kms_crypto_key_version":                       kms.DataSourceGoogleKmsCryptoKeyVersion(),
		"google_kms_key_ring":                                 kms.DataSourceGoogleKmsKeyRing(),
		"google_kms_secret":                                   kms.DataSourceGoogleKmsSecret(),
		"google_kms_secret_ciphertext":                        kms.DataSourceGoogleKmsSecretCiphertext(),
		"google_kms_secret_asymmetric":                        kms.DataSourceGoogleKmsSecretAsymmetric(),
		"google_firebase_android_app":                         firebase.DataSourceGoogleFirebaseAndroidApp(),
		"google_firebase_apple_app":                           firebase.DataSourceGoogleFirebaseAppleApp(),
		"google_firebase_hosting_channel":                     firebasehosting.DataSourceGoogleFirebaseHostingChannel(),
		"google_firebase_web_app":                             firebase.DataSourceGoogleFirebaseWebApp(),
		"google_folder":                                       resourcemanager.DataSourceGoogleFolder(),
		"google_folders":                                      resourcemanager.DataSourceGoogleFolders(),
		"google_folder_organization_policy":                   resourcemanager.DataSourceGoogleFolderOrganizationPolicy(),
		"google_logging_project_cmek_settings":                logging.DataSourceGoogleLoggingProjectCmekSettings(),
		"google_logging_sink":                                 logging.DataSourceGoogleLoggingSink(),
		"google_monitoring_notification_channel":              monitoring.DataSourceMonitoringNotificationChannel(),
		"google_monitoring_cluster_istio_service":             monitoring.DataSourceMonitoringServiceClusterIstio(),
		"google_monitoring_istio_canonical_service":           monitoring.DataSourceMonitoringIstioCanonicalService(),
		"google_monitoring_mesh_istio_service":                monitoring.DataSourceMonitoringServiceMeshIstio(),
		"google_monitoring_app_engine_service":                monitoring.DataSourceMonitoringServiceAppEngine(),
		"google_monitoring_uptime_check_ips":                  monitoring.DataSourceGoogleMonitoringUptimeCheckIps(),
		"google_netblock_ip_ranges":                           resourcemanager.DataSourceGoogleNetblockIpRanges(),
		"google_organization":                                 resourcemanager.DataSourceGoogleOrganization(),
		"google_privateca_certificate_authority":              privateca.DataSourcePrivatecaCertificateAuthority(),
		"google_project":                                      resourcemanager.DataSourceGoogleProject(),
		"google_projects":                                     resourcemanager.DataSourceGoogleProjects(),
		"google_project_organization_policy":                  resourcemanager.DataSourceGoogleProjectOrganizationPolicy(),
		"google_project_service":                              resourcemanager.DataSourceGoogleProjectService(),
		"google_pubsub_subscription":                          pubsub.DataSourceGooglePubsubSubscription(),
		"google_pubsub_topic":                                 pubsub.DataSourceGooglePubsubTopic(),
		"google_runtimeconfig_config":                         runtimeconfig.DataSourceGoogleRuntimeconfigConfig(),
		"google_runtimeconfig_variable":                       runtimeconfig.DataSourceGoogleRuntimeconfigVariable(),
		"google_secret_manager_secret":                        secretmanager.DataSourceSecretManagerSecret(),
		"google_secret_manager_secret_version":                secretmanager.DataSourceSecretManagerSecretVersion(),
		"google_secret_manager_secret_version_access":         secretmanager.DataSourceSecretManagerSecretVersionAccess(),
		"google_service_account":                              resourcemanager.DataSourceGoogleServiceAccount(),
		"google_service_account_access_token":                 resourcemanager.DataSourceGoogleServiceAccountAccessToken(),
		"google_service_account_id_token":                     resourcemanager.DataSourceGoogleServiceAccountIdToken(),
		"google_service_account_jwt":                          resourcemanager.DataSourceGoogleServiceAccountJwt(),
		"google_service_account_key":                          resourcemanager.DataSourceGoogleServiceAccountKey(),
		"google_sourcerepo_repository":                        sourcerepo.DataSourceGoogleSourceRepoRepository(),
		"google_spanner_instance":                             spanner.DataSourceSpannerInstance(),
		"google_sql_ca_certs":                                 sql.DataSourceGoogleSQLCaCerts(),
		"google_sql_tiers":                                    sql.DataSourceGoogleSQLTiers(),
		"google_sql_database_instance_latest_recovery_time":   sql.DataSourceSqlDatabaseInstanceLatestRecoveryTime(),
		"google_sql_backup_run":                               sql.DataSourceSqlBackupRun(),
		"google_sql_databases":                                sql.DataSourceSqlDatabases(),
		"google_sql_database":                                 sql.DataSourceSqlDatabase(),
		"google_sql_database_instance":                        sql.DataSourceSqlDatabaseInstance(),
		"google_sql_database_instances":                       sql.DataSourceSqlDatabaseInstances(),
		"google_service_networking_peered_dns_domain":         servicenetworking.DataSourceGoogleServiceNetworkingPeeredDNSDomain(),
		"google_storage_bucket":                               storage.DataSourceGoogleStorageBucket(),
		"google_storage_bucket_object":                        storage.DataSourceGoogleStorageBucketObject(),
		"google_storage_bucket_object_content":                storage.DataSourceGoogleStorageBucketObjectContent(),
		"google_storage_object_signed_url":                    storage.DataSourceGoogleSignedUrl(),
		"google_storage_project_service_account":              storage.DataSourceGoogleStorageProjectServiceAccount(),
		"google_storage_transfer_project_service_account":     storagetransfer.DataSourceGoogleStorageTransferProjectServiceAccount(),
		"google_tags_tag_key":                                 tags.DataSourceGoogleTagsTagKey(),
		"google_tags_tag_value":                               tags.DataSourceGoogleTagsTagValue(),
		"google_tpu_tensorflow_versions":                      tpu.DataSourceTpuTensorflowVersions(),
		"google_vpc_access_connector":                         vpcaccess.DataSourceVPCAccessConnector(),
		"google_redis_instance":                               redis.DataSourceGoogleRedisInstance(),
		"google_vertex_ai_index":                              vertexai.DataSourceVertexAIIndex(),
		"google_vmwareengine_network":                         vmwareengine.DataSourceVmwareengineNetwork(),
		"google_vmwareengine_private_cloud":                   vmwareengine.DataSourceVmwareenginePrivateCloud(),
		"google_vmwareengine_cluster":                         vmwareengine.DataSourceVmwareengineCluster(),
		// ####### END datasources ###########
		// ####### END handwritten datasources ###########
	},
		map[string]*schema.Resource{
			// ####### START generated IAM datasources ###########
			"google_access_context_manager_access_policy_iam_policy": tpgiamresource.DataSourceIamPolicy(accesscontextmanager.AccessContextManagerAccessPolicyIamSchema, accesscontextmanager.AccessContextManagerAccessPolicyIamUpdaterProducer),
			"google_api_gateway_api_iam_policy":                      tpgiamresource.DataSourceIamPolicy(apigateway.ApiGatewayApiIamSchema, apigateway.ApiGatewayApiIamUpdaterProducer),
			"google_api_gateway_api_config_iam_policy":               tpgiamresource.DataSourceIamPolicy(apigateway.ApiGatewayApiConfigIamSchema, apigateway.ApiGatewayApiConfigIamUpdaterProducer),
			"google_api_gateway_gateway_iam_policy":                  tpgiamresource.DataSourceIamPolicy(apigateway.ApiGatewayGatewayIamSchema, apigateway.ApiGatewayGatewayIamUpdaterProducer),
			"google_apigee_environment_iam_policy":                   tpgiamresource.DataSourceIamPolicy(apigee.ApigeeEnvironmentIamSchema, apigee.ApigeeEnvironmentIamUpdaterProducer),
			"google_artifact_registry_repository_iam_policy":         tpgiamresource.DataSourceIamPolicy(artifactregistry.ArtifactRegistryRepositoryIamSchema, artifactregistry.ArtifactRegistryRepositoryIamUpdaterProducer),
			"google_bigquery_table_iam_policy":                       tpgiamresource.DataSourceIamPolicy(bigquery.BigQueryTableIamSchema, bigquery.BigQueryTableIamUpdaterProducer),
			"google_bigquery_analytics_hub_data_exchange_iam_policy": tpgiamresource.DataSourceIamPolicy(bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIamSchema, bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIamUpdaterProducer),
			"google_bigquery_analytics_hub_listing_iam_policy":       tpgiamresource.DataSourceIamPolicy(bigqueryanalyticshub.BigqueryAnalyticsHubListingIamSchema, bigqueryanalyticshub.BigqueryAnalyticsHubListingIamUpdaterProducer),
			"google_bigquery_connection_iam_policy":                  tpgiamresource.DataSourceIamPolicy(bigqueryconnection.BigqueryConnectionConnectionIamSchema, bigqueryconnection.BigqueryConnectionConnectionIamUpdaterProducer),
			"google_bigquery_datapolicy_data_policy_iam_policy":      tpgiamresource.DataSourceIamPolicy(bigquerydatapolicy.BigqueryDatapolicyDataPolicyIamSchema, bigquerydatapolicy.BigqueryDatapolicyDataPolicyIamUpdaterProducer),
			"google_binary_authorization_attestor_iam_policy":        tpgiamresource.DataSourceIamPolicy(binaryauthorization.BinaryAuthorizationAttestorIamSchema, binaryauthorization.BinaryAuthorizationAttestorIamUpdaterProducer),
			"google_cloudbuildv2_connection_iam_policy":              tpgiamresource.DataSourceIamPolicy(cloudbuildv2.Cloudbuildv2ConnectionIamSchema, cloudbuildv2.Cloudbuildv2ConnectionIamUpdaterProducer),
			"google_cloudfunctions_function_iam_policy":              tpgiamresource.DataSourceIamPolicy(cloudfunctions.CloudFunctionsCloudFunctionIamSchema, cloudfunctions.CloudFunctionsCloudFunctionIamUpdaterProducer),
			"google_cloudfunctions2_function_iam_policy":             tpgiamresource.DataSourceIamPolicy(cloudfunctions2.Cloudfunctions2functionIamSchema, cloudfunctions2.Cloudfunctions2functionIamUpdaterProducer),
			"google_cloudiot_registry_iam_policy":                    tpgiamresource.DataSourceIamPolicy(cloudiot.CloudIotDeviceRegistryIamSchema, cloudiot.CloudIotDeviceRegistryIamUpdaterProducer),
			"google_cloud_run_service_iam_policy":                    tpgiamresource.DataSourceIamPolicy(cloudrun.CloudRunServiceIamSchema, cloudrun.CloudRunServiceIamUpdaterProducer),
			"google_cloud_run_v2_job_iam_policy":                     tpgiamresource.DataSourceIamPolicy(cloudrunv2.CloudRunV2JobIamSchema, cloudrunv2.CloudRunV2JobIamUpdaterProducer),
			"google_cloud_run_v2_service_iam_policy":                 tpgiamresource.DataSourceIamPolicy(cloudrunv2.CloudRunV2ServiceIamSchema, cloudrunv2.CloudRunV2ServiceIamUpdaterProducer),
			"google_cloud_tasks_queue_iam_policy":                    tpgiamresource.DataSourceIamPolicy(cloudtasks.CloudTasksQueueIamSchema, cloudtasks.CloudTasksQueueIamUpdaterProducer),
			"google_compute_backend_bucket_iam_policy":               tpgiamresource.DataSourceIamPolicy(compute.ComputeBackendBucketIamSchema, compute.ComputeBackendBucketIamUpdaterProducer),
			"google_compute_backend_service_iam_policy":              tpgiamresource.DataSourceIamPolicy(compute.ComputeBackendServiceIamSchema, compute.ComputeBackendServiceIamUpdaterProducer),
			"google_compute_disk_iam_policy":                         tpgiamresource.DataSourceIamPolicy(compute.ComputeDiskIamSchema, compute.ComputeDiskIamUpdaterProducer),
			"google_compute_image_iam_policy":                        tpgiamresource.DataSourceIamPolicy(compute.ComputeImageIamSchema, compute.ComputeImageIamUpdaterProducer),
			"google_compute_instance_iam_policy":                     tpgiamresource.DataSourceIamPolicy(compute.ComputeInstanceIamSchema, compute.ComputeInstanceIamUpdaterProducer),
			"google_compute_machine_image_iam_policy":                tpgiamresource.DataSourceIamPolicy(compute.ComputeMachineImageIamSchema, compute.ComputeMachineImageIamUpdaterProducer),
			"google_compute_region_backend_service_iam_policy":       tpgiamresource.DataSourceIamPolicy(compute.ComputeRegionBackendServiceIamSchema, compute.ComputeRegionBackendServiceIamUpdaterProducer),
			"google_compute_region_disk_iam_policy":                  tpgiamresource.DataSourceIamPolicy(compute.ComputeRegionDiskIamSchema, compute.ComputeRegionDiskIamUpdaterProducer),
			"google_compute_snapshot_iam_policy":                     tpgiamresource.DataSourceIamPolicy(compute.ComputeSnapshotIamSchema, compute.ComputeSnapshotIamUpdaterProducer),
			"google_compute_subnetwork_iam_policy":                   tpgiamresource.DataSourceIamPolicy(compute.ComputeSubnetworkIamSchema, compute.ComputeSubnetworkIamUpdaterProducer),
			"google_container_analysis_note_iam_policy":              tpgiamresource.DataSourceIamPolicy(containeranalysis.ContainerAnalysisNoteIamSchema, containeranalysis.ContainerAnalysisNoteIamUpdaterProducer),
			"google_data_catalog_entry_group_iam_policy":             tpgiamresource.DataSourceIamPolicy(datacatalog.DataCatalogEntryGroupIamSchema, datacatalog.DataCatalogEntryGroupIamUpdaterProducer),
			"google_data_catalog_policy_tag_iam_policy":              tpgiamresource.DataSourceIamPolicy(datacatalog.DataCatalogPolicyTagIamSchema, datacatalog.DataCatalogPolicyTagIamUpdaterProducer),
			"google_data_catalog_tag_template_iam_policy":            tpgiamresource.DataSourceIamPolicy(datacatalog.DataCatalogTagTemplateIamSchema, datacatalog.DataCatalogTagTemplateIamUpdaterProducer),
			"google_data_catalog_taxonomy_iam_policy":                tpgiamresource.DataSourceIamPolicy(datacatalog.DataCatalogTaxonomyIamSchema, datacatalog.DataCatalogTaxonomyIamUpdaterProducer),
			"google_data_fusion_instance_iam_policy":                 tpgiamresource.DataSourceIamPolicy(datafusion.DataFusionInstanceIamSchema, datafusion.DataFusionInstanceIamUpdaterProducer),
			"google_dataplex_asset_iam_policy":                       tpgiamresource.DataSourceIamPolicy(dataplex.DataplexAssetIamSchema, dataplex.DataplexAssetIamUpdaterProducer),
			"google_dataplex_datascan_iam_policy":                    tpgiamresource.DataSourceIamPolicy(dataplex.DataplexDatascanIamSchema, dataplex.DataplexDatascanIamUpdaterProducer),
			"google_dataplex_lake_iam_policy":                        tpgiamresource.DataSourceIamPolicy(dataplex.DataplexLakeIamSchema, dataplex.DataplexLakeIamUpdaterProducer),
			"google_dataplex_task_iam_policy":                        tpgiamresource.DataSourceIamPolicy(dataplex.DataplexTaskIamSchema, dataplex.DataplexTaskIamUpdaterProducer),
			"google_dataplex_zone_iam_policy":                        tpgiamresource.DataSourceIamPolicy(dataplex.DataplexZoneIamSchema, dataplex.DataplexZoneIamUpdaterProducer),
			"google_dataproc_autoscaling_policy_iam_policy":          tpgiamresource.DataSourceIamPolicy(dataproc.DataprocAutoscalingPolicyIamSchema, dataproc.DataprocAutoscalingPolicyIamUpdaterProducer),
			"google_dataproc_metastore_federation_iam_policy":        tpgiamresource.DataSourceIamPolicy(dataprocmetastore.DataprocMetastoreFederationIamSchema, dataprocmetastore.DataprocMetastoreFederationIamUpdaterProducer),
			"google_dataproc_metastore_service_iam_policy":           tpgiamresource.DataSourceIamPolicy(dataprocmetastore.DataprocMetastoreServiceIamSchema, dataprocmetastore.DataprocMetastoreServiceIamUpdaterProducer),
			"google_dns_managed_zone_iam_policy":                     tpgiamresource.DataSourceIamPolicy(dns.DNSManagedZoneIamSchema, dns.DNSManagedZoneIamUpdaterProducer),
			"google_gke_backup_backup_plan_iam_policy":               tpgiamresource.DataSourceIamPolicy(gkebackup.GKEBackupBackupPlanIamSchema, gkebackup.GKEBackupBackupPlanIamUpdaterProducer),
			"google_gke_hub_membership_iam_policy":                   tpgiamresource.DataSourceIamPolicy(gkehub.GKEHubMembershipIamSchema, gkehub.GKEHubMembershipIamUpdaterProducer),
			"google_gke_hub_feature_iam_policy":                      tpgiamresource.DataSourceIamPolicy(gkehub2.GKEHub2FeatureIamSchema, gkehub2.GKEHub2FeatureIamUpdaterProducer),
			"google_gke_hub_scope_iam_policy":                        tpgiamresource.DataSourceIamPolicy(gkehub2.GKEHub2ScopeIamSchema, gkehub2.GKEHub2ScopeIamUpdaterProducer),
			"google_healthcare_consent_store_iam_policy":             tpgiamresource.DataSourceIamPolicy(healthcare.HealthcareConsentStoreIamSchema, healthcare.HealthcareConsentStoreIamUpdaterProducer),
			"google_iap_app_engine_service_iam_policy":               tpgiamresource.DataSourceIamPolicy(iap.IapAppEngineServiceIamSchema, iap.IapAppEngineServiceIamUpdaterProducer),
			"google_iap_app_engine_version_iam_policy":               tpgiamresource.DataSourceIamPolicy(iap.IapAppEngineVersionIamSchema, iap.IapAppEngineVersionIamUpdaterProducer),
			"google_iap_tunnel_iam_policy":                           tpgiamresource.DataSourceIamPolicy(iap.IapTunnelIamSchema, iap.IapTunnelIamUpdaterProducer),
			"google_iap_tunnel_instance_iam_policy":                  tpgiamresource.DataSourceIamPolicy(iap.IapTunnelInstanceIamSchema, iap.IapTunnelInstanceIamUpdaterProducer),
			"google_iap_web_iam_policy":                              tpgiamresource.DataSourceIamPolicy(iap.IapWebIamSchema, iap.IapWebIamUpdaterProducer),
			"google_iap_web_backend_service_iam_policy":              tpgiamresource.DataSourceIamPolicy(iap.IapWebBackendServiceIamSchema, iap.IapWebBackendServiceIamUpdaterProducer),
			"google_iap_web_region_backend_service_iam_policy":       tpgiamresource.DataSourceIamPolicy(iap.IapWebRegionBackendServiceIamSchema, iap.IapWebRegionBackendServiceIamUpdaterProducer),
			"google_iap_web_type_app_engine_iam_policy":              tpgiamresource.DataSourceIamPolicy(iap.IapWebTypeAppEngineIamSchema, iap.IapWebTypeAppEngineIamUpdaterProducer),
			"google_iap_web_type_compute_iam_policy":                 tpgiamresource.DataSourceIamPolicy(iap.IapWebTypeComputeIamSchema, iap.IapWebTypeComputeIamUpdaterProducer),
			"google_notebooks_instance_iam_policy":                   tpgiamresource.DataSourceIamPolicy(notebooks.NotebooksInstanceIamSchema, notebooks.NotebooksInstanceIamUpdaterProducer),
			"google_notebooks_runtime_iam_policy":                    tpgiamresource.DataSourceIamPolicy(notebooks.NotebooksRuntimeIamSchema, notebooks.NotebooksRuntimeIamUpdaterProducer),
			"google_privateca_ca_pool_iam_policy":                    tpgiamresource.DataSourceIamPolicy(privateca.PrivatecaCaPoolIamSchema, privateca.PrivatecaCaPoolIamUpdaterProducer),
			"google_privateca_certificate_template_iam_policy":       tpgiamresource.DataSourceIamPolicy(privateca.PrivatecaCertificateTemplateIamSchema, privateca.PrivatecaCertificateTemplateIamUpdaterProducer),
			"google_pubsub_topic_iam_policy":                         tpgiamresource.DataSourceIamPolicy(pubsub.PubsubTopicIamSchema, pubsub.PubsubTopicIamUpdaterProducer),
			"google_runtimeconfig_config_iam_policy":                 tpgiamresource.DataSourceIamPolicy(runtimeconfig.RuntimeConfigConfigIamSchema, runtimeconfig.RuntimeConfigConfigIamUpdaterProducer),
			"google_secret_manager_secret_iam_policy":                tpgiamresource.DataSourceIamPolicy(secretmanager.SecretManagerSecretIamSchema, secretmanager.SecretManagerSecretIamUpdaterProducer),
			"google_scc_source_iam_policy":                           tpgiamresource.DataSourceIamPolicy(securitycenter.SecurityCenterSourceIamSchema, securitycenter.SecurityCenterSourceIamUpdaterProducer),
			"google_service_directory_namespace_iam_policy":          tpgiamresource.DataSourceIamPolicy(servicedirectory.ServiceDirectoryNamespaceIamSchema, servicedirectory.ServiceDirectoryNamespaceIamUpdaterProducer),
			"google_service_directory_service_iam_policy":            tpgiamresource.DataSourceIamPolicy(servicedirectory.ServiceDirectoryServiceIamSchema, servicedirectory.ServiceDirectoryServiceIamUpdaterProducer),
			"google_endpoints_service_iam_policy":                    tpgiamresource.DataSourceIamPolicy(servicemanagement.ServiceManagementServiceIamSchema, servicemanagement.ServiceManagementServiceIamUpdaterProducer),
			"google_endpoints_service_consumers_iam_policy":          tpgiamresource.DataSourceIamPolicy(servicemanagement.ServiceManagementServiceConsumersIamSchema, servicemanagement.ServiceManagementServiceConsumersIamUpdaterProducer),
			"google_sourcerepo_repository_iam_policy":                tpgiamresource.DataSourceIamPolicy(sourcerepo.SourceRepoRepositoryIamSchema, sourcerepo.SourceRepoRepositoryIamUpdaterProducer),
			"google_storage_bucket_iam_policy":                       tpgiamresource.DataSourceIamPolicy(storage.StorageBucketIamSchema, storage.StorageBucketIamUpdaterProducer),
			"google_tags_tag_key_iam_policy":                         tpgiamresource.DataSourceIamPolicy(tags.TagsTagKeyIamSchema, tags.TagsTagKeyIamUpdaterProducer),
			"google_tags_tag_value_iam_policy":                       tpgiamresource.DataSourceIamPolicy(tags.TagsTagValueIamSchema, tags.TagsTagValueIamUpdaterProducer),
			"google_vertex_ai_featurestore_iam_policy":               tpgiamresource.DataSourceIamPolicy(vertexai.VertexAIFeaturestoreIamSchema, vertexai.VertexAIFeaturestoreIamUpdaterProducer),
			"google_vertex_ai_featurestore_entitytype_iam_policy":    tpgiamresource.DataSourceIamPolicy(vertexai.VertexAIFeaturestoreEntitytypeIamSchema, vertexai.VertexAIFeaturestoreEntitytypeIamUpdaterProducer),
			"google_workstations_workstation_iam_policy":             tpgiamresource.DataSourceIamPolicy(workstations.WorkstationsWorkstationIamSchema, workstations.WorkstationsWorkstationIamUpdaterProducer),
			"google_workstations_workstation_config_iam_policy":      tpgiamresource.DataSourceIamPolicy(workstations.WorkstationsWorkstationConfigIamSchema, workstations.WorkstationsWorkstationConfigIamUpdaterProducer),
			// ####### END generated IAM datasources ###########
		},
		map[string]*schema.Resource{
			// ####### START non-generated IAM datasources ###########
			"google_bigtable_instance_iam_policy":       tpgiamresource.DataSourceIamPolicy(bigtable.IamBigtableInstanceSchema, bigtable.NewBigtableInstanceUpdater),
			"google_bigtable_table_iam_policy":          tpgiamresource.DataSourceIamPolicy(bigtable.IamBigtableTableSchema, bigtable.NewBigtableTableUpdater),
			"google_bigquery_dataset_iam_policy":        tpgiamresource.DataSourceIamPolicy(bigquery.IamBigqueryDatasetSchema, bigquery.NewBigqueryDatasetIamUpdater),
			"google_billing_account_iam_policy":         tpgiamresource.DataSourceIamPolicy(billing.IamBillingAccountSchema, billing.NewBillingAccountIamUpdater),
			"google_dataproc_cluster_iam_policy":        tpgiamresource.DataSourceIamPolicy(dataproc.IamDataprocClusterSchema, dataproc.NewDataprocClusterUpdater),
			"google_dataproc_job_iam_policy":            tpgiamresource.DataSourceIamPolicy(dataproc.IamDataprocJobSchema, dataproc.NewDataprocJobUpdater),
			"google_folder_iam_policy":                  tpgiamresource.DataSourceIamPolicy(resourcemanager.IamFolderSchema, resourcemanager.NewFolderIamUpdater),
			"google_healthcare_dataset_iam_policy":      tpgiamresource.DataSourceIamPolicy(healthcare.IamHealthcareDatasetSchema, healthcare.NewHealthcareDatasetIamUpdater),
			"google_healthcare_dicom_store_iam_policy":  tpgiamresource.DataSourceIamPolicy(healthcare.IamHealthcareDicomStoreSchema, healthcare.NewHealthcareDicomStoreIamUpdater),
			"google_healthcare_fhir_store_iam_policy":   tpgiamresource.DataSourceIamPolicy(healthcare.IamHealthcareFhirStoreSchema, healthcare.NewHealthcareFhirStoreIamUpdater),
			"google_healthcare_hl7_v2_store_iam_policy": tpgiamresource.DataSourceIamPolicy(healthcare.IamHealthcareHl7V2StoreSchema, healthcare.NewHealthcareHl7V2StoreIamUpdater),
			"google_kms_key_ring_iam_policy":            tpgiamresource.DataSourceIamPolicy(kms.IamKmsKeyRingSchema, kms.NewKmsKeyRingIamUpdater),
			"google_kms_crypto_key_iam_policy":          tpgiamresource.DataSourceIamPolicy(kms.IamKmsCryptoKeySchema, kms.NewKmsCryptoKeyIamUpdater),
			"google_spanner_instance_iam_policy":        tpgiamresource.DataSourceIamPolicy(spanner.IamSpannerInstanceSchema, spanner.NewSpannerInstanceIamUpdater),
			"google_spanner_database_iam_policy":        tpgiamresource.DataSourceIamPolicy(spanner.IamSpannerDatabaseSchema, spanner.NewSpannerDatabaseIamUpdater),
			"google_organization_iam_policy":            tpgiamresource.DataSourceIamPolicy(resourcemanager.IamOrganizationSchema, resourcemanager.NewOrganizationIamUpdater),
			"google_project_iam_policy":                 tpgiamresource.DataSourceIamPolicy(resourcemanager.IamProjectSchema, resourcemanager.NewProjectIamUpdater),
			"google_pubsub_subscription_iam_policy":     tpgiamresource.DataSourceIamPolicy(pubsub.IamPubsubSubscriptionSchema, pubsub.NewPubsubSubscriptionIamUpdater),
			"google_service_account_iam_policy":         tpgiamresource.DataSourceIamPolicy(resourcemanager.IamServiceAccountSchema, resourcemanager.NewServiceAccountIamUpdater),
			// ####### END non-generated IAM datasources ###########
		})
}

// Generated resources: 377
// Generated IAM resources: 237
// Total generated resources: 614
func ResourceMap() map[string]*schema.Resource {
	resourceMap, _ := ResourceMapWithErrors()
	return resourceMap
}

func ResourceMapWithErrors() (map[string]*schema.Resource, error) {
	return mergeResourceMaps(
		map[string]*schema.Resource{
			"google_folder_access_approval_settings":                         accessapproval.ResourceAccessApprovalFolderSettings(),
			"google_organization_access_approval_settings":                   accessapproval.ResourceAccessApprovalOrganizationSettings(),
			"google_project_access_approval_settings":                        accessapproval.ResourceAccessApprovalProjectSettings(),
			"google_access_context_manager_access_level":                     accesscontextmanager.ResourceAccessContextManagerAccessLevel(),
			"google_access_context_manager_access_level_condition":           accesscontextmanager.ResourceAccessContextManagerAccessLevelCondition(),
			"google_access_context_manager_access_levels":                    accesscontextmanager.ResourceAccessContextManagerAccessLevels(),
			"google_access_context_manager_access_policy":                    accesscontextmanager.ResourceAccessContextManagerAccessPolicy(),
			"google_access_context_manager_access_policy_iam_binding":        tpgiamresource.ResourceIamBinding(accesscontextmanager.AccessContextManagerAccessPolicyIamSchema, accesscontextmanager.AccessContextManagerAccessPolicyIamUpdaterProducer, accesscontextmanager.AccessContextManagerAccessPolicyIdParseFunc),
			"google_access_context_manager_access_policy_iam_member":         tpgiamresource.ResourceIamMember(accesscontextmanager.AccessContextManagerAccessPolicyIamSchema, accesscontextmanager.AccessContextManagerAccessPolicyIamUpdaterProducer, accesscontextmanager.AccessContextManagerAccessPolicyIdParseFunc),
			"google_access_context_manager_access_policy_iam_policy":         tpgiamresource.ResourceIamPolicy(accesscontextmanager.AccessContextManagerAccessPolicyIamSchema, accesscontextmanager.AccessContextManagerAccessPolicyIamUpdaterProducer, accesscontextmanager.AccessContextManagerAccessPolicyIdParseFunc),
			"google_access_context_manager_authorized_orgs_desc":             accesscontextmanager.ResourceAccessContextManagerAuthorizedOrgsDesc(),
			"google_access_context_manager_egress_policy":                    accesscontextmanager.ResourceAccessContextManagerEgressPolicy(),
			"google_access_context_manager_gcp_user_access_binding":          accesscontextmanager.ResourceAccessContextManagerGcpUserAccessBinding(),
			"google_access_context_manager_ingress_policy":                   accesscontextmanager.ResourceAccessContextManagerIngressPolicy(),
			"google_access_context_manager_service_perimeter":                accesscontextmanager.ResourceAccessContextManagerServicePerimeter(),
			"google_access_context_manager_service_perimeter_egress_policy":  accesscontextmanager.ResourceAccessContextManagerServicePerimeterEgressPolicy(),
			"google_access_context_manager_service_perimeter_ingress_policy": accesscontextmanager.ResourceAccessContextManagerServicePerimeterIngressPolicy(),
			"google_access_context_manager_service_perimeter_resource":       accesscontextmanager.ResourceAccessContextManagerServicePerimeterResource(),
			"google_access_context_manager_service_perimeters":               accesscontextmanager.ResourceAccessContextManagerServicePerimeters(),
			"google_active_directory_domain":                                 activedirectory.ResourceActiveDirectoryDomain(),
			"google_active_directory_domain_trust":                           activedirectory.ResourceActiveDirectoryDomainTrust(),
			"google_active_directory_peering":                                activedirectory.ResourceActiveDirectoryPeering(),
			"google_alloydb_backup":                                          alloydb.ResourceAlloydbBackup(),
			"google_alloydb_cluster":                                         alloydb.ResourceAlloydbCluster(),
			"google_alloydb_instance":                                        alloydb.ResourceAlloydbInstance(),
			"google_api_gateway_api":                                         apigateway.ResourceApiGatewayApi(),
			"google_api_gateway_api_iam_binding":                             tpgiamresource.ResourceIamBinding(apigateway.ApiGatewayApiIamSchema, apigateway.ApiGatewayApiIamUpdaterProducer, apigateway.ApiGatewayApiIdParseFunc),
			"google_api_gateway_api_iam_member":                              tpgiamresource.ResourceIamMember(apigateway.ApiGatewayApiIamSchema, apigateway.ApiGatewayApiIamUpdaterProducer, apigateway.ApiGatewayApiIdParseFunc),
			"google_api_gateway_api_iam_policy":                              tpgiamresource.ResourceIamPolicy(apigateway.ApiGatewayApiIamSchema, apigateway.ApiGatewayApiIamUpdaterProducer, apigateway.ApiGatewayApiIdParseFunc),
			"google_api_gateway_api_config":                                  apigateway.ResourceApiGatewayApiConfig(),
			"google_api_gateway_api_config_iam_binding":                      tpgiamresource.ResourceIamBinding(apigateway.ApiGatewayApiConfigIamSchema, apigateway.ApiGatewayApiConfigIamUpdaterProducer, apigateway.ApiGatewayApiConfigIdParseFunc),
			"google_api_gateway_api_config_iam_member":                       tpgiamresource.ResourceIamMember(apigateway.ApiGatewayApiConfigIamSchema, apigateway.ApiGatewayApiConfigIamUpdaterProducer, apigateway.ApiGatewayApiConfigIdParseFunc),
			"google_api_gateway_api_config_iam_policy":                       tpgiamresource.ResourceIamPolicy(apigateway.ApiGatewayApiConfigIamSchema, apigateway.ApiGatewayApiConfigIamUpdaterProducer, apigateway.ApiGatewayApiConfigIdParseFunc),
			"google_api_gateway_gateway":                                     apigateway.ResourceApiGatewayGateway(),
			"google_api_gateway_gateway_iam_binding":                         tpgiamresource.ResourceIamBinding(apigateway.ApiGatewayGatewayIamSchema, apigateway.ApiGatewayGatewayIamUpdaterProducer, apigateway.ApiGatewayGatewayIdParseFunc),
			"google_api_gateway_gateway_iam_member":                          tpgiamresource.ResourceIamMember(apigateway.ApiGatewayGatewayIamSchema, apigateway.ApiGatewayGatewayIamUpdaterProducer, apigateway.ApiGatewayGatewayIdParseFunc),
			"google_api_gateway_gateway_iam_policy":                          tpgiamresource.ResourceIamPolicy(apigateway.ApiGatewayGatewayIamSchema, apigateway.ApiGatewayGatewayIamUpdaterProducer, apigateway.ApiGatewayGatewayIdParseFunc),
			"google_apigee_addons_config":                                    apigee.ResourceApigeeAddonsConfig(),
			"google_apigee_endpoint_attachment":                              apigee.ResourceApigeeEndpointAttachment(),
			"google_apigee_env_keystore":                                     apigee.ResourceApigeeEnvKeystore(),
			"google_apigee_env_references":                                   apigee.ResourceApigeeEnvReferences(),
			"google_apigee_envgroup":                                         apigee.ResourceApigeeEnvgroup(),
			"google_apigee_envgroup_attachment":                              apigee.ResourceApigeeEnvgroupAttachment(),
			"google_apigee_environment":                                      apigee.ResourceApigeeEnvironment(),
			"google_apigee_environment_iam_binding":                          tpgiamresource.ResourceIamBinding(apigee.ApigeeEnvironmentIamSchema, apigee.ApigeeEnvironmentIamUpdaterProducer, apigee.ApigeeEnvironmentIdParseFunc),
			"google_apigee_environment_iam_member":                           tpgiamresource.ResourceIamMember(apigee.ApigeeEnvironmentIamSchema, apigee.ApigeeEnvironmentIamUpdaterProducer, apigee.ApigeeEnvironmentIdParseFunc),
			"google_apigee_environment_iam_policy":                           tpgiamresource.ResourceIamPolicy(apigee.ApigeeEnvironmentIamSchema, apigee.ApigeeEnvironmentIamUpdaterProducer, apigee.ApigeeEnvironmentIdParseFunc),
			"google_apigee_instance":                                         apigee.ResourceApigeeInstance(),
			"google_apigee_instance_attachment":                              apigee.ResourceApigeeInstanceAttachment(),
			"google_apigee_keystores_aliases_self_signed_cert":               apigee.ResourceApigeeKeystoresAliasesSelfSignedCert(),
			"google_apigee_nat_address":                                      apigee.ResourceApigeeNatAddress(),
			"google_apigee_organization":                                     apigee.ResourceApigeeOrganization(),
			"google_apigee_sync_authorization":                               apigee.ResourceApigeeSyncAuthorization(),
			"google_app_engine_application_url_dispatch_rules":               appengine.ResourceAppEngineApplicationUrlDispatchRules(),
			"google_app_engine_domain_mapping":                               appengine.ResourceAppEngineDomainMapping(),
			"google_app_engine_firewall_rule":                                appengine.ResourceAppEngineFirewallRule(),
			"google_app_engine_flexible_app_version":                         appengine.ResourceAppEngineFlexibleAppVersion(),
			"google_app_engine_service_network_settings":                     appengine.ResourceAppEngineServiceNetworkSettings(),
			"google_app_engine_service_split_traffic":                        appengine.ResourceAppEngineServiceSplitTraffic(),
			"google_app_engine_standard_app_version":                         appengine.ResourceAppEngineStandardAppVersion(),
			"google_artifact_registry_repository":                            artifactregistry.ResourceArtifactRegistryRepository(),
			"google_artifact_registry_repository_iam_binding":                tpgiamresource.ResourceIamBinding(artifactregistry.ArtifactRegistryRepositoryIamSchema, artifactregistry.ArtifactRegistryRepositoryIamUpdaterProducer, artifactregistry.ArtifactRegistryRepositoryIdParseFunc),
			"google_artifact_registry_repository_iam_member":                 tpgiamresource.ResourceIamMember(artifactregistry.ArtifactRegistryRepositoryIamSchema, artifactregistry.ArtifactRegistryRepositoryIamUpdaterProducer, artifactregistry.ArtifactRegistryRepositoryIdParseFunc),
			"google_artifact_registry_repository_iam_policy":                 tpgiamresource.ResourceIamPolicy(artifactregistry.ArtifactRegistryRepositoryIamSchema, artifactregistry.ArtifactRegistryRepositoryIamUpdaterProducer, artifactregistry.ArtifactRegistryRepositoryIdParseFunc),
			"google_backup_dr_management_server":                             backupdr.ResourceBackupDRManagementServer(),
			"google_beyondcorp_app_connection":                               beyondcorp.ResourceBeyondcorpAppConnection(),
			"google_beyondcorp_app_connector":                                beyondcorp.ResourceBeyondcorpAppConnector(),
			"google_beyondcorp_app_gateway":                                  beyondcorp.ResourceBeyondcorpAppGateway(),
			"google_biglake_catalog":                                         biglake.ResourceBiglakeCatalog(),
			"google_biglake_database":                                        biglake.ResourceBiglakeDatabase(),
			"google_bigquery_dataset":                                        bigquery.ResourceBigQueryDataset(),
			"google_bigquery_dataset_access":                                 bigquery.ResourceBigQueryDatasetAccess(),
			"google_bigquery_job":                                            bigquery.ResourceBigQueryJob(),
			"google_bigquery_routine":                                        bigquery.ResourceBigQueryRoutine(),
			"google_bigquery_table_iam_binding":                              tpgiamresource.ResourceIamBinding(bigquery.BigQueryTableIamSchema, bigquery.BigQueryTableIamUpdaterProducer, bigquery.BigQueryTableIdParseFunc),
			"google_bigquery_table_iam_member":                               tpgiamresource.ResourceIamMember(bigquery.BigQueryTableIamSchema, bigquery.BigQueryTableIamUpdaterProducer, bigquery.BigQueryTableIdParseFunc),
			"google_bigquery_table_iam_policy":                               tpgiamresource.ResourceIamPolicy(bigquery.BigQueryTableIamSchema, bigquery.BigQueryTableIamUpdaterProducer, bigquery.BigQueryTableIdParseFunc),
			"google_bigquery_analytics_hub_data_exchange":                    bigqueryanalyticshub.ResourceBigqueryAnalyticsHubDataExchange(),
			"google_bigquery_analytics_hub_data_exchange_iam_binding":        tpgiamresource.ResourceIamBinding(bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIamSchema, bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIamUpdaterProducer, bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIdParseFunc),
			"google_bigquery_analytics_hub_data_exchange_iam_member":         tpgiamresource.ResourceIamMember(bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIamSchema, bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIamUpdaterProducer, bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIdParseFunc),
			"google_bigquery_analytics_hub_data_exchange_iam_policy":         tpgiamresource.ResourceIamPolicy(bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIamSchema, bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIamUpdaterProducer, bigqueryanalyticshub.BigqueryAnalyticsHubDataExchangeIdParseFunc),
			"google_bigquery_analytics_hub_listing":                          bigqueryanalyticshub.ResourceBigqueryAnalyticsHubListing(),
			"google_bigquery_analytics_hub_listing_iam_binding":              tpgiamresource.ResourceIamBinding(bigqueryanalyticshub.BigqueryAnalyticsHubListingIamSchema, bigqueryanalyticshub.BigqueryAnalyticsHubListingIamUpdaterProducer, bigqueryanalyticshub.BigqueryAnalyticsHubListingIdParseFunc),
			"google_bigquery_analytics_hub_listing_iam_member":               tpgiamresource.ResourceIamMember(bigqueryanalyticshub.BigqueryAnalyticsHubListingIamSchema, bigqueryanalyticshub.BigqueryAnalyticsHubListingIamUpdaterProducer, bigqueryanalyticshub.BigqueryAnalyticsHubListingIdParseFunc),
			"google_bigquery_analytics_hub_listing_iam_policy":               tpgiamresource.ResourceIamPolicy(bigqueryanalyticshub.BigqueryAnalyticsHubListingIamSchema, bigqueryanalyticshub.BigqueryAnalyticsHubListingIamUpdaterProducer, bigqueryanalyticshub.BigqueryAnalyticsHubListingIdParseFunc),
			"google_bigquery_connection":                                     bigqueryconnection.ResourceBigqueryConnectionConnection(),
			"google_bigquery_connection_iam_binding":                         tpgiamresource.ResourceIamBinding(bigqueryconnection.BigqueryConnectionConnectionIamSchema, bigqueryconnection.BigqueryConnectionConnectionIamUpdaterProducer, bigqueryconnection.BigqueryConnectionConnectionIdParseFunc),
			"google_bigquery_connection_iam_member":                          tpgiamresource.ResourceIamMember(bigqueryconnection.BigqueryConnectionConnectionIamSchema, bigqueryconnection.BigqueryConnectionConnectionIamUpdaterProducer, bigqueryconnection.BigqueryConnectionConnectionIdParseFunc),
			"google_bigquery_connection_iam_policy":                          tpgiamresource.ResourceIamPolicy(bigqueryconnection.BigqueryConnectionConnectionIamSchema, bigqueryconnection.BigqueryConnectionConnectionIamUpdaterProducer, bigqueryconnection.BigqueryConnectionConnectionIdParseFunc),
			"google_bigquery_datapolicy_data_policy":                         bigquerydatapolicy.ResourceBigqueryDatapolicyDataPolicy(),
			"google_bigquery_datapolicy_data_policy_iam_binding":             tpgiamresource.ResourceIamBinding(bigquerydatapolicy.BigqueryDatapolicyDataPolicyIamSchema, bigquerydatapolicy.BigqueryDatapolicyDataPolicyIamUpdaterProducer, bigquerydatapolicy.BigqueryDatapolicyDataPolicyIdParseFunc),
			"google_bigquery_datapolicy_data_policy_iam_member":              tpgiamresource.ResourceIamMember(bigquerydatapolicy.BigqueryDatapolicyDataPolicyIamSchema, bigquerydatapolicy.BigqueryDatapolicyDataPolicyIamUpdaterProducer, bigquerydatapolicy.BigqueryDatapolicyDataPolicyIdParseFunc),
			"google_bigquery_datapolicy_data_policy_iam_policy":              tpgiamresource.ResourceIamPolicy(bigquerydatapolicy.BigqueryDatapolicyDataPolicyIamSchema, bigquerydatapolicy.BigqueryDatapolicyDataPolicyIamUpdaterProducer, bigquerydatapolicy.BigqueryDatapolicyDataPolicyIdParseFunc),
			"google_bigquery_data_transfer_config":                           bigquerydatatransfer.ResourceBigqueryDataTransferConfig(),
			"google_bigquery_bi_reservation":                                 bigqueryreservation.ResourceBigqueryReservationBiReservation(),
			"google_bigquery_capacity_commitment":                            bigqueryreservation.ResourceBigqueryReservationCapacityCommitment(),
			"google_bigquery_reservation":                                    bigqueryreservation.ResourceBigqueryReservationReservation(),
			"google_bigtable_app_profile":                                    bigtable.ResourceBigtableAppProfile(),
			"google_billing_budget":                                          billing.ResourceBillingBudget(),
			"google_binary_authorization_attestor":                           binaryauthorization.ResourceBinaryAuthorizationAttestor(),
			"google_binary_authorization_attestor_iam_binding":               tpgiamresource.ResourceIamBinding(binaryauthorization.BinaryAuthorizationAttestorIamSchema, binaryauthorization.BinaryAuthorizationAttestorIamUpdaterProducer, binaryauthorization.BinaryAuthorizationAttestorIdParseFunc),
			"google_binary_authorization_attestor_iam_member":                tpgiamresource.ResourceIamMember(binaryauthorization.BinaryAuthorizationAttestorIamSchema, binaryauthorization.BinaryAuthorizationAttestorIamUpdaterProducer, binaryauthorization.BinaryAuthorizationAttestorIdParseFunc),
			"google_binary_authorization_attestor_iam_policy":                tpgiamresource.ResourceIamPolicy(binaryauthorization.BinaryAuthorizationAttestorIamSchema, binaryauthorization.BinaryAuthorizationAttestorIamUpdaterProducer, binaryauthorization.BinaryAuthorizationAttestorIdParseFunc),
			"google_binary_authorization_policy":                             binaryauthorization.ResourceBinaryAuthorizationPolicy(),
			"google_certificate_manager_certificate":                         certificatemanager.ResourceCertificateManagerCertificate(),
			"google_certificate_manager_certificate_issuance_config":         certificatemanager.ResourceCertificateManagerCertificateIssuanceConfig(),
			"google_certificate_manager_certificate_map":                     certificatemanager.ResourceCertificateManagerCertificateMap(),
			"google_certificate_manager_certificate_map_entry":               certificatemanager.ResourceCertificateManagerCertificateMapEntry(),
			"google_certificate_manager_dns_authorization":                   certificatemanager.ResourceCertificateManagerDnsAuthorization(),
			"google_certificate_manager_trust_config":                        certificatemanager.ResourceCertificateManagerTrustConfig(),
			"google_cloud_asset_folder_feed":                                 cloudasset.ResourceCloudAssetFolderFeed(),
			"google_cloud_asset_organization_feed":                           cloudasset.ResourceCloudAssetOrganizationFeed(),
			"google_cloud_asset_project_feed":                                cloudasset.ResourceCloudAssetProjectFeed(),
			"google_cloudbuild_bitbucket_server_config":                      cloudbuild.ResourceCloudBuildBitbucketServerConfig(),
			"google_cloudbuild_trigger":                                      cloudbuild.ResourceCloudBuildTrigger(),
			"google_cloudbuildv2_connection_iam_binding":                     tpgiamresource.ResourceIamBinding(cloudbuildv2.Cloudbuildv2ConnectionIamSchema, cloudbuildv2.Cloudbuildv2ConnectionIamUpdaterProducer, cloudbuildv2.Cloudbuildv2ConnectionIdParseFunc),
			"google_cloudbuildv2_connection_iam_member":                      tpgiamresource.ResourceIamMember(cloudbuildv2.Cloudbuildv2ConnectionIamSchema, cloudbuildv2.Cloudbuildv2ConnectionIamUpdaterProducer, cloudbuildv2.Cloudbuildv2ConnectionIdParseFunc),
			"google_cloudbuildv2_connection_iam_policy":                      tpgiamresource.ResourceIamPolicy(cloudbuildv2.Cloudbuildv2ConnectionIamSchema, cloudbuildv2.Cloudbuildv2ConnectionIamUpdaterProducer, cloudbuildv2.Cloudbuildv2ConnectionIdParseFunc),
			"google_cloudfunctions_function_iam_binding":                     tpgiamresource.ResourceIamBinding(cloudfunctions.CloudFunctionsCloudFunctionIamSchema, cloudfunctions.CloudFunctionsCloudFunctionIamUpdaterProducer, cloudfunctions.CloudFunctionsCloudFunctionIdParseFunc),
			"google_cloudfunctions_function_iam_member":                      tpgiamresource.ResourceIamMember(cloudfunctions.CloudFunctionsCloudFunctionIamSchema, cloudfunctions.CloudFunctionsCloudFunctionIamUpdaterProducer, cloudfunctions.CloudFunctionsCloudFunctionIdParseFunc),
			"google_cloudfunctions_function_iam_policy":                      tpgiamresource.ResourceIamPolicy(cloudfunctions.CloudFunctionsCloudFunctionIamSchema, cloudfunctions.CloudFunctionsCloudFunctionIamUpdaterProducer, cloudfunctions.CloudFunctionsCloudFunctionIdParseFunc),
			"google_cloudfunctions2_function":                                cloudfunctions2.ResourceCloudfunctions2function(),
			"google_cloudfunctions2_function_iam_binding":                    tpgiamresource.ResourceIamBinding(cloudfunctions2.Cloudfunctions2functionIamSchema, cloudfunctions2.Cloudfunctions2functionIamUpdaterProducer, cloudfunctions2.Cloudfunctions2functionIdParseFunc),
			"google_cloudfunctions2_function_iam_member":                     tpgiamresource.ResourceIamMember(cloudfunctions2.Cloudfunctions2functionIamSchema, cloudfunctions2.Cloudfunctions2functionIamUpdaterProducer, cloudfunctions2.Cloudfunctions2functionIdParseFunc),
			"google_cloudfunctions2_function_iam_policy":                     tpgiamresource.ResourceIamPolicy(cloudfunctions2.Cloudfunctions2functionIamSchema, cloudfunctions2.Cloudfunctions2functionIamUpdaterProducer, cloudfunctions2.Cloudfunctions2functionIdParseFunc),
			"google_cloud_identity_group":                                    cloudidentity.ResourceCloudIdentityGroup(),
			"google_cloud_identity_group_membership":                         cloudidentity.ResourceCloudIdentityGroupMembership(),
			"google_cloud_ids_endpoint":                                      cloudids.ResourceCloudIdsEndpoint(),
			"google_cloudiot_device":                                         cloudiot.ResourceCloudIotDevice(),
			"google_cloudiot_registry":                                       cloudiot.ResourceCloudIotDeviceRegistry(),
			"google_cloudiot_registry_iam_binding":                           tpgiamresource.ResourceIamBinding(cloudiot.CloudIotDeviceRegistryIamSchema, cloudiot.CloudIotDeviceRegistryIamUpdaterProducer, cloudiot.CloudIotDeviceRegistryIdParseFunc),
			"google_cloudiot_registry_iam_member":                            tpgiamresource.ResourceIamMember(cloudiot.CloudIotDeviceRegistryIamSchema, cloudiot.CloudIotDeviceRegistryIamUpdaterProducer, cloudiot.CloudIotDeviceRegistryIdParseFunc),
			"google_cloudiot_registry_iam_policy":                            tpgiamresource.ResourceIamPolicy(cloudiot.CloudIotDeviceRegistryIamSchema, cloudiot.CloudIotDeviceRegistryIamUpdaterProducer, cloudiot.CloudIotDeviceRegistryIdParseFunc),
			"google_cloud_run_domain_mapping":                                cloudrun.ResourceCloudRunDomainMapping(),
			"google_cloud_run_service":                                       cloudrun.ResourceCloudRunService(),
			"google_cloud_run_service_iam_binding":                           tpgiamresource.ResourceIamBinding(cloudrun.CloudRunServiceIamSchema, cloudrun.CloudRunServiceIamUpdaterProducer, cloudrun.CloudRunServiceIdParseFunc),
			"google_cloud_run_service_iam_member":                            tpgiamresource.ResourceIamMember(cloudrun.CloudRunServiceIamSchema, cloudrun.CloudRunServiceIamUpdaterProducer, cloudrun.CloudRunServiceIdParseFunc),
			"google_cloud_run_service_iam_policy":                            tpgiamresource.ResourceIamPolicy(cloudrun.CloudRunServiceIamSchema, cloudrun.CloudRunServiceIamUpdaterProducer, cloudrun.CloudRunServiceIdParseFunc),
			"google_cloud_run_v2_job":                                        cloudrunv2.ResourceCloudRunV2Job(),
			"google_cloud_run_v2_job_iam_binding":                            tpgiamresource.ResourceIamBinding(cloudrunv2.CloudRunV2JobIamSchema, cloudrunv2.CloudRunV2JobIamUpdaterProducer, cloudrunv2.CloudRunV2JobIdParseFunc),
			"google_cloud_run_v2_job_iam_member":                             tpgiamresource.ResourceIamMember(cloudrunv2.CloudRunV2JobIamSchema, cloudrunv2.CloudRunV2JobIamUpdaterProducer, cloudrunv2.CloudRunV2JobIdParseFunc),
			"google_cloud_run_v2_job_iam_policy":                             tpgiamresource.ResourceIamPolicy(cloudrunv2.CloudRunV2JobIamSchema, cloudrunv2.CloudRunV2JobIamUpdaterProducer, cloudrunv2.CloudRunV2JobIdParseFunc),
			"google_cloud_run_v2_service":                                    cloudrunv2.ResourceCloudRunV2Service(),
			"google_cloud_run_v2_service_iam_binding":                        tpgiamresource.ResourceIamBinding(cloudrunv2.CloudRunV2ServiceIamSchema, cloudrunv2.CloudRunV2ServiceIamUpdaterProducer, cloudrunv2.CloudRunV2ServiceIdParseFunc),
			"google_cloud_run_v2_service_iam_member":                         tpgiamresource.ResourceIamMember(cloudrunv2.CloudRunV2ServiceIamSchema, cloudrunv2.CloudRunV2ServiceIamUpdaterProducer, cloudrunv2.CloudRunV2ServiceIdParseFunc),
			"google_cloud_run_v2_service_iam_policy":                         tpgiamresource.ResourceIamPolicy(cloudrunv2.CloudRunV2ServiceIamSchema, cloudrunv2.CloudRunV2ServiceIamUpdaterProducer, cloudrunv2.CloudRunV2ServiceIdParseFunc),
			"google_cloud_scheduler_job":                                     cloudscheduler.ResourceCloudSchedulerJob(),
			"google_cloud_tasks_queue":                                       cloudtasks.ResourceCloudTasksQueue(),
			"google_cloud_tasks_queue_iam_binding":                           tpgiamresource.ResourceIamBinding(cloudtasks.CloudTasksQueueIamSchema, cloudtasks.CloudTasksQueueIamUpdaterProducer, cloudtasks.CloudTasksQueueIdParseFunc),
			"google_cloud_tasks_queue_iam_member":                            tpgiamresource.ResourceIamMember(cloudtasks.CloudTasksQueueIamSchema, cloudtasks.CloudTasksQueueIamUpdaterProducer, cloudtasks.CloudTasksQueueIdParseFunc),
			"google_cloud_tasks_queue_iam_policy":                            tpgiamresource.ResourceIamPolicy(cloudtasks.CloudTasksQueueIamSchema, cloudtasks.CloudTasksQueueIamUpdaterProducer, cloudtasks.CloudTasksQueueIdParseFunc),
			"google_compute_address":                                         compute.ResourceComputeAddress(),
			"google_compute_autoscaler":                                      compute.ResourceComputeAutoscaler(),
			"google_compute_backend_bucket":                                  compute.ResourceComputeBackendBucket(),
			"google_compute_backend_bucket_iam_binding":                      tpgiamresource.ResourceIamBinding(compute.ComputeBackendBucketIamSchema, compute.ComputeBackendBucketIamUpdaterProducer, compute.ComputeBackendBucketIdParseFunc),
			"google_compute_backend_bucket_iam_member":                       tpgiamresource.ResourceIamMember(compute.ComputeBackendBucketIamSchema, compute.ComputeBackendBucketIamUpdaterProducer, compute.ComputeBackendBucketIdParseFunc),
			"google_compute_backend_bucket_iam_policy":                       tpgiamresource.ResourceIamPolicy(compute.ComputeBackendBucketIamSchema, compute.ComputeBackendBucketIamUpdaterProducer, compute.ComputeBackendBucketIdParseFunc),
			"google_compute_backend_bucket_signed_url_key":                   compute.ResourceComputeBackendBucketSignedUrlKey(),
			"google_compute_backend_service":                                 compute.ResourceComputeBackendService(),
			"google_compute_backend_service_iam_binding":                     tpgiamresource.ResourceIamBinding(compute.ComputeBackendServiceIamSchema, compute.ComputeBackendServiceIamUpdaterProducer, compute.ComputeBackendServiceIdParseFunc),
			"google_compute_backend_service_iam_member":                      tpgiamresource.ResourceIamMember(compute.ComputeBackendServiceIamSchema, compute.ComputeBackendServiceIamUpdaterProducer, compute.ComputeBackendServiceIdParseFunc),
			"google_compute_backend_service_iam_policy":                      tpgiamresource.ResourceIamPolicy(compute.ComputeBackendServiceIamSchema, compute.ComputeBackendServiceIamUpdaterProducer, compute.ComputeBackendServiceIdParseFunc),
			"google_compute_backend_service_signed_url_key":                  compute.ResourceComputeBackendServiceSignedUrlKey(),
			"google_compute_disk":                                            compute.ResourceComputeDisk(),
			"google_compute_disk_iam_binding":                                tpgiamresource.ResourceIamBinding(compute.ComputeDiskIamSchema, compute.ComputeDiskIamUpdaterProducer, compute.ComputeDiskIdParseFunc),
			"google_compute_disk_iam_member":                                 tpgiamresource.ResourceIamMember(compute.ComputeDiskIamSchema, compute.ComputeDiskIamUpdaterProducer, compute.ComputeDiskIdParseFunc),
			"google_compute_disk_iam_policy":                                 tpgiamresource.ResourceIamPolicy(compute.ComputeDiskIamSchema, compute.ComputeDiskIamUpdaterProducer, compute.ComputeDiskIdParseFunc),
			"google_compute_disk_resource_policy_attachment":                 compute.ResourceComputeDiskResourcePolicyAttachment(),
			"google_compute_external_vpn_gateway":                            compute.ResourceComputeExternalVpnGateway(),
			"google_compute_firewall":                                        compute.ResourceComputeFirewall(),
			"google_compute_forwarding_rule":                                 compute.ResourceComputeForwardingRule(),
			"google_compute_global_address":                                  compute.ResourceComputeGlobalAddress(),
			"google_compute_global_forwarding_rule":                          compute.ResourceComputeGlobalForwardingRule(),
			"google_compute_global_network_endpoint":                         compute.ResourceComputeGlobalNetworkEndpoint(),
			"google_compute_global_network_endpoint_group":                   compute.ResourceComputeGlobalNetworkEndpointGroup(),
			"google_compute_ha_vpn_gateway":                                  compute.ResourceComputeHaVpnGateway(),
			"google_compute_health_check":                                    compute.ResourceComputeHealthCheck(),
			"google_compute_http_health_check":                               compute.ResourceComputeHttpHealthCheck(),
			"google_compute_https_health_check":                              compute.ResourceComputeHttpsHealthCheck(),
			"google_compute_image":                                           compute.ResourceComputeImage(),
			"google_compute_image_iam_binding":                               tpgiamresource.ResourceIamBinding(compute.ComputeImageIamSchema, compute.ComputeImageIamUpdaterProducer, compute.ComputeImageIdParseFunc),
			"google_compute_image_iam_member":                                tpgiamresource.ResourceIamMember(compute.ComputeImageIamSchema, compute.ComputeImageIamUpdaterProducer, compute.ComputeImageIdParseFunc),
			"google_compute_image_iam_policy":                                tpgiamresource.ResourceIamPolicy(compute.ComputeImageIamSchema, compute.ComputeImageIamUpdaterProducer, compute.ComputeImageIdParseFunc),
			"google_compute_instance_iam_binding":                            tpgiamresource.ResourceIamBinding(compute.ComputeInstanceIamSchema, compute.ComputeInstanceIamUpdaterProducer, compute.ComputeInstanceIdParseFunc),
			"google_compute_instance_iam_member":                             tpgiamresource.ResourceIamMember(compute.ComputeInstanceIamSchema, compute.ComputeInstanceIamUpdaterProducer, compute.ComputeInstanceIdParseFunc),
			"google_compute_instance_iam_policy":                             tpgiamresource.ResourceIamPolicy(compute.ComputeInstanceIamSchema, compute.ComputeInstanceIamUpdaterProducer, compute.ComputeInstanceIdParseFunc),
			"google_compute_instance_group_named_port":                       compute.ResourceComputeInstanceGroupNamedPort(),
			"google_compute_interconnect_attachment":                         compute.ResourceComputeInterconnectAttachment(),
			"google_compute_machine_image":                                   compute.ResourceComputeMachineImage(),
			"google_compute_machine_image_iam_binding":                       tpgiamresource.ResourceIamBinding(compute.ComputeMachineImageIamSchema, compute.ComputeMachineImageIamUpdaterProducer, compute.ComputeMachineImageIdParseFunc),
			"google_compute_machine_image_iam_member":                        tpgiamresource.ResourceIamMember(compute.ComputeMachineImageIamSchema, compute.ComputeMachineImageIamUpdaterProducer, compute.ComputeMachineImageIdParseFunc),
			"google_compute_machine_image_iam_policy":                        tpgiamresource.ResourceIamPolicy(compute.ComputeMachineImageIamSchema, compute.ComputeMachineImageIamUpdaterProducer, compute.ComputeMachineImageIdParseFunc),
			"google_compute_managed_ssl_certificate":                         compute.ResourceComputeManagedSslCertificate(),
			"google_compute_network":                                         compute.ResourceComputeNetwork(),
			"google_compute_network_attachment":                              compute.ResourceComputeNetworkAttachment(),
			"google_compute_network_edge_security_service":                   compute.ResourceComputeNetworkEdgeSecurityService(),
			"google_compute_network_endpoint":                                compute.ResourceComputeNetworkEndpoint(),
			"google_compute_network_endpoint_group":                          compute.ResourceComputeNetworkEndpointGroup(),
			"google_compute_network_endpoints":                               compute.ResourceComputeNetworkEndpoints(),
			"google_compute_network_peering_routes_config":                   compute.ResourceComputeNetworkPeeringRoutesConfig(),
			"google_compute_node_group":                                      compute.ResourceComputeNodeGroup(),
			"google_compute_node_template":                                   compute.ResourceComputeNodeTemplate(),
			"google_compute_organization_security_policy":                    compute.ResourceComputeOrganizationSecurityPolicy(),
			"google_compute_organization_security_policy_association":        compute.ResourceComputeOrganizationSecurityPolicyAssociation(),
			"google_compute_organization_security_policy_rule":               compute.ResourceComputeOrganizationSecurityPolicyRule(),
			"google_compute_packet_mirroring":                                compute.ResourceComputePacketMirroring(),
			"google_compute_per_instance_config":                             compute.ResourceComputePerInstanceConfig(),
			"google_compute_public_advertised_prefix":                        compute.ResourceComputePublicAdvertisedPrefix(),
			"google_compute_public_delegated_prefix":                         compute.ResourceComputePublicDelegatedPrefix(),
			"google_compute_region_autoscaler":                               compute.ResourceComputeRegionAutoscaler(),
			"google_compute_region_backend_service":                          compute.ResourceComputeRegionBackendService(),
			"google_compute_region_backend_service_iam_binding":              tpgiamresource.ResourceIamBinding(compute.ComputeRegionBackendServiceIamSchema, compute.ComputeRegionBackendServiceIamUpdaterProducer, compute.ComputeRegionBackendServiceIdParseFunc),
			"google_compute_region_backend_service_iam_member":               tpgiamresource.ResourceIamMember(compute.ComputeRegionBackendServiceIamSchema, compute.ComputeRegionBackendServiceIamUpdaterProducer, compute.ComputeRegionBackendServiceIdParseFunc),
			"google_compute_region_backend_service_iam_policy":               tpgiamresource.ResourceIamPolicy(compute.ComputeRegionBackendServiceIamSchema, compute.ComputeRegionBackendServiceIamUpdaterProducer, compute.ComputeRegionBackendServiceIdParseFunc),
			"google_compute_region_commitment":                               compute.ResourceComputeRegionCommitment(),
			"google_compute_region_disk":                                     compute.ResourceComputeRegionDisk(),
			"google_compute_region_disk_iam_binding":                         tpgiamresource.ResourceIamBinding(compute.ComputeRegionDiskIamSchema, compute.ComputeRegionDiskIamUpdaterProducer, compute.ComputeRegionDiskIdParseFunc),
			"google_compute_region_disk_iam_member":                          tpgiamresource.ResourceIamMember(compute.ComputeRegionDiskIamSchema, compute.ComputeRegionDiskIamUpdaterProducer, compute.ComputeRegionDiskIdParseFunc),
			"google_compute_region_disk_iam_policy":                          tpgiamresource.ResourceIamPolicy(compute.ComputeRegionDiskIamSchema, compute.ComputeRegionDiskIamUpdaterProducer, compute.ComputeRegionDiskIdParseFunc),
			"google_compute_region_disk_resource_policy_attachment":          compute.ResourceComputeRegionDiskResourcePolicyAttachment(),
			"google_compute_region_health_check":                             compute.ResourceComputeRegionHealthCheck(),
			"google_compute_region_network_endpoint_group":                   compute.ResourceComputeRegionNetworkEndpointGroup(),
			"google_compute_region_per_instance_config":                      compute.ResourceComputeRegionPerInstanceConfig(),
			"google_compute_region_security_policy":                          compute.ResourceComputeRegionSecurityPolicy(),
			"google_compute_region_security_policy_rule":                     compute.ResourceComputeRegionSecurityPolicyRule(),
			"google_compute_region_ssl_certificate":                          compute.ResourceComputeRegionSslCertificate(),
			"google_compute_region_ssl_policy":                               compute.ResourceComputeRegionSslPolicy(),
			"google_compute_region_target_http_proxy":                        compute.ResourceComputeRegionTargetHttpProxy(),
			"google_compute_region_target_https_proxy":                       compute.ResourceComputeRegionTargetHttpsProxy(),
			"google_compute_region_target_tcp_proxy":                         compute.ResourceComputeRegionTargetTcpProxy(),
			"google_compute_region_url_map":                                  compute.ResourceComputeRegionUrlMap(),
			"google_compute_reservation":                                     compute.ResourceComputeReservation(),
			"google_compute_resource_policy":                                 compute.ResourceComputeResourcePolicy(),
			"google_compute_route":                                           compute.ResourceComputeRoute(),
			"google_compute_router":                                          compute.ResourceComputeRouter(),
			"google_compute_router_peer":                                     compute.ResourceComputeRouterBgpPeer(),
			"google_compute_router_nat":                                      compute.ResourceComputeRouterNat(),
			"google_compute_service_attachment":                              compute.ResourceComputeServiceAttachment(),
			"google_compute_snapshot":                                        compute.ResourceComputeSnapshot(),
			"google_compute_snapshot_iam_binding":                            tpgiamresource.ResourceIamBinding(compute.ComputeSnapshotIamSchema, compute.ComputeSnapshotIamUpdaterProducer, compute.ComputeSnapshotIdParseFunc),
			"google_compute_snapshot_iam_member":                             tpgiamresource.ResourceIamMember(compute.ComputeSnapshotIamSchema, compute.ComputeSnapshotIamUpdaterProducer, compute.ComputeSnapshotIdParseFunc),
			"google_compute_snapshot_iam_policy":                             tpgiamresource.ResourceIamPolicy(compute.ComputeSnapshotIamSchema, compute.ComputeSnapshotIamUpdaterProducer, compute.ComputeSnapshotIdParseFunc),
			"google_compute_ssl_certificate":                                 compute.ResourceComputeSslCertificate(),
			"google_compute_ssl_policy":                                      compute.ResourceComputeSslPolicy(),
			"google_compute_subnetwork":                                      compute.ResourceComputeSubnetwork(),
			"google_compute_subnetwork_iam_binding":                          tpgiamresource.ResourceIamBinding(compute.ComputeSubnetworkIamSchema, compute.ComputeSubnetworkIamUpdaterProducer, compute.ComputeSubnetworkIdParseFunc),
			"google_compute_subnetwork_iam_member":                           tpgiamresource.ResourceIamMember(compute.ComputeSubnetworkIamSchema, compute.ComputeSubnetworkIamUpdaterProducer, compute.ComputeSubnetworkIdParseFunc),
			"google_compute_subnetwork_iam_policy":                           tpgiamresource.ResourceIamPolicy(compute.ComputeSubnetworkIamSchema, compute.ComputeSubnetworkIamUpdaterProducer, compute.ComputeSubnetworkIdParseFunc),
			"google_compute_target_grpc_proxy":                               compute.ResourceComputeTargetGrpcProxy(),
			"google_compute_target_http_proxy":                               compute.ResourceComputeTargetHttpProxy(),
			"google_compute_target_https_proxy":                              compute.ResourceComputeTargetHttpsProxy(),
			"google_compute_target_instance":                                 compute.ResourceComputeTargetInstance(),
			"google_compute_target_ssl_proxy":                                compute.ResourceComputeTargetSslProxy(),
			"google_compute_target_tcp_proxy":                                compute.ResourceComputeTargetTcpProxy(),
			"google_compute_url_map":                                         compute.ResourceComputeUrlMap(),
			"google_compute_vpn_gateway":                                     compute.ResourceComputeVpnGateway(),
			"google_compute_vpn_tunnel":                                      compute.ResourceComputeVpnTunnel(),
			"google_container_analysis_note":                                 containeranalysis.ResourceContainerAnalysisNote(),
			"google_container_analysis_note_iam_binding":                     tpgiamresource.ResourceIamBinding(containeranalysis.ContainerAnalysisNoteIamSchema, containeranalysis.ContainerAnalysisNoteIamUpdaterProducer, containeranalysis.ContainerAnalysisNoteIdParseFunc),
			"google_container_analysis_note_iam_member":                      tpgiamresource.ResourceIamMember(containeranalysis.ContainerAnalysisNoteIamSchema, containeranalysis.ContainerAnalysisNoteIamUpdaterProducer, containeranalysis.ContainerAnalysisNoteIdParseFunc),
			"google_container_analysis_note_iam_policy":                      tpgiamresource.ResourceIamPolicy(containeranalysis.ContainerAnalysisNoteIamSchema, containeranalysis.ContainerAnalysisNoteIamUpdaterProducer, containeranalysis.ContainerAnalysisNoteIdParseFunc),
			"google_container_analysis_occurrence":                           containeranalysis.ResourceContainerAnalysisOccurrence(),
			"google_container_attached_cluster":                              containerattached.ResourceContainerAttachedCluster(),
			"google_billing_project_info":                                    corebilling.ResourceCoreBillingProjectInfo(),
			"google_database_migration_service_connection_profile":           databasemigrationservice.ResourceDatabaseMigrationServiceConnectionProfile(),
			"google_data_catalog_entry":                                      datacatalog.ResourceDataCatalogEntry(),
			"google_data_catalog_entry_group":                                datacatalog.ResourceDataCatalogEntryGroup(),
			"google_data_catalog_entry_group_iam_binding":                    tpgiamresource.ResourceIamBinding(datacatalog.DataCatalogEntryGroupIamSchema, datacatalog.DataCatalogEntryGroupIamUpdaterProducer, datacatalog.DataCatalogEntryGroupIdParseFunc),
			"google_data_catalog_entry_group_iam_member":                     tpgiamresource.ResourceIamMember(datacatalog.DataCatalogEntryGroupIamSchema, datacatalog.DataCatalogEntryGroupIamUpdaterProducer, datacatalog.DataCatalogEntryGroupIdParseFunc),
			"google_data_catalog_entry_group_iam_policy":                     tpgiamresource.ResourceIamPolicy(datacatalog.DataCatalogEntryGroupIamSchema, datacatalog.DataCatalogEntryGroupIamUpdaterProducer, datacatalog.DataCatalogEntryGroupIdParseFunc),
			"google_data_catalog_policy_tag":                                 datacatalog.ResourceDataCatalogPolicyTag(),
			"google_data_catalog_policy_tag_iam_binding":                     tpgiamresource.ResourceIamBinding(datacatalog.DataCatalogPolicyTagIamSchema, datacatalog.DataCatalogPolicyTagIamUpdaterProducer, datacatalog.DataCatalogPolicyTagIdParseFunc),
			"google_data_catalog_policy_tag_iam_member":                      tpgiamresource.ResourceIamMember(datacatalog.DataCatalogPolicyTagIamSchema, datacatalog.DataCatalogPolicyTagIamUpdaterProducer, datacatalog.DataCatalogPolicyTagIdParseFunc),
			"google_data_catalog_policy_tag_iam_policy":                      tpgiamresource.ResourceIamPolicy(datacatalog.DataCatalogPolicyTagIamSchema, datacatalog.DataCatalogPolicyTagIamUpdaterProducer, datacatalog.DataCatalogPolicyTagIdParseFunc),
			"google_data_catalog_tag":                                        datacatalog.ResourceDataCatalogTag(),
			"google_data_catalog_tag_template":                               datacatalog.ResourceDataCatalogTagTemplate(),
			"google_data_catalog_tag_template_iam_binding":                   tpgiamresource.ResourceIamBinding(datacatalog.DataCatalogTagTemplateIamSchema, datacatalog.DataCatalogTagTemplateIamUpdaterProducer, datacatalog.DataCatalogTagTemplateIdParseFunc),
			"google_data_catalog_tag_template_iam_member":                    tpgiamresource.ResourceIamMember(datacatalog.DataCatalogTagTemplateIamSchema, datacatalog.DataCatalogTagTemplateIamUpdaterProducer, datacatalog.DataCatalogTagTemplateIdParseFunc),
			"google_data_catalog_tag_template_iam_policy":                    tpgiamresource.ResourceIamPolicy(datacatalog.DataCatalogTagTemplateIamSchema, datacatalog.DataCatalogTagTemplateIamUpdaterProducer, datacatalog.DataCatalogTagTemplateIdParseFunc),
			"google_data_catalog_taxonomy":                                   datacatalog.ResourceDataCatalogTaxonomy(),
			"google_data_catalog_taxonomy_iam_binding":                       tpgiamresource.ResourceIamBinding(datacatalog.DataCatalogTaxonomyIamSchema, datacatalog.DataCatalogTaxonomyIamUpdaterProducer, datacatalog.DataCatalogTaxonomyIdParseFunc),
			"google_data_catalog_taxonomy_iam_member":                        tpgiamresource.ResourceIamMember(datacatalog.DataCatalogTaxonomyIamSchema, datacatalog.DataCatalogTaxonomyIamUpdaterProducer, datacatalog.DataCatalogTaxonomyIdParseFunc),
			"google_data_catalog_taxonomy_iam_policy":                        tpgiamresource.ResourceIamPolicy(datacatalog.DataCatalogTaxonomyIamSchema, datacatalog.DataCatalogTaxonomyIamUpdaterProducer, datacatalog.DataCatalogTaxonomyIdParseFunc),
			"google_dataform_repository":                                     dataform.ResourceDataformRepository(),
			"google_dataform_repository_release_config":                      dataform.ResourceDataformRepositoryReleaseConfig(),
			"google_dataform_repository_workflow_config":                     dataform.ResourceDataformRepositoryWorkflowConfig(),
			"google_data_fusion_instance":                                    datafusion.ResourceDataFusionInstance(),
			"google_data_fusion_instance_iam_binding":                        tpgiamresource.ResourceIamBinding(datafusion.DataFusionInstanceIamSchema, datafusion.DataFusionInstanceIamUpdaterProducer, datafusion.DataFusionInstanceIdParseFunc),
			"google_data_fusion_instance_iam_member":                         tpgiamresource.ResourceIamMember(datafusion.DataFusionInstanceIamSchema, datafusion.DataFusionInstanceIamUpdaterProducer, datafusion.DataFusionInstanceIdParseFunc),
			"google_data_fusion_instance_iam_policy":                         tpgiamresource.ResourceIamPolicy(datafusion.DataFusionInstanceIamSchema, datafusion.DataFusionInstanceIamUpdaterProducer, datafusion.DataFusionInstanceIdParseFunc),
			"google_data_loss_prevention_deidentify_template":                datalossprevention.ResourceDataLossPreventionDeidentifyTemplate(),
			"google_data_loss_prevention_inspect_template":                   datalossprevention.ResourceDataLossPreventionInspectTemplate(),
			"google_data_loss_prevention_job_trigger":                        datalossprevention.ResourceDataLossPreventionJobTrigger(),
			"google_data_loss_prevention_stored_info_type":                   datalossprevention.ResourceDataLossPreventionStoredInfoType(),
			"google_dataplex_asset_iam_binding":                              tpgiamresource.ResourceIamBinding(dataplex.DataplexAssetIamSchema, dataplex.DataplexAssetIamUpdaterProducer, dataplex.DataplexAssetIdParseFunc),
			"google_dataplex_asset_iam_member":                               tpgiamresource.ResourceIamMember(dataplex.DataplexAssetIamSchema, dataplex.DataplexAssetIamUpdaterProducer, dataplex.DataplexAssetIdParseFunc),
			"google_dataplex_asset_iam_policy":                               tpgiamresource.ResourceIamPolicy(dataplex.DataplexAssetIamSchema, dataplex.DataplexAssetIamUpdaterProducer, dataplex.DataplexAssetIdParseFunc),
			"google_dataplex_datascan":                                       dataplex.ResourceDataplexDatascan(),
			"google_dataplex_datascan_iam_binding":                           tpgiamresource.ResourceIamBinding(dataplex.DataplexDatascanIamSchema, dataplex.DataplexDatascanIamUpdaterProducer, dataplex.DataplexDatascanIdParseFunc),
			"google_dataplex_datascan_iam_member":                            tpgiamresource.ResourceIamMember(dataplex.DataplexDatascanIamSchema, dataplex.DataplexDatascanIamUpdaterProducer, dataplex.DataplexDatascanIdParseFunc),
			"google_dataplex_datascan_iam_policy":                            tpgiamresource.ResourceIamPolicy(dataplex.DataplexDatascanIamSchema, dataplex.DataplexDatascanIamUpdaterProducer, dataplex.DataplexDatascanIdParseFunc),
			"google_dataplex_lake_iam_binding":                               tpgiamresource.ResourceIamBinding(dataplex.DataplexLakeIamSchema, dataplex.DataplexLakeIamUpdaterProducer, dataplex.DataplexLakeIdParseFunc),
			"google_dataplex_lake_iam_member":                                tpgiamresource.ResourceIamMember(dataplex.DataplexLakeIamSchema, dataplex.DataplexLakeIamUpdaterProducer, dataplex.DataplexLakeIdParseFunc),
			"google_dataplex_lake_iam_policy":                                tpgiamresource.ResourceIamPolicy(dataplex.DataplexLakeIamSchema, dataplex.DataplexLakeIamUpdaterProducer, dataplex.DataplexLakeIdParseFunc),
			"google_dataplex_task":                                           dataplex.ResourceDataplexTask(),
			"google_dataplex_task_iam_binding":                               tpgiamresource.ResourceIamBinding(dataplex.DataplexTaskIamSchema, dataplex.DataplexTaskIamUpdaterProducer, dataplex.DataplexTaskIdParseFunc),
			"google_dataplex_task_iam_member":                                tpgiamresource.ResourceIamMember(dataplex.DataplexTaskIamSchema, dataplex.DataplexTaskIamUpdaterProducer, dataplex.DataplexTaskIdParseFunc),
			"google_dataplex_task_iam_policy":                                tpgiamresource.ResourceIamPolicy(dataplex.DataplexTaskIamSchema, dataplex.DataplexTaskIamUpdaterProducer, dataplex.DataplexTaskIdParseFunc),
			"google_dataplex_zone_iam_binding":                               tpgiamresource.ResourceIamBinding(dataplex.DataplexZoneIamSchema, dataplex.DataplexZoneIamUpdaterProducer, dataplex.DataplexZoneIdParseFunc),
			"google_dataplex_zone_iam_member":                                tpgiamresource.ResourceIamMember(dataplex.DataplexZoneIamSchema, dataplex.DataplexZoneIamUpdaterProducer, dataplex.DataplexZoneIdParseFunc),
			"google_dataplex_zone_iam_policy":                                tpgiamresource.ResourceIamPolicy(dataplex.DataplexZoneIamSchema, dataplex.DataplexZoneIamUpdaterProducer, dataplex.DataplexZoneIdParseFunc),
			"google_dataproc_autoscaling_policy":                             dataproc.ResourceDataprocAutoscalingPolicy(),
			"google_dataproc_autoscaling_policy_iam_binding":                 tpgiamresource.ResourceIamBinding(dataproc.DataprocAutoscalingPolicyIamSchema, dataproc.DataprocAutoscalingPolicyIamUpdaterProducer, dataproc.DataprocAutoscalingPolicyIdParseFunc),
			"google_dataproc_autoscaling_policy_iam_member":                  tpgiamresource.ResourceIamMember(dataproc.DataprocAutoscalingPolicyIamSchema, dataproc.DataprocAutoscalingPolicyIamUpdaterProducer, dataproc.DataprocAutoscalingPolicyIdParseFunc),
			"google_dataproc_autoscaling_policy_iam_policy":                  tpgiamresource.ResourceIamPolicy(dataproc.DataprocAutoscalingPolicyIamSchema, dataproc.DataprocAutoscalingPolicyIamUpdaterProducer, dataproc.DataprocAutoscalingPolicyIdParseFunc),
			"google_dataproc_metastore_federation":                           dataprocmetastore.ResourceDataprocMetastoreFederation(),
			"google_dataproc_metastore_federation_iam_binding":               tpgiamresource.ResourceIamBinding(dataprocmetastore.DataprocMetastoreFederationIamSchema, dataprocmetastore.DataprocMetastoreFederationIamUpdaterProducer, dataprocmetastore.DataprocMetastoreFederationIdParseFunc),
			"google_dataproc_metastore_federation_iam_member":                tpgiamresource.ResourceIamMember(dataprocmetastore.DataprocMetastoreFederationIamSchema, dataprocmetastore.DataprocMetastoreFederationIamUpdaterProducer, dataprocmetastore.DataprocMetastoreFederationIdParseFunc),
			"google_dataproc_metastore_federation_iam_policy":                tpgiamresource.ResourceIamPolicy(dataprocmetastore.DataprocMetastoreFederationIamSchema, dataprocmetastore.DataprocMetastoreFederationIamUpdaterProducer, dataprocmetastore.DataprocMetastoreFederationIdParseFunc),
			"google_dataproc_metastore_service":                              dataprocmetastore.ResourceDataprocMetastoreService(),
			"google_dataproc_metastore_service_iam_binding":                  tpgiamresource.ResourceIamBinding(dataprocmetastore.DataprocMetastoreServiceIamSchema, dataprocmetastore.DataprocMetastoreServiceIamUpdaterProducer, dataprocmetastore.DataprocMetastoreServiceIdParseFunc),
			"google_dataproc_metastore_service_iam_member":                   tpgiamresource.ResourceIamMember(dataprocmetastore.DataprocMetastoreServiceIamSchema, dataprocmetastore.DataprocMetastoreServiceIamUpdaterProducer, dataprocmetastore.DataprocMetastoreServiceIdParseFunc),
			"google_dataproc_metastore_service_iam_policy":                   tpgiamresource.ResourceIamPolicy(dataprocmetastore.DataprocMetastoreServiceIamSchema, dataprocmetastore.DataprocMetastoreServiceIamUpdaterProducer, dataprocmetastore.DataprocMetastoreServiceIdParseFunc),
			"google_datastore_index":                                         datastore.ResourceDatastoreIndex(),
			"google_datastream_connection_profile":                           datastream.ResourceDatastreamConnectionProfile(),
			"google_datastream_private_connection":                           datastream.ResourceDatastreamPrivateConnection(),
			"google_datastream_stream":                                       datastream.ResourceDatastreamStream(),
			"google_deployment_manager_deployment":                           deploymentmanager.ResourceDeploymentManagerDeployment(),
			"google_dialogflow_agent":                                        dialogflow.ResourceDialogflowAgent(),
			"google_dialogflow_entity_type":                                  dialogflow.ResourceDialogflowEntityType(),
			"google_dialogflow_fulfillment":                                  dialogflow.ResourceDialogflowFulfillment(),
			"google_dialogflow_intent":                                       dialogflow.ResourceDialogflowIntent(),
			"google_dialogflow_cx_agent":                                     dialogflowcx.ResourceDialogflowCXAgent(),
			"google_dialogflow_cx_entity_type":                               dialogflowcx.ResourceDialogflowCXEntityType(),
			"google_dialogflow_cx_flow":                                      dialogflowcx.ResourceDialogflowCXFlow(),
			"google_dialogflow_cx_intent":                                    dialogflowcx.ResourceDialogflowCXIntent(),
			"google_dialogflow_cx_page":                                      dialogflowcx.ResourceDialogflowCXPage(),
			"google_dialogflow_cx_webhook":                                   dialogflowcx.ResourceDialogflowCXWebhook(),
			"google_dns_managed_zone":                                        dns.ResourceDNSManagedZone(),
			"google_dns_managed_zone_iam_binding":                            tpgiamresource.ResourceIamBinding(dns.DNSManagedZoneIamSchema, dns.DNSManagedZoneIamUpdaterProducer, dns.DNSManagedZoneIdParseFunc),
			"google_dns_managed_zone_iam_member":                             tpgiamresource.ResourceIamMember(dns.DNSManagedZoneIamSchema, dns.DNSManagedZoneIamUpdaterProducer, dns.DNSManagedZoneIdParseFunc),
			"google_dns_managed_zone_iam_policy":                             tpgiamresource.ResourceIamPolicy(dns.DNSManagedZoneIamSchema, dns.DNSManagedZoneIamUpdaterProducer, dns.DNSManagedZoneIdParseFunc),
			"google_dns_policy":                                              dns.ResourceDNSPolicy(),
			"google_dns_response_policy":                                     dns.ResourceDNSResponsePolicy(),
			"google_dns_response_policy_rule":                                dns.ResourceDNSResponsePolicyRule(),
			"google_document_ai_processor":                                   documentai.ResourceDocumentAIProcessor(),
			"google_document_ai_processor_default_version":                   documentai.ResourceDocumentAIProcessorDefaultVersion(),
			"google_document_ai_warehouse_document_schema":                   documentaiwarehouse.ResourceDocumentAIWarehouseDocumentSchema(),
			"google_document_ai_warehouse_location":                          documentaiwarehouse.ResourceDocumentAIWarehouseLocation(),
			"google_essential_contacts_contact":                              essentialcontacts.ResourceEssentialContactsContact(),
			"google_filestore_backup":                                        filestore.ResourceFilestoreBackup(),
			"google_filestore_instance":                                      filestore.ResourceFilestoreInstance(),
			"google_filestore_snapshot":                                      filestore.ResourceFilestoreSnapshot(),
			"google_firebase_android_app":                                    firebase.ResourceFirebaseAndroidApp(),
			"google_firebase_apple_app":                                      firebase.ResourceFirebaseAppleApp(),
			"google_firebase_project":                                        firebase.ResourceFirebaseProject(),
			"google_firebase_project_location":                               firebase.ResourceFirebaseProjectLocation(),
			"google_firebase_web_app":                                        firebase.ResourceFirebaseWebApp(),
			"google_firebase_database_instance":                              firebasedatabase.ResourceFirebaseDatabaseInstance(),
			"google_firebase_extensions_instance":                            firebaseextensions.ResourceFirebaseExtensionsInstance(),
			"google_firebase_hosting_channel":                                firebasehosting.ResourceFirebaseHostingChannel(),
			"google_firebase_hosting_release":                                firebasehosting.ResourceFirebaseHostingRelease(),
			"google_firebase_hosting_site":                                   firebasehosting.ResourceFirebaseHostingSite(),
			"google_firebase_hosting_version":                                firebasehosting.ResourceFirebaseHostingVersion(),
			"google_firebase_storage_bucket":                                 firebasestorage.ResourceFirebaseStorageBucket(),
			"google_firestore_database":                                      firestore.ResourceFirestoreDatabase(),
			"google_firestore_document":                                      firestore.ResourceFirestoreDocument(),
			"google_firestore_field":                                         firestore.ResourceFirestoreField(),
			"google_firestore_index":                                         firestore.ResourceFirestoreIndex(),
			"google_game_services_game_server_cluster":                       gameservices.ResourceGameServicesGameServerCluster(),
			"google_game_services_game_server_config":                        gameservices.ResourceGameServicesGameServerConfig(),
			"google_game_services_game_server_deployment":                    gameservices.ResourceGameServicesGameServerDeployment(),
			"google_game_services_game_server_deployment_rollout":            gameservices.ResourceGameServicesGameServerDeploymentRollout(),
			"google_game_services_realm":                                     gameservices.ResourceGameServicesRealm(),
			"google_gke_backup_backup_plan":                                  gkebackup.ResourceGKEBackupBackupPlan(),
			"google_gke_backup_backup_plan_iam_binding":                      tpgiamresource.ResourceIamBinding(gkebackup.GKEBackupBackupPlanIamSchema, gkebackup.GKEBackupBackupPlanIamUpdaterProducer, gkebackup.GKEBackupBackupPlanIdParseFunc),
			"google_gke_backup_backup_plan_iam_member":                       tpgiamresource.ResourceIamMember(gkebackup.GKEBackupBackupPlanIamSchema, gkebackup.GKEBackupBackupPlanIamUpdaterProducer, gkebackup.GKEBackupBackupPlanIdParseFunc),
			"google_gke_backup_backup_plan_iam_policy":                       tpgiamresource.ResourceIamPolicy(gkebackup.GKEBackupBackupPlanIamSchema, gkebackup.GKEBackupBackupPlanIamUpdaterProducer, gkebackup.GKEBackupBackupPlanIdParseFunc),
			"google_gke_hub_membership":                                      gkehub.ResourceGKEHubMembership(),
			"google_gke_hub_membership_iam_binding":                          tpgiamresource.ResourceIamBinding(gkehub.GKEHubMembershipIamSchema, gkehub.GKEHubMembershipIamUpdaterProducer, gkehub.GKEHubMembershipIdParseFunc),
			"google_gke_hub_membership_iam_member":                           tpgiamresource.ResourceIamMember(gkehub.GKEHubMembershipIamSchema, gkehub.GKEHubMembershipIamUpdaterProducer, gkehub.GKEHubMembershipIdParseFunc),
			"google_gke_hub_membership_iam_policy":                           tpgiamresource.ResourceIamPolicy(gkehub.GKEHubMembershipIamSchema, gkehub.GKEHubMembershipIamUpdaterProducer, gkehub.GKEHubMembershipIdParseFunc),
			"google_gke_hub_feature":                                         gkehub2.ResourceGKEHub2Feature(),
			"google_gke_hub_feature_iam_binding":                             tpgiamresource.ResourceIamBinding(gkehub2.GKEHub2FeatureIamSchema, gkehub2.GKEHub2FeatureIamUpdaterProducer, gkehub2.GKEHub2FeatureIdParseFunc),
			"google_gke_hub_feature_iam_member":                              tpgiamresource.ResourceIamMember(gkehub2.GKEHub2FeatureIamSchema, gkehub2.GKEHub2FeatureIamUpdaterProducer, gkehub2.GKEHub2FeatureIdParseFunc),
			"google_gke_hub_feature_iam_policy":                              tpgiamresource.ResourceIamPolicy(gkehub2.GKEHub2FeatureIamSchema, gkehub2.GKEHub2FeatureIamUpdaterProducer, gkehub2.GKEHub2FeatureIdParseFunc),
			"google_gke_hub_membership_binding":                              gkehub2.ResourceGKEHub2MembershipBinding(),
			"google_gke_hub_membership_rbac_role_binding":                    gkehub2.ResourceGKEHub2MembershipRBACRoleBinding(),
			"google_gke_hub_namespace":                                       gkehub2.ResourceGKEHub2Namespace(),
			"google_gke_hub_scope":                                           gkehub2.ResourceGKEHub2Scope(),
			"google_gke_hub_scope_iam_binding":                               tpgiamresource.ResourceIamBinding(gkehub2.GKEHub2ScopeIamSchema, gkehub2.GKEHub2ScopeIamUpdaterProducer, gkehub2.GKEHub2ScopeIdParseFunc),
			"google_gke_hub_scope_iam_member":                                tpgiamresource.ResourceIamMember(gkehub2.GKEHub2ScopeIamSchema, gkehub2.GKEHub2ScopeIamUpdaterProducer, gkehub2.GKEHub2ScopeIdParseFunc),
			"google_gke_hub_scope_iam_policy":                                tpgiamresource.ResourceIamPolicy(gkehub2.GKEHub2ScopeIamSchema, gkehub2.GKEHub2ScopeIamUpdaterProducer, gkehub2.GKEHub2ScopeIdParseFunc),
			"google_gke_hub_scope_rbac_role_binding":                         gkehub2.ResourceGKEHub2ScopeRBACRoleBinding(),
			"google_gkeonprem_bare_metal_admin_cluster":                      gkeonprem.ResourceGkeonpremBareMetalAdminCluster(),
			"google_gkeonprem_bare_metal_cluster":                            gkeonprem.ResourceGkeonpremBareMetalCluster(),
			"google_gkeonprem_bare_metal_node_pool":                          gkeonprem.ResourceGkeonpremBareMetalNodePool(),
			"google_gkeonprem_vmware_cluster":                                gkeonprem.ResourceGkeonpremVmwareCluster(),
			"google_gkeonprem_vmware_node_pool":                              gkeonprem.ResourceGkeonpremVmwareNodePool(),
			"google_healthcare_consent_store":                                healthcare.ResourceHealthcareConsentStore(),
			"google_healthcare_consent_store_iam_binding":                    tpgiamresource.ResourceIamBinding(healthcare.HealthcareConsentStoreIamSchema, healthcare.HealthcareConsentStoreIamUpdaterProducer, healthcare.HealthcareConsentStoreIdParseFunc),
			"google_healthcare_consent_store_iam_member":                     tpgiamresource.ResourceIamMember(healthcare.HealthcareConsentStoreIamSchema, healthcare.HealthcareConsentStoreIamUpdaterProducer, healthcare.HealthcareConsentStoreIdParseFunc),
			"google_healthcare_consent_store_iam_policy":                     tpgiamresource.ResourceIamPolicy(healthcare.HealthcareConsentStoreIamSchema, healthcare.HealthcareConsentStoreIamUpdaterProducer, healthcare.HealthcareConsentStoreIdParseFunc),
			"google_healthcare_dataset":                                      healthcare.ResourceHealthcareDataset(),
			"google_healthcare_dicom_store":                                  healthcare.ResourceHealthcareDicomStore(),
			"google_healthcare_fhir_store":                                   healthcare.ResourceHealthcareFhirStore(),
			"google_healthcare_hl7_v2_store":                                 healthcare.ResourceHealthcareHl7V2Store(),
			"google_iam_access_boundary_policy":                              iam2.ResourceIAM2AccessBoundaryPolicy(),
			"google_iam_deny_policy":                                         iam2.ResourceIAM2DenyPolicy(),
			"google_iam_workload_identity_pool":                              iambeta.ResourceIAMBetaWorkloadIdentityPool(),
			"google_iam_workload_identity_pool_provider":                     iambeta.ResourceIAMBetaWorkloadIdentityPoolProvider(),
			"google_iam_workforce_pool":                                      iamworkforcepool.ResourceIAMWorkforcePoolWorkforcePool(),
			"google_iam_workforce_pool_provider":                             iamworkforcepool.ResourceIAMWorkforcePoolWorkforcePoolProvider(),
			"google_iap_app_engine_service_iam_binding":                      tpgiamresource.ResourceIamBinding(iap.IapAppEngineServiceIamSchema, iap.IapAppEngineServiceIamUpdaterProducer, iap.IapAppEngineServiceIdParseFunc),
			"google_iap_app_engine_service_iam_member":                       tpgiamresource.ResourceIamMember(iap.IapAppEngineServiceIamSchema, iap.IapAppEngineServiceIamUpdaterProducer, iap.IapAppEngineServiceIdParseFunc),
			"google_iap_app_engine_service_iam_policy":                       tpgiamresource.ResourceIamPolicy(iap.IapAppEngineServiceIamSchema, iap.IapAppEngineServiceIamUpdaterProducer, iap.IapAppEngineServiceIdParseFunc),
			"google_iap_app_engine_version_iam_binding":                      tpgiamresource.ResourceIamBinding(iap.IapAppEngineVersionIamSchema, iap.IapAppEngineVersionIamUpdaterProducer, iap.IapAppEngineVersionIdParseFunc),
			"google_iap_app_engine_version_iam_member":                       tpgiamresource.ResourceIamMember(iap.IapAppEngineVersionIamSchema, iap.IapAppEngineVersionIamUpdaterProducer, iap.IapAppEngineVersionIdParseFunc),
			"google_iap_app_engine_version_iam_policy":                       tpgiamresource.ResourceIamPolicy(iap.IapAppEngineVersionIamSchema, iap.IapAppEngineVersionIamUpdaterProducer, iap.IapAppEngineVersionIdParseFunc),
			"google_iap_brand":                                               iap.ResourceIapBrand(),
			"google_iap_client":                                              iap.ResourceIapClient(),
			"google_iap_tunnel_iam_binding":                                  tpgiamresource.ResourceIamBinding(iap.IapTunnelIamSchema, iap.IapTunnelIamUpdaterProducer, iap.IapTunnelIdParseFunc),
			"google_iap_tunnel_iam_member":                                   tpgiamresource.ResourceIamMember(iap.IapTunnelIamSchema, iap.IapTunnelIamUpdaterProducer, iap.IapTunnelIdParseFunc),
			"google_iap_tunnel_iam_policy":                                   tpgiamresource.ResourceIamPolicy(iap.IapTunnelIamSchema, iap.IapTunnelIamUpdaterProducer, iap.IapTunnelIdParseFunc),
			"google_iap_tunnel_instance_iam_binding":                         tpgiamresource.ResourceIamBinding(iap.IapTunnelInstanceIamSchema, iap.IapTunnelInstanceIamUpdaterProducer, iap.IapTunnelInstanceIdParseFunc),
			"google_iap_tunnel_instance_iam_member":                          tpgiamresource.ResourceIamMember(iap.IapTunnelInstanceIamSchema, iap.IapTunnelInstanceIamUpdaterProducer, iap.IapTunnelInstanceIdParseFunc),
			"google_iap_tunnel_instance_iam_policy":                          tpgiamresource.ResourceIamPolicy(iap.IapTunnelInstanceIamSchema, iap.IapTunnelInstanceIamUpdaterProducer, iap.IapTunnelInstanceIdParseFunc),
			"google_iap_web_iam_binding":                                     tpgiamresource.ResourceIamBinding(iap.IapWebIamSchema, iap.IapWebIamUpdaterProducer, iap.IapWebIdParseFunc),
			"google_iap_web_iam_member":                                      tpgiamresource.ResourceIamMember(iap.IapWebIamSchema, iap.IapWebIamUpdaterProducer, iap.IapWebIdParseFunc),
			"google_iap_web_iam_policy":                                      tpgiamresource.ResourceIamPolicy(iap.IapWebIamSchema, iap.IapWebIamUpdaterProducer, iap.IapWebIdParseFunc),
			"google_iap_web_backend_service_iam_binding":                     tpgiamresource.ResourceIamBinding(iap.IapWebBackendServiceIamSchema, iap.IapWebBackendServiceIamUpdaterProducer, iap.IapWebBackendServiceIdParseFunc),
			"google_iap_web_backend_service_iam_member":                      tpgiamresource.ResourceIamMember(iap.IapWebBackendServiceIamSchema, iap.IapWebBackendServiceIamUpdaterProducer, iap.IapWebBackendServiceIdParseFunc),
			"google_iap_web_backend_service_iam_policy":                      tpgiamresource.ResourceIamPolicy(iap.IapWebBackendServiceIamSchema, iap.IapWebBackendServiceIamUpdaterProducer, iap.IapWebBackendServiceIdParseFunc),
			"google_iap_web_region_backend_service_iam_binding":              tpgiamresource.ResourceIamBinding(iap.IapWebRegionBackendServiceIamSchema, iap.IapWebRegionBackendServiceIamUpdaterProducer, iap.IapWebRegionBackendServiceIdParseFunc),
			"google_iap_web_region_backend_service_iam_member":               tpgiamresource.ResourceIamMember(iap.IapWebRegionBackendServiceIamSchema, iap.IapWebRegionBackendServiceIamUpdaterProducer, iap.IapWebRegionBackendServiceIdParseFunc),
			"google_iap_web_region_backend_service_iam_policy":               tpgiamresource.ResourceIamPolicy(iap.IapWebRegionBackendServiceIamSchema, iap.IapWebRegionBackendServiceIamUpdaterProducer, iap.IapWebRegionBackendServiceIdParseFunc),
			"google_iap_web_type_app_engine_iam_binding":                     tpgiamresource.ResourceIamBinding(iap.IapWebTypeAppEngineIamSchema, iap.IapWebTypeAppEngineIamUpdaterProducer, iap.IapWebTypeAppEngineIdParseFunc),
			"google_iap_web_type_app_engine_iam_member":                      tpgiamresource.ResourceIamMember(iap.IapWebTypeAppEngineIamSchema, iap.IapWebTypeAppEngineIamUpdaterProducer, iap.IapWebTypeAppEngineIdParseFunc),
			"google_iap_web_type_app_engine_iam_policy":                      tpgiamresource.ResourceIamPolicy(iap.IapWebTypeAppEngineIamSchema, iap.IapWebTypeAppEngineIamUpdaterProducer, iap.IapWebTypeAppEngineIdParseFunc),
			"google_iap_web_type_compute_iam_binding":                        tpgiamresource.ResourceIamBinding(iap.IapWebTypeComputeIamSchema, iap.IapWebTypeComputeIamUpdaterProducer, iap.IapWebTypeComputeIdParseFunc),
			"google_iap_web_type_compute_iam_member":                         tpgiamresource.ResourceIamMember(iap.IapWebTypeComputeIamSchema, iap.IapWebTypeComputeIamUpdaterProducer, iap.IapWebTypeComputeIdParseFunc),
			"google_iap_web_type_compute_iam_policy":                         tpgiamresource.ResourceIamPolicy(iap.IapWebTypeComputeIamSchema, iap.IapWebTypeComputeIamUpdaterProducer, iap.IapWebTypeComputeIdParseFunc),
			"google_identity_platform_config":                                identityplatform.ResourceIdentityPlatformConfig(),
			"google_identity_platform_default_supported_idp_config":          identityplatform.ResourceIdentityPlatformDefaultSupportedIdpConfig(),
			"google_identity_platform_inbound_saml_config":                   identityplatform.ResourceIdentityPlatformInboundSamlConfig(),
			"google_identity_platform_oauth_idp_config":                      identityplatform.ResourceIdentityPlatformOauthIdpConfig(),
			"google_identity_platform_project_default_config":                identityplatform.ResourceIdentityPlatformProjectDefaultConfig(),
			"google_identity_platform_tenant":                                identityplatform.ResourceIdentityPlatformTenant(),
			"google_identity_platform_tenant_default_supported_idp_config":   identityplatform.ResourceIdentityPlatformTenantDefaultSupportedIdpConfig(),
			"google_identity_platform_tenant_inbound_saml_config":            identityplatform.ResourceIdentityPlatformTenantInboundSamlConfig(),
			"google_identity_platform_tenant_oauth_idp_config":               identityplatform.ResourceIdentityPlatformTenantOauthIdpConfig(),
			"google_kms_crypto_key":                                          kms.ResourceKMSCryptoKey(),
			"google_kms_crypto_key_version":                                  kms.ResourceKMSCryptoKeyVersion(),
			"google_kms_key_ring":                                            kms.ResourceKMSKeyRing(),
			"google_kms_key_ring_import_job":                                 kms.ResourceKMSKeyRingImportJob(),
			"google_kms_secret_ciphertext":                                   kms.ResourceKMSSecretCiphertext(),
			"google_logging_linked_dataset":                                  logging.ResourceLoggingLinkedDataset(),
			"google_logging_log_view":                                        logging.ResourceLoggingLogView(),
			"google_logging_metric":                                          logging.ResourceLoggingMetric(),
			"google_looker_instance":                                         looker.ResourceLookerInstance(),
			"google_memcache_instance":                                       memcache.ResourceMemcacheInstance(),
			"google_ml_engine_model":                                         mlengine.ResourceMLEngineModel(),
			"google_monitoring_alert_policy":                                 monitoring.ResourceMonitoringAlertPolicy(),
			"google_monitoring_service":                                      monitoring.ResourceMonitoringGenericService(),
			"google_monitoring_group":                                        monitoring.ResourceMonitoringGroup(),
			"google_monitoring_metric_descriptor":                            monitoring.ResourceMonitoringMetricDescriptor(),
			"google_monitoring_monitored_project":                            monitoring.ResourceMonitoringMonitoredProject(),
			"google_monitoring_notification_channel":                         monitoring.ResourceMonitoringNotificationChannel(),
			"google_monitoring_custom_service":                               monitoring.ResourceMonitoringService(),
			"google_monitoring_slo":                                          monitoring.ResourceMonitoringSlo(),
			"google_monitoring_uptime_check_config":                          monitoring.ResourceMonitoringUptimeCheckConfig(),
			"google_network_connectivity_service_connection_policy":          networkconnectivity.ResourceNetworkConnectivityServiceConnectionPolicy(),
			"google_network_management_connectivity_test":                    networkmanagement.ResourceNetworkManagementConnectivityTest(),
			"google_network_security_address_group":                          networksecurity.ResourceNetworkSecurityAddressGroup(),
			"google_network_security_authorization_policy":                   networksecurity.ResourceNetworkSecurityAuthorizationPolicy(),
			"google_network_security_client_tls_policy":                      networksecurity.ResourceNetworkSecurityClientTlsPolicy(),
			"google_network_security_gateway_security_policy":                networksecurity.ResourceNetworkSecurityGatewaySecurityPolicy(),
			"google_network_security_gateway_security_policy_rule":           networksecurity.ResourceNetworkSecurityGatewaySecurityPolicyRule(),
			"google_network_security_server_tls_policy":                      networksecurity.ResourceNetworkSecurityServerTlsPolicy(),
			"google_network_security_tls_inspection_policy":                  networksecurity.ResourceNetworkSecurityTlsInspectionPolicy(),
			"google_network_security_url_lists":                              networksecurity.ResourceNetworkSecurityUrlLists(),
			"google_network_services_edge_cache_keyset":                      networkservices.ResourceNetworkServicesEdgeCacheKeyset(),
			"google_network_services_edge_cache_origin":                      networkservices.ResourceNetworkServicesEdgeCacheOrigin(),
			"google_network_services_edge_cache_service":                     networkservices.ResourceNetworkServicesEdgeCacheService(),
			"google_network_services_endpoint_policy":                        networkservices.ResourceNetworkServicesEndpointPolicy(),
			"google_network_services_gateway":                                networkservices.ResourceNetworkServicesGateway(),
			"google_network_services_grpc_route":                             networkservices.ResourceNetworkServicesGrpcRoute(),
			"google_network_services_http_route":                             networkservices.ResourceNetworkServicesHttpRoute(),
			"google_network_services_mesh":                                   networkservices.ResourceNetworkServicesMesh(),
			"google_network_services_service_binding":                        networkservices.ResourceNetworkServicesServiceBinding(),
			"google_network_services_tcp_route":                              networkservices.ResourceNetworkServicesTcpRoute(),
			"google_network_services_tls_route":                              networkservices.ResourceNetworkServicesTlsRoute(),
			"google_notebooks_environment":                                   notebooks.ResourceNotebooksEnvironment(),
			"google_notebooks_instance":                                      notebooks.ResourceNotebooksInstance(),
			"google_notebooks_instance_iam_binding":                          tpgiamresource.ResourceIamBinding(notebooks.NotebooksInstanceIamSchema, notebooks.NotebooksInstanceIamUpdaterProducer, notebooks.NotebooksInstanceIdParseFunc),
			"google_notebooks_instance_iam_member":                           tpgiamresource.ResourceIamMember(notebooks.NotebooksInstanceIamSchema, notebooks.NotebooksInstanceIamUpdaterProducer, notebooks.NotebooksInstanceIdParseFunc),
			"google_notebooks_instance_iam_policy":                           tpgiamresource.ResourceIamPolicy(notebooks.NotebooksInstanceIamSchema, notebooks.NotebooksInstanceIamUpdaterProducer, notebooks.NotebooksInstanceIdParseFunc),
			"google_notebooks_location":                                      notebooks.ResourceNotebooksLocation(),
			"google_notebooks_runtime":                                       notebooks.ResourceNotebooksRuntime(),
			"google_notebooks_runtime_iam_binding":                           tpgiamresource.ResourceIamBinding(notebooks.NotebooksRuntimeIamSchema, notebooks.NotebooksRuntimeIamUpdaterProducer, notebooks.NotebooksRuntimeIdParseFunc),
			"google_notebooks_runtime_iam_member":                            tpgiamresource.ResourceIamMember(notebooks.NotebooksRuntimeIamSchema, notebooks.NotebooksRuntimeIamUpdaterProducer, notebooks.NotebooksRuntimeIdParseFunc),
			"google_notebooks_runtime_iam_policy":                            tpgiamresource.ResourceIamPolicy(notebooks.NotebooksRuntimeIamSchema, notebooks.NotebooksRuntimeIamUpdaterProducer, notebooks.NotebooksRuntimeIdParseFunc),
			"google_org_policy_custom_constraint":                            orgpolicy.ResourceOrgPolicyCustomConstraint(),
			"google_os_config_guest_policies":                                osconfig.ResourceOSConfigGuestPolicies(),
			"google_os_config_patch_deployment":                              osconfig.ResourceOSConfigPatchDeployment(),
			"google_os_login_ssh_public_key":                                 oslogin.ResourceOSLoginSSHPublicKey(),
			"google_privateca_ca_pool":                                       privateca.ResourcePrivatecaCaPool(),
			"google_privateca_ca_pool_iam_binding":                           tpgiamresource.ResourceIamBinding(privateca.PrivatecaCaPoolIamSchema, privateca.PrivatecaCaPoolIamUpdaterProducer, privateca.PrivatecaCaPoolIdParseFunc),
			"google_privateca_ca_pool_iam_member":                            tpgiamresource.ResourceIamMember(privateca.PrivatecaCaPoolIamSchema, privateca.PrivatecaCaPoolIamUpdaterProducer, privateca.PrivatecaCaPoolIdParseFunc),
			"google_privateca_ca_pool_iam_policy":                            tpgiamresource.ResourceIamPolicy(privateca.PrivatecaCaPoolIamSchema, privateca.PrivatecaCaPoolIamUpdaterProducer, privateca.PrivatecaCaPoolIdParseFunc),
			"google_privateca_certificate":                                   privateca.ResourcePrivatecaCertificate(),
			"google_privateca_certificate_authority":                         privateca.ResourcePrivatecaCertificateAuthority(),
			"google_privateca_certificate_template_iam_binding":              tpgiamresource.ResourceIamBinding(privateca.PrivatecaCertificateTemplateIamSchema, privateca.PrivatecaCertificateTemplateIamUpdaterProducer, privateca.PrivatecaCertificateTemplateIdParseFunc),
			"google_privateca_certificate_template_iam_member":               tpgiamresource.ResourceIamMember(privateca.PrivatecaCertificateTemplateIamSchema, privateca.PrivatecaCertificateTemplateIamUpdaterProducer, privateca.PrivatecaCertificateTemplateIdParseFunc),
			"google_privateca_certificate_template_iam_policy":               tpgiamresource.ResourceIamPolicy(privateca.PrivatecaCertificateTemplateIamSchema, privateca.PrivatecaCertificateTemplateIamUpdaterProducer, privateca.PrivatecaCertificateTemplateIdParseFunc),
			"google_public_ca_external_account_key":                          publicca.ResourcePublicCAExternalAccountKey(),
			"google_pubsub_schema":                                           pubsub.ResourcePubsubSchema(),
			"google_pubsub_subscription":                                     pubsub.ResourcePubsubSubscription(),
			"google_pubsub_topic":                                            pubsub.ResourcePubsubTopic(),
			"google_pubsub_topic_iam_binding":                                tpgiamresource.ResourceIamBinding(pubsub.PubsubTopicIamSchema, pubsub.PubsubTopicIamUpdaterProducer, pubsub.PubsubTopicIdParseFunc),
			"google_pubsub_topic_iam_member":                                 tpgiamresource.ResourceIamMember(pubsub.PubsubTopicIamSchema, pubsub.PubsubTopicIamUpdaterProducer, pubsub.PubsubTopicIdParseFunc),
			"google_pubsub_topic_iam_policy":                                 tpgiamresource.ResourceIamPolicy(pubsub.PubsubTopicIamSchema, pubsub.PubsubTopicIamUpdaterProducer, pubsub.PubsubTopicIdParseFunc),
			"google_pubsub_lite_reservation":                                 pubsublite.ResourcePubsubLiteReservation(),
			"google_pubsub_lite_subscription":                                pubsublite.ResourcePubsubLiteSubscription(),
			"google_pubsub_lite_topic":                                       pubsublite.ResourcePubsubLiteTopic(),
			"google_redis_cluster":                                           redis.ResourceRedisCluster(),
			"google_redis_instance":                                          redis.ResourceRedisInstance(),
			"google_resource_manager_lien":                                   resourcemanager.ResourceResourceManagerLien(),
			"google_runtimeconfig_config_iam_binding":                        tpgiamresource.ResourceIamBinding(runtimeconfig.RuntimeConfigConfigIamSchema, runtimeconfig.RuntimeConfigConfigIamUpdaterProducer, runtimeconfig.RuntimeConfigConfigIdParseFunc),
			"google_runtimeconfig_config_iam_member":                         tpgiamresource.ResourceIamMember(runtimeconfig.RuntimeConfigConfigIamSchema, runtimeconfig.RuntimeConfigConfigIamUpdaterProducer, runtimeconfig.RuntimeConfigConfigIdParseFunc),
			"google_runtimeconfig_config_iam_policy":                         tpgiamresource.ResourceIamPolicy(runtimeconfig.RuntimeConfigConfigIamSchema, runtimeconfig.RuntimeConfigConfigIamUpdaterProducer, runtimeconfig.RuntimeConfigConfigIdParseFunc),
			"google_secret_manager_secret":                                   secretmanager.ResourceSecretManagerSecret(),
			"google_secret_manager_secret_iam_binding":                       tpgiamresource.ResourceIamBinding(secretmanager.SecretManagerSecretIamSchema, secretmanager.SecretManagerSecretIamUpdaterProducer, secretmanager.SecretManagerSecretIdParseFunc),
			"google_secret_manager_secret_iam_member":                        tpgiamresource.ResourceIamMember(secretmanager.SecretManagerSecretIamSchema, secretmanager.SecretManagerSecretIamUpdaterProducer, secretmanager.SecretManagerSecretIdParseFunc),
			"google_secret_manager_secret_iam_policy":                        tpgiamresource.ResourceIamPolicy(secretmanager.SecretManagerSecretIamSchema, secretmanager.SecretManagerSecretIamUpdaterProducer, secretmanager.SecretManagerSecretIdParseFunc),
			"google_secret_manager_secret_version":                           secretmanager.ResourceSecretManagerSecretVersion(),
			"google_scc_mute_config":                                         securitycenter.ResourceSecurityCenterMuteConfig(),
			"google_scc_notification_config":                                 securitycenter.ResourceSecurityCenterNotificationConfig(),
			"google_scc_source":                                              securitycenter.ResourceSecurityCenterSource(),
			"google_scc_source_iam_binding":                                  tpgiamresource.ResourceIamBinding(securitycenter.SecurityCenterSourceIamSchema, securitycenter.SecurityCenterSourceIamUpdaterProducer, securitycenter.SecurityCenterSourceIdParseFunc),
			"google_scc_source_iam_member":                                   tpgiamresource.ResourceIamMember(securitycenter.SecurityCenterSourceIamSchema, securitycenter.SecurityCenterSourceIamUpdaterProducer, securitycenter.SecurityCenterSourceIdParseFunc),
			"google_scc_source_iam_policy":                                   tpgiamresource.ResourceIamPolicy(securitycenter.SecurityCenterSourceIamSchema, securitycenter.SecurityCenterSourceIamUpdaterProducer, securitycenter.SecurityCenterSourceIdParseFunc),
			"google_security_scanner_scan_config":                            securityscanner.ResourceSecurityScannerScanConfig(),
			"google_service_directory_endpoint":                              servicedirectory.ResourceServiceDirectoryEndpoint(),
			"google_service_directory_namespace":                             servicedirectory.ResourceServiceDirectoryNamespace(),
			"google_service_directory_namespace_iam_binding":                 tpgiamresource.ResourceIamBinding(servicedirectory.ServiceDirectoryNamespaceIamSchema, servicedirectory.ServiceDirectoryNamespaceIamUpdaterProducer, servicedirectory.ServiceDirectoryNamespaceIdParseFunc),
			"google_service_directory_namespace_iam_member":                  tpgiamresource.ResourceIamMember(servicedirectory.ServiceDirectoryNamespaceIamSchema, servicedirectory.ServiceDirectoryNamespaceIamUpdaterProducer, servicedirectory.ServiceDirectoryNamespaceIdParseFunc),
			"google_service_directory_namespace_iam_policy":                  tpgiamresource.ResourceIamPolicy(servicedirectory.ServiceDirectoryNamespaceIamSchema, servicedirectory.ServiceDirectoryNamespaceIamUpdaterProducer, servicedirectory.ServiceDirectoryNamespaceIdParseFunc),
			"google_service_directory_service":                               servicedirectory.ResourceServiceDirectoryService(),
			"google_service_directory_service_iam_binding":                   tpgiamresource.ResourceIamBinding(servicedirectory.ServiceDirectoryServiceIamSchema, servicedirectory.ServiceDirectoryServiceIamUpdaterProducer, servicedirectory.ServiceDirectoryServiceIdParseFunc),
			"google_service_directory_service_iam_member":                    tpgiamresource.ResourceIamMember(servicedirectory.ServiceDirectoryServiceIamSchema, servicedirectory.ServiceDirectoryServiceIamUpdaterProducer, servicedirectory.ServiceDirectoryServiceIdParseFunc),
			"google_service_directory_service_iam_policy":                    tpgiamresource.ResourceIamPolicy(servicedirectory.ServiceDirectoryServiceIamSchema, servicedirectory.ServiceDirectoryServiceIamUpdaterProducer, servicedirectory.ServiceDirectoryServiceIdParseFunc),
			"google_endpoints_service_iam_binding":                           tpgiamresource.ResourceIamBinding(servicemanagement.ServiceManagementServiceIamSchema, servicemanagement.ServiceManagementServiceIamUpdaterProducer, servicemanagement.ServiceManagementServiceIdParseFunc),
			"google_endpoints_service_iam_member":                            tpgiamresource.ResourceIamMember(servicemanagement.ServiceManagementServiceIamSchema, servicemanagement.ServiceManagementServiceIamUpdaterProducer, servicemanagement.ServiceManagementServiceIdParseFunc),
			"google_endpoints_service_iam_policy":                            tpgiamresource.ResourceIamPolicy(servicemanagement.ServiceManagementServiceIamSchema, servicemanagement.ServiceManagementServiceIamUpdaterProducer, servicemanagement.ServiceManagementServiceIdParseFunc),
			"google_endpoints_service_consumers_iam_binding":                 tpgiamresource.ResourceIamBinding(servicemanagement.ServiceManagementServiceConsumersIamSchema, servicemanagement.ServiceManagementServiceConsumersIamUpdaterProducer, servicemanagement.ServiceManagementServiceConsumersIdParseFunc),
			"google_endpoints_service_consumers_iam_member":                  tpgiamresource.ResourceIamMember(servicemanagement.ServiceManagementServiceConsumersIamSchema, servicemanagement.ServiceManagementServiceConsumersIamUpdaterProducer, servicemanagement.ServiceManagementServiceConsumersIdParseFunc),
			"google_endpoints_service_consumers_iam_policy":                  tpgiamresource.ResourceIamPolicy(servicemanagement.ServiceManagementServiceConsumersIamSchema, servicemanagement.ServiceManagementServiceConsumersIamUpdaterProducer, servicemanagement.ServiceManagementServiceConsumersIdParseFunc),
			"google_service_usage_consumer_quota_override":                   serviceusage.ResourceServiceUsageConsumerQuotaOverride(),
			"google_sourcerepo_repository":                                   sourcerepo.ResourceSourceRepoRepository(),
			"google_sourcerepo_repository_iam_binding":                       tpgiamresource.ResourceIamBinding(sourcerepo.SourceRepoRepositoryIamSchema, sourcerepo.SourceRepoRepositoryIamUpdaterProducer, sourcerepo.SourceRepoRepositoryIdParseFunc),
			"google_sourcerepo_repository_iam_member":                        tpgiamresource.ResourceIamMember(sourcerepo.SourceRepoRepositoryIamSchema, sourcerepo.SourceRepoRepositoryIamUpdaterProducer, sourcerepo.SourceRepoRepositoryIdParseFunc),
			"google_sourcerepo_repository_iam_policy":                        tpgiamresource.ResourceIamPolicy(sourcerepo.SourceRepoRepositoryIamSchema, sourcerepo.SourceRepoRepositoryIamUpdaterProducer, sourcerepo.SourceRepoRepositoryIdParseFunc),
			"google_spanner_database":                                        spanner.ResourceSpannerDatabase(),
			"google_spanner_instance":                                        spanner.ResourceSpannerInstance(),
			"google_sql_database":                                            sql.ResourceSQLDatabase(),
			"google_sql_source_representation_instance":                      sql.ResourceSQLSourceRepresentationInstance(),
			"google_storage_bucket_iam_binding":                              tpgiamresource.ResourceIamBinding(storage.StorageBucketIamSchema, storage.StorageBucketIamUpdaterProducer, storage.StorageBucketIdParseFunc),
			"google_storage_bucket_iam_member":                               tpgiamresource.ResourceIamMember(storage.StorageBucketIamSchema, storage.StorageBucketIamUpdaterProducer, storage.StorageBucketIdParseFunc),
			"google_storage_bucket_iam_policy":                               tpgiamresource.ResourceIamPolicy(storage.StorageBucketIamSchema, storage.StorageBucketIamUpdaterProducer, storage.StorageBucketIdParseFunc),
			"google_storage_bucket_access_control":                           storage.ResourceStorageBucketAccessControl(),
			"google_storage_default_object_access_control":                   storage.ResourceStorageDefaultObjectAccessControl(),
			"google_storage_hmac_key":                                        storage.ResourceStorageHmacKey(),
			"google_storage_object_access_control":                           storage.ResourceStorageObjectAccessControl(),
			"google_storage_transfer_agent_pool":                             storagetransfer.ResourceStorageTransferAgentPool(),
			"google_tags_tag_binding":                                        tags.ResourceTagsTagBinding(),
			"google_tags_tag_key":                                            tags.ResourceTagsTagKey(),
			"google_tags_tag_key_iam_binding":                                tpgiamresource.ResourceIamBinding(tags.TagsTagKeyIamSchema, tags.TagsTagKeyIamUpdaterProducer, tags.TagsTagKeyIdParseFunc),
			"google_tags_tag_key_iam_member":                                 tpgiamresource.ResourceIamMember(tags.TagsTagKeyIamSchema, tags.TagsTagKeyIamUpdaterProducer, tags.TagsTagKeyIdParseFunc),
			"google_tags_tag_key_iam_policy":                                 tpgiamresource.ResourceIamPolicy(tags.TagsTagKeyIamSchema, tags.TagsTagKeyIamUpdaterProducer, tags.TagsTagKeyIdParseFunc),
			"google_tags_tag_value":                                          tags.ResourceTagsTagValue(),
			"google_tags_tag_value_iam_binding":                              tpgiamresource.ResourceIamBinding(tags.TagsTagValueIamSchema, tags.TagsTagValueIamUpdaterProducer, tags.TagsTagValueIdParseFunc),
			"google_tags_tag_value_iam_member":                               tpgiamresource.ResourceIamMember(tags.TagsTagValueIamSchema, tags.TagsTagValueIamUpdaterProducer, tags.TagsTagValueIdParseFunc),
			"google_tags_tag_value_iam_policy":                               tpgiamresource.ResourceIamPolicy(tags.TagsTagValueIamSchema, tags.TagsTagValueIamUpdaterProducer, tags.TagsTagValueIdParseFunc),
			"google_tpu_node":                                                tpu.ResourceTPUNode(),
			"google_vertex_ai_dataset":                                       vertexai.ResourceVertexAIDataset(),
			"google_vertex_ai_endpoint":                                      vertexai.ResourceVertexAIEndpoint(),
			"google_vertex_ai_featurestore":                                  vertexai.ResourceVertexAIFeaturestore(),
			"google_vertex_ai_featurestore_iam_binding":                      tpgiamresource.ResourceIamBinding(vertexai.VertexAIFeaturestoreIamSchema, vertexai.VertexAIFeaturestoreIamUpdaterProducer, vertexai.VertexAIFeaturestoreIdParseFunc),
			"google_vertex_ai_featurestore_iam_member":                       tpgiamresource.ResourceIamMember(vertexai.VertexAIFeaturestoreIamSchema, vertexai.VertexAIFeaturestoreIamUpdaterProducer, vertexai.VertexAIFeaturestoreIdParseFunc),
			"google_vertex_ai_featurestore_iam_policy":                       tpgiamresource.ResourceIamPolicy(vertexai.VertexAIFeaturestoreIamSchema, vertexai.VertexAIFeaturestoreIamUpdaterProducer, vertexai.VertexAIFeaturestoreIdParseFunc),
			"google_vertex_ai_featurestore_entitytype":                       vertexai.ResourceVertexAIFeaturestoreEntitytype(),
			"google_vertex_ai_featurestore_entitytype_iam_binding":           tpgiamresource.ResourceIamBinding(vertexai.VertexAIFeaturestoreEntitytypeIamSchema, vertexai.VertexAIFeaturestoreEntitytypeIamUpdaterProducer, vertexai.VertexAIFeaturestoreEntitytypeIdParseFunc),
			"google_vertex_ai_featurestore_entitytype_iam_member":            tpgiamresource.ResourceIamMember(vertexai.VertexAIFeaturestoreEntitytypeIamSchema, vertexai.VertexAIFeaturestoreEntitytypeIamUpdaterProducer, vertexai.VertexAIFeaturestoreEntitytypeIdParseFunc),
			"google_vertex_ai_featurestore_entitytype_iam_policy":            tpgiamresource.ResourceIamPolicy(vertexai.VertexAIFeaturestoreEntitytypeIamSchema, vertexai.VertexAIFeaturestoreEntitytypeIamUpdaterProducer, vertexai.VertexAIFeaturestoreEntitytypeIdParseFunc),
			"google_vertex_ai_featurestore_entitytype_feature":               vertexai.ResourceVertexAIFeaturestoreEntitytypeFeature(),
			"google_vertex_ai_index":                                         vertexai.ResourceVertexAIIndex(),
			"google_vertex_ai_index_endpoint":                                vertexai.ResourceVertexAIIndexEndpoint(),
			"google_vertex_ai_metadata_store":                                vertexai.ResourceVertexAIMetadataStore(),
			"google_vertex_ai_tensorboard":                                   vertexai.ResourceVertexAITensorboard(),
			"google_vmwareengine_cluster":                                    vmwareengine.ResourceVmwareengineCluster(),
			"google_vmwareengine_network":                                    vmwareengine.ResourceVmwareengineNetwork(),
			"google_vmwareengine_private_cloud":                              vmwareengine.ResourceVmwareenginePrivateCloud(),
			"google_vpc_access_connector":                                    vpcaccess.ResourceVPCAccessConnector(),
			"google_workflows_workflow":                                      workflows.ResourceWorkflowsWorkflow(),
			"google_workstations_workstation":                                workstations.ResourceWorkstationsWorkstation(),
			"google_workstations_workstation_iam_binding":                    tpgiamresource.ResourceIamBinding(workstations.WorkstationsWorkstationIamSchema, workstations.WorkstationsWorkstationIamUpdaterProducer, workstations.WorkstationsWorkstationIdParseFunc),
			"google_workstations_workstation_iam_member":                     tpgiamresource.ResourceIamMember(workstations.WorkstationsWorkstationIamSchema, workstations.WorkstationsWorkstationIamUpdaterProducer, workstations.WorkstationsWorkstationIdParseFunc),
			"google_workstations_workstation_iam_policy":                     tpgiamresource.ResourceIamPolicy(workstations.WorkstationsWorkstationIamSchema, workstations.WorkstationsWorkstationIamUpdaterProducer, workstations.WorkstationsWorkstationIdParseFunc),
			"google_workstations_workstation_cluster":                        workstations.ResourceWorkstationsWorkstationCluster(),
			"google_workstations_workstation_config":                         workstations.ResourceWorkstationsWorkstationConfig(),
			"google_workstations_workstation_config_iam_binding":             tpgiamresource.ResourceIamBinding(workstations.WorkstationsWorkstationConfigIamSchema, workstations.WorkstationsWorkstationConfigIamUpdaterProducer, workstations.WorkstationsWorkstationConfigIdParseFunc),
			"google_workstations_workstation_config_iam_member":              tpgiamresource.ResourceIamMember(workstations.WorkstationsWorkstationConfigIamSchema, workstations.WorkstationsWorkstationConfigIamUpdaterProducer, workstations.WorkstationsWorkstationConfigIdParseFunc),
			"google_workstations_workstation_config_iam_policy":              tpgiamresource.ResourceIamPolicy(workstations.WorkstationsWorkstationConfigIamSchema, workstations.WorkstationsWorkstationConfigIamUpdaterProducer, workstations.WorkstationsWorkstationConfigIdParseFunc),
		},
		map[string]*schema.Resource{
			// ####### START handwritten resources ###########
			"google_app_engine_application":                 appengine.ResourceAppEngineApplication(),
			"google_apigee_sharedflow":                      apigee.ResourceApigeeSharedFlow(),
			"google_apigee_sharedflow_deployment":           apigee.ResourceApigeeSharedFlowDeployment(),
			"google_apigee_flowhook":                        apigee.ResourceApigeeFlowhook(),
			"google_apigee_keystores_aliases_pkcs12":        apigee.ResourceApigeeKeystoresAliasesPkcs12(),
			"google_apigee_keystores_aliases_key_cert_file": apigee.ResourceApigeeKeystoresAliasesKeyCertFile(),
			"google_bigquery_table":                         bigquery.ResourceBigQueryTable(),
			"google_bigtable_gc_policy":                     bigtable.ResourceBigtableGCPolicy(),
			"google_bigtable_instance":                      bigtable.ResourceBigtableInstance(),
			"google_bigtable_table":                         bigtable.ResourceBigtableTable(),
			"google_billing_subaccount":                     resourcemanager.ResourceBillingSubaccount(),
			"google_cloudfunctions_function":                cloudfunctions.ResourceCloudFunctionsFunction(),
			"google_composer_environment":                   composer.ResourceComposerEnvironment(),
			"google_compute_attached_disk":                  compute.ResourceComputeAttachedDisk(),
			"google_compute_instance":                       compute.ResourceComputeInstance(),
			"google_compute_disk_async_replication":         compute.ResourceComputeDiskAsyncReplication(),
			"google_compute_instance_from_machine_image":    compute.ResourceComputeInstanceFromMachineImage(),
			"google_compute_instance_from_template":         compute.ResourceComputeInstanceFromTemplate(),
			"google_compute_instance_group":                 compute.ResourceComputeInstanceGroup(),
			"google_compute_instance_group_manager":         compute.ResourceComputeInstanceGroupManager(),
			"google_compute_instance_template":              compute.ResourceComputeInstanceTemplate(),
			"google_compute_network_peering":                compute.ResourceComputeNetworkPeering(),
			"google_compute_project_default_network_tier":   compute.ResourceComputeProjectDefaultNetworkTier(),
			"google_compute_project_metadata":               compute.ResourceComputeProjectMetadata(),
			"google_compute_project_metadata_item":          compute.ResourceComputeProjectMetadataItem(),
			"google_compute_region_instance_group_manager":  compute.ResourceComputeRegionInstanceGroupManager(),
			"google_compute_region_instance_template":       compute.ResourceComputeRegionInstanceTemplate(),
			"google_compute_router_interface":               compute.ResourceComputeRouterInterface(),
			"google_compute_security_policy":                compute.ResourceComputeSecurityPolicy(),
			"google_compute_shared_vpc_host_project":        compute.ResourceComputeSharedVpcHostProject(),
			"google_compute_shared_vpc_service_project":     compute.ResourceComputeSharedVpcServiceProject(),
			"google_compute_target_pool":                    compute.ResourceComputeTargetPool(),
			"google_container_cluster":                      container.ResourceContainerCluster(),
			"google_container_node_pool":                    container.ResourceContainerNodePool(),
			"google_container_registry":                     containeranalysis.ResourceContainerRegistry(),
			"google_dataflow_job":                           dataflow.ResourceDataflowJob(),
			"google_dataflow_flex_template_job":             dataflow.ResourceDataflowFlexTemplateJob(),
			"google_dataproc_cluster":                       dataproc.ResourceDataprocCluster(),
			"google_dataproc_job":                           dataproc.ResourceDataprocJob(),
			"google_dialogflow_cx_version":                  dialogflowcx.ResourceDialogflowCXVersion(),
			"google_dialogflow_cx_environment":              dialogflowcx.ResourceDialogflowCXEnvironment(),
			"google_dns_record_set":                         dns.ResourceDnsRecordSet(),
			"google_endpoints_service":                      servicemanagement.ResourceEndpointsService(),
			"google_folder":                                 resourcemanager.ResourceGoogleFolder(),
			"google_folder_organization_policy":             resourcemanager.ResourceGoogleFolderOrganizationPolicy(),
			"google_logging_billing_account_sink":           logging.ResourceLoggingBillingAccountSink(),
			"google_logging_billing_account_exclusion":      logging.ResourceLoggingExclusion(logging.BillingAccountLoggingExclusionSchema, logging.NewBillingAccountLoggingExclusionUpdater, logging.BillingAccountLoggingExclusionIdParseFunc),
			"google_logging_billing_account_bucket_config":  logging.ResourceLoggingBillingAccountBucketConfig(),
			"google_logging_organization_sink":              logging.ResourceLoggingOrganizationSink(),
			"google_logging_organization_exclusion":         logging.ResourceLoggingExclusion(logging.OrganizationLoggingExclusionSchema, logging.NewOrganizationLoggingExclusionUpdater, logging.OrganizationLoggingExclusionIdParseFunc),
			"google_logging_organization_bucket_config":     logging.ResourceLoggingOrganizationBucketConfig(),
			"google_logging_folder_sink":                    logging.ResourceLoggingFolderSink(),
			"google_logging_folder_exclusion":               logging.ResourceLoggingExclusion(logging.FolderLoggingExclusionSchema, logging.NewFolderLoggingExclusionUpdater, logging.FolderLoggingExclusionIdParseFunc),
			"google_logging_folder_bucket_config":           logging.ResourceLoggingFolderBucketConfig(),
			"google_logging_project_sink":                   logging.ResourceLoggingProjectSink(),
			"google_logging_project_exclusion":              logging.ResourceLoggingExclusion(logging.ProjectLoggingExclusionSchema, logging.NewProjectLoggingExclusionUpdater, logging.ProjectLoggingExclusionIdParseFunc),
			"google_logging_project_bucket_config":          logging.ResourceLoggingProjectBucketConfig(),
			"google_monitoring_dashboard":                   monitoring.ResourceMonitoringDashboard(),
			"google_os_config_os_policy_assignment":         osconfig.ResourceOSConfigOSPolicyAssignment(),
			"google_project_service_identity":               resourcemanager.ResourceProjectServiceIdentity(),
			"google_service_networking_connection":          servicenetworking.ResourceServiceNetworkingConnection(),
			"google_sql_database_instance":                  sql.ResourceSqlDatabaseInstance(),
			"google_sql_ssl_cert":                           sql.ResourceSqlSslCert(),
			"google_sql_user":                               sql.ResourceSqlUser(),
			"google_organization_iam_custom_role":           resourcemanager.ResourceGoogleOrganizationIamCustomRole(),
			"google_organization_policy":                    resourcemanager.ResourceGoogleOrganizationPolicy(),
			"google_project":                                resourcemanager.ResourceGoogleProject(),
			"google_project_default_service_accounts":       resourcemanager.ResourceGoogleProjectDefaultServiceAccounts(),
			"google_project_service":                        resourcemanager.ResourceGoogleProjectService(),
			"google_project_iam_custom_role":                resourcemanager.ResourceGoogleProjectIamCustomRole(),
			"google_project_organization_policy":            resourcemanager.ResourceGoogleProjectOrganizationPolicy(),
			"google_project_usage_export_bucket":            compute.ResourceProjectUsageBucket(),
			"google_runtimeconfig_config":                   runtimeconfig.ResourceRuntimeconfigConfig(),
			"google_runtimeconfig_variable":                 runtimeconfig.ResourceRuntimeconfigVariable(),
			"google_service_account":                        resourcemanager.ResourceGoogleServiceAccount(),
			"google_service_account_key":                    resourcemanager.ResourceGoogleServiceAccountKey(),
			"google_service_networking_peered_dns_domain":   servicenetworking.ResourceGoogleServiceNetworkingPeeredDNSDomain(),
			"google_storage_bucket":                         storage.ResourceStorageBucket(),
			"google_storage_bucket_acl":                     storage.ResourceStorageBucketAcl(),
			"google_storage_bucket_object":                  storage.ResourceStorageBucketObject(),
			"google_storage_object_acl":                     storage.ResourceStorageObjectAcl(),
			"google_storage_default_object_acl":             storage.ResourceStorageDefaultObjectAcl(),
			"google_storage_notification":                   storage.ResourceStorageNotification(),
			"google_storage_transfer_job":                   storagetransfer.ResourceStorageTransferJob(),
			"google_tags_location_tag_binding":              tags.ResourceTagsLocationTagBinding(),
			// ####### END handwritten resources ###########
		},
		map[string]*schema.Resource{
			// ####### START non-generated IAM resources ###########
			"google_bigtable_instance_iam_binding":       tpgiamresource.ResourceIamBinding(bigtable.IamBigtableInstanceSchema, bigtable.NewBigtableInstanceUpdater, bigtable.BigtableInstanceIdParseFunc),
			"google_bigtable_instance_iam_member":        tpgiamresource.ResourceIamMember(bigtable.IamBigtableInstanceSchema, bigtable.NewBigtableInstanceUpdater, bigtable.BigtableInstanceIdParseFunc),
			"google_bigtable_instance_iam_policy":        tpgiamresource.ResourceIamPolicy(bigtable.IamBigtableInstanceSchema, bigtable.NewBigtableInstanceUpdater, bigtable.BigtableInstanceIdParseFunc),
			"google_bigtable_table_iam_binding":          tpgiamresource.ResourceIamBinding(bigtable.IamBigtableTableSchema, bigtable.NewBigtableTableUpdater, bigtable.BigtableTableIdParseFunc),
			"google_bigtable_table_iam_member":           tpgiamresource.ResourceIamMember(bigtable.IamBigtableTableSchema, bigtable.NewBigtableTableUpdater, bigtable.BigtableTableIdParseFunc),
			"google_bigtable_table_iam_policy":           tpgiamresource.ResourceIamPolicy(bigtable.IamBigtableTableSchema, bigtable.NewBigtableTableUpdater, bigtable.BigtableTableIdParseFunc),
			"google_bigquery_dataset_iam_binding":        tpgiamresource.ResourceIamBinding(bigquery.IamBigqueryDatasetSchema, bigquery.NewBigqueryDatasetIamUpdater, bigquery.BigqueryDatasetIdParseFunc),
			"google_bigquery_dataset_iam_member":         tpgiamresource.ResourceIamMember(bigquery.IamBigqueryDatasetSchema, bigquery.NewBigqueryDatasetIamUpdater, bigquery.BigqueryDatasetIdParseFunc),
			"google_bigquery_dataset_iam_policy":         tpgiamresource.ResourceIamPolicy(bigquery.IamBigqueryDatasetSchema, bigquery.NewBigqueryDatasetIamUpdater, bigquery.BigqueryDatasetIdParseFunc),
			"google_billing_account_iam_binding":         tpgiamresource.ResourceIamBinding(billing.IamBillingAccountSchema, billing.NewBillingAccountIamUpdater, billing.BillingAccountIdParseFunc),
			"google_billing_account_iam_member":          tpgiamresource.ResourceIamMember(billing.IamBillingAccountSchema, billing.NewBillingAccountIamUpdater, billing.BillingAccountIdParseFunc),
			"google_billing_account_iam_policy":          tpgiamresource.ResourceIamPolicy(billing.IamBillingAccountSchema, billing.NewBillingAccountIamUpdater, billing.BillingAccountIdParseFunc),
			"google_dataproc_cluster_iam_binding":        tpgiamresource.ResourceIamBinding(dataproc.IamDataprocClusterSchema, dataproc.NewDataprocClusterUpdater, dataproc.DataprocClusterIdParseFunc),
			"google_dataproc_cluster_iam_member":         tpgiamresource.ResourceIamMember(dataproc.IamDataprocClusterSchema, dataproc.NewDataprocClusterUpdater, dataproc.DataprocClusterIdParseFunc),
			"google_dataproc_cluster_iam_policy":         tpgiamresource.ResourceIamPolicy(dataproc.IamDataprocClusterSchema, dataproc.NewDataprocClusterUpdater, dataproc.DataprocClusterIdParseFunc),
			"google_dataproc_job_iam_binding":            tpgiamresource.ResourceIamBinding(dataproc.IamDataprocJobSchema, dataproc.NewDataprocJobUpdater, dataproc.DataprocJobIdParseFunc),
			"google_dataproc_job_iam_member":             tpgiamresource.ResourceIamMember(dataproc.IamDataprocJobSchema, dataproc.NewDataprocJobUpdater, dataproc.DataprocJobIdParseFunc),
			"google_dataproc_job_iam_policy":             tpgiamresource.ResourceIamPolicy(dataproc.IamDataprocJobSchema, dataproc.NewDataprocJobUpdater, dataproc.DataprocJobIdParseFunc),
			"google_folder_iam_binding":                  tpgiamresource.ResourceIamBinding(resourcemanager.IamFolderSchema, resourcemanager.NewFolderIamUpdater, resourcemanager.FolderIdParseFunc),
			"google_folder_iam_member":                   tpgiamresource.ResourceIamMember(resourcemanager.IamFolderSchema, resourcemanager.NewFolderIamUpdater, resourcemanager.FolderIdParseFunc),
			"google_folder_iam_policy":                   tpgiamresource.ResourceIamPolicy(resourcemanager.IamFolderSchema, resourcemanager.NewFolderIamUpdater, resourcemanager.FolderIdParseFunc),
			"google_folder_iam_audit_config":             tpgiamresource.ResourceIamAuditConfig(resourcemanager.IamFolderSchema, resourcemanager.NewFolderIamUpdater, resourcemanager.FolderIdParseFunc),
			"google_healthcare_dataset_iam_binding":      tpgiamresource.ResourceIamBinding(healthcare.IamHealthcareDatasetSchema, healthcare.NewHealthcareDatasetIamUpdater, healthcare.DatasetIdParseFunc, tpgiamresource.IamWithBatching),
			"google_healthcare_dataset_iam_member":       tpgiamresource.ResourceIamMember(healthcare.IamHealthcareDatasetSchema, healthcare.NewHealthcareDatasetIamUpdater, healthcare.DatasetIdParseFunc, tpgiamresource.IamWithBatching),
			"google_healthcare_dataset_iam_policy":       tpgiamresource.ResourceIamPolicy(healthcare.IamHealthcareDatasetSchema, healthcare.NewHealthcareDatasetIamUpdater, healthcare.DatasetIdParseFunc),
			"google_healthcare_dicom_store_iam_binding":  tpgiamresource.ResourceIamBinding(healthcare.IamHealthcareDicomStoreSchema, healthcare.NewHealthcareDicomStoreIamUpdater, healthcare.DicomStoreIdParseFunc, tpgiamresource.IamWithBatching),
			"google_healthcare_dicom_store_iam_member":   tpgiamresource.ResourceIamMember(healthcare.IamHealthcareDicomStoreSchema, healthcare.NewHealthcareDicomStoreIamUpdater, healthcare.DicomStoreIdParseFunc, tpgiamresource.IamWithBatching),
			"google_healthcare_dicom_store_iam_policy":   tpgiamresource.ResourceIamPolicy(healthcare.IamHealthcareDicomStoreSchema, healthcare.NewHealthcareDicomStoreIamUpdater, healthcare.DicomStoreIdParseFunc),
			"google_healthcare_fhir_store_iam_binding":   tpgiamresource.ResourceIamBinding(healthcare.IamHealthcareFhirStoreSchema, healthcare.NewHealthcareFhirStoreIamUpdater, healthcare.FhirStoreIdParseFunc, tpgiamresource.IamWithBatching),
			"google_healthcare_fhir_store_iam_member":    tpgiamresource.ResourceIamMember(healthcare.IamHealthcareFhirStoreSchema, healthcare.NewHealthcareFhirStoreIamUpdater, healthcare.FhirStoreIdParseFunc, tpgiamresource.IamWithBatching),
			"google_healthcare_fhir_store_iam_policy":    tpgiamresource.ResourceIamPolicy(healthcare.IamHealthcareFhirStoreSchema, healthcare.NewHealthcareFhirStoreIamUpdater, healthcare.FhirStoreIdParseFunc),
			"google_healthcare_hl7_v2_store_iam_binding": tpgiamresource.ResourceIamBinding(healthcare.IamHealthcareHl7V2StoreSchema, healthcare.NewHealthcareHl7V2StoreIamUpdater, healthcare.Hl7V2StoreIdParseFunc, tpgiamresource.IamWithBatching),
			"google_healthcare_hl7_v2_store_iam_member":  tpgiamresource.ResourceIamMember(healthcare.IamHealthcareHl7V2StoreSchema, healthcare.NewHealthcareHl7V2StoreIamUpdater, healthcare.Hl7V2StoreIdParseFunc, tpgiamresource.IamWithBatching),
			"google_healthcare_hl7_v2_store_iam_policy":  tpgiamresource.ResourceIamPolicy(healthcare.IamHealthcareHl7V2StoreSchema, healthcare.NewHealthcareHl7V2StoreIamUpdater, healthcare.Hl7V2StoreIdParseFunc),
			"google_kms_key_ring_iam_binding":            tpgiamresource.ResourceIamBinding(kms.IamKmsKeyRingSchema, kms.NewKmsKeyRingIamUpdater, kms.KeyRingIdParseFunc),
			"google_kms_key_ring_iam_member":             tpgiamresource.ResourceIamMember(kms.IamKmsKeyRingSchema, kms.NewKmsKeyRingIamUpdater, kms.KeyRingIdParseFunc),
			"google_kms_key_ring_iam_policy":             tpgiamresource.ResourceIamPolicy(kms.IamKmsKeyRingSchema, kms.NewKmsKeyRingIamUpdater, kms.KeyRingIdParseFunc),
			"google_kms_crypto_key_iam_binding":          tpgiamresource.ResourceIamBinding(kms.IamKmsCryptoKeySchema, kms.NewKmsCryptoKeyIamUpdater, kms.CryptoIdParseFunc),
			"google_kms_crypto_key_iam_member":           tpgiamresource.ResourceIamMember(kms.IamKmsCryptoKeySchema, kms.NewKmsCryptoKeyIamUpdater, kms.CryptoIdParseFunc),
			"google_kms_crypto_key_iam_policy":           tpgiamresource.ResourceIamPolicy(kms.IamKmsCryptoKeySchema, kms.NewKmsCryptoKeyIamUpdater, kms.CryptoIdParseFunc),
			"google_spanner_instance_iam_binding":        tpgiamresource.ResourceIamBinding(spanner.IamSpannerInstanceSchema, spanner.NewSpannerInstanceIamUpdater, spanner.SpannerInstanceIdParseFunc),
			"google_spanner_instance_iam_member":         tpgiamresource.ResourceIamMember(spanner.IamSpannerInstanceSchema, spanner.NewSpannerInstanceIamUpdater, spanner.SpannerInstanceIdParseFunc),
			"google_spanner_instance_iam_policy":         tpgiamresource.ResourceIamPolicy(spanner.IamSpannerInstanceSchema, spanner.NewSpannerInstanceIamUpdater, spanner.SpannerInstanceIdParseFunc),
			"google_spanner_database_iam_binding":        tpgiamresource.ResourceIamBinding(spanner.IamSpannerDatabaseSchema, spanner.NewSpannerDatabaseIamUpdater, spanner.SpannerDatabaseIdParseFunc),
			"google_spanner_database_iam_member":         tpgiamresource.ResourceIamMember(spanner.IamSpannerDatabaseSchema, spanner.NewSpannerDatabaseIamUpdater, spanner.SpannerDatabaseIdParseFunc),
			"google_spanner_database_iam_policy":         tpgiamresource.ResourceIamPolicy(spanner.IamSpannerDatabaseSchema, spanner.NewSpannerDatabaseIamUpdater, spanner.SpannerDatabaseIdParseFunc),
			"google_organization_iam_binding":            tpgiamresource.ResourceIamBinding(resourcemanager.IamOrganizationSchema, resourcemanager.NewOrganizationIamUpdater, resourcemanager.OrgIdParseFunc),
			"google_organization_iam_member":             tpgiamresource.ResourceIamMember(resourcemanager.IamOrganizationSchema, resourcemanager.NewOrganizationIamUpdater, resourcemanager.OrgIdParseFunc),
			"google_organization_iam_policy":             tpgiamresource.ResourceIamPolicy(resourcemanager.IamOrganizationSchema, resourcemanager.NewOrganizationIamUpdater, resourcemanager.OrgIdParseFunc),
			"google_organization_iam_audit_config":       tpgiamresource.ResourceIamAuditConfig(resourcemanager.IamOrganizationSchema, resourcemanager.NewOrganizationIamUpdater, resourcemanager.OrgIdParseFunc),
			"google_project_iam_policy":                  tpgiamresource.ResourceIamPolicy(resourcemanager.IamProjectSchema, resourcemanager.NewProjectIamUpdater, resourcemanager.ProjectIdParseFunc),
			"google_project_iam_binding":                 tpgiamresource.ResourceIamBinding(resourcemanager.IamProjectSchema, resourcemanager.NewProjectIamUpdater, resourcemanager.ProjectIdParseFunc, tpgiamresource.IamWithBatching),
			"google_project_iam_member":                  tpgiamresource.ResourceIamMember(resourcemanager.IamProjectSchema, resourcemanager.NewProjectIamUpdater, resourcemanager.ProjectIdParseFunc, tpgiamresource.IamWithBatching),
			"google_project_iam_audit_config":            tpgiamresource.ResourceIamAuditConfig(resourcemanager.IamProjectSchema, resourcemanager.NewProjectIamUpdater, resourcemanager.ProjectIdParseFunc, tpgiamresource.IamWithBatching),
			"google_pubsub_subscription_iam_binding":     tpgiamresource.ResourceIamBinding(pubsub.IamPubsubSubscriptionSchema, pubsub.NewPubsubSubscriptionIamUpdater, pubsub.PubsubSubscriptionIdParseFunc),
			"google_pubsub_subscription_iam_member":      tpgiamresource.ResourceIamMember(pubsub.IamPubsubSubscriptionSchema, pubsub.NewPubsubSubscriptionIamUpdater, pubsub.PubsubSubscriptionIdParseFunc),
			"google_pubsub_subscription_iam_policy":      tpgiamresource.ResourceIamPolicy(pubsub.IamPubsubSubscriptionSchema, pubsub.NewPubsubSubscriptionIamUpdater, pubsub.PubsubSubscriptionIdParseFunc),
			"google_service_account_iam_binding":         tpgiamresource.ResourceIamBinding(resourcemanager.IamServiceAccountSchema, resourcemanager.NewServiceAccountIamUpdater, resourcemanager.ServiceAccountIdParseFunc),
			"google_service_account_iam_member":          tpgiamresource.ResourceIamMember(resourcemanager.IamServiceAccountSchema, resourcemanager.NewServiceAccountIamUpdater, resourcemanager.ServiceAccountIdParseFunc),
			"google_service_account_iam_policy":          tpgiamresource.ResourceIamPolicy(resourcemanager.IamServiceAccountSchema, resourcemanager.NewServiceAccountIamUpdater, resourcemanager.ServiceAccountIdParseFunc),
			// ####### END non-generated IAM resources ###########
		},
		dclResources,
	)
}

func ProviderConfigure(ctx context.Context, d *schema.ResourceData, p *schema.Provider) (interface{}, diag.Diagnostics) {
	err := transport_tpg.HandleSDKDefaults(d)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	transport_tpg.HandleDCLCustomEndpointDefaults(d)

	config := transport_tpg.Config{
		Project:             d.Get("project").(string),
		Region:              d.Get("region").(string),
		Zone:                d.Get("zone").(string),
		UserProjectOverride: d.Get("user_project_override").(bool),
		BillingProject:      d.Get("billing_project").(string),
		UserAgent:           p.UserAgent("terraform-provider-google-beta", version.ProviderVersion),
	}

	// opt in extension for adding to the User-Agent header
	if ext := os.Getenv("GOOGLE_TERRAFORM_USERAGENT_EXTENSION"); ext != "" {
		ua := config.UserAgent
		config.UserAgent = fmt.Sprintf("%s %s", ua, ext)
	}

	if v, ok := d.GetOk("request_timeout"); ok {
		var err error
		config.RequestTimeout, err = time.ParseDuration(v.(string))
		if err != nil {
			return nil, diag.FromErr(err)
		}
	}

	if v, ok := d.GetOk("request_reason"); ok {
		config.RequestReason = v.(string)
	}

	// Check for primary credentials in config. Note that if neither is set, ADCs
	// will be used if available.
	if v, ok := d.GetOk("access_token"); ok {
		config.AccessToken = v.(string)
	}

	if v, ok := d.GetOk("credentials"); ok {
		config.Credentials = v.(string)
	}

	// only check environment variables if neither value was set in config- this
	// means config beats env var in all cases.
	if config.AccessToken == "" && config.Credentials == "" {
		config.Credentials = transport_tpg.MultiEnvSearch([]string{
			"GOOGLE_CREDENTIALS",
			"GOOGLE_CLOUD_KEYFILE_JSON",
			"GCLOUD_KEYFILE_JSON",
		})

		config.AccessToken = transport_tpg.MultiEnvSearch([]string{
			"GOOGLE_OAUTH_ACCESS_TOKEN",
		})
	}

	// Given that impersonate_service_account is a secondary auth method, it has
	// no conflicts to worry about. We pull the env var in a DefaultFunc.
	if v, ok := d.GetOk("impersonate_service_account"); ok {
		config.ImpersonateServiceAccount = v.(string)
	}

	delegates := d.Get("impersonate_service_account_delegates").([]interface{})
	if len(delegates) > 0 {
		config.ImpersonateServiceAccountDelegates = make([]string, len(delegates))
	}
	for i, delegate := range delegates {
		config.ImpersonateServiceAccountDelegates[i] = delegate.(string)
	}

	scopes := d.Get("scopes").([]interface{})
	if len(scopes) > 0 {
		config.Scopes = make([]string, len(scopes))
	}
	for i, scope := range scopes {
		config.Scopes[i] = scope.(string)
	}

	batchCfg, err := transport_tpg.ExpandProviderBatchingConfig(d.Get("batching"))
	if err != nil {
		return nil, diag.FromErr(err)
	}
	config.BatchingConfig = batchCfg

	// Generated products
	config.AccessApprovalBasePath = d.Get("access_approval_custom_endpoint").(string)
	config.AccessContextManagerBasePath = d.Get("access_context_manager_custom_endpoint").(string)
	config.ActiveDirectoryBasePath = d.Get("active_directory_custom_endpoint").(string)
	config.AlloydbBasePath = d.Get("alloydb_custom_endpoint").(string)
	config.ApiGatewayBasePath = d.Get("api_gateway_custom_endpoint").(string)
	config.ApigeeBasePath = d.Get("apigee_custom_endpoint").(string)
	config.AppEngineBasePath = d.Get("app_engine_custom_endpoint").(string)
	config.ArtifactRegistryBasePath = d.Get("artifact_registry_custom_endpoint").(string)
	config.BackupDRBasePath = d.Get("backup_dr_custom_endpoint").(string)
	config.BeyondcorpBasePath = d.Get("beyondcorp_custom_endpoint").(string)
	config.BiglakeBasePath = d.Get("biglake_custom_endpoint").(string)
	config.BigQueryBasePath = d.Get("big_query_custom_endpoint").(string)
	config.BigqueryAnalyticsHubBasePath = d.Get("bigquery_analytics_hub_custom_endpoint").(string)
	config.BigqueryConnectionBasePath = d.Get("bigquery_connection_custom_endpoint").(string)
	config.BigqueryDatapolicyBasePath = d.Get("bigquery_datapolicy_custom_endpoint").(string)
	config.BigqueryDataTransferBasePath = d.Get("bigquery_data_transfer_custom_endpoint").(string)
	config.BigqueryReservationBasePath = d.Get("bigquery_reservation_custom_endpoint").(string)
	config.BigtableBasePath = d.Get("bigtable_custom_endpoint").(string)
	config.BillingBasePath = d.Get("billing_custom_endpoint").(string)
	config.BinaryAuthorizationBasePath = d.Get("binary_authorization_custom_endpoint").(string)
	config.CertificateManagerBasePath = d.Get("certificate_manager_custom_endpoint").(string)
	config.CloudAssetBasePath = d.Get("cloud_asset_custom_endpoint").(string)
	config.CloudBuildBasePath = d.Get("cloud_build_custom_endpoint").(string)
	config.Cloudbuildv2BasePath = d.Get("cloudbuildv2_custom_endpoint").(string)
	config.CloudFunctionsBasePath = d.Get("cloud_functions_custom_endpoint").(string)
	config.Cloudfunctions2BasePath = d.Get("cloudfunctions2_custom_endpoint").(string)
	config.CloudIdentityBasePath = d.Get("cloud_identity_custom_endpoint").(string)
	config.CloudIdsBasePath = d.Get("cloud_ids_custom_endpoint").(string)
	config.CloudIotBasePath = d.Get("cloud_iot_custom_endpoint").(string)
	config.CloudRunBasePath = d.Get("cloud_run_custom_endpoint").(string)
	config.CloudRunV2BasePath = d.Get("cloud_run_v2_custom_endpoint").(string)
	config.CloudSchedulerBasePath = d.Get("cloud_scheduler_custom_endpoint").(string)
	config.CloudTasksBasePath = d.Get("cloud_tasks_custom_endpoint").(string)
	config.ComputeBasePath = d.Get("compute_custom_endpoint").(string)
	config.ContainerAnalysisBasePath = d.Get("container_analysis_custom_endpoint").(string)
	config.ContainerAttachedBasePath = d.Get("container_attached_custom_endpoint").(string)
	config.CoreBillingBasePath = d.Get("core_billing_custom_endpoint").(string)
	config.DatabaseMigrationServiceBasePath = d.Get("database_migration_service_custom_endpoint").(string)
	config.DataCatalogBasePath = d.Get("data_catalog_custom_endpoint").(string)
	config.DataformBasePath = d.Get("dataform_custom_endpoint").(string)
	config.DataFusionBasePath = d.Get("data_fusion_custom_endpoint").(string)
	config.DataLossPreventionBasePath = d.Get("data_loss_prevention_custom_endpoint").(string)
	config.DataplexBasePath = d.Get("dataplex_custom_endpoint").(string)
	config.DataprocBasePath = d.Get("dataproc_custom_endpoint").(string)
	config.DataprocMetastoreBasePath = d.Get("dataproc_metastore_custom_endpoint").(string)
	config.DatastoreBasePath = d.Get("datastore_custom_endpoint").(string)
	config.DatastreamBasePath = d.Get("datastream_custom_endpoint").(string)
	config.DeploymentManagerBasePath = d.Get("deployment_manager_custom_endpoint").(string)
	config.DialogflowBasePath = d.Get("dialogflow_custom_endpoint").(string)
	config.DialogflowCXBasePath = d.Get("dialogflow_cx_custom_endpoint").(string)
	config.DNSBasePath = d.Get("dns_custom_endpoint").(string)
	config.DocumentAIBasePath = d.Get("document_ai_custom_endpoint").(string)
	config.DocumentAIWarehouseBasePath = d.Get("document_ai_warehouse_custom_endpoint").(string)
	config.EssentialContactsBasePath = d.Get("essential_contacts_custom_endpoint").(string)
	config.FilestoreBasePath = d.Get("filestore_custom_endpoint").(string)
	config.FirebaseBasePath = d.Get("firebase_custom_endpoint").(string)
	config.FirebaseDatabaseBasePath = d.Get("firebase_database_custom_endpoint").(string)
	config.FirebaseExtensionsBasePath = d.Get("firebase_extensions_custom_endpoint").(string)
	config.FirebaseHostingBasePath = d.Get("firebase_hosting_custom_endpoint").(string)
	config.FirebaseStorageBasePath = d.Get("firebase_storage_custom_endpoint").(string)
	config.FirestoreBasePath = d.Get("firestore_custom_endpoint").(string)
	config.GameServicesBasePath = d.Get("game_services_custom_endpoint").(string)
	config.GKEBackupBasePath = d.Get("gke_backup_custom_endpoint").(string)
	config.GKEHubBasePath = d.Get("gke_hub_custom_endpoint").(string)
	config.GKEHub2BasePath = d.Get("gke_hub2_custom_endpoint").(string)
	config.GkeonpremBasePath = d.Get("gkeonprem_custom_endpoint").(string)
	config.HealthcareBasePath = d.Get("healthcare_custom_endpoint").(string)
	config.IAM2BasePath = d.Get("iam2_custom_endpoint").(string)
	config.IAMBetaBasePath = d.Get("iam_beta_custom_endpoint").(string)
	config.IAMWorkforcePoolBasePath = d.Get("iam_workforce_pool_custom_endpoint").(string)
	config.IapBasePath = d.Get("iap_custom_endpoint").(string)
	config.IdentityPlatformBasePath = d.Get("identity_platform_custom_endpoint").(string)
	config.KMSBasePath = d.Get("kms_custom_endpoint").(string)
	config.LoggingBasePath = d.Get("logging_custom_endpoint").(string)
	config.LookerBasePath = d.Get("looker_custom_endpoint").(string)
	config.MemcacheBasePath = d.Get("memcache_custom_endpoint").(string)
	config.MLEngineBasePath = d.Get("ml_engine_custom_endpoint").(string)
	config.MonitoringBasePath = d.Get("monitoring_custom_endpoint").(string)
	config.NetworkConnectivityBasePath = d.Get("network_connectivity_custom_endpoint").(string)
	config.NetworkManagementBasePath = d.Get("network_management_custom_endpoint").(string)
	config.NetworkSecurityBasePath = d.Get("network_security_custom_endpoint").(string)
	config.NetworkServicesBasePath = d.Get("network_services_custom_endpoint").(string)
	config.NotebooksBasePath = d.Get("notebooks_custom_endpoint").(string)
	config.OrgPolicyBasePath = d.Get("org_policy_custom_endpoint").(string)
	config.OSConfigBasePath = d.Get("os_config_custom_endpoint").(string)
	config.OSLoginBasePath = d.Get("os_login_custom_endpoint").(string)
	config.PrivatecaBasePath = d.Get("privateca_custom_endpoint").(string)
	config.PublicCABasePath = d.Get("public_ca_custom_endpoint").(string)
	config.PubsubBasePath = d.Get("pubsub_custom_endpoint").(string)
	config.PubsubLiteBasePath = d.Get("pubsub_lite_custom_endpoint").(string)
	config.RedisBasePath = d.Get("redis_custom_endpoint").(string)
	config.ResourceManagerBasePath = d.Get("resource_manager_custom_endpoint").(string)
	config.RuntimeConfigBasePath = d.Get("runtime_config_custom_endpoint").(string)
	config.SecretManagerBasePath = d.Get("secret_manager_custom_endpoint").(string)
	config.SecurityCenterBasePath = d.Get("security_center_custom_endpoint").(string)
	config.SecurityScannerBasePath = d.Get("security_scanner_custom_endpoint").(string)
	config.ServiceDirectoryBasePath = d.Get("service_directory_custom_endpoint").(string)
	config.ServiceManagementBasePath = d.Get("service_management_custom_endpoint").(string)
	config.ServiceUsageBasePath = d.Get("service_usage_custom_endpoint").(string)
	config.SourceRepoBasePath = d.Get("source_repo_custom_endpoint").(string)
	config.SpannerBasePath = d.Get("spanner_custom_endpoint").(string)
	config.SQLBasePath = d.Get("sql_custom_endpoint").(string)
	config.StorageBasePath = d.Get("storage_custom_endpoint").(string)
	config.StorageTransferBasePath = d.Get("storage_transfer_custom_endpoint").(string)
	config.TagsBasePath = d.Get("tags_custom_endpoint").(string)
	config.TPUBasePath = d.Get("tpu_custom_endpoint").(string)
	config.VertexAIBasePath = d.Get("vertex_ai_custom_endpoint").(string)
	config.VmwareengineBasePath = d.Get("vmwareengine_custom_endpoint").(string)
	config.VPCAccessBasePath = d.Get("vpc_access_custom_endpoint").(string)
	config.WorkflowsBasePath = d.Get("workflows_custom_endpoint").(string)
	config.WorkstationsBasePath = d.Get("workstations_custom_endpoint").(string)

	// Handwritten Products / Versioned / Atypical Entries
	config.CloudBillingBasePath = d.Get(transport_tpg.CloudBillingCustomEndpointEntryKey).(string)
	config.ComposerBasePath = d.Get(transport_tpg.ComposerCustomEndpointEntryKey).(string)
	config.ContainerBasePath = d.Get(transport_tpg.ContainerCustomEndpointEntryKey).(string)
	config.DataflowBasePath = d.Get(transport_tpg.DataflowCustomEndpointEntryKey).(string)
	config.IamCredentialsBasePath = d.Get(transport_tpg.IamCredentialsCustomEndpointEntryKey).(string)
	config.ResourceManagerV3BasePath = d.Get(transport_tpg.ResourceManagerV3CustomEndpointEntryKey).(string)
	config.RuntimeConfigBasePath = d.Get(transport_tpg.RuntimeConfigCustomEndpointEntryKey).(string)
	config.IAMBasePath = d.Get(transport_tpg.IAMCustomEndpointEntryKey).(string)
	config.ServiceNetworkingBasePath = d.Get(transport_tpg.ServiceNetworkingCustomEndpointEntryKey).(string)
	config.ServiceUsageBasePath = d.Get(transport_tpg.ServiceUsageCustomEndpointEntryKey).(string)
	config.BigtableAdminBasePath = d.Get(transport_tpg.BigtableAdminCustomEndpointEntryKey).(string)
	config.TagsLocationBasePath = d.Get(transport_tpg.TagsLocationCustomEndpointEntryKey).(string)

	// dcl
	config.ContainerAwsBasePath = d.Get(transport_tpg.ContainerAwsCustomEndpointEntryKey).(string)
	config.ContainerAzureBasePath = d.Get(transport_tpg.ContainerAzureCustomEndpointEntryKey).(string)

	stopCtx, ok := schema.StopContext(ctx)
	if !ok {
		stopCtx = ctx
	}
	if err := config.LoadAndValidate(stopCtx); err != nil {
		return nil, diag.FromErr(err)
	}

	return transport_tpg.ProviderDCLConfigure(d, &config), nil
}

func ValidateCredentials(v interface{}, k string) (warnings []string, errors []error) {
	if v == nil || v.(string) == "" {
		return
	}
	creds := v.(string)
	// if this is a path and we can stat it, assume it's ok
	if _, err := os.Stat(creds); err == nil {
		return
	}
	if _, err := googleoauth.CredentialsFromJSON(context.Background(), []byte(creds)); err != nil {
		errors = append(errors,
			fmt.Errorf("JSON credentials are not valid: %s", err))
	}

	return
}

func mergeResourceMaps(ms ...map[string]*schema.Resource) (map[string]*schema.Resource, error) {
	merged := make(map[string]*schema.Resource)
	duplicates := []string{}

	for _, m := range ms {
		for k, v := range m {
			if _, ok := merged[k]; ok {
				duplicates = append(duplicates, k)
			}

			merged[k] = v
		}
	}

	var err error
	if len(duplicates) > 0 {
		err = fmt.Errorf("saw duplicates in mergeResourceMaps: %v", duplicates)
	}

	return merged, err
}
