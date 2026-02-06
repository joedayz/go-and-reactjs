output "cloud_sql_connection_name" {
  description = "Cloud SQL instance connection name (for Cloud SQL Proxy or private IP)"
  value       = google_sql_database_instance.demo.connection_name
}

output "cloud_sql_public_ip" {
  description = "Public IP of Cloud SQL (if ipv4_enabled)"
  value       = google_sql_database_instance.demo.public_ip_address
}
