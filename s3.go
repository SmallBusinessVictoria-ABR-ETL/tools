package sbv_abr_etl

//func main() {
//	sess := session.Must(session.NewSession())
//
//	// Create a downloader with the session and default options
//	downloader := s3manager.NewDownloader(sess)
//
//	// Create a file to write the S3 Object contents to.
//	f, err := os.Create(filename)
//	if err != nil {
//		return fmt.Errorf("failed to create file %q, %v", filename, err)
//	}
//
//	// Write the contents of S3 Object to the file
//	n, err := downloader.Download(f, &s3.GetObjectInput{
//		Bucket: aws.String(os.Args[1]),
//		Key:    aws.String(os.Args[2]),
//	})
//	if err != nil {
//		return fmt.Errorf("failed to download file, %v", err)
//	}
//	fmt.Printf("file downloaded, %d bytes\n", n)
//}
