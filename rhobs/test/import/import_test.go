package test

import (
	"fmt"
	op "github.com/rhobs/obo-prometheus-operator/pkg/operator"
	"testing"
)

func TestDefalThanosImage(t *testing.T) {
	fmt.Printf("DefaultThanosImage is %s", op.DefaultThanosImage)
}
