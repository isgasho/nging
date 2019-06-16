
$(function(){
	App.init();
	App.initLeftNavAjax(window.activeURL,null);
	App.bottomFloat('.pagination');
	App.bottomFloat('.form-submit-group',0,true);
    if(window.errorMSG) App.message({title: App.i18n.SYS_INFO, text: window.errorMSG, class_name: "danger"});
	if(window.successMSG) App.message({title: App.i18n.SYS_INFO, text: window.successMSG, class_name: "success"});
	App.notifyListen();
});