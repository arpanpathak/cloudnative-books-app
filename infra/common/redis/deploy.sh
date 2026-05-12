#!/bin/bash
set -e

echo "=== Deploying Raw K8s Redis Cluster ==="
kubectl create namespace redis --dry-run=client -o yaml | kubectl apply -f -

# 1. Ask for password dynamically (Never stored in Git)
if ! kubectl get secret redis-auth -n redis >/dev/null 2>&1; then
    read -s -p "Enter a secure password for the Redis Cluster: " REDIS_PASSWORD
    echo ""
    # Creates the secret securely inside K8s directly from memory
    kubectl create secret generic redis-auth -n redis --from-literal=password="$REDIS_PASSWORD"
    echo "✅ Secret created securely inside Kubernetes!"
else
    echo "✅ Secret 'redis-auth' already exists in the cluster."
fi

# 2. Apply YAMLs
echo "🚀 Applying StatefulSets and Services..."
kubectl apply -f 01-configmap.yaml
kubectl apply -f 02-services.yaml
kubectl apply -f 03-statefulset.yaml

# 3. Wait for Pods
echo "⏳ Waiting for all 6 Redis pods to start (this takes a minute)..."
kubectl rollout status statefulset/redis-cluster -n redis

# 4. Initialize the Cluster Shards
echo "🔗 Extracting Pod IPs for Shard Initialization..."
POD_IPS=$(kubectl get pods -n redis -l app=redis-cluster -o jsonpath='{range.items[*]}{.status.podIP}:6379 {end}')

echo ""
echo "🔥 IMPORTANT FINAL STEP 🔥"
echo "To actually form the sharded cluster, run this command using the password you set:"
echo "kubectl exec -it redis-cluster-0 -n redis -- redis-cli -a <YOUR_PASSWORD> --cluster create $POD_IPS --cluster-replicas 1"
