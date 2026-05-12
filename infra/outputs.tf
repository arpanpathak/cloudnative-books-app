output "kubeconfig" {
  value     = linode_lke_cluster.kubegoat.kubeconfig
  sensitive = true
}

output "api_endpoints" {
  value = linode_lke_cluster.kubegoat.api_endpoints
}

output "id" {
  value = linode_lke_cluster.kubegoat.id
}

