/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { Growl } from 'primereact/growl';
import { TabPanel, TabView } from 'primereact/tabview';
import {
  pushErrorNotification,
  pushNotification,
  sleep
} from '../util/NotificationUtil';
import ExecutiveSummaryService from './ExecutiveSummaryService';
import { PrimaryButton } from 'pivotal-ui/react/buttons/buttons';
import { InputText } from 'primereact/components/inputtext/InputText';
import { DashboardCard } from '../shared/DashboardCard';
import { Spinner } from 'primereact/spinner';
import RecommendationTreeMap from './RecommendationTreeMap';
import RecommendationBar from './RecommendationBar';

export default class ExecutiveSummaryViewComponent extends Component {
  componentWillUnmount() {
    if (this.chart) {
      this.chart.dispose();
    }
  }

  constructor(props) {
    super(props);

    this.state = {
      applicationScores: [],
      numAppsByRun: 0,
      locByRun: 0,
      numFilesByRun: 0,
      readinessChartData: [],
      cardBackground: 'count card-background',
      activeIndex: 0,
      filteredData: [],
      filter: '',
      loading: true,
      selecectedApp: null,
      tableEditable: true
    };
    this.export = this.export.bind(this);
    this.defaultAppBizValue = this.defaultAppBizValue.bind(this);
    this.filterTable = this.filterTable.bind(this);
    this.appSelected = this.appSelected.bind(this);
    this.attributeChange = this.attributeChange.bind(this);
    this.numericEditor = this.numericEditor.bind(this);
    this.numbersMatch = this.numbersMatch.bind(this);
    this.performFilter = this.performFilter.bind(this);
    this.textEditor = this.textEditor.bind(this);
    this.techScoreTemplate = this.techScoreTemplate.bind(this);
    this.appTemplate = this.appTemplate.bind(this);
  }

  appSelected(item) {
    this.setState({ tableEditable: false }, () =>
      this.props.appNavigator(item.name)
    );
  }

  filterTable(filterValue) {
    if (this.state.tableEditable) {
      this.setState({ tableEditable: false }, () =>
        this.performFilter(filterValue)
      );
    } else {
      this.performFilter(filterValue);
    }
  }

  performFilter(filterValue) {
    if (this.state.filterTimer) {
      clearTimeout(this.state.filterTimer);
    }

    let fv = filterValue.toLowerCase();
    let filteredData = [];
    this.state.applicationScores.forEach(app => {
      if (app.name.toLowerCase().includes(fv)) {
        filteredData.push(app);
      } else if (!isNaN(parseFloat(fv)) && isFinite(fv)) {
        let flVal = parseFloat(fv);
        let dec = Math.floor(flVal);

        if (this.numbersMatch(flVal, dec, app.score)) {
          filteredData.push(app);
        } else if (this.numbersMatch(flVal, dec, app.businessvalue)) {
          filteredData.push(app);
        }
      } else if (app.businessdomain.toLowerCase().includes(fv)) {
        filteredData.push(app);
      } else if (app.recommendation.toLowerCase().includes(fv)) {
        filteredData.push(app);
      }
    });
    this.setState({ filteredData: filteredData }, () => {
      let filterTimer = setTimeout(() => {
        this.generateReadinessChart();
        this.setState({ tableEditable: true, filterTimer: null });
      }, 500);
      this.setState({ filter: filterValue, filterTimer: filterTimer });
    });
  }

  numbersMatch(flVal, dec, value) {
    if (dec === Math.floor(value)) {
      if (flVal === dec) {
        return true;
      } else {
        let decArray = flVal.toString().split('.');
        if (decArray.length > 1) {
          let fltDecimals = decArray[1];
          let decPlaces = fltDecimals.length || 0;
          decArray = value.toString().split('.');
          if (decArray.length > 1) {
            let scoreDecimals = decArray[1].substring(0, decPlaces);
            if (fltDecimals === scoreDecimals) {
              return true;
            }
          }
        }
      }
    }
  }

  export() {
    this.dt.exportCSV();
  }

  /* Populate summary and application scores for run */
  populateSummaryAndApplicationScores(selectedRun) {
    this.setState({ loading: true, filter: '' }, () => {
      this.fetchSummaryFindingsByRun(selectedRun.id);
      this.fetchAppScoresByRun(selectedRun.id);
    });
  }

  /* Populate application scores for run */
  fetchAppScoresByRun(runid) {
    ExecutiveSummaryService()
      .getApplicationScoresByRun(runid)
      .then(resp => this.setupAppScores(resp))
      .catch(err => pushErrorNotification(err, this.growl))
      .then(() => this.defaultAppBizValue());
  }

  /* Populate # of apps, lines of code, and # of files for run */
  fetchSummaryFindingsByRun(runid) {
    ExecutiveSummaryService()
      .getSummaryFindingsByRun(runid)
      .then(resp => {
        this.summaryStateSetter(resp);
      })
      .catch(err => pushErrorNotification(err, this.growl));
  }

  summaryStateSetter(resp) {
    this.setState({
      numAppsByRun: resp.applicationCount,
      locByRun: resp.codeLines,
      numFilesByRun: resp.totalFiles
    });
  }

  setupAppScores(resp) {
    let filteredData = [];
    if (resp.appScores) {
      resp.appScores.forEach(app => {
        filteredData.push(app);
      });
    }
    this.setState({
      applicationScores: resp.appScores,
      filteredData: filteredData,
      loading: false
    });
  }

  generateReadinessChart() {
    let chartData = [];
    if (this.state.filteredData.length > 0) {
      this.state.filteredData.forEach(app => {
        let chartRecord = {};

        //console.log("%s",JSON.stringify(app));

        //This is just for doing z-axis...not useful otherwise!
        let appRawScore = app.rawScore;

        if (appRawScore < 10) {
          appRawScore = 10;
        }
        chartRecord.y = app.businessvalue;
        chartRecord.x = app.score;
        chartRecord.a = app.name;
        chartRecord.name = app.name;
        chartRecord.z = appRawScore;
        chartRecord.s = app.slocCnt;
        chartRecord.size = app.slocCnt;
        chartRecord.fc = app.filesCnt;
        chartRecord.f = app.findings;
        chartRecord.r = app.recommendation;
        chartData.push(chartRecord);
      });
    }
    this.setState({ readinessChartData: chartData });
  }

  techScoreTemplate(rowData, column) {
    if (rowData.scoreModified) {
      return (
        <div className="tooltip">
          <span style={{ color: '#f27062' }}>{rowData.score}</span>
          <span className="tooltiptext pui-toggle-switch">
            {rowData.originalScore}
          </span>
        </div>
      );
    } else {
      return <span>{rowData.score}</span>;
    }
  }

  appTemplate(rowData) {
    return (
      <span className="item-link" onClick={() => this.appSelected(rowData)}>
        {rowData.name}
      </span>
    );
  }

  numericEditor(props) {
    if (props && props.rowIndex >= 0) {
      let rowData = props.value[props.rowIndex];
      let idPre = props.rowIndex + '-' + props.field + '-';
      return (
        <Spinner
          id={idPre + 'sp'}
          value={rowData[props.field]}
          style={{ width: '6em' }}
          step={0.05}
          min={0}
          max={10}
          onChange={e => this.attributeChange(props, e.value)}
        />
      );
    }
  }

  textEditor(props) {
    if (props && props.rowIndex >= 0) {
      let rowData = props.value[props.rowIndex];
      let idPre = props.rowIndex + '-' + props.field + '-';
      return (
        <InputText
          id={idPre + 'txt'}
          value={rowData[props.field]}
          style={{ width: '6em' }}
          onChange={e => this.attributeChange(props, e.target.value)}
        />
      );
    }
  }

  attributeChange(props, val) {
    if (this.state.sliderTimer) {
      clearTimeout(this.state.sliderTimer);
    }

    let i;
    for (i = 0; i < this.state.applicationScores.length; i++) {
      let app = this.state.applicationScores[i];
      if (app.name === props.value[props.rowIndex]['name']) {
        app[props.field] = val;
        let sliderTimer = setTimeout(() => {
          if (props.field === 'score') {
            app.scoreModified = app.score !== app.originalScore;
          }
          ExecutiveSummaryService()
            .updateApp(app)
            .catch(err => pushErrorNotification(err, this.growl));
          this.setState(
            { applicationScores: this.state.applicationScores },
            () => this.generateReadinessChart()
          );
        }, 250);
        this.setState({ sliderTimer: sliderTimer });
      }
    }
  }

  defaultAppBizValue() {
    this.state.applicationScores.forEach(app => {
      if (app.businessvalue < 0 || app.businessvalue > 10) {
        return (app.businessvalue = 0);
      }
    });
    this.generateReadinessChart(this.state.applicationScores);
    //this.setState({applicationScores: this.state.applicationScores});
  }

  componentDidMount() {
    if (typeof this.props.selectedRun.id != 'undefined') {
      this.populateSummaryAndApplicationScores(this.props.selectedRun);
    }
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    if (newProps.selectedRun !== this.props.selectedRun) {
      this.populateSummaryAndApplicationScores(newProps.selectedRun);
    }
  }

  async tableLoaded() {
    await sleep(100);
    this.setState({ loading: false });
    pushNotification(
      'Found [' + this.state.applicationScores + '] applications!',
      this.growl
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
              value={this.state.filter}
              onInput={e => {
                this.filterTable(e.target.value);
              }}
              placeholder="Search"
            />
          </div>
        </div>
      </div>
    );

    var appCnt = this.state.filteredData ? this.state.filteredData.length : 0;
    const footer = appCnt + ' applications (filtered) showing for portfolio...';

    const chartData = this.state.readinessChartData;

    return (
      <div className="ui-g-nopad">
        <Growl
          id="growl1"
          ref={el => (this.growl = el)}
          position="bottomright"
        />
        <div className="ui-g ui-fluid ui-g-nopad dashboard">
          <div className="ui-g-12 ui-g-nopad">
            <div className="ui-g-12 ui-md-3">
              <DashboardCard
                title={'APPS'}
                detail={'# of applications'}
                summaryType={this.state.numAppsByRun}
                backgroundclass={this.state.cardBackground}
              />
            </div>
            <div className="ui-g-12 ui-md-3">
              <DashboardCard
                title={'LOC'}
                detail={'Lines of code'}
                summaryType={this.state.locByRun}
                backgroundclass={this.state.cardBackground}
              />
            </div>
            <div className="ui-g-12 ui-md-3">
              <DashboardCard
                title={'FILES'}
                detail={'# of Files'}
                summaryType={this.state.numFilesByRun}
                backgroundclass={this.state.cardBackground}
              />
            </div>
            <div className="ui-g-12 ui-md-3">
              <DashboardCard
                title={'FINDINGS'}
                detail={'# of Findings'}
                summaryType={this.props.selectedRun.Findings}
                backgroundclass={this.state.cardBackground}
              />
            </div>
          </div>
          <TabView
            className="tabview chart-box"
            activeIndex={this.state.activeIndex}
            onTabChange={e => this.setState({ activeIndex: e.index })}
          >
            <TabPanel header="Scoring Summary">
              {/* <div className="ui-g-12 ui-g-nopad chart-box"> */}
              <div className="ui-g-12 ui-g-nopad">
                <div className="ui-g-12">
                  <DataTable
                    value={this.state.filteredData}
                    responsive={true}
                    scrollable={true}
                    scrollHeight="440px"
                    tableStyle={{ tableLayout: 'auto' }}
                    sortField="name"
                    sortOrder={1}
                    ref={el => {
                      this.dt = el;
                    }}
                    globalFilter={this.state.globalFilter}
                    header={header}
                    footer={footer}
                    style={{ textAlign: 'center' }}
                    loading={this.state.loading}
                    editable={this.state.tableEditable}
                  >
                    <Column
                      field="name"
                      header="Application"
                      headerStyle={{ textAlign: 'center' }}
                      sortable={true}
                      body={this.appTemplate}
                      style={{ textAlign: 'left' }}
                    />
                    <Column
                      field="slocCnt"
                      header="LOC"
                      sortable={true}
                      style={{ textAlign: 'center', width: '12%' }}
                    />
                    <Column
                      field="filesCnt"
                      header="# Files"
                      sortable={true}
                      style={{ textAlign: 'center', width: '12%' }}
                    />
                    <Column
                      field="rawScore"
                      header="Raw Score"
                      sortable={true}
                      style={{ textAlign: 'center', width: '12%' }}
                    />
                    <Column
                      field="model"
                      header="Scoring Model"
                      sortable={true}
                      style={{ textAlign: 'center', width: '12%' }}
                    />

                    {this.state.tableEditable && (
                      <Column
                        field="score"
                        header="Technical Score"
                        sortable={true}
                        editor={this.numericEditor}
                        body={this.techScoreTemplate}
                        style={{ textAlign: 'center', width: '12%' }}
                      />
                    )}
                    {!this.state.tableEditable && (
                      <Column
                        field="score"
                        header="Technical Score"
                        sortable={true}
                        body={this.techScoreTemplate}
                        style={{ textAlign: 'center', width: '12%' }}
                      />
                    )}

                    {/* 
                    Removed business value and 2x2 grid, keeping here
                    in case we get push back                   
                    {this.state.tableEditable && (
                      <Column
                        header="Business Value"
                        field="businessvalue"
                        editor={this.numericEditor}
                        style={{ textAlign: 'center', width: '8%' }}
                      />
                    )}
                    {!this.state.tableEditable && (
                      <Column
                        header="Business Value"
                        field="businessvalue"
                        style={{ textAlign: 'center', width: '8%' }}
                      />
                    )}

 */}

                    {/*                     {this.state.tableEditable && (
                      <Column
                        field="businessdomain"
                        header="Business Domain"
                        headerStyle={{ textAlign: 'center', width: '12%' }}
                        sortable={true}
                        style={{ textAlign: 'left', width: '12%' }}
                        editor={this.textEditor}
                      />
                    )}
                    {!this.state.tableEditable && (
                      <Column
                        field="businessdomain"
                        header="Business Domain"
                        headerStyle={{ textAlign: 'center', width: '12%' }}
                        sortable={true}
                        style={{ textAlign: 'left', width: '12%' }}
                      />
                    )} */}

                    {/*                     <Column
                      field="recommendation"
                      header="Recommendation"
                      headerStyle={{ textAlign: 'center', width: '15%' }}
                      sortable={true}
                      style={{ textAlign: 'left', width: '15%' }}
                    />
       */}
                  </DataTable>
                </div>

                {/*                 <div className="ui-g-12 ui-md-4 ui-g-nopad">
                  <div>
                    <ReactVisChartComponent
                      data={chartData}
                      bound={true}
                      height={570}
                      loading={this.state.loading}
                      labels={true}
                    />
                  </div>
                </div> */}
              </div>
            </TabPanel>
            <TabPanel header="Chart(s)">
              <div className="ui-g-12 ui-g-nopad">
                <div className="ui-g-6">
                  <RecommendationTreeMap
                    data={chartData}
                    labels={true}
                    loading={this.state.loading}
                  />
                </div>
                <div className="ui-g-6">
                  <RecommendationBar
                    data={chartData}
                    labels={true}
                    loading={this.state.loading}
                  />
                </div>
              </div>
            </TabPanel>
          </TabView>
        </div>
      </div>
    );
  }
}
