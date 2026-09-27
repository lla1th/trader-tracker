package main

import (
	"fmt"
	"time"
)

type TrustLevel int

const (
	TrustUnknown TrustLevel = iota
	TrustUntrusted
	TrustNeutral
	TrustTrusted
)

type Trader struct {
	SteamID     string
	Nickname    string
	Trust       TrustLevel
	WantedItems []string
}

type BanInfo struct {
	ExpiresAt time.Time
}

func main() {
	fmt.Println("Start")
}

func (t TrustLevel) String() string {
	switch t {
	case TrustUntrusted:
		return "untrusted"
	case TrustNeutral:
		return "neutral"
	case TrustTrusted:
		return "trusted"
	default:
		return "unknown"
	}
}

func (b BanInfo) IsExpired(now time.Time) bool {
	return !b.ExpiresAt.After(now)
}
