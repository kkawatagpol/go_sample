$(window).on('load',function(){
	// ログアウト処理
	$('button.logout').on('click', function(){
		location.href = '/logout/';
		return false;
	});
	// pagetop -----------------------------------------------------------
	var $pagetop = $("#pagetop");
	$pagetop.click(function () {
		$('body,html').animate({scrollTop: 0}, 1000);
		return false;
	});
	$(window).on("load scroll", function () {
		if ($(this).scrollTop() > 300) {
			$pagetop.fadeIn(300);
		} else {
			$pagetop.fadeOut(300);
		}
		scrollHeight = $(document).height();
		scrollPosition = $(window).height() + $(window).scrollTop();
		footHeight = $("#footer").height();
		if (scrollHeight - scrollPosition <= footHeight) {
			$pagetop.removeClass("fixed");
		} else {
			$pagetop.addClass("fixed");
		}
	});
});