

function showMyApps() {
	$.get("getAppList", function(data){
			var ownsave = [], j = 0, ilen = data.length;
			for(var i = 0; i < ilen; i ++){
				if(data[i].isshow == 1){
					ownsave[j] = data[i].id;
					j ++;
				}
			}
			
			var app_manage = new Vue({
				el : 'body',
				data : {
					apps : data,
					owned : ownsave,
					saved : _.clone(ownsave)
				},
				computed : {
					// 待选列表
					unselected_apps : function() {
						return this.extract(this.apps, this.owned, true);
					},
					// 已选列表
					selected_apps : function() {
						return this.extract(this.apps, this.owned);
					},
					// 保存后正式显示的列表
					saved_apps : function() {
						return this.extract(this.apps, this.saved);
					}
				},
				ready : function() {
//					var saved_apps = localStorage.getItem('apps');
//					if (saved_apps) {
//						var apps = JSON.parse(saved_apps);
//						this.owned = apps;
//						this.saved = _.clone(apps);
//					}
				},
				methods : {
					add : function(id, pos) {
						if (!Number.isNaN(id)) {
							if (this.owned.indexOf(id) > -1) {
								this.owned.splice(this.owned.indexOf(id), 1);
							}
							this.owned.push(id);
						}

						console.log(this.owned);
					},
					del : function(id, pos) {
						var index = this.owned.indexOf(id);
						if (index >= 0)
							this.owned.splice(index, 1);
					},
					extract : function(all, part, rev) {
						if (rev) {
							return _.filter(all, function(e) {
								return part.indexOf(e.id) < 0;
							});
						} else {
							return _.map(part, function(e) {
								var result = _.find(all, function(a) {
									return a.id == e;
								});
								return result;
							});
						}

					},
					saveStat : function() {
						this.saved = _.clone(this.owned);
						//localStorage.setItem('apps', JSON.stringify(this.saved));
						// 将 saved 变量中的值同步到数据库
						$.post("setSite",{sites:JSON.stringify(this.saved)}, function(data){
							
						})
						$('#app_modal').modal('hide');
					}
				}

			});

			// 设置应用元素的拖拽开始事件。由于元素会动态变化，因此必须使用 delegate
			$('#app_modal').delegate('.drag-box', 'dragstart', function(e) {
				var target = $(e.target)
				var id = target.data('id');
				e.originalEvent.dataTransfer.dropEffect = "move";
				e.originalEvent.dataTransfer.setData('text/plain', target.find(
						'p.dim-description').text());
				e.originalEvent.dataTransfer.setData('application/appid', id);
			});

			$('#myApp, #unselected_apps').on('dragover', function(e) {
				// console.log(e)
				e.preventDefault();
				e.originalEvent.dataTransfer.dropEffect = "move"
			});

			$('#myApp, #unselected_apps').on('drop', function(e) {
				e.preventDefault();
				var data = e.originalEvent.dataTransfer.getData("application/appid");
				if (e.target.id == 'myApp') {
					app_manage.add(Number(data));
				} else if ($(e.target).parent().attr('id') == 'unselected_apps') {
					app_manage.del(Number(data));
				} else {
					console.log(e.target)
				}
			});
			
		
	});
}



$(function () {
	showMyApps();
});
