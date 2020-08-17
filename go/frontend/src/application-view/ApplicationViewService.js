/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import axios from "axios";

const BASE_URL = "/api/runs";
export default () => {
  return {
    getApplicationScoreCard(runid, app, tags, includeFF) {
      const url =
        BASE_URL +
        "/" +
        runid +
        "/apps/" +
        app +
        "/scorecard?includeFF=" +
        includeFF;

      const requestObj = {
        tags: tags
      };

      return axios.post(url, requestObj).then(response => {
        if (response.data) {
          response.data.info !== undefined
            ? (response.data.info = response.data.info.toString())
            : (response.data.info = "n/a");

          response.data.low !== undefined
            ? (response.data.low = response.data.low.toString())
            : (response.data.low = "n/a");

          response.data.medium !== undefined
            ? (response.data.medium = response.data.medium.toString())
            : (response.data.medium = "n/a");

          response.data.high !== undefined
            ? (response.data.high = response.data.high.toString())
            : (response.data.high = "n/a");

          response.data.total !== undefined
            ? (response.data.total = (
                response.data.total - response.data.info
              ).toString())
            : (response.data.total = "n/a");
        }
        return response.data;
      });
    },
    getApplicationByRun(runid) {
      const url =
        BASE_URL +
        "/" +
        runid +
        "/summary/application_scores?rangeMin=1&rangeMax=10&reverse=false";
      if (runid === "" || runid === null) {
        return Promise.resolve({
          scores: {}
        });
      }
      return axios.get(url).then(response => {
        return response.data.scores;
      });
    },
    getApplicationScoreCardDetailsByLevel(runid, app, card, tags, includeFF) {
      const requestObj = {
        tags: tags
      };

      const url =
        BASE_URL +
        "/" +
        runid +
        "/apps/" +
        app +
        "/findings/scorecard/" +
        card +
        "?includeFF=" +
        includeFF;
      return axios.post(url, requestObj).then(response => response.data);
    },
    getLanguagesByApplicationRun(runid, app) {
      const url = BASE_URL + "/" + runid + "/apps/" + app + "/languages";
      return axios.get(url).then(response => response.data);
    },
    getApisByApplicationRun(runid, app) {
      const url = BASE_URL + "/" + runid + "/apps/" + app + "/apis";
      return axios.get(url).then(response => response.data);
    },
    getTagsForApplication(runid, app) {
      const url = BASE_URL + "/" + runid + "/apps/" + app + "/tags";
      return axios.get(url).then(response => response.data);
    }
  };
};
