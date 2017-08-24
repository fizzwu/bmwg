package throttling

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newTestHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		<-r.Context().Done()
	})
}

func setup(ctx context.Context) (*httptest.ResponseRecorder, *http.Request) {
	r := httptest.NewRequest("GET", "/health", nil)
	r = r.WithContext(ctx)
	return httptest.NewRecorder(), r
}

func TestCallNextWhenConnOK(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	hand := NewLimitHandler(2, newTestHandler(ctx))
	w, r := setup(ctx)

	go hand.ServeHTTP(w, r)
	cancel()

	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReturnBusyWhenLimitZero(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	hand := NewLimitHandler(0, newTestHandler(ctx))
	w, r := setup(ctx)

	time.AfterFunc(10*time.Millisecond, func() {
		cancel()
	})
	hand.ServeHTTP(w, r)
	assert.Equal(t, http.StatusTooManyRequests, w.Code)
}

func TestReturnOKWhenConnection2AndLimit2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	hand := NewLimitHandler(2, newTestHandler(ctx))
	w, r := setup(ctx)
	w2, r2 := setup(ctx2)

	time.AfterFunc(10*time.Millisecond, func() {
		cancel()
		cancel2()
	})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		hand.ServeHTTP(w, r)
		wg.Done()
	}()

	go func() {
		hand.ServeHTTP(w2, r2)
		wg.Done()
	}()

	wg.Wait()

	if w.Code != http.StatusOK && w2.Code != http.StatusOK {
		t.Fatalf("Both requests should be OK, request 1: %v, request 2: %v", w.Code, w2.Code)
	}
}

func TestReturnBusyWhenConnTouchsTheLimit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	hand := NewLimitHandler(1, newTestHandler(ctx))
	w, r := setup(ctx)
	w2, r2 := setup(ctx2)

	time.AfterFunc(10*time.Millisecond, func() {
		cancel()
		cancel2()
	})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		hand.ServeHTTP(w, r)
		wg.Done()
	}()

	go func() {
		hand.ServeHTTP(w2, r2)
		wg.Done()
	}()

	wg.Wait()

	if w.Code == http.StatusOK && w2.Code == http.StatusOK {
		t.Fatalf("one request should be busy, request 1: %v, request 2: %v", w.Code, w2.Code)
	}
}
