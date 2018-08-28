package tools

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

func SFTPGet(file, localFileName string) {

	username, err := Decrypt(os.Getenv("SFTP_USER_ENC"))
	if err != nil {
		log.Fatal("Failed to decrypt SFTP_USER_ENC")
	}
	password, err := Decrypt(os.Getenv("SFTP_PASS_ENC"))
	if err != nil {
		log.Fatal("Failed to decrypt SFTP_PASS_ENC")
	}

	var sshClient *ssh.Client

	config := ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%d", os.Getenv("SFTP_HOST"), os.Getenv("SFTP_PORT"))
	sshClient.Conn, err = ssh.Dial("tcp", addr, &config)
	if err != nil {
		log.Fatal("Failed to connect via sfpt")
	}

	// open an SFTP session over an existing ssh connection.
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	// check it's there
	fi, err := sftpClient.Lstat(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fi)

	remoteFile, err := sftpClient.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	localFile, err := os.Create(localFileName)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(localFile, remoteFile)
	remoteFile.Close()
	localFile.Close()
}
