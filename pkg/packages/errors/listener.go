package errors

import (
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
)

type Listener struct {
	started        atomic.Bool
	mu             sync.Mutex
	ctx            context.Context
	cancel         func()
	errChan        chan error
	listenerCtx    context.Context
	listenerCancel func()
}

func (l *Listener) GetErrChan() chan error {
	return l.errChan
}

func NewListener(ctx context.Context, cancel func()) *Listener {
	errChan := make(chan error)
	return &Listener{
		ctx:     ctx,
		cancel:  cancel,
		errChan: errChan,
	}
}

func (l *Listener) Start() *Error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.started.Load() {
		return NewInternalError("listener already started")
	}

	l.listenerCtx, l.listenerCancel = context.WithCancel(l.ctx)

	go l.errorsHandler()
	l.started.Store(true)
	return nil
}
func (l *Listener) Stop() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.started.Load() {
		return NewInternalError("listener stopped")
	}
	l.listenerCancel()
	close(l.errChan)
	l.started.Store(false)
	return nil
}

func (l *Listener) errorsHandler() {
	const SLEEP_DEFAULT_TIME = 300 * time.Millisecond

	for {
		select {
		case err := <-l.errChan:
			internalErr, ok := err.(*Error)
			if !ok {
				b, err := json.Marshal(err)
				if err == nil {
					logrus.Warning("[!] App golang error: ", string(b))
				} else {
					logrus.Warning("[!] App golang error: ", err)
				}
				continue
			}

			if internalErr.IsCriticalError() {
				b, err := json.Marshal(err)
				if err == nil {
					logrus.Error("[!] App critical error: ", string(b))
				} else {
					logrus.Error("[!] App critical error: ", err)
				}
				logrus.Debug(string(internalErr.Stack))
				l.cancel()
				return
			}

			b, err := json.Marshal(err)
			if err == nil {
				logrus.Warning("[!] App internal error: ", string(b))
			} else {
				logrus.Warning("[!] App internal error: ", err)
			}
			logrus.Debug(string(internalErr.Stack))
		case <-l.ctx.Done():
			return
		case <-l.listenerCtx.Done():
			return
		default:
			time.Sleep(SLEEP_DEFAULT_TIME)
		}
	}
}
