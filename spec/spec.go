package spec

type (
	// GlobalOptions represents the global options for all groups in topology
	// specification in topology.yaml
	GlobalOptions struct {
		User       string      `yaml:"user,omitempty" default:"tidb"`
		Group      string      `yaml:"group,omitempty"`
		SSHPort    int         `yaml:"ssh_port,omitempty" default:"22" validate:"ssh_port:editable"`
		SSHType    string      `yaml:"ssh_type,omitempty" default:"builtin"`
		TLSEnabled bool        `yaml:"enable_tls,omitempty"`
		DeployDir  string      `yaml:"deploy_dir,omitempty" default:"deploy"`
		DataDir    string      `yaml:"data_dir,omitempty" default:"data"`
		LogDir     string      `yaml:"log_dir,omitempty"`
		OS         string      `yaml:"os,omitempty" default:"linux"`
		Arch       string      `yaml:"arch,omitempty"`
		Custom     interface{} `yaml:"custom,omitempty" validate:"custom:ignore"`
	}

	// MonitoredOptions represents the monitored node configuration
	MonitoredOptions struct {
		NodeExporterPort     int    `yaml:"node_exporter_port,omitempty" default:"9100"`
		BlackboxExporterPort int    `yaml:"blackbox_exporter_port,omitempty" default:"9115"`
		DeployDir            string `yaml:"deploy_dir,omitempty"`
		DataDir              string `yaml:"data_dir,omitempty"`
		LogDir               string `yaml:"log_dir,omitempty"`
		NumaNode             string `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	}

	// ServerConfigs represents the server runtime configuration
	ServerConfigs struct {
		TiDB           map[string]interface{} `yaml:"tidb"`
		TiKV           map[string]interface{} `yaml:"tikv"`
		PD             map[string]interface{} `yaml:"pd"`
		Dashboard      map[string]interface{} `yaml:"tidb_dashboard"`
		TiFlash        map[string]interface{} `yaml:"tiflash"`
		TiFlashLearner map[string]interface{} `yaml:"tiflash-learner"`
		Pump           map[string]interface{} `yaml:"pump"`
		Drainer        map[string]interface{} `yaml:"drainer"`
		CDC            map[string]interface{} `yaml:"cdc"`
		TiKVCDC        map[string]interface{} `yaml:"kvcdc"`
		Grafana        map[string]string      `yaml:"grafana"`
	}

	// Specification represents the specification of topology.yaml
	Specification struct {
		GlobalOptions    GlobalOptions        `yaml:"global,omitempty" validate:"global:editable"`
		MonitoredOptions MonitoredOptions     `yaml:"monitored,omitempty" validate:"monitored:editable"`
		ServerConfigs    ServerConfigs        `yaml:"server_configs,omitempty" validate:"server_configs:ignore"`
		TiDBServers      []*TiDBSpec          `yaml:"tidb_servers"`
		TiKVServers      []*TiKVSpec          `yaml:"tikv_servers"`
		TiFlashServers   []*TiFlashSpec       `yaml:"tiflash_servers"`
		PDServers        []*PDSpec            `yaml:"pd_servers"`
		DashboardServers []*DashboardSpec     `yaml:"tidb_dashboard_servers,omitempty"`
		PumpServers      []*PumpSpec          `yaml:"pump_servers,omitempty"`
		Drainers         []*DrainerSpec       `yaml:"drainer_servers,omitempty"`
		CDCServers       []*CDCSpec           `yaml:"cdc_servers,omitempty"`
		TiKVCDCServers   []*TiKVCDCSpec       `yaml:"kvcdc_servers,omitempty"`
		TiSparkMasters   []*TiSparkMasterSpec `yaml:"tispark_masters,omitempty"`
		TiSparkWorkers   []*TiSparkWorkerSpec `yaml:"tispark_workers,omitempty"`
		Monitors         []*PrometheusSpec    `yaml:"monitoring_servers"`
		Grafanas         []*GrafanaSpec       `yaml:"grafana_servers,omitempty"`
		Alertmanagers    []*AlertmanagerSpec  `yaml:"alertmanager_servers,omitempty"`
	}
)
