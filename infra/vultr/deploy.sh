#!/bin/bash
set -e

echo "=== Vultr Kubernetes Deployment ==="

# 1. Ensure API Key is provided via terraform.tfvars
echo "Looking for API Key in terraform.tfvars..."
# 2. Initialize Terraform
echo "📦 Initializing Terraform..."
terraform init

# 3. Deploy
echo "🚀 Deploying Vultr Kubernetes Engine (VKE)..."
terraform apply -auto-approve

# 4. Success message
echo ""
echo "✅ Deployment complete!"
echo "To connect to your new Vultr cluster, run:"
echo "export KUBECONFIG=\$(pwd)/kubeconfig.yaml"
echo "kubectl get nodes"
