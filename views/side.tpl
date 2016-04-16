<nav class="navbar-default navbar-static-side" role="navigation">
    <div class="sidebar-collapse">
      <ul class="nav">
        <li class="nav-header" id="profile-menu">
          <div class="profile-element" id="logo">

            <img src="../static/img/logo.png" class="logo-img" alt="" width="55" height="47">

            <h2>MoonLord</h2>
          </div>
        </li>
		
		{{if eq .pagename "tranerest"}}
		<li class="active">
		{{else}}
		<li>
		{{end}}
          <a href="index">
            <i class="fa fa-search"></i> <span class="nav-label">相似轨迹查询</span>
          </a>
        </li>
		
        {{if eq .pagename "pointnerest"}}
		<li class="active">
		{{else}}
		<li>
		{{end}}
          <a href="pointnerest">
            <i class="fa fa-map-marker"></i> <span class="nav-label">点最近轨迹</span>
          </a>
        </li>
		
        {{if eq .pagename "addtra"}}
		<li class="active">
		{{else}}
		<li>
		{{end}}
          <a href="addtra">
            <i class="fa fa-plus-square-o"></i> <span class="nav-label">添加轨迹</span>
          </a>
        </li>
      </ul>

    </div>
  </nav>