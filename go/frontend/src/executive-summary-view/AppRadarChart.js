/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import 'react-vis/dist/style.css';
import { RadarChart, Hint } from 'react-vis';
import { format } from 'd3-format';

const scoreFormat = format('.2r');

const tipStyle = {
  display: 'flex',
  color: '#fff',
  background: '#00b5a3',
  alignItems: 'center',
  padding: '5px'
};

class AppRadarChart extends Component {
  constructor(props) {
    super(props);
    this.state = {
      data: [],
      hoveredCell: false
    };

    this.processChartData = this.processChartData.bind(this);
  }

  processChartData(data) {
    let maxSloc = 0;
    let maxFiles = 0;
    let maxFindings = 0;
    let maxRaw = 0;

    data.map(app => {
      maxSloc = app.s > maxSloc ? app.s : maxSloc;
      maxFiles = app.fc > maxFiles ? app.fc : maxFiles;
      maxFindings = app.f > maxFindings ? app.f : maxFindings;
      maxRaw = app.z > maxRaw ? app.z : maxRaw;
    });

    this.setState({
      data: data,
      maxSloc: maxSloc,
      maxFiles: maxFiles,
      maxFindings: maxFindings,
      maxRaw: maxRaw
    });
  }

  componentDidMount() {
    if (this.props.data) {
      this.processChartData(this.props.data);
    }
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    if (newProps.data) {
      this.processChartData(newProps.data);
    }
  }

  render() {
    const showChartElements = this.state.data && this.state.data.length > 0;
    const {
      data,
      maxSloc,
      maxFiles,
      maxFindings,
      maxRaw,
      hoveredCell
    } = this.state;
    return (
      <div>
        <div>
          <div className="ui-g-12 ui-g-nopad show-labels">
            <span>&nbsp;</span>
          </div>
          <div className="ui-g-12">
            {showChartElements && (
              <RadarChart
                animation
                data={data}
                startingAngle={0}
                renderAxesOverPolygons
                domains={[
                  {
                    name: 'Score',
                    domain: [0, 10],
                    tickFormat: t => scoreFormat(t),
                    getValue: d => d.x
                  },
                  {
                    name: 'BV',
                    domain: [0, 10],
                    tickFormat: t => scoreFormat(t),
                    getValue: d => d.y
                  },
                  { name: 'SLOC', domain: [1, maxSloc], getValue: d => d.s },
                  { name: 'Files', domain: [1, maxFiles], getValue: d => d.fc },
                  { name: 'Raw', domain: [0, maxRaw], getValue: d => d.z },
                  {
                    name: 'Findings',
                    domain: [0, maxFindings],
                    getValue: d => d.f
                  }
                ]}
                height={this.props.height ? this.props.height : 550}
                width={this.props.width ? this.props.width : 600}
                style={{
                  polygons: {
                    fillOpacity: 0,
                    strokeWidth: 1
                  },
                  axes: {
                    text: {
                      opacity: 1
                    }
                  },
                  labels: {
                    textAnchor: 'middle',
                    fontSize: '16px',
                    fontWeight: 'normal',
                    stroke: '#009fdf'
                  }
                }}
                margin={{
                  left: 120,
                  top: 50,
                  bottom: 50,
                  right: 80
                }}
                onSeriesMouseOver={data => {
                  this.setState({ hoveredCell: data.event[0] });
                }}
                onSeriesMouseOut={() => this.setState({ hoveredCell: false })}
              >
                {hoveredCell && (
                  <Hint value={{ x: 0, y: 0 }}>
                    <div style={tipStyle}>{hoveredCell.name}</div>
                  </Hint>
                )}
              </RadarChart>
            )}
          </div>
        </div>
      </div>
    );
  }
}

export default AppRadarChart;
