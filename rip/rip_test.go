package rip_test

import (
	"context"
	"github.com/sixleaveakkm/go-utils/rip"
	"log"
	"net/http"
)

func ExampleRunner_Run() {
	s := http.Server{
		Addr: ":8080",
	}

	setupOTel := func(ctx context.Context) (deferFn func() error, err error) {
		// Set up propagator.
		return nil, nil
	}

	exec := func() error {
		return s.ListenAndServe()
	}

	onShutdown := func() error {
		return s.Shutdown(context.Background())
	}

	err := rip.Builder().SetBeforeExec(setupOTel).SetExec(exec).SetOnShutdown(onShutdown).Run()
	if err != nil {
		log.Fatalln(err)
	}
}
