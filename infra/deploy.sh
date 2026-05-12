#!/usr/bin/env bash

# Exit immediately if a command exits with a non-zero status
set -e

echo "🚀 Initializing Terraform..."
terraform init

echo "🏗️  Applying Terraform configuration (auto-approving)..."
terraform apply -auto-approve

echo "🔑 Extracting and saving Kubeconfig to kubeconfig.yaml..."
# Attempting macOS base64 decode (-D), fallback to standard Linux (-d) if it fails
terraform output -raw kubeconfig | base64 -D > kubeconfig.yaml 2>/dev/null || terraform output -raw kubeconfig | base64 -d > kubeconfig.yaml

echo "✅ Cluster deployed successfully!"
echo "============================================================"
echo "⚠️  IMPORTANT: To use kubectl in your current terminal session,"
echo "you must run the following command to set your KUBECONFIG:"
echo ""
echo "    export KUBECONFIG=\$(pwd)/kubeconfig.yaml"
echo ""
echo "After that, verify the connection by running:"
echo "    kubectl get nodes"
echo "============================================================"
