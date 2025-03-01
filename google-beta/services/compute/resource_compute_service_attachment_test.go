// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccComputeServiceAttachment_serviceAttachmentBasicExampleUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeServiceAttachmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeServiceAttachment_serviceAttachmentBasicExampleFork(context),
			},
			{
				ResourceName:            "google_compute_service_attachment.psc_ilb_service_attachment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"target_service", "region"},
			},
			{
				Config: testAccComputeServiceAttachment_serviceAttachmentBasicExampleUpdate(context),
			},
			{
				ResourceName:            "google_compute_service_attachment.psc_ilb_service_attachment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"target_service", "region"},
			},
		},
	})
}

func testAccComputeServiceAttachment_serviceAttachmentBasicExampleFork(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_service_attachment" "psc_ilb_service_attachment" {
  name        = "tf-test-my-psc-ilb%{random_suffix}"
  region      = "us-west2"
  description = "A service attachment configured with Terraform"

  enable_proxy_protocol    = true
  connection_preference    = "ACCEPT_AUTOMATIC"
  nat_subnets              = [google_compute_subnetwork.psc_ilb_nat.id]
  target_service           = google_compute_forwarding_rule.psc_ilb_target_service.id
}

resource "google_compute_address" "psc_ilb_consumer_address" {
  name   = "tf-test-psc-ilb-consumer-address%{random_suffix}"
  region = "us-west2"

  subnetwork   = "default"
  address_type = "INTERNAL"
}

resource "google_compute_forwarding_rule" "psc_ilb_consumer" {
  name   = "tf-test-psc-ilb-consumer-forwarding-rule%{random_suffix}"
  region = "us-west2"

  target                = google_compute_service_attachment.psc_ilb_service_attachment.id
  load_balancing_scheme = "" # need to override EXTERNAL default when target is a service attachment
  network               = "default"
  ip_address            = google_compute_address.psc_ilb_consumer_address.id
}

resource "google_compute_forwarding_rule" "psc_ilb_target_service" {
  name   = "tf-test-producer-forwarding-rule%{random_suffix}"
  region = "us-west2"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.producer_service_backend.id
  all_ports             = true
  network               = google_compute_network.psc_ilb_network.name
  subnetwork            = google_compute_subnetwork.psc_ilb_producer_subnetwork.name
}

resource "google_compute_region_backend_service" "producer_service_backend" {
  name   = "tf-test-producer-service%{random_suffix}"
  region = "us-west2"

  health_checks = [google_compute_health_check.producer_service_health_check.id]
}

resource "google_compute_health_check" "producer_service_health_check" {
  name = "tf-test-producer-service-health-check%{random_suffix}"

  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_network" "psc_ilb_network" {
  name = "tf-test-psc-ilb-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "psc_ilb_producer_subnetwork" {
  name   = "tf-test-psc-ilb-producer-subnetwork%{random_suffix}"
  region = "us-west2"

  network       = google_compute_network.psc_ilb_network.id
  ip_cidr_range = "10.0.0.0/16"
}

resource "google_compute_subnetwork" "psc_ilb_nat" {
  name   = "tf-test-psc-ilb-nat%{random_suffix}"
  region = "us-west2"

  network       = google_compute_network.psc_ilb_network.id
  purpose       =  "PRIVATE_SERVICE_CONNECT"
  ip_cidr_range = "10.1.0.0/16"
}
`, context)
}

func testAccComputeServiceAttachment_serviceAttachmentBasicExampleUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_service_attachment" "psc_ilb_service_attachment" {
  name        = "tf-test-my-psc-ilb%{random_suffix}"
  region      = "us-west2"
  description = "A service attachment configured with Terraforms"

  enable_proxy_protocol    = true
  connection_preference    = "ACCEPT_MANUAL"
  nat_subnets              = [google_compute_subnetwork.psc_ilb_nat.id]
  target_service           = google_compute_forwarding_rule.psc_ilb_target_service.id

  consumer_reject_lists = ["673497134629", "482878270665"]
  consumer_accept_lists {
    project_id_or_num = "658859330310"
    connection_limit  = 4
  }

  reconcile_connections = false
}

resource "google_compute_address" "psc_ilb_consumer_address" {
  name   = "tf-test-psc-ilb-consumer-address%{random_suffix}"
  region = "us-west2"

  subnetwork   = "default"
  address_type = "INTERNAL"
}

resource "google_compute_forwarding_rule" "psc_ilb_consumer" {
  name   = "tf-test-psc-ilb-consumer-forwarding-rule%{random_suffix}"
  region = "us-west2"

  target                = google_compute_service_attachment.psc_ilb_service_attachment.id
  load_balancing_scheme = "" # need to override EXTERNAL default when target is a service attachment
  network               = "default"
  ip_address            = google_compute_address.psc_ilb_consumer_address.id
}

resource "google_compute_forwarding_rule" "psc_ilb_target_service" {
  name   = "tf-test-producer-forwarding-rule%{random_suffix}"
  region = "us-west2"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.producer_service_backend.id
  all_ports             = true
  network               = google_compute_network.psc_ilb_network.name
  subnetwork            = google_compute_subnetwork.psc_ilb_producer_subnetwork.name
}

resource "google_compute_region_backend_service" "producer_service_backend" {
  name   = "tf-test-producer-service%{random_suffix}"
  region = "us-west2"

  health_checks = [google_compute_health_check.producer_service_health_check.id]
}

resource "google_compute_health_check" "producer_service_health_check" {
  name = "tf-test-producer-service-health-check%{random_suffix}"

  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_network" "psc_ilb_network" {
  name = "tf-test-psc-ilb-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "psc_ilb_producer_subnetwork" {
  name   = "tf-test-psc-ilb-producer-subnetwork%{random_suffix}"
  region = "us-west2"

  network       = google_compute_network.psc_ilb_network.id
  ip_cidr_range = "10.0.0.0/16"
}

resource "google_compute_subnetwork" "psc_ilb_nat" {
  name   = "tf-test-psc-ilb-nat%{random_suffix}"
  region = "us-west2"

  network       = google_compute_network.psc_ilb_network.id
  purpose       =  "PRIVATE_SERVICE_CONNECT"
  ip_cidr_range = "10.1.0.0/16"
}
`, context)
}
