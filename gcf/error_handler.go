package gcf

import "github.com/getsentry/sentry-go"

func errorHandler(err error) error {
	if err != nil {
		sentry.CaptureException(err)
	}
	return err
}
