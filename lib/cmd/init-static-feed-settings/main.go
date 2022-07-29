package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/suzuito/village-go/pkg/inject"
)

func main() {
	ctx := context.Background()
	err := sentry.Init(sentry.ClientOptions{
		TracesSampleRate: 1.0,
	})
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	defer sentry.Flush(2 * time.Second)
	u, err := inject.NewUsecase(ctx)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	if err := u.InitStaticFeedSettings(ctx); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
