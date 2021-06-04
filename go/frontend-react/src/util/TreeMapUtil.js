/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

export default () => {
  return {
    generateApisTreeMap(resp) {
      let tempApisTreeMap = [];
      let top5TempApiTreeMap = [];
      let allOthersTempApiTreeMap = [];

      if (resp != null) {
        resp.forEach(app => {
          tempApisTreeMap.push({
            name: app.api,
            size: app.usageCount
          });
        });

        tempApisTreeMap.sort(function(a, z) {
          return a.size - z.size;
        });

        tempApisTreeMap.reverse();

        top5TempApiTreeMap = tempApisTreeMap.slice(0, 5);

        if (tempApisTreeMap.length > 5) {
          allOthersTempApiTreeMap = tempApisTreeMap.slice(
            5,
            tempApisTreeMap.length
          );
          let otherApiNameAggr = '';
          let otherSizeAggr = 0;
          for (let [index, val] of allOthersTempApiTreeMap.entries()) {
            if (index === 0) {
              otherApiNameAggr = val.name;
              otherSizeAggr = val.size;
            } else {
              otherApiNameAggr = otherApiNameAggr + ',' + val.name;
              otherSizeAggr = otherSizeAggr + val.size;
            }
          }

          top5TempApiTreeMap.push({
            name: 'Other',
            value: otherApiNameAggr,
            size: otherSizeAggr
          });
        }
      }
      return top5TempApiTreeMap;
    },
    generateLanguageTreeMap(resp) {
      let tempLanguageTreeMap = [];
      let top5TempLanguageTreeMap = [];
      let allOthersTempLanguageTreeMap = [];

      if (resp != null) {
        resp.forEach(app => {
          tempLanguageTreeMap.push({
            name: app.language,
            size: app.codeLines
          });
        });

        tempLanguageTreeMap.sort(function(a, z) {
          return a.size - z.size;
        });

        tempLanguageTreeMap.reverse();

        top5TempLanguageTreeMap = tempLanguageTreeMap.slice(0, 5);

        if (tempLanguageTreeMap.length > 5) {
          allOthersTempLanguageTreeMap = tempLanguageTreeMap.slice(
            5,
            tempLanguageTreeMap.length
          );
          let otherLangNameAggr = '';
          let otherSizeAggr = 0;
          for (let [index, val] of allOthersTempLanguageTreeMap.entries()) {
            if (index === 0) {
              otherLangNameAggr = val.name;
              otherSizeAggr = val.size;
            } else {
              otherLangNameAggr = otherLangNameAggr + ',' + val.name;
              otherSizeAggr = otherSizeAggr + val.size;
            }
          }

          top5TempLanguageTreeMap.push({
            name: 'Other',
            value: otherLangNameAggr,
            size: otherSizeAggr
          });
        }
      }
      return top5TempLanguageTreeMap;
    }
  };
};
