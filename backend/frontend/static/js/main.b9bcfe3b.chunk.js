(this.webpackJsonpfrontend=this.webpackJsonpfrontend||[]).push([[6],{10:function(e,n,t){"use strict";t.d(n,"d",(function(){return d})),t.d(n,"b",(function(){return g})),t.d(n,"a",(function(){return m}));var r=t(4),a=t.n(r),o=t(11),c=t(18),s=t(31),i=t.n(s),u=t(6),l=t(17),p=Object(c.b)({name:"jsonHolder",initialState:{value:{},config:{}},reducers:{jsonResponse(e,n){var t=n.payload;u.debug(t),e.value=t},configResponse(e,n){var t=n.payload;u.debug(t),e.config=t}}}),f=p.actions,d=f.jsonResponse,g=f.configResponse;n.c=p.reducer;var m=function(){return function(){var e=Object(o.a)(a.a.mark((function e(n){return a.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,O(v,n);case 2:case"end":return e.stop()}}),e)})));return function(n){return e.apply(this,arguments)}}()};function v(){return b.apply(this,arguments)}function b(){return(b=Object(o.a)(a.a.mark((function e(){var n,t,r;return a.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return n="http://127.0.0.1:"+l.a+"/api/close",e.next=3,i.a.get(n);case 3:return t=e.sent,r=t.data,e.abrupt("return",r);case 6:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function O(e,n){return E.apply(this,arguments)}function E(){return(E=Object(o.a)(a.a.mark((function e(n,t){var r;return a.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,e.next=3,n();case 3:r=e.sent,t(d(r)),e.next=10;break;case 7:e.prev=7,e.t0=e.catch(0),e.t0.response?t(d(e.t0.response)):e.t0.request?t(d(e.t0.request)):t(d(e.t0));case 10:case"end":return e.stop()}}),e,null,[[0,7]])})))).apply(this,arguments)}},17:function(e,n,t){"use strict";t.d(n,"a",(function(){return r}));var r=4999},33:function(e,n,t){"use strict";t.d(n,"a",(function(){return l}));var r=t(0),a=t.n(r),o=t(5),c=t(6),s=t.n(c),i=t(10),u=t(17),l=Object(r.createContext)(null);n.b=function(e){var n,t,r=e.children,p=Object(o.useDispatch)();return(n=new WebSocket("ws://127.0.0.1:"+u.a+"/api/ws")).onmessage=function(e){!function(e,n){var t=JSON.parse(e),r=t;if(void 0!==r)switch(r.type){case"display":s.a.debug(r.type+" : "+r.body),n(Object(i.d)(r));break;case"config":n(Object(i.b)(r.body));break;case"ping_pong":case"system":default:s.a.debug(r.type+" : "+r.body),n(Object(i.d)(r))}else n(Object(i.d)(t))}(e.data,p)},n.onopen=function(){c.warn("Websocket is now open")},n.onclose=function(e){c.warn("Socket Closed Connection: ",e)},n.onerror=function(e){c.error("Socket Error: ",e)},t={webSocket:n,sendMessage:function(e,t){!function(e,n,t){var r={type:n,body:t};e.send(JSON.stringify(r))}(n,e,t)}},a.a.createElement(l.Provider,{value:t},r)}},51:function(e,n,t){e.exports=t(95)},70:function(e,n,t){var r={".":13,"./":13,"./index":13,"./index.js":13};function a(e){var n=o(e);return t(n)}function o(e){if(!t.o(r,e)){var n=new Error("Cannot find module '"+e+"'");throw n.code="MODULE_NOT_FOUND",n}return r[e]}a.keys=function(){return Object.keys(r)},a.resolve=o,e.exports=a,a.id=70},74:function(e,n,t){var r={"./pages":[9,0],"./pages/":[9,0],"./pages/config":[22,1,2],"./pages/config.tsx":[22,1,2],"./pages/home":[21,3],"./pages/home.tsx":[21,3],"./pages/index":[9,0],"./pages/index.tsx":[9,0],"./pages/response":[23,1,4],"./pages/response.tsx":[23,1,4],"./pages/system":[24,5],"./pages/system.tsx":[24,5]};function a(e){if(!t.o(r,e))return Promise.resolve().then((function(){var n=new Error("Cannot find module '"+e+"'");throw n.code="MODULE_NOT_FOUND",n}));var n=r[e],a=n[0];return Promise.all(n.slice(1).map(t.e)).then((function(){return t(a)}))}a.keys=function(){return Object.keys(r)},a.id=74,e.exports=a},75:function(e,n,t){var r={"./pages":9,"./pages/":9,"./pages/config":22,"./pages/config.tsx":22,"./pages/home":21,"./pages/home.tsx":21,"./pages/index":9,"./pages/index.tsx":9,"./pages/response":23,"./pages/response.tsx":23,"./pages/system":24,"./pages/system.tsx":24};function a(e){var n=o(e);if(!t.m[n]){var r=new Error("Module '"+e+"' ('"+n+"') is not available (weak dependency)");throw r.code="MODULE_NOT_FOUND",r}return t(n)}function o(e){if(!t.o(r,e)){var n=new Error("Cannot find module '"+e+"'");throw n.code="MODULE_NOT_FOUND",n}return r[e]}a.keys=function(){return Object.keys(r)},a.resolve=o,a.id=75,e.exports=a},76:function(e,n,t){},95:function(e,n,t){"use strict";t.r(n);var r=t(0),a=t.n(r),o=t(20),c=t.n(o),s=t(19),i={HOME:"/",RESPONSE:"/response",CONFIG:"/config",SYSTEM:"/system"},u=function(){return a.a.createElement("nav",null,a.a.createElement("ul",null,a.a.createElement("li",null,a.a.createElement(s.NavLink,{to:i.HOME},"Home")),a.a.createElement("li",null,a.a.createElement(s.NavLink,{to:i.CONFIG},"Config")),a.a.createElement("li",null,a.a.createElement(s.NavLink,{to:i.RESPONSE},"Response")),a.a.createElement("li",{className:"float-right"},a.a.createElement(s.NavLink,{to:i.SYSTEM},"System"))))},l=t(29),p=t.n(l),f=t(5),d=p()((function(e){return Promise.all([t(74)("./pages".concat(e.page))]).then((function(e){return e[0]}))}),{chunkName:function(e){return e.page},resolve:function(e){return t(75).resolve("./pages".concat(e.page))}}),g=function(){var e=Object(f.useSelector)((function(e){return e.location})),n=e.type,t=e.routesMap;return a.a.createElement(d,{page:t[n]})},m=function(){return a.a.createElement("div",{className:"zeroMargin"},a.a.createElement(u,null),a.a.createElement("div",{className:"bodyCentering"},a.a.createElement(g,null)))},v=(t(76),t(32)),b=t(3),O=t(18),E=t(8),y=t(10),h=Object(E.connectRoutes)(i),w=h.middleware,x=h.enhancer,k=h.reducer,j=Object(b.c)({jsonHolder:y.c,location:k}),N=Object(O.a)({reducer:j,enhancers:function(e){return[x].concat(Object(v.a)(e))},middleware:[].concat(Object(v.a)(Object(O.c)({immutableCheck:{warnAfter:1e3},serializableCheck:{warnAfter:1e3}})),[w])}),S=t(33);c.a.render(a.a.createElement(f.Provider,{store:N},a.a.createElement(S.b,null,a.a.createElement(m,null))),document.getElementById("root"))}},[[51,7,8]]]);
//# sourceMappingURL=main.b9bcfe3b.chunk.js.map