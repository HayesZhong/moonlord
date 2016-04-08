
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>iTrip后台管理</title>
<!-- 外部CSS库 -->
<link href="/itrip-admin/plugin/bower_components/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">
<link href="/itrip-admin/plugin/bower_components/font-awesome/css/font-awesome.min.css" rel="stylesheet">
<link href="/itrip-admin/plugin/bower_components/amcharts3/amcharts/plugins/export/export.css" rel="stylesheet">
<link href="/itrip-admin/plugin/bower_components/bootstrap-switch/dist/css/bootstrap3/bootstrap-switch.min.css" rel="stylesheet">
<link href="/itrip-admin/plugin/css/build/animate.css" rel="stylesheet">
<link href="/itrip-admin/plugin/css/build/style.css" rel="stylesheet">
<link href="/itrip-admin/plugin/css/build/second_index.css"
	rel="stylesheet">
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
								<li class="active"><a data-toggle="tab" href="#tab-1">集群总体状态</a></li>
								<li class=""><a data-toggle="tab" href="#tab-2">各爬虫节点状态</a></li>
							</ul>
							<div class="tab-content">
								<div id="tab-1" class="tab-pane active">
									<div class="panel-body">
										<div class="row">
											<div class="col-lg-12">
												<div class="ibox">
													<div class="ibox-content">
														<table class="table table-striped text-left">
                      <tbody><tr><th>名称</th>
                      <th>数据量</th>
                      </tr><tr>
                        <td>集群健康状态</td>
                        <td class="text-navy"> <i class="fa fa-check"></i> 健康</td>
                        <!-- <td class="text-danger"> <i class="fa fa-times"></i> 失败</td> -->
                      </tr>
                     <c:forEach items="${clusterinfo }" var="item">
                      	<tr>
                      		<td>${item.key }</td>
	                        <td>${item.value }</td>
                      	</tr>
                      </c:forEach>	
                    </tbody></table>
													</div>
												</div>
											</div>
										</div>
									</div>
								</div>
								<div id="tab-2" class="tab-pane">
									<div class="panel-body">
										<div class="row">
											<div class="col-lg-12">
												<div class="ibox">
													<div class="ibox-content">
														<div>
															<div id="check_result_wrapper"
																class="dataTables_wrapper form-inline dt-bootstrap no-footer">
																<div class="row">
																	<div class="col-sm-6"></div>
																	<div class="col-sm-6"></div>
																</div>
																<div class="row">
																	<div class="col-sm-12">
																		<div class="dataTables_scroll">
																			<div class="dataTables_scrollHead"
																				style="overflow: hidden; position: relative; border: 0px; width: 100%;">
																				<div class="dataTables_scrollHeadInner"
																					style="box-sizing: content-box; width: 1311px; padding-right: 0px;">
																					<table
																						class="table table-striped table-bordered table-hover dataTable no-footer"
																						role="grid"
																						style="margin-left: 0px; width: 80%;">
																						<thead>
																							<tr>
																								<th width="170px" style="width: 170px; text-align: center;">IP</th>
																								<th width="150px" style="width: 150px; text-align: center;">状态</th>
																								<th width="150px" style="width: 150px; text-align: center;">允许爬取</th>
																								<th width="150px" style="width: 150px; text-align: center;">当前动作</th>
																								<th width="260px" style="width: 260px; text-align: center;">开始时间</th>
																								<th width="130px" style="width: 130px; text-align: center;">爬取页面数</th>
																								<th width="130px" style="width: 130px; text-align: center;">删除</th>
																							</tr>
																						</thead>
																						<tbody>
																							<c:forEach items="${nodesinfo }" var="node">
																								<tr id='${node[0] }' onclick="updatenodeinfo('${node[0] }')"  title="更新">
																									<c:forEach items="${node }" var="item" varStatus="index">
																										<c:choose>
																										<c:when test="${index.count != 3 }">
																											<td style="text-align: center;">${item }</td>
																										</c:when>
																										<c:otherwise>
																											<td style="text-align: center;">
																											<c:if test="${item eq '允许'}">
																				                      		<input id="switch"  name='allowswitch' type="checkbox" checked/>
																				                      		</c:if>
																				                      		<c:if test="${item eq '禁止'}">
																				                      		<input id="switch"  name='allowswitch' type="checkbox"/>     
																				                      		</c:if>
																				                      		</td>
																										</c:otherwise>
																										</c:choose>
																									</c:forEach>
																									<td style="text-align: center;">
																							 		<button class="btn btn-danger btn-xs" name="deletebutton" onclick="deleteSpider('${node[0] }')">删除</button>
																								 </td>
																								</tr>
																							 </c:forEach>
																							 
																						</tbody>
																					</table>
																				</div>
																			</div>
																		</div>
																	</div>
																</div>
																<div class="row">
																	<div class="col-sm-5">
																		<div class="dataTables_info" id="check_result_info"
																			role="status" aria-live="polite">显示第 0 至 ${fn:length(nodesinfo)}
																			项结果，共  ${fn:length(nodesinfo)} 项</div>
																	</div>
																	<div class="col-sm-7">
																		<div class="dataTables_paginate paging_simple_numbers"
																			id="check_result_paginate">
																			<ul class="pagination">
																				<li class="paginate_button previous disabled"
																					id="check_result_previous"><a href="#"
																					aria-controls="check_result" data-dt-idx="0"
																					tabindex="0">上页</a></li>
																				<li class="paginate_button next disabled"
																					id="check_result_next"><a href="#"
																					aria-controls="check_result" data-dt-idx="1"
																					tabindex="0">下页</a></li>
																			</ul>
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
	
	<script src="/itrip-admin/plugin/bower_components/bootstrap-switch/dist/js/bootstrap-switch.min.js"></script>
	

	<!-- 项目相关JS文件 -->
	<script src="/itrip-admin/plugin/js/inspinia.js"></script>
	<script src="/itrip-admin/plugin/js/pages/common.js"></script>

	<script src="/itrip-admin/plugin/js/pages/second_index.js"></script>
	<script>

	
		$.fn.bootstrapSwitch.defaults.size = 'mini';
		$.fn.bootstrapSwitch.defaults.onColor = 'success';
		$.fn.bootstrapSwitch.defaults.offColor = 'danger';
		$('input[name="allowswitch"]').bootstrapSwitch();
		$('input[name="allowswitch"]').on('switchChange.bootstrapSwitch', function(event, state) {
			  var ip = event.target.parentNode.parentNode.parentNode.parentNode.id;
			  $.get("setSpiderState",{ip:ip,state:state}, function(data){
					if(data.state == 1){
						$.get("getnodeinfof",{ip:ip}, function(nodeinfo){
							var td = $("tr[id='"+ip+"']").find("td").first();
							td = td.next().next().next();
							td.html("爬取中");
							td.next().html(nodeinfo.startTime);
							td.next().next().html(nodeinfo.pageNum+"个");
						},"json");						
					} else if(data.state == 0){
						var td = $("tr[id='"+ip+"']").find("td").first();
						td = td.next().next().next();
						td.html("");
						td.next().html("");
						td.next().next().html("");
					} else if(data.state == -1) {
						alert("更改失败,请重试!");
					}
				});
			});
		function deleteSpider(ip){
			var flag = window.confirm("确定删除ip为"+ip+"的爬虫节点？");
			if(flag == true) {
				$.get("deleteSpider",{ip:ip}, function(data){
					if(data.state == 1){
						$("tr[id='"+ip+"']").remove();
					} else if(data.state == 3) {
						alert("请先设置为禁止运行后再删除！");
					} else {
						alert("删除失败,请重试!");
					}
				});
			}
		}
		
		function updatenodeinfo(ip) {
			$.get("getnodeinfo",{ip:ip}, function(nodeinfo){
				var td = $("tr[id='"+ip+"']").find("td").first();
				td = td.next().next().next();
				//td.html("爬取中")
				//td.next().html(nodeinfo.startTime)
				td.next().next().html(nodeinfo.pageNum+"个");
			},"json");	
		}
	</script>


</body>

</html>
