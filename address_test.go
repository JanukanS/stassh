package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testExpress(t *testing.T) {
	addr1 := Address{username: "someone",
		ip:   "144.144.144.144",
		port: 22,
		desc: "test case"}

	expr1 := addr1.express()
	assert.Equal(t, expr1, "someone@144.144.144.144", "should equal")
}
