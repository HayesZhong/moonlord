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
		
									<input type="file" class="form-control" id="trafile" name="trafile"  placeholder="轨迹文件">
							</div>
							<a class="btn btn-success" href="#" onclick="uploadtrafile()" role="button">Search</a>
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
		<script src="../static/js/ajaxfileupload.js"></script>


	<script type="text/javascript"
		src="http://api.map.baidu.com/api?v=2.0&ak=V96wYEUYmulq4y8EgooYPGhEXhdBzabU"></script>
	<script type="text/javascript">
		

		var map = new BMap.Map("baidumap"); // 创建地图实例  
		var initPoint = new BMap.Point(116.404, 39.915); // 创建点坐标  
		map.centerAndZoom(initPoint, 19); // 初始化地图，设置中心点坐标和地图级别  
		map.enableScrollWheelZoom();
		map.addControl(new BMap.NavigationControl());
		map.addControl(new BMap.ScaleControl());
		map.addControl(new BMap.OverviewMapControl());
		map.addControl(new BMap.MapTypeControl());

		var colors = new Array("red","blue","yellow","black")

		function uploadtrafile() {
			 $.ajaxFileUpload({
                url: 'api/getsimtra', //用于文件上传的服务器端请求地址
                secureuri: false, //是否需要安全协议，一般设置为false
                fileElementId: 'trafile', //文件上传域的ID
                dataType: 'json', //返回值类型 一般设置为json
                success: function (data, status){  //服务器成功响应处理函数
                	if(data.status == "0") {
                		alert(data.message)
                	} else {
	                	map.clearOverlays()
	                	var result = data.tras
	                    for  (var i = 0; i < result.length; i++) {
							for (var j = 0; j < result[i].length-1; j++) {
								var polyline = new BMap.Polyline([new BMap.Point(result[i][j].lon, result[i][j].lat),
								                                  new BMap.Point(result[i][j+1].lon,result[i][j+1].lat)], {
									strokeColor : colors[i%4],
									strokeWeight : 6,
									strokeOpacity : 0.5
								});
								map.addOverlay(polyline);
							}
						}
					}
                },
                error: function (data, status, e) {//服务器响应失败处理函数     
                	alert (JSON.stringify(data))      
                    alert(e);
                }
            });
		}	
	</script>
</body>

</html>