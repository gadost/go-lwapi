package lwapi

import (
	"encoding/json"
	"fmt"
)

// List your Dedicated Servers. This api call supports pagination.
// Use the limit and offset query string parameters to paginate through all your dedicated servers.
// Every server object in the json response lists a few properties of a server.
// Use the single resouce api call to get more details for a single server.
func (a *Api) Servers(queryParams map[string]interface{}) (*Servers, error) {
	query := MakeQuery(queryParams)
	bodyResp, err := a.NewRequest(NilPayload, "/servers"+query, "GET")

	var r *Servers
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Use this API to get information about a single server.
func (a *Api) Server(serverID uint64) (*Server, error) {
	uri := fmt.Sprintf("/servers/%d", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Server
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Update the reference for a server.
func (a *Api) ServerReferenceUpdate(serverID uint64, params *Reference) (*Server, error) {
	uri := fmt.Sprintf("/servers/%d", serverID)
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *Server
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// This information is generated when running a hardware scan for your server.
// A hardware scan collects hardware information about your system.
func (a *Api) ServerHardwareInformation(serverID uint64) (*HardwareInformation, error) {
	uri := fmt.Sprintf("/servers/%d/hardwareInfo", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *HardwareInformation
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// List all IP Addresses associated with this server. Optionally filtered.
func (a *Api) ServerIPList(
	serverID uint64, queryParams map[string]interface{}) (*ServerIPs, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/ips%s", serverID, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *ServerIPs
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Get a single IP address associated with this server.
func (a *Api) ServerIP(serverID uint64, serverIP string) (*ServerIP, error) {
	uri := fmt.Sprintf("/servers/%d/ips/%s", serverID, serverIP)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *ServerIP
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Update the reverse lookup or DDoS detection profile for the ip address.
// DDoS detection profiles can only be changed if the IP address is protected using
// Advanced DDoS protection.
func (a *Api) ServerIPUpdate(serverID uint64, serverIP string, params *UpdateIPRequest) (
	*ServerIP, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/servers/%d/ips/%s", serverID, serverIP)
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *ServerIP
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Null the given IP address.
// It might take a few minutes before the change is propagated across the network.
func (a *Api) ServerIPNull(serverID uint64, serverIP string) (*ServerIP, error) {
	uri := fmt.Sprintf("/servers/%d/ips/%s/null", serverID, serverIP)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *ServerIP
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Remove an existing null route for the given IP address.
// It might take a few minutes before the change is propagated across the network.
func (a *Api) ServerIPUnNull(serverID uint64, serverIP string) (*ServerIP, error) {
	uri := fmt.Sprintf("/servers/%d/ips/%s/unnull", serverID, serverIP)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *ServerIP
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Show all null route history for any ips associated with this server.
func (a *Api) ServerIPNullHistory(
	serverID uint64, serverIP string, queryParams map[string]interface{}) (
	*NullHistory, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/ips/%s/nullRouteHistory%s", serverID, serverIP, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *NullHistory
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// List all network interfaces for this server, including their current status.
func (a *Api) ServerNetworkInterfaces(serverID uint64) (*NetworkInterfacesList, error) {
	uri := fmt.Sprintf("/servers/%d/networkInterfaces", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *NetworkInterfacesList
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Close all network interfaces for this server.
func (a *Api) ServerNetworkInterfacesClose(serverID uint64) (*NetworkInterfacesList, error) {
	uri := fmt.Sprintf("/servers/%d/networkInterfaces/close", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *NetworkInterfacesList
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Open all network interfaces of this server.
func (a *Api) ServerNetworkInterfacesOpen(serverID uint64) (*NetworkInterfacesList, error) {
	uri := fmt.Sprintf("/servers/%d/networkInterfaces/open", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *NetworkInterfacesList
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// List the network interfaces of the given type of this server, including their status.
func (a *Api) ServerNetworkInterface(serverID uint64, networkType NetworkType) (*NetworkInterface, error) {
	if err := networkType.Validate(); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/servers/%d/networkInterfaces/%s", serverID, networkType)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *NetworkInterface
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Close all network interfaces of this server.
func (a *Api) ServerNetworkInterfaceClose(serverID uint64, networkType NetworkType) (*Error, error) {
	if err := networkType.Validate(); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/servers/%d/networkInterfaces/%s/close", serverID, networkType)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Open all network interfaces of the given type for this server.
func (a *Api) ServerNetworkInterfaceOpen(serverID uint64, networkType NetworkType) (*Error, error) {
	if err := networkType.Validate(); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/servers/%d/networkInterfaces/%s/open", serverID, networkType)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// This API call will remove the dedicated server from the private network.
// It takes a few minutes before the server has been removed from the private network.
//To get the current status of the server you can call /bareMetals/v2/servers/{serverId}.
// While the server is being removed the status changes to REMOVING.
func (a *Api) ServerPrivateNetworkDelete(serverID uint64, privateNetworkID int) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/privateNetworks/%d", serverID, privateNetworkID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "DELETE")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// It takes a few minutes before the server has access to the private network.
// To get the current status of the server you can call api.Server(ID).
// Once the server is added to the private network the status changes from CONFIGURING to CONFIGURED.
func (a *Api) ServerPrivateNetworkAdd(serverID uint64, privateNetworkID int, linkSpeed int) (*Error, error) {
	ls := &LinkSpeed{
		LinkSpeed: linkSpeed,
	}
	if err := ls.Validate(); err != nil {
		return nil, err
	}

	payload, _ := json.Marshal(ls)

	uri := fmt.Sprintf("/servers/%d/privateNetworks/%d", serverID, privateNetworkID)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Delete a DHCP reservation for this server.
func (a *Api) ServerDHCPLeaseDelete(serverID uint64) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/leases", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "DELETE")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Please note that this will only show reservations for the public network interface.
func (a *Api) ServerDHCPLeases(serverID uint64) (*ServerDHCPLeases, error) {
	uri := fmt.Sprintf("/servers/%d/leases", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *ServerDHCPLeases
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// After rebooting your server it will acquire this DHCP reservation and boot from the specified bootfile url.
// Please note that this API call will not reboot or power cycle your server.
func (a *Api) ServerDHCPLeaseNew(serverID uint64, params *ServerDHCPLeaseNew) (*Error, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/leases", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Canceling an active job will trigger the onfail flow
// of the current job often resulting in a server reboot.
// If you do not want the server state to change expire the active job instead.
func (a *Api) ServerActiveJobCancel(serverID uint64) (*Job, error) {
	uri := fmt.Sprintf("/servers/%d/cancelActiveJob", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Expiring an active job will not have any influence on the current state
// of the server and is merely an administrative action.
// Often you want to cancel the job, resulting in a server reboot.
// In that case\nuse the /cancelActiveJob API call instead.
func (a *Api) ServerActiveJobExpire(serverID uint64) (*Job, error) {
	uri := fmt.Sprintf("/servers/%d/expireActiveJob", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// A hardware scan collects hardware related information from your server.
// A hardware scan will require a reboot of your server.
// The contents of your hard drive won't be altered in any way.
// After a successful hardware scan your server is booted back into the original operating system.
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

// Install your server with an Operating System and optional Control Panel.
// The default device / partitions to be used are operating system depended
// and can be retrieved via the /v2/operatingSystems/{operatingSystemId} endpoint.
// You are now able to target a specific diskset, like SATA1TB, SATA2TB, SSD256GB, etc.
// To see which disksets are available in your server check the /v2/servers/{serverId} endpoint
// and look for the corresponding diskset id from the hdd array.
func (a *Api) ServerInstallationLaunch(serverID uint64, params *InstallationJob) (*Job, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/install", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// A reset makes sure that your IPMI interface of your server is compatible with Leaseweb automation.
// An IPMI reset will require a reboot of your server.
// The contents of your hard drive won't be altered in any way.
// After a successful IPMI reset your server is booted back into the original operating system."
func (a *Api) ServerIPMIReset(serverID uint64, params *CallbackURL) (*Job, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/ipmiReset", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// List all jobs for this server.
func (a *Api) ServerJobs(serverID uint64) (*Jobs, error) {
	uri := fmt.Sprintf("/servers/%d/jobs", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Jobs
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Get a single job for this server.
func (a *Api) ServerJob(serverID uint64, jobID int) (*Job, error) {
	uri := fmt.Sprintf("/servers/%d/jobs/%d", serverID, jobID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Rescue mode allows you to trouble shoot your server in case your installed
// operating system is no longer reachable.
// You can supply a postInstallScript key in the body of the request which should contain
// a base64 encoded string with a valid script.
// This script will be executed as soon as rescue mode is launched and can be used to further automate
// the process. A requirement for the post install script is that it starts with a shebang
// line like #!/usr/bin/env bash.
// After a rescue mode is launched you can manually reboot the server.
// After this reboot the server will boot into the existing operating system.
// To get a list of available rescue images,
// you could do so by sending a GET request to /bareMetals/v2/rescueImages.
func (a *Api) ServerRescueMode(serverID uint64, params *RescueModeJob) (*Job, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/rescueMode", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Job
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// The credentials API allows you to store usernames and passwords securely.
// During (re)installations, rescue modes and ipmi resets the newly generated passwords
// are stored and can be retrieved using this API.
func (a *Api) ServerCredentials(
	serverID uint64, queryParams map[string]interface{}) (*Credentials, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/credentials%s", serverID, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Credentials
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Password will NOT be updated on the server.
// The ability to update credentials is for convenience only.
// It provides a secure way to communicate passwords with Leaseweb engineers in case support is required.
func (a *Api) ServerCredentialNew(serverID uint64, params *Credential) (*Credential, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/credentials", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Credential
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// List all the credentials filtered by the specified type that are associated with this server.
func (a *Api) ServerTypedCredentials(
	serverID uint64, credType CredType, queryParams map[string]interface{}) (*Credentials, error) {
	if err := credType.Validate(); err != nil {
		return nil, err
	}
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/credentials/%s%s", serverID, credType, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Credentials
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// This action is purely administrative and will only remove the username
// and password associated with this resource from our database.
func (a *Api) ServerUsernameCredentialDelete(
	serverID uint64, credType CredType, username string) (*Error, error) {
	if err := credType.Validate(); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/servers/%d/credentials/%s/%s", serverID, credType, username)
	bodyResp, err := a.NewRequest(NilPayload, uri, "DELETE")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// View the password for the given credential, identified by type and username.
// Auto generated credentials (during a re-install, rescue mode or ipmi reset can be found here).
func (a *Api) ServerUsernameTypedCredentials(
	serverID uint64, credType CredType, username string) (*Credential, error) {
	if err := credType.Validate(); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/servers/%d/credentials/%s/%s", serverID, credType, username)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Credential
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// The usernames or types cannot be changed.
// In order to change those remove this credentials and create a new one.
// This action is purely administrative and will only update the password
// associated with this resource in our database.
func (a *Api) ServerCredentialUpdate(
	serverID uint64, credType CredType, username string, params *Password) (*Credential, error) {
	if err := credType.Validate(); err != nil {
		return nil, err
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}

	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/credentials/%s/%s", serverID, credType, username)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Credential
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// At this moment only bandwidth information for the public interface is supported.
func (a *Api) ServerBandwidthMetrics(serverID uint64, params *BandwidthMetrics) (*Metrics, error) {
	var queryParams map[string]interface{}
	if ok, err := params.Validate(); err != nil {
		return nil, err
	} else {
		queryParams = ok
	}

	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/metrics/bandwidth%s", serverID, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Metrics
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// At this moment only bandwidth information for the public interface is supported.
func (a *Api) ServerDatatraficMetrics(serverID uint64, params *DatatrafficMetrics) (*Metrics, error) {
	var queryParams map[string]interface{}
	if ok, err := params.Validate(); err != nil {
		return nil, err
	} else {
		queryParams = ok
	}

	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/metrics/datatraffic%s", serverID, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Metrics
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// List all bandwith notification settings for this server.
func (a *Api) ServerBandwidthNotifications(
	serverID uint64, queryParams map[string]interface{}) (*BandwidthNotification, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/notificationSettings/bandwidth%s", serverID, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *BandwidthNotification
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Create a new bandwidth notification setting for this server.
func (a *Api) ServerBandwidthNotificationNew(
	serverID uint64, params *NotificationRequest) (*NotificationResponse, error) {
	if err := params.Validete(); err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/notificationSettings/bandwidth", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *NotificationResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Remove a Bandwidth Notification setting for this server.
func (a *Api) ServerBandwidthNotificationDelete(serverID uint64, notificationSettingId int) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/notificationSettings/bandwidth/%d", serverID, notificationSettingId)
	bodyResp, err := a.NewRequest(NilPayload, uri, "DELETE")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Get a bandwidth notification setting for this server.
func (a *Api) ServerBandwidthNotification(
	serverID uint64, notificationSettingId int) (*NotificationResponse, error) {
	uri := fmt.Sprintf("/servers/%d/notificationSettings/bandwidth/%d", serverID, notificationSettingId)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *NotificationResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Update an existing bandwidth notification setting for this server.
func (a *Api) ServerBandwidthNotificationUpdate(
	serverID uint64, notificationSettingId int, params *NotificationRequest) (*NotificationResponse, error) {
	if err := params.Validete(); err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/notificationSettings/bandwidth/%d", serverID, notificationSettingId)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *NotificationResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// List all datatraffic notification settings for this server.
func (a *Api) ServerDatatrafficNotifications(
	serverID uint64, queryParams map[string]interface{}) (*DatatrafficNotification, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/servers/%d/notificationSettings/datatraffic%s", serverID, query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *DatatrafficNotification
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Create a new datatraffic notification setting for this server.
func (a *Api) ServerDatatrafficNotificationNew(
	serverID uint64, params *DataTrafficNotificationRequest) (*NotificationResponse, error) {
	if err := params.Validete(); err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/notificationSettings/datatraffic", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *NotificationResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Delete the given datatraffic notification setting for this server.
func (a *Api) ServerDatatrafficNotificationDelete(serverID uint64, notificationSettingId int) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/notificationSettings/datatraffic/%d", serverID, notificationSettingId)
	bodyResp, err := a.NewRequest(NilPayload, uri, "DELETE")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Get a datatraffic notification setting for this server.
func (a *Api) ServerDatatrafficNotification(
	serverID uint64, notificationSettingId int) (*NotificationResponse, error) {
	uri := fmt.Sprintf("/servers/%d/notificationSettings/datatraffic/%d", serverID, notificationSettingId)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *NotificationResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Update an existing datatraffic notification setting for this server.
func (a *Api) ServerDatatrafficNotificationUpdate(
	serverID uint64,
	notificationSettingId int,
	params *DataTrafficNotificationRequest) (*NotificationResponse, error) {
	if err := params.Validete(); err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/notificationSettings/datatraffic/%d", serverID, notificationSettingId)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *NotificationResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Show all DDoS Protection related notification settings for this server.
// These settings control if you want to be notified via email in case a DDoS was mitigated.
func (a *Api) ServerDDoSNotification(serverID uint64) (*DDoSStatus, error) {
	uri := fmt.Sprintf("/servers/%d/notificationSettings/ddos", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *DDoSStatus
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Update your DDoS notification settings for this server.
func (a *Api) ServerDDoSNotificationUpdate(serverID uint64, params *DDoSStatus) (*Error, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/servers/%d/notificationSettings/ddos", serverID)
	bodyResp, err := a.NewRequest(payload, uri, "PUT")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Powercyle the server.
func (a *Api) ServerPowerCycle(serverID uint64) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/powerCycle", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// The server can either be ON or OFF.
// Servers can be powered on or off by using the respective /powerOn and /powerOff API calls.
// In addition servers can also be rebooted using the /powerCycle API call.
// The pdu object describes the power status from the power distribution unit (PDU) point of view.
// If your server is connected to multiple PDU ports the status property will report
// on if at least one PDU port has power.
// The ipmi object describes the power status by quering the remote management interface of your server.
// Note that pdu.status can report on but your server can still be powered off if
// it was shutdown via IPMI for example.
func (a *Api) ServerPowerStatus(serverID uint64) (*PowerStatus, error) {
	uri := fmt.Sprintf("/servers/%d/powerCycle", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *PowerStatus
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Power off the given server.
func (a *Api) ServerPowerOff(serverID uint64) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/powerOff", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Power on the given server.
func (a *Api) ServerPowerOn(serverID uint64) (*Error, error) {
	uri := fmt.Sprintf("/servers/%d/powerOn", serverID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// An id of a operating system can be supplied when (re)installing a dedicated server
// (for more information on how to install dedicated servers via the API refer to the API documentation).
func (a *Api) OSes(queryParams map[string]interface{}) (*OperatingSystems, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/operatingSystems%s", query)

	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *OperatingSystems
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// This detailed information shows default options when
// installing the given operating system on a dedicated server.
// For some operating systems these defaults can be adjusted when making the POST request to /install.
// If the configurable parameter is true these defaults can be adjusted by the client
func (a *Api) OS(operatingSystemId string, controlPanelId string) (*OSParams, error) {

	uri := fmt.Sprintf("/operatingSystems/%s?controlPanelId=%s", operatingSystemId, controlPanelId)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *OSParams
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// An id of a control panel can be supplied when (re)installing a dedicated server
// (for more information on how to install dedicated servers via the API refer to the API documentation).
func (a *Api) ControlPanels(queryParams map[string]interface{}) (*ControlPanels, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/controlPanels%s", query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *ControlPanels
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) RescueImages(queryParams map[string]interface{}) (*RescueImages, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/rescueImages%s", query)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *RescueImages
	json.Unmarshal(bodyResp, &r)
	return r, err
}
