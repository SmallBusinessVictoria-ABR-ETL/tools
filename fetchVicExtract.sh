#!/bin/bash

export AWS_REGION=ap-southeast-2
export EXTRACT_DATE="`date +%Y%m%d`"
envsubst < VicExtract.batch | /usr/bin/sshpass -p `kms-decrypt AQICAHjAjhV7d3YGxLXMWTRObCPHtjQT0joQ4ZkhoypbVJ9fIQFefuYdq1x049a/iPESUlFKAAAAaDBmBgkqhkiG9w0BBwagWTBXAgEAMFIGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMyY2jWUZOVygGcstEAgEQgCVKSUZMYnfxdQem2CEpMOqKgs30fzgCMv4E3ZcYvffcY9Ze7lZH` sftp -c aes256-cbc -o StrictHostKeyChecking=no -oKexAlgorithms=+diffie-hellman-group-exchange-sha256 `kms-decrypt AQICAHjAjhV7d3YGxLXMWTRObCPHtjQT0joQ4ZkhoypbVJ9fIQHuJuUm8IBYOZ3242iXQRjXAAAAezB5BgkqhkiG9w0BBwagbDBqAgEAMGUGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMxAxT8oW24rAJNbtiAgEQgDjwifBrEL3vHSY3LF9bs1fQaEbHk/tOoAkbTWpdg03NKJGdsW628pdFhH7AwtWxKmNo+njLlIZ+5w==`@180.149.195.60
vi .ssh