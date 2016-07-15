// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package echo

import (
	"time"

	"github.com/yarpc/yarpc-go"
	disp "github.com/yarpc/yarpc-go/crossdock/client/dispatcher"
	"github.com/yarpc/yarpc-go/crossdock/client/random"
	"github.com/yarpc/yarpc-go/crossdock/thrift/echo"
	"github.com/yarpc/yarpc-go/crossdock/thrift/echo/yarpc/echoclient"

	"github.com/crossdock/crossdock-go"
	"golang.org/x/net/context"
)

// Thrift implements the 'thrift' behavior.
func Thrift(t crossdock.T) {
	t = createEchoT("thrift", t)
	fatals := crossdock.Fatals(t)

	dispatcher := disp.Create(t)
	fatals.NoError(dispatcher.Start(), "could not start Dispatcher")
	defer dispatcher.Stop()

	client := echoclient.New(dispatcher.Channel("yarpc-test"))
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	token := random.String(5)
	pong, _, err := client.Echo(
		yarpc.NewReqMeta(ctx),
		&echo.Ping{Beep: token},
	)

	crossdock.Fatals(t).NoError(err, "call to Echo::echo failed: %v", err)
	crossdock.Assert(t).Equal(token, pong.Boop, "server said: %v", pong.Boop)
}