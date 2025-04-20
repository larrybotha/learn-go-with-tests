watch path:
  ( \
    cd {{ path }} && \
    watchexec --clear --exts go go test \
  )
