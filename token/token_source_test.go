package token

import "testing"

func TestGetTokenURL(t *testing.T) {
	const want = "https://api.bot.qq.com/app/getAppAccessToken"
	if got := getTokenURL(); got != want {
		t.Fatalf("getTokenURL() = %q, want %q", got, want)
	}
}
