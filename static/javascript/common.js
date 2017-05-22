function httpRequest(method, url, body) {
	return new Promise(function(resolve, reject){
		var req = new XMLHttpRequest;
		req.open(method, url);

		req.onload = function() {
			if (req.status == 200) {
				resolve(req.response);
			}
		};

		req.onerror = function() {
			reject(Error("Fail to httpRequest " + url));
		};

		req.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
		req.send(body);
	});
};
