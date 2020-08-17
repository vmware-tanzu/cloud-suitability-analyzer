/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios/index';
import MockAxios from 'axios-mock-adapter';

import ExecutiveSummaryViewComponent from './ExecutiveSummaryViewComponent';
import { shallow } from 'enzyme';

const mockAxios = new MockAxios(axios, { delayResponse: Math.random() * 100 });

it('renders without crashing', () => {
  const div = document.createElement('div');
  const run = { id: '1' };
  ReactDOM.render(<ExecutiveSummaryViewComponent selectedRun={run} />, div);
});

it('when summaryStateSetter then numAppsByRun,locByRun,numFilesByRun state set', () => {
  //given
  const run = { id: '1' };
  const wrapper = shallow(<ExecutiveSummaryViewComponent selectedRun={run} />);
  expect(wrapper.state('numAppsByRun')).toEqual(0);
  expect(wrapper.state('locByRun')).toEqual(0);
  expect(wrapper.state('numFilesByRun')).toEqual(0);

  //when
  wrapper.instance().summaryStateSetter({
    applicationCount: 50,
    codeLines: 100000,
    totalFiles: 100
  });

  //then
  expect(wrapper.state('numAppsByRun')).toEqual(50);
  expect(wrapper.state('locByRun')).toEqual(100000);
  expect(wrapper.state('numFilesByRun')).toEqual(100);
});

it('when componentDidMount then populateSummaryAndApplicationScores called 1 time', () => {
  //given
  const run = { id: '1' };
  const wrapper = shallow(<ExecutiveSummaryViewComponent selectedRun={run} />);
  const populateSummaryAndApplicationScoresSpy = jest.spyOn(
    wrapper.instance(),
    'populateSummaryAndApplicationScores'
  );
  //when
  wrapper.instance().componentDidMount();
  //then
  expect(populateSummaryAndApplicationScoresSpy).toHaveBeenCalledTimes(1);
});

it('when rendered then prime react components present', () => {
  const run = { id: '1' };
  const wrapper = shallow(<ExecutiveSummaryViewComponent selectedRun={run} />);
  expect(wrapper.find('Growl').exists()).toBeTruthy();
  expect(wrapper.find('DataTable').exists()).toBeTruthy();
});
