package uiza

type AnalyticTypeFilter string

type AnalyticMetric string

const (
	AnalyticTypeFilterCountry AnalyticTypeFilter = "country"
	AnalyticTypeFilterDevice  AnalyticTypeFilter = "device"
	AnalyticTypeFilterTitle   AnalyticTypeFilter = "title"
	AnalyticTypeFilterPlayer  AnalyticTypeFilter = "player"
)

const (
	AnalyticMetricPlaybackFailureScore      AnalyticMetric = "playback_failure_score"
	AnalyticMetricPlaybackFailurePercentage AnalyticMetric = "playback_failure_percentage"
	AnalyticMetricPageLoadTime              AnalyticMetric = "page_load_time"
	AnalyticMetricVideoStartupTime          AnalyticMetric = "video_startup_time"
	AnalyticMetricPlayerStartupTime         AnalyticMetric = "player_startup_time"
	AnalyticMetricAggregateStartupTime      AnalyticMetric = "aggregate_startup_time"
	AnalyticMetricSeekLatency               AnalyticMetric = "seek_latency"
	AnalyticMetricExitsBeforeVideoStart     AnalyticMetric = "exits_before_video_start"
	AnalyticMetricRebufferPercentage        AnalyticMetric = "rebuffer_percentage"
	AnalyticMetricRebufferFrequency         AnalyticMetric = "rebuffer_frequency"
	AnalyticMetricRebufferDuration          AnalyticMetric = "rebuffer_duration"
	AnalyticMetricRebufferCount             AnalyticMetric = "rebuffer_count"
	AnalyticMetricUpscalePercentage         AnalyticMetric = "upscale_percentage"
	AnalyticMetricDownscalePercentage       AnalyticMetric = "downscale_percentage"
	AnalyticMetricMaxUpscalePercentage      AnalyticMetric = "max_upscale_percentage"
	AnalyticMetricMaxDownscalePercentage    AnalyticMetric = "max_downscale_percentage"
)

type AnalyticTotalLineParams struct {
	Params    `form:"*"`
	StartDate *string         `form:"start_date"`
	EndDate   *string         `form:"end_date"`
	Metric    *AnalyticMetric `form:"metric"`
}
type AnalyticTotalLineResponse struct {
	Data []*AnalyticTotalLineData `json:"data"`
}

type AnalyticTotalLine struct {
	PlaybackFailureScore      float64 `json:"playback_failure_score"`
	PlaybackFailurePercentage float64 `json:"playback_failure_percentage"`
	PageLoadTime              float64 `json:"page_load_time"`
	VideoStartupTime          float64 `json:"video_startup_time"`
	PlayerStartupTime         float64 `json:"player_startup_time"`
	AggregateStartupTime      float64 `json:"aggregate_startup_time"`
	SeekLatency               float64 `json:"seek_latency"`
	ExitsBeforeVideoStart     float64 `json:"exits_before_video_start"`
	RebufferPercentage        float64 `json:"rebuffer_percentage"`
	RebufferFrequency         float64 `json:"rebuffer_frequency"`
	RebufferDuration          float64 `json:"rebuffer_duration"`
	RebufferCount             float64 `json:"rebuffer_count"`
	UpscalePercentage         float64 `json:"upscale_percentage"`
	DownscalePercentage       float64 `json:"downscale_percentage"`
	MaxUpscalePercentage      float64 `json:"max_upscale_percentage"`
	MaxDownscalePercentage    float64 `json:"max_downscale_percentage"`
}
type AnalyticTotalLineData struct {
	DateTime int64 `json:"date_time"`
	AnalyticTotalLine
}

type AnalyticTypeParams struct {
	Params     `form:"*"`
	StartDate  *string             `form:"start_date"`
	EndDate    *string             `form:"end_date"`
	TypeFilter *AnalyticTypeFilter `form:"type_filter"`
}

type AnalyticTypeResponse struct {
	Data []*AnalyticTypeData `json:"data"`
}

type AnalyticTypeData struct {
	Name             string  `json:"name"`
	TotalView        int64   `json:"total_view"`
	PercentageOfView float64 `json:"percentage_of_view"`
}

type AnalyticLineParams struct {
	Params    `form:"*"`
	StartDate *string         `form:"start_date"`
	EndDate   *string         `form:"end_date"`
	Type      *AnalyticMetric `form:"type"`
}

type AnalyticLineResponse struct {
	Data []*AnalyticLineData `json:"data"`
}

type AnalyticLineData struct {
	DateTime int64   `json:"date_time"`
	Value    float64 `json:"value"`
}
