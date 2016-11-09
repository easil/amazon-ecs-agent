// +build functional,windows

// Copyright 2014-2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package simpletest is an auto-generated set of tests defined by the json
// descriptions in testdata/simpletests.
//
// This file should not be edited; rather you should edit the generator instead
package simpletest

import (
	"os"
	"testing"
	"time"

	. "github.com/aws/amazon-ecs-agent/agent/functional_tests/util"
)

// TestDataVolume Check that basic data volumes work
func TestDataVolume(t *testing.T) {
	// Parallel is opt in because resource constraints could cause test failures
	// on smaller instances
	if os.Getenv("ECS_FUNCTIONAL_PARALLEL") != "" {
		t.Parallel()
	}
	agent := RunAgent(t, nil)
	defer agent.Cleanup()
	agent.RequireVersion(">=1.0.0")

	testTask, err := agent.StartTask(t, "datavolume-windows")
	if err != nil {
		t.Fatal("Could not start task", err)
	}
	timeout, err := time.ParseDuration("2m")
	if err != nil {
		t.Fatal("Could not parse timeout", err)
	}
	err = testTask.WaitStopped(timeout)
	if err != nil {
		t.Fatalf("Timed out waiting for task to reach stopped. Error %#v, task %#v", err, testTask)
	}

	if exit, ok := testTask.ContainerExitcode("exit"); !ok || exit != 42 {
		t.Errorf("Expected exit to exit with 42; actually exited (%v) with %v", ok, exit)
	}

}

// TestHostname Check that hostname works
func TestHostname(t *testing.T) {
	// Parallel is opt in because resource constraints could cause test failures
	// on smaller instances
	if os.Getenv("ECS_FUNCTIONAL_PARALLEL") != "" {
		t.Parallel()
	}
	agent := RunAgent(t, nil)
	defer agent.Cleanup()
	agent.RequireVersion(">=1.5.0")

	testTask, err := agent.StartTask(t, "hostname-windows")
	if err != nil {
		t.Fatal("Could not start task", err)
	}
	timeout, err := time.ParseDuration("2m")
	if err != nil {
		t.Fatal("Could not parse timeout", err)
	}
	err = testTask.WaitStopped(timeout)
	if err != nil {
		t.Fatalf("Timed out waiting for task to reach stopped. Error %#v, task %#v", err, testTask)
	}

	if exit, ok := testTask.ContainerExitcode("exit"); !ok || exit != 42 {
		t.Errorf("Expected exit to exit with 42; actually exited (%v) with %v", ok, exit)
	}

}

// TestSimpleExit Tests that the basic premis of this testing fromwork works (e.g. exit codes go through, etc)
func TestSimpleExit(t *testing.T) {
	// Parallel is opt in because resource constraints could cause test failures
	// on smaller instances
	if os.Getenv("ECS_FUNCTIONAL_PARALLEL") != "" {
		t.Parallel()
	}
	agent := RunAgent(t, nil)
	defer agent.Cleanup()
	agent.RequireVersion(">=1.0.0")

	testTask, err := agent.StartTask(t, "simple-exit-windows")
	if err != nil {
		t.Fatal("Could not start task", err)
	}
	timeout, err := time.ParseDuration("2m")
	if err != nil {
		t.Fatal("Could not parse timeout", err)
	}
	err = testTask.WaitStopped(timeout)
	if err != nil {
		t.Fatalf("Timed out waiting for task to reach stopped. Error %#v, task %#v", err, testTask)
	}

	if exit, ok := testTask.ContainerExitcode("exit"); !ok || exit != 42 {
		t.Errorf("Expected exit to exit with 42; actually exited (%v) with %v", ok, exit)
	}

}

// TestWorkingDir Check that working dir works
func TestWorkingDir(t *testing.T) {
	// Parallel is opt in because resource constraints could cause test failures
	// on smaller instances
	if os.Getenv("ECS_FUNCTIONAL_PARALLEL") != "" {
		t.Parallel()
	}
	agent := RunAgent(t, nil)
	defer agent.Cleanup()
	agent.RequireVersion(">=1.5.0")

	testTask, err := agent.StartTask(t, "working-dir-windows")
	if err != nil {
		t.Fatal("Could not start task", err)
	}
	timeout, err := time.ParseDuration("2m")
	if err != nil {
		t.Fatal("Could not parse timeout", err)
	}
	err = testTask.WaitStopped(timeout)
	if err != nil {
		t.Fatalf("Timed out waiting for task to reach stopped. Error %#v, task %#v", err, testTask)
	}

	if exit, ok := testTask.ContainerExitcode("exit"); !ok || exit != 42 {
		t.Errorf("Expected exit to exit with 42; actually exited (%v) with %v", ok, exit)
	}

}
