import React, { useEffect } from 'react';
import * as d3 from 'd3';

export default function TopologyEditor({nodes, onTopologyChange}){
    useEffect(()=>{
        const svg = d3.select('#topo-editor-svg').attr('width',1000).attr('height',600);
        svg.selectAll("*").remove();

        const drag = d3.drag().on("drag",(event,d)=>{
            d.x = event.x;
            d.y = event.y;
            d3.select(event.sourceEvent.target).attr("cx",d.x).attr("cy",d.y);
            onTopologyChange(nodes);
        });

        svg.selectAll('circle').data(nodes).enter().append('circle')
            .attr('r',12).attr('fill','orange')
            .attr('cx',d=>d.x).attr('cy',d=>d.y)
            .call(drag);
    },[nodes]);

    return <svg id="topo-editor-svg"></svg>;
}
