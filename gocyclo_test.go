package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"go/token"
	"strings"
	"testing"
)

func Test_writeStats_ignores_vendor_packages(t *testing.T){
	//Arrange
	*novendor = true
	var sortedStats = []stat{
		{
			PkgName:    "Ignored",
			FuncName:   "TestFunction",
			Complexity: 1,
			Pos:        token.Position{
				Filename: "vendor\\blaat.com\\functions",
				Offset:   0,
				Line:     0,
				Column:   0,
			},
		},
		{
			PkgName:    "NotIgnored",
			FuncName:   "TestFunction",
			Complexity: 1,
			Pos:        token.Position{
				Filename: "blaat.com\\functions",
				Offset:   0,
				Line:     0,
				Column:   0,
			},
		},
	} //Act
	w := &bytes.Buffer{}
	result := writeStats(w, sortedStats)

	//Assert
	assert.Equal(t,1, result)
	assert.True(t, strings.Contains(w.String(), "NotIgnored"))
	assert.False(t, strings.Contains(w.String(), "vendor"))
}

func Test_writeStats_shows_vendor_packages(t *testing.T){
	//Arrange
	*novendor = false
	sortedStats:= []stat{
		{
			PkgName:    "Ignored",
			FuncName:   "vendor/a/b/c.go",
			Complexity: 1,
			Pos:        token.Position{},
		},
		{
			PkgName:    "NotIgnored",
			FuncName:   "g/c.go",
			Complexity: 1,
			Pos:        token.Position{},
		},
	}

	//Act
	w := &bytes.Buffer{}
	result := writeStats(w, sortedStats)

	//Assert
	assert.Equal(t,2, result)
	assert.True(t, strings.Contains(w.String(), "NotIgnored"))
	assert.True(t, strings.Contains(w.String(), "vendor"))
}
