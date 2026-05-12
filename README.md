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

## 🔮 Future Roadmap
The vision is to be **cloud native from day 1**. Future plans include decoupling cross-cutting concerns by introducing:
- **Gateways & Ingress Controllers**
- **TLS Termination & mTLS**
- **AuthN & AuthZ**
- **Firewalls & NetworkPolicies**
- **Service Mesh & eBPF**
- **Custom Kubernetes Operators (Go)**: Replacing bash-based Redis initialization with a native Go Operator to automatically handle shard formulation, leader elections, and zero-downtime failovers natively inside the cluster.

## 🧠 Honest Distributed Systems Engineering
Distributed Systems are full of trade-offs. As a Senior Engineer, your responsibility goes beyond just coding, especially if you join startups with budget constraints. Big tech companies have unlimited money, but small and mid-size tech companies need to be frugal, and the scale is massive as well. 

Designing systems tied to a single cloud provider's proprietary tech, and not being able to move fast and adopt cutting-edge tech, poses a high risk in this free market economy.

One of the most useful transferable skills you can learn is to harness the power of Linux, Kubernetes, and projects created by the Cloud Native Computing Foundation (CNCF). The first step towards a cloud-native solution is to harness the power of decoupling cross-cutting concerns from your business logic codebase. Your business shouldn't suffer if one cloud provider goes bankrupt, shuts down, or starts charging aggressively. Design systems such that you can do a "lift & shift" flip-switch without making the lives of developers miserable.

### 🛠️ Learning by Doing (The Bare Metal Way)
If you truly want to get your hands dirty, buy an AMD Ryzen Mini PC with an NVMe SSD and install **Talos Linux**. There is no better way to learn than by implementing code that runs directly on bare metal.

If you want to build scalable database engines, master network programming, and explore advanced areas of systems engineering, I highly recommend reading:
1. *Designing Data-Intensive Applications* (Martin Kleppmann)
2. *Database Internals* (Alex Petrov)
3. *Understanding Distributed Systems* (Roberto Vitillo)

## 📜 License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
