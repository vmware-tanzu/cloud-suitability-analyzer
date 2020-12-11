/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { Growl } from 'primereact/growl';
import { Dropdown } from 'primereact/dropdown';

import { pushErrorNotification } from '../util/NotificationUtil';
import NavBarUtil from '../util/NavBarUtil';
import DropDownUtil from '../util/DropDownUtil';
import RunContainerService from './RunContainerService';
import { Menubar } from 'primereact/components/menubar/Menubar';
import ExecutiveSummaryViewComponent from '../executive-summary-view/ExecutiveSummaryViewComponent';
import PortfolioViewComponent from '../portfolio-view/PortfolioViewComponent';
import ApplicationViewComponent from '../application-view/ApplicationViewComponent';
import RulesTableComponent from '../rules-view/RulesTableComponent';
import DataViewComponent from '../data-view/DataViewComponent';
import RunDetails from './RunDetails';

function DetermineView(props) {
  const view = props.view;
  const run = props.selectedRun;
  const app = props.selectedApp;
  const tags = props.tags;
  const levels = props.levels;
  const bins = props.bins;

  // console.log('View [%s] Run[%s] App [%s] Levels [%s]',view, run.id, app, levels);

  if (view === 'Executive Summary') {
    return (
      <ExecutiveSummaryViewComponent
        selectedRun={run}
        appNavigator={props.appNavigator}
      />
    );
  } else if (view === 'Portfolio') {
    return <PortfolioViewComponent selectedRun={run} />;
  } else if (view === 'Application') {
    return (
      <ApplicationViewComponent
        selectedRun={run}
        appSetter={props.appSetter}
        levelSetter={props.levelSetter}
        tagSetter={props.tagSetter}
        binSetter={props.binSetter}
        application={app}
        appTagsSetter={props.appTagsSetter}
        appBinsSetter={props.appBinsSetter}
        levels={levels}
        tags={tags}
        bins={bins}
      />
    );
  } else if (view === 'Data') {
    return <DataViewComponent selectedRun={run} />;
  } else if (view === 'Rule') {
    return <RulesTableComponent />;
  } else {
    return (
      <h3>
        UNKNOWN VIEW{' '}
        <em>
          <i>{view}</i>
        </em>{' '}
        SELECTED!
      </h3>
    );
  }
}

export default class RunContainer extends Component {
  constructor(props) {
    super(props);
    this.onSelectRunChange = this.onSelectRunChange.bind(this);
    this.setSelectedApp = this.setSelectedApp.bind(this);
    this.setSelectedLevel = this.setSelectedLevel.bind(this);
    this.setSelectedTag = this.setSelectedTag.bind(this);
    this.setAppTags = this.setAppTags.bind(this);
    this.setSelectedBin = this.setSelectedBin.bind(this);
    this.setAppBins = this.setAppBins.bind(this);
    this.checkOtherBinsForTagAndSelect = this.checkOtherBinsForTagAndSelect.bind(
      this
    );
    this.navigateToApp = this.navigateToApp.bind(this);

    const app_param = new URLSearchParams(window.location.href.split('?')[1]).get("app");  
    if (app_param) { 
      this.state = {
        runIdItems: [],
        selectedRun: '',
        view: 'Application',
        lastView: 'unknown',
        selectedApp: app_param,
        tags: [],
        bins: [],
        levels: ['low', 'medium', 'high'],
        version: 'xxx.xxx.xxx'
      };
      // Remove parameter passed over
      window.history.pushState("object or string", "CSA", "/#/application")
    } else {
      this.state = {
        runIdItems: [],
        selectedRun: '',
        view: 'Executive Summary',
        lastView: 'unknown',
        selectedApp: '',
        tags: [],
        bins: [],
        levels: ['low', 'medium', 'high'],
        version: 'xxx.xxx.xxx'
      };
    }
  }

  navigateToApp(app) {
    console.log('Navigating to App [%s]', app);

    if (this.state.selectedApp === app) {
      this.updateView('Application');
      NavBarUtil().setNavBarLinkActive(2);
      window.location.hash = '/application';
    } else {
      this.setState(
        {
          selectedApp: app,
          tags: [],
          bins: [],
          levels: ['low', 'medium', 'high']
        },
        () => {
          this.updateView('Application');
          NavBarUtil().setNavBarLinkActive(2);
          window.location.hash = '/application';
        }
      );
    }
  }

  setSelectedApp(app, callback) {
    if (this.state.selectedApp !== app) {
      this.setState(
        {
          selectedApp: app,
          tags: [],
          bins: [],
          levels: ['low', 'medium', 'high']
        },
        callback
      );
    }
  }

  setAppTags(tags, callback) {
    this.setState({ tags: tags }, callback);
  }

  setSelectedTag(tag, callback) {
    this.state.tags.forEach(t => {
      if (tag.name === t.name) {
        t.selected = !tag.selected;
      }
    });
    if (!tag.selected) {
      this.state.bins.forEach(b => {
        b.matched.forEach(m => {
          if (m === tag.name) {
            b.selected = false;
          }
        });
      });
    }
    this.setState({ tags: this.state.tags, bins: this.state.bins }, callback);
  }

  setAppBins(bins, callback) {
    this.setState({ bins: bins }, callback);
  }

  setSelectedBin(bin, callback) {
    this.state.bins.forEach(b => {
      if (bin.name === b.name) {
        b.selected = !bin.selected;
      }
    });
    this.state.tags.forEach(t => {
      bin.matched.forEach(m => {
        if (m === t.name) {
          if (bin.selected) {
            t.selected = bin.selected;
          } else {
            t.selected = this.checkOtherBinsForTagAndSelect(bin.name, m);
          }
        }
      });
    });
    this.setState({ bins: this.state.bins, tags: this.state.tags }, callback);
  }

  checkOtherBinsForTagAndSelect(bin, tag) {
    return this.state.bins.some(b => {
      if (b.name !== bin && b.selected) {
        return b.matched.some(t => {
          if (t === tag) {
            return true;
          }
          return false;
        });
      }
      return false;
    });
  }

  setSelectedLevel(levels, callback) {
    // console.log('setting stored state for selected level => ' + levels);
    if (this.state.levels !== levels) {
      this.setState({ levels: levels }, callback);
    }
  }

  /* On run change, set selected Run state and populate summary and app scores */
  onSelectRunChange(e) {
    this.setState({
      selectedRun: e.value,
      levels: ['low', 'medium', 'high'],
      tags: [],
      bins: []
    });
  }

  updateView(view) {
    this.setState({ view: view, lastView: this.state.view });
  }

  /* Populate run dropdown, select latest run, and set selected run in dropdown */
  fetchDistinctRunIds() {
    RunContainerService()
      .getDistinctRuns()
      .then(resp => {
        this.setState((prevState, props) => {
          let items = DropDownUtil().generateRunItemsDropDown(resp);
          let run = DropDownUtil().setSelectedRun(items);
          return { runIdItems: items, selectedRun: run };
        });
      })
      .catch(err => pushErrorNotification(err, this.growl));
  }

  fetchVersionInfo() {
    RunContainerService()
      .getForgeVersion()
      .then(resp => {
        this.setState({ version: resp });
      })
      .catch(err => pushErrorNotification(err, this.growl));
  }

  startContainer() {
    this.fetchVersionInfo();
    this.fetchDistinctRunIds();
  }

  componentDidMount() {
    NavBarUtil().setNavBarLinkActive(0);
    this.startContainer();
  }

  render() {
    const items = [
      {
        label: 'Summary',
        command: () => {
          this.updateView('Executive Summary');
          NavBarUtil().setNavBarLinkActive(0);
          window.location.hash = '/summary';
        }
      },
      {
        label: 'Portfolio',
        command: () => {
          this.updateView('Portfolio');
          NavBarUtil().setNavBarLinkActive(1);
          window.location.hash = '/portfolio';
        }
      },
      {
        label: 'Application',
        command: () => {
          this.updateView('Application');
          NavBarUtil().setNavBarLinkActive(2);
          window.location.hash = '/application';
        }
      },
      {
        label: 'Data',
        command: () => {
          this.updateView('Data');
          NavBarUtil().setNavBarLinkActive(3);
          window.location.hash = '/data';
        }
      },
      {
        label: 'Rule',
        command: () => {
          this.updateView('Rule');
          NavBarUtil().setNavBarLinkActive(4);
          window.location.hash = '/rules';
        }
      }
    ];

    const selectedRun = this.state.selectedRun;
    const view = this.state.view;
    const selectedApp = this.state.selectedApp;
    const levels = this.state.levels;
    const tags = this.state.tags;
    const bins = this.state.bins;

    return (
      <div>
        <Growl
          id="growl1"
          ref={el => (this.growl = el)}
          position="bottomright"
        />
        <div className="ui-g">
          <div className="ui-g-12 ui-g-nopad">
            <div className="ui-g-1">
              <img
                id="logo"
                src="/VMware-logo.svg"
                width="60"
                height="60"
                alt="vmware-logo"
                style={{ padding: '1px' }}
              />
            </div>
            <div className="ui-g-6 ui-g-nopad">
              <div className="ui-g-12 ui-g-nopad">
                <div className="ui-g-6 ui-g-nopad">
                  <div className="ui-g-12">
                    <label htmlFor={'runDE'} className="csa-title">
                      Select Run
                    </label>
                  </div>
                  <div className="ui-g-12 ui-g-nopad">
                    <Dropdown
                      id={'runDE'}
                      value={this.state.selectedRun}
                      options={this.state.runIdItems}
                      onChange={this.onSelectRunChange}
                      placeholder={'Select'}
                      style={{ width: '100%' }}
                    />
                  </div>
                </div>
                <div className="ui-g-6 ui-g-nopad">
                  <div className="ui-g-12 ui-g-nopad">
                    <RunDetails selectedRun={this.state.selectedRun} />
                  </div>
                  <div className="ui-g-12 ui-g-nopad" />
                </div>
              </div>
            </div>
            <div className="ui-g-5">
              <div className="ui-g-12 ui-g-nopad">
                <div className="ui-g-12, version-title">
                  {this.state.version}
                </div>
                <Menubar
                  id={'menubar'}
                  model={items}
                  style={{
                    width: '100%',
                    border: 'none',
                    background: '#edf0f5',
                    padding: '0px 0px 20px 20px'
                  }}
                />
              </div>
            </div>
          </div>
          <div className="ui-g-12 ui-g-nopad">
            <DetermineView
              view={view}
              selectedRun={selectedRun}
              appSetter={this.setSelectedApp}
              levelSetter={this.setSelectedLevel}
              tagSetter={this.setSelectedTag}
              appTagsSetter={this.setAppTags}
              binSetter={this.setSelectedBin}
              appBinsSetter={this.setAppBins}
              selectedApp={selectedApp}
              appNavigator={this.navigateToApp}
              levels={levels}
              tags={tags}
              bins={bins}
            />
          </div>
        </div>
      </div>
    );
  }
}
