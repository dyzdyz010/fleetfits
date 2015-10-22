$(document).ready(function () {
	$('#fit-tags').tagsInput();

	$('#fit-parse').click(function () {
		var fitStr = $('#fit-str').val();
		if (fitStr == '') {return};
		// console.log(fitStr);
		$.post('/admin/fit/parse', fitStr, function (fit) {
			console.log(fit.items);
			$('#preview-ship-name').html(fit.ship);
			for (var i = 0; i < fit.items.length; i++) {
				$('#preview-equipments').append('<label class="col-xs-12">'+fit.items[i]+'</label>');
			};
		})
	});
});