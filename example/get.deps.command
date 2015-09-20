#!/usr/bin/env bash

set -e

here="$(dirname "$BASH_SOURCE")"
cd $here
source env.sh
cd _vendor
go get -d github.com/cloudimmunity/pdiscover


