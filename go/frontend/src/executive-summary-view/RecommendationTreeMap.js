/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import 'react-vis/dist/style.css';
import { Treemap } from 'react-vis';

const STYLES = {
  SVG: {
    stroke: '#ddd',
    strokeWidth: '0.25',
    strokeOpacity: 0.5
  },
  DOM: {
    border: 'thin solid #ddd'
  }
};

class RecommendationTreeMap extends Component {
  constructor(props) {
    super(props);
    this.state = {
      data: { name: '', hex: 1, children: [] },
      lastDrawLocation: null,
      selectedPointId: null,
      crosshairValues: false,
      displayLabels: true
    };

    this.processChartData = this.processChartData.bind(this);
  }

  processChartData(data) {
    let updatedChartData = {
      name: 'outcomes',
      hex: '#004775',
      children: [],
      style: {
        border: 'thin solid black'
      }
    };
    let PASchildren = {
      name: 'PAS',
      hex: '#04c775',
      children: [],
      style: {
        border: 'thin solid red'
      }
    };
    let PKSchildren = {
      name: 'PKS',
      hex: '#00a775',
      children: [],
      style: {
        border: 'thin solid blue'
      }
    };
    let OtherChildren = {
      name: 'OTHER',
      hex: '#007775',
      children: [],
      style: {
        border: 'thin solid green'
      }
    };
    updatedChartData.children.push(PASchildren);
    updatedChartData.children.push(PKSchildren);
    updatedChartData.children.push(OtherChildren);

    //Map the outcomes
    let outcomes = new Map();
    data.forEach(row => {
      let appItem = {
        name: row.a,
        hex: '#456c80',
        size: row.x,
        rec: row.r,
        children: []
      };
      let outcomeRows = outcomes.get(row.r);
      if (outcomeRows === undefined) {
        outcomes.set(row.r, {
          name: row.r,
          hex: '#00b5a3',
          children: []
        });
        outcomes.get(row.r).children.push(appItem);
      } else {
        outcomeRows.children.push(appItem);
      }
    });

    for (let oc of outcomes.entries()) {
      console.log(
        'binning outcome [%s] with [%d] entries',
        oc[0],
        oc[1].children.length
      );
      if (oc[0].toLowerCase().includes('pas')) {
        PASchildren.children.push(oc[1]);
      } else if (oc[0].toLowerCase().includes('pks')) {
        PKSchildren.children.push(oc[1]);
      } else {
        OtherChildren.children.push(oc[1]);
      }
    }

    this.setState({
      data: updatedChartData
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
    const showChartElements =
      this.state.data && this.state.data.children.length > 0;
    const { hoveredNode } = this.state;

    return (
      <div>
        <div>
          <div className="ui-g-12 ui-g-nopad show-labels">
            <span>&nbsp;</span>
          </div>
          <div className="ui-g-12">
            {showChartElements && (
              <Treemap
                {...{
                  className: 'nested-tree-example',
                  colorType: 'literal',
                  data: this.state.data,
                  getColor: d => d.hex,
                  margin: { left: 40, right: 40, top: 10, bottom: 40 },
                  height: this.props.height ? this.props.height : 500,
                  width: this.props.width ? this.props.width : 1000,
                  animation: {
                    damping: 9,
                    stiffness: 300
                  },
                  mode: 'circlePack',
                  style: STYLES['SVG'],
                  onLeafMouseOver: x => {
                    if (x.data !== undefined) {
                      if (x.data.rec !== undefined) {
                        this.setState({
                          hoveredNode: {
                            Name: x.data.name,
                            Rec: x.data.rec,
                            container: false
                          }
                        });
                      } else {
                        this.setState({
                          hoveredNode: {
                            Name: x.data.children.length,
                            Rec: x.data.name,
                            container: true
                          }
                        });
                      }
                    }
                  },
                  onLeafMouseOut: () => this.setState({ hoveredNode: false })
                  // hideRootNode: true
                }}
              />
            )}
          </div>
          <div className="ui-g-12 ui-g-nopad" style={{ textAlign: 'center' }}>
            {hoveredNode && (
              <div className="tree-tip">
                <span className="row">
                  {hoveredNode.container && (
                    <span className="lbl"># Apps:</span>
                  )}
                  {!hoveredNode.container && <span className="lbl">App:</span>}
                  <span className="val">{hoveredNode.Name}</span>
                  <span className="lbl">&nbsp;&nbsp;Recommendation:</span>
                  <span className="val">{hoveredNode.Rec}</span>
                </span>
              </div>
            )}
          </div>
        </div>
      </div>
    );
  }
}

export default RecommendationTreeMap;
