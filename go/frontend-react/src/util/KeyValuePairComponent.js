/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';

class KeyValuePairComponent extends Component {
  render() {
    const { field, value, height } = this.props;
    return (
      <div className="ui-g-12 ui-g-nopad" style={{ heigth: height }}>
        <div className="ui-g-2 ui-g-nopad kv-field">{field}:</div>
        <div className="ui-g-10 ui-g-nopad">
          <p className="kv-value">{value}</p>
        </div>
      </div>
    );
  }
}

export default KeyValuePairComponent;
