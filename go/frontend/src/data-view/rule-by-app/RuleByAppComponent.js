/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { InputText } from 'primereact/inputtext';
import { Growl } from 'primereact/growl';

import { pushErrorNotification, sleep } from '../../util/NotificationUtil';
import DataViewService from '../DataViewService';
import { PrimaryButton } from 'pivotal-ui/react/buttons/buttons';

export default class RuleByAppComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      cols: [],
      usageData: []
    };
    this.export = this.export.bind(this);
    this.tableLoaded = this.tableLoaded.bind(this);
    //    this.loadApiColumnData = this.loadApiColumnData.bind(this);
    this.populateRuleData = this.populateRuleData.bind(this);
  }

  export() {
    this.dt.exportCSV();
  }

  fetchRuleUsage(runid) {
    this.setState({ loading: true });
    DataViewService()
      .getRuleByAppUsage(runid)
      .then(resp => {
        this.setState({ cols: resp.cols });
        this.populateRuleData(resp);
      })
      .catch(err => pushErrorNotification(err, this.growl))
      .then(() => this.tableLoaded());
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    if (newProps.selectedRun !== this.props.selectedRun) {
      this.setState({ loading: true });
      this.fetchRuleUsage(newProps.selectedRun);
    }
  }

  populateRuleData(ruleScoreMatrix) {
    console.log('Populating usage data...');

    let cols = ruleScoreMatrix.cols;
    let data = ruleScoreMatrix.data;
    let usageData = [];
    data.forEach(app => {
      let ruleScoreRecord = { App: app.application };
      cols.forEach(col => {
        app.ruleusage.forEach(usage => {
          if (usage.rule === col) {
            ruleScoreRecord[col] = usage.count;
          }
        });
      });
      usageData.push(ruleScoreRecord);
    });

    this.setState({ usageData: usageData });

    console.log('Usage rows = [%d]', usageData.length);
  }

  componentDidMount() {
    this.fetchRuleUsage(this.props.selectedRun);
  }

  async tableLoaded() {
    await sleep(100);
    this.setState({ loading: false });
  }

  render() {
    const header = (
      <div>
        <div className="ui-g">
          <div className="ui-g-12 ui-md-6" style={{ textAlign: 'left' }}>
            <PrimaryButton type="button" aria-label="CSV" onClick={this.export}>
              CSV
            </PrimaryButton>
          </div>
          <div className="ui-g-12 ui-md-6" style={{ textAlign: 'right' }}>
            <InputText
              type="search"
              onInput={e => this.setState({ globalFilter: e.target.value })}
              placeholder="Search"
            />
          </div>
        </div>
      </div>
    );

    const cols = this.state.cols;

    let dynamicColumns = cols.map((col, i) => {
      if (col.toLowerCase() === 'app') {
        return (
          <Column
            key={col + i}
            field={col}
            header={col.toUpperCase()}
            headerStyle={{
              width: '400px',
              textAlign: 'center',
              fontWeight: 'bold',
              fontSize: '14px'
            }}
            style={{ width: '400px' }}
          />
        );
      } else {
        return (
          <Column
            key={col + i}
            field={col}
            header={col.toUpperCase()}
            headerStyle={{
              width: '100px',
              textAlign: 'center',
              fontWeight: 'bold',
              wordWrap: 'break-word',
              fontSize: '14px'
            }}
            style={{ width: '100px', textAlign: 'right' }}
          />
        );
      }
    });

    return (
      <div>
        <Growl ref={el => (this.growl = el)} position="bottomright" />
        <DataTable
          value={this.state.usageData}
          loading={this.state.loading}
          responsive={true}
          globalFilter={this.state.globalFilter}
          header={header}
          scrollable={true}
          scrollHeight="550px"
          scrollWidth="100%"
          ref={el => {
            this.dt = el;
          }}
        >
          {dynamicColumns}
        </DataTable>
      </div>
    );
  }
}
