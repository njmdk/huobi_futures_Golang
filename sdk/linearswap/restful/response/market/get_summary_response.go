package market

type GetSummaryResponse struct {
	Page                  Page              `json:"page"`
	Components            []SystemComponent `json:"components"`
	Incidents             []Incident        `json:"incidents"`
	ScheduledMaintenances []Maintenance     `json:"scheduled_maintenances"`
	Status                SystemStatus      `json:"status"`
}

type Page struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	TimeZone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
}

type SystemComponent struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Position           int64  `json:"position"`
	Description        string `json:"description"`
	Showcase           bool   `json:"showcase"`
	GroupID            string `json:"group_id"`
	PageID             string `json:"page_id"`
	Group              bool   `json:"group"`
	OnlyShowIfDegraded bool   `json:"only_show_if_degraded"`
}

type Component struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Position           int    `json:"position"`
	Description        string `json:"description"`
	Showcase           bool   `json:"showcase"`
	GroupID            string `json:"group_id"`
	PageID             string `json:"page_id"`
	Group              bool   `json:"group"`
	OnlyShowIfDegraded bool   `json:"only_show_if_degraded"`
}

type Incident struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Status          string           `json:"status"`
	CreatedAt       string           `json:"created_at"`
	UpdatedAt       string           `json:"updated_at"`
	MonitoringAt    string           `json:"monitoring_at"`
	ResolvedAt      string           `json:"resolved_at"`
	Impact          string           `json:"impact"`
	Shortlink       string           `json:"shortlink"`
	StartedAt       string           `json:"started_at"`
	PageID          string           `json:"page_id"`
	IncidentUpdates []IncidentUpdate `json:"incident_updates"`
	Components      []Component      `json:"components"`
}

type IncidentUpdate struct {
	ID                 string `json:"id"`
	Status             string `json:"status"`
	Body               string `json:"body"`
	IncidentID         string `json:"incident_id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	DisplayAt          string `json:"display_at"`
	AffectedComponents []struct {
		Code      string `json:"code"`
		Name      string `json:"name"`
		OldStatus string `json:"old_status"`
		NewStatus string `json:"new_status"`
	} `json:"affected_components"`
	DeliverNotifications bool   `json:"deliver_notifications"`
	CustomTweet          string `json:"custom_tweet"`
	TweetID              string `json:"tweet_id"`
}

type Maintenance struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Status          string           `json:"status"`
	CreatedAt       string           `json:"created_at"`
	UpdatedAt       string           `json:"updated_at"`
	MonitoringAt    string           `json:"monitoring_at"`
	ResolvedAt      string           `json:"resolved_at"`
	Impact          string           `json:"impact"`
	Shortlink       string           `json:"shortlink"`
	StartedAt       string           `json:"started_at"`
	PageID          string           `json:"page_id"`
	IncidentUpdates []IncidentUpdate `json:"incident_updates"`
	Components      []Component      `json:"components"`
	ScheduledFor    string           `json:"scheduled_for"`
	ScheduledUntil  string           `json:"scheduled_until"`
}

type SystemStatus struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}
