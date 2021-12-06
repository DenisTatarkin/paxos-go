module paxos_go

require (
	github.com/ahmetb/go-linq/v3 v3.2.0
	github.com/cristalhq/acmd v0.4.0
	google.golang.org/grpc v1.24.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20190311183353-d8887717615a // indirect
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.27.0

go 1.17
