#!/bin/bash
set -ex

SCRIPT_DIR=$(cd $(dirname "${BASH_SOURCE[0]}") >/dev/null 2>&1 && pwd)

cd "$SCRIPT_DIR"

pushd "provider"
  go build -o ../terraform-provider-cfn
popd
