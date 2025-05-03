watch-run path:
  watchexec --clear --timings --exts go --filter=justfile \
    go run {{ path }}

watch-tests path:
  watchexec --clear --timings --exts go --filter=justfile \
    'just test {{ path }}; just bench {{ path }}'

watch-checks path:
  watchexec --clear --timings --exts go --filter=justfile \
    'just vet {{ path }} && just errcheck {{ path }}'

vet path:
  @(cd {{ path }} && go vet)

errcheck path:
  @(cd {{ path }} && errcheck)

test path:
  @(cd {{ path }} && go test -race -v -cover)

# -bench -> run benchmarks
# -benchmem -> show bytes allocated per iteration,
#   and number of allocations per iteration
bench path:
  (cd {{ path }} && go test -race -bench=. -benchmem)

