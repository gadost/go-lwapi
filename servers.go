package lwapi

import (
	"encoding/json"
	"fmt"
)

var NilPayload = []byte("")

func (a *Api) Servers(queryParams ...map[string]interface{}) (*Servers, error) {
	query := MakeQuery(queryParams)
	bodyResp, err := a.NewRequest(NilPayload, "/servers"+query, "GET")

	var r *Servers
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) Server(serverID uint64) (*Server, error) {
	uri := fmt.Sprintf("/servers/%d", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Server
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerReferenceUpdate(serverID uint64, params *Reference) (*Server, error) {
	uri := fmt.Sprintf("/servers/%d", serverID)
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *Server
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerHardwareInformation(serverID uint64) (*HardwareInformation, error) {
	uri := fmt.Sprintf("/servers/%d/hardwareInfo", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *HardwareInformation
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerIPList(
	serverID uint64, queryParams ...map[string]interface{}) (*ServerIPs, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/ips%s", serverID, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *ServerIPs
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerIP(serverID uint64, serverIP string) (*ServerIP, error) {
	uri := fmt.Sprintf("/servers/%d/ips/%s", serverID, serverIP)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *ServerIP
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerIPUpdate(serverID uint64, serverIP string, params *UpdateIPRequest) (
	*ServerIP, error) {
	uri := fmt.Sprintf("/servers/%d/ips/%s", serverID, serverIP)
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *ServerIP
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerIPNull(serverID uint64, serverIP string) (*ServerIP, error) {
	uri := fmt.Sprintf("/servers/%d/ips/%s/null", serverID, serverIP)
	bodyResp, err := a.NewRequest(NilPayload, uri, "PUT")

	var r *ServerIP
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerIPUnNull(serverID uint64, serverIP string) (*ServerIP, error) {
	uri := fmt.Sprintf("/servers/%d/ips/%s/unnull", serverID, serverIP)
	bodyResp, err := a.NewRequest(NilPayload, uri, "PUT")

	var r *ServerIP
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerIPNullHistory(
	serverID uint64, serverIP string, queryParams ...map[string]interface{}) (
	*NullHistory, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/ips/%s/nullRouteHistory%s", serverID, serverIP, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *NullHistory
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// /bareMetals/v2/servers/{serverId}/networkInterfaces
func (a *Api) ServerNetworkInterfaces(serverID uint64) (*NetworkInterfacesList, error) {
	uri := fmt.Sprintf("/servers/%d/networkInterfaces", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *NetworkInterfacesList
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerNetworkInterfacesClose(serverID uint64) (*NetworkInterfacesList, error) {
	uri := fmt.Sprintf("/servers/%d/networkInterfaces/close", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *NetworkInterfacesList
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerNetworkInterfacesOpen(serverID uint64) (*NetworkInterfacesList, error) {
	uri := fmt.Sprintf("/servers/%d/networkInterfaces/open", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *NetworkInterfacesList
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerNetworkInterface(serverID uint64, networkType string) (*NetworkInterface, error) {
	uri := fmt.Sprintf("/servers/%d/networkInterfaces/%s", serverID, networkType)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *NetworkInterface
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerNetworkInterfaceClose(serverID uint64, networkType string) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/networkInterfaces/%s/close", serverID, networkType)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerNetworkInterfaceOpen(serverID uint64, networkType string) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/networkInterfaces/%s/open", serverID, networkType)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/privateNetworks/{privateNetworkId}
func (a *Api) ServerPrivateNetworkDelete(serverID uint64, privateNetworkID int) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/privateNetworks/%d", serverID, privateNetworkID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "DELETE")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerPrivateNetworkAdd(
	serverID uint64, privateNetworkID int, linkSpeed int) (*Error, error) {
	payload, _ := json.Marshal(&LinkSpeed{
		LinkSpeed: linkSpeed,
	})
	uri := fmt.Sprintf("/servers/%d/privateNetworks/%d", serverID, privateNetworkID)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/leases
func (a *Api) ServerDHCPLeaseDelete(serverID uint64) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/leases", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "DELETE")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerDHCPLeases(serverID uint64) (*ServerDHCPLeases, error) {
	uri := fmt.Sprintf("/servers/%d/leases", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *ServerDHCPLeases
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerDHCPLeaseNew(serverID uint64, params *ServerDHCPLeaseNew) (*Error, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/leases", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/cancelActiveJob

func (a *Api) ServerActiveJobCancel(serverID uint64) (*Job, error) {

	uri := fmt.Sprintf("/servers/%d/cancelActiveJob", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerActiveJobExpire(serverID uint64) (*Job, error) {

	uri := fmt.Sprintf("/servers/%d/expireActiveJob", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/hardwareScan
func (a *Api) ServerHardwareScan(serverID uint64, powerCycle bool, callbackUrl string) (*Job, error) {
	payload, _ := json.Marshal(&HardwareScanJob{
		PowerCycle:  powerCycle,
		CallbackUrl: callbackUrl,
	})
	uri := fmt.Sprintf("/servers/%d/hardwareScan", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/install
func (a *Api) ServerInstalationLaunch(serverID uint64, params *InstallationJob) (*Job, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/install", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/ipmiReset
func (a *Api) ServerIPMIReset(serverID uint64, params *CallbackURL) (*Job, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/ipmiReset", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerJobs(serverID uint64) (*Jobs, error) {
	uri := fmt.Sprintf("/servers/%d/jobs", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Jobs
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerJob(serverID uint64, jobID int) (*Job, error) {
	uri := fmt.Sprintf("/servers/%d/jobs/%d", serverID, jobID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/rescueMode
func (a *Api) ServerRescueMode(serverID uint64, params *RescueModeJob) (*Job, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/rescueMode", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/credentials
func (a *Api) ServerCredentials(
	serverID uint64, queryParams ...map[string]interface{}) (*Credentials, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/credentials%s", serverID, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Credentials
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/credentials
func (a *Api) ServerCredentialNew(serverID uint64, params *Credential) (*Credential, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/credentials", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Credential
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/credentials/{type}
func (a *Api) ServerTypedCredentials(
	serverID uint64, credType string, queryParams ...map[string]interface{}) (*Credentials, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/credentials/%s%s", serverID, credType, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Credentials
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/credentials/{type}/{username}
func (a *Api) ServerCredentialDelete(serverID uint64, credType string) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/credentials/%s", serverID, credType)
	bodyResp, err := a.NewRequest(NilPayload, uri, "DELETE")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ServerUserTypedCredentials(
	serverID uint64, credType string, username string) (*Credential, error) {
	uri := fmt.Sprintf("/servers/%d/credentials/%s/%s", serverID, credType, username)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Credential
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// https://api.leaseweb.com/bareMetals/v2/servers/{serverId}/credentials/{type}/{username}
func (a *Api) ServerCredentialUpdate(
	serverID uint64, credType string, username string, params *Password) (*Credential, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/credentials/%s/%s", serverID, credType, username)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Credential
	json.Unmarshal(bodyResp, &r)
	return r, err
}
