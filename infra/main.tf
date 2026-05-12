resource "linode_lke_cluster" "kubegoat" {
  label       = var.cluster_label
  k8s_version = "1.35"
  region      = var.region

  pool {
    type  = var.node_type
    count = 2
  }
}

