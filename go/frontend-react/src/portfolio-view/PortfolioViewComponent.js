/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { Chart } from 'primereact/chart';
import { Dropdown } from 'primereact/dropdown';
import { Toast } from 'primereact/toast';

import PortfolioViewService from './PortfolioViewService';
import { pushErrorNotification } from '../util/NotificationUtil';
import DropDownUtil from '../util/DropDownUtil';

export default class PortfolioViewComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      selectedApplication: '',
      selectedApi: '',
      selectedLanguage: '',
      top5LanguagesByLocData: [],
      top5ApisByScoreData: [],
      languageItems: [],
      appsByLanguageData: [],
      appsByApiData: [],
      apiItems: [],
      chartWidth: '600px',
      chartHeight: '300px',
      chartPadding: '0px 0px 0px 50px',
      dropDownWidth: '90%'
    };

    this.onSelectedApiChange = this.onSelectedApiChange.bind(this);
    this.onSelectedLanguageChange = this.onSelectedLanguageChange.bind(this);
  }

  onSelectedApiChange(e) {
    this.setState({ selectedApi: e.value }, () => {
      this.fetchAppsByApi(this.props.selectedRun, this.state.selectedApi);
    });
  }

  onSelectedLanguageChange(e) {
    this.setState({ selectedLanguage: e.value }, () => {
      this.fetchAppsByLanguage(
        this.props.selectedRun,
        this.state.selectedLanguage
      );
    });
  }

  populateLanguagesAndApis(selectedRun) {
    this.fetchTop5LanguagesByLoc(selectedRun);
    this.fetchTop5ApisByScore(selectedRun);
    this.fetchAllLanguages(selectedRun);
    this.fetchAllApis(selectedRun);
  }

  fetchAllLanguages(runid) {
    PortfolioViewService()
      .getLanguagesByLoc(runid)
      .then(resp => {
        this.setState(
          {
            languageItems: DropDownUtil().generateLanguageItemsDropDown(resp)
          },
          () => {
            this.setState({
              selectedLanguage: DropDownUtil().setSelectedLanguage(
                this.state.languageItems
              )
            });
          }
        );
      })
      .catch(err => pushErrorNotification(err, this.toast))
      .then(() => this.fetchAppsByLanguage(runid, this.state.selectedLanguage));
  }

  fetchAllApis(runid) {
    PortfolioViewService()
      .getApisByScore(runid)
      .then(resp => {
        this.setState(
          { apiItems: DropDownUtil().generateApiItemsDropDown(resp) },
          () => {
            this.setState({
              selectedApi: DropDownUtil().setSelectedApi(this.state.apiItems)
            });
          }
        );
      })
      .catch(err => pushErrorNotification(err, this.toast))
      .then(() => this.fetchAppsByApi(runid, this.state.selectedApi));
  }

  fetchTop5LanguagesByLoc(runid) {
    PortfolioViewService()
      .getLanguagesByLoc(runid)
      .then(resp => {
        if (resp != null) {
          this.setState({ top5LanguagesByLocData: resp.slice(0, 5) });
        }
      })
      .catch(err => pushErrorNotification(err, this.toast));
  }

  fetchTop5ApisByScore(runid) {
    PortfolioViewService()
      .getApisByScore(runid)
      .then(resp => {
        console.log('fetched top 5 apis with lenght => ' + resp.length);
        if (resp != null) {
          this.setState({ top5ApisByScoreData: resp.slice(0, 5) });
        }
      })
      .catch(err => pushErrorNotification(err, this.toast));
  }

  fetchAppsByLanguage(runid, language) {
    PortfolioViewService()
      .getAppsByLanguage(runid, language)
      .then(resp => {
        if (resp != null) {
          this.setState({ appsByLanguageData: resp.slice(0, 10) });
        }
      })
      .catch(err => pushErrorNotification(err, this.toast));
  }

  fetchAppsByApi(runid, api) {
    PortfolioViewService()
      .getAppsByApi(runid, api)
      .then(resp => {
        console.log('fetched apps by api!');
        this.setState({ appsByApiData: resp.slice(0, 10) });
      })
      .catch(err => pushErrorNotification(err, this.toast));
  }

  componentDidMount() {
    this.populateLanguagesAndApis(this.props.selectedRun.id);
  }

  UNSAFE_componentWillReceiveProps(newProps) {
    this.populateLanguagesAndApis(newProps.selectedRun.id);
  }

  render() {
    const DATASET_COLORS = [
      '#135680',
      '#1E88E5',
      '#00a79d',
      '#083330',
      '#229be6',
      '#1fccc0',
      '#d4d9d9',
      '#cc1f27',
      '#788e9a',
      '#00776d'
    ];

    const top5LanguageData = {
      labels: this.state.top5LanguagesByLocData.map(lang => lang['language']),
      datasets: [
        {
          label: 'LOC',
          backgroundColor: DATASET_COLORS,
          borderColor: DATASET_COLORS,
          data: this.state.top5LanguagesByLocData.map(lang => lang['codeLines'])
        }
      ]
    };

    const top5LanguageOptions = {
      title: {
        display: true,
        text: 'Top 5 Languages',
        fontSize: 16
      },
      legend: {
        display: false
      },
      animation: {
        duration: 0
      },
      scales: {
        yAxes: [
          {
            scaleLabel: {
              display: true,
              labelString: 'Language',
              fontSize: 16
            }
          }
        ],
        xAxes: [
          {
            scaleLabel: {
              display: true,
              labelString: 'Lines of Code',
              fontSize: 16
            },
            ticks: {
              beginAtZero: true
            }
          }
        ]
      }
    };

    const top5ApiData = {
      labels: this.state.top5ApisByScoreData.map(lang => lang['api']),
      datasets: [
        {
          label: 'Usages',
          backgroundColor: DATASET_COLORS,
          borderColor: DATASET_COLORS,
          data: this.state.top5ApisByScoreData.map(lang => lang['usageCount'])
        }
      ]
    };

    const top5ApiOptions = {
      title: {
        display: true,
        text: 'Top 5 APIs',
        fontSize: 16
      },
      legend: {
        display: false
      },
      animation: {
        duration: 0
      },
      scales: {
        yAxes: [
          {
            scaleLabel: {
              display: true,
              labelString: 'API',
              fontSize: 16
            }
          }
        ],
        xAxes: [
          {
            scaleLabel: {
              display: true,
              labelString: 'Usages',
              fontSize: 16
            },
            ticks: {
              beginAtZero: true
            }
          }
        ]
      }
    };

    const appUsingLanguageData = {
      labels: this.state.appsByLanguageData.map(lang => lang['application']),
      datasets: [
        {
          label: 'LOC',
          backgroundColor: DATASET_COLORS,
          borderColor: DATASET_COLORS,
          data: this.state.appsByLanguageData.map(lang => lang['codeLines'])
        }
      ]
    };

    const appUsingLanguageOptions = {
      title: {
        display: true,
        text: `Top 10 Apps using language: ${
          this.state.selectedLanguage != null ? this.state.selectedLanguage : ''
        }`,
        fontSize: 16
      },
      legend: {
        display: false
      },
      animation: {
        duration: 0
      },
      scales: {
        yAxes: [
          {
            scaleLabel: {
              display: true,
              labelString: 'Application',
              fontSize: 16
            }
          }
        ],
        xAxes: [
          {
            scaleLabel: {
              display: true,
              labelString: 'Lines of code',
              fontSize: 16
            },
            ticks: {
              beginAtZero: true
            }
          }
        ]
      }
    };

    const appUsingApiData = {
      labels: this.state.appsByApiData.map(lang => lang['application']),
      datasets: [
        {
          label: 'Usages',
          backgroundColor: DATASET_COLORS,
          borderColor: DATASET_COLORS,
          data: this.state.appsByApiData.map(lang => lang['usageCount'])
        }
      ]
    };

    const appUsingApiOptions = {
      title: {
        display: true,
        text: `Top 10 Apps using API: ${
          this.state.selectedApi != null ? this.state.selectedApi : ''
        }`,
        fontSize: 16
      },
      legend: {
        display: false
      },
      animation: {
        duration: 0
      },
      scales: {
        yAxes: [
          {
            scaleLabel: {
              display: true,
              labelString: 'Application',
              fontSize: 16
            }
          }
        ],
        xAxes: [
          {
            scaleLabel: {
              display: true,
              labelString: 'Usages',
              fontSize: 16
            },
            ticks: {
              beginAtZero: true
            }
          }
        ]
      }
    };

    return (
      <div>
        <p-toast ref={el => (this.toast = el)} position="bottomright" />
        <div className="ui-g ui-fluid dashboard">
          <div className="ui-g-12">
            <div className="ui-g-12 ui-g-nopad">
              <div className="ui-g-12 ui-md-6">
                <Chart
                  type="horizontalBar"
                  data={top5LanguageData}
                  options={top5LanguageOptions}
                  width={this.state.chartWidth}
                  height={this.state.chartHeight}
                  style={{ padding: this.state.chartPadding }}
                />
              </div>
              <div className="ui-g-12 ui-md-6">
                <Chart
                  type="horizontalBar"
                  data={top5ApiData}
                  options={top5ApiOptions}
                  width={this.state.chartWidth}
                  height={this.state.chartHeight}
                  style={{ padding: this.state.chartPadding }}
                />
              </div>
            </div>
          </div>
          <div className="ui-g-12">
            <div className="ui-g-12 ui-g-nopad">
              <div className="ui-g-12 ui-md-6">
                <div className="ui-g-12 ui-g-nopad">
                  <div className="card summary chart-dropdown">
                    <div>
                      <label htmlFor={'langD'} className="csa-title">
                        Select Language
                      </label>
                    </div>
                    <div>
                      <Dropdown
                        id={'langD'}
                        value={this.state.selectedLanguage}
                        options={this.state.languageItems}
                        onChange={this.onSelectedLanguageChange}
                        style={{ width: this.state.dropDownWidth }}
                      />
                    </div>
                  </div>
                </div>
                <div className="ui-g-12">
                  <Chart
                    type="horizontalBar"
                    data={appUsingLanguageData}
                    options={appUsingLanguageOptions}
                    width={this.state.chartWidth}
                    height={this.state.chartHeight}
                    style={{ padding: this.state.chartPadding }}
                  />
                </div>
              </div>
              <div className="ui-g-12 ui-md-6">
                <div className="ui-g-12 ui-g-nopad">
                  <div className="card summary chart-dropdown">
                    <div>
                      <label htmlFor={'apiD'} className="csa-title">
                        Select API
                      </label>
                    </div>
                    <div>
                      <Dropdown
                        id={'apiD'}
                        value={this.state.selectedApi}
                        options={this.state.apiItems}
                        onChange={this.onSelectedApiChange}
                        style={{ width: this.state.dropDownWidth }}
                      />
                    </div>
                  </div>
                  <div className="ui-g-12">
                    <Chart
                      type="horizontalBar"
                      data={appUsingApiData}
                      options={appUsingApiOptions}
                      width={this.state.chartWidth}
                      height={this.state.chartHeight}
                      style={{ padding: this.state.chartPadding }}
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
