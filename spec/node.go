package spec

// PDSpec represents the PD topology specification in topology.yaml
type PDSpec struct {
	Host                string `yaml:"host"`
	ListenHost          string `yaml:"listen_host,omitempty"`
	AdvertiseClientAddr string `yaml:"advertise_client_addr,omitempty"`
	AdvertisePeerAddr   string `yaml:"advertise_peer_addr,omitempty"`
	SSHPort             int    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported            bool   `yaml:"imported,omitempty"`
	Patched             bool   `yaml:"patched,omitempty"`
	IgnoreExporter      bool   `yaml:"ignore_exporter,omitempty"`
	// Use Name to get the name with a default value if it's empty.
	Name       string                 `yaml:"name,omitempty"`
	ClientPort int                    `yaml:"client_port,omitempty" default:"2379"`
	PeerPort   int                    `yaml:"peer_port,omitempty" default:"2380"`
	DeployDir  string                 `yaml:"deploy_dir,omitempty"`
	DataDir    string                 `yaml:"data_dir,omitempty"`
	LogDir     string                 `yaml:"log_dir,omitempty"`
	NumaNode   string                 `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	Config     map[string]interface{} `yaml:"config,omitempty" validate:"config:ignore"`
	Arch       string                 `yaml:"arch,omitempty"`
	OS         string                 `yaml:"os,omitempty"`
}

// TiDBSpec represents the TiDB topology specification in topology.yaml
type TiDBSpec struct {
	Host           string                 `yaml:"host"`
	ListenHost     string                 `yaml:"listen_host,omitempty"`
	AdvertiseAddr  string                 `yaml:"advertise_address,omitempty"`
	SSHPort        int                    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported       bool                   `yaml:"imported,omitempty"`
	Patched        bool                   `yaml:"patched,omitempty"`
	IgnoreExporter bool                   `yaml:"ignore_exporter,omitempty"`
	Port           int                    `yaml:"port,omitempty" default:"4000"`
	StatusPort     int                    `yaml:"status_port,omitempty" default:"10080"`
	DeployDir      string                 `yaml:"deploy_dir,omitempty"`
	LogDir         string                 `yaml:"log_dir,omitempty"`
	NumaNode       string                 `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	NumaCores      string                 `yaml:"numa_cores,omitempty" validate:"numa_cores:editable"`
	Config         map[string]interface{} `yaml:"config,omitempty" validate:"config:ignore"`
	Arch           string                 `yaml:"arch,omitempty"`
	OS             string                 `yaml:"os,omitempty"`
}

// TiKVSpec represents the TiKV topology specification in topology.yaml
type TiKVSpec struct {
	Host                string                 `yaml:"host"`
	ListenHost          string                 `yaml:"listen_host,omitempty"`
	AdvertiseAddr       string                 `yaml:"advertise_addr,omitempty"`
	SSHPort             int                    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported            bool                   `yaml:"imported,omitempty"`
	Patched             bool                   `yaml:"patched,omitempty"`
	IgnoreExporter      bool                   `yaml:"ignore_exporter,omitempty"`
	Port                int                    `yaml:"port,omitempty" default:"20160"`
	StatusPort          int                    `yaml:"status_port,omitempty" default:"20180"`
	AdvertiseStatusAddr string                 `yaml:"advertise_status_addr,omitempty"`
	DeployDir           string                 `yaml:"deploy_dir,omitempty"`
	DataDir             string                 `yaml:"data_dir,omitempty"`
	LogDir              string                 `yaml:"log_dir,omitempty"`
	Offline             bool                   `yaml:"offline,omitempty"`
	NumaNode            string                 `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	NumaCores           string                 `yaml:"numa_cores,omitempty" validate:"numa_cores:editable"`
	Config              map[string]interface{} `yaml:"config,omitempty" validate:"config:ignore"`
	Arch                string                 `yaml:"arch,omitempty"`
	OS                  string                 `yaml:"os,omitempty"`
}

// TiFlashSpec represents the TiFlash topology specification in topology.yaml
type TiFlashSpec struct {
	Host                 string                 `yaml:"host"`
	SSHPort              int                    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported             bool                   `yaml:"imported,omitempty"`
	Patched              bool                   `yaml:"patched,omitempty"`
	IgnoreExporter       bool                   `yaml:"ignore_exporter,omitempty"`
	TCPPort              int                    `yaml:"tcp_port,omitempty" default:"9000"`
	HTTPPort             int                    `yaml:"http_port,omitempty" default:"8123"`
	FlashServicePort     int                    `yaml:"flash_service_port,omitempty" default:"3930"`
	FlashProxyPort       int                    `yaml:"flash_proxy_port,omitempty" default:"20170"`
	FlashProxyStatusPort int                    `yaml:"flash_proxy_status_port,omitempty" default:"20292"`
	StatusPort           int                    `yaml:"metrics_port,omitempty" default:"8234"`
	DeployDir            string                 `yaml:"deploy_dir,omitempty"`
	DataDir              string                 `yaml:"data_dir,omitempty" validate:"data_dir:expandable"`
	LogDir               string                 `yaml:"log_dir,omitempty"`
	TmpDir               string                 `yaml:"tmp_path,omitempty"`
	Offline              bool                   `yaml:"offline,omitempty"`
	NumaNode             string                 `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	NumaCores            string                 `yaml:"numa_cores,omitempty" validate:"numa_cores:editable"`
	Config               map[string]interface{} `yaml:"config,omitempty" validate:"config:ignore"`
	LearnerConfig        map[string]interface{} `yaml:"learner_config,omitempty" validate:"learner_config:ignore"`
	Arch                 string                 `yaml:"arch,omitempty"`
	OS                   string                 `yaml:"os,omitempty"`
}

// PumpSpec represents the Pump topology specification in topology.yaml
type PumpSpec struct {
	Host           string                 `yaml:"host"`
	SSHPort        int                    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported       bool                   `yaml:"imported,omitempty"`
	Patched        bool                   `yaml:"patched,omitempty"`
	IgnoreExporter bool                   `yaml:"ignore_exporter,omitempty"`
	Port           int                    `yaml:"port,omitempty" default:"8250"`
	DeployDir      string                 `yaml:"deploy_dir,omitempty"`
	DataDir        string                 `yaml:"data_dir,omitempty"`
	LogDir         string                 `yaml:"log_dir,omitempty"`
	Offline        bool                   `yaml:"offline,omitempty"`
	NumaNode       string                 `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	Config         map[string]interface{} `yaml:"config,omitempty" validate:"config:ignore"`
	Arch           string                 `yaml:"arch,omitempty"`
	OS             string                 `yaml:"os,omitempty"`
}

// DrainerSpec represents the Drainer topology specification in topology.yaml
type DrainerSpec struct {
	Host           string                 `yaml:"host"`
	SSHPort        int                    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported       bool                   `yaml:"imported,omitempty"`
	Patched        bool                   `yaml:"patched,omitempty"`
	IgnoreExporter bool                   `yaml:"ignore_exporter,omitempty"`
	Port           int                    `yaml:"port,omitempty" default:"8249"`
	DeployDir      string                 `yaml:"deploy_dir,omitempty"`
	DataDir        string                 `yaml:"data_dir,omitempty"`
	LogDir         string                 `yaml:"log_dir,omitempty"`
	CommitTS       *int64                 `yaml:"commit_ts,omitempty" validate:"commit_ts:editable"` // do not use it anymore, exist for compatibility
	Offline        bool                   `yaml:"offline,omitempty"`
	NumaNode       string                 `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	Config         map[string]interface{} `yaml:"config,omitempty" validate:"config:ignore"`
	Arch           string                 `yaml:"arch,omitempty"`
	OS             string                 `yaml:"os,omitempty"`
}

// TiKVCDCSpec represents the TiKVCDC topology specification in topology.yaml
type TiKVCDCSpec struct {
	Host           string `yaml:"host"`
	SSHPort        int    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported       bool   `yaml:"imported,omitempty"`
	Patched        bool   `yaml:"patched,omitempty"`
	IgnoreExporter bool   `yaml:"ignore_exporter,omitempty"`
	Port           int    `yaml:"port,omitempty" default:"8600"`
	DeployDir      string `yaml:"deploy_dir,omitempty"`
	DataDir        string `yaml:"data_dir,omitempty"`
	LogDir         string `yaml:"log_dir,omitempty"`
	Offline        bool   `yaml:"offline,omitempty"`
	GCTTL          int64  `yaml:"gc-ttl,omitempty" validate:"gc-ttl:editable"`
	TZ             string `yaml:"tz,omitempty" validate:"tz:editable"`
	NumaNode       string `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	Arch           string `yaml:"arch,omitempty"`
	OS             string `yaml:"os,omitempty"`
}

// CDCSpec represents the CDC topology specification in topology.yaml
type CDCSpec struct {
	Host           string                 `yaml:"host"`
	SSHPort        int                    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported       bool                   `yaml:"imported,omitempty"`
	Patched        bool                   `yaml:"patched,omitempty"`
	IgnoreExporter bool                   `yaml:"ignore_exporter,omitempty"`
	Port           int                    `yaml:"port,omitempty" default:"8300"`
	DeployDir      string                 `yaml:"deploy_dir,omitempty"`
	DataDir        string                 `yaml:"data_dir,omitempty"`
	LogDir         string                 `yaml:"log_dir,omitempty"`
	Offline        bool                   `yaml:"offline,omitempty"`
	GCTTL          int64                  `yaml:"gc-ttl,omitempty" validate:"gc-ttl:editable"`
	TZ             string                 `yaml:"tz,omitempty" validate:"tz:editable"`
	TiCDCClusterID string                 `yaml:"ticdc_cluster_id"`
	NumaNode       string                 `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	Config         map[string]interface{} `yaml:"config,omitempty" validate:"config:ignore"`
	Arch           string                 `yaml:"arch,omitempty"`
	OS             string                 `yaml:"os,omitempty"`
}

// TiSparkMasterSpec is the topology specification for TiSpark master node
type TiSparkMasterSpec struct {
	Host           string                 `yaml:"host"`
	ListenHost     string                 `yaml:"listen_host,omitempty"`
	SSHPort        int                    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported       bool                   `yaml:"imported,omitempty"`
	Patched        bool                   `yaml:"patched,omitempty"`
	IgnoreExporter bool                   `yaml:"ignore_exporter,omitempty"`
	Port           int                    `yaml:"port,omitempty" default:"7077"`
	WebPort        int                    `yaml:"web_port,omitempty" default:"8080"`
	DeployDir      string                 `yaml:"deploy_dir,omitempty"`
	JavaHome       string                 `yaml:"java_home,omitempty" validate:"java_home:editable"`
	SparkConfigs   map[string]interface{} `yaml:"spark_config,omitempty" validate:"spark_config:ignore"`
	SparkEnvs      map[string]string      `yaml:"spark_env,omitempty" validate:"spark_env:ignore"`
	Arch           string                 `yaml:"arch,omitempty"`
	OS             string                 `yaml:"os,omitempty"`
}

// TiSparkWorkerSpec is the topology specification for TiSpark slave nodes
type TiSparkWorkerSpec struct {
	Host           string `yaml:"host"`
	ListenHost     string `yaml:"listen_host,omitempty"`
	SSHPort        int    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported       bool   `yaml:"imported,omitempty"`
	Patched        bool   `yaml:"patched,omitempty"`
	IgnoreExporter bool   `yaml:"ignore_exporter,omitempty"`
	Port           int    `yaml:"port,omitempty" default:"7078"`
	WebPort        int    `yaml:"web_port,omitempty" default:"8081"`
	DeployDir      string `yaml:"deploy_dir,omitempty"`
	JavaHome       string `yaml:"java_home,omitempty" validate:"java_home:editable"`
	Arch           string `yaml:"arch,omitempty"`
	OS             string `yaml:"os,omitempty"`
}

// PrometheusSpec represents the Prometheus Server topology specification in topology.yaml
type PrometheusSpec struct {
	Host                 string                 `yaml:"host"`
	SSHPort              int                    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported             bool                   `yaml:"imported,omitempty"`
	Patched              bool                   `yaml:"patched,omitempty"`
	IgnoreExporter       bool                   `yaml:"ignore_exporter,omitempty"`
	Port                 int                    `yaml:"port,omitempty" default:"9090"`
	NgPort               int                    `yaml:"ng_port,omitempty" validate:"ng_port:editable"` // ng_port is usable since v5.3.0 and default as 12020 since v5.4.0, so the default value is set in spec.go/AdjustByVersion
	DeployDir            string                 `yaml:"deploy_dir,omitempty"`
	DataDir              string                 `yaml:"data_dir,omitempty"`
	LogDir               string                 `yaml:"log_dir,omitempty"`
	NumaNode             string                 `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	Retention            string                 `yaml:"storage_retention,omitempty" validate:"storage_retention:editable"`
	Arch                 string                 `yaml:"arch,omitempty"`
	OS                   string                 `yaml:"os,omitempty"`
	RuleDir              string                 `yaml:"rule_dir,omitempty" validate:"rule_dir:editable"`
	AdditionalScrapeConf map[string]interface{} `yaml:"additional_scrape_conf,omitempty" validate:"additional_scrape_conf:ignore"`
	ScrapeInterval       string                 `yaml:"scrape_interval,omitempty" validate:"scrape_interval:editable"`
	ScrapeTimeout        string                 `yaml:"scrape_timeout,omitempty" validate:"scrape_timeout:editable"`
}

// GrafanaSpec represents the Grafana topology specification in topology.yaml
type GrafanaSpec struct {
	Host            string            `yaml:"host"`
	SSHPort         int               `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported        bool              `yaml:"imported,omitempty"`
	Patched         bool              `yaml:"patched,omitempty"`
	IgnoreExporter  bool              `yaml:"ignore_exporter,omitempty"`
	Port            int               `yaml:"port,omitempty" default:"3000"`
	DeployDir       string            `yaml:"deploy_dir,omitempty"`
	Config          map[string]string `yaml:"config,omitempty" validate:"config:ignore"`
	Arch            string            `yaml:"arch,omitempty"`
	OS              string            `yaml:"os,omitempty"`
	DashboardDir    string            `yaml:"dashboard_dir,omitempty" validate:"dashboard_dir:editable"`
	Username        string            `yaml:"username,omitempty" default:"admin" validate:"username:editable"`
	Password        string            `yaml:"password,omitempty" default:"admin" validate:"password:editable"`
	AnonymousEnable bool              `yaml:"anonymous_enable,omitempty" default:"false" validate:"anonymous_enable:editable"`
	RootURL         string            `yaml:"root_url,omitempty" validate:"root_url:editable"`
	Domain          string            `yaml:"domain,omitempty" validate:"domain:editable"`
	DefaultTheme    string            `yaml:"default_theme,omitempty" validate:"default_theme:editable"`
	OrgName         string            `yaml:"org_name,omitempty" validate:"org_name:editable"`
	OrgRole         string            `yaml:"org_role,omitempty" validate:"org_role:editable"`
}

// AlertmanagerSpec represents the AlertManager topology specification in topology.yaml
type AlertmanagerSpec struct {
	Host           string `yaml:"host"`
	SSHPort        int    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Imported       bool   `yaml:"imported,omitempty"`
	Patched        bool   `yaml:"patched,omitempty"`
	IgnoreExporter bool   `yaml:"ignore_exporter,omitempty"`
	WebPort        int    `yaml:"web_port,omitempty" default:"9093"`
	ClusterPort    int    `yaml:"cluster_port,omitempty" default:"9094"`
	ListenHost     string `yaml:"listen_host,omitempty" validate:"listen_host:editable"`
	DeployDir      string `yaml:"deploy_dir,omitempty"`
	DataDir        string `yaml:"data_dir,omitempty"`
	LogDir         string `yaml:"log_dir,omitempty"`
	NumaNode       string `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	Arch           string `yaml:"arch,omitempty"`
	OS             string `yaml:"os,omitempty"`
	ConfigFilePath string `yaml:"config_file,omitempty" validate:"config_file:editable"`
}

// DashboardSpec represents the Dashboard topology specification in topology.yaml
type DashboardSpec struct {
	Host           string                 `yaml:"host"`
	SSHPort        int                    `yaml:"ssh_port,omitempty" validate:"ssh_port:editable"`
	Version        string                 `yaml:"version,omitempty"`
	Patched        bool                   `yaml:"patched,omitempty"`
	IgnoreExporter bool                   `yaml:"ignore_exporter,omitempty"`
	Port           int                    `yaml:"port,omitempty" default:"12333"`
	DeployDir      string                 `yaml:"deploy_dir,omitempty"`
	DataDir        string                 `yaml:"data_dir,omitempty"`
	LogDir         string                 `yaml:"log_dir,omitempty"`
	NumaNode       string                 `yaml:"numa_node,omitempty" validate:"numa_node:editable"`
	Config         map[string]interface{} `yaml:"config,omitempty" validate:"config:ignore"`
	Arch           string                 `yaml:"arch,omitempty"`
	OS             string                 `yaml:"os,omitempty"`
}
