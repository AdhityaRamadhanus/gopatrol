// config.js must be included BEFORE this file!

// Configure access to storage
checkup.storage.setup(checkup.config.storage);

// Once the DOM is loaded, go ahead and render the graphs
// (if it hasn't been done already).
document.addEventListener('DOMContentLoaded', function () {
	checkup.domReady = true;

	checkup.dom.favicon = document.getElementById("favicon");
	checkup.dom.status = document.getElementById("overall-status");
	checkup.dom.statustext = document.getElementsByClassName("overall-status-text")[0];
	checkup.dom.lastcheck = document.getElementById("info-lastcheck");
	checkup.dom.timeline = document.getElementById("timeline");

	if (!checkup.graphsMade) makeGraphs();
}, false);

// Immediately begin downloading check files, and keep page updated 
checkup.storage.getChecksWithin(checkup.config.timeframe, processNewCheckFile, allCheckFilesLoaded);
setInterval(function() {
	checkup.storage.getNewChecks(processNewCheckFile, allCheckFilesLoaded);
}, checkup.config.refresh_interval * 1000);

// Update "time ago" tags every so often
setInterval(function() {
	var times = document.querySelectorAll("time.dynamic");
	for (var i = 0; i < times.length; i++) {
		var timeEl = times[i];
		timeEl.innerHTML = timeEl.getAttribute("datetime")
	}
}, 5000);


function processNewCheckFile (json, filename) {
	checkup.checks.push(json);

	// update the timestamp of the last check file's timestamp
	var dashLoc = filename.indexOf("-");
	if (dashLoc > 0) {
		var checkTs = Number(filename.substr(0, dashLoc));
		if (checkTs > checkup.lastCheckTs) {
			checkup.lastCheckTs = checkTs;
		}
	}

	// iterate each result and store/process it
	for (var j = 0; j < json.length; j++) {
		var result = json[j];

		// Save stats with the result so we don't have to recompute them later
		result.stats = checkup.computeStats(result);

		checkup.orderedResults.push(result); // will sort later, more efficient that way

		if (!checkup.groupedResults[result.timestamp])
			checkup.groupedResults[result.timestamp] = [result];
		else
			checkup.groupedResults[result.timestamp].push(result);

		if (!checkup.results[result.endpoint])
			checkup.results[result.endpoint] = [result];
		else
			checkup.results[result.endpoint].push(result);

		var chart = checkup.charts[result.endpoint] || checkup.makeChart(result.title);
		chart.results.push(result);

		var ts = checkup.unixNanoToD3Timestamp(result.timestamp);
		chart.series.min.push({ timestamp: ts, rtt: Math.round(result.stats.min) });
		chart.series.med.push({ timestamp: ts, rtt: Math.round(result.stats.median) });
		chart.series.max.push({ timestamp: ts, rtt: Math.round(result.stats.max) });
		if (result.threshold)
			chart.series.threshold.push({ timestamp: ts, rtt: result.threshold });

		checkup.charts[result.endpoint] = chart;
		checkup.charts[result.endpoint].endpoint = result.endpoint;

		for (var s in chart.series) {
			chart.series[s].sort(function(a, b) {
			  return a.timestamp - b.timestamp;
			});
		}

		if (!checkup.lastResultTs || ts > checkup.lastResultTs) {
			checkup.lastResultTs = ts;
			checkup.dom.lastcheck.innerHTML = checkup.makeTimeTag(checkup.lastResultTs);
		}
	}

	if (checkup.domReady) {
		makeGraphs();
	}
}

function allCheckFilesLoaded (numChecksLoaded, numResultsLoaded) {
	// Sort the result lists
	checkup.orderedResults.sort(function(a, b) { return a.timestamp - b.timestamp; });
	for (var endpoint in checkup.results)
		checkup.results[endpoint].sort(function(a, b) { return a.timestamp - b.timestamp; });
	
	// Create events for the timeline

	var newEvents = [];
	var statuses = {}; // keyed by endpoint

	// First load the last known status of each endpoint
	for (var i = checkup.events.length-1; i >= 0; i--) {
		var result = checkup.events[i].result;
		if (!statuses[result.endpoint])
			statuses[result.endpoint] = checkup.events[i].status;
	}

	// Then go through the new results and look for new events
	for (var i = checkup.orderedResults.length-numResultsLoaded; i < checkup.orderedResults.length; i++) {
		var result = checkup.orderedResults[i];

		var status = "healthy";
		if (result.degraded) status = "degraded";
		else if (result.down) status = "down";

		if (status != statuses[result.endpoint]) {
			// New event because status changed
			newEvents.push({
				id: checkup.eventCounter++,
				result: result,
				status: status
			});
		}
		if (result.message) {
			// New event because message posted
			newEvents.push({
				id: checkup.eventCounter++,
				result: result,
				status: status,
				message: result.message
			});
		}

		statuses[result.endpoint] = status;
	}

	checkup.events = checkup.events.concat(newEvents);

	function renderTime(ns) {
		var d = new Date(ns * 1e-6);
		var hours = d.getHours();
		var ampm = "AM";
		if (hours > 12) {
			hours -= 12;
			ampm = "PM";
		}
		return hours+":"+checkup.leftpad(d.getMinutes(), 2, "0")+" "+ampm;
	}

	// Render events
	for (var i = 0; i < newEvents.length; i++) {
		var e = newEvents[i];

		// Save this event to the chart's event series so it will render on the graph
		var imgFile = "ok.png"; 
		var imgWidth = 10;
		var imgHeight = 10; // the different icons look smaller/larger because of their shape
		if (e.status == "down") { 
			imgFile = "incident.png"; 
			imgWidth = 10; 
			imgHeight = 10; 
		} else if (e.status == "degraded") { 
			imgFile = "degraded.png"; 
			imgWidth = 10; 
			imgHeight = 10;
		}
		var chart = checkup.charts[e.result.endpoint];
		chart.series.events.push({
			timestamp: checkup.unixNanoToD3Timestamp(e.result.timestamp),
			rtt: e.result.stats.median,
			eventid: e.id,
			imgFile: imgFile,
			imgWidth: imgWidth,
			imgHeight: imgHeight
		});

		// Render event to timeline
		var evtElem = document.createElement("div");
		evtElem.setAttribute("data-eventid", e.id);
		evtElem.classList.add("event-item");
		evtElem.classList.add("event-id-"+e.id);
		evtElem.classList.add(checkup.color[e.status]);
		if (e.message) {
			evtElem.classList.add("message");
			evtElem.innerHTML = '<div class="message-head">'+checkup.makeTimeTag(e.result.timestamp*1e-6)+' ago</div>';
			evtElem.innerHTML += '<div class="message-body">'+e.message+'</div>';
		} else {
			evtElem.classList.add("event");
			evtElem.innerHTML = '<span class="time">'+renderTime(e.result.timestamp)+'</span> '+e.result.title+" "+e.status;
		}
		checkup.dom.timeline.insertBefore(evtElem, checkup.dom.timeline.childNodes[0]);
	}

	// Update DOM now that we have the whole picture

	// Update overall status
	var overall = "healthy";
	for (var endpoint in checkup.results) {
		if (overall == "down") break;
		var lastResult = checkup.results[endpoint][checkup.results[endpoint].length-1];
		if (lastResult) {
			if (lastResult.down)
				overall = "down";
			else if (lastResult.degraded)
				overall = "degraded";
		}
	}

	if (overall == "healthy") {
		checkup.dom.favicon.href = "images/status-green.png";
		checkup.dom.status.className = "green";
		checkup.dom.statustext.innerHTML = checkup.config.status_text.healthy || "System Nominal";
	} else if (overall == "degraded") {
		checkup.dom.favicon.href = "images/status-yellow.png";
		checkup.dom.status.className = "yellow";
		checkup.dom.statustext.innerHTML = checkup.config.status_text.degraded || "Sub-Optimal";
	} else if (overall == "down") {
		checkup.dom.favicon.href = "images/status-red.png";
		checkup.dom.status.className = "red";
		checkup.dom.statustext.innerHTML = checkup.config.status_text.down || "Outage";
	} else {
		checkup.dom.favicon.href = "images/status-gray.png";
		checkup.dom.status.className = "gray";
		checkup.dom.statustext.innerHTML = checkup.config.status_text.unknown || "Status Unknown";
	}


	// Detect big gaps in any of the charts, and if there is one, show an explanation.
	var bigGap = false;
	var lastTimeDiff;
	for (var key in checkup.charts) {
		// We expect results to be chronologically ordered, but since they are downloaded
		// in an arbitrary order due to network conditions, we have to sort to be sure.
		checkup.charts[key].results.sort(function(a, b) {
			return a.timestamp - b.timestamp;
		});
		for (var k = 1; k < checkup.charts[key].results.length; k++) {
			var timeDiff = Math.abs(checkup.charts[key].results[k].timestamp - checkup.charts[key].results[k-1].timestamp);
			bigGap = lastTimeDiff && timeDiff > lastTimeDiff * 10;
			lastTimeDiff = timeDiff;
			if (bigGap) {
				document.getElementById("big-gap").style.display = 'block';
				break;
			}
		}
		if (bigGap) break;
	}
	if (!bigGap) {
		document.getElementById("big-gap").style.display = 'none';
	}

	makeGraphs(); // must render graphs again after we've filled in the event series
	updateInfoBar();

}

function updateInfoBar () {
	let checkLen = checkup.checks.length
	let totalChecksElm = document.getElementById('info-totalchecks')
	totalChecksElm.innerHTML = checkLen
	let healthyEndpointElm = document.getElementById('info-totalhealthy')
	healthyEndpointElm.innerHTML = checkup.checks[checkLen - 1].filter((check) => check.healthy).length
	let downEndpointElm = document.getElementById('info-totaldown')
	downEndpointElm.innerHTML = checkup.checks[checkLen - 1].filter((check) => !check.healthy).length
}

function makeGraphs () {
	if (!checkup.placeholdersRemoved && checkup.checks.length > 0) {
		// Remove placeholder to make way for the charts;
		// placeholder necessary to give space in absense of charts.
		let placeHolderElm = document.getElementById("chart-placeholder");
		if (placeHolderElm) placeHolderElm.remove(); 
		checkup.placeholdersRemoved = true;
	}

	for (var endpoint in checkup.charts) {
		makeGraph(checkup.charts[endpoint], endpoint);
  }
	checkup.graphsMade = true;
}

function makeGraph (chart, endpoint) {
	// Render chart to page if first time seeing this endpoint
	if (!chart.elem) renderChart(chart);

	// Define scale for x axis
	let xMax = d3.max(chart.data, c => d3.max(c, d => d.timestamp));
	let xMin = d3.min(chart.data, c => d3.min(c, d => d.timestamp));
	chart.xScale.domain([xMin, xMax]).nice();

	// Define scale for y axis
	let yMax = d3.max(chart.data, c => d3.max(c, d => d.rtt));
	let yMin = d3.min(chart.data, c => d3.min(c, d => d.rtt));
	chart.yScale.domain([0, yMax]).nice();

	chart.xAxis = d3.svg.axis()
		.scale(chart.xScale)
		.ticks(5)
		.outerTickSize(0)
		.orient("bottom");

	chart.yAxis = d3.svg.axis()
		.scale(chart.yScale)
		.tickFormat(checkup.formatDuration)
		.outerTickSize(0)
		.ticks(5)
		.orient("left");

	if (chart.svg.selectAll(".x.axis")[0].length == 0) {
		chart.svg.insert("g", ":first-child")
			.attr("class", "x axis")
			.attr("transform", "translate(0," + chart.height + ")")
			.call(chart.xAxis);
	} else {
		chart.svg.selectAll(".x.axis")
			.transition()
			.duration(checkup.animDuration)
			.call(chart.xAxis);
	}

	if (chart.svg.selectAll(".y.axis")[0].length == 0) {
		chart.svg.insert("g", ":first-child")
			.attr("class", "y axis")
			.call(chart.yAxis);
	} else {
		chart.svg.selectAll(".y.axis")
			.transition()
			.duration(checkup.animDuration)
			.call(chart.yAxis);
	}

	chart.lines = chart.lineGroup.selectAll(".line").data(chart.data);
	chart.events = chart.eventGroup.selectAll("image").data(chart.series.events);

	// transition from old paths to new paths
	chart.lines
		.transition()
		.duration(checkup.animDuration)
		.attr("d", chart.line);

	chart.events
		.transition()
		.duration(checkup.animDuration);

	// enter any new data (lines)
	chart.lines.enter()
	  .append("path")
		.attr("class", (d) => {
			switch (d) {
				case chart.series.min:
					return "min line";
				case chart.series.med:
					return "main line";
				case chart.series.max:
					return "max line";
				case chart.series.threshold:
					return "tolerance line";
				default:
					return "line"
			}
		})
		.attr("d", chart.line);

	// enter any new data (events)
	chart.events.enter().append("svg:image")
		.attr("width", (d, i) => d.imgWidth || 0)
		.attr("height", (d, i) => d.imgHeight || 0)
		.attr("xlink:href", (d, i) => "images/"+d.imgFile)
		.attr("x", (d, i) => chart.xScale(d.timestamp) - (d.imgWidth/2))
		.attr("y", (d, i) => chart.yScale(d.rtt) - (d.imgHeight/2))
		.attr("data-eventid", (d, i) => d.eventid)
		.attr("class", (d, i) => "event-item event-id-"+d.eventid)
		.on("mouseover", (e) => {
			let events = document.querySelectorAll(".event-item:not(.event-id-"+e.eventid+")");
			events.forEach((event) => {
				event.style.opacity = ".25";
			})
		})
		.on("mouseout", (e) => {
			let events = document.querySelectorAll(".event-item:not(.event-id-"+e.eventid+")");
			events.forEach((event) => {
				event.style.opacity = "";
			})
		});

	// exit any old data
	chart.lines
		.exit()
		.remove();
}

function renderChart (chart) {
	// Outer div is a wrapper that we use for layout
	let chartContainerElm = document.createElement('div');
	chartContainerElm.className = "chart-container chart-50";
	if (document.getElementsByClassName('chart-container').length == 0) {
		chartContainerElm.className = "chart-container chart-100";
	} else {
		// It's possible that a chart was created that, at the time,
		// was the only one, but now it is too wide, since there are
		// at least two charts. Resize the wide one to be smaller.
		let wideChartElm = document.querySelector('.chart-container.chart-100');
		if (wideChartElm) wideChartElm.className = "chart-container chart-50";
	}

	// Div to contain the endpoint / title
	let chartTitleElm = document.createElement('div');
	chartTitleElm.className = "chart-title-container";

	let chartTitleLink = document.createElement('a');
	chartTitleLink.appendChild(document.createTextNode(chart.title))
	chartTitleLink.href = chart.endpoint
	chartTitleLink.title = chart.title

	let chartTitleText = document.createElement('h3'); 
	chartTitleText.className = "pull-left chart-title"
	chartTitleText.appendChild(chartTitleLink);
	chartTitleElm.appendChild(chartTitleText);

	let chartTitleCol = document.createElement('div');
	chartTitleCol.className = "col-md-12";
	let chartTitleRow = document.createElement('div');
	chartTitleRow.className = "row";

	chartTitleCol.appendChild(chartTitleElm)
	chartTitleRow.appendChild(chartTitleCol)

	// Inner div is used to contain the actual svg tag
	let chartSvgElm = document.createElement('div');
	chartSvgElm.className = "chart";
	let chartSvgRow = document.createElement('div');
	chartSvgRow.className = "row";
	chartSvgRow.appendChild(chartSvgElm)

	chartContainerElm
		.appendChild(chartTitleRow)
	
	chartContainerElm
		.appendChild(chartSvgRow)

	// Inject elements into DOM
	document
		.getElementById('chart-grid')
		.appendChild(chartContainerElm);

	// Save it with the chart and use D3 to set up its svg element.
	chart.elem = chartSvgElm;
	chart.svgTag = d3.select(chart.elem)
	  .append("svg")
	    .attr("id", chart.id)
		.attr("preserveAspectRatio", "xMinYMin meet")
		.attr("viewBox", "0 0 " + checkup.CHART_WIDTH + " " + checkup.CHART_HEIGHT);

	chart.margin = {top: 20, right: 20, bottom: 40, left: 55};
	chart.width = checkup.CHART_WIDTH - chart.margin.left - chart.margin.right;
	chart.height = checkup.CHART_HEIGHT - chart.margin.top - chart.margin.bottom;

	// chart.svgTag is the svg tag, but chart.svg
	// is the group where we actually put the lines.
	chart.svg = chart.svgTag
	  .append("g")
	  	.attr("class", "chart-data")
		.attr("transform", "translate(" + chart.margin.left + "," + chart.margin.top + ")");

	chart.xScale = d3.time.scale()
		.range([0, chart.width]);

	chart.yScale = d3.scale.linear()
		.range([chart.height, 0]);

	chart.line = d3.svg.line()
		.x(d => chart.xScale(d.timestamp))
		.y(d => chart.yScale(d.rtt))
		.interpolate("linear"); // linear, monotone, or basis

	chart.lineGroup = chart.svg
	  .append("g")
		.attr("class", "lines")

	// Focus Markers
	let focus = chart.svg
	  .append("g")
		.attr("class", "focus")
		.style("display", "none");

	focus.append("circle")
		.attr("r", 3);

	// Focus Text
	let textRtt = focus.append("text")
		.attr("x", 9)
		.attr("dy", ".35em")
		.attr("class", "focus-text rtt");
	let textTs = focus.append("text")
		.attr("x", 9)
		.attr("dy", ".35em")
		.attr("class", "focus-text ts");
	
	// Next we build an overlay to cover the data area,
	// so when the mouse hovers it we can show the point.
	let bisectDate = d3.bisector(d => d.timestamp).left;
	var overlay;
	chart.svg.append("rect")
		.attr("class", "overlay")
		.attr("width", chart.width)
		.attr("height", chart.height)
		.on('mouseout', () => focus.style('display', 'none'))
		.on('mousemove', function () { // need closure to chart and overlay
			let x0 = chart.xScale.invert(d3.mouse(this)[0]);
			let dateMid = bisectDate(chart.series.med, x0, 1);
			let dataLeft = chart.series.med[dateMid - 1];
			let dataRight = chart.series.med[dateMid];
			let data = (dataLeft && dataRight) 
				? (x0 - dataLeft.timestamp > dataRight.timestamp - x0 ? dataRight : dataLeft) 
				: (dataLeft || dataRight);
			let xloc = chart.xScale(data.timestamp);			
			if (xloc > overlay.width.animVal.value - 50){
				textRtt.attr("transform", "translate(-60, 10)");
				textTs.attr("transform", "translate(-60, 23)");
			} else {
				textRtt.attr("transform", "translate(0, 10)");
				textTs.attr("transform", "translate(0, 23)");
			}
			focus
				.attr("transform", "translate(" + xloc + "," + chart.yScale(data.rtt) + ")")
				.style('display', 'inline');
			focus.select(".focus-text.rtt").text(checkup.formatDuration(data.rtt));
			focus.select(".focus-text.ts").text(checkup.compactDateTimeString(data.timestamp));
		})
	overlay	= document.querySelector("#"+chart.id+" .overlay");

	chart.eventGroup = chart.svg
	  .append("g")
		.attr("class", "events");
}
