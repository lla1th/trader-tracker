package main

import (
	"fmt"
	"time"
)

type TrustLevel int

type Trader struct {
	SteamID string
	Nickname string
	Trust TrustLevel
	WantedItems []string
}

type BanInfo struct {
	Active bool
	ExpiresAt time.Time
}

func main() {
	fmt.Println("Hello world!")
}
