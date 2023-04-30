#!/bin/sh

test_and_coverage() {
  go test -v ./... -coverprofile coverage.out -covermode count
  go tool cover -func coverage.out
  echo "Quality Gate: checking test coverage is above threshold ..."
  echo "Threshold             : $TEST_COVERAGE_THRESHOLD %"
  totalCoverage=$(go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+')
  echo "Current test coverage : $totalCoverage %"
  if awk "BEGIN {exit !($totalCoverage >= $TEST_COVERAGE_THRESHOLD)}"; then
    echo "OK"
  else
    echo "Current test coverage is below threshold. Please add more unit tests or adjust threshold to a lower value."
    echo "Failed"
    exit 1
  fi
}

vet() {
  go vet ./...
}


"$@"