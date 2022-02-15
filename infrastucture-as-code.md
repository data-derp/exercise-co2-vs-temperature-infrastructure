## Infrastructure as Code (IAC)
Sometimes clicking through the UI is challenging and prone to mistakes. Use Terraform and custom modules to apply the same changes!

1. [Setup](#setup)
2. [Data Ingestion IAC](#data-ingestion-iac)
3. [Data Transformation IAC](#data-transformation-iac)
4. [Data Workflow IAC](#data-workflow-iac)
5. [Destroy everything](#destroy-everything)

## Setup
1. [Install dependencies](#install-depdenencies)
2. [Ensure you have an active AWS CLI Session](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html)


### Install Depdenencies
#### Non-Containerised
1. Install Python (version 3.7), for example with [Pyenv](https://github.com/pyenv/pyenv)
2. [Install Terraform](https://www.terraform.io/downloads) (required version 1.1.3)

#### Containerised
```bash
docker run -it -v $(pwd):/app ghcr.io/data-derp/python-aws-terraform-container:master bash
````
All workloads can be run within this container once you have [set up an active AWS CLI Session](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html)

## Data Ingestion IAC
```bash
cd iac/data-ingestion

# Change these variables
export PROJECT_NAME=awesome-project
export MODULE_NAME=awesome-module
export AWS_DEFAULT_REGION=eu-central-1

terraform init && terraform apply -var "project-name=${PROJECT_NAME}" -var "module-name=${MODULE_NAME}" -auto-approve
```

## Data Transformation IAC
```bash
cd iac/data-transformation

# Change these variables
export PROJECT_NAME=awesome-project
export MODULE_NAME=awesome-module
export AWS_DEFAULT_REGION=eu-central-1

terraform init && terraform apply -var "project-name=${PROJECT_NAME}" -var "module-name=${MODULE_NAME}" -auto-approve
```

## Data Workflow IAC
```bash
cd iac/data-workflow

# Change these variables
export PROJECT_NAME=awesome-project
export MODULE_NAME=awesome-module
export AWS_DEFAULT_REGION=eu-central-1

terraform init && terraform apply -var "project-name=${PROJECT_NAME}" -var "module-name=${MODULE_NAME}" -auto-approve
```
Run the Workflow in the AWS Console and watch it turn green!

## Destroy Everything
To destroy all of the resources you created using Terraform, run this from the root of the repository:
```bash
export AWS_DEFAULT_REGION=eu-central-1

for p in data-workflow data-transformation data-ingestion
do
 pushd "./iac/${p}" > /dev/null
    terraform destroy -auto-approve
 popd > /dev/null
done
```
    
