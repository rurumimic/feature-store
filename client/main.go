package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type getOnlineFeaturesRequest struct {
	Features       []string         `json:"features,omitempty"`
	FeatureService string           `json:"feature_service,omitempty"`
	Entities       map[string][]any `json:"entities"`
}

func main() {
	addr := flag.String("addr", "http://localhost:6566", "Feast feature server base URL")
	featuresFlag := flag.String("features", "", "Comma-separated feature references")
	featureService := flag.String("feature-service", "", "Feature service name (e.g., driver_activity_v1)")
	timeout := flag.Duration("timeout", 5*time.Second, "Request timeout")
	flag.Parse()

	features := defaultFeatures()
	if *featuresFlag != "" {
		features = splitCSV(*featuresFlag)
	}

	reqBody := getOnlineFeaturesRequest{
		Features: features,
		Entities: defaultEntities(),
	}
	if *featureService != "" {
		reqBody.Features = nil
		reqBody.FeatureService = *featureService
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		exitf("marshal request: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(*addr, "/")+"/get-online-features", bytes.NewReader(payload))
	if err != nil {
		exitf("create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		exitf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		exitf("read response: %v", err)
	}

	if resp.StatusCode >= 300 {
		exitf("request failed with %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}

	fmt.Println(string(body))
}

func defaultFeatures() []string {
	return []string{
		"driver_hourly_stats:acc_rate",
		"transformed_conv_rate:conv_rate_plus_val1",
		"transformed_conv_rate:conv_rate_plus_val2",
	}
}

func defaultEntities() map[string][]any {
	return map[string][]any{
		"driver_id":    {1001, 1002},
		"val_to_add":   {1000, 1001},
		"val_to_add_2": {2000, 2002},
	}
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item == "" {
			continue
		}
		out = append(out, item)
	}
	return out
}

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
