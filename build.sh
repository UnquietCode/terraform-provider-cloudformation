#!/bin/bash
set -ex

SCRIPT_DIR=$(cd $(dirname "${BASH_SOURCE[0]}") >/dev/null 2>&1 && pwd)

cd "$SCRIPT_DIR"

pushd "generator"
  python3 generate.py
popd

pushd "provider"
  go build
popd