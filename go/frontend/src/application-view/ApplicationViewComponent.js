/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { Dropdown } from 'primereact/dropdown';
import { Toast } from 'primereact/toast';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { InputText } from 'primereact/inputtext';
import CustomTreeMapContent from '../util/CustomTreeMapContent';
import { ResponsiveContainer, Treemap } from 'recharts';
import {
  pushErrorNotification,
  pushNotification,
  sleep
} from '../util/NotificationUtil';
import DropDownUtil from '../util/DropDownUtil';
import TreeMapUtil from '../util/TreeMapUtil';
import { PrimaryButton } from 'pivotal-ui/react/buttons/buttons';
import ApplicationViewService from './ApplicationViewService';
import ScoreCard from '../shared/ScoreCard';
import DataViewService from '../data-view/DataViewService';
import FindingDetails from '../util/FindingDetails';
import { Sidebar } from 'primereact/sidebar';
import { TabPanel, TabView } from 'primereact/tabview';
import { DashboardCard } from '../shared/DashboardCard';
import { Checkbox } from 'primereact/checkbox';
import { MultiSelect } from 'primereact/multiselect';

const levelItems = [
  { label: 'All', value: 'All' },
//  { label: 'Info', value: 'info' },
  { label: 'Low', value: 'low' },
  { label: 'Medium', value: 'medium' },
  { label: 'High', value: 'high' }
];

const emptyScoreCard = {
  application: 'n/a',
  tags: [],
  app: { score: '0', rawScore: '0' },
  //info: '0',
  low: '0',
  medium: '0',
  high: '0',
  total: '0'
};

export default class ApplicationViewComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      applications: [],
      applicationItems: [],
      applicationScoreCard: emptyScoreCard,
      applicationDetails: [],
      applicationLevelDetails: [],
      languageTreeMap: [],
      apiTreeMap: [],
      findings: [],
      loading: true,
      finding: {},
      sidebarVisible: false,
      sidebarPosition: 'left',
      activeIndex: 0,
      currentApp: { slocCnt: 0, filesCnt: 0 },
      cardBackground: 'count card-background',
      includeFF: false
    };
    this.onSelectLevelChange = this.onSelectLevelChange.bind(this);
    this.onSelectApplicationChange = this.onSelectApplicationChange.bind(this);
    this.export = this.export.bind(this);
    this.tableLoaded = this.tableLoaded.bind(this);
    this.filenameTemplate = this.filenameTemplate.bind(this);
    this.idTemplate = this.idTemplate.bind(this);
    this.recipeTemplate = this.recipeTemplate.bind(this);
    this.tagTemplate = this.tagTemplate.bind(this);
    this.getFindingDetails = this.getFindingDetails.bind(this);
    this.showFindingSideBar = this.showFindingSideBar.bind(this);
    this.filterLevel = this.filterLevel.bind(this);
    this.onSelectedTagChange = this.onSelectedTagChange.bind(this);
    this.onSelectedBinChange = this.onSelectedBinChange.bind(this);
    this.onIncludeChange = this.onIncludeChange.bind(this);
    this.effortTemplate = this.effortTemplate.bind(this);
  }

  export() {
    this.dt.exportCSV();
  }

  onSelectApplicationChange(e) {
    this.setState(
      { loading: true, currentApp: { slocCnt: 0, filesCnt: 0 } },
      () => {
        this.props.appSetter(e.value, () =>
          this.populateAppStateData(
            this.props.selectedRun.id,
            this.props.application,
            this.props.levels,
            this.props.tags,
            this.props.bins,
            this.state.includeFF
          )
        );
      }
    );
  }

  onSelectLevelChange(e) {
    this.setState({ findings: [] });
    this.props.levelSetter(e.value, () => {
      this.filterLevel(e.value);
    });
    this.fetchApplicationScoreCardDetails(
      this.props.selectedRun.id,
      this.props.application,
      e.value,
      this.props.tags,
      this.state.includeFF
    );
  }

  onSelectedTagChange(tag) {
    this.setState({ loading: true }, () => {
      this.props.tagSetter(tag, () => {
        this.populateAppStateData(
          this.props.selectedRun.id,
          this.props.application,
          this.props.levels,
          this.props.tags,
          this.props.bins,
          this.state.includeFF
        );
      });
    });
  }

  onSelectedBinChange(bin) {
    this.setState({ loading: true }, () => {
      this.props.binSetter(bin, () => {
        this.populateAppStateData(
          this.props.selectedRun.id,
          this.props.application,
          this.props.levels,
          this.props.tags,
          this.props.bins,
          this.state.includeFF
        );
      });
    });
  }

  onIncludeChange(e) {
    console.log('Include FF changing...now checked [%s]', e.checked);

    this.setState(
      { includeFF: e.checked },
      this.populateAppStateData(
        this.props.selectedRun.id,
        this.props.application,
        this.props.levels,
        this.props.tags,
        this.props.bins,
        e.checked
      )
    );
  }

  filterLevel(levels) {
    if (levels.includes('All')) {
      levels = null;
    }
    if (this.dt) {
      this.dt.filter(levels, 'level', 'in');
    }
  }

  filenameTemplate(rowData) {
    return (
      <div className="tooltip">
        {rowData.filename}
        <span className="tooltiptext pui-toggle-switch">{rowData.fqn}</span>
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
      //console.log('Finding [' + rowData.id + '] recipes => ' + JSON.stringify(rowData.recipes));

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

  tagTemplate(rowData) {
    if (rowData.tags) {
      //console.log('Finding [' + rowData.id + '] recipes => ' + JSON.stringify(rowData.recipes));

      return (
        <span
          className="item-link"
          onClick={() => this.getFindingDetails(rowData.id, 'left')}
        >
          {rowData.tags.length}
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

  effortTemplate(rowData) {
    let className = '';
    let note = '';
    if (rowData.effort > 10) {
      className = 'crit-background';
      note = 'Critical finding! Effort > 10';
    } else if (rowData.effort < 0) {
      className = 'boost-background';
      note = 'Positive finding! Effort < 0';
    } else if (
      rowData.effort === 0 &&
      (rowData.note && rowData.note.includes('squash'))
    ) {
      className = 'squash-background';
      note = rowData.note;
    }

    return (
      <div className="tooltip">
        <span className={className}>{rowData.effort}</span>
        <span className="tooltiptext pui-toggle-switch">{note}</span>
      </div>
    );
  }

  fetchApplications(runid, app, tags) {
    // console.log(
    //   'AppView-fetching applications for runid [%s] app[%s]',
    //   runid,
    //   app
    // );

    ApplicationViewService()
      .getApplicationByRun(runid)
      .then(resp => {
        this.setState(
          {
            applications: resp.appScores,
            applicationItems: DropDownUtil().generateApplicationItemsDropDown(
              resp.appScores
            )
          },
          () => {
            this.props.appSetter(
              DropDownUtil().setSelectedApplication(
                this.state.applicationItems,
                app
              )
            );
          }
        );
      })
      .then(() => {
        this.populateAppStateData(
          runid,
          this.props.application,
          ['low', 'medium', 'high'],
          this.props.tags,
          this.props.bins,
          this.state.includeFF
        );
      })
      .catch(err => pushErrorNotification(err, this.toast));
  }

  populateAppStateData(
    selectedRun,
    selectedApplication,
    selectedLevels,
    selectedTags,
    selectedBins,
    includeFF
  ) {
    // console.log(
    //   'AppView - populateAppStateData run[%s] app[%s] level[%s] tags[%s]',
    //   selectedRun,
    //   selectedApplication,
    //   selectedLevels,
    //   JSON.stringify(selectedTags)
    // );

    if (selectedApplication !== '') {
      this.fetchApplicationScoreCard(
        selectedRun,
        selectedApplication,
        selectedTags,
        selectedBins,
        includeFF
      );
      this.fetchApplicationScoreCardDetails(
        selectedRun,
        selectedApplication,
        selectedLevels,
        selectedTags,
        includeFF
      );
      this.fetchLanguagesByApp(selectedRun, selectedApplication);
      this.fetchApisByApp(selectedRun, selectedApplication);
    } else {
      this.setState({
        loading: false,
        applicationScoreCard: emptyScoreCard,
        languageTreeMap: [],
        apiTreeMap: [],
        findings: []
      });
    }
  }

  fetchApplicationScoreCardDetails(runid, app, levels, tags, includeFF) {
    console.log('IncludeFF => %s', this.state.includeFF);

    ApplicationViewService()
      .getApplicationScoreCardDetailsByLevel(
        runid,
        app,
        levels,
        tags,
        includeFF
      )
      .then(resp => this.setState({ findings: resp }))
      .then(() => this.filterLevel(this.props.levels))
      .catch(err => pushErrorNotification(err, this.toast))
      .then(() => this.tableLoaded());
  }

  fetchApplicationScoreCard(selectedRun, application, tags, bins, includeFF) {
    ApplicationViewService()
      .getApplicationScoreCard(selectedRun, application, tags, includeFF)
      .then(resp => {
        this.setupScoreCard(resp, application, tags, bins);
      })
      .catch(err => pushErrorNotification(err, this.toast));
  }

  fetchLanguagesByApp(runid, app) {
    ApplicationViewService()
      .getLanguagesByApplicationRun(runid, app)
      .then(resp => {
        if (resp != null) {
          this.setState({
            languageTreeMap: TreeMapUtil().generateLanguageTreeMap(resp)
          });
        }
      })
      .catch(err => pushErrorNotification(err, this.toast));
  }

  fetchApisByApp(runid, app) {
    ApplicationViewService()
      .getApisByApplicationRun(runid, app)
      .then(resp => {
        if (resp != null) {
          this.setState({
            apiTreeMap: TreeMapUtil().generateApisTreeMap(resp)
          });
        }
      })
      .catch(err => pushErrorNotification(err, this.toast));
  }

  setupScoreCard(scoreCard, application, tags, bins) {
    //console.log('App View => processing new scorecard!');
    let currentApp = { slocCnt: 0, filesCnt: 0 };
    this.state.applications.forEach(app => {
      if (application === app.name) {
        scoreCard.app = app;
        currentApp = app;

        if (tags.length > 0) {
          // console.log(
          //   'Updating scorecard with stateful tags [%s]',
          //   JSON.stringify(tags)
          // );
          scoreCard.tags = tags;
        } else {
          // console.log(
          //   'Updating scorecard with tags from app [%s]',
          //   JSON.stringify(tags)
          // );
          scoreCard.tags = [];
          if (app.tags && app.tags.length > 0) {
            app.tags.forEach(tag => {
              scoreCard.tags.push({ name: tag.Value, selected: false });
            });
          }
          this.props.appTagsSetter(scoreCard.tags);
        }

        if (bins.length > 0) {
           console.log('Updating scorecard with stateful bins [%s]', JSON.stringify(bins));
          scoreCard.bins = bins;
        } else {
          console.log('Bins are zero');
          scoreCard.bins = [];
          if (app.bins && app.bins.length > 0) {
            app.bins.forEach(bin => {
              bin.selected = false;
              scoreCard.bins.push(bin);
            });
          }
          this.props.appBinsSetter(scoreCard.bins);
        }
      }
    });
    this.setState({ applicationScoreCard: scoreCard, currentApp: currentApp });
  }

  componentDidMount() {
    console.log('AppView - didMount');

    //   console.log('AppView - didMount run [%s] app[%s] level[%s] tags: %s',
    // this.props.selectedRun.id,
    // this.props.application,
    // this.props.level,
    // JSON.stringify(this.props.tags));

    this.setState({ loading: true }, () =>
      this.fetchApplications(
        this.props.selectedRun.id,
        this.props.application,
        this.props.tags
      )
    );
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    console.log('AppView - willReceiveProps');
    if (newProps.selectedRun.id !== this.props.selectedRun.id) {
      // console.log('AppView - run id or tags has changed.... existing [%s] new [%s] existing-tags[%s] newTags[%s]',
      //   this.props.selectedRun.id,
      //   newProps.selectedRun.id,
      //   JSON.stringify(this.props.tags),
      //   JSON.stringify(newProps.tags)
      // );

      this.setState({ loading: true }, () =>
        this.fetchApplications(
          newProps.selectedRun.id,
          newProps.application,
          newProps.tags
        )
      );
    }
  }

  async tableLoaded() {
    await sleep(100);
    this.setState({ loading: false });
    pushNotification(
      'Found [' + this.state.findings.length + '] findings!',
      this.toast
    );
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

  // rowClassName(rowData) {
  //   return {
  //     'state-highlight-red': rowData.effort > 10,
  //     'state-highlight-green': rowData.effort < 0
  //   };
  // }

  render() {
    const checked = this.state.includeFF;
    const header = (
      <div>
        <div className="ui-g">
          <div className="ui-g-12 ui-md-1" style={{ textAlign: 'left' }}>
            <PrimaryButton type="button" aria-label="CSV" onClick={this.export}>
              CSV
            </PrimaryButton>
          </div>
          <div className="ui-g-12 ui-md-5" style={{ textAlign: 'left' }}>
            <Checkbox
              inputId="cb1"
              onChange={this.onIncludeChange}
              checked={checked}
            />
            <label htmlFor="cb1">Include File Findings</label>
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

    const levelFilter = (
      <MultiSelect
        value={this.props.levels}
        options={levelItems}
        onChange={this.onSelectLevelChange}
        style={{ width: '100%' }}
      />
    );

    return (
      <div>
        <Toast ref={el => (this.toast = el)} position="bottomright" />
        <div className="ui-g ui-fluid dashboard">
          <div className="ui-g-12 ui-md-6">
            <div className="card summary">
              <div>
                <label htmlFor={'apiDA'} className="csa-title">
                  Select Application
                </label>
              </div>
              <div>
                <Dropdown
                  id={'appDA'}
                  value={this.props.application}
                  options={this.state.applicationItems}
                  onChange={this.onSelectApplicationChange}
                  style={{ width: '90%' }}
                  className={'af-dropdown'}
                  filter={true}
                />
              </div>
            </div>
          </div>
          <div className="ui-g-12 ui-md-3">
            <DashboardCard
              title={'SLOC'}
              detail={'Lines of Code'}
              summaryType={this.state.currentApp.slocCnt}
              backgroundclass={this.state.cardBackground}
            />
          </div>
          <div className="ui-g-12 ui-md-3">
            <DashboardCard
              title={'Files'}
              detail={'# of Files'}
              summaryType={this.state.currentApp.filesCnt}
              backgroundclass={this.state.cardBackground}
            />
          </div>
        </div>
        <TabView
          className="tabview"
          activeIndex={this.state.activeIndex}
          onTabChange={e => this.setState({ activeIndex: e.index })}
        >
          <TabPanel header="Findings">
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

            <ScoreCard
              applicationScoreCard={this.state.applicationScoreCard}
              tagSetter={this.onSelectedTagChange}
              binSetter={this.onSelectedBinChange}
            />

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
              rowsPerPageOptions={[100, 250, 500, 1000]}
              //rowClassName={this.rowClassName}
              ref={el => {
                this.dt = el;
              }}
            >
              <Column
                field="id"
                header="ID"
                sortable={true}
                body={this.idTemplate}
                style={{ textAlign: 'center', width: '75px' }}
              />
              <Column
                field="level"
                header="Level"
                sortable={true}
                filter="true"
                filterElement={levelFilter}
                style={{ textAlign: 'center', width: '100px' }}
              />
              <Column
                field="category"
                header="Category"
                sortable={true}
                filter="true"
                filterMatchMode="contains"
                style={{ textAlign: 'center', width: '150px' }}
              />
              <Column
                field="value"
                header="Value"
                filter="true"
                filterMatchMode="contains"
                style={{ width: '700px' }}
              />
              <Column
                field="filename"
                header="Filename"
                sortable={true}
                filter="true"
                filterMatchMode="contains"
                body={this.filenameTemplate}
                style={{ width: '400px' }}
              />
              <Column
                field="line"
                header="Line #"
                sortable={true}
                filter="true"
                filterMatchMode="contains"
                style={{ width: '75px' }}
              />
              <Column
                field="effort"
                header="Effort"
                sortable={true}
                filter="true"
                filterMatchMode="contains"
                style={{ width: '75px' }}
                body={this.effortTemplate}
              />
              <Column
                field="tags"
                header="# Tag"
                body={this.tagTemplate}
                style={{ width: '75px' }}
              />
              <Column
                field="recipes"
                header="# Recipes"
                body={this.recipeTemplate}
                style={{ width: '100px' }}
              />
            </DataTable>
          </TabPanel>
          <TabPanel header="Languages">
            <div className="ui-g ui-fluid">
              <div className="ui-g-12">
                <label>
                  <b>Languages for Application:</b> {this.props.application}
                </label>
                <ResponsiveContainer width="99%" height={600}>
                  <Treemap
                    data={
                      this.state.languageTreeMap.length > 0
                        ? this.state.languageTreeMap
                        : [{}]
                    }
                    dataKey="size"
                    stroke="#fff"
                    fill="#1bb3a8"
                    isAnimationActive={false}
                    isUpdateAnimationActive={false}
                    content={<CustomTreeMapContent colors={['#1bb3a8']} />}
                  />
                </ResponsiveContainer>
              </div>
            </div>
          </TabPanel>
          <TabPanel header="API's">
            <div className="ui-g ui-fluid">
              <div className="ui-g-12">
                <label>
                  <b>API's For Application:</b> {this.props.application}
                </label>
                <ResponsiveContainer width="99%" height={500}>
                  <Treemap
                    data={
                      this.state.apiTreeMap.length > 0
                        ? this.state.apiTreeMap
                        : [{}]
                    }
                    dataKey="size"
                    stroke="#fff"
                    fill="#135680"
                    isAnimationActive={false}
                    isUpdateAnimationActive={false}
                    content={<CustomTreeMapContent colors={['#135680']} />}
                  />
                </ResponsiveContainer>
              </div>
            </div>
          </TabPanel>
        </TabView>
      </div>
    );
  }
}
