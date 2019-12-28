locals {
  docker_image = "gcr.io/${var.gcp_project}/${var.context_name}:${var.image_tag}"
}

resource "google_cloud_run_service" "cloud_run" {
  name     = "cloudrun-test"
  location = "us-central1"

  template {
    spec {
      containers {
        image = local.docker_image
        env {
          name  = "DJANGO_SETTINGS_MODULE"
          value = var.django_settings_module
        }
      }
    }
  }
}

data "google_iam_policy" "cloud_run_policy" {
  binding {
    role    = "roles/run.invoker"
    members = ["allUsers"]
  }
}

resource "google_cloud_run_service_iam_policy" "cloud_run_policy" {
  location    = google_cloud_run_service.cloud_run.location
  project     = google_cloud_run_service.cloud_run.project
  service     = google_cloud_run_service.cloud_run.name
  policy_data = data.google_iam_policy.cloud_run_policy.policy_data
}
