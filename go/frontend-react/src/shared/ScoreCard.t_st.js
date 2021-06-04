/*
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import React from 'react';
import ReactDOM from 'react-dom';
import ScoreCard from './ScoreCard';

it('renders without crashing', () => {
  const div = document.createElement('div');
  const applicationScoreCardMock = {
    application: 'src',
    info: '2',
    low: '8',
    medium: '1',
    high: '0',
    total: '11',
    app: {
      score: 7.6,
      rawScore: 4567,
      tags: ['tommy', 'mickey']
    }
  };

  ReactDOM.render(
    <ScoreCard applicationScoreCard={applicationScoreCardMock} />,
    div
  );
  ReactDOM.unmountComponentAtNode(div);
});
