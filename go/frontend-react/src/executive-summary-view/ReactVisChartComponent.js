/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { Checkbox } from 'primereact/checkbox';
import 'react-vis/dist/style.css';
import {
  Hint,
  FlexibleWidthXYPlot,
  Highlight,
  HorizontalGridLines,
  MarkSeries,
  VerticalGridLines,
  XAxis,
  YAxis,
  LabelSeries,
  PolygonSeries
} from 'react-vis';

const markerColor = '#009fdf';
const markerSelected = '#f27062';
const chartBackground = '#00b5a3';

class ReactVisChartComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      data: [],
      lastDrawLocation: null,
      selectedPointId: null,
      crosshairValues: false,
      displayLabels: true
    };

    this.findValue = this.findValue.bind(this);
    this.processChartData = this.processChartData.bind(this);
    this.determineColor = this.determineColor.bind(this);
    this.displayLabelsCBValueChange = this.displayLabelsCBValueChange.bind(
      this
    );
  }

  findValue(item, e) {
    //console.log('finding Value => %s', JSON.stringify(item));
    this.setState({
      selectedPointId: item.id,
      crosshairValues: item
    });
  }

  displayLabelsCBValueChange(e) {
    this.setState({ displayLabels: e.checked });
  }

  processChartData(data) {
    let updatedChartData = [];
    data.forEach(row => {
      updatedChartData.push({
        x: row.x,
        y: row.y,
        size: row.z,
        name: row.a,
        label: row.a,
        r: row.r
      });
    });
    this.setState({
      data: updatedChartData.map((d, id) => ({ ...d, id }))
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

  determineColor(id) {
    return markerColor;
  }

  render() {
    const lastDrawLocation = this.state.lastDrawLocation;
    const selectedPointId = this.state.selectedPointId;
    const crosshairValues = this.state.crosshairValues;
    const displayLabels = this.state.displayLabels;
    const scaleBound = !!this.props.bound;
    const showChartElements = this.state.data && this.state.data.length > 0;

    return (
      <div>
        <div>
          <div className="ui-g-12 ui-g-nopad show-labels">
            <Checkbox
              inputId="cb1"
              onChange={this.displayLabelsCBValueChange}
              checked={displayLabels}
            />
            <label htmlFor="cb1">Show Labels</label>
          </div>
          <div className="ui-g-12">
            <FlexibleWidthXYPlot
              margin={{ left: 40, right: 40, top: 5, bottom: 40 }}
              height={this.props.height ? this.props.height : 550}
              animation={'wobbly'}
              xDomain={
                lastDrawLocation
                  ? [lastDrawLocation.left, lastDrawLocation.right]
                  : scaleBound
                    ? [0, 10]
                    : null
              }
              yDomain={
                lastDrawLocation
                  ? [lastDrawLocation.bottom, lastDrawLocation.top]
                  : scaleBound
                    ? [0, 10]
                    : null
              }
              onValueMouseOut={() => this.setState({ crosshairValues: false })}
              onMouseLeave={() =>
                this.setState({ crosshairValues: false, selectedPointId: -1 })
              }
            >
              <PolygonSeries
                data={[
                  { x: 0, y: 0 },
                  { x: 10, y: 0 },
                  { x: 10, y: 10 },
                  { x: 0, y: 10 }
                ]}
                style={{
                  strokeWidth: 0.5,
                  strokeOpacity: 1,
                  opacity: 0.4,
                  fill: { chartBackground }
                }}
              />

              <VerticalGridLines />
              <HorizontalGridLines />
              <XAxis
                title="Technical Fitness"
                className="axis-title"
                position="start"
              />
              <YAxis
                title="Business Value"
                className="axis-title"
                orientation="right"
              />
              {/*<LabelSeries*/}
              {/*data={labelData}*/}
              {/*labelAnchorX="middle"*/}
              {/*labelAnchorY="middle"*/}
              {/*/>*/}
              {displayLabels && (
                <LabelSeries data={this.state.data} labelAnchorX="middle" />
              )}
              {showChartElements && (
                <MarkSeries
                  className="mark-series-example"
                  colorType="literal"
                  stroke="#138078"
                  strokeWidth={1.5}
                  size={8}
                  data={this.state.data}
                  onNearestXY={(value, e) => {
                    this.findValue(value, e);
                  }}
                  getColor={({ id }) =>
                    selectedPointId === id
                      ? markerSelected
                      : this.determineColor(id)
                  }
                  getOpacity={({ id }) =>
                    selectedPointId === id ? '1.0' : '1'
                  }
                />
              )}
              {showChartElements && (
                <Highlight
                  align={{ vertical: 'top', horizontal: 'left' }}
                  onBrushEnd={area => this.setState({ lastDrawLocation: area })}
                  onDrag={area => {
                    this.setState({
                      lastDrawLocation: {
                        bottom:
                          lastDrawLocation.bottom + (area.top - area.bottom),
                        left: lastDrawLocation.left - (area.right - area.left),
                        right:
                          lastDrawLocation.right - (area.right - area.left),
                        top: lastDrawLocation.top + (area.top - area.bottom)
                      }
                    });
                  }}
                />
              )}
              {crosshairValues &&
                showChartElements && (
                  <Hint value={crosshairValues}>
                    <div
                      align={{ horizonal: 'left', vertical: 'top' }}
                      className="chart-tip"
                    >
                      <span className="heading">{crosshairValues.name}</span>
                      <span className="row">
                        <span className="lbl">Score:&nbsp;</span>
                        <span className="val">{crosshairValues.x}</span>
                      </span>
                      <span className="row">
                        <span className="lbl">BV:&nbsp;</span>
                        <span className="val">{crosshairValues.y}</span>
                      </span>
                      <span className="row">
                        <span className="lbl">Raw:&nbsp;</span>
                        <span className="val">{crosshairValues.size}</span>
                      </span>
                      <span className="row">
                        <span className="lbl">Recommendation:&nbsp;</span>
                        <span className="val">{crosshairValues.r}</span>
                      </span>
                    </div>
                  </Hint>
                )}
            </FlexibleWidthXYPlot>
          </div>
        </div>
      </div>
    );
  }
}

export default ReactVisChartComponent;
