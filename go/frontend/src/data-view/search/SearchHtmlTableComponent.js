/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React, { Component } from 'react';
import { InputText } from 'primereact/inputtext';
import { SelectButton } from 'primereact/selectbutton';
import { Sidebar } from 'primereact/sidebar';
import DataViewService from '../DataViewService';
import {
  pushErrorNotification,
  pushNotification,
  sleep
} from '../../util/NotificationUtil';
import { Growl } from 'primereact/components/growl/Growl';
//import KeyValuePairComponent from "../../util/KeyValuePairComponent";
import FindingDetails from '../../util/FindingDetails';

export default class SearchHtmlTableComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      searchResults: [],
      hits: [],
      loading: false,
      queryText: '',
      searchMessage: 'No search results!',
      qType: 'query-string',
      maxResults: 50,
      sidebarVisible: false,
      finding: {},
      lastSearch: '',
      hitCnt: 0,
      totalHits: 0,
      tookMS: '',
      index: { exists: false }
    };
    this.export = this.export.bind(this);
    this.tableLoaded = this.tableLoaded.bind(this);
    this.performSearch = this.performSearch.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleKeyDown = this.handleKeyDown.bind(this);
    this.getSpans = this.getSpans.bind(this);
    this.processMatches = this.processMatches.bind(this);
    this.processMatch = this.processMatch.bind(this);
    this.processSearchResult = this.processSearchResult.bind(this);
    this.getFindingDetails = this.getFindingDetails.bind(this);
    this.showFindingSideBar = this.showFindingSideBar.bind(this);
    this.determineIndexExists = this.determineIndexExists.bind(this);
  }

  searchForFindings(runid, query, queryType, maxResults) {
    DataViewService()
      .searchRunFindings(runid, query, queryType, maxResults)
      .then(resp => this.processSearchResult(resp))
      .then(() => this.tableLoaded())
      .catch(err => {
        if (err.response) {
          pushErrorNotification(err.response.data, this.growl);
        } else {
          pushErrorNotification(err, this.growl);
        }
      });
  }

  export() {
    this.dt.exportCSV();
  }

  performSearch(query) {
    if (query && query !== this.state.queryText) {
      this.setState({ queryText: query });
    } else {
      query = this.state.queryText;
    }

    this.setState({ loading: true });
    this.searchForFindings(
      this.props.selectedRun,
      query,
      this.state.qType,
      this.state.maxResults
    );
  }

  handleChange(evt) {
    this.setState({ queryText: evt.target.value });
  }

  handleKeyDown(evt) {
    if (evt.keyCode === 13 && this.state.queryText.length > 0) {
      return this.performSearch();
    }
  }

  async tableLoaded() {
    await sleep(100);
    this.setState({ loading: false });
    if (this.state.searchResults) {
      var hits = 0;
      if (this.state.searchResults.Hits) {
        hits = this.state.searchResults.Hits.length;
      }

      pushNotification('Search returned [' + hits + '] findings!', this.growl);
    }
  }

  getSpans(hit) {
    let span = 0;

    if (hit.MatchedOn) {
      for (var j = 0; j < hit.MatchedOn.length; j++) {
        span += hit.MatchedOn[j].Matches.length;
      }
    } else {
      span = 1;
    }
    return span;
  }

  processMatch(rowid, match) {
    var html = [];
    var span = match.Matches.length;

    for (var k = 0; k < span; k++) {
      var tdId = rowid + '-' + k;
      if (k === 0) {
        html.push(
          <td key={tdId + 'f'} width="5%" rowSpan={span}>
            <span className="finding-data">{match.FieldName}</span>
          </td>
        );
      }
      html.push(
        <td key={tdId + 'v'}>
          <span
            className="finding-data"
            dangerouslySetInnerHTML={{ __html: match.Matches[k] }}
          />
        </td>
      );
    }

    return html;
  }

  processMatches(span, indx, hit, matches) {
    var html = [];

    var idClass = 'id-score odd';

    if (indx === 0 || indx % 2 === 0) {
      idClass = 'id-score even';
    }

    if (matches) {
      for (var j = 0; j < matches.length; j++) {
        var rowid = hit.ID + '-' + j;
        if (j > 0) {
          html.push(
            <tr key={rowid}>{this.processMatch(rowid, matches[j])}</tr>
          );
        } else {
          html.push(
            <tr key={rowid}>
              <td rowSpan={span} className={idClass} width="5%">
                <span
                  className="item-link"
                  onClick={() => this.getFindingDetails(hit.ID)}
                >
                  {hit.ID}
                </span>
              </td>
              {this.processMatch(rowid, matches[j])}
              <td rowSpan={span} className={idClass} width="15%" padding="">
                <span className="finding-data">{hit.Score}</span>
              </td>
            </tr>
          );
        }
      }
    } else {
      html.push(
        <tr key={rowid}>
          <td rowSpan={span} className={idClass} width="5%">
            <span
              className="item-link"
              onClick={() => this.getFindingDetails(hit.ID)}
            >
              {hit.ID}
            </span>
          </td>
          <td>&nbsp;</td>
          <td>&nbsp;</td>
          <td rowSpan={span} className={idClass} width="15%">
            <span className="finding-data">{hit.Score}</span>
          </td>
        </tr>
      );
    }
    return html;
  }

  processSearchResult(result) {
    var msg = 'Your search returned no results...';
    var hitCnt = 0;
    var totalHits = 0;
    var tookMS = '';
    if (result) {
      totalHits = result.Totalhits;
      tookMS = result.TookMs;

      if (result.Hits) {
        hitCnt = result.Hits.length;
      }
    }
    this.setState({
      searchMessage: msg,
      searchResults: result,
      queryText: '',
      lastSearch: this.state.queryText,
      hitCnt: hitCnt,
      totalHits: totalHits,
      tookMS: tookMS
    });
  }

  createTableBody(hits) {
    var html = [];
    for (var i = 0; i < hits.length; i++) {
      var span = this.getSpans(hits[i]);
      html.push(this.processMatches(span, i, hits[i], hits[i].MatchedOn));
    }
    return html;
  }

  createTableFooter() {
    if (this.state.hitCnt === 0 && this.state.lastSearch === '') {
      return (
        <tfoot className="search-footer">
          <tr>
            <td colSpan="4">&nbsp;</td>
          </tr>
        </tfoot>
      );
    }

    return (
      <tfoot className="search-footer">
        <tr>
          <td colSpan="4">
            Search returned [{this.state.hitCnt}] of [{this.state.totalHits}]
            total Findings matching your search terms [<span
              className="finding-query"
              onClick={() => this.performSearch(this.state.lastSearch)}
            >
              {this.state.lastSearch}
            </span>] in {this.state.tookMS}
          </td>
        </tr>
      </tfoot>
    );
  }

  getFindingDetails(findingId) {
    DataViewService()
      .getFinding(findingId)
      .then(resp => this.showFindingSideBar(resp))
      .catch(err => {
        if (err.response) {
          pushErrorNotification(err.response.data, this.growl);
        } else {
          pushErrorNotification(err, this.growl);
        }
      });
  }

  showFindingSideBar(finding) {
    this.setState({
      finding: finding,
      sidebarVisible: true
    });
  }

  determineIndexExists(run) {
    DataViewService()
      .getRunIndexDetails(run)
      .then(resp => {
        this.setState({ index: resp.index });
      })
      .catch(err => {
        if (err.response) {
          pushErrorNotification(err.response.data, this.growl);
        } else {
          pushErrorNotification(err, this.growl);
        }
        this.setState({ index: { exists: false } });
      });
  }

  componentDidUpdate(prevProps) {
    if (this.props.selectedRun !== prevProps.selectedRun) {
      this.determineIndexExists(this.props.selectedRun);
    }
  }

  componentDidMount() {
    this.determineIndexExists(this.props.selectedRun);
  }

  render() {
    const qTypes = [
      { label: 'Query-String', value: 'query-string' },
      { label: 'Regex', value: 'regex' },
      { label: 'Fuzzy', value: 'fuzzy' },
      { label: '*', value: 'wildcard' }
    ];

    const maxResults = [
      { label: '50', value: 50 },
      { label: '100', value: 100 },
      { label: '1000', value: 1000 }
    ];

    var body = (
      <tbody>
        <tr>
          <td width="5%">&nbsp;</td>
          <td width="5%">&nbsp;</td>
          <td width="75%">&nbsp;</td>
          <td width="15%">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          </td>
        </tr>
      </tbody>
    );

    if (this.state.searchResults.Hits) {
      body = (
        <tbody>{this.createTableBody(this.state.searchResults.Hits)}</tbody>
      );
    }

    var footer = this.createTableFooter();

    let searchBody = (
      <div>
        <br />
        <br />
        <h4>
          <span style={{ color: 'red' }}>
            {' '}
            Index for run {this.props.selectedRun} does not exist! Please re-run
            forge with text indexing enabled to utilize search feature!
          </span>
        </h4>
      </div>
    );

    if (this.state.index.exists) {
      searchBody = (
        <div>
          <Sidebar
            position="right"
            className="ui-sidebar-md"
            blockScroll={true}
            visible={this.state.sidebarVisible}
            baseZIndex={1000000}
            onHide={e => this.setState({ sidebarVisible: false })}
          >
            <FindingDetails color="#138078" finding={this.state.finding} />
          </Sidebar>
          <div className="ui-g">
            <div
              className="ui-g-12 ui-md-12"
              style={{ textAlign: 'left', padding: '0px 0px 10px 0px' }}
            >
              <div className="ui-inputgroup">
                <span
                  style={{ width: '20%', flex: 'none' }}
                  className="ui-inputgroup-addon"
                >
                  <SelectButton
                    id="query-type-btn"
                    value={this.state.qType}
                    options={qTypes}
                    tooltip="Select the type of search to conduct"
                    onChange={e => this.setState({ qType: e.value })}
                  />
                </span>
                <InputText
                  id="query-input"
                  placeholder="Please enter your search criteria..."
                  value={this.state.queryText}
                  onChange={this.handleChange}
                  onKeyDown={this.handleKeyDown}
                  style={{ width: '65%', flex: 'none' }}
                />
                <span
                  style={{ width: '15%', flex: 'none' }}
                  className="ui-inputgroup-addon"
                >
                  <SelectButton
                    id="max-results-btn"
                    value={this.state.maxResults}
                    options={maxResults}
                    tooltip="Limits your search results to 'N' findings!"
                    onChange={e => this.setState({ maxResults: e.value })}
                  />
                </span>
              </div>
            </div>
          </div>
          <div className="ui-g">
            <table className="search-table ui-datatable">
              <thead className="search-header">
                <tr>
                  <th rowSpan="2" width="5%">
                    ID
                  </th>
                  <th colSpan="2" width="80%">
                    Matched on
                  </th>
                  <th rowSpan="2" width="15%">
                    Rank
                  </th>
                </tr>
                <tr>
                  <th width="5%">Field</th>
                  <th width="95%">Value</th>
                </tr>
              </thead>
              {body}
              {footer}
            </table>
          </div>
        </div>
      );
    }

    return (
      <div>
        <Growl ref={el => (this.growl = el)} position="bottomright" />
        {searchBody}
      </div>
    );
  }
}
