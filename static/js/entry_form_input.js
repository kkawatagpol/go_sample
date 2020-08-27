$(function() {
	$(".form .terms .scroll").mCustomScrollbar();

	var ua = navigator.userAgent;
	if (ua.indexOf('iPhone') < 0 && ua.indexOf('Android') < 0) {
		$('.telhref span').each(function () {
			$(this).unwrap();
		});
	}
});
