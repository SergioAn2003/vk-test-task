#!/bin/bash
export ROOT=../../..
source variables.sh
export DEBUG=true
export DEV=true

go clean -testcache
cd ../internal/
echo '-- тест usecase --'
for s in $(go list ./usecase/test/...); do if ! go test -failfast -p 1 $s; then break; fi; done 2>&1 | grep -v '/usr/bin/ld: /usr/lib/oracle/11.2/client64/lib/libnnz11.so'
echo '-- тест repository --'
export CONF_PATH=../../../../../config/conf.yaml
for s in $(go list ./repository/oracle/test/...); do if ! go test -failfast -p 1 $s; then break; fi; done 2>&1 | grep -v '/usr/bin/ld: /usr/lib/oracle/11.2/client64/lib/libnnz11.so'