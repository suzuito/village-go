package gcf

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

func Test() {
	sentry.CaptureException(fmt.Errorf("dummy error"))
}
