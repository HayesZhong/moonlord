<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>MoonLord</title>
	<!-- 外部CSS库 -->
	<link
	href="../static/bower_components/bootstrap/dist/css/bootstrap.min.css"
	rel="stylesheet">
	<link
	href="../static/bower_components/datetimepicker/jquery.datetimepicker.css"
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
		<div id="page-wrapper" class="gray-bg" style="min-height: 10px;">
			<!-- {{template "head.tpl" .}} -->
			<div class="wrapper wrapper-content" style="padding: 5px 10px 1px;">
				<div  class="row show-grid">
					<div class="col-md-12" >
						<form class="form-inline" style="float:right">
							<div class="form-group">
								<label class="sr-only" for="loninput">经度</label>
								<div class="input-group">
									<input type="text" class="form-control" id="loninput" placeholder="经度">
									<div class="input-group-addon">度</div>
								</div>
							</div>
							<div class="form-group">
								<label class="sr-only" for="latinput">纬度</label>
								<div class="input-group">
									<input type="text" class="form-control" id="latinput" placeholder="纬度">
									<div class="input-group-addon">度</div>
								</div>
							</div>
							<div class="form-group">
							<div class="input-group">
								<input id="datetimepicker" size="18" class="form-control" type="text" placeholder="日期"></div>
							</div>
							<div class="form-group">
								<label class="sr-only" for="limitinput">条数</label>
								<div class="input-group">
									<input type="text" class="form-control" size="10" id="limitinput" placeholder="Limit">
								</div>
							</div>
							<a class="btn btn-success" href="#" onclick="searchtras()" role="button">Search</a>
						</form>
					</div>
				</div>
				<div class="row show-grid">

					<div class="col-md-12" id="baidumap"></div>
				</div>
			</div>
			<div class="footer">
				<div class="text-center"> <strong>&copy;2016</strong>
					河海大学 MoonLord团队
				</div>
			</div>

		</div>
	</div>

	<!-- 第三方JS库 -->
	<script src="../static/bower_components/jquery/dist/jquery.js"></script>
	<script
		src="../static/bower_components/bootstrap/dist/js/bootstrap.js"></script>

	<script src="../static/bower_components/datetimepicker/build/jquery.datetimepicker.full.js"></script>

	<script type="text/javascript"
		src="http://api.map.baidu.com/api?v=2.0&ak=V96wYEUYmulq4y8EgooYPGhEXhdBzabU"></script>
	<script type="text/javascript">
		$('#datetimepicker').datetimepicker({
            	lang:'ch',        //语言选择中文
		      format:"Y-m-d-h-i-s",      //格式化日期
		      yearStart:1970,     //设置最小年份
		      yearEnd:2018,        //设置最大年份
        });

		var map = new BMap.Map("baidumap"); // 创建地图实例  
		var initPoint = new BMap.Point(116.404, 39.915); // 创建点坐标  
		map.centerAndZoom(initPoint, 11); // 初始化地图，设置中心点坐标和地图级别  
		map.enableScrollWheelZoom();
		map.addControl(new BMap.NavigationControl());
		map.addControl(new BMap.ScaleControl());
		map.addControl(new BMap.OverviewMapControl());
		map.addControl(new BMap.MapTypeControl());

		

		map.addEventListener("rightclick", function(e){
			map.clearOverlays() 
 			var point = new BMap.Point(e.point.lng,e.point.lat);
 			map.addOverlay(new BMap.Marker(point));
 			$("#loninput").val(e.point.lng)
 			$("#latinput").val(e.point.lat)
		});

		var colors = new Array("red","blue","yellow","black")
		function searchtras() {
			map.clearOverlays()
			
			var lon = $("#loninput").val()
			var lat = $("#latinput").val()
			var time = $("#datetimepicker").val()
			var limit = $("#limitinput").val()
			var point = new BMap.Point(Number(lon),Number(lat));
			
			map.addOverlay(new BMap.Marker(point));	
			
			// 将标注添加到地图中
			$.getJSON("api/nearestmtras?limit="+limit+"&time="+time+"&lat="+lat+"&lon="+lon, function(result) {
				if(result.status == 0) {
					map.centerAndZoom(point, 15);
					alert(result.message)
				} else {
					for  (var i = 0; i < result.mtras.length; i++) {
						for (var j = 0; j < result.mtras[i].length-1; j++) {
							var polyline = new BMap.Polyline([new BMap.Point(result.mtras[i][j].lon, result.mtras[i][j].lat),
							                                  new BMap.Point(result.mtras[i][j+1].lon,result.mtras[i][j+1].lat)], {
								strokeColor : colors[i%4],
								strokeWeight : 6,
								strokeOpacity : 0.5
							});
							map.addOverlay(polyline);
						}
					}
					map.centerAndZoom( new BMap.Point(result.center.lon,result.center.lat), result.zoom);				
				}								
			});
		}

		
	</script></body>

</html>