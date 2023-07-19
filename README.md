### Terraform provider for reg.ru DNS records

Example usage:

```hcl
provider "regru" {
    api_username = "test@exmpl.com"
    api_password = "fo0b4r6az"
}

resource "regru_dns_record" "test" {
    zone = "xmplar.com"
    name = "ex"
    type = "A"
    record = "1.1.1.1"
}
```