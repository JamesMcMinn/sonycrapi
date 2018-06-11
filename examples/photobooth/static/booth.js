var client = new ReconnectingWebSocket('ws://' + location.host + '/liveview');

var canvas = document.getElementById('videoCanvas');
var context = canvas.getContext('2d');

var currentPreviewName = '';
var countingDown = false;

var image = new Image();
image.onload = function() {
	context.drawImage(
		image,
		0,
		0,
		image.width,
		image.height, // source rectangle
		0,
		0,
		canvas.width,
		canvas.height
	); // destination rectangle
};

client.onmessage = function onmessage(event) {
	if (event.data instanceof Blob) {
		var reader = new FileReader();
		reader.onload = () => (image.src = 'data:image/png;base64,' + reader.result);
		reader.readAsText(event.data);
	}
};

function countDown() {
	countDowning = true;
	var timeLeft = 3;
	countdownElement = document.getElementById('countdown');
	countdownElement.innerHTML = timeLeft;

	countdownElement.className = 'visible';
	var interval = setInterval(function() {
		timeLeft--;

		countdownElement.innerHTML = timeLeft ? timeLeft : 'Smile!';
		if (timeLeft <= 0) {
			clearInterval(interval);
			takePhoto();
		}
	}, 1000);
}

function getReady() {
	if (!countingDown) {
		countDown();
	}
}

function takePhoto() {
	console.log('Take photo');

	httpGetAsync('/takephoto', function(json) {
		countingDown = false;
		console.log(json)
		json = JSON.parse(json);

		if (json.Error) {
			alert('Something went wrong. Go find James and let him know: takePhoto is broken:', json.Error);
			console.log(json.Error);
			return;
		}

		console.log(json);

		document.getElementById('countdown').className = 'hidden';
		showPreview(json.PhotoName);
	});
}

function showPreview(imageURL) {
	currentPreviewName = imageURL;
	imageURL = '/photo/' + imageURL;

	previewElement = document.getElementById('preview');
	previewPhoto = document.getElementById('previewPhoto');
	previewBorder = document.getElementById('previewBorder');
	loading = document.getElementById('loading');
	spinner = document.getElementById('spinner');

	previewPhoto.className = '';
	previewBorder.className = 'hidden';
	loading.className = 'visible';
	previewElement.className = 'visible';
	spinner.className = 'spinner';

	preloadImage(imageURL, function(img) {
		document.getElementById('previewPhoto').src = imageURL;
		previewBorder.className = 'visible';
		previewPhoto.className = 'fadeIn';
		loading.className = 'hidden';
		spinner.className = '';
	});
}

function another() {
	previewElement = document.getElementById('preview');
	previewElement.className = 'hidden';
}

function print() {
	previewElement = document.getElementById('preview');
	previewElement.className = 'hidden';

	printingElement = document.getElementById('printing');
	printingElement.className = 'visible';
	setTimeout(function() {
		printingElement.className = 'hidden';
	}, 10000);

	console.log('Print photo');
	httpGetAsync('/print/' + currentPreviewName, function() {});
}
