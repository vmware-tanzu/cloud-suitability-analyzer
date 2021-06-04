/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';

import RulesService from './RulesService';
const mockRulesDataJson = require('../../public/mock/rules');

const mockRulesData = jest.fn();
mockRulesData.mockReturnValue(mockRulesDataJson);

jest.mock('./RulesService', () => {
  return jest.fn().mockImplementation(() => {
    return { getRules: mockRulesData };
  });
});

beforeEach(() => {
  RulesService.mockClear();
  mockRulesData.mockClear();
});

it('ruleservice created and called on creation', () => {
  const rulesService = new RulesService();
  expect(rulesService).toBeTruthy();
  expect(RulesService).toHaveBeenCalledTimes(1);
});

it('We can check if the consumer called a method on the class instance', () => {
  const rulesService = new RulesService();
  const mockRules = mockRulesDataJson;
  rulesService.getRules();
  expect(mockRulesData()).toEqual(mockRules);
});
