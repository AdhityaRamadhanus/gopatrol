checkup.config = {
	// How much history to show on the status page. Long durations and
	// frequent checks make for slow loading, so be conservative.
	"timeframe": 30 * time.Minute,

	// How often, in seconds, to pull new checks and update the page.
	"refresh_interval": 60,

	// Configure read-only access to stored checks. Currently, S3 is
	// supported. These credentials will be visible to everyone, so
	// use keys with ONLY read access!
	"storage": {
		"url": "logs"
	},

	// The text to display along the top bar depending on overall status.
	"status_text": {
		"healthy": "Situation Normal",
		"degraded": "Degraded Service",
		"down": "Service Disruption"
	}
};
