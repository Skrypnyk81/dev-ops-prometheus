module "gke_cluster" {
  source         = "github.com/Skrypnyk81/dev-ops-prometheus/tf-google-gke-cluster"
  GOOGLE_REGION  = var.GOOGLE_REGION
  GOOGLE_PROJECT = var.GOOGLE_PROJECT
  GKE_NUM_NODES  = 2
}
