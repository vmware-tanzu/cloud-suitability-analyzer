/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import axios from 'axios';

const BASE_URL = '/api/runs';
export default () => {
  return {
    /* /api/runs/1/data/api_detailed_usage */
    getApiDetailedUsage(runid) {
      const url = BASE_URL + '/' + runid + '/data/api_detailed_usage';
      return axios.get(url).then(response => response.data);
    },
    /* /api/runs/1/data/app_api_usage */
    getApiByAppUsage(runid) {
      const url = BASE_URL + '/' + runid + '/data/app_api_usage';
      return axios.get(url).then(response => response.data);
    },
    /* /api/runs/1/data/rule_detailed_score */
    getRuleByAppUsage(runid) {
      const url = BASE_URL + '/' + runid + '/data/app_rule_score';
      return axios.get(url).then(response => response.data);
    } /* /api/runs/1/data/api_summary */,
    getApiSummaryUsage(runid) {
      const url = BASE_URL + '/' + runid + '/data/api_summary';
      return axios.get(url).then(response => response.data);
    },
    /* /api/runs/:runid/data/sloc */
    getSourceCodeData(runid) {
      const url = BASE_URL + '/' + runid + '/data/sloc';
      return axios.get(url).then(response => response.data);
    },
    /*/api/runs/:runid/data/annotations */
    getAnnotationData(runid) {
      const url = BASE_URL + '/' + runid + '/data/annotations';
      return axios.get(url).then(response => response.data);
    },
    /* /api/runs/:runid/data/thirdParty */
    getThirdPartyData(runid) {
      const url = BASE_URL + '/' + runid + '/data/thirdParty';
      return axios.get(url).then(response => response.data);
    },
    /* /api/runs/:runid/findings */
    getRunFindings(runid) {
      const url = BASE_URL + '/' + runid + '/findings';
      return axios.get(url).then(response => response.data);
    },
    /* /api/runs/1/rule-metrics */
    getRuleMetrics(runid) {
      const url = BASE_URL + '/' + runid + '/rule-metrics';
      return axios.get(url).then(response => response.data);
    },
    getFinding(id) {
      const url = '/api/findings/' + id;
      return axios.get(url).then(response => response.data);
    },
    /* /api/runs/:runid/search */
    searchRunFindings(runid, query, queryType, maxResults) {
      const queryObj = {
        query: query,
        max: maxResults,
        type: queryType
      };

      const url = BASE_URL + '/' + runid + '/search';
      return axios.post(url, queryObj).then(response => response.data);
    },
    getRunIndexDetails(runid) {
      const url = BASE_URL + '/' + runid + '/index';
      return axios.get(url).then(response => response.data);
    }
  };
};
