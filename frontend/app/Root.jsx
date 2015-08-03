import React from 'react';

import { Router } from 'react-router';

import { applyMiddleware, createStore, combineReducers, compose } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import { loggerMiddleware } from './middleware';
import { promiseMiddleware } from './utils/redux';

import * as stores from './stores';


// Create Redux
let middleware = [thunk, promiseMiddleware];

// In production, we want to use just the middleware.
// In development, we want to use some store enhancers from redux-devtools.
// UglifyJS will eliminate the dead code depending on the build environment.
let createStoreWithMiddleware;

if (process.env.NODE_ENV === 'production') {
  createStoreWithMiddleware = applyMiddleware(middleware)(createStore);
} else {
  createStoreWithMiddleware = compose(
    applyMiddleware(middleware),
    require('redux-devtools').devTools(),
    require('redux-devtools').persistState(
      window.location.href.match(/[?&]debug_session=([^&]+)\b/)
    ),
    createStore
  );
}

// Build reducers.
const reducer = combineReducers(stores);

// Build final store.
const store = createStoreWithMiddleware(reducer);


// The main application class.
export default class Root extends React.Component {
  static propTypes = {
    history: React.PropTypes.object.isRequired,
  }

  render() {
    return (
      <Provider store={store}>
        {renderRoutes.bind(null, this.props.history)}
      </Provider>
    );
  }
}


import routes from './routes';
function renderRoutes(history) {
  let children = [
    <Router history={history} key="router">
      {routes}
    </Router>
  ];

  if (process.env.NODE_ENV !== 'production') {
    const { DevTools, DebugPanel, LogMonitor } = require('redux-devtools/lib/react');

    children.push(
      <DebugPanel top right bottom key="debugPanel" style={{zIndex: 9999}}>
        <DevTools store={store} monitor={LogMonitor}/>
      </DebugPanel>
    );
  }

  return (
    <div>
      {children}
    </div>
  );
}
