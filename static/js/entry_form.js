/* checkPolicy
----------------------------------------------*/
$(function() {
	var $check_policy = $('#checkPolicy'),
		$btn_confirm = $('.confirm button');
	
	var _set_check_policy = function(){
		if ( $check_policy.prop('checked') == false ) {
			$btn_confirm.attr('disabled', 'disabled');
		} else {
			$btn_confirm.removeAttr('disabled');
		}
	};
	
	$check_policy.on('click', _set_check_policy);
	_set_check_policy();
});
