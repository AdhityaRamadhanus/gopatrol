package templates

var CaddyFile = `{{.URL}} {  
    root /statuspage
    log /root/.caddy/logs/checkup.log {
        rotate {
            size 100 # Rotate after 100 MB
            age  14  # Keep log files for 14 days
            keep 10  # Keep at most 10 log files
        }
    }
    errors /root/.caddy/errors/checkup.log
}`
