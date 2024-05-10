package tess

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func Test_isCI(t *testing.T) {
	os.Clearenv()

	t.Run("no CI env should be false", func(t *testing.T) {
		require.False(t, isCI())
	})

	t.Run("CI env set to 'true'", func(t *testing.T) {
		_ = os.Setenv("CI", "true")
		defer os.Clearenv()
		require.True(t, isCI())
	})

	t.Run("CI env set to '1'", func(t *testing.T) {
		_ = os.Setenv("CI", "1")
		defer os.Clearenv()
		require.True(t, isCI())
	})

	t.Run("CI env set to 'false'", func(t *testing.T) {
		_ = os.Setenv("CI", "false")
		defer os.Clearenv()
		require.False(t, isCI())
	})

	t.Run("CI env set to meaningless string", func(t *testing.T) {
		_ = os.Setenv("CI", "foo")
		defer os.Clearenv()
		require.False(t, isCI())
	})
}

func TestIsSameTime(t *testing.T) {
	JST := time.FixedZone("Asia/Tokyo", 9*3600)

	t1 := time.Date(2024, 1, 2, 1, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 2, 1, 0, 0, 100, time.UTC)
	t3 := time.Date(2024, 1, 2, 10, 0, 0, 100, JST)
	t4 := time.Date(2024, 1, 2, 1, 0, 1, 0, time.UTC)

	t.Run("100 nsec drift consider as same", func(t *testing.T) {
		require.True(t, IsSameTime(t1, t2))
	})

	t.Run("100 nsec drift in different timezone consider as same", func(t *testing.T) {
		require.True(t, IsSameTime(t1, t3))
	})

	t.Run("1 sec drift consider different", func(t *testing.T) {
		require.False(t, IsSameTime(t1, t4))
	})

	t.Run("1 sec drift consider same if set drift to 2 sec", func(t *testing.T) {
		require.True(t, IsSameTime(t1, t4, time.Second*2))
	})

	t.Run("1 sec drift consider different as threshold is 1 sec", func(t *testing.T) {
		require.False(t, IsSameTime(t1, t4, time.Second))
	})

	t.Run("test Diff", func(t *testing.T) {
		require.Equal(t, "", Diff(t1, t2))
	})
}
