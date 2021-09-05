
function removeIPv4Domain(a,k) {
  var c = a.parent(), i = a.data('index');
  var lastIndex = c.find('.input-group[data-index]:last').data('index');
  a.remove();
  if (lastIndex == i) {
    return;
  }
  var prefix = 'DNSServices['+k+'][IPv4Domains][';
  resortDomain(c, prefix);
}
function removeIPv6Domain(a,k) {
  var c = a.parent(), i = a.data('index');
  var lastIndex = c.find('.input-group[data-index]:last').data('index');
  a.remove();
  if (lastIndex == i) {
    return;
  }
  var prefix = 'DNSServices['+k+'][IPv6Domains][';
  resortDomain(c, prefix);
}
function resortDomain(c, prefix){
    c.find('.input-group[data-index]').each(function(index){
        var $this = $(this);
        var oldIndex = $this.data('index');
        if(index == oldIndex) return;
        $this.attr('data-index',index);
        $this.find('[name^="'+prefix+'"]').each(function(){
          var name = $(this).attr('name');
          name = name.substring(prefix.length);
          name = name.replace(/^[0-9]+/,index);
          name = prefix+name;
          $(this).attr('name',name);
        })
    });
}
function addIPv4Domain(a,k,supportLine) {
  var c = a.parent();
  var i = c.find('.input-group[data-index]:last').data('index')+1;
  var t = template('tmpl-domain-row',{k:k,domainK:i,supportLine:supportLine,ipVer:4});
  c.append(t);
}
function addIPv6Domain(a,k,supportLine) {
  var c = a.parent();
  var i = c.find('.input-group[data-index]:last').data('index')+1;
  var t = template('tmpl-domain-row',{k:k,domainK:i,supportLine:supportLine,ipVer:6});
  c.append(t);
}
var ipv4NetInterfaceIPRule = 'IPv4[NetInterface][Filter][Include]',ipv6NetInterfaceIPRule = 'IPv6[NetInterface][Filter][Include]';
function insertNetIfaceRegexpTag(ipVer){
    var name = ipVer==6?ipv6NetInterfaceIPRule:ipv4NetInterfaceIPRule;
    var value = $('input[name="'+name+'"]').val();
    if(/^regexp:/.test(value)) return;
    $('input[name="'+name+'"]').val('regexp:'+value);
}
$(function(){
    // $('#ddns-form').off().on('submit',function(e){
    //     e.preventDefault();
    //     $.post(window.location.href,$(this).serialize(),function(r){
    //         if(r.Code==1){
    //             $('#search-result').text(r.Data);
    //             return;
    //         }
    //     },'json');
    // });
    $('input[name="IPv6[Type]"],input[name="IPv4[Type]"]').on('click',function(){
      var name = $(this).attr('name');
      $('div[rel="'+this.id+'"]').show();
      $('input[name="'+name+'"]:not(:checked)').each(function(){
        $('div[rel="'+this.id+'"]').hide();
      });
    });
    $('input[name="IPv6[Type]"]:checked,input[name="IPv4[Type]"]:checked').trigger('click');
    $('.provider-switch-onoff').on('click',function(){
        var rel = $(this).attr('rel'), on = $(this).val()=='1';
        on ? $('#'+rel).removeClass('hide') : $('#'+rel+':not(.hide)').addClass('hide');
    });
    $('.provider-switch-onoff:checked').trigger('click');

    var notifyTemplateName = 'NotifyTemplate[html]';
    $('textarea[name="NotifyTemplate[html]"],textarea[name="NotifyTemplate[markdown]"]').on('focus',function(){
        notifyTemplateName = $(this).attr('name');
    });
    $('#notify-template-tag-values code').on('click',function(){
        App.insertAtCursor($('textarea[name="'+notifyTemplateName+'"]')[0],$(this).text());
    });
    $('input[name="IPv4[NetInterface][Filter][Include]"],input[name="IPv4[NetInterface][Filter][Exclude]"]').on('focus',function(){
        ipv4NetInterfaceIPRule = $(this).attr('name');
    });
    $('input[name="IPv6[NetInterface][Filter][Include]"],input[name="IPv6[NetInterface][Filter][Exclude]"]').on('focus',function(){
        ipv6NetInterfaceIPRule = $(this).attr('name');
    });
});