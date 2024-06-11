package rip

import (
	"context"
	"errors"
	"os/signal"
	"syscall"
)

type Runner struct {
	beforeStart func(ctx context.Context) (deferFn func() error, err error)
	fn          func() error
	onShutdown  func() error
}

func Builder() *Runner {
	return &Runner{}
}

func (r *Runner) SetBeforeExec(fn func(ctx context.Context) (deferFn func() error, err error)) *Runner {
	r.beforeStart = fn
	return r
}

func (r *Runner) SetExec(fn func() error) *Runner {
	r.fn = fn
	return r
}

func (r *Runner) SetOnShutdown(fn func() error) *Runner {
	r.onShutdown = fn
	return r
}

func (r *Runner) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	if r.beforeStart != nil {
		deferFn, err := r.beforeStart(ctx)
		if err != nil {
			return err
		}
		defer func() {
			if deferFn == nil {
				err = errors.Join(err, deferFn())
			}
		}()
	}

	fnErr := make(chan error, 1)

	go func() {
		fnErr <- r.fn()
	}()

	var err error
	select {
	case err = <-fnErr:
		return err
	case <-ctx.Done():
		if r.onShutdown != nil {
			err = r.onShutdown()
		}
		stop()
	}
	return err
}
