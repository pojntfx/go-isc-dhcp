package svc

//go:generate mkdir -p ../proto/generated
//go:generate sh -c "protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"
//go:generate rm -rf dhcp statikDhcpd statikDhclient
//go:generate git clone https://gitlab.isc.org/isc-projects/dhcp.git
//go:generate sh -c "cd dhcp; ./configure; make; cd .."
//go:generate statik -src dhcp/server -include dhcpd -p statikDhcpd
//go:generate statik -src dhcp/client -include dhclient -p statikDhclient
