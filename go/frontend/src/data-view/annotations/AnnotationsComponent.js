/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { InputText } from 'primereact/inputtext';
import DataViewService from '../DataViewService';
import { pushErrorNotification, sleep } from '../../util/NotificationUtil';
import { Growl } from 'primereact/components/growl/Growl';
import { PrimaryButton } from 'pivotal-ui/react/buttons/buttons';

export default class AnnotationsComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      findings: [],
      loading: true
    };
    this.export = this.export.bind(this);
    this.tableLoaded = this.tableLoaded.bind(this);
  }

  fetchAnnotationData(runid) {
    DataViewService()
      .getAnnotationData(runid)
      .then(resp => this.setState({ findings: resp }))
      .catch(err => pushErrorNotification(err, this.growl))
      .then(() => this.tableLoaded());
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    if (newProps.selectedRun !== this.props.selectedRun) {
      this.setState({ loading: true });
      this.fetchAnnotationData(newProps.selectedRun);
    }
  }

  componentDidMount() {
    this.fetchAnnotationData(this.props.selectedRun);
  }

  export() {
    this.dt.exportCSV();
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
        <Growl ref={el => (this.growl = el)} position="bottomright" />
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
        >
          <Column
            field="application"
            header="Application"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="value"
            header="Value"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="pattern"
            header="Pattern"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="advice"
            header="Advice"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="effort"
            header="Effort"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="level"
            header="Level"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
        </DataTable>
      </div>
    );
  }
}
