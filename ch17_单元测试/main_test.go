package main

import (
	"testing"
)

func TestInit(t *testing.T) {
	t.Log("heh")

	helper := PersonHelper{}
	helper.init("pleuvoir")
	t.Log(helper.Name)
}
