/**

S3 Storage Adapter for Checkup.js

**/

var checkup = checkup || {};

checkup.storage = (function() {
	var bucket, bucketName, region;

	// getCheckFileList gets the list of check files within
	// the given timeframe (as a unit of nanoseconds) to
	// download. Only check files are added to the list; other
	// file types will be ignored.
	function getCheckFileList(timeframe, callback) {
		var allObjects = [];

		function getObjectsAfter(marker) {
			bucket.listObjects({
				Marker: marker
			}, function(err, data) {
				allObjects = allObjects.concat(data.Contents);
				if (data.IsTruncated) {
					getObjectsAfter(data.Contents[data.Contents.length-1].Key);
				} else {
					// TODO: Can this be converted to a map() function?
					var keys = [];
					for (var i = 0; i < allObjects.length; i++) {
						if (allObjects[i].Key.indexOf(checkup.checkFileSuffix) != -1)
							keys.push(allObjects[i].Key);
					}
					callback(keys);
				}
			});
		}

		getObjectsAfter("" + (time.Now() - timeframe))
	};

	// setup prepares this storage unit to operate.
	this.setup = function(cfg) {
		AWS.config.update({accessKeyId: cfg.AccessKeyID, secretAccessKey: cfg.SecretAccessKey})
		bucket = new AWS.S3({
			params: {
				Bucket: cfg.BucketName,
			}
		});
		bucketName = cfg.BucketName;
		region = cfg.Region;
	};

	// getChecksWithin gets all the checks within timeframe as a unit
	// of nanoseconds, and executes callback for each check file.
	this.getChecksWithin = function(timeframe, fileCallback, doneCallback) {
		var checksLoaded = 0, resultsLoaded = 0;
		getCheckFileList(timeframe, function(list) {
			if (list.length == 0 && (typeof doneCallback === 'function')) {
				doneCallback(checksLoaded);
			} else {
				for (var i = 0; i < list.length; i++) {
					var url;
					if (region && region !== "" && region !== "us-east-1") {
						url = "https://s3-"+region+".amazonaws.com/"+bucketName+"/"+list[i];
					} else {
						url = "https://s3.amazonaws.com/"+bucketName+"/"+list[i];
					}
					checkup.getJSON(url, function(filename) {
						return function(json, url) {
							checksLoaded++;
							resultsLoaded += json.length;
							if (typeof fileCallback === 'function')
								fileCallback(json, filename);
							if (checksLoaded >= list.length && (typeof doneCallback === 'function'))
								doneCallback(checksLoaded, resultsLoaded);
						};
					}(list[i]));
				}
			}
		});
	};

	// getNewChecks gets any checks since the timestamp on the file name
	// of the youngest check file that has been downloaded. If no check
	// files have been downloaded, no new check files will be loaded.
	this.getNewChecks = function(fileCallback, doneCallback) {
		if (!checkup.lastCheckTs == null)
			return;
		var timeframe = time.Now() - checkup.lastCheckTs;
		return this.getChecksWithin(timeframe, fileCallback, doneCallback);
	};

	return this;
})();
