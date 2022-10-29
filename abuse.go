package lwapi

import (
	"encoding/json"
	"fmt"
)

// List reports
func (a *AbuseApi) Reports(queryParams map[string]interface{}) (*Reports, error) {
	query := MakeQuery(queryParams)
	bodyResp, err := a.NewRequest(NilPayload, "/reports"+query, "GET")

	var r *Reports
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Inspect a report
func (a *AbuseApi) Report(reportID string) (*Report, error) {
	uri := fmt.Sprintf("/reports/%s", reportID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	var r *Report
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Inspect a report messages
func (a *AbuseApi) Messages(reportID string, queryParams map[string]interface{}) (*Messages, error) {
	query := MakeQuery(queryParams)
	uri := fmt.Sprintf("/reports/%s/messages", reportID)
	payload, _ := json.Marshal(query)
	bodyResp, err := a.NewRequest(payload, uri, "GET")

	var r *Messages
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Create new message
func (a *AbuseApi) NewMessage(reportID string, params *Message) ([]string, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/reports/%s/messages", reportID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r []string
	json.Unmarshal(bodyResp, &r)
	return r, err
}

// Inspect a report attachments
// The compliance team sometimes add an attachment to a message.
// You can use this endpoint to get the attachment. The content-type of the response depends on the content of the attachment.
// https://api.leaseweb.com/abuse/v1/reports/{reportId}/messageAttachments/{fileId}
func (a *AbuseApi) MessageAttachments(reportID string, fileID string) ([]byte, error) {
	uri := fmt.Sprintf("/reports/%s/messageAttachments/%s", reportID, fileID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	return bodyResp, err
}

// Inspect a report attachment
// Use this endpoint to get an attachment which was created with the abuse report.
// The content-type of the response depends on the content of the attachment.
// https://api.leaseweb.com/abuse/v1/reports/{reportId}/reportAttachments/{fileId}
func (a *AbuseApi) ReportAttachments(reportID string, fileID string) ([]byte, error) {
	uri := fmt.Sprintf("/reports/%s/reportAttachments/%s", reportID, fileID)
	bodyResp, err := a.NewRequest(NilPayload, uri, "GET")

	return bodyResp, err
}

// Resolve report
func (a *AbuseApi) Resolve(reportID string, params *Resolutions) (*Error, error) {
	payload, _ := json.Marshal(params)
	uri := fmt.Sprintf("/reports/%s/resolve", reportID)
	bodyResp, err := a.NewRequest(payload, uri, "POST")

	var r *Error
	json.Unmarshal(bodyResp, &r)
	return r, err
}
