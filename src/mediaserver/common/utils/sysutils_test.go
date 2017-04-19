package utils

import (
	"testing"
)

func TestListFileInfo(t *testing.T) {

	infos := listFileInfo("/home/mao/test", nil)

	for _, i := range infos {
		t.Logf("name: %s, time: %d", i.Name(), i.ModTime().Unix())
	}
}
