package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	testCases := []struct {
		version_requirement string
		version             string
		expectResult        bool
	}{
		{
			version_requirement: ">=1.3.0",
			version:             "refs/tags/1.4.0",
			expectResult:        true,
		},
		{
			version_requirement: ">=1.3.0",
			version:             "refs/tags/1.3.0",
			expectResult:        true,
		},
		{
			version_requirement: ">=1.3.0",
			version:             "refs/tags/1.2.0",
			expectResult:        false,
		},
		{
			version_requirement: ">1.3.0",
			version:             "refs/tags/1.4.0",
			expectResult:        true,
		},
		{
			version_requirement: ">1.3.0",
			version:             "refs/tags/1.3.0",
			expectResult:        false,
		},
		{
			version_requirement: ">1.3.0",
			version:             "refs/tags/1.2.0",
			expectResult:        false,
		},
		{
			version_requirement: "1.3.0",
			version:             "refs/tags/1.2.0",
			expectResult:        false,
		},
		{
			version_requirement: "1.3.0",
			version:             "refs/tags/1.3.0",
			expectResult:        true,
		},
		{
			version_requirement: "",
			version:             "refs/tags/1.3.0",
			expectResult:        true,
		},
	}
	for _, tCase := range testCases {
		assert.Equal(t, tCase.expectResult, versionRequired(tCase.version_requirement, tCase.version), "version check result not match")
	}
}
