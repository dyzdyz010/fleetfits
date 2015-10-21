<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<title>Fit: Tengu</title>

	<!-- Bootstrap -->
	<!-- <link href="/static/src/bower_components/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet"> -->
	<link href="/static/src/css/theme.min.css" rel="stylesheet">
</head>
<body>
	<header>
		<h1 class="logo text-center">Fit: HMLgu</h1>
	</header>

	<div class="container">
		<div class="row">
			<div class="col-xs-12">
				{{if eq .FitID "new"}}
				<textarea placeholder="Paste your fit from game" style="width: 80%;"></textarea>
				{{else}}
				<h1>Hell oworld!</h1>
				{{end}}
			</div>
		</div>
	</div>
</body>
</html>