# CO2 vs. Temperature (Infrastructure)
In this exercise, we assume that you've completed the [production-code](https://github.com/data-derp/exercise-co2-vs-temperature-production-code) and have managed to push those artifacts to an AWS S3 Bucket. If not, see [Fresh Start](#fresh-start). This exercise focuses on using those artifacts as part of an AWS Glue workflow.

**NOTE:** The following exercises follow the same concept as [the production-code exercise](https://github.com/data-derp/exercise-co2-vs-temperature-production-code) where a `project-name` and `module-name` are used consistently to create resources. 

In all examples in ,
* `project-name` = **awesome-project**
* `module-name` = **awesome-module**

Where these are used, you'll want to pick your own unique `project-name` and `module-name`.

## Prerequisites
* An AWS Account and IAM User with permissions to create AWS Glue and Athena resources and read S3 buckets
* AWS CLI access

## Quickstart
1. If you don't already have artifacts for Data Ingestion and Transformation in your S3 bucket follow the [Fresh Start instructions](#fresh-start)
2. Create AWS Resources for [Data Ingestion](./data-ingestion.md)
3. Create AWS Resources for [Data Transformation](./data-transformation.md)
4. Create AWS Resources for [Data Workflow](./data-workflow.md)
5. Create AWS Resources via [Infrastructure as Code](./infrastucture-as-code.md)

## Fresh Start
If you don't have the artifacts
1. [Create an S3 bucket](https://github.com/data-derp/s3-bucket-aws-cloudformation)
2. Upload the artifacts ([ensure you have an active AWS CLI Session](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html))
    ```bash
    # Change these variables
    PROJECT_NAME=awesome-project
    MODULE_NAME=awesome-module
   
    ./go upload-artifacts "${PROJECT_NAME}-${MODULE_NAME}"
    ```

   ```