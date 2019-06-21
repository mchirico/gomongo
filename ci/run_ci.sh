#!/bin/bash

set -eu

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
export fly_target=${fly_target:-mce}
echo "Concourse API target ${fly_target}"
echo "Tutorial $(basename $DIR)"

pushd $DIR
  fly -t ${fly_target} set-pipeline -p gomongo-pipeline -c build-golang-pipeline.yml -n -v mongo_connection_string=$MONGO_CONNECTION_STRING -v mongo_database=$MONGO_DATABASE
  fly -t ${fly_target} unpause-pipeline -p gomongo-pipeline
#  fly -t ${fly_target} trigger-job -w -j tutorial-pipeline/job-hello-world
popd

echo -e "\n\n                  Common commands:"
echo -e "**************************************\n\n"
echo -e "\n"
echo -e "                           fly -t mce watch --job gomongo-pipeline/unit"
echo -e "                           fly -t mce builds|grep 'gomongo-pipeline'"
echo -e "                           fly -t mce destroy-pipeline -p gomongo-pipeline -n"
echo -e "                           fly -t mce workers -d "
echo -e "                            "
echo -e "                           To login to a running container: "
echo -e "                           fly -t mce intercept --job gomongo-pipeline/unit "
echo -e "\n"
echo -e "\n"

