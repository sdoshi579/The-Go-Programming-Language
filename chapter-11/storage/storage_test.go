package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	const user = "joe@example.org"
	usage[user] = 980 * 1000 * 1000
	checkQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified want %s", notifiedUser, user)
	}
	if !strings.Contains(notifiedMsg, "98% of your quota") {
		t.Errorf("unexpected notified msg <<%s>>", notifiedMsg)
	}

}
