provider "aws" {}

data "aws_caller_identity" "current" {}

module "data-workflow" {
  source = "github.com/data-derp/terraform-modules//glue-workflow"

  project-name = var.project-name
  module-name = var.module-name
}