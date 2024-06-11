package configer_test

import (
	"fmt"
	"github.com/sixleaveakkm/go-utils/configer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"os"
	"testing"
	"time"
)

type InsideConfig struct {
	String  string `env:"S1"`
	String2 string `env:"S2"`
	String3 string `env:"S3"`
}

type Config struct {
	String  string        `env:"S1"`
	String2 string        `env:"S2"`
	String3 string        `env:"S3"`
	Bool1   bool          `env:"B1" default:"true"`
	Bool2   bool          `env:"B2" default:"false"`
	Int     int           `env:"I1" default:"10"`
	Dur     time.Duration `env:"D1" default:"1s"`
	URL     url.URL       `env:"D2" default:"http://example.com"`
	Inside  InsideConfig  `envPrefix:"INSIDE_"`
}

func TestParse(t *testing.T) {
	c := Config{
		String:  "string",
		String2: "string2",
		Inside: InsideConfig{
			String:  "inside string",
			String2: "inside string2",
		},
	}

	_ = os.Setenv("S2", "String2")
	_ = os.Setenv("S3", "String3")
	_ = os.Setenv("INSIDE_S1", "Inside 1")
	_ = os.Setenv("INSIDE_S3", "Inside 3")

	err := configer.Parse(&c)
	require.NoError(t, err)

	assert.Equal(t, "string", c.String)
	assert.Equal(t, "String2", c.String2)
	assert.Equal(t, "String3", c.String3)
	assert.Equal(t, false, c.Bool1)
	assert.Equal(t, "Inside 1", c.Inside.String)
	assert.Equal(t, "inside string2", c.Inside.String2)
	assert.Equal(t, "Inside 3", c.Inside.String3)
}

func TestBuilder(t *testing.T) {
	c := Config{
		String:  "string",
		String2: "string2",
		Inside: InsideConfig{
			String:  "inside string",
			String2: "inside string2",
		},
	}

	_ = os.Setenv("S2", "String2")
	_ = os.Setenv("S3", "String3")
	_ = os.Setenv("INSIDE_S1", "Inside 1")
	_ = os.Setenv("INSIDE_S3", "Inside 3")

	configer.For(c).LoadYamlFile("config.yaml")
}

func TestRoutine(t *testing.T) {
	os.Setenv("S1", "s")
	var s string
	set := func() {
		s = fmt.Sprintf("%s %s", os.Getenv("S1"), os.Getenv("S2"))
	}

	go func() {
		os.Setenv("S2", "d")
	}()

	set()

	time.Sleep(20 * time.Millisecond)
	fmt.Println(s)

}
