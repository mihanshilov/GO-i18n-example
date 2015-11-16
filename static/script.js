function setCookie(cname, cvalue, exdays) {
    var d = new Date();
    d.setTime(d.getTime() + (exdays*24*60*60*1000));
    var expires = "expires="+d.toUTCString();
    document.cookie = cname + "=" + cvalue + "; " + expires;
}

function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1);
        if (c.indexOf(name) == 0) return c.substring(name.length,c.length);
    }
    return "";
}

$(function(){
    const cookieName = "lang";
    var lang = getCookie(cookieName);
    if(!lang)
    {
        lang = "en";
    }

    var $localeSelector = $("#lng");
    $localeSelector.val(lang);
    $localeSelector.change(function(){
        var option = $(this).find('option:selected').val();
        setCookie(cookieName, option, 10000);
        document.location.href = document.location.href;
    });
});