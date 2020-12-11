/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { InputText } from 'primereact/inputtext';
import { Growl } from 'primereact/growl';

import { RulesService } from './RulesService';
import NavBarUtil from '../util/NavBarUtil';
import { pushErrorNotification } from '../util/NotificationUtil';
import { PrimaryButton } from 'pivotal-ui/react/buttons/buttons';

export default class RulesTableComponent extends Component {
  constructor() {
    super();
    this.state = {
      rules: []
    };
    this.rowExpansionTemplate = this.rowExpansionTemplate.bind(this);
    this.tagFilter = this.tagFilter.bind(this);
    this.export = this.export.bind(this);
  }

  fetchRules() {
    const rulesService = new RulesService();
    rulesService
      .getRules()
      .then(rulesResponse => {
        rulesResponse.map(rule => {
          const tags =
            rule.Tags != null
              ? rule.Tags.map(tag => '[' + tag.Value + ']\n')
              : 'none';
          return (rule.Tags = tags);
        });

        this.setState({ rules: rulesResponse });
      })
      .catch(err => pushErrorNotification(err, this.growl));
  }

  tagFilter(value, filter) {
    const cellValue = '' + value;
    return cellValue.includes(filter);
  }

  rowExpansionTemplate(data) {
    const patterns =
      data.Patterns != null
        ? data.Patterns.map((pattern, index) => (
            <li key={index}>{pattern.Value}</li>
          ))
        : 'none';
    const recipes =
      data.Recipes != null
        ? data.Recipes.map((recipe, index) => (
            <li key={index}>
              <a href={recipe.URI}>{recipe.URI}</a>
            </li>
          ))
        : 'none';
    const defaultPattern =
      data.DefaultPattern != null ? data.DefaultPattern : 'none';

    return (
      <div className="ui-g ui-fluid">
        <div className="ui-g-12 ui-md-9">
          <div className="ui-g">
            <div className="ui-md-2" style={{ fontWeight: 'bold' }}>
              Default Pattern:{' '}
            </div>
            <div className="ui-md-10" style={{ fontWeight: 'bold' }}>
              {defaultPattern}
            </div>

            <div className="ui-md-2" style={{ fontWeight: 'bold' }}>
              Patterns:{' '}
            </div>
            <div className="ui-md-10" style={{ fontWeight: 'bold' }}>
              {patterns}
            </div>

            <div className="ui-md-2" style={{ fontWeight: 'bold' }}>
              Recipes:{' '}
            </div>
            <div className="ui-md-10" style={{ fontWeight: 'bold' }}>
              {recipes}
            </div>
          </div>
        </div>
      </div>
    );
  }

  componentDidMount() {
    NavBarUtil().setNavBarLinkActive(4);
    this.fetchRules();
  }

  componentWillUnmount() {
    NavBarUtil().unsetNavBarLinkActive(4);
  }

  export() {
    this.dtr.exportCSV();
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
            <i className="fa fa-search" style={{ margin: '4px 4px 0 0' }} />
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
          value={this.state.rules}
          responsive={true}
          globalFilter={this.state.globalFilter}
          header={header}
          expandedRows={this.state.expandedRows}
          onRowToggle={e => this.setState({ expandedRows: e.data })}
          rowExpansionTemplate={this.rowExpansionTemplate}
          tableStyle={{ tableLayout: 'auto' }}
          ref={el => {
            this.dtr = el;
          }}
          paginator={true}
          rows={25}
          rowsPerPageOptions={[10, 25, 50, 100]}
        >
          <Column expander={true} style={{ width: '2em' }} />
          <Column
            field="Name"
            header="Name"
            sortable={true}
            filter="true"
            filterMatchMode="contains"
          />
          <Column
            field="Tags"
            header="Tags"
            sortable={true}
            filter="true"
            filterMatchMode="custom"
            filterFunction={this.tagFilter}
          />
          <Column
            field="FileType"
            header="FileType"
            sortable={true}
            filter="true"
          />
          <Column
            field="Target"
            header="Target"
            sortable={true}
            filter="true"
          />
          <Column field="Type" header="Type" sortable={true} filter="true" />
          <Column
            field="Advice"
            header="Advice"
            sortable={true}
            filter="true"
          />
          <Column
            field="Effort"
            header="Effort"
            sortable={true}
            filter="true"
          />
          <Column
            field="Category"
            header="Category"
            sortable={true}
            filter="true"
          />
          <Column
            field="Criticality"
            header="Criticality"
            sortable={true}
            filter="true"
          />
        </DataTable>
      </div>
    );
  }
}
