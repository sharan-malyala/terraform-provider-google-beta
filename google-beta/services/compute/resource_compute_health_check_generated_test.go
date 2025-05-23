// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeHealthCheck_healthCheckTcpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckTcpExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.tcp-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckTcpExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "tcp-health-check" {
  name = "tf-test-tcp-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckTcpFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckTcpFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.tcp-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckTcpFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "tcp-health-check" {
  name        = "tf-test-tcp-health-check%{random_suffix}"
  description = "Health check via tcp"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  tcp_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    request            = "ARE YOU HEALTHY?"
    proxy_header       = "NONE"
    response           = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckSslExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckSslExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.ssl-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckSslExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "ssl-health-check" {
  name = "tf-test-ssl-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  ssl_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckSslFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckSslFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.ssl-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckSslFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "ssl-health-check" {
  name        = "tf-test-ssl-health-check%{random_suffix}"
  description = "Health check via ssl"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  ssl_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    request            = "ARE YOU HEALTHY?"
    proxy_header       = "NONE"
    response           = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttpExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.http-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttpExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "http-health-check" {
  name = "tf-test-http-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttpFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttpFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.http-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttpFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "http-health-check" {
  name        = "tf-test-http-health-check%{random_suffix}"
  description = "Health check via http"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  http_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    host               = "1.2.3.4"
    request_path       = "/mypath"
    proxy_header       = "NONE"
    response           = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttpsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttpsExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.https-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttpsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "https-health-check" {
  name = "tf-test-https-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  https_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttpsFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttpsFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.https-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttpsFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "https-health-check" {
  name        = "tf-test-https-health-check%{random_suffix}"
  description = "Health check via https"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  https_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    host               = "1.2.3.4"
    request_path       = "/mypath"
    proxy_header       = "NONE"
    response           = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttp2Example(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttp2Example(context),
			},
			{
				ResourceName:      "google_compute_health_check.http2-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttp2Example(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "http2-health-check" {
  name = "tf-test-http2-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  http2_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttp2FullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttp2FullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.http2-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttp2FullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "http2-health-check" {
  name        = "tf-test-http2-health-check%{random_suffix}"
  description = "Health check via http2"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  http2_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    host               = "1.2.3.4"
    request_path       = "/mypath"
    proxy_header       = "NONE"
    response           = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckGrpcExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckGrpcExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.grpc-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckGrpcExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "grpc-health-check" {
  name = "tf-test-grpc-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  grpc_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckGrpcFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckGrpcFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.grpc-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckGrpcFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "grpc-health-check" {
  name = "tf-test-grpc-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  grpc_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    grpc_service_name  = "testservice"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckGrpcWithTlsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckGrpcWithTlsExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.grpc-with-tls-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckGrpcWithTlsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "grpc-with-tls-health-check" {
  provider = google-beta

  name = "tf-test-grpc-with-tls-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  grpc_tls_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckGrpcWithTlsFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckGrpcWithTlsFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.grpc-with-tls-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckGrpcWithTlsFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "grpc-with-tls-health-check" {
  provider = google-beta

  name        = "tf-test-grpc-with-tls-health-check%{random_suffix}"
  description = "Health check via grpc with TLS"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  grpc_tls_health_check {
    port_specification = "USE_FIXED_PORT"
    port = "443"
    grpc_service_name  = "testservice"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckWithLoggingExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckWithLoggingExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.health-check-with-logging",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckWithLoggingExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "health-check-with-logging" {
  provider = google-beta

  name = "tf-test-tcp-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  tcp_health_check {
    port = "22"
  }

  log_config {
    enable = true
  }
}
`, context)
}

func TestAccComputeHealthCheck_computeHealthCheckHttpSourceRegionsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_computeHealthCheckHttpSourceRegionsExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.http-health-check-with-source-regions",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_computeHealthCheckHttpSourceRegionsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "http-health-check-with-source-regions" {
  name = "tf-test-http-health-check%{random_suffix}"
  check_interval_sec = 30

  http_health_check {
    port = 80
    port_specification = "USE_FIXED_PORT"
  }

  source_regions = ["us-west1", "us-central1", "us-east5"]
}
`, context)
}

func TestAccComputeHealthCheck_computeHealthCheckHttpsSourceRegionsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_computeHealthCheckHttpsSourceRegionsExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.https-health-check-with-source-regions",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_computeHealthCheckHttpsSourceRegionsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "https-health-check-with-source-regions" {
  name = "tf-test-https-health-check%{random_suffix}"
  check_interval_sec = 30

  https_health_check {
    port = 80
    port_specification = "USE_FIXED_PORT"
  }

  source_regions = ["us-west1", "us-central1", "us-east5"]
}
`, context)
}

func TestAccComputeHealthCheck_computeHealthCheckTcpSourceRegionsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_computeHealthCheckTcpSourceRegionsExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.tcp-health-check-with-source-regions",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_computeHealthCheckTcpSourceRegionsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_health_check" "tcp-health-check-with-source-regions" {
  name = "tf-test-tcp-health-check%{random_suffix}"
  check_interval_sec = 30

  tcp_health_check {
    port = 80
    port_specification = "USE_FIXED_PORT"
  }

  source_regions = ["us-west1", "us-central1", "us-east5"]
}
`, context)
}

func testAccCheckComputeHealthCheckDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_health_check" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/healthChecks/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeHealthCheck still exists at %s", url)
			}
		}

		return nil
	}
}
