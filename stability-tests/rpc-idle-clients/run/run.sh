#!/bin/bash
rm -rf /tmp/zuad-temp

NUM_CLIENTS=128
zuad --devnet --appdir=/tmp/zuad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
ZUAD_PID=$!
ZUAD_KILLED=0
function killZuadIfNotKilled() {
  if [ $ZUAD_KILLED -eq 0 ]; then
    kill $ZUAD_PID
  fi
}
trap "killZuadIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $ZUAD_PID

wait $ZUAD_PID
ZUAD_EXIT_CODE=$?
ZUAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Zuad exit code: $ZUAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ZUAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
