/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import KeyValuePairComponent from './KeyValuePairComponent';
import ListUtils from './ListUtils';

class FindingDetails extends Component {
  render() {
    const { color, finding } = this.props;

    let tags = [];

    if (finding.tags) {
      tags = ListUtils().renderList(finding.tags);
    }

    let recipes = [];

    if (finding.recipes) {
      recipes = ListUtils().renderLinks(finding.recipes);
    }

    return (
      <div>
        <h1 style={{ color: color }}>{finding.id}</h1>
        <hr className="finding-rule" />
        <div className="ui-g ui-g-nopad">
          <div className="ui-g-12">
            <KeyValuePairComponent field="File" value={finding.filename} />
            <KeyValuePairComponent field="FQN" value={finding.fqn} />
            <KeyValuePairComponent field="Line" value={finding.line} />
            <KeyValuePairComponent field="Rule" value={finding.rule} />
            <KeyValuePairComponent field="Pattern" value={finding.pattern} />
            <KeyValuePairComponent field="Value" value={finding.value} />
            <KeyValuePairComponent field="Effort" value={finding.effort} />
            <KeyValuePairComponent field="Category" value={finding.category} />
            <KeyValuePairComponent
              field="Application"
              value={finding.application}
            />
          </div>
          <div className="ui-g-12">
            <h2 className="finding-advice">Notes</h2>
            <hr className="finding-rule" />
            <p>{finding.note}</p>
          </div>
          <div className="ui-g-12">
            <h2 className="finding-advice">Advice</h2>
            <hr className="finding-rule" />
            <p>{finding.advice}</p>
          </div>
          <div className="ui-g-12">
            <h2>Tags</h2>
            <hr className="finding-rule" />
            <ul>{tags}</ul>
          </div>
          <div className="ui-g-12">
            <h2>Recipes</h2>
            <hr className="finding-rule" />
            <ul>{recipes}</ul>
          </div>
        </div>
      </div>
    );
  }
}

export default FindingDetails;
