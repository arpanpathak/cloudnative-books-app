#!/bin/bash
set -e

echo "=== Vultr Kubernetes Deployment ==="

# 1. Ensure API Key is provided
if [ -z "$VULTR_API_KEY" ]; then
    echo "🚨 VULTR_API_KEY environment variable is not set!"
    echo "Go to https://my.vultr.com/settings/#settingsapi to generate an API key."
    echo "IMPORTANT: Make sure to 'Allow All IPv4' or whitelist your specific IP on Vultr's API page!"
    read -p "Enter your Vultr API Key: " VULTR_API_KEY
    export VULTR_API_KEY
fi

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
