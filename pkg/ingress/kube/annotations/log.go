package annotations

const (
    // annotations key
    enableEnvoyAccessLoggers = "enable-envoy-access-loggers"
    logPath                  = ""
)

type log struct {
    Path string
}

func (l log) Parse(annotations Annotations, config *Ingress, _ *GlobalContext) error {

}

func needLogConfig(annotations Annotations) bool {
    return false
}
