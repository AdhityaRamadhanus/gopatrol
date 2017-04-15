package templates

var ConfigFSJS = `checkup.config = {
	"timeframe": 1 * time.Day,
	"refresh_interval": 60,
	"storage": {
		"url": "logs"
	},
	"status_text": {
		"healthy": "Situation Normal",
		"degraded": "Degraded Service",
		"down": "Service Disruption"
	}
};`

var ConfigS3JS = `checkup.config = {
	"timeframe": 1 * time.Day,
	"refresh_interval": 60,
	"storage": {
		"AccessKeyID": "{{.AccessKeyID}}",
		"SecretAccessKey": "{{.SecretAccessKey}}",
		"Region": "{{.Region}}",
		"BucketName": "{{.Bucket}}"
	},
	"status_text": {
		"healthy": "Situation Normal",
		"degraded": "Degraded Service",
		"down": "Service Disruption"
	}
};`

var IndexHTML = `<!DOCTYPE html>
<html>
	<head>
		<title>Status Page</title>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<script src="js/d3.v3.min.js" charset="utf-8"></script>
		<script src="js/{{.Type}}.js"></script>
		<script src="js/checkup.js"></script>
		<script src="js/config.js"></script>
		<script src="js/statuspage.js"></script>
		<link rel="icon" href="images/favicon.png" id="favicon">
		<link rel="stylesheet" href="https://rawgit.com/allinurl/goaccess/master/resources/css/bootstrap.min.css">		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
		<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:400,300,700">
		<link rel="stylesheet" href="css/style.css">
	</head>
	<body>
		<div class="app container">
			<header id="overall-status" class="gray">
				<div id="divOverallText">
					<i class="fa fa-stethoscope" style="font-size:36px;color: #FFF;"></i>
					<span class="overall-status-text">Loading</id>
				</div>
				<div class="infobar">
					<div class="item">
						<b>Last check:</b> <span id="info-lastcheck">&mdash;</span>
					</div>
				</div>
			</header>

			<main>
				<div id="chart-grid">
					
					<span id="chart-placeholder">&nbsp;</span>
				</div>
				<div id="timeline">
					<div id="big-gap">
						There is a big gap of time where no checkups were performed, so some graphs may look distorted.
					</div>
					<div id="bg-line"></div>
				</div>
			</main>
		</div>
		<footer>
			Powered by <img src="images/checkup.png" id="checkup-logo">
		</footer>
	</body>
</html>`
