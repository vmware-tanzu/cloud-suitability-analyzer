/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import axios from 'axios';

const BASE_URL = '/api/runs';
export default () => {
  return {
    getApplicationScoresByRun(runid) {
      let scores = { normalizedScore: 0, appScores: [] };

      if (runid === '' || runid === null) {
        return Promise.resolve({
          scores
        });
      }
      const url =
        BASE_URL +
        '/' +
        runid +
        '/summary/application_scores?rangeMin=0&rangeMax=10&reverse=false';

      return axios.get(url).then(response => {
        if (response.data.scores && response.data.scores.appScores) {
          response.data.scores.appScores.map(app => {
            return response.data.scores;
          });
          return response.data.scores;
        }
        return scores;
      });
    },
    getSummaryFindingsByRun(runid) {
      if (runid === '' || runid === null) {
        return Promise.resolve({
          applicationCount: 0,
          codeLines: 0,
          totalFiles: 0,
          findings: 0
        });
      }
      const url = BASE_URL + '/' + runid + '/summary/run_slocs';
      return axios.get(url).then(response => {
        return response.data;
      });
    },
    updateApp(app) {
      const url = BASE_URL + '/' + app.runId + '/apps/' + app.appId + '/';
      return axios.post(url, app).then(response => response.data);
    }
  };
};
