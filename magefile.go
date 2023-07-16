//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

func Build() error {
	fmt.Println("build test_app...")

	appBuildPath, ok := os.LookupEnv("TEST_APP_BUILD_PATH")
	if !ok {
		appBuildPath = "test_app"
	}

	err := sh.RunWith(
		nil, "go", "build", "-o", appBuildPath,
	)
	if err != nil {
		return fmt.Errorf("failed to build test_app: %v", err)
	}
	return nil
}
