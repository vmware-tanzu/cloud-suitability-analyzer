/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { InputText } from 'primereact/inputtext';
import DataViewService from '../DataViewService';
import {
  pushErrorNotification,
  pushNotification,
  sleep
} from '../../util/NotificationUtil';
import { Toast } from 'primereact/toast';
import { PrimaryButton } from 'pivotal-ui/react/buttons/buttons';

export default class RuleMetricsComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      metrics: [],
      loading: true
    };
    this.export = this.export.bind(this);
    this.tableLoaded = this.tableLoaded.bind(this);
  }

  fetchRunMetrics(runid) {
    DataViewService()
      .getRuleMetrics(runid)
      .then(resp => this.setState({ metrics: resp.metrics }))
      .catch(err => {
        if (err.response) {
          pushErrorNotification(err.response.data, this.toast);
        } else {
          pushErrorNotification(err, this.toast);
        }
      })
      .then(() => this.tableLoaded());
  }

  export() {
    this.dt.exportCSV();
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    if (newProps.selectedRun !== this.props.selectedRun) {
      this.setState({ loading: true });
      this.fetchRunMetrics(newProps.selectedRun);
    }
  }

  componentDidMount() {
    this.fetchRunMetrics(this.props.selectedRun);
  }

  async tableLoaded() {
    await sleep(100);
    this.setState({ loading: false });
    pushNotification(
      'Found [' + this.state.metrics.length + '] metrics!',
      this.toast
    );
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
              onInput={e =>
                this.setState({
                  globalFilter: e.target.value
                })
              }
              placeholder="Search"
            />
          </div>
        </div>
      </div>
    );

    return (
      <div>
        <Toast ref={el => (this.toast = el)} position="bottomright" />

        <DataTable
          value={this.state.metrics}
          loading={this.state.loading}
          responsive={true}
          globalFilter={this.state.globalFilter}
          autoLayout={true}
          header={header}
          tableStyle={{ tableLayout: 'auto' }}
          scrollable={true}
          scrollHeight="550px"
          ref={el => {
            this.dt = el;
          }}
        >
          <Column
            field="rule"
            header="Rule"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="checks"
            header="Checks"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="patternChecks"
            header="PatternChecks"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
            body={this.filenameTemplate}
          />
          <Column
            field="hits"
            header="# Hits"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="totalTime"
            header="Total Time"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="longest"
            header="Longest"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="shortest"
            header="Shortest"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="avgRule"
            header="Avg-Rule"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="avgPat"
            header="Avg-Pattern"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="avgHit"
            header="Avg-Hit"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
        </DataTable>
      </div>
    );
  }
}
