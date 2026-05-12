# ☁️ Vultr Kubernetes Engine (VKE) IaC

This module provisions a managed Kubernetes cluster on Vultr using Terraform.

## 🚀 Deployment

1. **Set API Key:**
   Ensure you have generated a Vultr API Key and whitelisted your IP address in the Vultr API settings.
   ```bash
   export VULTR_API_KEY="your-api-key"
   ```

2. **Run Deployment:**
   Execute the wrapper script to initialize and apply the Terraform plan.
   ```bash
   ./deploy.sh
   ```
   *The script will automatically download and decode the `kubeconfig.yaml` for you.*

## 🧹 Cleanup
To stop billing, destroy the cluster when finished:
```bash
terraform destroy -auto-approve
```
