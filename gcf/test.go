package gcf

import (
	"fmt"
	"net/http"

	"github.com/getsentry/sentry-go"
)

func Test(w http.ResponseWriter, r *http.Request) {
	sentry.CaptureException(fmt.Errorf("dummy error"))
}
