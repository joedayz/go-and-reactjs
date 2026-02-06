# Terraform + GCP — Demo

1. Instalar [Terraform](https://www.terraform.io/downloads) y [Google Cloud SDK](https://cloud.google.com/sdk/docs/install).

2. Autenticación GCP:
   ```bash
   gcloud auth application-default login
   gcloud config set project TU_PROJECT_ID
   ```

3. Inicializar y aplicar:
   ```bash
   cd infra/terraform
   terraform init
   terraform plan -var="project_id=TU_PROJECT_ID" -var="db_password=TU_PASSWORD_SEGURA"
   terraform apply -var="project_id=TU_PROJECT_ID" -var="db_password=TU_PASSWORD_SEGURA"
   ```

4. Conectar el backend a Cloud SQL:
   - Usar la IP pública o Cloud SQL Proxy.
   - `DATABASE_URL=postgres://app:TU_PASSWORD@IP_PUBLICA:5432/demo?sslmode=require`

Para no dejar recursos activos (evitar costos):
   ```bash
   terraform destroy -var="project_id=TU_PROJECT_ID" -var="db_password=..."
   ```
