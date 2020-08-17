/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import KeyValuePairComponent from '../util/KeyValuePairComponent';

export default class RunDetails extends Component {
  render() {
    return (
      <div className="ui-g ui-g-nopad card run-summary">
        <div className="ui-g-12 ui-g-nopad">
          <KeyValuePairComponent
            field="User"
            value={this.props.selectedRun.User}
          />
          <KeyValuePairComponent
            field="Request Date"
            value={this.props.selectedRun.requestDate}
          />
          <KeyValuePairComponent
            field="Target"
            value={this.props.selectedRun.Target}
          />
          <KeyValuePairComponent
            field="Runtime"
            value={this.props.selectedRun.Runtime}
          />
        </div>
      </div>
    );
  }
}
