#!/bin/sh
benchstat -delta-test none testdata/bench_before/${1} testdata/bench_after/${1} | tee testdata/bench_results/${1}
