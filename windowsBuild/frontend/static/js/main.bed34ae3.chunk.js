(this.webpackJsonpfrontend=this.webpackJsonpfrontend||[]).push([[0],{25:function(e,n,t){e.exports=t(55)},54:function(e,n,t){},55:function(e,n,t){"use strict";t.r(n);var a=t(0),r=t.n(a),c=t(8),o=t.n(c),s=t(24),u=t(5),l=t(9),i=t.n(l),p=t(4),d=t.n(p),f=t(6),b=t(7),v=t(13),m=t.n(v),g=t(1),E=t.n(g),y=Object(b.b)({name:"jsonHolder",initialState:{value:{},config:{}},reducers:{jsonResponse(e,n){var t=n.payload;g.debug(t),e.value=t},configResponse(e,n){var t=n.payload;g.debug(t),e.config=t}}}),w=y.actions,j=w.jsonResponse,h=w.configResponse,k=y.reducer;function O(){return x.apply(this,arguments)}function x(){return(x=Object(f.a)(d.a.mark((function e(){var n,t;return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return"http://127.0.0.1:5001/api/close",e.next=3,m.a.get("http://127.0.0.1:5001/api/close");case 3:return n=e.sent,t=n.data,e.abrupt("return",t);case 6:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function C(e,n){return S.apply(this,arguments)}function S(){return(S=Object(f.a)(d.a.mark((function e(n,t){var a;return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,e.next=3,n();case 3:a=e.sent,t(j(a)),e.next=10;break;case 7:e.prev=7,e.t0=e.catch(0),e.t0.response?t(j(e.t0.response)):e.t0.request?t(j(e.t0.request)):t(j(e.t0));case 10:case"end":return e.stop()}}),e,null,[[0,7]])})))).apply(this,arguments)}var R=Object(a.createContext)(null),_=function(e){var n,t,a=e.children,c=Object(u.b)();return(n=new WebSocket("ws://127.0.0.1:4999/api/ws")).onmessage=function(e){!function(e,n){var t=JSON.parse(e),a=t;if(void 0!==a)switch(a.type){case"display":E.a.debug(a.type+" : "+a.body),n(j(a));break;case"config":n(h(a.body));break;case"ping_pong":case"system":default:E.a.debug(a.type+" : "+a.body),n(j(a))}else n(j(t))}(e.data,c)},n.onopen=function(){g.warn("Websocket is now open")},n.onclose=function(e){g.warn("Socket Closed Connection: ",e)},n.onerror=function(e){g.error("Socket Error: ",e)},t={webSocket:n,sendMessage:function(e,t){!function(e,n,t){var a={type:n,body:t};e.send(JSON.stringify(a))}(n,e,t)}},r.a.createElement(R.Provider,{value:t},a)},M=function(){var e=Object(u.b)(),n=Object(u.c)((function(e){return e.jsonHolder})).value,t=Object(a.useState)("not set"),c=Object(s.a)(t,2),o=c[0],l=c[1],p=Object(a.useContext)(R);return r.a.createElement("div",null,r.a.createElement("div",null,r.a.createElement(i.a,{src:n,theme:"solarized"}),r.a.createElement("button",{onClick:function(){null===p||void 0===p||p.sendMessage("data",{})}},"Get Data")),r.a.createElement("form",{onSubmit:function(e){e.preventDefault(),null===p||void 0===p||p.sendMessage("user",o)}},r.a.createElement("label",null,"Message Backend"),r.a.createElement("input",{value:o,onChange:function(e){var n=e.target.value;l(n)}}),r.a.createElement("button",{type:"submit"},"Send over websocket")),r.a.createElement("div",null,r.a.createElement("label",null,"To Close Backend server"),r.a.createElement("button",{onClick:function(){e(function(){var e=Object(f.a)(d.a.mark((function e(n){return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,C(O,n);case 2:case"end":return e.stop()}}),e)})));return function(n){return e.apply(this,arguments)}}())}},"Close")))},H=function(){var e=Object(u.c)((function(e){return e.jsonHolder})).config,n=Object(a.useContext)(R);return r.a.createElement("div",null,r.a.createElement(i.a,{src:e,theme:"solarized",onEdit:function(e){if(typeof e.existing_value!==typeof e.new_value)return E.a.warn(e),!1;E.a.warn(typeof e.existing_value),E.a.warn(typeof e.new_value),E.a.warn(e),null===n||void 0===n||n.sendMessage("config",e.updated_src)}}))},J=function(){return r.a.createElement("div",{className:"container"},r.a.createElement("div",{className:"container"},r.a.createElement("h1",null,"Response"),r.a.createElement("p",null,"API Response for reading"),r.a.createElement(M,null),r.a.createElement("p",null,"App Config"),r.a.createElement(H,null)))},N=(t(54),t(3)),B=Object(N.c)({jsonHolder:k}),q=Object(b.a)({reducer:B});o.a.render(r.a.createElement(u.a,{store:q},r.a.createElement(_,null,r.a.createElement(J,null))),document.getElementById("root"))}},[[25,1,2]]]);
//# sourceMappingURL=main.bed34ae3.chunk.js.map