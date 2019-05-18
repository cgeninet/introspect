package introspect

import (
	"encoding/json"
	"testing"
)

type Period struct {
	Seconds *json.Number `json:"seconds,omitempty"`
	Text    *string      `json:"text,omitempty"`
	Value   *string      `json:"value,omitempty"`
	Name    *string      `json:"name,omitempty"`
	Unit    *string      `json:"unit,omitempty"`
}

type LogSet struct {
	ID   *json.Number `json:"id,omitempty"`
	Name *string      `json:"name,omitempty"`
}

type TimeRange struct {
	To   *json.Number `json:"to,omitempty"`
	From *json.Number `json:"from,omitempty"`
	Live *bool        `json:"live,omitempty"`
}

type QueryConfig struct {
	LogSet        *LogSet    `json:"logset,omitempty"`
	TimeRange     *TimeRange `json:"timeRange,omitempty"`
	QueryString   *string    `json:"queryString,omitempty"`
	QueryIsFailed *bool      `json:"queryIsFailed,omitempty"`
}

type ThresholdCount struct {
	Ok               *json.Number `json:"ok,omitempty"`
	Critical         *json.Number `json:"critical,omitempty"`
	Warning          *json.Number `json:"warning,omitempty"`
	Unknown          *json.Number `json:"unknown,omitempty"`
	CriticalRecovery *json.Number `json:"critical_recovery,omitempty"`
	WarningRecovery  *json.Number `json:"warning_recovery,omitempty"`
	Period           *Period      `json:"period,omitenmpty"`
	TimeAggregator   *string      `json:"timeAggregator,omitempty"`
}

type ThresholdWindows struct {
	RecoveryWindow *string `json:"recovery_window,omitempty"`
	TriggerWindow  *string `json:"trigger_window,omitempty"`
}

type NoDataTimeframe int

type Options struct {
	NoDataTimeframe   NoDataTimeframe   `json:"no_data_timeframe,omitempty"`
	NotifyAudit       bool              `json:"notify_audit,omitempty"`
	NotifyNoData      bool              `json:"notify_no_data,omitempty"`
	RenotifyInterval  int               `json:"renotify_interval,omitempty"`
	NewHostDelay      int               `json:"new_host_delay,omitempty"`
	EvaluationDelay   int               `json:"evaluation_delay,omitempty"`
	Silenced          map[string]int    `json:"silenced,omitempty"`
	TimeoutH          int               `json:"timeout_h,omitempty"`
	EscalationMessage string            `json:"escalation_message,omitempty"`
	Thresholds        *ThresholdCount   `json:"thresholds,omitempty"`
	ThresholdWindows  *ThresholdWindows `json:"threshold_windows,omitempty"`
	IncludeTags       bool              `json:"include_tags,omitempty"`
	RequireFullWindow bool              `json:"require_full_window,omitempty"`
	Locked            bool              `json:"locked,omitempty"`
	EnableLogsSample  bool              `json:"enable_logs_sample,omitempty"`
	QueryConfig       *QueryConfig      `json:"queryConfig,omitempty"`
}

type TriggeringValue struct {
	FromTs *int `json:"from_ts,omitempty"`
	ToTs   *int `json:"to_ts,omitempty"`
	Value  *int `json:"value,omitempty"`
}

type GroupData struct {
	LastNoDataTs    int              `json:"last_nodata_ts,omitempty"`
	LastNotifiedTs  int              `json:"last_notified_ts,omitempty"`
	LastResolvedTs  int              `json:"last_resolved_ts,omitempty"`
	LastTriggeredTs int              `json:"last_triggered_ts,omitempty"`
	Name            string           `json:"name,omitempty"`
	Status          string           `json:"status,omitempty"`
	TriggeringValue *TriggeringValue `json:"triggering_value,omitempty"`
}

type State struct {
	Groups map[string]GroupData `json:"groups,omitempty"`
}

type Monitor struct {
	Creator              *Creator `json:"creator,omitempty"`
	Type                 string   `json:"type,omitempty"`
	Query                string   `json:"query,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Message              string   `json:"message,omitempty"`
	OverallState         string   `json:"overall_state,omitempty"`
	OverallStateModified string   `json:"overall_state_modified,omitempty"`
	Tags                 []string `json:"tags"`
	Options              *Options `json:"options,omitempty"`
	State                State    `json:"state,omitempty"`
}

type Creator struct {
	Email  string `json:"email,omitempty"`
	Handle string `json:"handle,omitempty"`
	Name   string `json:"name,omitempty"`
}

func TestStruct(t *testing.T) {

	o := &Options{
		NotifyNoData:      true,
		NotifyAudit:       false,
		Locked:            false,
		NoDataTimeframe:   120,
		NewHostDelay:      600,
		RequireFullWindow: true,
		Silenced:          map[string]int{},
	}
	f := &Monitor{
		Message:      "Test message",
		Query:        "max(last_1m):max:custom.zookeeper.isleader{env:prod} < 1",
		Name:         "Monitor name",
		Options:      o,
		Type:         "metric alert",
		Tags:         make([]string, 0),
		OverallState: "No Data",
	}

	m := NewStruct(f)
	k := m.Keys()

	if m.TypeOf("NIL") != "nil" {
		t.Errorf("TypeOf should be nil for unknown path, but got %s", m.TypeOf("NIL"))
	}
	if m.TypeOf("Monitor.Name") != "string" {
		t.Errorf("TypeOf should be string, but got %s", m.TypeOf("Monitor.Name"))
	}
	if m.Value("Monitor.Name") != "Monitor name" {
		t.Errorf("Value should be 'Monitor name', but got %s", m.Value("Monitor.Name"))
	}
	if len(k) != 21 {
		t.Errorf("Number of keys should be 21, but got %d", len(k))
	}
}
