package sth_about_adrr

import (
	"fmt"
	"testing"
)

func TestTryAdrr(t *testing.T) {
	var aWorker Worker
	TryAdrr(&aWorker)
	fmt.Println(aWorker.Sb)
}
