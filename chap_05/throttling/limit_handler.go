package throttling

import "net/http"

// LimitHandler is a middleware that limits the connection number
type LimitHandler struct {
	connsChan chan struct{}
	next      http.Handler
}

// NewLimitHandler returns a new LimitHandler
func NewLimitHandler(maxConnNum int, next http.Handler) *LimitHandler {
	cc := make(chan struct{}, maxConnNum)

	for i := 0; i < maxConnNum; i++ {
		cc <- struct{}{}
	}

	return &LimitHandler{
		connsChan: cc,
		next:      next,
	}
}

func (h *LimitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	select {
	case <-h.connsChan:
		h.next.ServeHTTP(w, r)
		h.connsChan <- struct{}{} // release the lock
	default: // return busy if we cannot get item from the channel
		http.Error(w, "Busy", http.StatusTooManyRequests)
	}
}
