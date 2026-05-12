resource "vultr_kubernetes" "k8s" {
  region  = var.region
  label   = var.cluster_name
  version = var.k8s_version

  node_pools {
    node_quantity = var.node_count
    plan          = var.plan
    label         = "default-pool"
    auto_scaler   = false
  }
}

# Vultr returns the kubeconfig as a base64 encoded string.
# This resource automatically decodes it and saves it locally.
resource "local_sensitive_file" "kubeconfig" {
  depends_on = [vultr_kubernetes.k8s]
  filename   = "${path.module}/kubeconfig.yaml"
  content    = base64decode(vultr_kubernetes.k8s.kube_config)
}
