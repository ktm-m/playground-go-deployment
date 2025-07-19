# Prerequisites

- Create `app-secret.yaml` in the `deployment` directory with the following content:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
  namespace: app
type: Opaque
stringData:
  POSTGRES_DB: "your_database_name"
  POSTGRES_USER: "your_database_user"
  POSTGRES_PASSWORD: "your_database_password"
  REDIS_PASSWORD: "your_redis_password"
```

---

# To run the application

1. Build the Docker image and push it to your container registry:

```bash
docker build -t your_docker_registry/your_image_name:latest .
docker push your_docker_registry/your_image_name:latest
```

2. Edit the `backend-deployment.yaml` file to set the correct image name.

```yaml
image: your_image_name:latest
```

3. Apply the Kubernetes manifests (make sure your Kubernetes cluster is set up and kubectl is configured):

```bash
kubectl apply -k ./deployment/
```

4. Verify that the pods are running:

```bash
kubectl get pods -n app
```

5. Test backend health check endpoint via nginx:

```bash
kubectl port-forward svc/app-nginx 8081:80 -n app
curl http://localhost:8081/api/health
```
