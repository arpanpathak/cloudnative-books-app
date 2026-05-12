# Cloud-Native Books App 🚀

A multi-cloud, cloud-native book server app to demonstrate the true power of Kubernetes.

This repository serves as a comprehensive masterclass in modern Platform Engineering, showcasing end-to-end deployment of a Go microservice across multiple cloud providers. 

## 🏗️ Architecture
- **Microservices**: A high-performance, thread-safe Go REST API.
- **Containerization**: Multi-stage Distroless Docker builds for ultimate security and minimal footprint.
- **Kubernetes**: Production-grade manifests featuring Deployments, LoadBalancers, and Horizontal Pod Autoscaling (HPA).
- **Infrastructure as Code (IaC)**: Automated multi-cloud cluster provisioning using Terraform.

## 📂 Repository Structure

| Directory | Description |
|-----------|-------------|
| [`/apps`](./apps) | Application source code, Dockerfiles, and Kubernetes manifests. |
| [`/infra`](./infra) | Terraform modules for automated cluster provisioning across different clouds. |

## 🚀 Quick Start
1. **Provision Infrastructure**: Navigate to `/infra` and choose your preferred cloud provider (Linode, Azure, Vultr) to spin up a Kubernetes cluster.
2. **Deploy Application**: Navigate to `/apps/books-api` to build the container and deploy the manifests to your newly created cluster.

## 📜 License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
