package tests

import (
	"io"
	"os"
	"testing"

	"github.com/secDre4mer/gyp"
	"github.com/secDre4mer/gyp/pb"
	"github.com/stretchr/testify/assert"
)

// These are just utilities

func openTestFile(t *testing.T, fname string) io.Reader {
	f, err := os.Open(fname)
	assert.NoError(t, err, `Couldn't open file "%s"`, fname)
	return f
}

func parseTestFile(t *testing.T, fname string) (*pb.RuleSet, error) {
	f := openTestFile(t, fname)
	rs, err := gyp.Parse(f)
	return rs.AsProto(), err
}
