/*
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 */

import './App.css';
import React, { Component } from 'react';
import RunContainer from './run-container/RunContainer';
import { Footer } from './view-template/Footer';

class App extends Component {
  render() {
    return (
      <div style={{ backgroundColor: '#edf0f5' }}>
        <div>
          <RunContainer />
        </div>
        <div>
          <Footer />
        </div>
      </div>
    );
  }
}

export default App;
