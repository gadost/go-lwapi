package lwapi

import (
	"encoding/json"
	"fmt"
)

func (a *Api) VirtualServers(queryParams map[string]interface{}) (*VirtualServers, error) {
	query := MakeQuery(queryParams)
	bodyResp, err := a.NewRequest(NilPayload, query, "GET")

	var r *VirtualServers
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) VirtualServer(vserverID uint64) (*VirtualServer, error) {
	uri := fmt.Sprintf("/%d", vserverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *VirtualServer
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) VirtualServerReferenceUpdate(vserverID uint64, params *Reference) (*VirtualServer, error) {
	uri := fmt.Sprintf("/%d", vserverID)
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *VirtualServer
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) VirtualServerPowerOff(vserverID uint64) (*Error, error) {
	uri := fmt.Sprintf("/%d/powerOff", vserverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) VirtualServerPowerOn(vserverID uint64) (*Error, error) {
	uri := fmt.Sprintf("/%d/powerOn", vserverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) VirtualServerReboot(vserverID uint64) (*Error, error) {
	uri := fmt.Sprintf("/%d/reboot", vserverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) VirtualServerReinstall(vserverID uint64, operatingSystemId string) (*Error, error) {
	uri := fmt.Sprintf("/%d/reinstall", vserverID)
	payload, _ := json.Marshal(&OperatingSystemID{
		OperatingSystemID: operatingSystemId,
	})
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// This API call will only update the password displayed for your machine in the customer portal.
// It will not make any changes to the root password of your machine.
func (a *Api) VirtualServerCredentialUpdate(
	vserverID uint64, params *VirtualServerCredentialUpdate) (*Error, error) {

	if err := params.Validate(); err != nil {
		return nil, err
	}

	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/%d/credentials", vserverID)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// You will only be able to retrieve the last password that we configured in your server or control panel.
// If you changed it, the information retrieved by this API call will not work.
// The password is not returned in this call, you must use the endpoint to get a credential by
// the username to retrieve it.
func (a *Api) VirtualServerCredentials(
	serverID uint64, credType VirtualCredType, queryParams map[string]interface{}) (*Credentials, error) {
	if err := credType.Validate(); err != nil {
		return nil, err
	}
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/%d/credentials/%s%s", serverID, credType, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Credentials
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) VirtualServerUserCredentials(
	vserverID uint64, credType VirtualCredType, username string) (*Credential, error) {
	if err := credType.Validate(); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/%d/credentials/%s/%s", vserverID, credType, username)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Credential
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) VirtualServerDatatraficMetrics(
	vserverID uint64, params *DatatrafficMetrics) (*VirtualServerDatatraficMetrics, error) {
	var queryParams map[string]interface{}
	if ok, err := params.Validate(); err != nil {
		return nil, err
	} else {
		queryParams = ok
	}

	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/%d/metrics/datatraffic%s", vserverID, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *VirtualServerDatatraficMetrics
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) VirtualServerTemplates(vserverID uint64) (*Templates, error) {
	uri := fmt.Sprintf("/%d/templates", vserverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Templates
	json.Unmarshal(bodyResp, &r)
	return r, err
}
