variable "linode_token" {
  type      = string
  sensitive = true
}

variable "region" {
  type    = string
  default = "us-sea"
}

variable "cluster_label" {
  type    = string
  default = "kubegoat"
}

variable "node_type" {
  type    = string
  default = "g6-standard-1"
}
