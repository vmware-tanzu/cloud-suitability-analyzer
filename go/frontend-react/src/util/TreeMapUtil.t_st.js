/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import TreeMapUtil from './TreeMapUtil';

it('generateApisTreeMap is Valid', () => {
  const resp = [
    { api: 'JsfAPI', usageCount: 1000 },
    { api: 'JspAPI', usageCount: 2000 }
  ];

  expect(TreeMapUtil().generateApisTreeMap(resp)).toEqual([
    { name: 'JspAPI', size: 2000 },
    { name: 'JsfAPI', size: 1000 }
  ]);
});

it('generateApisTreeMap receives null response then empty array returned', () => {
  const resp = null;
  expect(TreeMapUtil().generateApisTreeMap(resp)).toEqual([]);
});

it('generateApisTreeMap with allothers category when number of apis > 5', () => {
  const resp = [
    { api: 'JsfAPI', usageCount: 1000 },
    { api: 'Servlet', usageCount: 2000 },
    { api: 'config', usageCount: 3000 },
    { api: 'websphere', usageCount: 4000 },
    { api: 'JspAPI', usageCount: 5000 },
    { api: 'JPA', usageCount: 1 },
    { api: 'StrutsApi', usageCount: 2 },
    { api: 'ejbAPI', usageCount: 3 }
  ];

  const expectedResult = [
    { name: 'JspAPI', size: 5000 },
    { name: 'websphere', size: 4000 },
    { name: 'config', size: 3000 },
    { name: 'Servlet', size: 2000 },
    { name: 'JsfAPI', size: 1000 },
    { name: 'Other', size: 6, value: 'ejbAPI,StrutsApi,JPA' }
  ];

  expect(TreeMapUtil().generateApisTreeMap(resp)).toEqual(expectedResult);
});

it('generateLanguageTreeMap is Valid and returns top5TempLanguageTreeMap when size less than 5', () => {
  const resp = [
    { language: 'Java', codeLines: 1000 },
    { language: 'XML', codeLines: 2000 }
  ];

  expect(TreeMapUtil().generateLanguageTreeMap(resp)).toEqual([
    { name: 'XML', size: 2000 },
    { name: 'Java', size: 1000 }
  ]);
});

it('generateLanguageTreeMap receives null response then object with top 5 and all other empty arrays returned', () => {
  const resp = null;
  expect(TreeMapUtil().generateLanguageTreeMap(resp)).toEqual([]);
});

it('generateLanguageTreeMap with allothers category when number of languages > 5', () => {
  const resp = [
    { language: 'Java', codeLines: 1000 },
    { language: '.net', codeLines: 2000 },
    { language: 'javascript', codeLines: 3000 },
    { language: 'R', codeLines: 4000 },
    { language: 'Go', codeLines: 5000 },
    { language: 'Html', codeLines: 1 },
    { language: 'bourne shell', codeLines: 2 },
    { language: 'XML', codeLines: 3 }
  ];

  const expectedResult = [
    { name: 'Go', size: 5000 },
    { name: 'R', size: 4000 },
    { name: 'javascript', size: 3000 },
    { name: '.net', size: 2000 },
    { name: 'Java', size: 1000 },
    { name: 'Other', size: 6, value: 'XML,bourne shell,Html' }
  ];

  expect(TreeMapUtil().generateLanguageTreeMap(resp)).toEqual(expectedResult);
});
