package test

import (
	"testing"

	"github.com/zhangjiacheng-iHealth/IHCommunity/package/web/auth"
)

func TestAuth(t *testing.T) {
	auth.DeletePolicy("1", "", "*")
}
