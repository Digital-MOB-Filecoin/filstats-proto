proto:
	protoc -I ./ --go_out=plugins=grpc:./ --go_opt=paths=source_relative ./*.proto
