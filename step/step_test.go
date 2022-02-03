package step

import (
	"testing"
	"time"

	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_runtime(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want string
	}{
		{
			name: "formats duration < 1m",
			d:    45 * time.Second,
			want: "45s",
		},
		{
			name: "formats duration < 1h",
			d:    (45 * time.Minute) + (30 * time.Second),
			want: "45m 30s",
		},
		{
			name: "formats duration > 1h",
			d:    (2 * time.Hour) + (45 * time.Minute) + (30 * time.Second),
			want: "2h 45m 30s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runtime(tt.d); got != tt.want {
				t.Errorf("runtime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEASBuilder_Run(t *testing.T) {
	client := new(MockEASClient)
	client.On("Build", mock.Anything).Return(nil).Once()

	clientBuilder := new(MockEASClientBuilder)
	clientBuilder.On("Build", mock.Anything, mock.Anything).Return(client).Once()

	step := NewEASBuilder(nil, log.NewLogger(), clientBuilder)
	err := step.Run(Config{})
	require.NoError(t, err)

	clientBuilder.AssertExpectations(t)
	client.AssertExpectations(t)
}
