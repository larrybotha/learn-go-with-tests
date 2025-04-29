watch-run path:
  watchexec --clear --exts go go run {{ path }}

watch-tests path:
  watchexec --clear --exts go just test {{ path }}

watch-checks path:
  watchexec --clear --exts go 'just vet {{ path }} && just errcheck {{ path }}'

vet path:
  @(cd {{ path }} && go vet)

errcheck path:
  @(cd {{ path }} && errcheck)

test path:
  @(cd {{ path }} && go test -v -cover)

bench path:
  # -bench -> run benchmarks
  # -benchmem -> show bytes allocated per iteration,
  #   and number of allocations per iteration
  (cd {{ path }} && go test -bench=. -benchmem)

