package peano

import (
	"testing"
)

func TestGWLog(t *testing.T) {
	InitLogger("debug", "./test.log")
	Errorf("2+5=%d", 7)
	Debugf("debug log")
	Infof("info log = %s", "infolevel")
	Warnf("this is a warning %d", "franco")
}
