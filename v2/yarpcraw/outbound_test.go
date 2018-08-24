// Copyright (c) 2018 Uber Technologies, Inc.
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

package yarpcraw

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/uber/tchannel-go/testutils/testreader"
	"go.uber.org/yarpc/v2"
	"go.uber.org/yarpc/v2/internal/clientconfig"
	"go.uber.org/yarpc/v2/yarpctransporttest"
)

func TestCall(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	caller := "caller"
	service := "service"

	tests := []struct {
		procedure    string
		headers      map[string]string
		body         []byte
		responseBody [][]byte
		responseErr  error

		want        []byte
		wantErr     string
		wantHeaders map[string]string
	}{
		{
			procedure:    "foo",
			body:         []byte{1, 2, 3},
			responseBody: [][]byte{{4}, {5}, {6}},
			want:         []byte{4, 5, 6},
		},
		{
			procedure:    "foo",
			body:         []byte{1, 2, 3},
			responseBody: [][]byte{{4}, {5}, {6}},
			responseErr:  errors.New("bar"),
			want:         []byte{4, 5, 6},
			wantErr:      "bar",
		},
		{
			procedure:    "bar",
			body:         []byte{1, 2, 3},
			responseBody: [][]byte{{4}, {5}, nil, {6}},
			wantErr:      "error set by user",
		},
		{
			procedure:    "headers",
			headers:      map[string]string{"x": "y"},
			body:         []byte{},
			responseBody: [][]byte{},
			want:         []byte{},
			wantHeaders:  map[string]string{"a": "b"},
		},
	}

	for _, tt := range tests {
		outbound := yarpctransporttest.NewMockUnaryOutbound(mockCtrl)
		client := New(clientconfig.MultiOutbound(caller, service,
			yarpc.Outbounds{
				Unary: outbound,
			}))

		writer, responseBody := testreader.ChunkReader()
		for _, chunk := range tt.responseBody {
			writer <- chunk
		}
		close(writer)

		outbound.EXPECT().Call(gomock.Any(),
			yarpctransporttest.NewRequestMatcher(t,
				&yarpc.Request{
					Caller:    caller,
					Service:   service,
					Procedure: tt.procedure,
					Headers:   yarpc.HeadersFromMap(tt.headers),
					Encoding:  Encoding,
					Body:      bytes.NewReader(tt.body),
				}),
		).Return(
			&yarpc.Response{
				Body:    ioutil.NopCloser(responseBody),
				Headers: yarpc.HeadersFromMap(tt.wantHeaders),
			}, tt.responseErr)

		var (
			opts       []yarpc.CallOption
			resHeaders map[string]string
		)

		for k, v := range tt.headers {
			opts = append(opts, yarpc.WithHeader(k, v))
		}
		opts = append(opts, yarpc.ResponseHeaders(&resHeaders))

		resBody, err := client.Call(ctx, tt.procedure, tt.body, opts...)
		if tt.wantErr != "" {
			if assert.Error(t, err) {
				assert.Equal(t, err.Error(), tt.wantErr)
			}
		} else {
			assert.NoError(t, err)
		}
		if tt.want != nil {
			assert.Equal(t, tt.want, resBody)
		}
		if tt.wantHeaders != nil {
			assert.Equal(t, tt.wantHeaders, resHeaders)
		}
	}
}
