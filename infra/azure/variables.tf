variable "resource_group_name" {
  type    = string
  default = "rg-kubegoat-aks"
}

variable "location" {
  type    = string
  default = "eastus"
}

variable "cluster_name" {
  type    = string
  default = "kubegoat-aks"
}

variable "node_count" {
  type    = number
  default = 2
}

variable "vm_size" {
  type    = string
  # Standard_D2s_v3 gives you 2 vCPUs and 8GB of RAM per node. 
  # This is the "sweet spot" for running Istio and heavy workloads.
  default = "Standard_D2s_v3"
}
