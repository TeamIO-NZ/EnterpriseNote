(this["webpackJsonpnote-ui"]=this["webpackJsonpnote-ui"]||[]).push([[0],{582:function(e,t,n){"use strict";n.r(t);var a=n(2),r=n(0),i=n.n(r),c=n(16),s=n.n(c),o=n(8),l=n(617),u=n(21),d=n(205),j=n(206),h=n(612),b=n(613),f=n(120),g=n(187),O=n(284),p=n.n(O),m=n(615),x=n(209),v=n(297),w=n(591),y=n(42),S=function e(t,n,a,r,i,c,s,o){Object(y.a)(this,e),this.name=void 0,this.password=void 0,this.loggedIn=void 0,this.token=void 0,this.email=void 0,this.gender=void 0,this.id=void 0,this.userSettings=void 0,this.name=t,this.password=n,this.loggedIn=a,this.token=r,this.email=i,this.gender=c,this.id=s,this.userSettings=o},N=Object(r.createContext)({user:new S("","",!1,"","","",0,-1),setUser:function(e){return console.warn("No User Provider")}}),C=function(){return Object(r.useContext)(N)},k=n(161),E=Object(d.a)((function(e){return Object(j.a)({root:{flexGrow:1},menuButton:{marginRight:e.spacing(2)},title:{flexGrow:1},avatar:{backgroundColor:k.a[100],color:k.a[600]}})}));function I(){var e=C(),t=e.user,n=e.setUser,r=Object(u.g)(),c=E(),s=i.a.useState(null),l=Object(o.a)(s,2),d=l[0],j=l[1],O=Boolean(d),y=Object(u.h)(),N="/login"===y.pathname||"/register"===y.pathname,k=function(){j(null)};return Object(a.jsx)("div",{className:c.root,children:Object(a.jsx)(h.a,{position:"static",children:Object(a.jsxs)(b.a,{children:[Object(a.jsx)(g.a,{edge:"start",className:c.menuButton,color:"inherit","aria-label":"menu",children:Object(a.jsx)(p.a,{})}),Object(a.jsx)(f.a,{variant:"h6",className:c.title,children:"Enterprise Note"}),""!==t.token&&Object(a.jsxs)("div",{children:[Object(a.jsx)(w.a,{"aria-label":"account of ".concat(t.name),"aria-controls":"menu-appbar",onClick:function(e){j(e.currentTarget)},className:c.avatar,children:t.name.charAt(0).toUpperCase()}),Object(a.jsxs)(v.a,{id:"menu-appbar",anchorEl:d,anchorOrigin:{vertical:"top",horizontal:"right"},keepMounted:!0,transformOrigin:{vertical:"top",horizontal:"right"},open:O,onClose:k,children:[Object(a.jsx)(m.a,{onClick:function(){r.push("/"),k()},children:"My Notes"}),Object(a.jsx)(m.a,{onClick:function(){r.push("/users"),k()},children:"UserList"}),Object(a.jsx)(m.a,{onClick:function(){n(new S("","",!1,"","","",-1,-1)),r.push("/login"),k()},children:"Logout"})]})]}),!N&&""===t.token&&Object(a.jsx)(x.a,{color:"inherit",onClick:function(){r.push("/login")}&&k,children:"Login"})]})})})}var T=n(283),R=n(41),D=n(623),L=n(159),P=n.n(L),z=n(616),A=n(98),U=n.n(A),B=n(158),W=n(71),V=n(285),F=n.n(V).a.create({baseURL:"http://192.168.0.64:8082/api/v1/",headers:{"Content-type":"application/json"}}),H=new(function(){function e(){Object(y.a)(this,e)}return Object(W.a)(e,[{key:"getAll",value:function(){return F.get("/users")}},{key:"get",value:function(e){return F.get("/user/".concat(e))}},{key:"getByName",value:function(e){return F.get("/user/name/".concat(e))}},{key:"getByEmail",value:function(e){return F.get("/user/email/".concat(e))}},{key:"create",value:function(e){return console.log("Create me:"),console.log(e),F.post("/user",e)}},{key:"update",value:function(e,t){return F.put("/user/".concat(e),t)}},{key:"delete",value:function(e){return F.delete("/user/".concat(e))}},{key:"login",value:function(e,t){return F.get("/login/".concat(e,"/").concat(t))}},{key:"getData",value:function(){var e=Object(B.a)(U.a.mark((function e(){return U.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return console.log("Getting Data"),e.next=3,this.getAll().then((function(e){return console.log(e),e.data}));case 3:return e.abrupt("return",e.sent);case 4:case"end":return e.stop()}}),e,this)})));return function(){return e.apply(this,arguments)}}()}]),e}()),q=n(622),J=n(218);function G(){return Object(a.jsxs)(f.a,{variant:"body2",color:"textSecondary",align:"center",children:["Copyright \xa9 ",Object(a.jsx)(R.b,{color:"inherit",to:"https://github.com/TeamIO-NZ/EnterpriseNote",children:"Enterprise Note"})," ",(new Date).getFullYear(),"."]})}var Z=n(72),$=Object(d.a)((function(e){return{paper:{marginTop:e.spacing(8),display:"flex",flexDirection:"column",alignItems:"center"},avatar:{margin:e.spacing(1),backgroundColor:e.palette.secondary.main},form:{width:"100%",marginTop:e.spacing(1)},submit:{margin:e.spacing(3,0,2)}}}));function M(){var e=$(),t=i.a.useState(!1),n=Object(o.a)(t,2),r=n[0],c=n[1],s=i.a.useState(!1),d=Object(o.a)(s,2),j=d[0],h=d[1],b=i.a.useState(!1),g=Object(o.a)(b,2),O=g[0],p=g[1],m=C(),v=m.user,y=m.setUser,N=Object(u.g)(),k=Object(Z.a)();-1!==v.id&&k()&&N.push("/"),console.log("Before Handles");var E=function(e,t){H.login(e,t).then((function(n){if(k()){p(!1);var a=n.data.data.data;if(a){var r=a.token;console.log(r),atob(a.token)===e+t&&(y(new S(a.name,a.password,!0,a.token,a.email,a.gender,a.userId,a.usersettings)),N.push("/"))}else c(!0),h(!0)}}))};return Object(a.jsxs)(z.a,{component:"main",maxWidth:"xs",children:[Object(a.jsx)(l.a,{}),Object(a.jsxs)("div",{className:e.paper,children:[Object(a.jsx)(w.a,{className:e.avatar,children:Object(a.jsx)(P.a,{})}),Object(a.jsx)(f.a,{component:"h1",variant:"h5",children:"Login"}),Object(a.jsxs)("form",{className:e.form,noValidate:!0,children:[Object(a.jsx)(T.a,{variant:"outlined",margin:"normal",required:!0,fullWidth:!0,id:"username",label:"Username",name:"username",autoComplete:"username",error:r,onChange:function(e){return function(e){console.log("handleUsername"),/\s/g.test(e.target.value)?c(!0):(c(!1),v.name=e.target.value.toLowerCase())}(e)},disabled:O,autoFocus:!0}),Object(a.jsx)(T.a,{variant:"outlined",margin:"normal",required:!0,fullWidth:!0,name:"password",label:"Password",type:"password",id:"password",error:j,autoComplete:"current-password",onChange:function(e){return function(e){console.log("HandlePassword"),""===e.target.value?(h(!0),console.log("invalid password")):(h(!1),v.password=e.target.value)}(e)},disabled:O}),Object(a.jsx)(x.a,{type:"submit",fullWidth:!0,variant:"contained",color:"primary",disabled:r&&j||O,className:e.submit,onClick:function(e){return function(e){e.preventDefault(),E(v.name,v.password)}(e)},children:O?Object(a.jsx)(q.a,{size:25}):"Login"})]}),Object(a.jsxs)(J.a,{container:!0,children:[Object(a.jsx)(J.a,{item:!0,xs:!0}),Object(a.jsx)(J.a,{item:!0,children:Object(a.jsx)(R.b,{to:"/register",children:"Don't have an account? Register"})})]})]}),Object(a.jsx)(D.a,{mt:8,children:Object(a.jsx)(G,{})})]})}var Y=function e(t,n,a,r,i,c,s){Object(y.a)(this,e),this.id=void 0,this.title=void 0,this.desc=void 0,this.content=void 0,this.owner=void 0,this.viewers=void 0,this.editors=void 0,this.id=t,this.title=n,this.desc=a,this.content=r,this.owner=i,this.viewers=c,this.editors=s},_=n(644),K=n(286),Q=n(119),X=n(17),ee=n(648),te=n(288),ne=n.n(te),ae=n(626),re=n(300),ie=n(287),ce=n.n(ie),se=new(function(){function e(){Object(y.a)(this,e)}return Object(W.a)(e,[{key:"getAll",value:function(){return F.get("/notes")}},{key:"getAllUserHasAccessTo",value:function(e){return F.get("/usersnotes/".concat(e))}},{key:"get",value:function(e){return F.get("/note/".concat(e))}},{key:"create",value:function(e){return F.post("/note",e)}},{key:"update",value:function(e,t,n){return F.put("/note/".concat(e,"/").concat(n),t)}},{key:"delete",value:function(e){return console.log("Deleting note"),F.delete("/note/".concat(e))}},{key:"getData",value:function(){var e=Object(B.a)(U.a.mark((function e(t){return U.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.getAllUserHasAccessTo(t).then((function(e){return e.data}));case 2:return e.abrupt("return",e.sent);case 3:case"end":return e.stop()}}),e,this)})));return function(t){return e.apply(this,arguments)}}()}]),e}()),oe=n(57),le=n(627),ue=n(628),de=Object(d.a)((function(e){return Object(j.a)({appBar:{position:"relative"},title:{marginLeft:e.spacing(2),flex:1},editor:{margin:40},search:Object(Q.a)({position:"relative",borderRadius:e.shape.borderRadius,backgroundColor:Object(X.d)(e.palette.common.white,.15),"&:hover":{backgroundColor:Object(X.d)(e.palette.common.white,.25)},margin:3,width:"100%"},e.breakpoints.up("sm"),{marginLeft:e.spacing(3),width:"auto"}),searchIcon:{padding:e.spacing(0,2),height:"100%",position:"absolute",pointerEvents:"none",display:"flex",alignItems:"center",justifyContent:"center"},inputRoot:{color:"inherit"},inputInput:Object(Q.a)({padding:e.spacing(1,1,1,0),paddingLeft:"calc(1em + ".concat(e.spacing(4),"px)"),transition:e.transitions.create("width"),width:"100%"},e.breakpoints.up("md"),{width:"20ch"})})})),je=i.a.forwardRef((function(e,t){return Object(a.jsx)(ae.a,Object(K.a)({direction:"up",ref:t},e))}));function he(e){var t=e.open,n=e.setOpen,c=e.note,s=e.setShouldRefresh,l=de(),u=i.a.useState(c.content),d=Object(o.a)(u,2),j=d[0],O=d[1],p=i.a.useState(c.title),m=Object(o.a)(p,2),v=m[0],w=m[1],y=i.a.useState(c.desc),S=Object(o.a)(y,2),N=S[0],k=S[1],E=i.a.useState(JSON.stringify(Object(oe.convertToRaw)(oe.EditorState.createEmpty().getCurrentContent()))),I=Object(o.a)(E,2),T=I[0],R=I[1],D=C(),L=D.user,P=(D.setUser,function(){s(String(Date.now())),n(!1)});Object(r.useEffect)((function(){""==c.content?(console.log("fuckin empty"),c.content=btoa(JSON.stringify(Object(oe.convertToRaw)(oe.EditorState.createEmpty().getCurrentContent())))):console.log("not empty"),R(atob(c.content)),w(c.title),k(c.desc)}),[c]);var z=function(){c.title=v,c.desc=N,c.content=j,"number"===typeof c.id&&c.id>0?se.update(c.id,c,L.id).then((function(){s(String(Date.now()))})):se.create(c).then((function(e){console.log("create: ".concat(e))})).then((function(){s(String(Date.now()))}))};return Object(a.jsx)("div",{children:Object(a.jsxs)(ee.a,{fullScreen:!0,open:t,onClose:P,TransitionComponent:je,children:[Object(a.jsx)(h.a,{className:l.appBar,children:Object(a.jsxs)(b.a,{children:[Object(a.jsx)(g.a,{edge:"start",color:"inherit",onClick:function(){return P()},"aria-label":"close",children:Object(a.jsx)(ne.a,{})}),Object(a.jsxs)(f.a,{variant:"h6",className:l.title,children:[Object(a.jsxs)("div",{className:l.search,children:[Object(a.jsx)("div",{className:l.searchIcon,children:Object(a.jsx)(le.a,{})}),Object(a.jsx)(re.a,{placeholder:"Title",classes:{root:l.inputRoot,input:l.inputInput},inputProps:{"aria-label":"title"},onChange:function(e){w(e.target.value)},defaultValue:c.title})]}),Object(a.jsxs)("div",{className:l.search,children:[Object(a.jsx)("div",{className:l.searchIcon,children:Object(a.jsx)(ue.a,{})}),Object(a.jsx)(re.a,{placeholder:"Description",classes:{root:l.inputRoot,input:l.inputInput},inputProps:{"aria-label":"description"},onChange:function(e){k(e.target.value)},defaultValue:c.desc})]})]}),Object(a.jsx)(x.a,{color:"inherit",onClick:function(){return z(),void P()},children:"save"})]})}),Object(a.jsx)("div",{className:l.editor,children:Object(a.jsx)(ce.a,{inlineToolbar:!0,onChange:function(e){O(btoa(JSON.stringify(Object(oe.convertToRaw)(e.getCurrentContent()))))},label:"Type something here...",onSave:z,defaultValue:T})})]})})}var be=n(651),fe=n(635),ge=n(637),Oe=n(649),pe=n(640),me=n(641),xe=n(208),ve=n(188),we=n(636),ye=n(638),Se=n(639),Ne=n(296),Ce=n(289),ke=n.n(Ce),Ee=n(629),Ie=Object(d.a)((function(e){return Object(j.a)({avatar:{backgroundColor:k.a[100],color:k.a[600]}})}));function Te(e){var t=e.note,n=Ie(),c=i.a.useState([]),s=Object(o.a)(c,2),l=s[0],u=s[1],d=i.a.useState([]),j=Object(o.a)(d,2),h=j[0],b=j[1],g=i.a.useState([]),O=Object(o.a)(g,2),p=O[0],m=O[1],x=Object(Z.a)(),v=function(e){return e.length>0?Object(a.jsx)(Ee.a,{max:4,classes:{avatar:n.avatar},children:e.map((function(e){return i.a.cloneElement(Object(a.jsx)(Oe.a,{title:e.name,children:Object(a.jsx)(w.a,{alt:e.name,className:n.avatar,children:e.name.charAt(0).toUpperCase().toString()})}),{key:e.id})}))}):null};return Object(r.useEffect)((function(){if(H.get(t.owner).then((function(e){console.log(e);var t=[];if(void 0===e.data||null===e.data)return[];var n=e.data.data;return t.push(new S(n.name,"",!1,"","","",n.userId,-1)),t})).then((function(e){u(e)})),null!=t.editors&&t.editors.length>0){var e=Object.assign([],h);b([]),t.editors.forEach((function(t){0!==t&&H.get(t).then((function(t){if(x()){var n=t.data.data;e.push(new S(n.name,"",!1,"","","",n.userId,-1)),b(e)}}))}))}if(null!=t.viewers&&t.viewers.length>0){var n=Object.assign([],p);m([]),t.viewers.forEach((function(e){0!==e&&H.get(e).then((function(e){if(x()){var t=e.data.data;n.push(new S(t.name,"",!1,"","","",t.userId,-1)),m(n)}}))}))}}),[t]),Object(a.jsxs)(J.a,{container:!0,spacing:3,alignItems:"stretch",justify:"flex-start",children:[Object(a.jsxs)(J.a,{item:!0,xs:3,sm:3,children:[Object(a.jsx)(f.a,{children:"Owner"}),v(l)]}),Object(a.jsxs)(J.a,{item:!0,xs:3,sm:3,children:[Object(a.jsx)(f.a,{children:"Editors"}),v(h)]}),Object(a.jsxs)(J.a,{item:!0,xs:3,sm:3,children:[Object(a.jsx)(f.a,{children:"Viewer"}),v(p)]})]})}var Re,De=n(291),Le=n.n(De),Pe=n(292),ze=n.n(Pe),Ae=n(631),Ue=n(633),Be=n(630),We=n(634),Ve=n(645),Fe=n(632),He=new(function(){function e(){Object(y.a)(this,e)}return Object(W.a)(e,[{key:"get",value:function(e){return F.get("/usersettings/".concat(e))}},{key:"create",value:function(e){return F.post("/usersettings",e)}},{key:"update",value:function(e,t){return F.put("/usersettings/".concat(e),t)}},{key:"delete",value:function(e){return console.log("Deleting note"),F.delete("/usersettings/".concat(e))}}]),e}()),qe=Object(d.a)((function(e){return Object(j.a)({avatar:{backgroundColor:k.a[100],color:k.a[600]},actions:{marginLeft:40},listText:{marginRight:250},formControl:{margin:e.spacing(1),minWidth:120,maxWidth:300}})}));function Je(e){var t=qe(),n=e.onClose,c=e.open,s=e.note,l=e.setNote,u=e.requestRefresh,d=i.a.useState([]),j=Object(o.a)(d,2),h=j[0],b=j[1],g=Object(Z.a)(),O=i.a.useState("fucking work m8"),p=Object(o.a)(O,2),m=p[0],v=p[1];null==s.editors&&(s.editors=[]),null==s.viewers&&(s.viewers=[]),Object(r.useEffect)((function(){g()&&H.getAll().then((function(e){if(console.log(e),g()){b([]);var t=[];e.data.data.forEach((function(e){t.push({name:e.name,userId:e.userId,role:y(e.userId)})})),console.log(t),b(t)}}))}),[s,m]);var y=function(e){return s.owner==e?Re.Owner:s.editors.includes(e)?Re.Editor:s.viewers.includes(e)?Re.Viewer:Re.None},S=function(e,t){if(s.owner!==t.userId){var n=Object.assign([],h);N(n[n.indexOf(t)]),e.target.value==Re.None?n[n.indexOf(t)].role=Re.None:e.target.value==Re.Viewer?(n[n.indexOf(t)].role=Re.Viewer,s.viewers.push(t.userId)):e.target.value==Re.Editor&&(n[n.indexOf(t)].role=Re.Editor,s.editors.push(t.userId)),b(n)}},N=function(e){s.viewers=s.viewers.filter((function(t){return t!==e.userId})),s.editors=s.editors.filter((function(t){return t!==e.userId}))};return Object(a.jsxs)(ee.a,{onClose:function(){console.log(s),se.update(s.id,s,s.owner).then((function(){u(String(Date.now())),n()}))},"aria-labelledby":"share-dialog-title",open:c,children:[Object(a.jsxs)(Be.a,{id:"share-dialog-title",children:["Share Note",Object(a.jsx)(x.a,{className:t.actions,onClick:function(){He.get(s.owner).then((function(e){return console.log(e),400==e.data.code})).then((function(e){e?He.create({editors:s.editors,viewers:s.viewers,id:s.owner}):He.update(s.owner,{editors:s.editors,viewers:s.viewers})}))},children:"Save Preset"}),Object(a.jsx)(x.a,{onClick:function(){He.get(s.owner).then((function(e){if(console.log(e),200===e.data.code){void 0==e.data.data.editors&&(e.data.data.editors=[]),void 0==e.data.data.viewers&&(e.data.data.viewers=[]);var t=Object.assign({},s);t.editors=e.data.data.editors,t.viewers=e.data.data.viewers,l(t),v(String(Date.now()))}}))},children:"Load Preset"})]}),Object(a.jsxs)(xe.a,{children:[Object(a.jsxs)(ve.a,{children:[Object(a.jsx)(Ae.a,{children:Object(a.jsx)(w.a,{className:t.avatar,children:Object(a.jsx)(Fe.a,{})})}),Object(a.jsx)(Ue.a,{className:t.listText,children:Object(a.jsx)("b",{children:"Users"})}),Object(a.jsx)(We.a,{children:Object(a.jsx)(f.a,{children:Object(a.jsx)("b",{children:"Roles"})})})]},"fuck"),h.map((function(e){return Object(a.jsxs)(ve.a,{children:[Object(a.jsx)(Ae.a,{children:Object(a.jsx)(w.a,{className:t.avatar,children:e.name.charAt(0).toUpperCase()})}),Object(a.jsx)(Ue.a,{primary:e.name,className:t.listText}),Object(a.jsxs)(We.a,{children:[Object(a.jsx)(Oe.a,{title:Re.None,children:Object(a.jsx)(Ve.a,{checked:e.role==Re.None,onChange:function(t){return S(t,e)},value:Re.None,color:"default",name:"radio-button-demo",inputProps:{"aria-label":Re.None},disabled:e.userId===s.owner,size:"small"})}),Object(a.jsx)(Oe.a,{title:Re.Viewer,children:Object(a.jsx)(Ve.a,{checked:e.role==Re.Viewer,onChange:function(t){return S(t,e)},value:Re.Viewer,color:"default",name:"radio-button-demo",inputProps:{"aria-label":Re.Viewer},disabled:e.userId===s.owner,size:"small"})}),Object(a.jsx)(Oe.a,{title:Re.Editor,children:Object(a.jsx)(Ve.a,{checked:e.role==Re.Editor,onChange:function(t){return S(t,e)},value:Re.Editor,color:"default",name:"radio-button-demo",inputProps:{"aria-label":Re.Editor},disabled:e.userId===s.owner,size:"small"})})]})]},"".concat(e.userId,"-").concat(e.name))}))]})]})}!function(e){e.None="None",e.Viewer="Viewer",e.Editor="Editor",e.Owner="Owner"}(Re||(Re={}));var Ge=Object(d.a)((function(e){return Object(j.a)({root:{flexGrow:1,marginLeft:40},title:{margin:e.spacing(4,0,2)},heading:{fontSize:e.typography.pxToRem(15),flexBasis:"33.33%",flexShrink:0},secondaryHeading:{fontSize:e.typography.pxToRem(15),color:e.palette.text.secondary},item:{width:"100%"}})}));function Ze(e,t,n){var r=n.expanded,c=n.handleExpand,s=n.classes,o=n.handleOpenEditor,l=n.deleteNote,u=n.user,d=n.handleShare;return t.map((function(t,n){return i.a.cloneElement(e,{key:n},Object(a.jsxs)(be.a,{className:s.item,expanded:r==="panel".concat(n),onChange:c("panel".concat(n)),children:[Object(a.jsxs)(fe.a,{expandIcon:Object(a.jsx)(we.a,{}),"aria-controls":"panel".concat(n,"bh-content"),id:"panel".concat(n,"bh-header"),children:[Object(a.jsx)(f.a,{className:s.heading,children:t.title}),Object(a.jsx)(f.a,{className:s.secondaryHeading,children:t.desc})]}),Object(a.jsxs)(ge.a,{children:[Object(a.jsx)(Te,{note:t}),t.owner===u.id&&Object(a.jsx)(a.Fragment,{children:Object(a.jsx)(Oe.a,{title:"Share Note",children:Object(a.jsx)(g.a,{size:"small",onClick:function(){return d(t)},children:Object(a.jsx)(ze.a,{})})})}),(t.owner===u.id||t.editors.includes(u.id))&&Object(a.jsxs)(a.Fragment,{children:[Object(a.jsx)(Oe.a,{title:"Delete Note",children:Object(a.jsx)(g.a,{size:"small",onClick:function(){return l(t.id)},children:Object(a.jsx)(ye.a,{})})}),Object(a.jsx)(Oe.a,{title:"Edit Note",children:Object(a.jsx)(g.a,{size:"small",onClick:function(){return o(n)},children:Object(a.jsx)(Se.a,{})})})]})]}),Object(a.jsx)(pe.a,{}),Object(a.jsx)(me.a,{children:Object(a.jsx)(ke.a,{plugins:[Le.a],children:Object(Ne.a)(JSON.parse(atob(t.content)))})})]}))}))}function $e(e){var t=e.shouldRefresh,n=e.setShouldRefresh,c=e.notes,s=e.setNotes,l=e.setActiveNote,u=e.setEditorOpen,d=C().user,j=Ge(),h=i.a.useState(!1),b=Object(o.a)(h,1)[0],f=i.a.useState(!1),g=Object(o.a)(f,2),O=g[0],p=g[1],m=Object(Z.a)(),x=i.a.useState(!1),v=Object(o.a)(x,2),w=v[0],y=v[1],S=i.a.useState(!1),N=Object(o.a)(S,2),k=N[0],E=N[1],I=i.a.useState(new Y(0,"","","",0,[],[])),T=Object(o.a)(I,2),R=T[0],D=T[1];return Object(r.useEffect)((function(){console.log("list refresh"),y(!1),se.getData(d.id).then((function(e){var t=[];return void 0===e.data||null===e.data?[]:(e.data.forEach((function(e){console.log(e.id),t.push(new Y(e.id,e.title,e.desc,e.content,e.owner,e.viewers,e.editors))})),t)})).then((function(e){m()&&(s(e),y(!0))}))}),[d.id,m,t]),w?Object(a.jsxs)(a.Fragment,{children:[Object(a.jsx)(xe.a,{dense:b,children:Ze(Object(a.jsx)(ve.a,{}),c,{expanded:O,handleExpand:function(e){return function(t,n){p(!!n&&e)}},classes:j,handleOpenEditor:function(e){l(c[e]),u(!0)},deleteNote:function(e){console.log("Click"),se.delete(e).then((function(e){console.log(e),n(String(Date.now()))}))},user:d,handleShare:function(e){D(e),E(!0)}})}),Object(a.jsx)(Je,{open:k,onClose:function(){return E(!1)},note:R,setNote:D,requestRefresh:n})]}):Object(a.jsx)("div",{style:{display:"flex",justifyContent:"center",alignItems:"center",height:"100vh"},children:Object(a.jsx)(q.a,{})})}var Me=n(152),Ye=n(620),_e=n(647),Ke=n(621),Qe=n(642),Xe=n(643),et=Object(d.a)((function(e){return Object(j.a)({root:{padding:"2px 4px",display:"flex",alignItems:"center",width:"100%",marginTop:40},input:{marginLeft:e.spacing(1),flex:1},iconButton:{padding:10},divider:{height:28,margin:4},formControl:{margin:e.spacing(1),minWidth:120},hint:{marginLeft:10}})})),tt=[{value:"Prefix",tooltip:"A sentence with a given prefix or suffix"},{value:"Suffix"},{value:"Phone",tooltip:"A phone number with a give area code and/or consecutive number pattern"},{value:"Email",tooltip:"An email address on a domain that is only partially provided"},{value:"Three Words",tooltip:"Text that contains at least three of the following case sensitive words. [meeting, minutes, agenda, action, attendees, apologies]"},{value:"3 Letter word",tooltip:"A 3+ letter word thats all caps"}];function nt(e){var t=e.setDisplayNotes,n=e.notes,r=et(),c=i.a.useState("Pefix"),s=Object(o.a)(c,2),l=s[0],u=s[1],d=i.a.useState(!1),j=Object(o.a)(d,2),h=j[0],b=j[1],f=i.a.useState(0),O=Object(o.a)(f,2),p=O[0],x=O[1],v=i.a.useState(!1),w=Object(o.a)(v,2),y=w[0],S=w[1],N=i.a.useState(""),C=Object(o.a)(N,2),k=C[0],E=C[1],I=/([A-Z]{3,})/g,T=/(meeting|minutes|agenda|action|attendees|apologies)/g,R=/[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@([a-zA-Z0-9]){3,}(\.)*([a-zA-Z0-9])*/g,D=/^([0-9]){2}$/g,L=/^([0-9]){1,}$/g,P=function(e,t,n){var a=[];return e.forEach((function(e){if("Prefix"===t)e.content.split(".").forEach((function(t){n&&t.startsWith(n)&&a.push(e)}));else if("Suffix"===t)e.content.split(".").forEach((function(t){n&&t.endsWith(n)&&a.push(e)}));else if("Phone"===t){var r=!1;k.split(" ").forEach((function(t){1!==z(t,D)||r?1===z(t,L)&&(a.includes(e)||a.push(e)):(r=!0,a.push(e))}))}else"Email"===t?z(e.content,R)>0&&a.push(e):"Three Words"===t?z(e.content,T)>=3&&a.push(e):"3 Letter word"===t&&z(e.content,I)>0&&a.push(e)})),a},z=function(e,t){return((e||"").match(t)||[]).length},A=function(e){var a=[];n.forEach((function(t){e.forEach((function(e){e.id===t.id&&a.push(t)}))})),t(a)},U=function(e){return oe.EditorState.createWithContent(Object(oe.convertFromRaw)(JSON.parse(atob(e)))).getCurrentContent().getPlainText("")};return Object(a.jsxs)(a.Fragment,{children:[Object(a.jsxs)(Me.a,{component:"form",className:r.root,children:[Object(a.jsx)(re.a,{className:r.input,placeholder:"Search",inputProps:{"aria-label":"search"},id:"search-field",onChange:function(e){E(e.target.value)},disabled:h}),y&&Object(a.jsx)(g.a,{className:r.iconButton,type:"reset","aria-label":"reset",onClick:function(e){return function(e){e.preventDefault(),x(0),S(!1),t(n)}(e)},children:Object(a.jsx)(Qe.a,{})}),!y&&Object(a.jsx)(g.a,{type:"submit",className:r.iconButton,"aria-label":"search",onClick:function(e){return function(e){e.preventDefault();var t=[];n.forEach((function(e){t.push({id:e.id,content:U(e.content)})}));var a=P(t,l,k);x(a.length),A(a),S(!0)}(e)},children:Object(a.jsx)(Xe.a,{})}),Object(a.jsx)(pe.a,{className:r.divider,orientation:"vertical"}),Object(a.jsx)(Ye.a,{className:r.formControl,children:Object(a.jsx)(_e.a,{value:l,onChange:function(e){u(e.target.value),"Prefix"===e.target.value||"Suffix"===e.target.value||"Phone"===e.target.value?b(!1):b(!0)},inputProps:{"aria-label":"search type"},children:tt.map((function(e){return Object(a.jsx)(m.a,{value:e.value,children:e.value},e.value)}))})})]}),Object(a.jsx)(Ke.a,{className:r.hint,children:y&&"".concat(p," results match your search.")})]})}var at=Object(d.a)((function(e){return Object(j.a)({root:{flexGrow:1,marginLeft:40},title:{margin:e.spacing(4,0,2)},heading:{fontSize:e.typography.pxToRem(15),flexBasis:"33.33%",flexShrink:0},secondaryHeading:{fontSize:e.typography.pxToRem(15),color:e.palette.text.secondary},item:{width:"100%"}})}));function rt(){var e=C().user,t=Object(u.g)(),n=at(),c=i.a.useState([]),s=Object(o.a)(c,2),l=s[0],d=s[1],j=i.a.useState([]),h=Object(o.a)(j,2),b=h[0],O=h[1],p=i.a.useState(!1),m=Object(o.a)(p,2),x=m[0],v=m[1],w=i.a.useState(new Y(0,"","","",e.id,[0],[0])),y=Object(o.a)(w,2),S=y[0],N=y[1],k=i.a.useState("change-me-to-refresh"),E=Object(o.a)(k,2),I=E[0],T=E[1];return Object(r.useEffect)((function(){O(l)}),[l]),e.loggedIn?Object(a.jsxs)("div",{className:n.root,children:[Object(a.jsxs)(J.a,{container:!0,spacing:2,direction:"row",justify:"center",alignItems:"center",children:[Object(a.jsx)(J.a,{item:!0,xs:9,md:9,children:Object(a.jsxs)(J.a,{container:!0,spacing:2,direction:"row",justify:"center",alignItems:"center",children:[Object(a.jsx)(J.a,{item:!0,xs:3,children:Object(a.jsxs)(f.a,{variant:"h6",className:n.title,children:["Your Notes",Object(a.jsx)(g.a,{onClick:function(){return N(new Y(0,"","","",e.id,[],[])),void v(!0)},children:Object(a.jsx)(_.a,{})})]})}),Object(a.jsx)(J.a,{item:!0,xs:6,children:Object(a.jsx)(nt,{notes:l,setDisplayNotes:O})}),Object(a.jsx)(J.a,{item:!0,xs:3})]})}),Object(a.jsx)(J.a,{item:!0,xs:9,md:9,children:Object(a.jsx)($e,{shouldRefresh:I,setShouldRefresh:T,notes:b,setNotes:d,setActiveNote:N,setEditorOpen:v})})]}),Object(a.jsx)(he,{open:x,setOpen:v,note:S,setShouldRefresh:T})]}):(t.push("/login"),Object(a.jsx)("div",{}))}var it=Object(d.a)((function(e){return{paper:{marginTop:e.spacing(8),display:"flex",flexDirection:"column",alignItems:"center"},avatar:{margin:e.spacing(1),backgroundColor:e.palette.secondary.main},form:{width:"100%",marginTop:e.spacing(1)},submit:{margin:e.spacing(3,0,2)}}}));function ct(){var e=it(),t=i.a.useState(""),n=Object(o.a)(t,2),r=n[0],c=n[1],s=i.a.useState(""),d=Object(o.a)(s,2),j=d[0],h=d[1],b=i.a.useState(""),g=Object(o.a)(b,2),O=g[0],p=g[1],m=i.a.useState(!1),v=Object(o.a)(m,2),y=v[0],N=v[1],k=i.a.useState(!1),E=Object(o.a)(k,2),I=E[0],L=E[1],A=i.a.useState(!1),U=Object(o.a)(A,2),B=U[0],W=U[1],V=i.a.useState(!1),F=Object(o.a)(V,2),Z=F[0],$=F[1],M=i.a.useState(!1),Y=Object(o.a)(M,2),_=Y[0],K=Y[1],Q=C(),X=Q.user,ee=Q.setUser,te=Object(u.g)(),ne=/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;X.id>0&&te.push("/");return Object(a.jsxs)(z.a,{component:"main",maxWidth:"xs",children:[Object(a.jsx)(l.a,{}),Object(a.jsxs)("div",{className:e.paper,children:[Object(a.jsx)(w.a,{className:e.avatar,children:Object(a.jsx)(P.a,{})}),Object(a.jsx)(f.a,{component:"h1",variant:"h5",children:"Register"}),Object(a.jsxs)("form",{className:e.form,noValidate:!0,children:[Object(a.jsx)(T.a,{variant:"outlined",margin:"normal",required:!0,fullWidth:!0,id:"username",label:"Username",name:"username",error:y,onChange:function(e){if(/\s/g.test(e.target.value)||""===e.target.value)N(!0);else{var t=e.target.value.toLowerCase();H.getByName(t).then((function(e){""===e.data.data.name?(c(t),N(!1)):N(!0)}))}},disabled:_,autoFocus:!0}),Object(a.jsx)(T.a,{variant:"outlined",margin:"normal",required:!0,fullWidth:!0,id:"email",label:"Email",name:"email",error:I,onChange:function(e){if(ne.test(e.target.value)){var t=e.target.value.toLowerCase();H.getByEmail(t).then((function(e){""===e.data.data.email?(p(t),L(!1)):L(!0)}))}else L(!0)},disabled:_,autoFocus:!0}),Object(a.jsx)(T.a,{variant:"outlined",margin:"normal",required:!0,fullWidth:!0,name:"password",label:"Password",type:"password",id:"password",error:B,onChange:function(e){""===e.target.value?(W(!0),console.log("invalid password")):(W(!1),h(e.target.value))},disabled:_}),Object(a.jsx)(T.a,{variant:"outlined",margin:"normal",required:!0,fullWidth:!0,name:"repeat-password",label:"Repeat Password",type:"password",id:"repeat-password",error:Z,onChange:function(e){""===e.target.value||(e.target.value,0)?$(!0):W(!1)},disabled:_}),Object(a.jsx)(x.a,{type:"submit",fullWidth:!0,variant:"contained",color:"primary",disabled:y||B||_,className:e.submit,onClick:function(e){K(!0),e.preventDefault(),X.name=r,X.password=j,X.email=O,X.id=0,H.create(X).then((function(e){console.log(e),console.log(X),H.login(r,j).then((function(e){console.log("login response"),console.log(e);var t=e.data.data.data,n=t.token;console.log(n),atob(t.token)===r+j&&(console.log("Time to login"),ee(new S(t.name,t.password,!0,t.token,t.email,t.gender,t.userId,-1)),console.log(X),console.log("Time to push to home"),te.push("/"))}))}))},children:_?Object(a.jsx)(q.a,{size:25}):"Register"})]}),Object(a.jsxs)(J.a,{container:!0,children:[Object(a.jsx)(J.a,{item:!0,xs:!0}),Object(a.jsx)(J.a,{item:!0,children:Object(a.jsx)(R.b,{to:"/login",children:"Already have an account? Login"})})]})]}),Object(a.jsx)(D.a,{mt:8,children:Object(a.jsx)(G,{})})]})}var st=n(293),ot=n.n(st),lt=n(294),ut=n.n(lt),dt=Object(d.a)((function(e){return Object(j.a)({root:{flexGrow:1,marginLeft:40},title:{margin:e.spacing(4,0,2)},heading:{fontSize:e.typography.pxToRem(15),flexBasis:"33.33%",flexShrink:0},secondaryHeading:{fontSize:e.typography.pxToRem(15),color:e.palette.text.secondary},item:{width:"100%"}})}));function jt(e,t,n){var r=n.expanded,c=n.handleExpand,s=n.classes,o=n.deleteUser;return t.map((function(t){return i.a.cloneElement(e,{key:t.id},Object(a.jsxs)(ot.a,{className:s.item,expanded:r==="panel".concat(t.id),onChange:c("panel".concat(t.id)),children:[Object(a.jsxs)(ut.a,{expandIcon:Object(a.jsx)(we.a,{}),"aria-controls":"panel".concat(t.id,"bh-content"),id:"panel".concat(t.id,"bh-header"),children:[Object(a.jsx)(f.a,{className:s.heading,children:t.name}),Object(a.jsx)(f.a,{className:s.secondaryHeading,children:t.email})]}),Object(a.jsxs)(ge.a,{children:[Object(a.jsx)(Oe.a,{title:"Delete User",children:Object(a.jsx)(g.a,{onClick:function(){return o(t.id)},size:"small",children:Object(a.jsx)(ye.a,{})})}),Object(a.jsx)(Oe.a,{title:"Edit User",children:Object(a.jsx)(g.a,{size:"small",children:Object(a.jsx)(Se.a,{})})})]}),Object(a.jsx)(pe.a,{}),Object(a.jsx)(me.a,{children:Object(a.jsx)(f.a,{className:s.heading,children:t.gender})})]}))}))}function ht(){var e=dt(),t=i.a.useState(!1),n=Object(o.a)(t,2),c=n[0],s=n[1],l=i.a.useState([]),d=Object(o.a)(l,2),j=d[0],h=d[1],b=i.a.useState(!1),g=Object(o.a)(b,1)[0],O=C().user,p=i.a.useState("change-me-to-refresh"),m=Object(o.a)(p,2),x=m[0],v=m[1],w=i.a.useState(!1),y=Object(o.a)(w,2),N=y[0],k=y[1];return Object(r.useEffect)((function(){H.getData().then((function(e){var t=[];return void 0===e.data||null===e.data?[]:(e.data.forEach((function(e){t.push(new S(e.name,"",!1,"",e.email,e.gender,e.userId,-1))})),t)})).then((function(e){h(e),s(!0)}))}),[x]),O.loggedIn?c?Object(a.jsx)("div",{className:e.root,children:Object(a.jsxs)(J.a,{container:!0,spacing:2,direction:"row",justify:"center",alignItems:"center",children:[Object(a.jsxs)(J.a,{item:!0,xs:12,md:6,children:[Object(a.jsx)(f.a,{variant:"h6",className:e.title,children:"List of Users"}),Object(a.jsx)(xe.a,{dense:g,children:jt(Object(a.jsx)(ve.a,{}),j,{expanded:N,handleExpand:function(e){return function(t,n){k(!!n&&e)}},classes:e,deleteUser:function(e){console.log("Click"),H.delete(e).then((function(e){console.log(e),v(String(Date.now()))}))}})})]}),Object(a.jsx)(J.a,{item:!0,xs:12,md:12})]})}):Object(a.jsx)("div",{children:"Loading..."}):Object(a.jsx)(u.a,{to:"/login"})}var bt=n(275),ft=n(646),gt=n(295),Ot=new(function(){function e(){Object(y.a)(this,e)}return Object(W.a)(e,[{key:"get",value:function(){return F.get("/")}}]),e}()),pt=Object(d.a)((function(e){return Object(j.a)({root:{width:"100%","& > * + *":{marginTop:e.spacing(2)}}})}));function mt(){var e=pt(),t=i.a.useState(!1),n=Object(o.a)(t,2),c=n[0],s=n[1];return Object(r.useEffect)((function(){Object(gt.setInterval)((function(){Ot.get().then((function(){c&&s(!1)})).catch((function(){c||s(!0)}))}),5e3)})),Object(a.jsx)("div",{className:e.root,children:Object(a.jsx)(bt.a,{in:c,children:Object(a.jsx)(ft.a,{severity:"error",children:"Connection to backend lost, attempting to reconnect."})})})}var xt=function(){var e=i.a.useState(new S("","",!1,"","","",-1,-1)),t=Object(o.a)(e,2),n=t[0],r=t[1];return console.log("App Level User:"),console.log(n),Object(a.jsx)(N.Provider,{value:{user:n,setUser:r},children:Object(a.jsxs)(i.a.Fragment,{children:[Object(a.jsx)(l.a,{}),Object(a.jsx)(I,{}),Object(a.jsx)(mt,{}),Object(a.jsxs)(u.d,{children:[Object(a.jsx)(u.b,{exact:!0,path:["/","/home"],component:rt}),Object(a.jsx)(u.b,{exact:!0,path:"/login",component:M}),Object(a.jsx)(u.b,{exact:!0,path:"/register",component:ct}),Object(a.jsx)(u.b,{exact:!0,path:"/users",component:ht})]})]})})},vt=function(e){e&&e instanceof Function&&n.e(3).then(n.bind(null,652)).then((function(t){var n=t.getCLS,a=t.getFID,r=t.getFCP,i=t.getLCP,c=t.getTTFB;n(e),a(e),r(e),i(e),c(e)}))};s.a.render(Object(a.jsx)(i.a.StrictMode,{children:Object(a.jsx)(R.a,{basename:"/web/",children:Object(a.jsx)(xt,{})})}),document.getElementById("root")),vt()}},[[582,1,2]]]);
//# sourceMappingURL=main.ee61e389.chunk.js.map