package scheduler

import (
	"context"
	"fmt"
	"time"
)

func Run(p time.Duration, o time.Duration) {
	ctx := context.Background() //  .TODO()
	fmt.Println("Let's start:", time.Now())

	RunScheduler(ctx, p, o, func(t time.Time) {
		// fmt.Println("Time:", time.Now())

		message, open := IsOpenMarket()
		fmt.Println(message, open)
	})

}

// Schedule calls function `f` with a period `p` offsetted by `o`.
// Scheduler(ctx, time.Minute*2, time.Minute, fn()) Run every 2 minutes, starting 1 minute after the first run of this code
// if the code started 52:41.26, then first run will be at 53:00 followed by another run at 55:00, and so on
// Scheduler(ctx, time.Hour, time.Hour, func(t time.Time){}) Run every head of each hour, like 20:00:00, 21:00:00, ...
func RunScheduler(ctx context.Context, p time.Duration, o time.Duration, f func(time.Time)) {
	// Position the first execution
	first := time.Now().Truncate(p).Add(o)
	if first.Before(time.Now()) {
		first = first.Add(p)
	}
	firstC := time.After(first.Sub(time.Now()))

	// Receiving from a nil channel blocks forever
	t := &time.Ticker{C: nil}

	for {
		select {
		case v := <-firstC:
			// The ticker has to be started before f as it can take some time to finish
			t = time.NewTicker(p)
			f(v)
		case v := <-t.C:
			f(v)
		case <-ctx.Done():
			t.Stop()
			return
		}
	}

}

func IsOpenMarket() (string, bool) {
	var flag bool
	var message string

	now := time.Now()
	weekday := now.Weekday()
	hour := now.Hour()
	minute := now.Minute()

	if weekday == 0 || weekday == 6 {
		message = fmt.Sprintf("Haftasonu, Saat: %d:%d - Borsa kapali\n", hour, minute)
		flag = false
	} else {
		message = fmt.Sprintf("Haftaici, Saat: %d:%d - ", hour, minute)
		if hour < 9 || hour > 19 {
			message += fmt.Sprintf("Borsa kapali\n")
			flag = false
		} else if hour == 9 && minute < 55 {
			message += fmt.Sprintf("Borsa kapali\n")
			flag = false
		} else if hour == 18 && minute > 5 {
			message += fmt.Sprintf("Borsa kapali\n")
			flag = false
		} else {
			message += fmt.Sprintf("Borsa acik\n")
			flag = true
		}
	}

	return message, flag
}
