package asyncexec_test

import (
	"github.com/sixleaveakkm/go-utils/asyncexec"
	"github.com/sixleaveakkm/go-utils/toy"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	var ids []int
	for i := 0; i < 2000; i++ {
		ids = append(ids, i)
	}

	tr := toy.NewTimeRecorder()

	type user struct {
		ID    int
		Name  string
		Email string
	}

	getUserByID := func(id int) (user, error) {
		// mock
		time.Sleep(100 * time.Millisecond)
		tr.Record()
		return user{
			ID:   id,
			Name: "foo",
		}, nil
	}

	//_ = asyncexec.Do(getUserByID).For(ids).MaxRps(500).SetWorker(500).Await()
	_ = asyncexec.Do(getUserByID).For(ids).MaxRps(500).Await()

	tr.Graph(100 * time.Millisecond)

}
