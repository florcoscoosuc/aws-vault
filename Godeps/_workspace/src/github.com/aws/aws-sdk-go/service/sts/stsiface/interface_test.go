// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package stsiface_test

import (
	"testing"

	"github.com/99designs/aws-vault/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/sts"
	"github.com/99designs/aws-vault/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/sts/stsiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*stsiface.STSAPI)(nil), sts.New(nil))
}