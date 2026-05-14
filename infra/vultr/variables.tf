variable "region" {
  description = "The Vultr region to deploy the cluster in (e.g., ewr, sjc, sea, blr)"
  type        = string
  default     = "sea" # Seattle region for low latency.
}

variable "cluster_name" {
  description = "Name of the Vultr Kubernetes cluster"
  type        = string
  default     = "kubegoat-vultr"
}

variable "k8s_version" {
  description = "The Kubernetes version to install"
  type        = string
  default     = "v1.35.2+1" 
}

variable "node_count" {
  description = "Number of worker nodes"
  type        = number
  default     = 3
}

variable "plan" {
  description = "The Vultr instance plan (vc2-4c-8gb = 4vCPU / 8GB RAM)"
  type        = string
  default     = "vc2-4c-8gb" 
}

variable "vultr_api_key" {
  description = "The Vultr API Key"
  type        = string
  sensitive   = true
}
