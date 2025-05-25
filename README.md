# DevOpsify the go web application

## Go Web Application
A simple website built using Golang and the `net/http` package.

### Running the Server
To start the server locally:
```bash
go run main.go
```
Access the website at: `http://localhost:8080/`

## DevOps Process
This project demonstrates how to apply DevOps practices to a go-webapp, including:
- Dockerization (multi-stage builds)
- CI with GitHub Actions
- CD with Argo CD on Kubernetes (EKS)

### Architecture Overview
![image](https://github.com/Akshata0211/go-webapp-k8s-cicd/blob/main/images/wrokflow.png)

###  Containerization with Docker

Use Docker to containerize the application for consistent builds and deployment.:

```bash
#Build the Docker image
docker build -t <your-docker-username>/go-web-app .

#Run the container
docker run -p 8080:8080 <your-docker-username>/go-web-app

#Push image to Docker Hub
docker push <your-docker-username>/go-web-app
```

### Continuous Integration (CI)

#### Implemented using GitHub Actions.

CI Workflow steps:
- Checkout source code
- Set up Go environment
- Run tests
- Build Docker image
- Push Docker image to Docker Hub

### Continuous Deployment (CD)

Argo CD for GitOps-based CD
Argo CD syncs Kubernetes apps from a Git repository and automatically deploys them to your EKS cluster.

#### Tool Installation (Required)
#Install the following tools on your system:
```bash
# AWS CLI
# kubectl
# eksctl
# helm
```

### Kubernetes Deployment with Helm

Create and deploy Helm chart:
```bash
# Install app using Helm
helm install go-web-app ./go-web-app-chart

# Uninstall app
helm uninstall go-web-app
```

### Migrate Kubernetes YAML to Helm:
##### Copy `k8s/manifests/*` into `helm/go-web-app-chart/templates/` and:
- Replace hardcoded Docker image tag with: {{ .Values.image.tag }}
- Update values.yaml accordingly

### EKS Cluster Setup

```bash
#Create an EKS cluster using eksctl:
eksctl create cluster --name demo-cluster --region us-east-1
# To delete:
eksctl delete cluster --name demo-cluster --region us-east-1
```
###  Kubernetes Management

```bash
# Apply all manifests
kubectl apply -f k8s/manifests/

# Get all resources
kubectl get all

# Delete everything
kubectl delete all --all
```

### NGINX Ingress on AWS

```bash
#Install Ingress controller:
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.11.1/deploy/static/provider/aws/deploy.yaml
kubectl get pods -n ingress-nginx
kubectl get svc -n ingress-nginx
```

###  Get Ingress LoadBalancer IP

```bash
#Run the following to retrieve the public IP/hostname:
nslookup <your-ingress-load-balancer-address>
```

#### To access your website domain:
- Map Domain Locally (Temporary) to test before DNS is fully configured:
- Open `/etc/hosts` on your local system
- Add an entry like:
```bash
<Ingress-IP>   mywebapp.example.com
#Replace `mywebapp.example.com` with your custom domain.
```
- Edit EC2 security group to allow NodePort traffic for testing
- Access via: `http://<eks-node-external-ip>:<nodeport>`

### Argo CD Setup

```bash
#Install Argo CD
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

#Expose Argo CD via LoadBalancer
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
kubectl get svc argocd-server -n argocd

#Get Argo CD Admin Password
kubectl get secret argocd-initial-admin-secret -n argocd -o yaml
# Decode password
echo <base64-password> | base64 --decode
```

### Connect Your Git Repo to Argo CD

##### In Argo CD UI:
- Create a new application
- Set sync policy to Automatic and enable Self-Heal
- Use the Git repo as source
- Provide path: helm/go-web-app-chart to point to your Helm chart

#### Notes

- Use CI to automatically update image tag in values.yaml on every push
- Argo CD picks up this tag and deploys it to EKS cluster
