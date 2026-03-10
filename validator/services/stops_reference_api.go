package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"main/lib"
)

type stopsReferenceAPIResponse struct {
	Data    []map[string]any `json:"data"`
	StopIds []string         `json:"stop_ids"`
}

var (
	stopsReferenceAPIMutex sync.RWMutex
	stopsReferenceAPISet   map[string]struct{}
	debugNoCacheOnce       sync.Once
	debugCacheActiveOnce   sync.Once
)

func IsStopIDAllowedByReferenceAPI(stopID string) (allowed bool, enabled bool) {
	stopsReferenceAPIMutex.RLock()
	defer stopsReferenceAPIMutex.RUnlock()

	if stopsReferenceAPISet == nil {
		// #region agent log
		debugLog("run-1", "H4", "services/stops_api.go:IsStopIDAllowedByReferenceAPI", "API cache not enabled; skipping external stop check", map[string]any{
			"cacheEnabled": false,
		}, &debugNoCacheOnce)
		// #endregion
		return true, false
	}

	_, ok := stopsReferenceAPISet[stopID]
	// #region agent log
	debugLog("run-1", "H4", "services/stops_api.go:IsStopIDAllowedByReferenceAPI", "API cache enabled; performing external stop check", map[string]any{
		"cacheEnabled": true,
		"cacheSize":    len(stopsReferenceAPISet),
	}, &debugCacheActiveOnce)
	// #endregion
	return ok, true
}

func LoadStopsFromAPI() error {
	apiURL := strings.TrimSpace(os.Getenv("STOPS_REFERENCE_API_URL"))
	// #region agent log
	debugLogAlways("run-1", "H2", "services/stops_api.go:LoadStopsReferenceFromAPI", "Starting API preload for stop ", map[string]any{
		"apiURLConfigured":  apiURL != "",
		"apiURLLength":      len(apiURL),
		"timeoutConfigured": strings.TrimSpace(os.Getenv("STOPS_REFERENCE_API_TIMEOUT")) != "",
	})
	// #endregion
	if apiURL == "" {
		lib.AppLogger.Debug("STOPS_REFERENCE_API_URL not set; skipping API reference preload.")
		return nil
	}

	timeout := 5 * time.Second
	if rawTimeout := strings.TrimSpace(os.Getenv("STOPS_REFERENCE_API_TIMEOUT")); rawTimeout != "" {
		parsed, err := time.ParseDuration(rawTimeout)
		if err != nil {
			disableStopsReferenceCache("invalid STOPS_REFERENCE_API_TIMEOUT")
			lib.AppLogger.Info(fmt.Sprintf("WARNING: Invalid STOPS_REFERENCE_API_TIMEOUT; skipping API reference preload: %v", err))
			return nil
		}
		timeout = parsed
	}

	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		disableStopsReferenceCache("failed to create API request")
		lib.AppLogger.Info(fmt.Sprintf("WARNING: Failed to create API request; skipping API reference preload: %v", err))
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		// #region agent log
		debugLogAlways("run-1", "H3", "services/stops_api.go:LoadStopsReferenceFromAPI", "API preload request failed", map[string]any{
			"error": err.Error(),
		})
		// #endregion
		disableStopsReferenceCache("failed to call API")
		lib.AppLogger.Info(fmt.Sprintf("WARNING: Failed to call API %q; skipping API reference preload: %v", apiURL, err))
		return nil
	}
	defer res.Body.Close()

	// #region agent log
	debugLogAlways("run-1", "H3", "services/stops_api.go:LoadStopsReferenceFromAPI", "API preload response received", map[string]any{
		"statusCode": res.StatusCode,
	})
	// #endregion
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		disableStopsReferenceCache("non-2xx API response")
		lib.AppLogger.Info(fmt.Sprintf("WARNING: API %q returned status %d; skipping API reference preload", apiURL, res.StatusCode))
		return nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		disableStopsReferenceCache("failed to read API response body")
		lib.AppLogger.Info(fmt.Sprintf("WARNING: Failed to read API response body; skipping API reference preload: %v", err))
		return nil
	}

	set, err := parseStopReferencePayload(body)
	if err != nil {
		// #region agent log
		debugLogAlways("run-1", "H3", "services/stops_api.go:LoadStopsReferenceFromAPI", "API preload payload parse failed", map[string]any{
			"error": err.Error(),
		})
		// #endregion
		disableStopsReferenceCache("failed to parse API response body")
		lib.AppLogger.Info(fmt.Sprintf("WARNING: Failed to parse API response body; skipping API reference preload: %v", err))
		return nil
	}

	stopsReferenceAPIMutex.Lock()
	stopsReferenceAPISet = set
	stopsReferenceAPIMutex.Unlock()
	// #region agent log
	debugLogAlways("run-1", "H3", "services/stops_api.go:LoadStopsReferenceFromAPI", "API preload cache populated", map[string]any{
		"cacheSize": len(set),
	})
	// #endregion

	lib.AppLogger.Info(fmt.Sprintf("Loaded %d stop reference IDs from API.", len(set)))
	return nil
}

func parseStopReferencePayload(body []byte) (map[string]struct{}, error) {
	var directList []string
	if err := json.Unmarshal(body, &directList); err == nil {
		return toStopSet(directList), nil
	}

	var payload stopsReferenceAPIResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}

	if len(payload.StopIds) > 0 {
		return toStopSet(payload.StopIds), nil
	}

	if len(payload.Data) > 0 {
		ids := make([]string, 0, len(payload.Data))
		for _, item := range payload.Data {
			rawID, ok := item["stop_id"]
			if !ok {
				continue
			}
			id, ok := rawID.(string)
			if !ok || strings.TrimSpace(id) == "" {
				continue
			}
			ids = append(ids, strings.TrimSpace(id))
		}
		return toStopSet(ids), nil
	}

	return map[string]struct{}{}, nil
}

func toStopSet(values []string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue
		}
		set[trimmed] = struct{}{}
	}
	return set
}

func disableStopsReferenceCache(reason string) {
	stopsReferenceAPIMutex.Lock()
	stopsReferenceAPISet = nil
	stopsReferenceAPIMutex.Unlock()
	debugLogAlways("run-1", "H7", "services/stops_api.go:disableStopsReferenceCache", "API reference cache disabled", map[string]any{
		"reason": reason,
	})
}

func debugLog(runID, hypothesisID, location, message string, data map[string]any, once *sync.Once) {
	once.Do(func() {
		debugLogAlways(runID, hypothesisID, location, message, data)
	})
}

func debugLogAlways(runID, hypothesisID, location, message string, data map[string]any) {
	f, err := os.OpenFile("/home/joao/tml/validator/.cursor/debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return
	}
	defer f.Close()

	payload := map[string]any{
		"runId":        runID,
		"hypothesisId": hypothesisID,
		"location":     location,
		"message":      message,
		"data":         data,
		"timestamp":    time.Now().UnixMilli(),
	}
	encoded, err := json.Marshal(payload)
	if err != nil {
		return
	}
	_, _ = f.Write(append(encoded, '\n'))
}
