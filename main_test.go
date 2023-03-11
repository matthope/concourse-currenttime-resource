package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_run(t *testing.T) {
	t.Parallel()

	type args struct {
		args []string
	}

	type want struct {
		err    bool
		stdout string
	}

	now := time.Date(2000, time.January, 1, 2, 3, 4, 5, time.UTC)

	tests := map[string]struct {
		args args
		want want
	}{
		"check": {
			args: args{args: []string{"check"}},
			want: want{err: false, stdout: "[{\"time\":\"2000-01-01T02:03:04Z\"}]\n"},
		},
		"out": {
			args: args{args: []string{"out"}},
			want: want{err: false, stdout: "{\"version\":{\"time\":\"2000-01-01T02:03:04Z\"}}\n"},
		},
		"in": {
			args: args{args: []string{"in"}},
			want: want{err: false, stdout: "{\"version\":{\"time\":\"2000-01-01T02:03:04Z\"}}\n"},
		},
		"unknown": {
			args: args{args: []string{"unknown"}},
			want: want{err: false, stdout: "{\"version\":{\"time\":\"2000-01-01T02:03:04Z\"}}\n"},
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			stdout := &bytes.Buffer{}

			if err := run(tt.args.args, now, stdout); (err != nil) != tt.want.err {
				t.Errorf("run() error = %v, wantErr %v", err, tt.want.err)

				return
			}

			assert.Equal(t, tt.want.stdout, stdout.String())
		})
	}
}
