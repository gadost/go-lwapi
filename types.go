package lwapi

import (
	"net/http"
	"time"
)

type Error struct {
	ErrorCode     string `json:"errorCode"`
	ErrorMessage  string `json:"errorMessage"`
	CorrelationID string `json:"correlationId"`
	UserMessage   string `json:"userMessage"`
	Reference     string `json:"reference"`
}

type Api struct {
	BaseURL     string
	conn        *http.Client
	Token       string
	ServiceType string
}

type Token struct {
	Token string
}

type Metadata struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"totalCount"`
}

type Contract struct {
	CustomerID     string `json:"customerId"`
	DeliveryStatus string `json:"deliveryStatus"`
	ID             string `json:"id"`
	Reference      string `json:"reference"`
	SalesOrgID     string `json:"salesOrgId"`
}

type FeatureAvailability struct {
	Automation       bool `json:"automation"`
	IpmiReboot       bool `json:"ipmiReboot"`
	PowerCycle       bool `json:"powerCycle"`
	PrivateNetwork   bool `json:"privateNetwork"`
	RemoteManagement bool `json:"remoteManagement"`
}

type Location struct {
	Rack  string `json:"rack"`
	Site  string `json:"site"`
	Suite string `json:"suite"`
	Unit  string `json:"unit"`
}

type Port struct {
	Name string `json:"name"`
	Port string `json:"port"`
}

type InterfaceParams struct {
	Gateway string `json:"gateway"`
	IP      string `json:"ip"`
	Mac     string `json:"mac"`
	Ports   []Port `json:"ports"`
}

type NetworkInterfaces struct {
	Internal         InterfaceParams `json:"internal"`
	Public           InterfaceParams `json:"public"`
	RemoteManagement InterfaceParams `json:"remoteManagement"`
}

type PowerPort struct {
	Name string `json:"name"`
	Port string `json:"port"`
}

type PrivateNetwork struct {
	ID        string `json:"id"`
	LinkSpeed int    `json:"linkSpeed"`
	Status    string `json:"status"`
	Subnet    string `json:"subnet"`
	VlanID    string `json:"vlanId"`
}

type Rack struct {
	Type string `json:"type"`
}

type Servers struct {
	Metadata Metadata `json:"_metadata"`
	Servers  []Server `json:"servers"`
	Error
}

type Server struct {
	AssetID             string              `json:"assetId"`
	Contract            Contract            `json:"contract"`
	FeatureAvailability FeatureAvailability `json:"featureAvailability,omitempty"`
	ID                  string              `json:"id"`
	Location            Location            `json:"location"`
	NetworkInterfaces   NetworkInterfaces   `json:"networkInterfaces"`
	PowerPorts          []PowerPort         `json:"powerPorts,omitempty"`
	PrivateNetworks     []PrivateNetwork    `json:"privateNetworks"`
	Rack                Rack                `json:"rack"`
	Error
}

type Reference struct {
	Reference string `json:"reference"`
}

type Firmware struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	Vendor      string `json:"vendor"`
	Version     string `json:"version"`
}

type Motherboard struct {
	Product string `json:"product"`
	Serial  string `json:"serial"`
	Vendor  string `json:"vendor"`
}

type HardwareChassis struct {
	Description string      `json:"description"`
	Firmware    Firmware    `json:"firmware"`
	Motherboard Motherboard `json:"motherboard"`
	Product     string      `json:"product"`
	Serial      string      `json:"serial"`
	Vendor      string      `json:"vendor"`
}

type CPUCapabilities struct {
	Cpufreq string `json:"cpufreq"`
	Ht      string `json:"ht"`
	Vmx     bool   `json:"vmx"`
	X8664   string `json:"x86-64"`
}

type CPUSettings struct {
	Cores        string `json:"cores"`
	Enabledcores string `json:"enabledcores"`
	Threads      string `json:"threads"`
}

type CPU struct {
	Capabilities CPUCapabilities `json:"capabilities"`
	Description  string          `json:"description"`
	Hz           string          `json:"hz"`
	SerialNumber string          `json:"serial_number"`
	Settings     CPUSettings     `json:"settings"`
	Slot         string          `json:"slot"`
	Vendor       string          `json:"vendor"`
}

type HardwareInformation struct {
	ID            string                    `json:"id"`
	ParserVersion string                    `json:"parserVersion"`
	Result        HardwareInformationResult `json:"result"`
	ScannedAt     time.Time                 `json:"scannedAt"`
	ServerID      string                    `json:"serverId"`

	Error
}

type HardwareInformationResult struct {
	Chassis HardwareChassis `json:"chassis"`
	CPU     []CPU           `json:"cpu"`
	Disks   []Disk          `json:"disks"`
	Ipmi    IPMI            `json:"ipmi"`
	Memory  []Memory        `json:"memory"`
	Network []Network       `json:"network"`
}

type Disk struct {
	Description  string   `json:"description"`
	ID           string   `json:"id"`
	Product      string   `json:"product"`
	SerialNumber string   `json:"serial_number"`
	Size         string   `json:"size"`
	Smartctl     Smartctl `json:"smartctl"`
	Vendor       string   `json:"vendor"`
}

type Smartctl struct {
	AtaVersion      string             `json:"ata_version"`
	Attributes      SmartctlAttributes `json:"attributes"`
	DeviceModel     string             `json:"device_model"`
	ExecutionStatus string             `json:"execution_status"`
	FirmwareVersion string             `json:"firmware_version"`
	IsSas           bool               `json:"is_sas"`
	OverallHealth   string             `json:"overall_health"`
	Rpm             string             `json:"rpm"`
	SataVersion     string             `json:"sata_version"`
	SectorSize      string             `json:"sector_size"`
	SerialNumber    string             `json:"serial_number"`
	SmartErrorLog   string             `json:"smart_error_log"`
	SmartSupport    SmartSupport       `json:"smart_support"`
	SmartctlVersion string             `json:"smartctl_version"`
	UserCapacity    string             `json:"user_capacity"`
}

type SmartctlAttributes struct {
	PowerOnHours        PowerOnHours        `json:"Power_On_Hours"`
	ReallocatedSectorCt ReallocatedSectorCt `json:"Reallocated_Sector_Ct"`
}

type PowerOnHours struct {
	Flag       string `json:"flag"`
	ID         string `json:"id"`
	RawValue   string `json:"raw_value"`
	Thresh     string `json:"thresh"`
	Type       string `json:"type"`
	Updated    string `json:"updated"`
	Value      string `json:"value"`
	WhenFailed string `json:"when_failed"`
	Worst      string `json:"worst"`
}

type ReallocatedSectorCt struct {
	Flag       string `json:"flag"`
	ID         string `json:"id"`
	RawValue   string `json:"raw_value"`
	Thresh     string `json:"thresh"`
	Type       string `json:"type"`
	Updated    string `json:"updated"`
	Value      string `json:"value"`
	WhenFailed string `json:"when_failed"`
	Worst      string `json:"worst"`
}

type SmartSupport struct {
	Available bool `json:"available"`
	Enabled   bool `json:"enabled"`
}

type IPMI struct {
	Defgateway string `json:"defgateway"`
	Firmware   string `json:"firmware"`
	Ipaddress  string `json:"ipaddress"`
	Ipsource   string `json:"ipsource"`
	Macaddress string `json:"macaddress"`
	Subnetmask string `json:"subnetmask"`
	Vendor     string `json:"vendor"`
}

type Memory struct {
	ClockHz      string `json:"clock_hz"`
	Description  string `json:"description"`
	ID           string `json:"id"`
	SerialNumber string `json:"serial_number"`
	SizeBytes    string `json:"size_bytes"`
}

type Network struct {
	Capabilities NetworkCapabilities `json:"capabilities"`
	Lldp         Lldp                `json:"lldp"`
	LogicalName  string              `json:"logical_name"`
	MacAddress   string              `json:"mac_address"`
	Product      string              `json:"product"`
	Settings     Settings            `json:"settings"`
	Vendor       string              `json:"vendor"`
}

type NetworkCapabilities struct {
	Autonegotiation string      `json:"autonegotiation"`
	BusMaster       string      `json:"bus_master"`
	CapList         string      `json:"cap_list"`
	Ethernet        string      `json:"ethernet"`
	LinkSpeeds      interface{} `json:"link_speeds"`
	Msi             string      `json:"msi"`
	Msix            string      `json:"msix"`
	Pciexpress      string      `json:"pciexpress"`
	Physical        string      `json:"physical"`
	Pm              string      `json:"pm"`
	Tp              string      `json:"tp"`
}

type Lldp struct {
	Chassis LldpChassis `json:"chassis"`
	Port    LldpPort    `json:"port"`
	Vlan    LldpVlan    `json:"vlan"`
}

type LldpPort struct {
	AutoNegotiation AutoNegotiation `json:"auto_negotiation"`
	Description     string          `json:"description"`
}
type LldpChassis struct {
	Description string `json:"description"`
	MacAddress  string `json:"mac_address"`
	Name        string `json:"name"`
}

type AutoNegotiation struct {
	Enabled   string `json:"enabled"`
	Supported string `json:"supported"`
}

type LldpVlan struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Name  string `json:"name"`
}

type Settings struct {
	Autonegotiation string `json:"autonegotiation"`
	Broadcast       string `json:"broadcast"`
	Driver          string `json:"driver"`
	Driverversion   string `json:"driverversion"`
	Duplex          string `json:"duplex"`
	Firmware        string `json:"firmware"`
	IP              string `json:"ip"`
	Latency         string `json:"latency"`
	Link            string `json:"link"`
	Multicast       string `json:"multicast"`
	Port            string `json:"port"`
	Speed           string `json:"speed"`
}

type DDoS struct {
	DetectionProfile string `json:"detectionProfile"`
	ProtectionType   string `json:"protectionType"`
}

type ServerIP struct {
	Ddos          DDoS   `json:"ddos"`
	FloatingIP    bool   `json:"floatingIp"`
	Gateway       string `json:"gateway"`
	IP            string `json:"ip"`
	MainIP        bool   `json:"mainIp"`
	NetworkType   string `json:"networkType"`
	NullRouted    bool   `json:"nullRouted"`
	ReverseLookup string `json:"reverseLookup"`
	Version       int    `json:"version"`
	Error
}

type ServerIPs struct {
	Metadata Metadata   `json:"_metadata"`
	Ips      []ServerIP `json:"ips"`
	Error
}

type UpdateIPRequest struct {
	// DetectionProfile one of "ADVANCED_DEFAULT" "ADVANCED_LOW_UDP" "ADVANCED_MED_UDP"
	DetectionProfile string `json:"detectionProfile"`
	ReverseLookup    string `json:"reverseLookup"`
}

type NullHistory struct {
	Metadata   Metadata    `json:"_metadata"`
	NullRoutes []NullRoute `json:"nullRoutes"`

	Error
}

type NullRoute struct {
	AutomatedUnnullingAt time.Time `json:"automatedUnnullingAt"`
	Comment              string    `json:"comment"`
	IP                   string    `json:"ip"`
	NullLevel            int       `json:"nullLevel"`
	NulledAt             time.Time `json:"nulledAt"`
}

type NetworkInterfacesList struct {
	Metadata          Metadata           `json:"_metadata"`
	NetworkInterfaces []NetworkInterface `json:"networkInterfaces"`

	Error
}
type NetworkInterface struct {
	LinkSpeed       string `json:"linkSpeed"`
	OperStatus      string `json:"operStatus"`
	Status          string `json:"status"`
	SwitchInterface string `json:"switchInterface"`
	SwitchName      string `json:"switchName"`
	Type            string `json:"type"`
}

type LinkSpeed struct {
	LinkSpeed int `json:"linkSpeed"`
}

type ServerDHCPLeases struct {
	Metadata Metadata          `json:"_metadata"`
	Leases   []ServerDHCPLease `json:"leases"`
	Error
}

type ServerDHCPLease struct {
	Bootfile          string            `json:"bootfile"`
	CreatedAt         time.Time         `json:"createdAt"`
	Gateway           string            `json:"gateway"`
	Hostname          string            `json:"hostname"`
	IP                string            `json:"ip"`
	LastClientRequest LastClientRequest `json:"lastClientRequest"`
	Mac               string            `json:"mac"`
	Netmask           string            `json:"netmask"`
	Site              string            `json:"site"`
	UpdatedAt         time.Time         `json:"updatedAt"`
}

type LastClientRequest struct {
	RelayAgent interface{} `json:"relayAgent"`
	Type       string      `json:"type"`
	UserAgent  string      `json:"userAgent"`
}

type ServerDHCPLeaseNew struct {
	Bootfile string `json:"bootfile"`
	Hostname string `json:"hostname,omitempty"`
}

type Job struct {
	CreatedAt string      `json:"createdAt"`
	Flow      string      `json:"flow"`
	IsRunning bool        `json:"isRunning"`
	Node      string      `json:"node"`
	Payload   JobPayload  `json:"payload"`
	Progress  JobProgress `json:"progress"`
	ServerID  string      `json:"serverId"`
	Status    string      `json:"status"`
	Tasks     []Task      `json:"tasks"`
	Type      string      `json:"type"`
	UpdatedAt string      `json:"updatedAt"`
	UUID      string      `json:"uuid"`
}

type JobPayload struct {
	Configurable      bool                `json:"configurable"`
	Device            string              `json:"device"`
	FileserverBaseURL string              `json:"fileserverBaseUrl"`
	JobType           string              `json:"jobType"`
	NumberOfDisks     interface{}         `json:"numberOfDisks"`
	OperatingSystemID string              `json:"operatingSystemId"`
	Os                PayloadOS           `json:"os"`
	Partitions        []PayloadPartitions `json:"partitions"`
	Pop               string              `json:"pop"`
	PowerCycle        bool                `json:"powerCycle"`
	RaidLevel         interface{}         `json:"raidLevel"`
	ServerID          string              `json:"serverId"`
	Timezone          string              `json:"timezone"`
	X                 int                 `json:"x"`
}

type PayloadPartitions struct {
	Filesystem string `json:"filesystem"`
	Size       int    `json:"size"`
}

type StatusTimestamps struct {
	Canceled time.Time `json:"CANCELED"`
	Pending  time.Time `json:"PENDING"`
	Waiting  time.Time `json:"WAITING"`
}

type JobProgress struct {
	Canceled   int `json:"canceled"`
	Expired    int `json:"expired"`
	Failed     int `json:"failed"`
	Finished   int `json:"finished"`
	Inprogress int `json:"inprogress"`
	Pending    int `json:"pending"`
	Percentage int `json:"percentage"`
	Total      int `json:"total"`
	Waiting    int `json:"waiting"`
}

type PayloadOS struct {
	Architecture string `json:"architecture"`
	Family       string `json:"family"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Version      string `json:"version"`
}

type Task struct {
	Description      string           `json:"description"`
	ErrorMessage     string           `json:"errorMessage"`
	Flow             string           `json:"flow"`
	OnError          string           `json:"onError"`
	Status           string           `json:"status"`
	StatusTimestamps StatusTimestamps `json:"statusTimestamps"`
	UUID             string           `json:"uuid"`
}

type HardwareScanJob struct {
	PowerCycle  bool   `json:"powerCycle"`
	CallbackUrl string `json:"callbackUrl"`
}

type InstallationJob struct {
	ControlPanelID    string      `json:"controlPanelId,omitempty"`
	Device            string      `json:"device,omitempty"`
	Hostname          string      `json:"hostname,omitempty"`
	OperatingSystemID string      `json:"operatingSystemId"`
	Partitions        []Partition `json:"partitions,omitempty"`
	SSHKeys           string      `json:"sshKeys,omitempty"`
}

type Partition struct {
	Bootable   bool   `json:"bootable,omitempty"`
	Filesystem string `json:"filesystem"`
	Mountpoint string `json:"mountpoint,omitempty"`
	Primary    bool   `json:"primary,omitempty"`
	Size       string `json:"size"`
}

type CallbackURL struct {
	CallbackURL string `json:"callbackUrl"`
}

type Jobs struct {
	Metadata Metadata `json:"_metadata"`
	Jobs     []Job    `json:"jobs"`

	Error
}

type RescueModeJob struct {
	CallbackURL   string `json:"callbackUrl,omitempty"`
	PowerCycle    bool   `json:"powerCycle,omitempty"`
	RescueImageID string `json:"rescueImageId"`
	SSHKeys       string `json:"sshKeys,omitempty"`

	Error
}

type Credentials struct {
	Metadata    Metadata     `json:"_metadata"`
	Credentials []Credential `json:"credentials"`

	Error
}

type Credential struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`

	Error
}

type Password struct {
	Password string `json:"password"`
}

type MetricsMetadata struct {
	Aggregation string    `json:"aggregation"`
	From        time.Time `json:"from"`
	Granularity string    `json:"granularity"`
	To          time.Time `json:"to"`
}

type Metrics struct {
	Metadata MetricsMetadata `json:"_metadata"`
	Metrics  Metric          `json:"metrics"`

	Error
}

type Metric struct {
	DownPublic MetricsValues `json:"DOWN_PUBLIC"`
	UpPublic   MetricsValues `json:"UP_PUBLIC"`
}
type MetricsValues struct {
	Unit   string         `json:"unit"`
	Values []MetricsValue `json:"values"`
}

type MetricsValue struct {
	Timestamp time.Time `json:"timestamp"`
	Value     int       `json:"value"`
}

type BandwidthNotification struct {
	Metadata                      Metadata              `json:"_metadata"`
	BandwidthNotificationSettings []NotificationSetting `json:"bandwidthNotificationSettings"`
}

type DatatrafficNotification struct {
	Metadata                        Metadata              `json:"_metadata"`
	DatatrafficNotificationSettings []NotificationSetting `json:"datatrafficNotificationSettings"`
}

type Action struct {
	LastTriggeredAt time.Time `json:"lastTriggeredAt"`
	Type            string    `json:"type"`
}

type NotificationSetting struct {
	Actions             []Action  `json:"actions"`
	Frequency           string    `json:"frequency"`
	ID                  string    `json:"id"`
	LastCheckedAt       time.Time `json:"lastCheckedAt"`
	Threshold           string    `json:"threshold"`
	ThresholdExceededAt time.Time `json:"thresholdExceededAt"`
	Unit                string    `json:"unit"`
}

type NotificationRequest struct {
	Frequency string `json:"frequency"`
	Threshold string `json:"threshold"`
	Unit      string `json:"unit"`
}

type DataTrafficNotificationRequest struct {
	Frequency string `json:"frequency"`
	Threshold string `json:"threshold"`
	Unit      string `json:"unit"`
}

type NotificationResponse struct {
	Actions             []Action  `json:"actions"`
	Frequency           string    `json:"frequency"`
	ID                  string    `json:"id"`
	LastCheckedAt       time.Time `json:"lastCheckedAt"`
	Threshold           string    `json:"threshold"`
	ThresholdExceededAt time.Time `json:"thresholdExceededAt"`
	Unit                string    `json:"unit"`

	Error
}

type DDoSStatus struct {
	Nulling   string `json:"nulling"`
	Scrubbing string `json:"scrubbing"`

	Error
}

type PowerStatus struct {
	Ipmi IPMIPowerStatus `json:"ipmi"`
	Pdu  PduPowerStatus  `json:"pdu"`

	Error
}

type IPMIPowerStatus struct {
	Status string `json:"status"`
}

type PduPowerStatus struct {
	Status string `json:"status"`
}

type OperatingSystems struct {
	Metadata         Metadata          `json:"_metadata"`
	OperatingSystems []OperatingSystem `json:"operatingSystems"`

	Error
}

type OperatingSystem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OSParams struct {
	Architecture         string           `json:"architecture"`
	Configurable         bool             `json:"configurable"`
	Defaults             OSParamsDefaults `json:"defaults"`
	Family               string           `json:"family"`
	Features             []string         `json:"features"`
	ID                   string           `json:"id"`
	Name                 string           `json:"name"`
	SupportedBootDevices []string         `json:"supportedBootDevices"`
	SupportedFileSystems []string         `json:"supportedFileSystems"`
	Type                 string           `json:"type"`
	Version              string           `json:"version"`

	Error
}

type OSParamsDefaults struct {
	Device     string      `json:"device"`
	Partitions []Partition `json:"partitions"`
}

type ControlPanels struct {
	Metadata      Metadata       `json:"_metadata"`
	ControlPanels []ControlPanel `json:"controlPanels"`

	Error
}

type ControlPanel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RescueImages struct {
	Metadata     Metadata      `json:"_metadata"`
	RescueImages []RescueImage `json:"rescueImages"`

	Error
}

type RescueImage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NetworkType string
type CredType string
type ControlPanelID int

type BandwidthMetrics struct {
	From        string
	To          string
	Aggregation string
}

type DatatrafficMetrics struct {
	From        string
	To          string
	Aggregation string
}
