#!/bin/bash

APPDIR=/tmp/zuad-temp
ZUAD_RPC_PORT=29587

rm -rf "${APPDIR}"

zuad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${ZUAD_RPC_PORT}" --profile=6061 &
ZUAD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${ZUAD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $ZUAD_PID

wait $ZUAD_PID
ZUAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Zuad exit code: $ZUAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ZUAD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
