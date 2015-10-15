var app = {}
app.init = function() {
	$("#qrcode-create").bind("click", function() {
		qr_string = $("#qrstring").val()
		if (qr_string.length == 0) {
			app.error("二维码文字为空")
			return
		}
		if (qr_string.length > 120) {
			app.error("二维码文字过长")
			return
		}
		$(".qrcode-help").hide()
		$(".qrcode-loading").show()
		$("#qrcode-qrcode").attr("src", "/api?qr_string="+qr_string)
		$("#qrcode-qrcode").show()
	})
	$("#qrcode-qrcode").bind("load", function() {
		$(".qrcode-loading").hide()
		var _this = $(this)
		_this.css("visibility", "visible")
	})
}
app.error = function(str) {
	$("#qrcode-error").html(str)
	$("#qrcode-error").show("fast")
	$("#qrcode-error").delay(5000).hide("slow");
}
