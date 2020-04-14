// Copyright (c) 2020 Uber Technologies, Inc.
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

package transport_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/yarpcerrors"
)

type testHandler struct {
	fn func(ctx context.Context) error
}

func (h testHandler) Handle(ctx context.Context, req *transport.Request, _ transport.ResponseWriter) error {
	return h.fn(ctx)
}

func TestUnaryHandlerInvokerWrappedCtxError(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()
	<-ctx.Done()

	handler := testHandler{fn: func(ctx context.Context) error {
		require.Equal(t, context.DeadlineExceeded, ctx.Err(), "unexpected context error")
		return fmt.Errorf("wrapped %w", ctx.Err())
	}}

	invokeReq := transport.UnaryInvokeRequest{
		Context:   ctx,
		StartTime: time.Now(),
		Request:   &transport.Request{},
		Handler:   handler,
	}

	err := transport.InvokeUnaryHandler(invokeReq)
	require.Error(t, err, "expected wrapped error")

	st := yarpcerrors.FromError(err)
	assert.Equal(t, yarpcerrors.CodeDeadlineExceeded.String(), st.Code().String(), "unexpeceted error")
}
