/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { InputText } from 'primereact/inputtext';
import { Toast } from 'primereact/toast';

import { pushErrorNotification, sleep } from '../../util/NotificationUtil';
import DataViewService from '../DataViewService';
import { PrimaryButton } from 'pivotal-ui/react/buttons/buttons';

export default class ApiByAppComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      cols: [],
      usageData: []
    };
    this.export = this.export.bind(this);
    this.tableLoaded = this.tableLoaded.bind(this);
    this.populateApiData = this.populateApiData.bind(this);
  }

  export() {
    this.dt.exportCSV();
  }

  fetchApiUsage(runid) {
    this.setState({ loading: true });
    DataViewService()
      .getApiByAppUsage(runid)
      .then(resp => {
        this.setState({ cols: resp.cols });
        this.populateApiData(resp);
      })
      .catch(err => pushErrorNotification(err, this.toast))
      .then(() => this.tableLoaded());
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    if (newProps.selectedRun !== this.props.selectedRun) {
      this.setState({ loading: true });
      this.fetchApiUsage(newProps.selectedRun);
    }
  }

  populateApiData(apiUsageMatrix) {
    console.log('Populating usage data...');

    let cols = apiUsageMatrix.cols;
    let data = apiUsageMatrix.data;
    let usageData = [];
    data.forEach(app => {
      let apiUsageRecord = { App: app.application };
      cols.forEach(col => {
        app.apiusage.forEach(usage => {
          if (usage.api === col) {
            apiUsageRecord[col] = usage.usageCount;
          }
        });
      });
      usageData.push(apiUsageRecord);
    });

    this.setState({ usageData: usageData });
  }

  componentDidMount() {
    this.fetchApiUsage(this.props.selectedRun);
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
        <Toast ref={el => (this.toast = el)} position="bottomright" />
        <DataTable
          value={this.state.usageData}
          loading={this.state.loading}
          responsive={true}
          globalFilter={this.state.globalFilter}
          header={header}
          // tableStyle={{ tableLayout: 'auto' }}

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
