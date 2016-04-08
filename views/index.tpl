<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>MoonLord</title>
<!-- 外部CSS库 -->
<link
	href="../static/bower_components/bootstrap/dist/css/bootstrap.min.css"
	rel="stylesheet">
<link
	href="../static/bower_components/font-awesome/css/font-awesome.min.css"
	rel="stylesheet">

<link href="../static/css/build/animate.css" rel="stylesheet">
<link href="../static/css/build/style.css" rel="stylesheet">
<link href="../static/css/build/second_index.css" rel="stylesheet">
<style type="text/css">
#baidumap {
	height: 600px
}
</style>
</head>

<body>

	<div id="wrapper">
		{{template "side.tpl" .}}
		<div id="page-wrapper" class="gray-bg">
			{{template "head.tpl" .}}
			<div class="wrapper wrapper-content">
				<div class="row">
					<div class="col-md-12" id="baidumap"></div>
				</div>
			</div>
			<div class="footer">
				<div class="text-center">
					<strong>&copy;2016</strong> 河海大学 MoonLord团队
				</div>
			</div>

		</div>
	</div>


	<!-- 第三方JS库 -->
	<script src="../static/bower_components/jquery/dist/jquery.min.js"></script>
	<script
		src="../static/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
	<script type="text/javascript"
		src="http://api.map.baidu.com/api?v=2.0&ak=V96wYEUYmulq4y8EgooYPGhEXhdBzabU"></script>
	<script type="text/javascript">
		var map = new BMap.Map("baidumap"); // 创建地图实例  
		var point = new BMap.Point(116.404, 39.915); // 创建点坐标  
		map.centerAndZoom(point, 19); // 初始化地图，设置中心点坐标和地图级别  
		map.enableScrollWheelZoom();

		$.getJSON("api/getonetras?index={{.index}}", function(result) {
			
			for (var i = 0; i < result.length-1; i++) {
				var polyline = new BMap.Polyline([new BMap.Point(result[i].lon, result[i].lat),
				                                  new BMap.Point(result[i+1].lon, result[i+1].lat)], {
					strokeColor : "blue",
					strokeWeight : 6,
					strokeOpacity : 0.5
				});
				
				var polylineFilter = new BMap.Polyline([new BMap.Point(result[i].flon, result[i].flat),
				                                  new BMap.Point(result[i+1].flon, result[i+1].flat)], {
					strokeColor : "red",
					strokeWeight : 6,
					strokeOpacity : 0.5
				});
				
				map.addOverlay(polyline);
				map.addOverlay(polylineFilter);
			}
			var center = 0;
			if(result.length % 2 == 0) {
				center = result.length / 2
			} else {
				center = (result.length-1) / 2
			}
			map.centerAndZoom(new BMap.Point(result[center].lon, result[center].lat), 19);
		});

		map.addControl(new BMap.NavigationControl());
		map.addControl(new BMap.ScaleControl());
		map.addControl(new BMap.OverviewMapControl());
		map.addControl(new BMap.MapTypeControl());
	</script>

</body>

</html>
