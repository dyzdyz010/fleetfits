<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<title>EVE Fit</title>

	<!-- jQuery Tags Input -->
	<link rel="stylesheet" type="text/css" href="/static/src/bower_components/jquery.tagsinput/src/jquery.tagsinput.css">

	<!-- Bootstrap -->
	<link href="/static/src/bower_components/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">
	<link href="/static/src/css/theme.min.css" rel="stylesheet">
</head>
<body>
	<header>
		<h1 class="logo text-center">Fit conversion</h1>
	</header>

	<div class="container">
		<div class="row" style="margin-top: 30px;">
			<div class="col-xs-6">
				<form class="form-horizontal" method="post" action="/">
					<fieldset>
						<legend>Fit String</legend>
<!-- 						<div class="form-group">
							<label for="fit-name" class="col-xs-2 control-label" style="font-weight: bold;">Name</label>
							<div class="col-md-10 col-xs-12">
								<input type="text" class="form-control" id="fit-name" placeholder="Fit Name" required>
							</div>
						</div> -->
						<div class="form-group">
							<label for="fit-str" class="col-xs-2 control-label" style="font-weight: bold;">Fit</label>
							<div class="col-md-10 col-xs-12">
								<textarea class="form-control" rows="10" id="fit-str" placeholder="Paste your fit from game" name="fitStr">{{.FitStr}}</textarea>
							</div>
						</div>
						<div class="form-group">
							<div class="col-lg-10 col-lg-offset-2">
								<input type="submit" class="btn btn-primary" name="submit" value="Submit">
							</div>
						</div>
					</fieldset>
				</form>
			</div>
			<div class="col-xs-6">
				<form class="form-horizontal">
					<fieldset>
						<legend>Parsed: </legend>
						<div class="form-group">
							<label class="col-xs-2 control-label" style="font-weight: bold;">Ship</label>
							<div class="col-md-10 col-xs-12">
								<div class="well well-sm row" id="preview-ship-name">{{.ShipName}}</div>
							</div>
						</div>
						<div class="form-group">
							<label class="col-xs-2 control-label" style="font-weight: bold;">In-game Link</label>
							<div class="col-md-10 col-xs-12">
								<div class="well well-sm row" id="fit-link">
									<a href="javascript:CCPEVE.showFitting('{{.FitLink}}');">{{.FitName}}</a>
								</div>
							</div>
						</div>
<!-- 						<div class="form-group">
							<label class="col-xs-2 control-label" style="font-weight: bold;">Equipments</label>
							<div class="col-md-10 col-xs-12">
								<div class="well well-sm row" id="preview-equipments"></div>
							</div>
						</div> -->
					</fieldset>
				</form>
			</div>
		</div>
	</div>

	<script src="/static/src/bower_components/jquery/dist/jquery.min.js"></script>
</body>
</html>