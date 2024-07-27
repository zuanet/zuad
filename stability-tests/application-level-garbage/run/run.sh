#!/bin/bash
rm -rf /tmp/zuad-temp

zuad --devnet --appdir=/tmp/zuad-temp --profile=6061 --loglevel=debug &
ZUAD_PID=$!
ZUAD_KILLED=0
function killZuadIfNotKilled() {
    if [ $ZUAD_KILLED -eq 0 ]; then
      kill $ZUAD_PID
    fi
}
trap "killZuadIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:46609 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $ZUAD_PID

wait $ZUAD_PID
ZUAD_KILLED=1
ZUAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Zuad exit code: $ZUAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ZUAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
