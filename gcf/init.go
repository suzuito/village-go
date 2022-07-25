package gcf

import (
	"context"
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/suzuito/village-go/pkg/inject"
	"github.com/suzuito/village-go/pkg/usecase"
)

var u *usecase.Usecase

func init() {
	ctx := context.Background()
	var err error
	err = sentry.Init(sentry.ClientOptions{
		TracesSampleRate: 1.0,
	})
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	u, err = inject.NewUsecase(ctx)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
