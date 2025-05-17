package cron_match

import (
	"github.com/sixleaveakkm/go-utils/ctime"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCronFilter_StartDurationMatches(t *testing.T) {
	_ = os.Setenv("TZ", "UTC")
	t.Run("with nth day of the month", func(t *testing.T) {
		cronFilter, err := NewCronFilter(map[string]any{
			"start":    "0 12 * * SAT#2",
			"duration": "10h",
		})
		require.NoError(t, err)

		tt := []struct {
			tm       time.Time
			expected bool
		}{
			{
				time.Date(2024, 2, 10, 20, 40, 0, 0, ctime.JST),
				false,
			}, {
				time.Date(2024, 2, 10, 21, 0, 0, 0, ctime.JST),
				true,
			}, {
				time.Date(2024, 2, 11, 2, 0, 0, 0, ctime.JST),
				true,
			}, {
				time.Date(2024, 2, 11, 7, 30, 0, 0, ctime.JST),
				false,
			}, {
				time.Date(2024, 2, 17, 21, 10, 0, 0, ctime.JST),
				false,
			}, {
				time.Date(2024, 2, 3, 21, 10, 0, 0, ctime.JST),
				false,
			},
		}

		for _, s := range tt {
			if got := cronFilter.Matches(s.tm); got != s.expected {
				t.Errorf("%s expected %v but got %v", s.tm.String(), s.expected, got)
			}
		}
	})

	t.Run("with no nth day of month", func(t *testing.T) {
		cronFilter, err := NewCronFilter(map[string]any{
			"start":    "0 12 * * SAT",
			"duration": "10h",
		})
		require.NoError(t, err)

		tt := []struct {
			tm       time.Time
			expected bool
		}{
			{
				time.Date(2024, 2, 10, 20, 40, 0, 0, ctime.JST),
				false,
			}, {
				time.Date(2024, 2, 10, 21, 0, 0, 0, ctime.JST),
				true,
			}, {
				time.Date(2024, 2, 11, 2, 0, 0, 0, ctime.JST),
				true,
			}, {
				time.Date(2024, 2, 11, 7, 30, 0, 0, ctime.JST),
				false,
			}, {
				time.Date(2024, 2, 17, 21, 10, 0, 0, ctime.JST),
				true,
			}, {
				time.Date(2024, 2, 3, 21, 10, 0, 0, ctime.JST),
				true,
			},
		}

		for _, s := range tt {
			if got := cronFilter.Matches(s.tm); got != s.expected {
				t.Errorf("%s expected %v but got %v", s.tm.String(), s.expected, got)
			}
		}
	})
}

func TestCronFilter_StartEndMatches(t *testing.T) {
	_ = os.Setenv("TZ", "UTC")
	t.Run("start end", func(t *testing.T) {
		cronFilter, err := NewCronFilter(map[string]any{
			"start": "0 12 * * SAT",
			"end":   "0 22 * * SAT",
		})
		require.NoError(t, err)

		tt := []struct {
			tm       time.Time
			expected bool
		}{
			{
				time.Date(2024, 2, 10, 20, 40, 0, 0, ctime.JST),
				false,
			}, {
				time.Date(2024, 2, 10, 21, 0, 0, 0, ctime.JST),
				true,
			}, {
				time.Date(2024, 2, 11, 2, 0, 0, 0, ctime.JST),
				true,
			}, {
				time.Date(2024, 2, 11, 7, 30, 0, 0, ctime.JST),
				false,
			}, {
				time.Date(2024, 2, 17, 21, 10, 0, 0, ctime.JST),
				true,
			}, {
				time.Date(2024, 2, 3, 21, 10, 0, 0, ctime.JST),
				true,
			},
		}

		for _, s := range tt {
			if got := cronFilter.Matches(s.tm); got != s.expected {
				t.Errorf("%s expected %v but got %v", s.tm.String(), s.expected, got)
			}
		}
	})
}
