package daemon

// import (
// 	"log"
// 	"os"

// 	checkup "github.com/AdhityaRamadhanus/gopatrol"
// 	"github.com/AdhityaRamadhanus/gopatrol/config"
// 	"github.com/pkg/errors"
// 	"github.com/urfave/cli"
// )

// // Helper
// func setupFSDaemon(c *cli.Context) (err error) {
// 	defer func() {
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}()
// 	// Create directory for logs
// 	// ignore the error
// 	log.Println("Setting up directory")
// 	os.MkdirAll("./checkup_config/logs", 0777)
// 	// setup checkup.json
// 	log.Println("Creating checkup.json")
// 	checkup := checkup.Checkup{
// 		Checkers: []checkup.Checker{},
// 		Storage: checkup.FS{
// 			Dir: c.String("log"),
// 		},
// 	}
// 	jsonBytes, err := checkup.MarshalJSON()
// 	if err != nil {
// 		return err
// 	}
// 	file, err := os.OpenFile(config.DefaultCheckupJSON, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
// 	defer file.Close()
// 	if err != nil {
// 		return errors.Wrap(err, "Failed opening checkup.json file")
// 	}
// 	if _, err = file.Write(jsonBytes); err != nil {
// 		return errors.Wrap(err, "Failed to write checkup.json")
// 	}
// 	log.Println("Success!")
// 	return nil
// }

// // Helper
// func setupS3Daemon(c *cli.Context) (err error) {
// 	defer func() {
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}()
// 	s3config := config.S3Config{
// 		AccessKeyID:     c.String("accesskeyid"),
// 		SecretAccessKey: c.String("secretaccesskey"),
// 		Region:          c.String("region"),
// 		Bucket:          c.String("bucket"),
// 	}
// 	var s3Error error
// 	if len(s3config.AccessKeyID) == 0 {
// 		log.Println("Please Provide AccessKeyID")
// 		s3Error = errors.New("S3 Config Error")
// 	}

// 	if len(s3config.SecretAccessKey) == 0 {
// 		log.Println("Please Provide Secret AccessKey")
// 		s3Error = errors.New("S3 Config Error")
// 	}

// 	if len(s3config.Region) == 0 {
// 		log.Println("Please Provide S3 Region")
// 		s3Error = errors.New("S3 Config Error")
// 	}

// 	if len(s3config.Bucket) == 0 {
// 		log.Println("Please Provide S3 Bucket")
// 		s3Error = errors.New("S3 Config Error")
// 	}

// 	if s3Error != nil {
// 		log.Println(s3Error)
// 		return s3Error
// 	}

// 	// Create directory for logs
// 	// ignore the error
// 	log.Println("Setting up directory")
// 	os.MkdirAll("./checkup_config/logs", 0777)
// 	// setup checkup.json
// 	log.Println("Creating checkup.json")
// 	checkup := checkup.Checkup{
// 		Checkers: []checkup.Checker{},
// 		Storage: checkup.S3{
// 			AccessKeyID:     s3config.AccessKeyID,
// 			SecretAccessKey: s3config.SecretAccessKey,
// 			Region:          s3config.Region,
// 			Bucket:          s3config.Bucket,
// 		},
// 	}
// 	jsonBytes, err := checkup.MarshalJSON()
// 	if err != nil {
// 		return err
// 	}
// 	file, err := os.OpenFile(config.DefaultCheckupJSON, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
// 	defer file.Close()
// 	if err != nil {
// 		return errors.Wrap(err, "Failed opening checkup.json file")
// 	}
// 	if _, err = file.Write(jsonBytes); err != nil {
// 		return errors.Wrap(err, "Failed to write checkup.json")
// 	}
// 	log.Println("Success!")
// 	return nil
// }
