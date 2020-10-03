Generate:
	openssl genrsa -out https-server.key 2048
	openssl ecparam -genkey -name secp384r1 -out https-server.key
	openssl req -new -x509 -sha256 -key https-server.key -out https-server.crt -days 3650

OtherGenerate: # from https://medium.com/@tbusser/creating-a-browser-trusted-self-signed-ssl-certificate-2709ce43fd15
	# root certificate
	openssl genrsa -des3 -out rootCA.key 2048
	openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 365 -out rootCA.pem
	# ssl certificate
