package test

import (
	"fmt"
	"testing"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/docker"
)

func TestDockerTask(t *testing.T) {

	var tag string = "day-two-task/v0"
	var container_name string = "day-two-task"

	// Build the Dockerfile
	buildOptions := &docker.BuildOptions{Tags: []string{tag}}
	docker.Build(t, "../app", buildOptions)
	fmt.Println("Build Successful")

	// Run the container
	opts := &docker.RunOptions{Entrypoint: "", Command: []string{"run", "-p", "5002"}, EnvironmentVariables: []string{"FLASK_APP=main.py"} , Name: container_name, Remove: true, OtherOptions: []string{"-d"}}
	docker.Run(t, tag, opts)


	stopOpts := &docker.StopOptions{
		Time:   0,
		Logger: &logger.Logger{},
	}
	docker.Stop(t, []string{container_name}, stopOpts)
	fmt .Println("Run & Stop Successful")
}
