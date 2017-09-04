package sentry

import "github.com/getsentry/raven-go"

type Publisher struct {
	env       string
	SentryDSN string
}

func (p *Publisher) Publish(errs error, reqBody []byte, withStackTrace bool) {
	go func() {
		raven.CaptureError(errs, map[string]string{
			"env":   p.env,
			"error": errs.Error(),
		})
	}()
}
