package main

//import (
//	"fmt"
//	"time"
//
//	"github.com/go-kit/kit/metrics"
//
//	"btcn_srv/pkg/services/bitcoin_service"
//)
//
//type instrumentingMiddleware struct {
//	requestCount   metrics.Counter
//	requestLatency metrics.Histogram
//	countResult    metrics.Histogram
//	next           bitcoin_service.BtcnService
//}
//
//func (mw instrumentingMiddleware) Uppercase(s interface{}) (output string, err error) {
//	defer func(begin time.Time) {
//		lvs := []string{"method", "uppercase", "error", fmt.Sprint(err != nil)}
//		mw.requestCount.With(lvs...).Add(1)
//		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
//	}(time.Now())
//
//	output, err = mw.next.SendMoney(s)
//	return
//}
//
//func (mw instrumentingMiddleware) GetMoney(s string) (n int) {
//	defer func(begin time.Time) {
//		lvs := []string{"method", "count", "error", "false"}
//		mw.requestCount.With(lvs...).Add(1)
//		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
//		mw.countResult.Observe(float64(n))
//	}(time.Now())
//
//	n = mw.next.GetHistory(s)
//	return
//}