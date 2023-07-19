### Terraform provider for reg.ru DNS records

Example usage:

```hcl
provider "regru" {
    api_username = "test@exmpl.com"
    api_password = "fo0b4r6az"
}


// will create A record ex.xmplar.com -> 1.1.1.1
resource "regru_dns_record" "test-a" {
    zone = "xmplar.com"
    name = "ex"
    type = "A"
    record = "1.1.1.1"
}

// will create CNAME record exit.xmplar.com -> exit.xmplar.com
resource "regru_dns_record" "test-cname" {
    zone = "xmplar.com"
    name = "exit"
    type = "CNAME"
    record = "ex.mplar.com"
}

// will create MX record xmplar.com -> 10 mx.yandex.net
resource "regru_dns_record" "test-mx" {
    zone = "xmplar.com"
    name = "@"
    type = "MX"
    record = "10 mx.yandex.net"
}

// will create TXT record _acme-challenge.xmplar.com -> apchIhba
resource "regru_dns_record" "test-txt" {
    zone = "xmplar.com"
    name = "_acme-challenge"
    type = "TXT"
    record = "apchIhba"
}
```