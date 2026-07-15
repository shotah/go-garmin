package garmin

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestPersonalRecordsList(t *testing.T) {
	body := `[{"typeId":3,"value":1200,"activityId":1,"prStartTimeGmtFormatted":"2026-01-01 12:00:00"}]`
	client := testAuthedClient(t, roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if !strings.Contains(r.URL.Path, "/personalrecord-service/personalrecord/prs/anonymous") {
			t.Errorf("path = %s", r.URL.Path)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body)), Header: make(http.Header)}, nil
	}))
	prs, err := client.PersonalRecords.List(context.Background(), "anonymous")
	if err != nil || len(prs.Entries) != 1 || prs.Entries[0].TypeID != 3 {
		t.Fatalf("prs=%+v err=%v", prs, err)
	}
}

func TestLifestyleGetDaily(t *testing.T) {
	body := `{"dailyLogsReport":[{"behaviourId":1,"name":"Alcohol","logStatus":"YES","category":"LIFESTYLE"}],"completionStats":[]}`
	client := testAuthedClient(t, roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if !strings.HasSuffix(r.URL.Path, "/lifestylelogging-service/dailyLog/2026-07-14") {
			t.Errorf("path = %s", r.URL.Path)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body)), Header: make(http.Header)}, nil
	}))
	log, err := client.Lifestyle.GetDaily(context.Background(), time.Date(2026, 7, 14, 0, 0, 0, 0, time.UTC))
	if err != nil || len(log.DailyLogsReport) != 1 || log.DailyLogsReport[0].BehaviourID != 1 {
		t.Fatalf("log=%+v err=%v", log, err)
	}
}

func TestBloodPressureGetRange(t *testing.T) {
	body := `{"measurementSummaries":[{"startDate":"2026-07-01","endDate":"2026-07-15","numOfMeasurements":1,"measurements":[{"version":1,"systolic":120,"diastolic":80}]}],"categoryStats":null}`
	client := testAuthedClient(t, roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if !strings.Contains(r.URL.Path, "/bloodpressure-service/bloodpressure/range/") {
			t.Errorf("path = %s", r.URL.Path)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body)), Header: make(http.Header)}, nil
	}))
	bp, err := client.BloodPressure.GetRange(context.Background(), time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC), time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC))
	if err != nil || len(bp.MeasurementSummaries) != 1 {
		t.Fatalf("bp=%+v err=%v", bp, err)
	}
}

func TestTrainingPlanListJSON(t *testing.T) {
	raw := `{"trainingPlanList":[{"trainingPlanId":9,"name":"5K","trainingPlanCategory":"FBT_ADAPTIVE"}]}`
	var list TrainingPlanList
	if err := json.Unmarshal([]byte(raw), &list); err != nil {
		t.Fatal(err)
	}
	if len(list.TrainingPlanList) != 1 || list.TrainingPlanList[0].TrainingPlanCategory != "FBT_ADAPTIVE" {
		t.Fatalf("%+v", list)
	}
}

func TestBadgeListUnmarshalArrayOrWrap(t *testing.T) {
	var bare BadgeList
	if err := json.Unmarshal([]byte(`[{"badgeId":1,"badgeName":"A"}]`), &bare); err != nil || len(bare.Entries) != 1 {
		t.Fatalf("bare: %+v %v", bare, err)
	}
	var wrap BadgeList
	if err := json.Unmarshal([]byte(`{"badgeList":[{"badgeId":2,"badgeName":"B"}]}`), &wrap); err != nil || wrap.Entries[0].BadgeID != 2 {
		t.Fatalf("wrap: %+v %v", wrap, err)
	}
}

func TestHillScoreStatsAndRacePredictionsPaths(t *testing.T) {
	seen := map[string]bool{}
	client := testAuthedClient(t, roundTripFunc(func(r *http.Request) (*http.Response, error) {
		seen[r.URL.Path+"?"+r.URL.RawQuery] = true
		body := `{"avg":1,"groupMap":{}}`
		if strings.Contains(r.URL.Path, "racepredictions") {
			body = `[{"calendarDate":"2026-07-15","time5K":1200}]`
		}
		if strings.Contains(r.URL.Path, "fitnessage") {
			body = `{"fitnessAge":35,"chronologicalAge":40}`
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body)), Header: make(http.Header)}, nil
	}))
	ctx := context.Background()
	start := time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC)
	if _, err := client.Metrics.GetHillScoreStats(ctx, start, end, AggregationDaily); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Metrics.GetRacePredictionsDaily(ctx, "anonymous", start, end); err != nil {
		t.Fatal(err)
	}
	if _, err := client.FitnessAge.GetDaily(ctx, end); err != nil {
		t.Fatal(err)
	}
	if !seen["/metrics-service/metrics/hillscore/stats?startDate=2026-07-01&endDate=2026-07-15&aggregation=daily"] {
		t.Errorf("missing hill stats; seen=%v", seen)
	}
	if !seen["/metrics-service/metrics/racepredictions/daily/anonymous?fromCalendarDate=2026-07-01&toCalendarDate=2026-07-15"] {
		t.Errorf("missing race daily; seen=%v", seen)
	}
	if !seen["/fitnessage-service/fitnessage/2026-07-15?"] && !seen["/fitnessage-service/fitnessage/2026-07-15"] {
		// path key may not include ?
		found := false
		for k := range seen {
			if strings.Contains(k, "/fitnessage-service/fitnessage/2026-07-15") {
				found = true
			}
		}
		if !found {
			t.Errorf("missing fitness age daily; seen=%v", seen)
		}
	}
}
