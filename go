#!/bin/bash

set -e
script_dir=$(cd "$(dirname "$0")" ; pwd -P)

goal_upload-data-source() {
    bucket_name="${1}"
    if [ -z "${bucket_name}" ]; then
      echo "BUCKET_NAME not supplied. Usage <func> bucket_name"
      exit 1
    fi

    ingestion_input_dir="${script_dir}/inputs/data-source"
    mkdir -p ${ingestion_input_dir}
    curl -o "${ingestion_input_dir}/EmissionsByCountry.csv" "https://raw.githubusercontent.com/data-derp/exercise-co2-vs-temperature-databricks/master/data-ingestion/input-data/EmissionsByCountry.csv"
    curl -o "${ingestion_input_dir}/GlobalTemperatures.csv" "https://raw.githubusercontent.com/data-derp/exercise-co2-vs-temperature-databricks/master/data-ingestion/input-data/GlobalTemperatures.csv"
    curl -o "${ingestion_input_dir}/TemperaturesByCountry.csv" "https://raw.githubusercontent.com/data-derp/exercise-co2-vs-temperature-databricks/master/data-ingestion/input-data/TemperaturesByCountry.csv"

    aws s3 cp ${ingestion_input_dir} s3://${bucket_name}/data-source/ --recursive
}

goal_upload-artifacts() {
    bucket_name="${1}"
    if [ -z "${bucket_name}" ]; then
      echo "BUCKET_NAME not supplied. Usage <func> bucket_name"
      exit 1
    fi

   aws s3 cp ${script_dir}/inputs/data-ingestion/ s3://${bucket_name}/data-ingestion-artifacts/ --recursive
   aws s3 cp ${script_dir}/inputs/data-transformation/ s3://${bucket_name}/data-transformation-artifacts/ --recursive
}

TARGET=${1:-}
if type -t "goal_${TARGET}" &>/dev/null; then
  "goal_${TARGET}" ${@:2}
else
  echo "Usage: $0 <goal>

goal:
    upload-data-source                    - Fetches and uploads data-source files to AWS S3 bucket
    upload-artifacts                      - Uploads artifact files to AWS S3 bucket
"
  exit 1
fi
