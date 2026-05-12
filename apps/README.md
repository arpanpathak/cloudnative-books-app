# 📦 Microservices & Applications

This directory houses the source code, containerization logic, and Kubernetes orchestration manifests for the microservices in this repository.

## 📚 Services

### [Books API](./books-api)
A high-performance Go REST API demonstrating:
- Thread-safe, collection-oriented REST patterns.
- Multi-stage, zero-vulnerability `distroless` Docker builds.
- Kubernetes Horizontal Pod Autoscaler (HPA) for dynamic load management.
- Cloud Load Balancer integration for public ingress.

Please see the [Books API README](./books-api/README.md) for detailed instructions on building, pushing, deploying, and load testing the service.
