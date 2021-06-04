/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { TabPanel, TabView } from 'primereact/tabview';

import ApiUsageDetailedComponent from './api-usage-detailed/ApiUsageDetailedComponent';
import ApiUsageSummaryComponent from './api-usage-summary/ApiUsageSummaryComponent';
import AnnotationsComponent from './annotations/AnnotationsComponent';
import SourceCodeComponent from './source-code/SourceCodeComponent';
import ThirdPartyComponent from './third-party/ThirdPartyComponent';
import FindingsComponent from './findings-all/FindingsComponent';
import SearchHtmlTableComponent from './search/SearchHtmlTableComponent';
import DataViewService from './DataViewService';
import { pushErrorNotification } from '../util/NotificationUtil';
import ApiByAppComponent from './api-by-app/ApiByAppComponent';
import RuleMetricsComponent from './rule-metrics/RuleMetricsComponent';
import RuleByAppComponent from './rule-by-app/RuleByAppComponent';

export default class DataViewComponent extends Component {
  constructor(props) {
    super(props);
    this.determineIndexExists = this.determineIndexExists.bind(this);

    this.state = {
      activeIndex: 0,
      index: { exists: false }
    };
  }

  determineIndexExists(run) {
    DataViewService()
      .getRunIndexDetails(run)
      .then(resp => {
        if (resp.index.exists) {
          this.setState({ index: resp.index });
        } else {
          this.setState({ index: resp.index, activeIndex: 0 });
        }
      })
      .catch(err => {
        if (err.response) {
          pushErrorNotification(err.response.data, this.toast);
        } else {
          pushErrorNotification(err, this.toast);
        }
        this.setState({ index: { exists: false, activeIndex: 0 } });
      });
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    if (this.props.selectedRun.id !== newProps.selectedRun.id) {
      this.determineIndexExists(newProps.selectedRun.id);
    }
  }

  componentDidMount() {
    this.determineIndexExists(this.props.selectedRun.id);
  }

  render() {
    const indexexists = this.state.index.exists;

    return (
      <div>
        <div>
          <TabView
            activeIndex={this.state.activeIndex}
            onTabChange={e => this.setState({ activeIndex: e.index })}
          >
            <TabPanel header="API By App">
              <ApiByAppComponent selectedRun={this.props.selectedRun.id} />
            </TabPanel>
            <TabPanel header="API Usage (Detailed)">
              <ApiUsageDetailedComponent
                selectedRun={this.props.selectedRun.id}
              />
            </TabPanel>
            <TabPanel header="API Usage (Summary)">
              <ApiUsageSummaryComponent
                selectedRun={this.props.selectedRun.id}
              />
            </TabPanel>
            <TabPanel header="Annotations">
              <AnnotationsComponent selectedRun={this.props.selectedRun.id} />
            </TabPanel>
            <TabPanel header="Third Party Libs">
              <ThirdPartyComponent selectedRun={this.props.selectedRun.id} />
            </TabPanel>
            <TabPanel header="Source Code">
              <SourceCodeComponent selectedRun={this.props.selectedRun.id} />
            </TabPanel>
            <TabPanel header="All Findings">
              <FindingsComponent selectedRun={this.props.selectedRun.id} />
            </TabPanel>
            <TabPanel header="Rule By App">
              <RuleByAppComponent selectedRun={this.props.selectedRun.id} />
            </TabPanel>
            <TabPanel header="Rule Metrics">
              <RuleMetricsComponent selectedRun={this.props.selectedRun.id} />
            </TabPanel>
            <TabPanel header="Search" disabled={!indexexists}>
              <SearchHtmlTableComponent
                selectedRun={this.props.selectedRun.id}
              />
            </TabPanel>
          </TabView>
        </div>
      </div>
    );
  }
}
