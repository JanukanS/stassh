package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpress(t *testing.T) {
	addr1 := Address{username: "someone",
		ip:   "144.144.144.144",
		port: 22,
		desc: "test case"}

	expr1 := addr1.express()
	assert.Equal(t, expr1, "someone@144.144.144.144", "should equal")

	addr1.port = 23
	expr2 := addr1.express()
	assert.Equal(t, expr2, "someone@144.144.144.144:23", "Should be Equal")

	addr1.username = ""
	expr3 := addr1.express()
	assert.Equal(t, expr3, "144.144.144.144:23", "Should be equal")
}
