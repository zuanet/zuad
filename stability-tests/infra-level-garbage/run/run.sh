#!/bin/bash
rm -rf /tmp/zuad-temp

zuad --devnet --appdir=/tmp/zuad-temp --profile=6061 &
ZUAD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:46609 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $ZUAD_PID

wait $ZUAD_PID
ZUAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Zuad exit code: $ZUAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ZUAD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
