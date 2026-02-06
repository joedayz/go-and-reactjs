# Kubernetes — Demo

1. Construir imágenes y cargar en el cluster (ej. minikube):
   ```bash
   eval $(minikube docker-env)
   docker build -t demo-backend:latest ./backend
   docker build -t demo-frontend:latest ./frontend
   ```

2. Crear el secret para la DB (ajustar la URL):
   ```bash
   kubectl create secret generic demo-db-secret --from-literal=url='postgres://user:pass@host:5432/demo?sslmode=disable'
   ```

3. Aplicar manifests:
   ```bash
   kubectl apply -f backend-deployment.yaml
   kubectl apply -f frontend-deployment.yaml
   ```

4. Acceder al frontend (LoadBalancer):
   ```bash
   kubectl get svc demo-frontend
   ```
