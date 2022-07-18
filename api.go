package lwapi

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	APIVer                   = "v2"
	BaseURL                  = "https://api.leaseweb.com"
	ServiceTypeDedicated     = "bareMetals"
	ServiceTypeVirtualServer = "virtualServers"
	ServiceTypeCloud         = "cloud"
)

var NilPayload []byte

//New Api
func New(token string) *Token {
	return &Token{
		Token: token,
	}

}

func (t *Token) VirtualServers() *Api {
	return &Api{
		BaseURL:     BaseURL,
		conn:        connect(),
		Token:       t.Token,
		ServiceType: ServiceTypeVirtualServer,
	}
}

func (t *Token) DedicatedServers() *Api {
	return &Api{
		BaseURL:     BaseURL,
		conn:        connect(),
		Token:       t.Token,
		ServiceType: ServiceTypeDedicated,
	}
}

func connect() *http.Client {
	tl := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{Transport: tl}
}

func (a *Api) entrypoint() string {
	switch a.ServiceType {
	case ServiceTypeDedicated:
		return fmt.Sprintf("%s/%s/%s", a.BaseURL, ServiceTypeDedicated, APIVer)
	case ServiceTypeVirtualServer:
		return fmt.Sprintf("%s/%s/%s/%s", a.BaseURL, ServiceTypeCloud, APIVer, ServiceTypeVirtualServer)
	default:
		return ""
	}
}

func (a *Api) NewRequest(payload []byte, uri string, reqType string) ([]byte, error) {
	body := bytes.NewReader(payload)
	req, err := http.NewRequest(reqType, a.entrypoint()+uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-lsw-auth", a.Token)

	resp, err := a.conn.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)

	return bodyResp, err
}

func MakeQuery(params map[string]interface{}) string {
	query := "?"
	if len(params) > 0 {
		for k, v := range params {
			if fmt.Sprint(v) != "" {
				query += fmt.Sprintf("%s=%v&", k, v)
			}
		}
	} else {
		query = ""
	}
	if len(query) != 0 {
		if query[len(query)-1:] == "&" {
			query = query[:len(query)-1]
		}
	}

	return query
}

func (u *UpdateIPRequest) Validate() error {
	switch u.DetectionProfile {
	case "ADVANCED_DEFAULT":
	case "ADVANCED_LOW_UDP":
	case "ADVANCED_MED_UDP":
	case "":
	default:
		return errors.New("DetectionProfile should be one of ADVANCED_DEFAULT ADVANCED_LOW_UDP ADVANCED_MED_UDP")
	}
	return nil
}

func (n NetworkType) Validate() error {
	switch n {
	case "public":
	case "internal":
	case "remoteManagement":
	default:
		return errors.New("networkType should be one of public internal remoteManagement")
	}
	return nil
}

func (l LinkSpeed) Validate() error {
	switch l.LinkSpeed {
	case 100:
	case 1000:
	case 10000:
	default:
		return errors.New("linkSpeed should be one of 100 1000 10000")
	}
	return nil
}

func (i InstallationJob) Validate() error {
	switch i.OperatingSystemID {
	case "":
		return errors.New("OperatingSystemID can't be empty")
	default:
	}
	return nil
}

func (r RescueModeJob) Validate() error {
	switch r.RescueImageID {
	case "":
		return errors.New("RescueImageID can't be empty")
	default:
		return nil
	}
}

func (c Credential) Validate() error {
	switch c.Type {
	case "OPERATING_SYSTEM":
	case "CONTROL_PANEL":
	case "REMOTE_MANAGEMENT":
	case "RESCUE_MODE":
	case "SWITCH":
	case "PDU":
	case "FIREWALL":
	case "LOAD_BALANCER":
	default:
		return errors.New("credType should be one of OPERATING_SYSTEM CONTROL_PANEL REMOTE_MANAGEMENT RESCUE_MODE SWITCH PDU FIREWALL LOAD_BALANCER")
	}
	switch c.Password {
	case "":
		return errors.New("password can't be empty")
	}
	switch c.Username {
	case "":
		return errors.New("username can't be empty")
	}

	return nil
}

func (c CredType) Validate() error {
	switch c {
	case "OPERATING_SYSTEM":
	case "CONTROL_PANEL":
	case "REMOTE_MANAGEMENT":
	case "RESCUE_MODE":
	case "SWITCH":
	case "PDU":
	case "FIREWALL":
	case "LOAD_BALANCER":
	default:
		return errors.New("credType should be one of OPERATING_SYSTEM CONTROL_PANEL REMOTE_MANAGEMENT RESCUE_MODE SWITCH PDU FIREWALL LOAD_BALANCER")
	}
	return nil
}

func (c VirtualCredType) Validate() error {
	switch c {
	case "OPERATING_SYSTEM":
	case "CONTROL_PANEL":

	default:
		return errors.New("credType should be one of OPERATING_SYSTEM CONTROL_PANEL")
	}
	return nil
}

func (p Password) Validate() error {
	switch p.Password {
	case "":
		return errors.New("password can't be empty")
	default:
		return nil
	}
}

func (m BandwidthMetrics) Validate() (map[string]interface{}, error) {
	q := make(map[string]interface{})
	if m.From == "" || m.To == "" || m.Aggregation == "" {
		return nil, errors.New("from, to, aggregation can't be empty")
	}
	switch m.Aggregation {
	case "AVG":
	case "95TH":
	default:
		return q, errors.New("aggregation should be one of AVG 95TH")
	}

	q["from"] = m.From
	q["to"] = m.To
	q["aggregation"] = m.Aggregation
	if m.Granularity != "" {
		q["granularity"] = m.Granularity
	}
	return q, nil
}

func (m DatatrafficMetrics) Validate() (map[string]interface{}, error) {
	if m.From == "" || m.To == "" || m.Aggregation == "" {
		return nil, errors.New("from, to, aggregation can't be empty")
	}
	switch m.Aggregation {
	case "SUM":
	default:
		return nil, errors.New("aggregation should be one of AVG 95TH")
	}

	q := make(map[string]interface{})
	q["from"] = m.From
	q["to"] = m.To
	q["aggregation"] = m.Aggregation
	if m.Granularity != "" {
		q["granularity"] = m.Granularity
	}
	return q, nil
}

func (n NotificationRequest) Validete() error {
	switch n.Frequency {
	case "DAILY":
	case "WEEKLY":
	case "MONTHLY":
	default:
		return errors.New("frequency should be one of DAILY WEEKLY MONTHLY")
	}
	switch n.Threshold {
	case "":
		return errors.New("threshold can't be empty")
	default:
	}
	switch n.Unit {
	case "Gbps":
	case "Mbps":
	default:
		return errors.New("unit should be one of Gbps Mbps")
	}
	return nil
}

func (n DataTrafficNotificationRequest) Validete() error {
	switch n.Frequency {
	case "DAILY":
	case "WEEKLY":
	case "MONTHLY":
	default:
		return errors.New("frequency should be one of DAILY WEEKLY MONTHLY")
	}
	switch n.Threshold {
	case "":
		return errors.New("threshold can't be empty")
	default:
	}
	switch n.Unit {
	case "GB":
	case "MB":
	case "TB":
	default:
		return errors.New("unit should be one of GB MB TB")
	}
	return nil
}

func FormatISO8601(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05Z07:00")
}

func FormatRFC3339(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05Z07:00")
}

func (v VirtualServerCredentialUpdate) Validate() error {
	switch v.Password {
	case "":
		return errors.New("password can't be empty")
	default:
	}

	switch v.Username {
	case "":
		return errors.New("username can't be empty")
	default:

	}

	switch v.Type {
	case "OPERATING_SYSTEM":
	case "CONTROL_PANEL":
	default:
		return errors.New("type should be one of CONTROL_PANEL OPERATING_SYSTEM")
	}
	return nil
}
