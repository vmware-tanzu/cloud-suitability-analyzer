/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';
import MockAxios from 'axios-mock-adapter';
import { shallow } from 'enzyme/build/index';

import PortfolioViewComponent from './PortfolioViewComponent';

const mockAxios = new MockAxios(axios, { delayResponse: Math.random() * 100 });

it('renders without crashing', () => {
  const div = document.createElement('div');
  const run = { id: '1' };
  ReactDOM.render(<PortfolioViewComponent selectedRun={run} />, div);
  ReactDOM.unmountComponentAtNode(div);
});

it('when componentDidMount then populateLanguagesAndApis called 1 time', () => {
  //given
  const run = { id: '1' };
  const wrapper = shallow(<PortfolioViewComponent selectedRun={run} />);
  const populateLanguagesAndApisSpy = jest.spyOn(
    wrapper.instance(),
    'populateLanguagesAndApis'
  );
  //when
  wrapper.instance().componentDidMount();
  //then
  expect(populateLanguagesAndApisSpy).toHaveBeenCalledTimes(1);
});

it('when onSelectedApiChange then selectedApi state set and fetchAppsByApi called', () => {
  //given
  const run = { id: '1' };
  const wrapper = shallow(<PortfolioViewComponent selectedRun={run} />);
  const fetchAppsByApiSpy = jest.spyOn(wrapper.instance(), 'fetchAppsByApi');
  expect(wrapper.state('selectedApi')).toEqual('');
  //when
  wrapper.instance().onSelectedApiChange({ value: 'JspAPI' });
  //then
  expect(wrapper.state('selectedApi')).toEqual('JspAPI');
  expect(fetchAppsByApiSpy).toHaveBeenCalledWith({ id: '1' }, 'JspAPI');
});

it('when onSelectedLanguageChange then selectedLanguage state set and fetchAppsByLanguage called', () => {
  //given
  const run = { id: '1' };
  const wrapper = shallow(<PortfolioViewComponent selectedRun={run} />);
  const fetchAppsByLanguageSpy = jest.spyOn(
    wrapper.instance(),
    'fetchAppsByLanguage'
  );
  expect(wrapper.state('selectedLanguage')).toEqual('');
  //when
  wrapper.instance().onSelectedLanguageChange({ value: 'Java' });
  //then
  expect(wrapper.state('selectedLanguage')).toEqual('Java');
  expect(fetchAppsByLanguageSpy).toHaveBeenCalledWith({ id: '1' }, 'Java');
});

it('when populateLanguagesAndApis then fetchTop5LanguagesByLoc,fetchTop5ApisByScore,fetchAllLanguages,fetchAllApis called', () => {
  //given
  const run = { id: '1' };
  const wrapper = shallow(<PortfolioViewComponent selectedRun={run} />);
  const fetchTop5LanguagesByLocSpy = jest.spyOn(
    wrapper.instance(),
    'fetchTop5LanguagesByLoc'
  );
  const fetchTop5ApisByScoreSpy = jest.spyOn(
    wrapper.instance(),
    'fetchTop5ApisByScore'
  );
  const fetchAllLanguagesSpy = jest.spyOn(
    wrapper.instance(),
    'fetchAllLanguages'
  );
  const fetchAllApisSpy = jest.spyOn(wrapper.instance(), 'fetchAllApis');
  //when
  wrapper.instance().populateLanguagesAndApis(1);
  //then
  expect(fetchTop5LanguagesByLocSpy).toHaveBeenCalledWith(1);
  expect(fetchTop5ApisByScoreSpy).toHaveBeenCalledWith(1);
  expect(fetchAllLanguagesSpy).toHaveBeenCalledWith(1);
  expect(fetchAllApisSpy).toHaveBeenCalledWith(1);
});

it('when rendered then prime react components present', () => {
  const run = { id: '1' };
  const wrapper = shallow(<PortfolioViewComponent selectedRun={run} />);
  expect(wrapper.find('Growl').exists()).toBeTruthy();
  expect(wrapper.find('Chart').exists()).toBeTruthy();
  expect(wrapper.find('Dropdown').exists()).toBeTruthy();
});
