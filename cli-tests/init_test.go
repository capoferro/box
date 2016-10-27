package main

import (
	"bytes"
	"os"
	. "testing"

	"github.com/erikh/testcli"

	. "gopkg.in/check.v1"
)

type cliSuite struct{}

var _ = Suite(&cliSuite{})

func TestCLI(t *T) {
	TestingT(t)
}

func (s *cliSuite) SetUpTest(c *C) {
	os.Setenv("NO_CACHE", "1")
}

func build(content string, extraArgs ...string) *testcli.Cmd {
	c := testcli.Command("box", extraArgs...)
	if content != "" {
		buf := bytes.NewBufferString(content)
		c.SetStdin(buf)
	}
	c.Run()

	return c
}

func checkSuccess(c *C, cmd *testcli.Cmd) {
	c.Assert(cmd.Success(), Equals, true, Commentf("stdout:\n%s\nstderr:\n%s\n", cmd.Stdout(), cmd.Stderr()))
}

func checkFailure(c *C, cmd *testcli.Cmd) {
	c.Assert(cmd.Failure(), Equals, true, Commentf("stdout:\n%s\nstderr:\n%s\n", cmd.Stdout(), cmd.Stderr()))
}