/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import 'react-vis/dist/style.css';
import { Sunburst } from 'react-vis';
import { LabelSeries } from 'react-vis';
import { EXTENDED_DISCRETE_COLOR_RANGE } from 'react-vis/dist/theme';

const LABEL_STYLE = {
  fontSize: '16px',
  fontWeight: 'bold',
  textAnchor: 'middle'
};

/**
 * Recursively work backwards from highlighted node to find path of value nodes
 * @param {Object} node - the current node being considered
 * @returns {Array} an array of strings describing the key route to the current node
 */
function getKeyPath(node) {
  if (!node.parent) {
    return ['RECOMMENDATION'];
  }

  return [(node.data && node.data.name) || node.name].concat(
    getKeyPath(node.parent)
  );
}

/**
 * Recursively modify data depending on whether or not each cell has been selected by the hover/highlight
 * @param {Object} data - the current node being considered
 * @param {Object|Boolean} keyPath - a map of keys that are in the highlight path
 * if this is false then all nodes are marked as selected
 * @returns {Object} Updated tree structure
 */
function updateData(data, keyPath, depth) {
  if (data.children) {
    data.children.map(child => updateData(child, keyPath, depth + 1));
  }
  // add a fill to all the uncolored cells
  if (!data.hex) {
    data.style = {
      fill: EXTENDED_DISCRETE_COLOR_RANGE[depth]
    };
  }
  data.style = {
    ...data.style,
    fillOpacity: keyPath && !keyPath[data.name] ? 0.2 : 1
  };

  return data;
}

class RecommendationSunBurst extends Component {
  constructor(props) {
    super(props);
    this.state = {
      data: { name: '', color: 1, children: [] },
      pathValue: false,
      finalValue: 'RECOMMENDATIONS',
      clicked: false
    };

    this.processChartData = this.processChartData.bind(this);
  }

  processChartData(data) {
    let updatedChartData = {
      name: 'outcomes',
      hex: '#004775',
      children: []
    };
    let PASchildren = {
      name: 'PAS',
      hex: '#04c775',
      children: []
    };
    let PKSchildren = {
      name: 'PKS',
      hex: '#00a775',
      children: []
    };
    let OtherChildren = {
      name: 'OTHER',
      hex: '#007775',
      children: []
    };
    updatedChartData.children.push(PASchildren);
    updatedChartData.children.push(PKSchildren);
    updatedChartData.children.push(OtherChildren);

    //Map the outcomes
    let outcomes = new Map();
    data.forEach(row => {
      let appItem = { name: row.a, hex: '#456c80', value: row.x, children: [] };
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

    let updatedData = updateData(updatedChartData, false);
    this.setState({
      data: updatedData,
      decoratedData: updatedData
    });
  }

  componentDidMount() {
    if (this.props.data) {
      this.processChartData(this.props.data);
    }
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    //console.log("ReactVisChart - Will Receive Props!");
    //console.log("New Props => \n%s\n\nOld Props =>\n%s", JSON.stringify(newProps), JSON.stringify(this.props));
    if (newProps.data) {
      this.processChartData(newProps.data);
    }
  }

  render() {
    const decoratedData = this.state.decoratedData;

    const showChartElements =
      this.state.data && this.state.data.children.length > 0;
    const { clicked, data, finalValue, pathValue } = this.state;
    return (
      <div>
        <div>
          <div className="ui-g-12 ui-g-nopad show-labels">
            <span>&nbsp;</span>
          </div>
          <div className="ui-g-12">
            {showChartElements && (
              <Sunburst
                animation
                className="sunburst"
                hideRootNode
                onValueMouseOver={node => {
                  if (clicked) {
                    return;
                  }
                  const path = getKeyPath(node).reverse();
                  const depth = path.length - 1;
                  const pathAsMap = path.reduce((res, row) => {
                    res[row] = true;
                    return res;
                  }, {});
                  this.setState({
                    finalValue: path[depth],
                    pathValue: path.join(' > '),
                    data: updateData(decoratedData, pathAsMap, depth)
                  });
                }}
                onValueMouseOut={() =>
                  clicked
                    ? () => {}
                    : this.setState({
                        pathValue: false,
                        finalValue: false,
                        data: updateData(decoratedData, false, 0)
                      })
                }
                onValueClick={() => this.setState({ clicked: !clicked })}
                style={{
                  stroke: '#ddd',
                  strokeOpacity: 0.3,
                  strokeWidth: '0.5'
                }}
                //colorType="literal"
                getSize={d => d.value}
                getColor={d => d.hex}
                data={data}
                height={500}
                width={850}
              >
                {finalValue && (
                  <LabelSeries
                    data={[
                      { x: 0, y: 0, label: finalValue, style: LABEL_STYLE }
                    ]}
                  />
                )}
              </Sunburst>
            )}
            <div className="ui-g-12 ui-g-nopad" style={{ textAlign: 'center' }}>
              {pathValue}
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default RecommendationSunBurst;
