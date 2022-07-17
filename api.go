package lwapi

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

const (
	APIVer                   = "v2"
	BaseURL                  = "https://api.leaseweb.com"
	ServiceTypeDedicated     = "bareMetals"
	ServiceTypeVirtualServer = "virtualServers"
)

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
		return fmt.Sprintf("%s/%s", a.BaseURL, ServiceTypeVirtualServer)
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

func MakeQuery(params []map[string]interface{}) string {
	query := "?"
	if len(params) > 0 {
		for _, p := range params {
			for k, v := range p {
				if fmt.Sprint(v) != "" {
					query += fmt.Sprintf("%s=%s", k, v)
				}
			}
		}
	} else {
		query = ""
	}

	return query
}
