# 🌍 Infrastructure as Code (IaC)

This directory contains the Terraform modules required to automatically provision Kubernetes clusters across various cloud providers.

## ☁️ Supported Cloud Providers

| Provider | Description | Path |
|----------|-------------|------|
| **Linode** | Linode Kubernetes Engine (LKE) clusters. | [`/`](./) |
| **Azure** | Azure Kubernetes Service (AKS) clusters. | [`/azure`](./azure) |
| **Vultr** | Vultr Kubernetes Engine (VKE) clusters. | [`/vultr`](./vultr) |

## 🛠️ Usage
Each cloud provider directory contains a self-contained Terraform module and a convenient `deploy.sh` wrapper script. 

1. Navigate to the cloud provider of your choice.
2. Export your respective Cloud API Key.
3. Run the deployment script.

```bash
cd vultr
./deploy.sh
```

> **Security Note:** All credentials, API keys, Terraform State (`.tfstate`), and generated `kubeconfig` files are strictly ignored via `.gitignore` to prevent accidental leakage.
