(this.webpackJsonpfrontend=this.webpackJsonpfrontend||[]).push([[0],{25:function(e,t,n){e.exports=n(55)},54:function(e,t,n){},55:function(e,t,n){"use strict";n.r(t);var r=n(0),a=n.n(r),c=n(8),o=n.n(c),s=n(13),u=n(4),l=n(24),i=n.n(l),p=n(3),f=n.n(p),v=n(7),d=n(6),b=n(12),m=n.n(b),h=n(5),E=Object(d.b)({name:"jsonHolder",initialState:{value:{}},reducers:{jsonResponse(e,t){var n=t.payload;h.debug(n),e.value=n}}}),j=E.actions.jsonResponse,k=E.reducer;function O(){return w.apply(this,arguments)}function w(){return(w=Object(v.a)(f.a.mark((function e(){var t,n;return f.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return"http://127.0.0.1:5000/ping",e.next=3,m.a.get("http://127.0.0.1:5000/ping");case 3:return t=e.sent,n=t.data,e.abrupt("return",n);case 6:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function g(){return x.apply(this,arguments)}function x(){return(x=Object(v.a)(f.a.mark((function e(){var t,n;return f.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return"http://127.0.0.1:5000/api/close",e.next=3,m.a.get("http://127.0.0.1:5000/api/close");case 3:return t=e.sent,n=t.data,e.abrupt("return",n);case 6:case"end":return e.stop()}}),e)})))).apply(this,arguments)}var S=Object(r.createContext)(null),y=function(e){var t,n,r=e.children,c=Object(u.b)();return(t=new WebSocket("ws://127.0.0.1:5000/api/ws")).onmessage=function(e){var t=JSON.parse(e.data);c(j(t))},t.onopen=function(){h.warn("Websocket is now open")},t.onclose=function(e){h.warn("Socket Closed Connection: ",e)},t.onerror=function(e){h.error("Socket Error: ",e)},n={webSocket:t,sendMessage:function(e){var n={type:"sendMessage",data:e};t.send(JSON.stringify(n))}},a.a.createElement(S.Provider,{value:n},r)},C=function(){var e=Object(u.b)(),t=Object(u.c)((function(e){return e.jsonHolder})).value,n=Object(r.useState)(void 0),c=Object(s.a)(n,2),o=c[0],l=c[1];Object(r.useEffect)((function(){return function(){void 0!==o&&clearInterval(o)}}),[o]);var p=function(){void 0!==o&&(clearInterval(o),l(void 0))},d=Object(r.useState)("not set"),b=Object(s.a)(d,2),m=b[0],h=b[1],E=Object(r.useContext)(S);return a.a.createElement("div",null,a.a.createElement("div",null,a.a.createElement(i.a,{src:t,theme:"solarized"}),a.a.createElement("button",{onClick:function(){void 0===o&&l(setInterval((function(){e(function(){var e=Object(v.a)(f.a.mark((function e(t){var n;return f.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,e.next=3,O();case 3:n=e.sent,t(j(n)),e.next=10;break;case 7:e.prev=7,e.t0=e.catch(0),e.t0.response?t(j(e.t0.response)):e.t0.request?t(j(e.t0.request)):t(j(e.t0));case 10:case"end":return e.stop()}}),e,null,[[0,7]])})));return function(t){return e.apply(this,arguments)}}())}),1e3))}},"Start"),a.a.createElement("button",{onClick:p},"Stop Refresh")),a.a.createElement("form",{onSubmit:function(e){e.preventDefault(),null===E||void 0===E||E.sendMessage(m)}},a.a.createElement("label",null,"Message Backend"),a.a.createElement("input",{value:m,onChange:function(e){var t=e.target.value;h(t)}}),a.a.createElement("button",{type:"submit"},"Send over websocket")),a.a.createElement("div",null,a.a.createElement("label",null,"To Close Backend server"),a.a.createElement("button",{onClick:function(){p(),e(function(){var e=Object(v.a)(f.a.mark((function e(t){var n;return f.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,e.next=3,g();case 3:n=e.sent,t(j(n)),e.next=10;break;case 7:e.prev=7,e.t0=e.catch(0),e.t0.response?t(j(e.t0.response)):e.t0.request?t(j(e.t0.request)):t(j(e.t0));case 10:case"end":return e.stop()}}),e,null,[[0,7]])})));return function(t){return e.apply(this,arguments)}}())}},"Close")))},I=function(){return a.a.createElement("div",{className:"container"},a.a.createElement("div",{className:"container"},a.a.createElement("h1",null,"Response"),a.a.createElement("p",null,"API Response for reading"),a.a.createElement(C,null)))},R=(n(54),n(2)),q=Object(R.c)({jsonHolder:k}),J=Object(d.a)({reducer:q});o.a.render(a.a.createElement(u.a,{store:J},a.a.createElement(y,null,a.a.createElement(I,null))),document.getElementById("root"))}},[[25,1,2]]]);
//# sourceMappingURL=main.1a073ccb.chunk.js.map