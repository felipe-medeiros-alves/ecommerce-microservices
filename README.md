## E-commerce Microservices

### Getting Started
    
```bash
git clone https://github.com/felipe-medeiros-alves/e-commerce-backend.git
```

```bash
cd e-commerce-backend
go mod download
```

### Deployment

#### Build Docker Images:

```bash
docker-compose build
docker-compose up
```

#### Run with Docker Compose:

```bash
docker-compose up
```
#### Push Docker Images to Registry (Docker Hub):

```bash
docker tag auth-service your-dockerhub-username/auth-service:latest
docker push your-dockerhub-username/auth-service:latest
docker tag product-service your-dockerhub-username/product-service:latest
docker push your-dockerhub-username/product-service:latest
docker tag user-service your-dockerhub-username/user-service:latest
docker push your-dockerhub-username/user-service:latest
```

#### Deploy on Kubernetes:

```bash
kubectl apply -f k8s/auth-service.yaml
kubectl apply -f k8s/product-service.yaml
kubectl apply -f k8s/user-service.yaml
```
