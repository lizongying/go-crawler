import{p as a,q as c,t as d,N as l,O as p,Q as _,x as u,S as k,U as h,c as o,y as n,X as f}from"./index-1ddded62.js";import{_ as g}from"./0b6d9653.js";import{_ as v}from"./8025f07a.js";import"./92337dc1.js";const w={key:0,style:{"font-weight":"bold"}},B={key:1},I=n("a",null,"Run",-1),C=n("a",null,"Delete",-1),N={class:"ant-dropdown-link"},O={__name:"SchedulesView",emits:["router—change"],setup(S){const m=[{title:"Spider",dataIndex:"spider",sorter:(t,e)=>t.spider>e.spider,width:200},{title:"Schedule",dataIndex:"schedule",sorter:(t,e)=>t.schedule>e.schedule},{title:"Action",dataIndex:"action",width:200,fixed:"right"}],y=[{id:"1",spider:"John Brown",schedule:"every day"},{id:"2",spider:"Jim Green",schedule:"every day"},{id:"3",spider:"Joe Black",schedule:"every day"}];return(t,e)=>{const i=v,x=g;return a(),c(x,{columns:m,"data-source":y,scroll:{x:"100%"}},{headerCell:d(({column:s})=>[["spider","schedule"].includes(s.dataIndex)?(a(),l("span",w,p(s.title),1)):_("",!0)]),bodyCell:d(({column:s,record:r})=>[s.dataIndex==="spider"?(a(),c(u(k),{key:0,to:"/spiders?name="+r.spider,onClick:e[0]||(e[0]=V=>t.$emit("router—change","3"))},{default:d(()=>[h(p(r.spider),1)]),_:2},1032,["to"])):s.dataIndex==="action"?(a(),l("span",B,[I,o(i,{type:"vertical"}),C,o(i,{type:"vertical"}),n("a",N,[h(" More "),o(u(f))])])):_("",!0)]),_:1})}}};export{O as default};
