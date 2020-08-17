/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from "react";
import { Card } from "primereact/card";
import { ToggleButton } from "primereact/togglebutton";

class ScoreCard extends Component {
  constructor(props) {
    super(props);
    this.tagsHeader = this.tagsHeader.bind(this);
    this.binsHeader = this.binsHeader.bind(this);
    this.tagChange = this.tagChange.bind(this);
    this.binChange = this.binChange.bind(this);
  }

  binsHeader(scorecard) {
    if (scorecard && typeof scorecard.bins != "undefined") {
      return (
        <div className="ui-g-12 ui-g-nopad">
          {scorecard.bins.map(bin => {
            return (
              <div key={bin.name} className="scorecard-btn">
                <ToggleButton
                  className="scorecard-bin"
                  onLabel={bin.name}
                  offLabel={bin.name}
                  checked={bin.selected}
                  onChange={() => {
                    this.binChange(bin);
                  }}
                />
              </div>
            );
          })}
        </div>
      );
    } else {
      return <div className="ui-g-12">&nbsp;</div>;
    }
  }

  tagsHeader(scorecard) {
    if (scorecard && typeof scorecard.tags != "undefined") {
      return (
        <div className="ui-g-12 ui-g-nopad">
          {scorecard.tags.map(tag => {
            return (
              <div key={tag.name} className="scorecard-btn">
                <ToggleButton
                  className="scorecard-tag"
                  onLabel={tag.name}
                  offLabel={tag.name}
                  checked={tag.selected}
                  onChange={() => {
                    this.tagChange(tag);
                  }}
                />
              </div>
            );
          })}
        </div>
      );
    } else {
      return <div className="ui-g-12">&nbsp;</div>;
    }
  }

  binChange(bin) {
    console.log("Bin Click event! bin => %s", JSON.stringify(bin));
    this.props.binSetter(bin);
  }

  tagChange(tag) {
    console.log("Tag Click event! tag => %s", JSON.stringify(tag));
    this.props.tagSetter(tag);
  }

  // componentDidMount() {
  //     console.log('ScoreCard - didMount [%s]',JSON.stringify(this.props.applicationScoreCard));
  //     this.setState({scorecard:this.props.applicationScoreCard});
  // }
  //
  // componentWillReceiveProps(newProps) {
  //     console.log('ScoreCard - willReceiveProps')
  //     if (newProps.applicationScoreCard.application !== this.props.applicationScoreCard.application) {
  //         console.log('ScoreCard App Changed!.... existing [%s] new [%s]',JSON.stringify(this.props.applicationScoreCard),JSON.stringify(newProps.applicationScoreCard));
  //         this.setState({scorecard:newProps.applicationScoreCard});
  //     }
  // }

  render() {
    let applicationScoreCard = this.props.applicationScoreCard;
    const newTotal =
      parseInt(applicationScoreCard.low, 10) +
      parseInt(applicationScoreCard.medium, 10) +
      parseInt(applicationScoreCard.high, 10);
    // console.log(
    //   'Scorecard => rendering with sc => %s',
    //   JSON.stringify(applicationScoreCard)
    // );
    const binsHeader = this.binsHeader(applicationScoreCard);
    const tagsHeader = this.tagsHeader(applicationScoreCard);

    return (
      <div className="ui-g ui-fluid ui-g-nopad">
        <div className="ui-g-12 ui-g-nopad">{binsHeader}</div>
        <div className="ui-g-12 ui-g-nopad">{tagsHeader}</div>
        {/* <div className="ui-g-12 ui-md-2">
          <Card
            title="Info"
            subTitle={applicationScoreCard.info}
            style={{
              backgroundColor: '#229be6',
              border: '#229be6',
              textAlign: 'center'
            }}
          >
            <i className="fa fa-info-circle" style={{ fontSize: '60px' }} />
          </Card>
        </div> */}

        <div className="ui-g-12 ui-md-2">
          <Card
            title="Low"
            subTitle={applicationScoreCard.low}
            style={{
              backgroundColor: "#1bb3a8",
              border: "#1bb3a8",
              textAlign: "center"
            }}
          >
            <i
              className="fa fa-thermometer-quarter"
              style={{ fontSize: "60px" }}
            />
          </Card>
        </div>

        <div className="ui-g-12 ui-md-2">
          <Card
            title="Medium"
            subTitle={applicationScoreCard.medium}
            style={{
              backgroundColor: "orange",
              border: "orange",
              textAlign: "center"
            }}
          >
            <i
              className="fa fa-thermometer-half"
              style={{ fontSize: "60px" }}
            />
          </Card>
        </div>

        <div className="ui-g-12 ui-md-2">
          <Card
            title="High"
            subTitle={applicationScoreCard.high}
            style={{
              backgroundColor: "#f2242f",
              border: "#f2242f",
              textAlign: "center"
            }}
          >
            <i
              className="fa fa-thermometer-full"
              style={{ fontSize: "60px" }}
            />
          </Card>
        </div>

        <div className="ui-g-12 ui-md-2">
          <Card
            title="Total"
            subTitle={newTotal}
            //subTitle={applicationScoreCard.total}
            style={{
              backgroundColor: "#607580",
              border: "#607580",
              textAlign: "center"
            }}
          >
            <i className="fa fa-heartbeat" style={{ fontSize: "60px" }} />
          </Card>
        </div>
        <div className="ui-g-12 ui-md-2 tooltip">
          <Card
            title="Score"
            subTitle={applicationScoreCard.app.score.toString()}
            style={{
              backgroundColor: "#138078",
              border: "#607580",
              textAlign: "center"
            }}
          >
            <span className="scoretooltiptext">
              {applicationScoreCard.app.rawScore}
            </span>
            <i className="fa fa-bullseye" style={{ fontSize: "60px" }} />
          </Card>
        </div>
      </div>
    );
  }
}

export default ScoreCard;
