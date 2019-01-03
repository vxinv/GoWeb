var OK = 0;
var MSIE = navigator.userAgent.indexOf('MSIE') > -1 ? true : false;
var MSIE6 = navigator.userAgent.indexOf('MSIE 6') > -1 ? true : false;
var Moz = navigator.userAgent.indexOf('Firefox') > -1 ? true : false;
var Opera = navigator.userAgent.indexOf('Opera') > -1 ? true : false;
var Safari = navigator.userAgent.indexOf('Safari') > -1 ? true : false;
var Action = { Create: 0, List: 1, Grid: 2, View: 3, Edit: 4, Delete: 5, Disable: 6, Enable: 7 };
var UserType = { All: 0, Mine: 1, Staff: 2, Share: 3 };
var DefaultSplitor = "|";
var $ = function (id) { return document.getElementById(id); };
var $P = function (id) { return parent.document.getElementById(id); };
var $T = function (id) { return top.document.getElementById(id); };
var $N = function (id) { return document.getElementsByName(id); };
window.query = function (val) { var uri = window.location.search; var re = new RegExp("" + val + "=([^&?]*)", "ig"); return ((uri.match(re)) ? (uri.match(re)[0].substr(val.length + 1)) : ""); };
window.cookie = { set: function (sName, sValue) { var date = new Date(); var year = date.getFullYear() + 1; var newDate = new Date(year, date.getMonth(), date.getDay()); document.cookie = sName + "=" + escape(sValue) + "; expires=" + newDate.toGMTString(); }, get: function (sName) { var aCookie = document.cookie.split("; "); for (var i = 0; i < aCookie.length; i++) { var aCrumb = aCookie[i].split("="); if (sName == aCrumb[0]) return unescape(aCrumb[1]); } return ""; } };
String.prototype.trim = function () { var reExtraSpace; try { switch (arguments[0].toLowerCase()) { case "comma": reExtraSpace = /^,*(.*?),*$/; break; case "bigcomma": reExtraSpace = /^，*(.*?)，*$/; break; case "dollar": reExtraSpace = /^\$*(.*?)\$*$/; break; default: reExtraSpace = /^\s*(.*?)\s+$/; break; } } catch (e) { reExtraSpace = /^\s*(.*?)\s+$/; } return this.replace(reExtraSpace, "$1"); };
String.prototype.encodeHTML = function () { var val = this.toString().trim(); if (val == '') return '&nbsp;'; val = val.replace(/&/g, "&amp;"); val = val.replace(/</g, "&lt;"); val = val.replace(/>/g, "&gt;"); val = val.replace(/\n/g, '<br />'); return val; }
Array.prototype.indexOf = function (s) { for (var i = 0; i < this.length; i++) { if (this[i] == s) { return i; break; } } };
Date.prototype.dateValue = function () { if (this.getFullYear() != 1900) { var rst = ''; rst += this.getFullYear() + '-'; rst += (this.getMonth() + 1) < 10 ? '0' + (this.getMonth() + 1) + '-' : (this.getMonth() + 1) + '-'; rst += this.getDate() < 10 ? '0' + this.getDate() : this.getDate(); return rst; } else { return ''; } }
var http = {
    _objPool: [], _getInstance: function () 
    { for (var i = 0; i < this._objPool.length; i++) { if (this._objPool[i].readyState == 0) { return this._objPool[i]; } } 
    this._objPool[this._objPool.length] = this._createObj(); return this._objPool[this._objPool.length - 1]; },
    _createObj: function () { if (window.XMLHttpRequest) { var objXMLHttp = new XMLHttpRequest(); } 
    else { var MSXML = ['Microsoft.XMLHTTP', 'MSXML2.XMLHTTP.5.0', 'MSXML2.XMLHTTP.4.0', 'MSXML2.XMLHTTP.3.0', 'MSXML2.XMLHTTP'];
     for (var n = 0; n < MSXML.length; n++) { try { var objXMLHttp = new ActiveXObject(MSXML[n]); break; } catch (e) { } } }
      if (objXMLHttp.readyState == null) { objXMLHttp.readyState = 0; objXMLHttp.addEventListener("load", function () 
      { objXMLHttp.readyState = 4; if (typeof objXMLHttp.onreadystatechange == "function") { objXMLHttp.onreadystatechange(); } }, false); }
       return objXMLHttp; },
    load: function (method, url, data, callback) {
        var objXMLHttp = this._getInstance(); var mode = arguments.length > 4 ? arguments[4] : true;
        with (objXMLHttp) {
            if (url.indexOf("?") > 0) { url += "&randnum=" + Math.random(); } else { url += "?randnum=" + Math.random(); }
            open(method, url, mode); setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=utf-8');
            send(data); if (mode) {
                onreadystatechange = function () {
                    if (objXMLHttp.readyState == 4 && (objXMLHttp.status == 200 || objXMLHttp.status == 304)) { callback(objXMLHttp); }
                    else if (objXMLHttp.readyState == 4 && (objXMLHttp.status != 200 && objXMLHttp.status != 304)) {
                        if (confirm("发生一个错误，单击“确定”立即显示错误信息，单击“取消”转到错误页面")) { alert(objXMLHttp.responseText); } else { location = url + "&" + data; }
                    }
                }
            } else {
                if (objXMLHttp.readyState == 4 && (objXMLHttp.status == 200 || objXMLHttp.status == 304)) { return callback(objXMLHttp); }
                else if (objXMLHttp.readyState == 4 && (objXMLHttp.status != 200 && objXMLHttp.status != 304)) {
                    if (confirm("发生一个错误，单击“确定”立即显示错误信息，单击“取消”转到错误页面")) { try { alert(objXMLHttp.responseText); } catch (e) { alert(objXMLHttp.responseText); } } else { location = url + "&" + data; }
                }
            }
        }
    }
}
var JSON = JSON || {}; (function () {
    function f(n) { return n < 10 ? '0' + n : n; }
    if (typeof Date.prototype.toJSON !== 'function') {
        Date.prototype.toJSON = function (key) {
            return isFinite(this.valueOf()) ? this.getUTCFullYear() + '-' +
                f(this.getUTCMonth() + 1) + '-' +
                f(this.getUTCDate()) + 'T' +
                f(this.getUTCHours()) + ':' +
                f(this.getUTCMinutes()) + ':' +
                f(this.getUTCSeconds()) + 'Z' : null;
        }; String.prototype.toJSON = Number.prototype.toJSON = Boolean.prototype.toJSON = function (key) { return this.valueOf(); };
    }
    var cx = /[\u0000\u00ad\u0600-\u0604\u070f\u17b4\u17b5\u200c-\u200f\u2028-\u202f\u2060-\u206f\ufeff\ufff0-\uffff]/g, escapable = /[\\\"\x00-\x1f\x7f-\x9f\u00ad\u0600-\u0604\u070f\u17b4\u17b5\u200c-\u200f\u2028-\u202f\u2060-\u206f\ufeff\ufff0-\uffff]/g, gap, indent, meta = { '\b': '\\b', '\t': '\\t', '\n': '\\n', '\f': '\\f', '\r': '\\r', '"': '\\"', '\\': '\\\\' }, rep; function quote(string) { escapable.lastIndex = 0; return escapable.test(string) ? '"' + string.replace(escapable, function (a) { var c = meta[a]; return typeof c === 'string' ? c : '\\u' + ('0000' + a.charCodeAt(0).toString(16)).slice(-4); }) + '"' : '"' + string + '"'; }
    function str(key, holder) {
        var i, k, v, length, mind = gap, partial, value = holder[key]; if (value && typeof value === 'object' && typeof value.toJSON === 'function') { value = value.toJSON(key); }
        if (typeof rep === 'function') { value = rep.call(holder, key, value); }
        switch (typeof value) {
            case 'string': return quote(value); case 'number': return isFinite(value) ? String(value) : 'null'; case 'boolean': case 'null': return String(value); case 'object': if (!value) { return 'null'; }
                gap += indent; partial = []; if (Object.prototype.toString.apply(value) === '[object Array]') {
                    length = value.length; for (i = 0; i < length; i += 1) { partial[i] = str(i, value) || 'null'; }
                    v = partial.length === 0 ? '[]' : gap ? '[\n' + gap +
                        partial.join(',\n' + gap) + '\n' +
                        mind + ']' : '[' + partial.join(',') + ']'; gap = mind; return v;
                }
                if (rep && typeof rep === 'object') { length = rep.length; for (i = 0; i < length; i += 1) { k = rep[i]; if (typeof k === 'string') { v = str(k, value); if (v) { partial.push(quote(k) + (gap ? ': ' : ':') + v); } } } } else { for (k in value) { if (Object.hasOwnProperty.call(value, k)) { v = str(k, value); if (v) { partial.push(quote(k) + (gap ? ': ' : ':') + v); } } } }
                v = partial.length === 0 ? '{}' : gap ? '{\n' + gap + partial.join(',\n' + gap) + '\n' +
                    mind + '}' : '{' + partial.join(',') + '}'; gap = mind; return v;
        }
    }
    //if(typeof JSON.stringify!=='function'){
    JSON.stringify = function (value, replacer, space) {
        var i; gap = ''; indent = ''; if (typeof space === 'number') { for (i = 0; i < space; i += 1) { indent += ' '; } } else if (typeof space === 'string') { indent = space; }
        rep = replacer; if (replacer && typeof replacer !== 'function' && (typeof replacer !== 'object' || typeof replacer.length !== 'number')) { throw new Error('JSON.stringify'); }
        return str('', { '': value });
    };
    //}
    if (typeof JSON.parse !== 'function') {
    JSON.parse = function (text, reviver) {
        var j; function walk(holder, key) {
            var k, v, value = holder[key]; if (value && typeof value === 'object') { for (k in value) { if (Object.hasOwnProperty.call(value, k)) { v = walk(value, k); if (v !== undefined) { value[k] = v; } else { delete value[k]; } } } }
            return reviver.call(holder, key, value);
        }
        cx.lastIndex = 0; if (cx.test(text)) {
            text = text.replace(cx, function (a) {
                return '\\u' +
                    ('0000' + a.charCodeAt(0).toString(16)).slice(-4);
            });
        }
        if (/^[\],:{}\s]*$/.test(text.replace(/\\(?:["\\\/bfnrt]|u[0-9a-fA-F]{4})/g, '@').replace(/"[^"\\\n\r]*"|true|false|null|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?/g, ']').replace(/(?:^|:|,)(?:\s*\[)+/g, '').replace(/new Date\(\]\)/g, ']'))) { j = eval('(' + text + ')'); return typeof reviver === 'function' ? walk({ '': j }, '') : j; }
        throw new SyntaxError('JSON.parse');
    };
    }
}());
var Class = { create: function () { return function () { this.initialize.apply(this, arguments); }; } };
Function.prototype.bind = function () { var wc = this, a = $A(arguments), o = a.shift(); return function () { wc.apply(o, a.concat($A(arguments))); }; };
Object.extend = function (a, b) { for (var i in b) a[i] = b[i]; return a; };
Object.extend(Object, {
    addEvent: function (a, b, c, d) { if (a.attachEvent) a.attachEvent(b[0], c); else a.addEventListener(b[1] || b[0].replace(/^on/, ""), c, d || false); return c; },
    delEvent: function (a, b, c, d) { if (a.detachEvent) a.detachEvent(b[0], c); else a.removeEventListener(b[1] || b[0].replace(/^on/, ""), c, d || false); return c; },
    reEvent: function () { return window.event ? window.event : (function (o) { do { o = o.caller; } while (o && !/^\[object[ A-Za-z]*Event\]$/.test(o.arguments[0])); return o.arguments[0]; })(this.reEvent); }
});
function TabPage() { this.CurrentTab; }
TabPage.prototype = {
    set: function (menus, divs, openClass, closeClass) {
        var _this = this;
        if (menus.length != divs.length) {
            alert("菜单层数量和内容层数量不一样!");
            return false;
        }
        _this.CurrentTab = _this.$(divs[0]).id;
        for (var i = 0; i < menus.length; i++) {
            _this.$(menus[i]).value = i;
            _this.$(menus[i]).onclick = function () {
                for (var j = 0; j < menus.length; j++) {
                    _this.$(menus[j]).className = closeClass;
                    _this.$(divs[j]).style.display = "none";
                }
                _this.$(menus[this.value]).className = openClass;
                _this.$(divs[this.value]).style.display = "block";
                _this.CurrentTab = _this.$(divs[this.value]).id;
            }
        }
    },
    $: function (oid) { if (typeof (oid) == "string") return document.getElementById(oid); return oid; }
}
window.showErrorTo = function (o, s) { o.innerHTML = s; o.style.visibility = 'visible'; setTimeout(function () { o.style.visibility = 'hidden'; }, 2000); };
window.fromJSON = function (s) {
    var reviver = arguments.length > 1 ? arguments[1] : null;
    try {
        if (typeof reviver === "Function") {
            return JSON.parse(s, reviver);
        } else {
            return JSON.parse(s);
        }
    } catch (e) {
        if (s.substr(0, 1) == '(') return eval(s);
        else return eval('(' + s + ')');
    }
};
window.toJSON = function (o) {
    if (typeof o !== 'object') return null;
    var encode = arguments.length > 1 && typeof (arguments[1]) == 'boolean' ? arguments[3] : true;
    var replacer = arguments.length > 2 && typeof (arguments[2]) == 'object' ? arguments[1] : null;
    var space = arguments.length > 3 ? arguments[3] : null;
    if (encode) return encodeURIComponent(JSON.stringify(o, replacer, space));
    else return JSON.stringify(o, replacer, space);
}
window.insertHTML = function (where, el, html) {
    if (el.insertAdjacentHTML) {
        switch (where) {
            case "beforeBegin": el.insertAdjacentHTML(where, html); return el.previousSibling; break;
            case "afterBegin": el.insertAdjacentHTML(where, html); return el.firstChild; break;
            case "beforeEnd": el.insertAdjacentHTML(where, html); return el.lastChild; break;
            case "afterEnd": el.insertAdjacentHTML(where, html); return el.nextSibling; break;
        }
        throw '未知的插入点 -> "' + where + '"';
    } else {
        var range = el.ownerDocument.createRange(); var frag;
        switch (where) {
            case "beforeBegin": range.setStartBefore(el); frag = range.createContextualFragment(html); el.parentNode.insertBefore(frag, el); return el.previousSibling; break;
            case "afterBegin": if (el.firstChild) { range.setStartBefore(el.firstChild); frag = range.createContextualFragment(html); el.insertBefore(frag, el.firstChild); return el.firstChild; } else { el.innerHTML = html; return el.firstChild; } break;
            case "beforeEnd": if (el.lastChild) { range.setStartAfter(el.lastChild); frag = range.createContextualFragment(html); el.appendChild(frag); return el.lastChild; } else { el.innerHTML = html; return el.lastChild; } break;
            case "afterEnd": range.setStartAfter(el); frag = range.createContextualFragment(html); el.parentNode.insertBefore(frag, el.nextSibling); return el.nextSibling; break;
        }
        throw '未知的插入点 -> "' + where + '"';
    }
}
window.getAppPath = function (app) {
    var rst = '';
    var tmp = document.location.pathname.toString();
    if (/\/.*\//i.test(tmp)) {
        var m = tmp.match(/\/.*\//i).toString().toLowerCase();
        if (m.indexOf(app.toLowerCase()) > -1) { return '' }
        else { return m.split('/').length >= 4 ? '../' + app : app; }
    } else {
        //没有建虚拟目录的情况
        return '/' + app;
    }
}
window.halt = false;
String.prototype.padLeft = function (l, s) { var val = this.toString(); if (val.length < l) { for (var i = 0; i < l - val.length; i++) val = s + val; } return val; }
window.outerHTML = function (elmt) {
    var h = [], arg = arguments[1] ? arguments[1].match(/(\w+ ?=?('|"| )?.[^ '"=]+('|")?)/ig) : [];
    var p = [];
    for (var i = 0; i < arg.length; i++) {
        p[p.length] = arg[i].split('=')[0].trim().toLowerCase();
    }
    h[h.length] = '<' + elmt.tagName;
    for (var i = 0; i < elmt.attributes.length; i++) {
        if (elmt.attributes[i].specified && p.indexOf(elmt.attributes[i].name.toLowerCase()) < 0) {
            h[h.length] = ' ' + elmt.attributes[i].name + '="' + elmt.attributes[i].value + '"';
        }
    }
    //附加的其它属性
    if (arguments[1]) h[h.length] = ' ' + arguments[1];
    h[h.length] = '>' + elmt.innerHTML + '</' + elmt.tagName + '>';
    return h.join('');
}
String.prototype.isDate = function () {
    var str = this;
    var reg = /^(\d{4})(?:-|\/)(\d{1,2})(?:-|\/)(\d{1,2})(?:\s(\d{1,2})\:(\d{1,2})(?:\:(\d{1,2}))?)?$/;
    var m = str.match(reg);
    try {
        var tmp = 'new Date(';
        for (var i = 1; i < m.length; i++) { m[i] = (m[i] == '' || typeof m[i] == 'undefined') ? 0 : i == 2 ? m[i] - 1 : m[i]; tmp += m[i] + ','; }
        tmp = tmp.substr(0, tmp.length - 1) + ');';
        var dt = eval(tmp);
        if (dt.getFullYear() != m[1]) return false;
        if (dt.getMonth() != m[2]) return false;
        if (dt.getDate() != m[3]) return false;
        if (dt.getHours() != m[4]) return false;
        if (dt.getMinutes() != m[5]) return false;
        if (dt.getSeconds() != m[6]) return false;
        return true;
    } catch (e) { return false; }
}


window.setMaskOn = function () {
    var mask = document.createElement('div');
    mask.id = '__JWGL_Mask__';
    mask.style.cssText = 'position:absolute;top:0px;left:0px;right:0px;bottom:0px;overflow:auto;z-index:1000;' +
        'filter:Alpha(Opacity=30);opacity:0.4;-moz-opacity:0.4;background-color:#ffffff;';
    var selObj = document.getElementsByTagName('select');
    if (MSIE) for (var i = 0; i < selObj.length; i++) selObj[i].style.visibility = 'hidden';
    document.body.appendChild(mask);
}
window.setMaskOff = function () {
    var selObj = document.getElementsByTagName('select');
    if (MSIE) for (var i = 0; i < selObj.length; i++) selObj[i].style.visibility = 'visible';
    if ($('__JWGL_Mask__')) document.body.removeChild($('__JWGL_Mask__'));
}