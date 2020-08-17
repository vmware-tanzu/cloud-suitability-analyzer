/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';
import ReactDOM from 'react-dom';
import ApplicationViewComponent from './ApplicationViewComponent';
import axios from 'axios';
import MockAxios from 'axios-mock-adapter';
import { shallow } from 'enzyme';

const mockAxios = new MockAxios(axios, { delayResponse: Math.random() * 100 });

it('renders without crashing', () => {
  const div = document.createElement('div');
  const run = { id: '1' };
  ReactDOM.render(<ApplicationViewComponent selectedRun={run} />, div);
  ReactDOM.unmountComponentAtNode(div);
});

it('when populateAppStateData state set and populateSummaryAndApplicationScores called', () => {
  //given
  const run = { id: '2' };
  const wrapper = shallow(<ApplicationViewComponent selectedRun={run} />);
  const fetchApplicationScoreCardSpy = jest.spyOn(
    wrapper.instance(),
    'fetchApplicationScoreCard'
  );
  const fetchApplicationScoreCardDetailsSpy = jest.spyOn(
    wrapper.instance(),
    'fetchApplicationScoreCardDetails'
  );
  const fetchLanguagesByAppSpy = jest.spyOn(
    wrapper.instance(),
    'fetchLanguagesByApp'
  );
  const fetchApisByAppSpy = jest.spyOn(wrapper.instance(), 'fetchApisByApp');
  //when
  wrapper.instance().populateAppStateData(2, 'app1', 'info', [], [], false);
  //then
  expect(fetchApplicationScoreCardSpy).toHaveBeenCalledWith(
    2,
    'app1',
    [],
    [],
    false
  );
  expect(fetchApplicationScoreCardDetailsSpy).toHaveBeenCalledWith(
    2,
    'app1',
    'info',
    [],
    false
  );
  expect(fetchLanguagesByAppSpy).toHaveBeenCalledWith(2, 'app1');
  expect(fetchApisByAppSpy).toHaveBeenCalledWith(2, 'app1');
});
