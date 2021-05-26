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
import FindingDetails from '../../util/FindingDetails';
import { Sidebar } from 'primereact/sidebar';

export default class FindingsComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      findings: [],
      loading: true,
      finding: {},
      sidebarVisible: false,
      sidebarPosition: 'left'
    };
    this.export = this.export.bind(this);
    this.tableLoaded = this.tableLoaded.bind(this);
    this.filenameTemplate = this.filenameTemplate.bind(this);
    this.idTemplate = this.idTemplate.bind(this);
    this.recipeTemplate = this.recipeTemplate.bind(this);
    this.getFindingDetails = this.getFindingDetails.bind(this);
    this.showFindingSideBar = this.showFindingSideBar.bind(this);
  }

  fetchRunFindings(runid) {
    DataViewService()
      .getRunFindings(runid)
      .then(resp => this.setState({ findings: resp }))
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
      this.fetchRunFindings(newProps.selectedRun);
    }
  }

  componentDidMount() {
    this.fetchRunFindings(this.props.selectedRun);
  }

  async tableLoaded() {
    await sleep(100);
    this.setState({ loading: false });
    pushNotification(
      'Found [' + this.state.findings.length + '] findings!',
      this.toast
    );
  }

  filenameTemplate(rowData) {
    return (
      <div className="tooltip">
        {rowData.filename}
        <span className="tooltiptext">{rowData.fqn}</span>
      </div>
    );
  }

  idTemplate(rowData) {
    return (
      <span
        className="item-link"
        onClick={() => this.getFindingDetails(rowData.id, 'right')}
      >
        {rowData.id}
      </span>
    );
  }

  recipeTemplate(rowData) {
    if (rowData.recipes) {
      console.log(
        'Finding [' +
          rowData.id +
          '] recipes => ' +
          JSON.stringify(rowData.recipes)
      );

      return (
        <span
          className="item-link"
          onClick={() => this.getFindingDetails(rowData.id, 'left')}
        >
          {rowData.recipes.length}
        </span>
      );
    } else {
      return (
        <span
          className="item-link"
          onClick={() => this.getFindingDetails(rowData.id, 'left')}
        >
          0
        </span>
      );
    }
  }

  getFindingDetails(findingId, position) {
    DataViewService()
      .getFinding(findingId)
      .then(resp => this.showFindingSideBar(resp, position))
      .catch(err => {
        if (err.response) {
          pushErrorNotification(err.response.data, this.toast);
        } else {
          pushErrorNotification(err, this.toast);
        }
      });
  }

  showFindingSideBar(finding, position) {
    this.setState({
      finding: finding,
      sidebarVisible: true,
      sidebarPosition: position
    });
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
        <p-toast ref={el => (this.toast = el)} position="bottomright" />
        <Sidebar
          position={this.state.sidebarPosition}
          className="ui-sidebar-md"
          blockScroll={true}
          visible={this.state.sidebarVisible}
          baseZIndex={1000000}
          onHide={e => this.setState({ sidebarVisible: false })}
        >
          <FindingDetails color="#138078" finding={this.state.finding} />
        </Sidebar>
        <DataTable
          value={this.state.findings}
          loading={this.state.loading}
          responsive={true}
          globalFilter={this.state.globalFilter}
          autoLayout={true}
          header={header}
          tableStyle={{ tableLayout: 'auto' }}
          paginator={true}
          rows={100}
          rowsPerPageOptions={[50, 100, 250, 500]}
          ref={el => {
            this.dt = el;
          }}
        >
          <Column
            field="id"
            header="ID"
            sortable={true}
            body={this.idTemplate}
          />
          <Column
            field="application"
            header="Application"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />

          <Column
            field="category"
            header="Category"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column field="value" header="Value" filterMatchMode="contains" />
          <Column
            field="filename"
            header="Filename"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
            body={this.filenameTemplate}
          />
          <Column
            field="line"
            header="Line #"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="effort"
            header="Score"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="recipes"
            header="# Recipes"
            body={this.recipeTemplate}
          />
        </DataTable>
      </div>
    );
  }
}
