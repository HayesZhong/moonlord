
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
<style type="text/css">
#ex1Slider .slider-selection {
	background: #BABABA;
}
</style>
</head>

<body>

	<div id="wrapper">
		<jsp:include flush="true" page="side.jsp"></jsp:include>


		<div id="page-wrapper" class="gray-bg">
			<jsp:include flush="true" page="head.jsp"></jsp:include>
			<div class="wrapper wrapper-content animated fadeInDown">
				<div class="row">
					<div class="col-lg-12">
						<div class="tabs-container">
							<ul class="nav nav-tabs">
								<li class="active"><a data-toggle="tab" href="#tab-1">集群总体设置</a></li>
							</ul>
							<div class="tab-content">
								<div id="tab-1" class="tab-pane active">
									<div class="panel-body">
										<div class="row">
											<div class="col-lg-12">
												<div class="ibox">
													<div class="ibox-content">
													<form action="setclusterinfo.do" method="post">
														<table style="font-size:12px;"class="table table-striped text-left">
                      <tbody>
                      <tr><th>名称</th>
                      <th>参数</th>
                      </tr>
                      <tr>
                      	<td>集群启停</td>
                      	<td><select name="isrunning" class="form-control" style="width:100px;">
                      		<c:if test="${clusterinfo.cluster_status == 1}">
                      		<option value="1">启动</option>
                      		<option value="0">停止</option>
                      		</c:if>
                      		<c:if test="${clusterinfo.cluster_status == 0}">
                      		<option value="0">停止</option>
                      		<option value="1">启动</option>                     		
                      		</c:if>
   							
                      	</select></td>
                      	</tr>
                      	<tr>
                      	<td>爬取开始时间</td>
                      	<td>
                      	<div style="width:100px;" class="input-group clockpicker">
						    <input type="text" class="form-control" value="${clusterinfo.beginTimeList }:00" name="begin_time">
						    <span class="input-group-addon">
						        <span class="glyphicon glyphicon-time"></span>
						    </span>
						</div>
                      	</td>
					
                      </tr>
                      <tr>
                      	<td >爬虫节点状态更新时间间隔</td>
                      	<td style="padding-top:25px;">
                      	<input  name="tick" id="ex8" data-slider-id='ex1Slider' type="text" data-slider-min="1" data-slider-max="10" data-slider-step="1" data-slider-value="${clusterinfo.tick }"/>
                      	</td>
                      </tr>
                     <tr>
                     	<td style="float:right;"><button class="btn btn-default navbar-input"  onclick="quexiao"><a style="color:#fff;" href="clustersetting">取消</a></button></td>
                     	<td> <input type="submit" class="btn btn-primary navbar-input" value="保存"/></td>
                     </tr>
                     
                    </tbody></table>
                    	 
                    </form>
													</div>
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
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
	<script type="text/javascript">
	$('.clockpicker').clockpicker();
	$("#ex8").slider({
		tooltip: 'always'
	});

	</script>

</body>

</html>
