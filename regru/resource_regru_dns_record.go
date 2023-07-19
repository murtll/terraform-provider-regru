package regru

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRegruDNSRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceRegruDNSRecordCreate,
		Read:   resourceRegruDNSRecordRead,
		Delete: resourceRegruDNSRecordDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"record": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"zone": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceRegruDNSRecordCreate(d *schema.ResourceData, m interface{}) error {
	record_type := d.Get("type").(string)
	record_name := d.Get("name").(string)
	value := d.Get("record").(string)
	zone := d.Get("zone").(string)

	c := m.(*Client)
	baseRequest := CreateRecordRequest{
		Username:          c.username,
		Password:          c.password,
		Domains:           []Domain{{DName: zone}},
		SubDomain:         record_name,
		OutputContentType: "plain",
	}

	var request interface{}

	switch strings.ToUpper(record_type) {
	case "A":
		request = CreateARecordRequest{
			CreateRecordRequest: baseRequest,
			IPAddr:              value,
		}
	case "CNAME":
		request = CreateCnameRecordRequest{
			CreateRecordRequest: baseRequest,
			CanonicalName:       value,
		}
	case "MX":
		fields := strings.Fields(value)

		request = CreateMxRecordRequest{
			CreateRecordRequest: baseRequest,
			MailServer:          fields[1],
			Priority:            fields[0],
		}
	case "TXT":
		request = CreateTxtRecordRequest{
			CreateRecordRequest: baseRequest,
			Text:                value,
		}
	default:
		return fmt.Errorf("invalid record type '%s'", record_type)
	}

	action := fmt.Sprintf("add_%s", strings.ToLower(record_type))
	if strings.ToLower(record_type) == "a" {
		action = "add_alias"
	}

	resp, err := c.doRequest(request, "zone", action)
	if err != nil {
		return err
	}
	if resp.HasError() != nil {
		return resp.HasError()
	}
	d.SetId(strings.Join([]string{record_name, zone}, "."))
	return nil
}

func resourceRegruDNSRecordRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRegruDNSRecordDelete(d *schema.ResourceData, m interface{}) error {
	record_type := d.Get("type").(string)
	record_name := d.Get("name").(string)
	value := d.Get("record").(string)
	zone := d.Get("zone").(string)

	c := m.(*Client)

	request := DeleteRecordRequest{
		Username:          c.username,
		Password:          c.password,
		Domains:           []Domain{{DName: zone}},
		SubDomain:         record_name,
		Content:           value,
		RecordType:        strings.ToUpper(record_type),
		OutputContentType: "plain",
	}

	resp, err := c.doRequest(request, "zone", "remove_record")
	if err != nil {
		return err
	}

	return resp.HasError()
}
