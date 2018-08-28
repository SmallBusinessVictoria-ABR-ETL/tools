

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

## Install

```bash
go get github.com/SmallBusinessVictoria-ABR-ETL/tools
```

## Secure SFTP Batch Get

Use AWS KMS encrypt to encrypt the username and password environments  

```bash
export SFTP_USER_ENC=AQICAHitM6SAPtEvY+DLu+YrFfAk4jBguxTikUS6Sqpc7bzoUAHZvf0S6o0wIh8zPZ4ZbQ4FAAAAezB5BgkqhkiG9w0BBwagbDBqAgEAMGUGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMbti5hLZMSNWtEUemAgEQgDhHeKM6LwPyxkg9ryeTi/kMYBMC9OIYCJSPRajUDtEyT5UTWVUB4ln4qkzt8trMjPexfGBlCIvXwQ==
export SFTP_PASS_ENC=AQICAHitM6SAPtEvY+DLu+YrFfAk4jBguxTikUS6Sqpc7bzoUAEEAWka3vXRzDgxZzCSnR5BAAAAaDBmBgkqhkiG9w0BBwagWTBXAgEAMFIGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQM5a4hhxWXlLRh1xSIAgEQgCX/TXvTaYwHL8yIGSbv1rhUOLzNQ3BHW0d/95d5vfZIpe3CH8y/
export SFTP_HOST=180.149.195.60
export SFTP_PORT=22
export AWS_REGION_KMS=us-east-1

sftp-get "AllStates_ABR Data/Sent/VIC_ABR Extract.zip" "`date +%Y%m%d`-VIC_ABR_Extract.zip"
```