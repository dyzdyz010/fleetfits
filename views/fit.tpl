<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<title>Fit: Tengu</title>

	<!-- jQuery Tags Input -->
	<link rel="stylesheet" type="text/css" href="/static/src/bower_components/jquery.tagsinput/src/jquery.tagsinput.css">

	<!-- Bootstrap -->
	<link href="/static/src/bower_components/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">
	<link href="/static/src/css/theme.min.css" rel="stylesheet">
</head>
<body>
	<header>
		<h1 class="logo text-center">Configure Fit</h1>
	</header>

	<div class="container">
		<div class="row" style="margin-top: 30px;">
			<div class="col-md-6 col-xs-12">
				<form class="form-horizontal">
					<fieldset>
						<legend>Fit: HMLgu</legend>
						<div class="form-group">
							<label for="fit-name" class="col-xs-2 control-label" style="font-weight: bold;">Name</label>
							<div class="col-md-10 col-xs-12">
								<input type="text" class="form-control" id="fit-name" placeholder="Fit Name" required>
							</div>
						</div>
						<div class="form-group">
							<label for="fit-str" class="col-xs-2 control-label" style="font-weight: bold;">Fit</label>
							<div class="col-md-10 col-xs-12">
								<textarea class="form-control" rows="10" id="fit-str" placeholder="Paste your fit from game"></textarea>
							</div>
						</div>
						<div class="form-group">
							<label for="fit-group" class="col-xs-2 control-label" style="font-weight: bold;">Group</label>
							<div class="col-md-10 col-xs-12">
								<label style="margin-top: 20px; margin-right: 10px;">
									<input type="checkbox" checked="checked"> Imperium
								</label>
								<label style="margin-top: 20px; margin-right: 10px;">
									<input type="checkbox" checked="checked"> FCON
								</label>

								<label style="margin-top: 20px; margin-right: 10px;">
									<input type="checkbox" checked="checked"> NN-S
								</label>							</div>
						</div>
						<div class="form-group">
							<label for="fit-tags" class="col-xs-2 control-label" style="font-weight: bold;">Tags</label>
							<div class="col-md-10 col-xs-12">
								<input name="tags" id="fit-tags">

								<label style="margin-top: 20px;">
									<input type="checkbox" checked="checked"> Active
								</label>
							</div>
						</div>
						<div class="form-group">
							<div class="col-lg-10 col-lg-offset-2">
								<a href="/" class="btn btn-default">Cancel</a>
								<button type="submit" class="btn btn-primary">Submit</button>
								<button type="button" class="btn btn-default" id="fit-parse">Parse</button>
							</div>
						</div>
					</fieldset>
				</form>
			</div>
			<div class="col-md-6 col-xs-12">
				<form class="form-horizontal">
					<fieldset>
						<legend>Parsed: </legend>
						<div class="form-group">
							<label class="col-xs-2 control-label" style="font-weight: bold;">Ship</label>
							<div class="col-md-10 col-xs-12">
								<div class="well well-sm row" id="preview-ship-name"></div>
							</div>
						</div>
						<div class="form-group">
							<label class="col-xs-2 control-label" style="font-weight: bold;">Equipments</label>
							<div class="col-md-10 col-xs-12">
								<div class="well well-sm row" id="preview-equipments">
<!-- 									<label class="col-xs-12">Co-Processor II</label>
									<label class="col-xs-12">Co-Processor II</label>
									<label class="col-xs-12">Co-Processor II</label>
									<label class="col-xs-12">Co-Processor II</label>
									<label class="col-xs-12">Loki</label>
									<label class="col-xs-12">Loki</label>
									<label class="col-xs-12">Loki</label>
									<label class="col-xs-12">Loki</label>
 -->								</div>
							</div>
						</div>
					</fieldset>
				</form>
			</div>
		</div>
	</div>

	<script src="/static/src/bower_components/jquery/dist/jquery.min.js"></script>
	<script src="/static/src/bower_components/jquery.tagsinput/src/jquery.tagsinput.js"></script>
	<script src="/static/src/js/tag-edit.js"></script>
</body>
</html>