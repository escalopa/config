package config_go

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTempFile(t *testing.T) (*os.File, func() error) {
	file, err := os.Create("/tmp/test.txt")
	require.NoError(t, err)

	defer func() {
		err := file.Close()
		require.NoError(t, err)
	}()

	_, err = file.WriteString("value_file")
	require.NoError(t, err)

	return file, func() error { return os.Remove(file.Name()) }
}

func TestConfigGet(t *testing.T) {
	file, clean := createTempFile(t)
	defer clean()

	tests := []struct {
		name    string
		prepare func()
		clean   func() error
		key     string
		want    string
	}{
		{
			name: "using secret file",
			prepare: func() {
				_ = os.Setenv("key_file", file.Name())
			},
			clean: func() error { return os.Unsetenv("key_file") },
			key:   "key_file",
			want:  "value_file",
		},
		{
			name: "using env variable",
			prepare: func() {
				_ = os.Setenv("key", "value_env")
			},
			clean: func() error {
				return os.Unsetenv("key")
			},
			key:  "key",
			want: "value_env",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			defer func() { tt.clean() }()
			if got := Get(tt.key); got != tt.want {
				t.Errorf("Config.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
