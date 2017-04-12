checkup.config = {
	"timeframe": 1 * time.Day,
	"refresh_interval": 60,
	"storage": {
		"AccessKeyID": "test",
		"SecretAccessKey": "test",
		"Region": "test",
		"BucketName": "test"
	},
	"status_text": {
		"healthy": "Situation Normal",
		"degraded": "Degraded Service",
		"down": "Service Disruption"
	}
};