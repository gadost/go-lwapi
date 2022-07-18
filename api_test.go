package lwapi_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gadost/go-lwapi"
	"github.com/stretchr/testify/assert"
)

func TestMakeQuery(t *testing.T) {
	params := make(map[string]interface{})
	params["offset"] = "10"
	params["count"] = 3
	s := lwapi.MakeQuery(params)

	t.Log(s)
}

var v = make(map[string]interface{})

func TestServers(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":2,"offset":80,"totalCount":132},"servers":[{"assetId":"627293","contract":{"customerId":"32923828192","deliveryStatus":"ACTIVE","id":"674382","reference":"database.server","salesOrgId":"2300"},"featureAvailability":{"automation":true,"ipmiReboot":false,"powerCycle":true,"privateNetwork":true,"remoteManagement":false},"id":"12345","location":{"rack":"A83","site":"AMS-01","suite":"99","unit":"16-17"},"networkInterfaces":{"internal":{"gateway":"10.22.192.12","ip":"10.22.192.3","mac":"AA:BB:CC:DD:EE:FF","ports":[{"name":"EVO-AABB-01","port":"30"}]},"public":{"gateway":"95.211.162.62","ip":"95.211.162.0","mac":"AA:AC:CC:88:EE:E4","ports":[]},"remoteManagement":{"gateway":"10.22.192.126","ip":"10.22.192.1","mac":"AA:AC:CC:88:EE:E4","ports":[]}},"powerPorts":[{"name":"EVO-JV12-APC02","port":"10"}],"privateNetworks":[{"id":"1","linkSpeed":1000,"status":"CONFIGURED","subnet":"127.0.0.80\/24","vlanId":"2120"}],"rack":{"type":"SHARED"}},{"assetId":"627294","contract":{"customerId":"32923828192","deliveryStatus":"ACTIVE","id":"929282","reference":"web.server","salesOrgId":"2300"},"featureAvailability":{"automation":false,"powerCycle":false,"privateNetwork":false,"remoteManagement":false},"id":"47854","location":{"rack":"13","site":"AMS-01","suite":"A6","unit":"18"},"networkInterfaces":{"internal":null,"public":null,"remoteManagement":null},"privateNetworks":[{"id":"2","linkSpeed":1000,"status":"CONFIGURED","subnet":"127.0.0.80\/24","vlanId":"2130"}],"rack":{"type":"SHARED"}}]}`))

	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	v := make(map[string]interface{})
	s, e := api.Servers(v)
	if e != nil {
		t.Error(e)
	}
	//assert.Equal(, s.Metadata.Limit, 2)
	assert.Equal(t, s.Servers[0].AssetID, "627293")
	assert.Equal(t, s.Metadata.TotalCount, 132)

}

func TestServer(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"assetId":"627294","contract":{"billingCycle":12,"billingFrequency":"MONTH","contractTerm":12,"currency":"EUR","customerId":"32923828192","deliveryStatus":"ACTIVE","endsAt":"2017-10-01T01:00:00+0100","id":"674382","networkTraffic":{"datatrafficLimit":100,"datatrafficUnit":"TB|Mbps","trafficType":"PREMIUM|VOLUME","type":"95|FLATFEE|DATATRAFFIC"},"pricePerFrequency":49,"privateNetworks":[{"id":"1","linkSpeed":1000,"status":"CONFIGURED","subnet":"127.0.0.80\/24","vlanId":"2120"}],"reference":"database.server","salesOrgId":"2300","sla":"BRONZE","softwareLicenses":[{"currency":"EUR","name":"WINDOWS_2012_R2_SERVER","price":12.12}],"startsAt":"2014-01-01T01:00:00+0100"},"featureAvailability":{"automation":true,"ipmiReboot":false,"powerCycle":true,"privateNetwork":true,"remoteManagement":false},"id":"12345","location":{"rack":"13","site":"AMS-01","suite":"A6","unit":"16-17"},"networkInterfaces":{"internal":{"gateway":"123.123.123.126","ip":"123.123.123.123\/27","mac":"AA:BB:CC:DD:EE:FF","ports":[]},"public":{"gateway":"123.123.123.126","ip":"123.123.123.123\/27","mac":"AA:BB:CC:DD:EE:FF","ports":[{"name":"EVO-JV12-1","port":"33"}]},"remoteManagement":null},"powerPorts":[{"name":"EVO-JV12-APC02","port":"10"}],"rack":{"type":"PRIVATE"},"serialNumber":"JDK18291JK","specs":{"chassis":"Dell R210 II","cpu":{"quantity":4,"type":"Intel Xeon E3-1220"},"hardwareRaidCapable":true,"hdd":[{"amount":2,"id":"SATA2TB","performanceType":null,"size":2,"type":"SATA","unit":"TB"}],"pciCards":[{"description":"2x10GE UTP card"},{"description":"2x30GE UTP card"}],"ram":{"size":32,"unit":"GB"}}}`))

	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.Server(12345)
	if e != nil {
		t.Error(e)
	}
	//assert.Equal(, s.Metadata.Limit, 2)
	assert.Equal(t, s.AssetID, "627294")
	assert.Equal(t, s.Contract.CustomerID, "32923828192")

}

func TestServerReferenceUpdate(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345")
		assert.Equal(t, req.Method, "PUT")
		// Send response to be tested
		rw.Write([]byte(`{"assetId":"627294","contract":{"billingCycle":12,"billingFrequency":"MONTH","contractTerm":12,"currency":"EUR","customerId":"32923828192","deliveryStatus":"ACTIVE","endsAt":"2017-10-01T01:00:00+0100","id":"674382","networkTraffic":{"datatrafficLimit":100,"datatrafficUnit":"TB|Mbps","trafficType":"PREMIUM|VOLUME","type":"95|FLATFEE|DATATRAFFIC"},"pricePerFrequency":49,"privateNetworks":[{"id":"1","linkSpeed":1000,"status":"CONFIGURED","subnet":"127.0.0.80\/24","vlanId":"2120"}],"reference":"database.server","salesOrgId":"2300","sla":"BRONZE","softwareLicenses":[{"currency":"EUR","name":"WINDOWS_2012_R2_SERVER","price":12.12}],"startsAt":"2014-01-01T01:00:00+0100"},"featureAvailability":{"automation":true,"ipmiReboot":false,"powerCycle":true,"privateNetwork":true,"remoteManagement":false},"id":"12345","location":{"rack":"13","site":"AMS-01","suite":"A6","unit":"16-17"},"networkInterfaces":{"internal":{"gateway":"123.123.123.126","ip":"123.123.123.123\/27","mac":"AA:BB:CC:DD:EE:FF","ports":[]},"public":{"gateway":"123.123.123.126","ip":"123.123.123.123\/27","mac":"AA:BB:CC:DD:EE:FF","ports":[{"name":"EVO-JV12-1","port":"33"}]},"remoteManagement":null},"powerPorts":[{"name":"EVO-JV12-APC02","port":"10"}],"rack":{"type":"PRIVATE"},"serialNumber":"JDK18291JK","specs":{"chassis":"Dell R210 II","cpu":{"quantity":4,"type":"Intel Xeon E3-1220"},"hardwareRaidCapable":true,"hdd":[{"amount":2,"id":"SATA2TB","performanceType":null,"size":2,"type":"SATA","unit":"TB"}],"pciCards":[{"description":"2x10GE UTP card"},{"description":"2x30GE UTP card"}],"ram":{"size":32,"unit":"GB"}}}`))

	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerReferenceUpdate(12345, &lwapi.Reference{
		Reference: "test",
	})
	if e != nil {
		t.Error(e)
	}
	//assert.Equal(, s.Metadata.Limit, 2)
	assert.Equal(t, s.AssetID, "627294")
	assert.Equal(t, s.Contract.CustomerID, "32923828192")

}

func TestServerHardwareInformation(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/hardwareInfo")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"id":"2378237","parserVersion":"3.6","result":{"chassis":{"description":"Rack Mount Chassis","firmware":{"date":"07\/01\/2013","description":"BIOS","vendor":"HP","version":"J01"},"motherboard":{"product":"","serial":"","vendor":""},"product":"ProLiant DL120 G7 (647339-B21)","serial":"CZ33109CHV","vendor":"HP"},"cpu":[{"capabilities":{"cpufreq":"CPU Frequency scaling","ht":"HyperThreading","vmx":false,"x86-64":"64bits extensions (x86-64)"},"description":"Intel(R) Xeon(R) CPU E31230","hz":"2792640000","serial_number":"","settings":{"cores":"4","enabledcores":"4","threads":"8"},"slot":"Proc 1","vendor":"Intel Corp."}],"disks":[{"description":"ATA Disk","id":"disk:0","product":"Hitachi HDS72302","serial_number":"MS77215W07S6SA","size":"2000398934016","smartctl":{"ata_version":"ATA8-ACS T13\/1699-D revision 4","attributes":{"Power_On_Hours":{"flag":"0x0012","id":"  9","raw_value":"39832","thresh":"000","type":"Old_age","updated":"Always","value":"095","when_failed":"-","worst":"095"},"Reallocated_Sector_Ct":{"flag":"0x0033","id":"  5","raw_value":"0","thresh":"005","type":"Pre-fail","updated":"Always","value":"100","when_failed":"-","worst":"100"}},"device_model":"Hitachi HDS723020BLE640","execution_status":"0","firmware_version":"MX4OAAB0","is_sas":false,"overall_health":"PASSED","rpm":"7200 rpm","sata_version":"SATA 3.0, 6.0 Gb\/s (current: 6.0 Gb\/s)","sector_size":"512 bytes logical, 4096 bytes physical","serial_number":"MS77215W07S6SA","smart_error_log":"No Errors Logged","smart_support":{"available":true,"enabled":true},"smartctl_version":"6.2","user_capacity":"2,000,398,934,016 bytes [2.00 TB]"},"vendor":"Hitachi"}],"ipmi":{"defgateway":"10.19.79.126","firmware":"1.88","ipaddress":"10.19.79.67","ipsource":"DHCP Address","macaddress":"28:92:4a:33:48:e8","subnetmask":"255.255.255.192","vendor":"Hewlett-Packard"},"memory":[{"clock_hz":"1333000000","description":"DIMM DDR3 Synchronous 1333 MHz (0.8 ns)","id":"memory\/bank:0","serial_number":"8369AF58","size_bytes":"4294967296"},{"clock_hz":"1333000000","description":"DIMM DDR3 Synchronous 1333 MHz (0.8 ns)","id":"memory\/bank:1","serial_number":"8369B174","size_bytes":"4294967296"}],"network":[{"capabilities":{"autonegotiation":"Auto-negotiation","bus_master":"bus mastering","cap_list":"PCI capabilities listing","ethernet":"","link_speeds":{"1000bt-fd":"1Gbit\/s (full duplex)","100bt":"100Mbit\/s","100bt-fd":"100Mbit\/s (full duplex)","10bt":"10Mbit\/s","10bt-fd":"10Mbit\/s (full duplex)"},"msi":"Message Signalled Interrupts","msix":"MSI-X","pciexpress":"PCI Express","physical":"Physical interface","pm":"Power Management","tp":"twisted pair"},"lldp":{"chassis":{"description":"Juniper Networks, Inc. ex3300-48t Ethernet Switch, kernel JUNOS 15.1R5.5, Build date: 2016-11-25 16:02:59 UTC Copyright (c) 1996-2016 Juniper Networks, Inc.","mac_address":"4c:16:fc:3a:84:c0","name":"EVO-NS19-1"},"port":{"auto_negotiation":{"enabled":"yes","supported":"yes"},"description":"ge-0\/0\/2.0"},"vlan":{"id":"0","label":"VLAN","name":"default"}},"logical_name":"eth0","mac_address":"28:92:4a:33:48:e6","product":"82574L Gigabit Network Connection","settings":{"autonegotiation":"on","broadcast":"yes","driver":"e1000e","driverversion":"3.2.6-k","duplex":"full","firmware":"2.1-2","ip":"212.32.230.67","latency":"0","link":"yes","multicast":"yes","port":"twisted pair","speed":"1Gbit\/s"},"vendor":"Intel Corporation"}]},"scannedAt":"2017-09-27T14:21:01Z","serverId":"62264"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerHardwareInformation(12345)
	if e != nil {
		t.Error(e)
	}
	//assert.Equal(, s.Metadata.Limit, 2)
	assert.Equal(t, s.ID, "2378237")
	assert.Equal(t, s.Result.Chassis.Firmware.Vendor, "HP")

}

func TestServerIPList(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/ips")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":10,"offset":0,"totalCount":2},"ips":[{"ddos":{"detectionProfile":"ADVANCED_LOW_UDP","protectionType":"ADVANCED"},"floatingIp":false,"gateway":"12.123.123.254","ip":"12.123.123.1\/24","mainIp":true,"networkType":"PUBLIC","nullRouted":true,"reverseLookup":"domain.example.com","version":4},{"ddos":{"detectionProfile":"STANDARD_DEFAULT","protectionType":"STANDARD"},"floatingIp":false,"gateway":"2001:db8:85a3::8a2e:370:1","ip":"2001:db8:85a3::8a2e:370:7334\/64","mainIp":false,"networkType":"REMOTE_MANAGEMENT","nullRouted":false,"reverseLookup":"domain.example.com","version":6}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerIPList(12345, v)
	if e != nil {
		t.Error(e)
	}
	//assert.Equal(, s.Metadata.Limit, 2)
	assert.Equal(t, s.Metadata.Offset, 0)
	assert.Equal(t, s.Ips[0].Ddos.ProtectionType, "ADVANCED")
}

func TestServerIP(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/ips/127.0.0.1")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"ddos":{"detectionProfile":"ADVANCED_LOW_UDP","protectionType":"ADVANCED"},"floatingIp":false,"gateway":"12.123.123.254","ip":"12.123.123.1\/24","mainIp":true,"networkType":"PUBLIC","nullRouted":false,"reverseLookup":"domain.example.com","version":4}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerIP(12345, "127.0.0.1")
	if e != nil {
		t.Error(e)
	}
	//assert.Equal(, s.Metadata.Limit, 2)
	assert.Equal(t, s.Ddos.DetectionProfile, "ADVANCED_LOW_UDP")
	assert.Equal(t, s.FloatingIP, false)

}

func TestServerIPUpdate(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/ips/127.0.0.1")
		assert.Equal(t, req.Method, "PUT")
		// Send response to be tested
		rw.Write([]byte(`{"ddos":{"detectionProfile":"ADVANCED_DEFAULT","protectionType":"ADVANCED"},"floatingIp":false,"gateway":"12.123.123.254","ip":"12.123.123.1\/24","mainIp":true,"networkType":"PUBLIC","nullRouted":false,"reverseLookup":"domain.example.com","version":4}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerIPUpdate(12345, "127.0.0.1", &lwapi.UpdateIPRequest{
		DetectionProfile: "ADVANCED_DEFAULT",
		ReverseLookup:    "example.com",
	})
	if e != nil {
		t.Error(e)
	}
	//assert.Equal(, s.Metadata.Limit, 2)
	assert.Equal(t, s.Ddos.DetectionProfile, "ADVANCED_DEFAULT")
	assert.Equal(t, s.FloatingIP, false)
}

func TestServerIPNull(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/ips/127.0.0.1/null")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"ddos":{"detectionProfile":"ADVANCED_DEFAULT","protectionType":"ADVANCED"},"floatingIp":false,"gateway":"12.123.123.254","ip":"12.123.123.1\/24","mainIp":true,"networkType":"PUBLIC","nullRouted":false,"reverseLookup":"domain.example.com","version":4}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerIPNull(12345, "127.0.0.1")
	if e != nil {
		t.Error(e)
	}
	//assert.Equal(, s.Metadata.Limit, 2)
	assert.Equal(t, s.Ddos.DetectionProfile, "ADVANCED_DEFAULT")
	assert.Equal(t, s.FloatingIP, false)

}

func TestServerIPUnNull(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/ips/127.0.0.1/unnull")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"ddos":{"detectionProfile":"ADVANCED_DEFAULT","protectionType":"ADVANCED"},"floatingIp":false,"gateway":"12.123.123.254","ip":"12.123.123.1\/24","mainIp":true,"networkType":"PUBLIC","nullRouted":false,"reverseLookup":"domain.example.com","version":4}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerIPUnNull(12345, "127.0.0.1")
	if e != nil {
		t.Error(e)
	}
	//assert.Equal(, s.Metadata.Limit, 2)
	assert.Equal(t, s.Ddos.DetectionProfile, "ADVANCED_DEFAULT")
	assert.Equal(t, s.FloatingIP, false)

}

func TestServerIPNullHistory(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/ips/127.0.0.1/nullRouteHistory")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":10,"offset":0,"totalCount":1},"nullRoutes":[{"automatedUnnullingAt":"2016-08-12T07:45:33+00:00","comment":"Device Null Route related to DDoS Mitigation","ip":"1.1.1.1\/32","nullLevel":3,"nulledAt":"2016-08-12T07:40:27+00:00","ticketId":"282912"}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerIPNullHistory(12345, "127.0.0.1", v)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.NullRoutes[0].IP, "1.1.1.1/32")

}

func TestServerNetworkInterfaces(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/networkInterfaces")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":10,"offset":0,"totalCount":1},"networkInterfaces":[{"linkSpeed":"100Mbps","operStatus":"OPEN","status":"OPEN","switchInterface":"33","switchName":"EVO-AA11-1","type":"PUBLIC"}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerNetworkInterfaces(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.NetworkInterfaces[0].Status, "OPEN")

}

var err = `{"errorCode":"APP00800","errorMessage":"The connection with the DB cannot be established.","correlationId":"550e8400-e29b-41d4-a716-446655440000","userMessage":"Cannot handle your request at the moment. Please try again later.","reference":"https:\/\/developer.leaseweb.com\/errors\/APP00800"}`

func TestServerNetworkInterfacesClose(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/networkInterfaces/close")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerNetworkInterfacesClose(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")

}

func TestServerNetworkInterfacesOpen(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/networkInterfaces/open")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerNetworkInterfacesOpen(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")

}

func TestServerNetworkInterface(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/networkInterfaces/public")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"linkSpeed":"100Mbps","operStatus":"OPEN","status":"OPEN","switchInterface":"33","switchName":"EVO-JV12-1","type":"PUBLIC"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerNetworkInterface(12345, "public")
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Status, "OPEN")

}

func TestServerNetworkInterfaceClose(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/networkInterfaces/public/close")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerNetworkInterfaceClose(12345, "public")
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")

}

func TestServerNetworkInterfaceOpen(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/networkInterfaces/public/open")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerNetworkInterfaceOpen(12345, "public")
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")

}

func TestServerPrivateNetworkDelete(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/privateNetworks/1")
		assert.Equal(t, req.Method, "DELETE")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerPrivateNetworkDelete(12345, 1)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")

}

func TestServerPrivateNetworkAdd(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/privateNetworks/1")
		assert.Equal(t, req.Method, "PUT")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerPrivateNetworkAdd(12345, 1, 100)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")

}

func TestServerDHCPLeaseDelete(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/leases")
		assert.Equal(t, req.Method, "DELETE")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDHCPLeaseDelete(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")

}

func TestServerDHCPLeases(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/leases")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDHCPLeases(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")

}

func TestServerDHCPLeaseNew(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/leases")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDHCPLeaseNew(12345, &lwapi.ServerDHCPLeaseNew{
		Bootfile: "http://ss.ss",
		Hostname: "example.com",
	})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")

}

func TestServerActiveJobCancel(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/cancelActiveJob")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"createdAt":"2021-01-09T08:54:06+0000","flow":"#stop","isRunning":false,"node":"80:18:44:E0:AF:C4!JGNTQ92","payload":{"configurable":true,"device":"SATA2TB","fileserverBaseUrl":"","jobType":"install","numberOfDisks":null,"operatingSystemId":"UBUNTU_20_04_64BIT","os":{"architecture":"64bit","family":"ubuntu","name":"Ubuntu 20.04 LTS (Focal Fossa) (amd64)","type":"linux","version":"20.04"},"partitions":[{"filesystem":"swap","size":4096}],"pop":"AMS-01","powerCycle":true,"raidLevel":null,"serverId":"99944","timezone":"UTC","x":1},"progress":{"canceled":1,"expired":0,"failed":0,"finished":0,"inprogress":0,"pending":0,"percentage":0,"total":1,"waiting":0},"serverId":"99944","status":"CANCELED","tasks":[{"description":"dummy","errorMessage":"The job was canceled by the api consumer","flow":"tasks","onError":"break","status":"CANCELED","statusTimestamps":{"CANCELED":"2021-01-09T08:54:15+00:00","PENDING":"2021-01-09T08:54:06+00:00","WAITING":"2021-01-09T08:54:06+00:00"},"uuid":"085ce145-39bd-4cb3-8e2b-53f17a97a463"}],"type":"install","updatedAt":"2021-01-09T08:54:15+0000","uuid":"c77d8a6b-d255-4744-8b95-8bf4af6f8b48"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerActiveJobCancel(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Flow, "#stop")

}

func TestServerActiveJobExpire(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/expireActiveJob")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"createdAt":"2021-01-09T08:54:06+0000","flow":"#stop","isRunning":false,"node":"80:18:44:E0:AF:C4!JGNTQ92","payload":{"configurable":true,"device":"SATA2TB","fileserverBaseUrl":"","jobType":"install","numberOfDisks":null,"operatingSystemId":"UBUNTU_20_04_64BIT","os":{"architecture":"64bit","family":"ubuntu","name":"Ubuntu 20.04 LTS (Focal Fossa) (amd64)","type":"linux","version":"20.04"},"partitions":[{"filesystem":"swap","size":4096}],"pop":"AMS-01","powerCycle":true,"raidLevel":null,"serverId":"99944","timezone":"UTC","x":1},"progress":{"canceled":1,"expired":0,"failed":0,"finished":0,"inprogress":0,"pending":0,"percentage":0,"total":1,"waiting":0},"serverId":"99944","status":"CANCELED","tasks":[{"description":"dummy","errorMessage":"The job was canceled by the api consumer","flow":"tasks","onError":"break","status":"CANCELED","statusTimestamps":{"CANCELED":"2021-01-09T08:54:15+00:00","PENDING":"2021-01-09T08:54:06+00:00","WAITING":"2021-01-09T08:54:06+00:00"},"uuid":"085ce145-39bd-4cb3-8e2b-53f17a97a463"}],"type":"install","updatedAt":"2021-01-09T08:54:15+0000","uuid":"c77d8a6b-d255-4744-8b95-8bf4af6f8b48"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerActiveJobExpire(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Flow, "#stop")
}

func TestServerHardwareScan(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/hardwareScan")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"createdAt":"2021-01-09T08:54:06+0000","flow":"#stop","isRunning":false,"node":"80:18:44:E0:AF:C4!JGNTQ92","payload":{"configurable":true,"device":"SATA2TB","fileserverBaseUrl":"","jobType":"install","numberOfDisks":null,"operatingSystemId":"UBUNTU_20_04_64BIT","os":{"architecture":"64bit","family":"ubuntu","name":"Ubuntu 20.04 LTS (Focal Fossa) (amd64)","type":"linux","version":"20.04"},"partitions":[{"filesystem":"swap","size":4096}],"pop":"AMS-01","powerCycle":true,"raidLevel":null,"serverId":"99944","timezone":"UTC","x":1},"progress":{"canceled":1,"expired":0,"failed":0,"finished":0,"inprogress":0,"pending":0,"percentage":0,"total":1,"waiting":0},"serverId":"99944","status":"CANCELED","tasks":[{"description":"dummy","errorMessage":"The job was canceled by the api consumer","flow":"tasks","onError":"break","status":"CANCELED","statusTimestamps":{"CANCELED":"2021-01-09T08:54:15+00:00","PENDING":"2021-01-09T08:54:06+00:00","WAITING":"2021-01-09T08:54:06+00:00"},"uuid":"085ce145-39bd-4cb3-8e2b-53f17a97a463"}],"type":"install","updatedAt":"2021-01-09T08:54:15+0000","uuid":"c77d8a6b-d255-4744-8b95-8bf4af6f8b48"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerHardwareScan(12345, true, "")
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Flow, "#stop")
}

func TestServerInstallationLaunch(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/install")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"createdAt":"2021-01-09T08:54:06+0000","flow":"#stop","isRunning":false,"node":"80:18:44:E0:AF:C4!JGNTQ92","payload":{"configurable":true,"device":"SATA2TB","fileserverBaseUrl":"","jobType":"install","numberOfDisks":null,"operatingSystemId":"UBUNTU_20_04_64BIT","os":{"architecture":"64bit","family":"ubuntu","name":"Ubuntu 20.04 LTS (Focal Fossa) (amd64)","type":"linux","version":"20.04"},"partitions":[{"filesystem":"swap","size":4096}],"pop":"AMS-01","powerCycle":true,"raidLevel":null,"serverId":"99944","timezone":"UTC","x":1},"progress":{"canceled":1,"expired":0,"failed":0,"finished":0,"inprogress":0,"pending":0,"percentage":0,"total":1,"waiting":0},"serverId":"99944","status":"CANCELED","tasks":[{"description":"dummy","errorMessage":"The job was canceled by the api consumer","flow":"tasks","onError":"break","status":"CANCELED","statusTimestamps":{"CANCELED":"2021-01-09T08:54:15+00:00","PENDING":"2021-01-09T08:54:06+00:00","WAITING":"2021-01-09T08:54:06+00:00"},"uuid":"085ce145-39bd-4cb3-8e2b-53f17a97a463"}],"type":"install","updatedAt":"2021-01-09T08:54:15+0000","uuid":"c77d8a6b-d255-4744-8b95-8bf4af6f8b48"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerInstallationLaunch(12345, &lwapi.InstallationJob{OperatingSystemID: "s"})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Flow, "#stop")
}

func TestServerIPMIReset(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/ipmiReset")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"createdAt":"2021-01-09T08:54:06+0000","flow":"#stop","isRunning":false,"node":"80:18:44:E0:AF:C4!JGNTQ92","payload":{"configurable":true,"device":"SATA2TB","fileserverBaseUrl":"","jobType":"install","numberOfDisks":null,"operatingSystemId":"UBUNTU_20_04_64BIT","os":{"architecture":"64bit","family":"ubuntu","name":"Ubuntu 20.04 LTS (Focal Fossa) (amd64)","type":"linux","version":"20.04"},"partitions":[{"filesystem":"swap","size":4096}],"pop":"AMS-01","powerCycle":true,"raidLevel":null,"serverId":"99944","timezone":"UTC","x":1},"progress":{"canceled":1,"expired":0,"failed":0,"finished":0,"inprogress":0,"pending":0,"percentage":0,"total":1,"waiting":0},"serverId":"99944","status":"CANCELED","tasks":[{"description":"dummy","errorMessage":"The job was canceled by the api consumer","flow":"tasks","onError":"break","status":"CANCELED","statusTimestamps":{"CANCELED":"2021-01-09T08:54:15+00:00","PENDING":"2021-01-09T08:54:06+00:00","WAITING":"2021-01-09T08:54:06+00:00"},"uuid":"085ce145-39bd-4cb3-8e2b-53f17a97a463"}],"type":"install","updatedAt":"2021-01-09T08:54:15+0000","uuid":"c77d8a6b-d255-4744-8b95-8bf4af6f8b48"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerIPMIReset(12345, &lwapi.CallbackURL{})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Flow, "#stop")
}

func TestServerJobs(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/jobs")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":10,"offset":0,"totalCount":1},"jobs":[{"createdAt":"2018-01-09T10:38:12+0000","flow":"tasks","isRunning":true,"node":"80:18:44:E0:AF:C4!JGNTQ92","progress":{"canceled":0,"expired":0,"failed":0,"finished":0,"inprogress":0,"pending":1,"percentage":0,"total":1,"waiting":0},"serverId":"99944","status":"ACTIVE","type":"install","updatedAt":"2018-01-09T10:38:12+0000","uuid":"3a867358-5b4b-44ee-88ac-4274603ef641"}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerJobs(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Jobs[0].Flow, "tasks")
}

func TestServerJob(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/jobs/1")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"createdAt":"2021-01-09T08:54:06+0000","flow":"#stop","isRunning":false,"node":"80:18:44:E0:AF:C4!JGNTQ92","payload":{"configurable":true,"device":"SATA2TB","fileserverBaseUrl":"","jobType":"install","numberOfDisks":null,"operatingSystemId":"UBUNTU_20_04_64BIT","os":{"architecture":"64bit","family":"ubuntu","name":"Ubuntu 20.04 LTS (Focal Fossa) (amd64)","type":"linux","version":"20.04"},"partitions":[{"filesystem":"swap","size":4096}],"pop":"AMS-01","powerCycle":true,"raidLevel":null,"serverId":"99944","timezone":"UTC","x":1},"progress":{"canceled":1,"expired":0,"failed":0,"finished":0,"inprogress":0,"pending":0,"percentage":0,"total":1,"waiting":0},"serverId":"99944","status":"CANCELED","tasks":[{"description":"dummy","errorMessage":"The job was canceled by the api consumer","flow":"tasks","onError":"break","status":"CANCELED","statusTimestamps":{"CANCELED":"2021-01-09T08:54:15+00:00","PENDING":"2021-01-09T08:54:06+00:00","WAITING":"2021-01-09T08:54:06+00:00"},"uuid":"085ce145-39bd-4cb3-8e2b-53f17a97a463"}],"type":"install","updatedAt":"2021-01-09T08:54:15+0000","uuid":"c77d8a6b-d255-4744-8b95-8bf4af6f8b48"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerJob(12345, 1)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Flow, "#stop")
}

func TestServerRescueMode(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/rescueMode")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"createdAt":"2021-01-09T08:54:06+0000","flow":"#stop","isRunning":false,"node":"80:18:44:E0:AF:C4!JGNTQ92","payload":{"configurable":true,"device":"SATA2TB","fileserverBaseUrl":"","jobType":"install","numberOfDisks":null,"operatingSystemId":"UBUNTU_20_04_64BIT","os":{"architecture":"64bit","family":"ubuntu","name":"Ubuntu 20.04 LTS (Focal Fossa) (amd64)","type":"linux","version":"20.04"},"partitions":[{"filesystem":"swap","size":4096}],"pop":"AMS-01","powerCycle":true,"raidLevel":null,"serverId":"99944","timezone":"UTC","x":1},"progress":{"canceled":1,"expired":0,"failed":0,"finished":0,"inprogress":0,"pending":0,"percentage":0,"total":1,"waiting":0},"serverId":"99944","status":"CANCELED","tasks":[{"description":"dummy","errorMessage":"The job was canceled by the api consumer","flow":"tasks","onError":"break","status":"CANCELED","statusTimestamps":{"CANCELED":"2021-01-09T08:54:15+00:00","PENDING":"2021-01-09T08:54:06+00:00","WAITING":"2021-01-09T08:54:06+00:00"},"uuid":"085ce145-39bd-4cb3-8e2b-53f17a97a463"}],"type":"install","updatedAt":"2021-01-09T08:54:15+0000","uuid":"c77d8a6b-d255-4744-8b95-8bf4af6f8b48"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerRescueMode(12345, &lwapi.RescueModeJob{RescueImageID: "s"})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Flow, "#stop")
}

func TestServerCredentials(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/credentials")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":10,"offset":0,"totalCount":4},"credentials":[{"type":"REMOTE_MANAGEMENT","username":"admin"},{"type":"REMOTE_MANAGEMENT","username":"root"},{"type":"OPERATING_SYSTEM","username":"root"},{"type":"OPERATING_SYSTEM","username":"user"}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerCredentials(12345, v)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Credentials[0].Type, "REMOTE_MANAGEMENT")
}

func TestServerCredentialNew(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/credentials")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{
			"password": "mys3cr3tp@ssw0rd",
			"type": "OPERATING_SYSTEM",
			"username": "root"
		  }`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerCredentialNew(12345, &lwapi.Credential{
		Type:     "OPERATING_SYSTEM",
		Username: "user",
		Password: "pass",
		Error:    lwapi.Error{},
	})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Password, "mys3cr3tp@ssw0rd")
}

func TestServerTypedCredentials(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/credentials/RESCUE_MODE")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{
			"_metadata": {
			  "limit": 10,
			  "offset": 0,
			  "totalCount": 1
			},
			"credentials": [
			  {
				"type": "OPERATING_SYSTEM",
				"username": "root"
			  }
			]
		  }`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerTypedCredentials(12345, "RESCUE_MODE", v)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Credentials[0].Username, "root")
}

func TestServerUsernameCredentialDelete(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/credentials/RESCUE_MODE/s")
		assert.Equal(t, req.Method, "DELETE")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerUsernameCredentialDelete(12345, "RESCUE_MODE", "s")
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")
}

func TestServerUsernameTypedCredentials(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/credentials/RESCUE_MODE/s")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{
			"password": "mys3cr3tp@ssw0rd",
			"type": "OPERATING_SYSTEM",
			"username": "root"
		  }`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerUsernameTypedCredentials(12345, "RESCUE_MODE", "s")
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Password, "mys3cr3tp@ssw0rd")
}

func TestServerCredentialUpdate(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/credentials/RESCUE_MODE/s")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{
			"password": "mys3cr3tp@ssw0rd",
			"type": "OPERATING_SYSTEM",
			"username": "root"
		  }`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerCredentialUpdate(12345, "RESCUE_MODE", "s", &lwapi.Password{Password: "pass"})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Password, "mys3cr3tp@ssw0rd")
}

func TestServerBandwidthMetrics(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		//assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/metrics/bandwidth?from=2016-10-20T09:00:00Z&to=2016-10-20T09:00:00Z&aggregation=95TH")
		assert.Equal(t, req.Method, "GET")
		assert.Equal(t, req.URL.Path, "/bareMetals/v2/servers/12345/metrics/bandwidth")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"aggregation":"AVG","from":"2016-10-20T09:00:00Z","granularity":"HOUR","to":"2016-10-20T11:00:00Z"},"metrics":{"DOWN_PUBLIC":{"unit":"bps","values":[{"timestamp":"2016-10-20T09:00:00Z","value":202499},{"timestamp":"2016-10-20T10:00:00Z","value":29900}]},"UP_PUBLIC":{"unit":"bps","values":[{"timestamp":"2016-10-20T09:00:00Z","value":43212393},{"timestamp":"2016-10-20T10:00:00Z","value":12342929}]}}}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerBandwidthMetrics(12345, &lwapi.BandwidthMetrics{
		From:        "2016-10-20T09:00:00Z",
		To:          "2016-10-20T09:00:00Z",
		Aggregation: "95TH",
	})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Metrics.DownPublic.Unit, "bps")
}

func TestServerDatatraficMetrics(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.Path, "/bareMetals/v2/servers/12345/metrics/datatraffic")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"aggregation":"SUM","from":"2016-10-20T09:00:00Z","granularity":"HOUR","to":"2016-10-20T11:00:00Z"},"metrics":{"DOWN_PUBLIC":{"unit":"B","values":[{"timestamp":"2016-10-20T09:00:00Z","value":202499},{"timestamp":"2016-10-20T10:00:00Z","value":29900}]},"UP_PUBLIC":{"unit":"B","values":[{"timestamp":"2016-10-20T09:00:00Z","value":43212393},{"timestamp":"2016-10-20T10:00:00Z","value":12342929}]}}}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDatatraficMetrics(12345, &lwapi.DatatrafficMetrics{
		From:        "2016-10-20T09:00:00Z",
		To:          "2016-10-20T09:00:00Z",
		Aggregation: "SUM",
	})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Metrics.DownPublic.Unit, "B")
}

func TestServerBandwidthNotifications(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/bandwidth")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":10,"offset":0,"totalCount":2},"bandwidthNotificationSettings":[{"actions":[{"lastTriggeredAt":"2021-03-16T01:01:44+00:00","type":"EMAIL"}],"frequency":"WEEKLY","id":"12345","lastCheckedAt":"2021-03-16T01:01:41+00:00","threshold":"1","thresholdExceededAt":"2021-03-16T01:01:41+00:00","unit":"Gbps"},{"actions":[{"lastTriggeredAt":"2021-03-16T01:01:44+00:00","type":"EMAIL"}],"frequency":"DAILY","id":"123456","lastCheckedAt":"2021-03-16T01:01:41+00:00","threshold":"1","thresholdExceededAt":"2021-03-16T01:01:41+00:00","unit":"Mbps"}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerBandwidthNotifications(12345, v)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.BandwidthNotificationSettings[0].Actions[0].Type, "EMAIL")
}

func TestServerBandwidthNotificationNew(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/bandwidth")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"actions":[{"lastTriggeredAt":"2021-03-16T01:01:44+00:00","type":"EMAIL"}],"frequency":"WEEKLY","id":"12345","lastCheckedAt":"2021-03-16T01:01:41+00:00","threshold":"1","thresholdExceededAt":"2021-03-16T01:01:41+00:00","unit":"Gbps"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerBandwidthNotificationNew(12345, &lwapi.NotificationRequest{
		Frequency: "DAILY",
		Threshold: "s",
		Unit:      "Mbps",
	})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Actions[0].Type, "EMAIL")
}

func TestServerBandwidthNotificationDelete(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/bandwidth/1")
		assert.Equal(t, req.Method, "DELETE")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerBandwidthNotificationDelete(12345, 1)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")
}

func TestServerBandwidthNotification(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/bandwidth/1")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"actions":[{"lastTriggeredAt":"2021-03-16T01:01:44+00:00","type":"EMAIL"}],"frequency":"WEEKLY","id":"12345","lastCheckedAt":"2021-03-16T01:01:41+00:00","threshold":"1","thresholdExceededAt":"2021-03-16T01:01:41+00:00","unit":"Gbps"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerBandwidthNotification(12345, 1)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Actions[0].Type, "EMAIL")
}

func TestServerBandwidthNotificationUpdate(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/bandwidth/1")
		assert.Equal(t, req.Method, "PUT")
		// Send response to be tested
		rw.Write([]byte(`{"actions":[{"lastTriggeredAt":"2021-03-16T01:01:44+00:00","type":"EMAIL"}],"frequency":"WEEKLY","id":"12345","lastCheckedAt":"2021-03-16T01:01:41+00:00","threshold":"1","thresholdExceededAt":"2021-03-16T01:01:41+00:00","unit":"Gbps"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerBandwidthNotificationUpdate(12345, 1, &lwapi.NotificationRequest{
		Frequency: "DAILY",
		Threshold: "s",
		Unit:      "Mbps",
	})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Actions[0].Type, "EMAIL")
}

func TestServerDatatrafficNotifications(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/datatraffic")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":10,"offset":0,"totalCount":2},"datatrafficNotificationSettings":[{"actions":[{"lastTriggeredAt":null,"type":"EMAIL"}],"frequency":"WEEKLY","id":"12345","lastCheckedAt":null,"threshold":"1","thresholdExceededAt":null,"unit":"MB"},{"actions":[{"lastTriggeredAt":null,"type":"EMAIL"}],"frequency":"DAILY","id":"123456","lastCheckedAt":null,"threshold":"1","thresholdExceededAt":null,"unit":"GB"}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDatatrafficNotifications(12345, v)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.DatatrafficNotificationSettings[0].Actions[0].Type, "EMAIL")
}

func TestServerDatatrafficNotificationNew(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/datatraffic")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(`{"actions":[{"lastTriggeredAt":"2021-03-16T01:01:44+00:00","type":"EMAIL"}],"frequency":"WEEKLY","id":"12345","lastCheckedAt":"2021-03-16T01:01:41+00:00","threshold":"1","thresholdExceededAt":"2021-03-16T01:01:41+00:00","unit":"Gbps"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDatatrafficNotificationNew(12345, &lwapi.DataTrafficNotificationRequest{
		Frequency: "DAILY",
		Threshold: "s",
		Unit:      "MB",
	})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Actions[0].Type, "EMAIL")
}

func TestServerDatatrafficNotificationDelete(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/datatraffic/1")
		assert.Equal(t, req.Method, "DELETE")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDatatrafficNotificationDelete(12345, 1)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")
}

func TestServerDatatrafficNotification(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/datatraffic/1")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"actions":[{"lastTriggeredAt":"2021-03-16T01:01:44+00:00","type":"EMAIL"}],"frequency":"WEEKLY","id":"12345","lastCheckedAt":"2021-03-16T01:01:41+00:00","threshold":"1","thresholdExceededAt":"2021-03-16T01:01:41+00:00","unit":"Gbps"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDatatrafficNotification(12345, 1)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Actions[0].Type, "EMAIL")
}

func TestServerDatatrafficNotificationUpdate(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/datatraffic/1")
		assert.Equal(t, req.Method, "PUT")
		// Send response to be tested
		rw.Write([]byte(`{"actions":[{"lastTriggeredAt":"2021-03-16T01:01:44+00:00","type":"EMAIL"}],"frequency":"WEEKLY","id":"12345","lastCheckedAt":"2021-03-16T01:01:41+00:00","threshold":"1","thresholdExceededAt":"2021-03-16T01:01:41+00:00","unit":"Gbps"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDatatrafficNotificationUpdate(12345, 1, &lwapi.DataTrafficNotificationRequest{
		Frequency: "DAILY",
		Threshold: "s",
		Unit:      "GB",
	})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Actions[0].Type, "EMAIL")
}

func TestServerDDoSNotification(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/ddos")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{
			"nulling": "ENABLED",
			"scrubbing": "DISABLED"
		  }`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDDoSNotification(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Nulling, "ENABLED")
}

func TestServerDDoSNotificationUpdate(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/notificationSettings/ddos")
		assert.Equal(t, req.Method, "PUT")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerDDoSNotificationUpdate(12345, &lwapi.DDoSStatus{})
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")
}

func TestServerPowerCycle(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/powerCycle")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerPowerCycle(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")
}

func TestServerPowerStatus(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/powerCycle")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{
			"ipmi": {
			  "status": "off"
			},
			"pdu": {
			  "status": "on"
			}
		  }`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerPowerStatus(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Ipmi.Status, "off")
}

func TestServerPowerOn(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/powerOn")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerPowerOn(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")
}

func TestServerPowerOff(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/servers/12345/powerOff")
		assert.Equal(t, req.Method, "POST")
		// Send response to be tested
		rw.Write([]byte(err))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ServerPowerOff(12345)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ErrorCode, "APP00800")
}

func TestOSes(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/operatingSystems")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":100,"offset":0,"totalCount":26},"operatingSystems":[{"id":"ALMALINUX_8_64BIT","name":"AlmaLinux 8 (x86_64)"},{"id":"CENTOS_7_64BIT","name":"CentOS 7 (x86_64)"},{"id":"ROCKY_LINUX_8_64BIT","name":"Rocky Linux 8 (x86_64)"},{"id":"DEBIAN_10_64BIT","name":"Debian 10 (amd64)"},{"id":"DEBIAN_9_64BIT","name":"Debian 9 (amd64)"},{"id":"UBUNTU_20_04_64BIT","name":"Ubuntu 20.04 LTS (Focal Fossa) (amd64)"},{"id":"UBUNTU_18_04_64BIT","name":"Ubuntu 18.04 LTS (Bionic Beaver) (amd64)"},{"id":"WINDOWS_SERVER_2019_STANDARD_64BIT","name":"Windows Server 2019 Standard (x64)"},{"id":"WINDOWS_SERVER_2019_DATACENTER_64BIT","name":"Windows Server 2019 Datacenter (x64)"},{"id":"WINDOWS_SERVER_2016_STANDARD_64BIT","name":"Windows Server 2016 Standard (x64)"},{"id":"WINDOWS_SERVER_2016_DATACENTER_64BIT","name":"Windows Server 2016 Datacenter (x64)"},{"id":"WINDOWS_SERVER_2012_R2_STANDARD_64BIT","name":"Windows Server 2012 R2 Standard (x64)"},{"id":"WINDOWS_SERVER_2012_R2_DATACENTER_64BIT","name":"Windows Server 2012 R2 Datacenter (x64)"},{"id":"WINDOWS_SERVER_2012_STANDARD_64BIT","name":"Windows Server 2012 Standard (x64) (Only available to existing license holders)"},{"id":"WINDOWS_SERVER_2012_DATACENTER_64BIT","name":"Windows Server 2012 Datacenter (x64) (Only available to existing license holders)"},{"id":"ESXI_7_0_64BIT_CSTM","name":"ESXi 7.0 Targeted (x86_64)"},{"id":"ESXI_7_0_64BIT","name":"ESXi 7.0 (x86_64)"},{"id":"ESXI_6_7_64BIT_CSTM","name":"ESXi 6.7 Targeted (x86_64)"},{"id":"ESXI_6_7_64BIT","name":"ESXi 6.7 (x86_64)"},{"id":"ESXI_6_5_64BIT_CSTM","name":"ESXi 6.5 Targeted (x86_64)"},{"id":"ESXI_6_5_64BIT","name":"ESXi 6.5 (x86_64)"},{"id":"ESXI_6_0_64BIT_CSTM","name":"ESXi 6.0 Targeted (x86_64)"},{"id":"ESXI_6_0_64BIT","name":"ESXi 6.0 (x86_64)"},{"id":"FREEBSD_12_64BIT","name":"FreeBSD 12.3 (amd64)"}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.OSes(v)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.OperatingSystems[0].ID, "ALMALINUX_8_64BIT")
}

func TestOS(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/operatingSystems/s?controlPanelId=s")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"architecture":"64bit","configurable":true,"defaults":{"device":"SATA_SAS","partitions":[{"bootable":true,"filesystem":"ext2","mountpoint":"\/boot","primary":true,"size":1024},{"filesystem":"swap","size":4096},{"filesystem":"ext4","mountpoint":"\/tmp","size":4096},{"filesystem":"ext4","mountpoint":"\/","primary":true,"size":"*"}]},"family":"ubuntu","features":["PARTITIONING","SW_RAID","TIMEZONE","HOSTNAME","SSH_KEYS","POST_INSTALL_SCRIPTS"],"id":"UBUNTU_20_04_64BIT","name":"Ubuntu 20.04 LTS (Focal Fossa) (amd64)","supportedBootDevices":["SATA_SAS","NVME"],"supportedFileSystems":["ext2","ext3","ext4","xfs","swap"],"type":"linux","version":"20.04"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.OS("s", "s")
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.Architecture, "64bit")
}

func TestControlPanels(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/controlPanels")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":10,"offset":0,"totalCount":2},"controlPanels":[{"id":"CPANEL_PREMIER_100","name":"cPanel Premier 100"},{"id":"CPANEL_PREMIER_150","name":"cPanel Premier 150"},{"id":"CPANEL_PREMIER_200","name":"cPanel Premier 200"},{"id":"CPANEL_PREMIER_250","name":"cPanel Premier 250"},{"id":"PLESK_DEDSER_WEB_ADMIN","name":"Plesk Web Admin 10 Domains"},{"id":"PLESK_DEDSER_WEB_PRO","name":"Plesk Web Pro 30 Domains"},{"id":"PLESK_DEDSER_WEB_HOST","name":"Plesk Web Host Unlimited Domains"},{"id":"VESTA","name":"Vesta CP"}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.ControlPanels(v)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.ControlPanels[0].ID, "CPANEL_PREMIER_100")
}

func TestRescueImages(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		rw.Header().Set("Content-Type", "application/json")
		assert.Equal(t, req.URL.String(), "/bareMetals/v2/rescueImages")
		assert.Equal(t, req.Method, "GET")
		// Send response to be tested
		rw.Write([]byte(`{"_metadata":{"limit":10,"offset":0,"totalCount":2},"rescueImages":[{"id":"GRML","name":"GRML Linux Rescue Image (amd64)"},{"id":"CENTOS_7","name":"CentOS 7 Linux Rescue Image (amd64)"},{"id":"FREEBSD","name":"FreeBSD Rescue Image (amd64)"}]}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := lwapi.New("testtoken").DedicatedServers()
	api.BaseURL = server.URL

	s, e := api.RescueImages(v)
	if e != nil {
		t.Error(e)
	}

	assert.Equal(t, s.RescueImages[0].ID, "GRML")
}
