package annotations

const (
	// annotations key
	enableEnvoyAccessLoggers = "enable-envoy-access-loggers"
	envoyAccessLogPath       = "envoy-access-log-path"
	//envoyAccessLogFormat     = "envoy-access-log-format"

	// annotation value
	defaultEnvoyAccessLogPath = "/var/log/proxy/proxy.log"
	//defaultEnvoyAccessLogFormat = `
	//    {
	//        "authority": "%REQ(:AUTHORITY)%",
	//        "bytes_received": "%BYTES_RECEIVED%",
	//        "bytes_sent": "%BYTES_SENT%",
	//        "downstream_local_address": "%DOWNSTREAM_LOCAL_ADDRESS%",
	//        "downstream_remote_address": "%DOWNSTREAM_REMOTE_ADDRESS%",
	//        "duration": "%DURATION%",
	//        "method": "%REQ(:METHOD)%",
	//        "path": "%REQ(X-ENVOY-ORIGINAL-PATH?:PATH)%",
	//        "protocol": "%PROTOCOL%",
	//        "request_id": "%REQ(X-REQUEST-ID)%",
	//        "requested_server_name": "%REQUESTED_SERVER_NAME%",
	//        "response_code": "%RESPONSE_CODE%",
	//        "response_flags": "%RESPONSE_FLAGS%",
	//        "route_name": "%ROUTE_NAME%",
	//        "start_time": "%START_TIME%",
	//        "trace_id": "%REQ(X-B3-TRACEID)%",
	//        "upstream_cluster": "%UPSTREAM_CLUSTER%",
	//        "upstream_host": "%UPSTREAM_HOST%",
	//        "upstream_local_address": "%UPSTREAM_LOCAL_ADDRESS%",
	//        "upstream_service_time": "%RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)%",
	//        "upstream_transport_failure_reason": "%UPSTREAM_TRANSPORT_FAILURE_REASON%",
	//        "user_agent": "%REQ(USER-AGENT)%",
	//        "x_forwarded_for": "%REQ(X-FORWARDED-FOR)%"
	//    }
	//`
)

var _ Parser = &log{}

type log struct {
	Enabled bool
	LogPath string
}

func (l log) Parse(annotations Annotations, config *Ingress, _ *GlobalContext) error {
	// check enableEnvoyAccessLoggers first
	if !needLogConfig(annotations) {
		return nil
	}

	// enable envoy access loggers
	enable, _ := annotations.ParseBoolForHigress(enableEnvoyAccessLoggers)
	if !enable {
		return nil
	}

	logConfig := &log{
		Enabled: true,
		LogPath: defaultEnvoyAccessLogPath,
	}
	defer func() {
		config.Log = logConfig
	}()

	// access log path
	if logPath, err := annotations.ParseStringForHigress(envoyAccessLogPath); err == nil {
		logConfig.LogPath = logPath
	}
	return nil
}

func needLogConfig(annotations Annotations) bool {
	return annotations.HasHigress(enableEnvoyAccessLoggers)
}
