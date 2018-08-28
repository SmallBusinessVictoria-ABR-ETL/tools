FROM ubuntu:18.04
RUN apt update && apt install -y sshpass openssh
COPY password-sftp.sh /password-sftp.sh
ENTRYPOINT password-sftp.sh