# https://github.com/trpc-group/trpc-cmdline/blob/483259c22a2fd1e477c7ebc4b5b1572a5a47010c/README.zh_CN.md
# https://github.com/trpc-group/trpc-cmdline/blob/483259c22a2fd1e477c7ebc4b5b1572a5a47010c/docs/examples/example-2/README.zh_CN.md
# https://github.com/bufbuild/protoc-gen-validate/blob/main/validate/validate.proto
# https://github.com/bufbuild/protoc-gen-validate/blob/v1.0.2/README.md
# https://github.com/trpc-group/trpc/blob/main/trpc/validate/validate.proto

find ./blog/v1 -type f ! -name 'blog.proto' -exec rm -f {} +
trpc create --protofile=blog/v1/blog.proto --rpconly --validate --nogomod -o blog/v1
