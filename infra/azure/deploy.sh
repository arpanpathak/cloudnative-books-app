#!/usr/bin/env bash
set -e

SUBSCRIPTION_ID="83965f97-3369-41e6-aff5-763219a67144"

echo "☁️  Verifying Azure CLI..."
if ! command -v az &> /dev/null; then
    echo "❌ Azure CLI is not installed. Please install it using: brew install azure-cli"
    exit 1
fi

echo "🔐 Setting Azure Subscription context..."
# Check if we have a valid access token, otherwise trigger az login
if ! az account get-access-token > /dev/null 2>&1; then
    echo "Azure token expired. Opening browser to log you in..."
    az login --scope https://graph.microsoft.com/.default
fi

# This ensures any manual 'az' commands you run also use the right subscription
az account set --subscription "$SUBSCRIPTION_ID"

echo "🚀 Initializing Terraform..."
terraform init

echo "🏗️  Applying Terraform configuration (this will take a few minutes for Azure)..."
terraform apply -auto-approve

echo "✅ Azure AKS Cluster deployed successfully!"
echo "============================================================"
echo "🔗 Automatically adding cluster to your ~/.kube/config..."
az aks get-credentials --resource-group "rg-kubegoat-aks" --name "kubegoat-aks" --overwrite-existing

echo "============================================================"
echo "🎉 All Done! The Azure cluster is live and configured."
echo "You can now use 'kubectx' to switch back and forth between"
echo "this cluster and your Linode cluster!"
echo ""
echo "Verify connection by running:"
echo "    kubectl get nodes"
echo "============================================================"
