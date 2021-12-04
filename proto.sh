for file in proto/*.proto; do
  ./protoc -I proto $file --go_out=plugins=grpc:pb/
done