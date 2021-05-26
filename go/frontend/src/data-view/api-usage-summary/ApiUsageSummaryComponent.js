/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { InputText } from 'primereact/inputtext';
import { Toast } from 'primereact/toast';
import DataViewService from '../DataViewService';
import { pushErrorNotification, sleep } from '../../util/NotificationUtil';
import { PrimaryButton } from 'pivotal-ui/react/buttons/buttons';

export default class ApiUsageSummaryComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      findings: [],
      loading: true
    };
    this.export = this.export.bind(this);
    this.tableLoaded = this.tableLoaded.bind(this);
  }

  export() {
    this.dt.exportCSV();
  }

  fetchApiSummaryUsage(runid) {
    DataViewService()
      .getApiSummaryUsage(runid)
      .then(resp => this.setState({ findings: resp }))
      .catch(err => pushErrorNotification(err, this.toast))
      .then(() => this.tableLoaded());
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    if (newProps.selectedRun !== this.props.selectedRun) {
      this.setState({ loading: true });
      this.fetchApiSummaryUsage(newProps.selectedRun);
    }
  }

  componentDidMount() {
    this.fetchApiSummaryUsage(this.props.selectedRun);
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

    return (
      <div>
        <p-toast ref={el => (this.toast = el)} position="bottomright" />
        <DataTable
          value={this.state.findings}
          loading={this.state.loading}
          responsive={true}
          globalFilter={this.state.globalFilter}
          header={header}
          tableStyle={{ tableLayout: 'auto' }}
          ref={el => {
            this.dt = el;
          }}
          paginator={true}
          rows={50}
          rowsPerPageOptions={[10, 25, 50, 100]}
          sortField="usageCount"
          sortOrder={-1}
        >
          <Column
            field="api"
            header="API"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="usageCount"
            header="Usage Count"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
        </DataTable>
      </div>
    );
  }
}
