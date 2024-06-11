package asyncexec

import (
	"context"
	"fmt"
	. "github.com/sixleaveakkm/go-utils"
	"github.com/sixleaveakkm/go-utils/errz"
	"github.com/sixleaveakkm/go-utils/ptr"
	"math"

	"golang.org/x/time/rate"
)

type AsyncErr[T any] struct {
	Input T
	Err   error
}

func (ae AsyncErr[T]) Error() string {
	return fmt.Sprintf("error on input: %v, Err: %v", ae.Input, ae.Err)
}

func asyncWarpErr[T any](input T, err error) error {
	if err == nil {
		return nil
	}
	return AsyncErr[T]{
		Input: input,
		Err:   err,
	}
}

const UnlimitedWorker = math.MaxInt

type AsyncDoFn[T any, R any] func(e T) (R, error)

type AsyncPool[T any, R any] struct {
	elements *[]T
	fn       AsyncDoFn[T, R]
	ctx      context.Context
	limiter  *rate.Limiter
	worker   int
}

// Do create an async worker pool to execute function for a list of elements with rate limit.
// It will execute when Await() is called.
// The number of worker should be set according to your rate requirement.
// Usually for function like http request with response less than 1 second, set worker number equals to rps is enough.
// In some case you want to reach max rps, set worker to UnlimitedWorker will not wait for function response,
// which will create a goroutine for each running task, may have performance issue when rps is too high.
func Do[T any, R any](fn AsyncDoFn[T, R]) *AsyncPool[T, R] {
	return &AsyncPool[T, R]{
		fn:     fn,
		worker: UnlimitedWorker,
		ctx:    context.TODO(),
	}
}

func (ap *AsyncPool[T, R]) WithContext(ctx context.Context) *AsyncPool[T, R] {
	ap.ctx = ctx
	return ap
}

func (ap *AsyncPool[T, R]) For(elements []T) *AsyncPool[T, R] {
	ap.elements = &elements
	return ap
}

func (ap *AsyncPool[T, R]) MaxRps(c float64) *AsyncPool[T, R] {
	ap.limiter = rate.NewLimiter(rate.Limit(c), 1)
	return ap
}

func (ap *AsyncPool[T, R]) SetLimiter(l *rate.Limiter) *AsyncPool[T, R] {
	ap.limiter = l
	return ap
}

// SetWorker set the number of worker. Worker number is suggest to be greater than rps.
// For time cost task and want maximum rate, set worker to UnlimitedWorker which will create a goroutine
// for each task. Notice this will create more goroutine than worker mode.
func (ap *AsyncPool[T, R]) SetWorker(n int) *AsyncPool[T, R] {
	ap.worker = n
	return ap
}

func (ap *AsyncPool[T, R]) AwaitN1() ([]R, error) {
	if err := ap.checkParameter(); err != nil {
		return nil, err
	}
	errs := errz.New()
	rs := ap.Await()
	list := make([]R, 0, len(rs))
	for _, r := range rs {
		if r.UnwrapError() != nil {
			errs.Append(r.UnwrapError())
		} else {
			list = append(list, r.Unwrap())
		}
	}
	return list, errs.Err()
}

func (ap *AsyncPool[T, R]) Await() []*Result[R] {
	if err := ap.checkParameter(); err != nil {
		return []*Result[R]{
			ptr.Ptr(Err[R](err)),
		}
	}
	reqChan := make(chan *T, len(*ap.elements))
	resultChan := make(chan *Result[R], len(*ap.elements))

	if ap.worker != UnlimitedWorker {
		for i := 0; i < ap.worker; i++ {
			go func() {
				for e := range reqChan {
					_ = ap.limiter.Wait(ap.ctx)
					res, err := ap.fn(*e)
					resultChan <- ptr.Ptr(ResultFrom(res, asyncWarpErr(e, err)))
				}
			}()
		}
	} else {
		go func() {
			for e := range reqChan {
				_ = ap.limiter.Wait(ap.ctx)
				go func(e *T) {
					res, err := ap.fn(*e)
					resultChan <- ptr.Ptr(ResultFrom(res, asyncWarpErr(e, err)))
				}(e)
			}
		}()
	}

	for i := range *ap.elements {
		reqChan <- &(*ap.elements)[i]
	}
	close(reqChan)

	var results []*Result[R]
	for j := 0; j < len(*ap.elements); j++ {
		result := <-resultChan
		results = append(results, result)
	}
	return results
}

func (ap *AsyncPool[T, R]) checkParameter() error {
	if ap.elements == nil {
		return fmt.Errorf("elements is nil")
	}
	if ap.fn == nil {
		return fmt.Errorf("fn is nil")
	}
	if ap.limiter == nil {
		return fmt.Errorf("rps is nil")
	}
	if ap.worker <= 0 {
		ap.worker = UnlimitedWorker
	}
	return nil
}
