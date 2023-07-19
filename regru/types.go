package regru

import "fmt"

const successResult = "success"

type APIResponse struct {
	Result string `json:"result"`

	Answer *Answer `json:"answer,omitempty"`

	ErrorCode string `json:"error_code,omitempty"`
	ErrorText string `json:"error_text,omitempty"`
}

func (a APIResponse) Error() string {
	return fmt.Sprintf("API %s: %s: %s", a.Result, a.ErrorCode, a.ErrorText)
}

func (a APIResponse) HasError() error {
	if a.Result != successResult {
		return a
	}

	if a.Answer != nil {
		for _, domResp := range a.Answer.Domains {
			if domResp.Result != successResult {
				return domResp
			}
		}
	}

	return nil
}

type Answer struct {
	Domains []DomainResponse `json:"domains,omitempty"`
}

type DomainResponse struct {
	Result string `json:"result"`

	DName     string `json:"dname"`

	ErrorCode string `json:"error_code,omitempty"`
	ErrorText string `json:"error_text,omitempty"`
}

func (d DomainResponse) Error() string {
	return fmt.Sprintf("API %s: %s: %s", d.Result, d.ErrorCode, d.ErrorText)
}

type CreateRecordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`

	Domains           []Domain `json:"domains,omitempty"`
	SubDomain         string   `json:"subdomain,omitempty"`
	OutputContentType string   `json:"output_content_type,omitempty"`
}

type CreateARecordRequest struct {
	CreateRecordRequest
	IPAddr string `json:"ipaddr,omitempty"`
}

type CreateCnameRecordRequest struct {
	CreateRecordRequest
	CanonicalName string `json:"canonical_name,omitempty"`
}

type CreateTxtRecordRequest struct {
	CreateRecordRequest
	Text string `json:"text,omitempty"`
}

type CreateMxRecordRequest struct {
	CreateRecordRequest
	MailServer string `json:"mail_server,omitempty"`
	Priority   string `json:"priority,omitempty"`
}

type DeleteRecordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`

	Domains           []Domain `json:"domains,omitempty"`
	SubDomain         string   `json:"subdomain,omitempty"`
	Content           string   `json:"content,omitempty"`
	RecordType        string   `json:"record_type,omitempty"`
	OutputContentType string   `json:"output_content_type,omitempty"`
}

type Domain struct {
	DName string `json:"dname"`
}
