#!/bin/bash
rm -rf /tmp/zuad-temp

zuad --devnet --appdir=/tmp/zuad-temp --profile=6061 --loglevel=debug &
ZUAD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $ZUAD_PID

wait $ZUAD_PID
ZUAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Zuad exit code: $ZUAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ZUAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
