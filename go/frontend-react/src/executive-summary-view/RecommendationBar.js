/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import 'react-vis/dist/style.css';
import {
  VerticalBarSeries,
  HorizontalGridLines,
  VerticalGridLines,
  XAxis,
  YAxis,
  FlexibleWidthXYPlot
} from 'react-vis';

class RecommendationBar extends Component {
  constructor(props) {
    super(props);
    this.state = {
      data: []
    };

    this.processChartData = this.processChartData.bind(this);
  }

  processChartData(data) {
    let updatedChartData = [];
    //Map the outcomes
    let outcomes = new Map();
    data.forEach(row => {
      let cnt = outcomes.get(row.r);
      if (cnt === undefined) {
        outcomes.set(row.r, { cnt: 1 });
      } else {
        cnt.cnt++;
      }
    });

    for (let oc of outcomes.entries()) {
      console.log('Setting recommendation [%s] to cnt[%d]', oc[0], oc[1].cnt);
      updatedChartData.push({ x: oc[0], y: oc[1].cnt });
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
    const showChartElements = this.state.data && this.state.data.length > 0;
    const data = this.state.data;
    return (
      <div>
        <div>
          <div className="ui-g-12 ui-g-nopad show-labels">
            <span>&nbsp;</span>
          </div>
          <div className="ui-g-12">
            {showChartElements && (
              <FlexibleWidthXYPlot
                margin={{ left: 40, right: 40, top: 5, bottom: 40 }}
                height={this.props.height ? this.props.height : 550}
                xType="ordinal"
                xDistance={100}
              >
                <VerticalGridLines />
                <HorizontalGridLines />
                <XAxis />
                <YAxis />
                <VerticalBarSeries
                  className="vertical-bar"
                  color="#00b5a3"
                  data={data}
                />
              </FlexibleWidthXYPlot>
            )}
          </div>
        </div>
      </div>
    );
  }
}

export default RecommendationBar;
