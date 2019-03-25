## Total Line
Get data grouped by hour (data refresh every 5 minutes). Track video playback on any metric performance, so you can know exactly what’s happening on every user’s device and debug more effectively.

About grouped by hour algorithm, Uiza currently support up to 16 days (it means when your time range is lower than 16 days, data response will be grouped by hour. Otherwise, it will return and to be grouped by day). In case your requested timerange doesn't have data, API won't show it in response.

See details [here](https://docs.uiza.io/#total-line).

```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/analytic"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

metric := uiza.AnalyticMetricRebufferCount
params := &uiza.AnalyticTotalLineParams{
	StartDate:  uiza.String("2018-11-01 08:00"),
	EndDate:    uiza.String("2019-11-19 14:00"),
	Metric:    &metric,
}

response, _ := analytic.GetTotalLine(params)
for 1_, v := range response {
	log.Printf("%v", v.DateTime)
	log.Printf("%v", v.RebufferCount)
}
```
Example Response
```golang
[
   {
       "date_time": 1551312000000,
       "rebuffer_count": 372.55555555555554
   },
   {
       "date_time": 1551398400000,
       "rebuffer_count": 0.9090909090909091
   }
]
```
## Type
Get data base on 4 type of filter: country, device, title, player

See details [here](https://docs.uiza.io/#type).

```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/analytic"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

analyticTypeFilter := uiza.AnalyticTypeFilterCountry
params := &uiza.AnalyticTypeParams{
	StartDate: uiza.String("2019-01-01"),
	EndDate:   uiza.String("2019-02-28"),
	TypeFilter: &analyticTypeFilter,
}
response, _ := analytic.GetType(params)
for _, v := range response {
		log.Printf("%v\n", v)
	}
```

Example Response
```golang
[
   {
       "name": "Vietnam",
       "total_view": 15,
       "percentage_of_view": 0.625
   },
   {
       "name": "Other",
       "total_view": 9,
       "percentage_of_view": 0.375
   }
]
```
## Line
Get data grouped by hour. Get total view in time range. This help you to draw a line chart to visualize data

About grouped by hour algorithm, Uiza currently support upto 16 days (it's mean when your time range is lower than 16 days, data response will be grouped by hour. Otherwise, it will return and to be grouped by day). In case your requested timerange doesn't have data, API won't show it in response.

See details [here](https://docs.uiza.io/#line).

```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/analytic"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

rebufferCount := uiza.AnalyticMetricRebufferCount
params := &uiza.AnalyticLineParams{
	StartDate: uiza.String("2019-01-01"),
	EndDate:   uiza.String("2019-02-28"),
	Type:&rebufferCount,
}
response, _ := analytic.GetLine(params)
for _, v := range response {
	log.Printf("%v\n", v)
}
```
Example Response
```golang
[
    {
        "date_time": 1551312000000,
        "value": 372.55555555555554
    }
]
```