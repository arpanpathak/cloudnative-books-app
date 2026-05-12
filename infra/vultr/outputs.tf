output "cluster_id" {
  description = "The ID of the Vultr Kubernetes cluster"
  value       = vultr_kubernetes.k8s.id
}

output "kubeconfig_instructions" {
  description = "Instructions to connect to the cluster"
  value       = "Run: export KUBECONFIG=${path.module}/kubeconfig.yaml"
}
