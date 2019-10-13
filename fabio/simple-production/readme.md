# Fabio - Simple Production Setup

In this test, we'll setup Fabio to listen for HTTP & HTTPS on port 80 & 443.

But first, let's prepare a domain name for testing: `www.test.yk`.
1. Edit `/etc/hosts` to include that domain. The line would look like this: ```127.0.0.1 www.test.yk```
2. Create a certificate with the help of [minica](https://github.com/jsha/minica). Run `minica --domains www.test.yk`

Let's run the services from `/basic` test and start the web services, excluding consul & fabio:
```docker-compose up -d web1 web2 web3```

In this test, we'll use [Consul binary](https://www.consul.io/downloads.html) & [Fabio binary](https://github.com/fabiolb/fabio/releases) without docker.

Start consul like this (v1.6.1):

```consul agent -data-dir=/tmp/consul -server -ui -bootstrap-expect=1```

Then run fabio with following command:

```fabio -registry.consul.addr=127.0.0.1:8500 -registry.consul.register.addr=127.0.0.1:9998 -ui.addr=0.0.0.0:9998 -proxy.addr=":80,:443;cs=mycert" -proxy.cs="cs=mycert;type=consul;cert=http://127.0.0.1:8500/v1/kv/certificates"```

What are those parameters?
1. Consul
   - data-dir=/tmp/consul   Specify data directory
   - server                 Run consul as server
   - ui                     Enable UI access (on port :8500)
   - bootstrap-expect=1     Run consul cluster with 1 quorum (it's used for single server setup)

2. Fabio
   - proxy.addr=":80,:443;cs=mycert"                 Tell Fabio to listen on :80 and :443 using certificate provider `mycert`
   - proxy.cs="cs=mycert;type=consul;cert=http://127.0.0.1:8500/v1/kv/certificates"
     Specify certificate provider with name `mycert` using Consul at given address.
   - registry.consul.addr=127.0.0.1:8500             Specify location of Consul (default is 127.0.0.1:8500). Can be ommited if you have Consul agent on 127.0.0.1
   - registry.consul.register.addr=127.0.0.1:9998    Specify the location of Fabio so that consul can find it.
   - ui.addr=0.0.0.0:9998                            Specify Fabio UI address.
   - More config can be found at [https://fabiolb.net/ref/](https://fabiolb.net/ref/)

You could also put all that to config file (eg: `fabio.cfg`):
```
proxy.addr = :80,:443;cs=mycert
proxy.cs = cs=mycert;type=consul;cert=http://127.0.0.1:8500/v1/kv/certificates
registry.consul.addr = 127.0.0.1:8500
registry.consul.register.addr = 127.0.0.1:9998
ui.addr = 0.0.0.0:9998
```
Run fabio with `fabio -cfg fabio.cfg`

## Testing:

If you add this line to `/fabio/config` key/value store in Consul:
```route add web1 / http://127.0.0.1:9001```
you'll be able to browse `http://www.test.yk` and see a result. But it will complain if you try to open: `https://www.test.yk`.
HTTPS require that you have a certificate set up.

Copy `cert.pem` and `key.pem` into Consul kv, let's name it `/certificates/www.test.yk.pem` (certificate must end with `.pem`). The content would look like this:
```
-----BEGIN CERTIFICATE-----
MIIDMjCCAhqgAwIBAgIIGVAcUGQhkp0wDQYJKoZIhvcNAQELBQAwIDEeMBwGA1UE
AxMVbWluaWNhIHJvb3QgY2EgMWRhY2FkMCAXDTE5MTAxMzA1NDg0NVoYDzIxMDkx
MDEzMDU0ODQ1WjAWMRQwEgYDVQQDEwt3d3cudGVzdC55azCCASIwDQYJKoZIhvcN
AQEBBQADggEPADCCAQoCggEBAM31f//IVtuCiyIed0JCf245ld646gtyJ0Q0BIlc
d5+fl6Bw1syoZm7PMStShsYXd4PSrRDnUOtvVKHaZcXOCtodWhYNlgZS6nq3Yjev
9U2PoxJalnyslgBIhND/kZA8+ou6zbVRtsmRtbrys3+bbS0TIcA0VNa7e+MgBhHT
tclLu21iJCLkRmIpkKXgVVW2ozG4fz/kwHCTse0LWyXaR/+SeIRhv342e36BBOYv
uG+WcEX/Ci0X9kKrjiGxi7GsecYSEedrGcoZMJ7aQ0nn+2PeknGmQ795byprCo91
7FYqqvemDjAUAra5BODuW8y4Z2HmFauskVdkjPXxD6poElMCAwEAAaN4MHYwDgYD
VR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNV
HRMBAf8EAjAAMB8GA1UdIwQYMBaAFE8GRlcGLfac1cbMEAK012JHl0QBMBYGA1Ud
EQQPMA2CC3d3dy50ZXN0LnlrMA0GCSqGSIb3DQEBCwUAA4IBAQACxDciPEImzveW
158Ej0ZGJRJibCVjkG7vm71FwG4yS6J+prlY3P0TxoGLhKjKuGo+Up97rw2WbOKr
z4Ozy4i263HN6kkmc4e1L7hoPiVI89qz1k3sxN93xIrEwGiLcXS9fGW5CY0P/iKr
sEoYZQs8dkS+k62ID0NTTmt1/Elxw30x3XPlcMQl4h4v4ROVsIBvnX42t0X1yW0V
MbE1YfZk8y9a2SuEsCRpWvaK4fKPFJstPN+pOKkucS2YekRugCxRZRQp8AH28TWa
ewjHQ9klaQjsFfrVrI6NnSA4MkuBUXBqUNL5V/QBCPOf6SA12MXxeVQatoI0H7eU
sNJHCcPq
-----END CERTIFICATE-----
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzfV//8hW24KLIh53QkJ/bjmV3rjqC3InRDQEiVx3n5+XoHDW
zKhmbs8xK1KGxhd3g9KtEOdQ629Uodplxc4K2h1aFg2WBlLqerdiN6/1TY+jElqW
fKyWAEiE0P+RkDz6i7rNtVG2yZG1uvKzf5ttLRMhwDRU1rt74yAGEdO1yUu7bWIk
IuRGYimQpeBVVbajMbh/P+TAcJOx7QtbJdpH/5J4hGG/fjZ7foEE5i+4b5ZwRf8K
LRf2QquOIbGLsax5xhIR52sZyhkwntpDSef7Y96ScaZDv3lvKmsKj3XsViqq96YO
MBQCtrkE4O5bzLhnYeYVq6yRV2SM9fEPqmgSUwIDAQABAoIBABYfySAw5SOvYkLI
AwebHRU6Gl9SfdG88XriG5ql1DPRcPhLJsfVTyuYFjARLWkaHDyM5QprzsV7sBuy
2jvlZkUH4iV8UCFdxtJn4KUawb8TLvFJyeCaqYJeR+YrjfdwlXltF6vim0AL5vmu
GJocmB2/cGC1PmfSu/wp9Hz1wRFeX7GU4PpwQnKaKVmhfzOKpt8YoKTIhfQyZZ+J
oDUirJ7GKINIBHeglEppHkbbHXNopXvBl713uF7Yhy7k6ESMCY3cdQ+fnwnhJ5Pb
RP600X4ka8GeR0lzFAiX1FixUD/hul4uZORVZMAD5hYMZxlZHdFgClcO4bvku/U1
Cw3seAECgYEA0BqqIjyd9V5vtfm3/SkaY0BbvxjBd5IKSQddsgOye/F9Xw+1Hp51
mSdnfk+OsoYT349nlnuk1RfbgIYwgx7Q4qlAP/F8A9zwOkRodXYiA/13KmpIMDIC
lOCogAa3+WXbbqbS6rjM83qDYCIgTRdE3QQoiF8SKPa5kwXLpPwaMVsCgYEA/Vxx
g6t3TWtpteFKWbyK/5ad7HcYa+XL9jVqLdVAK2unc+BKmee3JESTZEMWZM1cbHr9
JJmdY9ZqmPlRpVIekYQqJneCVZS3dGz1jgvaUqNHoNPNdi5AQNh0sFVWD2eCpQva
AjP+lXBCOL9L1XtC/S9F2mxEj4ymnjsqD88xvGkCgYEAyK1N7yIBOMJOe29J64kd
dyRy1L6sof9kh7PguG80SK1BNtBQ2iv4Py5ucLGLa8A7ndQOEmE9PHh7JV9BnM+0
oz6PRJo7+wWtaqLZEJxQhQSBS5ed8UvojWRvWLYh5xBAIF4i+lIm4Yv88FE4UN7l
ezQtWgRD4Ni7b3mhPYIWSA0CgYB70lsTy4/hwVYHcpRgqNmRse16bHX9/W+h41cC
EU9sKQ/MfNhYwTrrOayC+pqOJyM9TRo8cerOqTKtkmOJlUmlOl9TL2L+KlFCUCHu
CvLnIi9WdUzbrhzu1BqrNvl5S9A4k0M3gmuwYw2qKCuKqNQDYsAT0IftVAL2H9od
odgfyQKBgFVL9Gcw+dpAWloCnsRB9fteKlUinm0q4hcOE11v3trL+Y2V8RoXJGbv
vxZXIVmH2iMVNEOewuKei1CyHdaAyuLCsPphEf1TKvxcP55eWkGr3ya5DX3K59B8
Lf6kdwaNtsNmHFGLJKRL8iXR3xyLhmbFAPZ0MMcSQFVQi7lX+Lbk
-----END RSA PRIVATE KEY-----
```

Now you'll be able to open `https://www.test.yk`. Even though your certificate will not be trusted by browser, the test shows the convenience of updating SSL Certificate in Fabio using Consul.

If you want more realism for trusted local certificate, checkout [mkcert](https://github.com/FiloSottile/mkcert).
First run `mkcert -install`. Next run `mkcert "*.test.yk"`, it will generate two files `./_wildcard.test.yk.pem` and `./_wildcard.test.yk-key.pem`.
You can use `cat ./_wildcard.test.yk.pem ./_wildcard.test.yk-key.pem` and copy the result to Consul key/value store, replacing the previous certificate.

To make sure all request are using HTTPS, add these line to Fabio config in Consul key/value store (`/fabio/config`).
```
route add http-redirect www.test.yk:80 https://www.test.yk/ opts "redirect=301"
```