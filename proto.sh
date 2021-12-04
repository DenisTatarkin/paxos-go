./protoc -I proto proto/proposal.proto --go_out=plugins=grpc:pb/
./protoc -I proto proto/phase_a.proto --go_out=plugins=grpc:pb/
./protoc -I proto proto/phase_b.proto --go_out=plugins=grpc:pb/