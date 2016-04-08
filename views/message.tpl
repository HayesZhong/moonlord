
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>集群设置</title>
<!-- 外部CSS库 -->
<link href="/itrip-admin/plugin/bower_components/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">
<link href="/itrip-admin/plugin/bower_components/font-awesome/css/font-awesome.min.css" rel="stylesheet">
<link href="/itrip-admin/plugin/bower_components/amcharts3/amcharts/plugins/export/export.css" rel="stylesheet">

<link href="/itrip-admin/plugin/bower_components/clockpicker/assets/clockpicker.css" rel="stylesheet">
<link href="/itrip-admin/plugin/bower_components/bootstrap-slider/css/bootstrap-slider.min.css" rel="stylesheet">

<link href="/itrip-admin/plugin/css/build/animate.css" rel="stylesheet">
<link href="/itrip-admin/plugin/css/build/style.css" rel="stylesheet">
<link href="/itrip-admin/plugin/css/build/second_index.css" rel="stylesheet">

</head>

<body>

	<div id="wrapper">
		<jsp:include flush="true" page="side.jsp"></jsp:include>


		<div id="page-wrapper" class="gray-bg">
			<jsp:include flush="true" page="head.jsp"></jsp:include>
			<div class="wrapper wrapper-content animated fadeInDown">
				<div class="row">
					<div class="col-lg-12">
						${message}
					</div>

				</div>

			</div>
			<!-- /.modal-content -->
		</div>
		<!-- /.modal-dialog -->
	</div>
	<!-- /.modal -->


	<!-- 第三方JS库 -->
	<script
		src="/itrip-admin/plugin/bower_components/jquery/dist/jquery.min.js"></script>
	<script
		src="/itrip-admin/plugin/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
	<script
		src="/itrip-admin/plugin/bower_components/slimscroll/jquery.slimscroll.min.js"></script>

	<!-- AMchart相关 -->
	<script
		src="/itrip-admin/plugin/bower_components/amcharts3/amcharts/amcharts.js"></script>
	<script
		src="/itrip-admin/plugin/bower_components/amcharts3/amcharts/plugins/dataloader/dataloader.min.js"></script>
	<script
		src="/itrip-admin/plugin/bower_components/amcharts3/amcharts/plugins/export/export.min.js"></script>
	<script
		src="/itrip-admin/plugin/bower_components/amcharts3/amcharts/serial.js"></script>
	<script src="/itrip-admin/plugin/js/vendor/amchart-lang.js"></script>


	<script src="/itrip-admin/plugin/bower_components/lodash/lodash.min.js"></script>
	<script src="/itrip-admin/plugin/bower_components/vue/dist/vue.min.js"></script>
	<script src="/itrip-admin/plugin/bower_components/clockpicker/assets/clockpicker.js"></script>
	<script src="/itrip-admin/plugin/bower_components/bootstrap-slider/js/bootstrap-slider.min.js"></script>
	
	<!-- 项目相关JS文件 -->
	<script src="/itrip-admin/plugin/js/inspinia.js"></script>
	<script src="/itrip-admin/plugin/js/pages/common.js"></script>
	
	<script src="/itrip-admin/plugin/js/pages/second_index.js"></script>


</body>

</html>
