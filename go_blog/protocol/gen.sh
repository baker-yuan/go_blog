#trpc create -p template/template.proto -o out

# https://github.com/trpc-group/trpc-cmdline/blob/483259c22a2fd1e477c7ebc4b5b1572a5a47010c/README.zh_CN.md
# https://github.com/trpc-group/trpc-cmdline/blob/483259c22a2fd1e477c7ebc4b5b1572a5a47010c/docs/examples/example-2/README.zh_CN.md
# https://github.com/trpc-group/trpc-cmdline/blob/483259c22a2fd1e477c7ebc4b5b1572a5a47010c/docs/README.zh_CN.md
# https://github.com/bufbuild/protoc-gen-validate/blob/main/validate/validate.proto
# https://github.com/bufbuild/protoc-gen-validate/blob/v1.0.2/README.md


#find ./template -type f ! -name 'template.proto' -exec rm -f {} +
#trpc create --protofile=template/template.proto --validate=true --rpconly -o template


#find ./blog -type f ! -name 'blog.proto' -exec rm -f {} +
#trpc create --protofile=blog/blog.proto --validate=true --rpconly --alias --swagger -o blog

#find ./auth -type f ! -name 'auth.proto' -exec rm -f {} +
#trpc create --protofile=auth/auth.proto --validate=true --rpconly --alias --swagger -o auth
#
#find ./user -type f ! -name 'user.proto' -exec rm -f {} +
#trpc create --protofile=user/user.proto --validate=true --rpconly --alias --swagger -o user

find ./datasync -type f ! -name 'datasync.proto' -exec rm -f {} +
trpc create --protofile=datasync/datasync.proto --validate=true --rpconly --alias -o datasync
