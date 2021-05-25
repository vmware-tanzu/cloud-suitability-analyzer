/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { Column } from 'primereact/components/column/Column';
import { DataTable } from 'primereact/components/datatable/DataTable';
import { Toast } from 'primereact/toast';
import { InputText } from 'primereact/components/inputtext/InputText';

import DataViewService from '../DataViewService';
import { pushErrorNotification, sleep } from '../../util/NotificationUtil';
import { PrimaryButton } from 'pivotal-ui/react/buttons/buttons';

export default class SourceCodeComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      findings: [],
      loading: true
    };
    this.export = this.export.bind(this);
    this.tableLoaded = this.tableLoaded.bind(this);
  }

  fetchSourceCodeData(runid) {
    DataViewService()
      .getSourceCodeData(runid)
      .then(resp => this.setState({ findings: resp }))
      .catch(err => pushErrorNotification(err, this.toast))
      .then(() => this.tableLoaded());
  }

  export() {
    this.dt.exportCSV();
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    if (newProps.selectedRun !== this.props.selectedRun) {
      this.setState({ loading: true });
      this.fetchSourceCodeData(newProps.selectedRun);
    }
  }

  componentDidMount() {
    this.fetchSourceCodeData(this.props.selectedRun);
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
        <Toast ref={el => (this.toast = el)} position="bottomright" />
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
          sortField="codeLines"
          sortOrder={-1}
        >
          {/*Language     string `json:"language"`*/}
          {/*TotalFiles   int    `json:"totalFiles"`*/}
          {/*BlankLines   int    `json:"blankLines"`*/}
          {/*CommentLines int    `json:"commentLines"`*/}
          {/*CodeLines    int    `json:"codeLines"`*/}

          <Column
            field="language"
            header="Language"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column field="totalFiles" header="# Files" sortable={true} />
          <Column field="blankLines" header="Blank Lines" sortable={true} />
          <Column field="commentLines" header="Comment Lines" sortable={true} />
          <Column field="codeLines" header="Lines of Code" sortable={true} />
        </DataTable>
      </div>
    );
  }
}
