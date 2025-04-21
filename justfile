watch path:
  watchexec --exts go just test {{ path }}

test path:
  (cd {{ path }} && go test -v)

bench path:
  # -bench -> run benchmarks
  # -benchmem -> show bytes allocated per iteration,
  #   and number of allocations per iteration
  (cd {{ path }} && go test -bench=. -benchmem)

