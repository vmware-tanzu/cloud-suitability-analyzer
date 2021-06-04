/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';
import DropDownUtil from './DropDownUtil';

it('generateRunItemsDropDown is Valid and sorted in reverse order', () => {
  const alias1 = 'target1';
  const alias2 = 'target2';
  const run1 = { id: '1', Alias: alias1 };
  const run2 = { id: '2', Alias: alias2 };

  const resp = [run1, run2];

  expect(DropDownUtil().generateRunItemsDropDown(resp)).toEqual([
    { label: '2 - ' + alias2, value: run2 },
    { label: '1 - ' + alias1, value: run1 }
  ]);
});

it('generateRunItemsDropDown receives null response then empty array returned', () => {
  const resp = null;
  expect(DropDownUtil().generateRunItemsDropDown(resp)).toEqual([]);
});

it('setSelectedRun receives null response then empty array returned', () => {
  const resp = [{ value: 1 }, { value: 2 }];
  expect(DropDownUtil().setSelectedRun(resp)).toEqual(1);
});

it('setSelectedRun receives null response then empty string returned', () => {
  const resp = null;
  expect(DropDownUtil().setSelectedRun(resp)).toEqual('');
});

it('generateApplicationItemsDropDown is Valid', () => {
  const resp = [{ name: 'app-1' }, { name: 'app-2' }];

  expect(DropDownUtil().generateApplicationItemsDropDown(resp)).toEqual([
    { label: 'app-1', value: 'app-1' },
    { label: 'app-2', value: 'app-2' }
  ]);
});

it('generateApplicationItemsDropDown receives null response then empty array returned', () => {
  const resp = null;
  expect(DropDownUtil().generateApplicationItemsDropDown(resp)).toEqual([]);
});

it('setSelectedApplication valid response then return value valid', () => {
  const resp = [{ value: '1' }, { value: '2' }];
  expect(DropDownUtil().setSelectedApplication(resp)).toEqual('1');
});

it('setSelectedApplication receives null response then empty string returned', () => {
  const resp = null;
  expect(DropDownUtil().setSelectedApplication(resp)).toEqual('');
});

it('generateLanguageItemsDropDown is Valid', () => {
  const resp = [{ language: 'java' }, { language: 'cobol' }];

  expect(DropDownUtil().generateLanguageItemsDropDown(resp)).toEqual([
    { label: 'java', value: 'java' },
    { label: 'cobol', value: 'cobol' }
  ]);
});

it('generateLanguageItemsDropDown receives null response then empty array returned', () => {
  const resp = null;
  expect(DropDownUtil().generateLanguageItemsDropDown(resp)).toEqual([]);
});

it('setSelectedLanguage valid response then return value valid', () => {
  const resp = [{ value: 1 }, { value: 2 }];
  expect(DropDownUtil().setSelectedLanguage(resp)).toEqual(1);
});

it('setSelectedLanguage receives null response then empty string returned', () => {
  const resp = null;
  expect(DropDownUtil().setSelectedLanguage(resp)).toEqual('');
});

it('generateApiItemsDropDown is Valid', () => {
  const resp = [{ api: 'java' }, { api: 'cobol' }];

  expect(DropDownUtil().generateApiItemsDropDown(resp)).toEqual([
    { label: 'java', value: 'java' },
    { label: 'cobol', value: 'cobol' }
  ]);
});

it('generateApiItemsDropDown receives null response then empty array returned', () => {
  const resp = null;
  expect(DropDownUtil().generateApiItemsDropDown(resp)).toEqual([]);
});

it('setSelectedApi valid response then return value valid', () => {
  const resp = [{ value: 1 }, { value: 2 }];
  expect(DropDownUtil().setSelectedApi(resp)).toEqual(1);
});

it('setSelectedApi receives null response then empty string returned', () => {
  const resp = null;
  expect(DropDownUtil().setSelectedApi(resp)).toEqual('');
});
