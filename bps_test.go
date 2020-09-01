package bpsrestgo

import (
	"fmt"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {

	bps := BPS(os.Getenv("BPS_HOST"), os.Getenv("BPS_USER"), os.Getenv("BPS_PASS"))
	err := bps.connect()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("BPS("+bps.host+") connection attempted, resulted in session:", bps.sessionID, "and API key:", bps.apiKey)

	err := bps.disconnect()
	if err != nil {
		fmt.Println(err)
	}

}
