/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';
import MockAxios from 'axios-mock-adapter';
import { shallow } from 'enzyme';

import RulesTableComponent from './RulesTableComponent';

const mockAxios = new MockAxios(axios, { delayResponse: Math.random() * 100 });

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<RulesTableComponent />, div);
});

it('tagFilter by valid partial string isTruthy', () => {
  const value = 'test';
  const filter = 'es';
  const wrapper = shallow(<RulesTableComponent />);
  expect(wrapper.instance().tagFilter(value, filter)).toEqual(true);
});

it('tagFilter filters by Invalid partial string isFalsey', () => {
  const value = 'test';
  const filter = 'tt';
  const wrapper = shallow(<RulesTableComponent />);
  expect(wrapper.instance().tagFilter(value, filter)).toEqual(false);
});

it('when componentDidMount then fetchDistinctRunIds called 1 time', () => {
  //given
  const wrapper = shallow(<RulesTableComponent />);
  const fetchRulesSpy = jest.spyOn(wrapper.instance(), 'fetchRules');
  //when
  wrapper.instance().componentDidMount();
  //then
  expect(fetchRulesSpy).toHaveBeenCalledTimes(1);
});
