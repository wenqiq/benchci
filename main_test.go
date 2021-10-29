package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
)

func TestVersionCompare(t *testing.T) {
	testCases := []struct {
		versionRequirement string
		version            string
		expectResult       bool
	}{
		{
			versionRequirement: ">=1.9.0",
			version:            "1.10.0",
			expectResult:       true,
		},
		{
			versionRequirement: ">=1.3.0",
			version:            "1.3.0",
			expectResult:       true,
		},
		{
			versionRequirement: ">=1.3.0",
			version:            "1.2.0",
			expectResult:       false,
		},
		{
			versionRequirement: ">v1.3.0",
			version:            "v1.4.0",
			expectResult:       true,
		},
		{
			versionRequirement: ">1.3.0",
			version:            "v1.3.0",
			expectResult:       false,
		},
		{
			versionRequirement: ">v1.3.0",
			version:            "1.2.0",
			expectResult:       false,
		},
		{
			versionRequirement: "1.3.0",
			version:            "1.2.0",
			expectResult:       false,
		},
		{
			versionRequirement: "1.3.0",
			version:            "1.3.0",
			expectResult:       true,
		},
		{
			versionRequirement: "",
			version:            "1.3.0",
			expectResult:       true,
		},
	}
	for _, tCase := range testCases {
		assert.Equal(t, tCase.expectResult, versionRequired(tCase.versionRequirement, tCase.version), "version check result not match")
	}
}

func TestGetReleaseVersion(t *testing.T) {
	r, err := git.PlainOpen("./../antrea")
	assert.NoError(t, err, "")
	prevVersionTag, err := getLatestRelease(r)
	assert.NoError(t, err, "")
	fmt.Println(prevVersionTag.Hash(), prevVersionTag.String(), prevVersionTag.Name().Short())
}
