version := 0.1.4
path := $$HOME/.terraform.d/plugins/github.com/murtll/regru/${version}/linux_amd64

build:
	mkdir -p ${path}
	go build -o ${path}/terraform-provider-regru_${version}