# **UTIL Send-GMail**

Utility send-gmail - the utility is designed to send mail to gmail.

## **Installation**
- install [golang](https://go.dev/) 1.16+
- go get github.com/xxandev/util-send-gmail
- cd ..../util-send-gmail
- make [ build | arm6 | arm7 | arm8 | linux64 | linux32 | win64 | win32 | win64i | win32i ] or go build .

## **Run**
- create config file
```yaml
gmail:
  login: you_mail@gmail.com
  password: you_password
```
- run 
```bash
.../send-gmail -to 'mail_1@gmail.com,mail_2@gmail.com' -subj hello -body 'hello world'
```
OR
```bash
.../send-gmail \
    -login you_mail@gmail.com \
    -pass you_password \
    -to 'mail_1@gmail.com,mail_2@gmail.com' \
    -subj hello \
    -body 'hello world'
```