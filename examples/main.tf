terraform {
    required_providers {
        regru = {
            version = "~> 0.1.0"
            source  = "github.com/murtll/regru"
        }
    }
}

provider "regru" {}

resource "regru_dns_record" "test" {
    zone = "esskeetiter.ru"
    name = "test"
    type = "A"
    record = "1.1.1.1"
}