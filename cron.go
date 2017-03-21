package cron

import (
	"github.com/golang-plus/errors"
	"github.com/golang-plus/jobs"

	"github.com/gorhill/cronexpr"
)

// Expression represents a cron expression.
type Expression struct {
	*cronexpr.Expression
}

func predefined(expr string) jobs.Expression {
	return &Expression{
		Expression: cronexpr.MustParse(expr),
	}
}

// Predefined cron expressions.
var (
	Annually = predefined("@annually") // every year at midnight in the morning of January 1
	Yearly   = predefined("@yearly")   // every year at midnight in the morning of January 1
	Monthly  = predefined("@monthly")  // every month at midnight in the morning of the first of the month
	Weekly   = predefined("@weekly")   // every week at midnight in the morning of Sunday
	Daily    = predefined("@daily")    // every day at midnight
	Hourly   = predefined("@hourly")   // every hour at the beginning of the hour
)

// New returns a new cron expression.
// We use github.com/gorhill/cronexpr for parsing cron express inside.
func New(expr string) (jobs.Expression, error) {
	cronexpr, err := cronexpr.Parse(expr)
	if err != nil {
		return nil, errors.Wrapf(err, "could not parse cron expression %q", expr)
	}

	return &Expression{
		Expression: cronexpr,
	}, nil
}
