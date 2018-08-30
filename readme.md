
## Dev Requirements

 * IAM Role with access to:
   * KMS:Encrypt (credentials)
 * 

## Server Requirements

 * git
 * go
 * Ubuntu Ec2 Instance
   * Elastic IP whitelists with ATO
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

## Install/Updatep

```bash
go get -u github.com/SmallBusinessVictoria-ABR-ETL/tools/...
```

## Encrypt username and passwords

```bash
aws --region ap-southeast-2 kms list-keys
aws --region ap-southeast-2 kms encrypt --key-id <key-id> --plaintext <username or password>
```

## Secure SFTP Batch Get

Use AWS KMS encrypt to encrypt the username and password environments  

```bash
export SFTP_USER_ENC=AQICAHjAjhV7d3YGxLXMWTRObCPHtjQT0joQ4ZkhoypbVJ9fIQHuJuUm8IBYOZ3242iXQRjXAAAAezB5BgkqhkiG9w0BBwagbDBqAgEAMGUGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMxAxT8oW24rAJNbtiAgEQgDjwifBrEL3vHSY3LF9bs1fQaEbHk/tOoAkbTWpdg03NKJGdsW628pdFhH7AwtWxKmNo+njLlIZ+5w==
export SFTP_PASS_ENC=AQICAHjAjhV7d3YGxLXMWTRObCPHtjQT0joQ4ZkhoypbVJ9fIQFefuYdq1x049a/iPESUlFKAAAAaDBmBgkqhkiG9w0BBwagWTBXAgEAMFIGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMyY2jWUZOVygGcstEAgEQgCVKSUZMYnfxdQem2CEpMOqKgs30fzgCMv4E3ZcYvffcY9Ze7lZH
export SFTP_HOST=180.149.195.60
export SFTP_PORT=22
export AWS_REGION=ap-southeast-2

go run sftp-get/app.go "AllStates_ABR Data/Sent/VIC_ABR Extract.zip" "`date +%Y%m%d`-VIC_ABR_Extract.zip"
```


```bash

export AWS_REGION=ap-southeast-2 
export EXTRACT_DATE="`date +%Y%m%d`" 
envsubst < VicExtract.batch | sshpass -p `go run kms-decrypt/app.go AQICAHjAjhV7d3YGxLXMWTRObCPHtjQT0joQ4ZkhoypbVJ9fIQFefuYdq1x049a/iPESUlFKAAAAaDBmBgkqhkiG9w0BBwagWTBXAgEAMFIGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMyY2jWUZOVygGcstEAgEQgCVKSUZMYnfxdQem2CEpMOqKgs30fzgCMv4E3ZcYvffcY9Ze7lZH` sftp -c aes256-cbc -o StrictHostKeyChecking=no -oKexAlgorithms=+diffie-hellman-group-exchange-sha256 `go run kms-decrypt/app.go AQICAHjAjhV7d3YGxLXMWTRObCPHtjQT0joQ4ZkhoypbVJ9fIQHuJuUm8IBYOZ3242iXQRjXAAAAezB5BgkqhkiG9w0BBwagbDBqAgEAMGUGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMxAxT8oW24rAJNbtiAgEQgDjwifBrEL3vHSY3LF9bs1fQaEbHk/tOoAkbTWpdg03NKJGdsW628pdFhH7AwtWxKmNo+njLlIZ+5w==`@180.149.195.60

```


## Diff (Updated + New)

```bash
aws s3 cp s3://sbv-abr-etl/FACT/AGENCY/date=2018-08-23/state=full/full.txt.gz ./previous.txt.gz 
diff-abr previous.txt VIC*_ABR_Agency_Data.txt update.txt new.txt
gzip update.txt
gzip new.txt
aws s3 cp update.txt.gz s3://sbv-abr-etl/FACT/AGENCY/date=`date +%Y-%m-%d`/state=update/update.txt.gz
aws s3 cp new.txt.gz s3://sbv-abr-etl/FACT/AGENCY/date=`date +%Y-%m-%d`/state=new/new.txt.gz
aws s3 cp VIC*_ABR_Agency_Data.txt s3://sbv-abr-etl/FACT/AGENCY/date=`date +%Y-%m-%d`/state=full/full.txt.gz
```