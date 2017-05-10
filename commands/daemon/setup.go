package daemon

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path"

	"github.com/AdhityaRamadhanus/gopatrol/config"
	"github.com/AdhityaRamadhanus/gopatrol/templates"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"
	"github.com/urfave/cli"
)

func SetupFS(c *cli.Context) (err error) {
	// defer func() {
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	// // Create directory for logs
	// // ignore the error
	// setupDir()

	// // setup checkup.json
	// log.Println("Creating checkup.json")
	// checkup := checkup.Checkup{
	// 	Checkers: []checkup.Checker{},
	// 	Storage: checkup.FS{
	// 		Dir: path.Join(config.DefaultStatuspageDir, "logs"),
	// 	},
	// }
	// jsonBytes, err := checkup.MarshalJSON()
	// if err != nil {
	// 	return err
	// }
	// file, err := os.OpenFile(path.Join(config.DefaultCheckupDir, "checkup.json"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	// defer file.Close()
	// if err != nil {
	// 	return errors.Wrap(err, "Failed opening checkup.json file")
	// }
	// if _, err = file.Write(jsonBytes); err != nil {
	// 	return errors.Wrap(err, "Failed to write checkup.json")
	// }

	// // setup Caddyfile
	// // Read the template
	// log.Println("Creating Caddyfile")
	// caddyTemplate, err := template.New("caddyfile").Parse(templates.CaddyFile)
	// if err != nil {
	// 	return errors.Wrap(err, "Error parsing caddy template config")
	// }
	// caddyFile, err := os.OpenFile(path.Join(config.DefaultCaddyDir, "Caddyfile"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0660)
	// defer caddyFile.Close()
	// if err != nil {
	// 	return errors.Wrap(err, "Error opening Caddyfile")
	// }
	// // Execute the template
	// if err := caddyTemplate.Execute(caddyFile, struct{ URL string }{URL: c.String("url")}); err != nil {
	// 	return errors.Wrap(err, "Failed writing to Caddyfile")
	// }

	// // setup config status page
	// log.Println("Creating index.html for status page")

	// dstIndexHTML, err := os.OpenFile(path.Join(config.DefaultStatuspageDir, "index.html"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	// if _, err = dstIndexHTML.Write(templates.IndexHTML); err != nil {
	// 	return errors.Wrap(err, "Failed to write index.html")
	// }

	// log.Println("Creating minified app.js for status page")

	// // Create app.js
	// if err := buildAppMinJs(path.Join(config.DefaultStatuspageDir, "js/app.min.js"), "fs", nil); err != nil {
	// 	return errors.Wrap(err, "failed to create app.min.js")
	// }

	// log.Println("Creating minified style.css for status page")

	// // Create app.js
	// if err := buildStyleCSS(path.Join(config.DefaultStatuspageDir, "css/style.min.css")); err != nil {
	// 	return errors.Wrap(err, "failed to create style.min.css")
	// }

	// setupImage()

	// log.Println("Success!")
	return nil
}

func SetupS3(c *cli.Context) (err error) {
	// defer func() {
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }()

	// s3config := config.S3Config{
	// 	AccessKeyID:     c.String("accesskeyid"),
	// 	SecretAccessKey: c.String("secretaccesskey"),
	// 	Region:          c.String("region"),
	// 	Bucket:          c.String("bucket"),
	// }
	// var s3Error error
	// if len(s3config.AccessKeyID) == 0 {
	// 	log.Println("Please Provide AccessKeyID")
	// 	s3Error = errors.New("S3 Config Error")
	// }

	// if len(s3config.SecretAccessKey) == 0 {
	// 	log.Println("Please Provide Secret AccessKey")
	// 	s3Error = errors.New("S3 Config Error")
	// }

	// if len(s3config.Region) == 0 {
	// 	log.Println("Please Provide S3 Region")
	// 	s3Error = errors.New("S3 Config Error")
	// }

	// if len(s3config.Bucket) == 0 {
	// 	log.Println("Please Provide S3 Bucket")
	// 	s3Error = errors.New("S3 Config Error")
	// }

	// if s3Error != nil {
	// 	return s3Error
	// }
	// // Create directory for logs
	// // ignore the error
	// setupDir()
	// // setup checkup.json
	// log.Println("Creating checkup.json")
	// checkup := checkup.Checkup{
	// 	Checkers: []checkup.Checker{},
	// 	Storage: checkup.S3{
	// 		AccessKeyID:     s3config.AccessKeyID,
	// 		SecretAccessKey: s3config.SecretAccessKey,
	// 		Region:          s3config.Region,
	// 		Bucket:          s3config.Bucket,
	// 	},
	// }
	// jsonBytes, err := checkup.MarshalJSON()
	// if err != nil {
	// 	return err
	// }
	// file, err := os.OpenFile(path.Join(config.DefaultCheckupDir, "checkup.json"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	// defer file.Close()
	// if err != nil {
	// 	return errors.Wrap(err, "Failed opening checkup.json file")
	// }
	// if _, err = file.Write(jsonBytes); err != nil {
	// 	return errors.Wrap(err, "Failed to write checkup.json")
	// }

	// // setup Caddyfile
	// // Read the template
	// log.Println("Creating Caddyfile")
	// caddyTemplate, err := template.New("caddyfile").Parse(templates.CaddyFile)
	// if err != nil {
	// 	return errors.Wrap(err, "Error parsing caddy template config")
	// }
	// caddyFile, err := os.OpenFile(path.Join(config.DefaultCaddyDir, "Caddyfile"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0660)
	// defer caddyFile.Close()
	// if err != nil {
	// 	return errors.Wrap(err, "Error opening Caddyfile")
	// }
	// // Execute the template
	// if err := caddyTemplate.Execute(caddyFile, struct{ URL string }{URL: c.String("url")}); err != nil {
	// 	return errors.Wrap(err, "Failed writing to Caddyfile")
	// }

	// // setup config status page
	// log.Println("Creating index.html for status page")

	// dstIndexHTML, err := os.OpenFile(path.Join(config.DefaultStatuspageDir, "index.html"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	// if _, err = dstIndexHTML.Write(templates.IndexHTML); err != nil {
	// 	return errors.Wrap(err, "Failed to write index.html")
	// }

	// log.Println("Creating minified app.js for status page")

	// // Create app.js
	// if err := buildAppMinJs(path.Join(config.DefaultStatuspageDir, "js/app.min.js"), "s3", s3config); err != nil {
	// 	return errors.Wrap(err, "failed to create app.min.js")
	// }

	// log.Println("Creating minified style.css for status page")

	// // Create app.js
	// if err := buildStyleCSS(path.Join(config.DefaultStatuspageDir, "css/style.min.css")); err != nil {
	// 	return errors.Wrap(err, "failed to create style.min.css")
	// }

	// setupImage()

	// log.Println("Success!")
	return nil
}

func setupDir() {
	log.Println("Setting up directory")
	os.MkdirAll(path.Join(config.DefaultCheckupDir), 0777)
	os.MkdirAll(path.Join(config.DefaultStatuspageDir, "images"), 0777)
	os.MkdirAll(path.Join(config.DefaultStatuspageDir, "css"), 0777)
	os.MkdirAll(path.Join(config.DefaultStatuspageDir, "js"), 0777)
	os.MkdirAll(path.Join(config.DefaultStatuspageDir, "logs"), 0777)
	os.MkdirAll(path.Join(config.DefaultCaddyDir, "logs"), 0777)
	os.MkdirAll(path.Join(config.DefaultCaddyDir, "errors"), 0777)
}

func setupImage() error {
	imagesGray, err := os.OpenFile(path.Join(config.DefaultStatuspageDir, "images/status-gray.png"), os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	if _, err := imagesGray.Write(templates.StatusGrayPNG); err != nil {
		return err
	}

	imagesGreen, err := os.OpenFile(path.Join(config.DefaultStatuspageDir, "images/status-green.png"), os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	if _, err := imagesGreen.Write(templates.StatusGreenPNG); err != nil {
		return err
	}

	imagesYellow, err := os.OpenFile(path.Join(config.DefaultStatuspageDir, "images/status-yellow.png"), os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	if _, err := imagesYellow.Write(templates.StatusYellowPNG); err != nil {
		return err
	}

	imagesRed, err := os.OpenFile(path.Join(config.DefaultStatuspageDir, "images/status-red.png"), os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	if _, err := imagesRed.Write(templates.StatusRedPNG); err != nil {
		return err
	}
	return nil
}

func buildAppMinJs(dst, name string, configData interface{}) error {
	var b bytes.Buffer
	_, err := b.Write(templates.CheckupJS)
	if err != nil {
		return err
	}
	switch name {
	case "fs":
		configJsTemplate, err := template.New("config.js").Parse(templates.ConfigFSJS)
		if err != nil {
			return err
		}
		if err := configJsTemplate.Execute(&b, configData); err != nil {
			return err
		}
		_, err = b.Write(templates.FSJS)
		if err != nil {
			return err
		}
	case "s3":
		configJsTemplate, err := template.New("config.js").Parse(templates.ConfigS3JS)
		if err != nil {
			return err
		}
		if err := configJsTemplate.Execute(&b, configData); err != nil {
			return err
		}
		_, err = b.Write(templates.S3JS)
		if err != nil {
			return err
		}
	}
	_, err = b.Write(templates.StatusPageJS)
	if err != nil {
		return err
	}

	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	appMinJs, err := os.OpenFile(dst, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	minBytes, err := m.Bytes("text/javascript", b.Bytes())
	if err != nil {
		return err
	}
	_, err = appMinJs.Write(minBytes)
	return err
}

func buildStyleCSS(dst string) error {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	styleCSS, err := os.OpenFile(dst, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	minBytes, err := m.Bytes("text/css", templates.StyleCSS)
	if err != nil {
		return err
	}
	_, err = styleCSS.Write(minBytes)
	return err
}
