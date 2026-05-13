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

# 4. Merge kubeconfig automatically
echo "🔗 Merging Vultr cluster into ~/.kube/config..."
export KUBECONFIG=~/.kube/config:$(pwd)/kubeconfig.yaml
kubectl config view --flatten > ~/.kube/merged_config
mv ~/.kube/merged_config ~/.kube/config
unset KUBECONFIG
chmod 600 ~/.kube/config

echo ""
echo "✅ Deployment complete! Your ~/.kube/config is updated."
echo "To switch to your new Vultr cluster context, run:"
echo "kubectl config get-contexts"
echo "kubectl config use-context <vultr-context-name>"
