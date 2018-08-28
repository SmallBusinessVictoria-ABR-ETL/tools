

## Server Requirements

 * git
 * go
 * EC2 Role with access to:
   * KMS:Decrypt (credentials)
   * KMS:Encrypt (data)
   * KMS:Decrypt (data)
   * S3:PutObject      
   * S3:ListBucket      
   * S3:GetObject      

### GO bin config

```bash
# .bash_profile/.profile
# ...
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

## Install/Update

```bash
go get -u github.com/SmallBusinessVictoria-ABR-ETL/tools
```

## Secure SFTP Batch Get

Use AWS KMS encrypt to encrypt the username and password environments  

```bash
export SFTP_USER_ENC=AQICAHjAjhV7d3YGxLXMWTRObCPHtjQT0joQ4ZkhoypbVJ9fIQHuJuUm8IBYOZ3242iXQRjXAAAAezB5BgkqhkiG9w0BBwagbDBqAgEAMGUGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMxAxT8oW24rAJNbtiAgEQgDjwifBrEL3vHSY3LF9bs1fQaEbHk/tOoAkbTWpdg03NKJGdsW628pdFhH7AwtWxKmNo+njLlIZ+5w==
export SFTP_PASS_ENC=AQICAHjAjhV7d3YGxLXMWTRObCPHtjQT0joQ4ZkhoypbVJ9fIQFefuYdq1x049a/iPESUlFKAAAAaDBmBgkqhkiG9w0BBwagWTBXAgEAMFIGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMyY2jWUZOVygGcstEAgEQgCVKSUZMYnfxdQem2CEpMOqKgs30fzgCMv4E3ZcYvffcY9Ze7lZH
export SFTP_HOST=180.149.195.60
export SFTP_PORT=22
export AWS_REGION=ap-southeast-2

sftp-get "AllStates_ABR Data/Sent/VIC_ABR Extract.zip" "`date +%Y%m%d`-VIC_ABR_Extract.zip"
```


