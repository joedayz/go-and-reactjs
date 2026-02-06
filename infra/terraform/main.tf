# Terraform + GCP — Demo para perfil
# Requisitos: cuenta GCP, gcloud auth application-default login, terraform >= 1.0

terraform {
  required_version = ">= 1.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

variable "project_id" {
  description = "GCP Project ID"
  type        = string
}

variable "region" {
  description = "GCP region"
  type        = string
  default     = "us-central1"
}

provider "google" {
  project = var.project_id
  region  = var.region
}

# Cloud SQL (PostgreSQL) — requisito "Bases de datos: PostgreSQL, MySQL o Cloud SQL"
resource "google_sql_database_instance" "demo" {
  name             = "demo-profile-db"
  database_version = "POSTGRES_15"
  region           = var.region

  settings {
    tier = "db-f1-micro" # Cambiar a db-g1-small en producción

    ip_configuration {
      ipv4_enabled    = true
      private_network = null
    }

    backup_configuration {
      enabled                        = true
      start_time                     = "03:00"
      point_in_time_recovery_enabled = false
    }
  }

  deletion_protection = false
}

resource "google_sql_database" "app" {
  name     = "demo"
  instance = google_sql_database_instance.demo.name
}

resource "google_sql_user" "app" {
  name     = "app"
  instance = google_sql_database_instance.demo.name
  password = var.db_password
}

variable "db_password" {
  description = "Password for Cloud SQL user"
  type        = string
  sensitive   = true
}

# Opcional: GKE cluster para Kubernetes (requisito K8s + GCP)
# Descomentar si quieres demostrar GKE
# resource "google_container_cluster" "demo" {
#   name     = "demo-profile-cluster"
#   location = var.region
#
#   initial_node_count = 1
#   node_config {
#     machine_type = "e2-micro"
#   }
# }
