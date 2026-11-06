package profiling

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	pyroscope "github.com/grafana/pyroscope-go"
)

const (
	enabledEnvironmentVariable       = "PROFILING_ENABLED"
	serverAddressEnvironmentVariable = "PROFILING_SERVER_ADDRESS"
)

// Start enables continuous profiling only when the runtime flag is explicitly
// true. It uses push mode and does not register a public pprof HTTP endpoint.
func Start(serviceName string) (func() error, bool, error) {
	enabledValue := strings.TrimSpace(os.Getenv(enabledEnvironmentVariable))
	if enabledValue == "" {
		return nil, false, nil
	}
	enabled, err := strconv.ParseBool(enabledValue)
	if err != nil {
		return nil, false, fmt.Errorf("parse %s: %w", enabledEnvironmentVariable, err)
	}
	if !enabled {
		return nil, false, nil
	}

	serverAddress := strings.TrimSpace(os.Getenv(serverAddressEnvironmentVariable))
	if serverAddress == "" {
		return nil, true, fmt.Errorf("%s is required when %s=true", serverAddressEnvironmentVariable, enabledEnvironmentVariable)
	}

	profiler, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "courier." + serviceName,
		ServerAddress:   serverAddress,
		Tags: map[string]string{
			"service_name": serviceName,
			"environment":  environment(),
		},
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
	})
	if err != nil {
		return nil, true, fmt.Errorf("start pyroscope profiler: %w", err)
	}

	return profiler.Stop, true, nil
}

func environment() string {
	if value := strings.TrimSpace(os.Getenv("APP_ENV")); value != "" {
		return value
	}

	return "local"
}
