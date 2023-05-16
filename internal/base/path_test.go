package base

import (
	"bytes"
	"fmt"
	"testing"
)

var (
	src = []byte(`import (
		"context"
		"phanes/store"

		"github.com/lizhiqpxv/proto/base"
		"github.com/lizhiqpxv/proto/dto"
		log "go-micro.dev/v4/logger"
	)
	
	var User = &user{}
	
	type user struct {
		user store.IUser
	}
	`)
)

func Test_readfileForLine(t *testing.T) {
	bufs := bytes.NewReader(src)
	buf := readfileForLine(bufs, []string{"phanes", "hello"})
	fmt.Println(string(buf))

}
