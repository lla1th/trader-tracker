package main

import (
	"testing"
	"time"
)

func TestBanInfo_IsExpired(t *testing.T) {
	base := time.Date(2026, 9, 28, 22, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		ban  BanInfo
		now  time.Time
		want bool
	}{
		{
			name: "well before expiry",
			ban:  BanInfo{ExpiresAt: base},
			now:  base.Add(-72 * time.Hour),
			want: false,
		},
		{
			name: "just before expiry",
			ban:  BanInfo{ExpiresAt: base},
			now:  base.Add(-time.Nanosecond),
			want: false,
		},
		{
			name: "exactly at expiry",
			ban:  BanInfo{ExpiresAt: base},
			now:  base,
			want: true,
		},
		{
			name: "after expiry",
			ban:  BanInfo{ExpiresAt: base},
			now:  base.Add(24 * time.Hour),
			want: true,
		},
		{
			name: "zero value ban",
			ban:  BanInfo{},
			now:  base,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ban.IsExpired(tt.now)
			if got != tt.want {
				t.Errorf("IsExpired(now=%v) = %v, want %v", tt.now, got, tt.want)
			}
		})
	}
}
