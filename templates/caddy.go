package templates

var CaddyFile = `{{.URL}} {  
    root statuspage
    log caddy_config/logs/gopatrol.log
    errors caddy_config/errors/gopatrol.log
}`
