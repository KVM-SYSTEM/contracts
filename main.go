package main

import "os/exec"

func main() {
	// protoc -I proto/core_server proto/core_server/material.proto --go_out=./gen/core_server --go-grpc_out=./gen/core_server --go-grpc_opt=paths=source_relative

	cmd := exec.Command(
		"protoc",
		"-I", "proto/core_server",
		"proto/core_server/materials.proto",
		"--go_out=./gen",
		"--go-grpc_out=./gen",
		"--go-grpc_opt=paths=source_relative",
	)

	out, err := cmd.CombinedOutput()

	if err != nil {
		println(string(out))
		println(err.Error())
		return
	}

	cmd = exec.Command(
		"protoc",
		"-I", "proto/estimate_server",
		"proto/estimate_server/estimates.proto",
		"--go_out=./gen",
		"--go-grpc_out=./gen",
		"--go-grpc_opt=paths=source_relative",
	)

	out, err = cmd.CombinedOutput()

	if err != nil {
		println(string(out))
		println(err.Error())
		return
	}

	println(string(out))
}
