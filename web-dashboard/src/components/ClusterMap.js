import React, { useEffect, useState } from 'react';
import * as d3 from 'd3';
import { fetchClusterData } from '../services/api';

export default function ClusterMap() {
    const [nodes, setNodes] = useState([]);

    useEffect(() => {
        async function loadData() {
            const data = await fetchClusterData();
            setNodes(data.nodes);
            renderGraph(data.nodes);
        }
        loadData();
    }, []);

    function renderGraph(nodes) {
        const svg = d3.select('#cluster-svg')
            .attr('width', 1000)
            .attr('height', 600);

        svg.selectAll("*").remove();

        const simulation = d3.forceSimulation(nodes)
            .force("charge", d3.forceManyBody().strength(-50))
            .force("center", d3.forceCenter(500, 300))
            .force("link", d3.forceLink().id(d => d.id).distance(50));

        const nodeElements = svg.selectAll('circle')
            .data(nodes)
            .enter().append('circle')
            .attr('r', 10)
            .attr('fill', 'steelblue');

        simulation.nodes(nodes).on('tick', () => {
            nodeElements
                .attr('cx', d => d.x)
                .attr('cy', d => d.y);
        });
    }

    return <svg id="cluster-svg"></svg>;
}
