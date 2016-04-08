
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>后台公告</title>
  <!-- 外部CSS库 -->
  <link href="/itrip-admin/plugin/bower_components/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">
  <link href="/itrip-admin/plugin/bower_components/font-awesome/css/font-awesome.min.css" rel="stylesheet">

  <link href="/itrip-admin/plugin/css/build/animate.css" rel="stylesheet">
  <link href="/itrip-admin/plugin/css/build/style.css" rel="stylesheet">
  <link href="/itrip-admin/plugin/css/build/info.css" rel="stylesheet">
</head>

<body>

<div id="wrapper">
  <jsp:include flush="true" page="side.jsp"></jsp:include> 

  <div id="page-wrapper" class="gray-bg">
    <jsp:include flush="true" page="head.jsp"></jsp:include> 
    <div class="wrapper wrapper-content animated fadeInDown">
      <div class="row">
        <div class="col-lg-12">
          <div class="m-t-lg">
            <div class="ibox">
              <div class="ibox-title">
                <h5>最新日志</h5>
                <div class="ibox-tools">
                  <a class="collapse-link">
                    <i class="fa fa-chevron-up"></i>
                  </a>
                </div>
              </div>
              <div class="ibox-content">
                <table class="footable table table-stripped" style="margin:0" data-page-size="15">
                  <thead>
                  <tr>
                    <th>级别</th>
                    <th>内容</th>
                    <th>时间</th>
                  </tr>
                  </thead>
                  <tbody>
                   <tr>
                        <td width="80px" class="text-danger">严重</td>
                        <td><a href="#">爬虫节点10.293.19.1停止运行!</a></td>
                        <td width="138px">2016-3-23 13:34:02</td>
                      </tr>
                      <tr>
                        <td class="text-warning">警告</td>
                        <td><a href="#">10.293.19.1 inite node fail!</a></td>
                        <td>2016-3-23 13:34:01</td>
                      </tr>
                      <tr>
                        <td class="text-warning">警告</td>
                        <td>
                          <a href="#">10.293.19.1 regist to zookeeper fail!</a></td>
                        <td>2016-3-23 13:33:59</td>
                      </tr>
                      <tr>
                        <td class="text-warning">警告</td>
                        <td><a href="#">10.293.19.1 no ZKClient available!</a></td>
                        <td>>2016-3-23 13:33.57</td>
                      </tr>
                      <tr>
                        <td class="text-warning">警告</td>
                        <td><a href="#">10.293.19.2 update clusterInfo tick time change fail</a></td>
                        <td>2016-3-23 09:17:33</td>
                      </tr>
                      <tr>
                        <td class="text-warning">警告</td>
                        <td><a href="#">10.293.19.2 regist to zookeeper fail!</a></td>
                        <td>2016-3-23 09:17:31</td>
                      </tr>
                  </tbody>
                  <tfoot>
                  <tr>
                    <td colspan="3">
                      <ul class="pagination pull-right">
                        <li><a href="#first">«</a></li>
                        <li><a href="#prev">‹</a></li>
                        <li class="active"><a data-page="0" href="#">1</a></li>
                        <li class="disabled"><a data-page="next" href="#next">›</a></li>
                        <li class="disabled"><a data-page="last" href="#last">»</a></li>
                      </ul>
                    </td>
                  </tr>
                  </tfoot>
                </table>


              </div>
            </div>

          </div>


        </div>
      </div>
    </div>
  </div>
  <div class="footer">
    <div class="text-center">
      <strong>&copy;2015</strong> 河海大学 华能集团
    </div>
  </div>

</div>
</div>

<!-- 第三方JS库 -->
<script src="/itrip-admin/plugin/bower_components/jquery/dist/jquery.min.js"></script>
<script src="/itrip-admin/plugin/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
<script src="/itrip-admin/plugin/bower_components/slimscroll/jquery.slimscroll.min.js"></script>
<script src="/itrip-admin/plugin/bower_components/Chart.js/Chart.min.js"></script>

<!-- 项目相关JS文件 -->
<script src="/itrip-admin/plugin/js/inspinia.js"></script>
<script src="/itrip-admin/plugin/js/pages/common.js"></script>
<script src="/itrip-admin/plugin/js/pages/info.js"></script>


</body>

</html>
