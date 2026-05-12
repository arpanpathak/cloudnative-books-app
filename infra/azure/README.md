# ☁️ Azure Kubernetes Service (AKS) IaC

This module provisions a managed Kubernetes cluster on Microsoft Azure using Terraform.

## 🚀 Deployment

1. **Authenticate with Azure:**
   You must be logged in via the Azure CLI before running Terraform.
   ```bash
   az login --scope https://graph.microsoft.com/.default
   ```

2. **Run Deployment:**
   Execute the wrapper script to initialize and apply the Terraform plan.
   ```bash
   ./deploy.sh
   ```

## 🧹 Cleanup
To avoid recurring charges, destroy the infrastructure when finished:
```bash
terraform destroy -auto-approve
```
