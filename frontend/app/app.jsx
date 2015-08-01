import React from 'react';

import HashHistory from 'react-router/lib/HashHistory';
import { Router } from 'react-router';

import { applyMiddleware, createStore, combineReducers, compose } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import { loggerMiddleware } from './middleware';
import { promiseMiddleware } from './utils/redux';


// Create Redux
let middleware = [thunk, promiseMiddleware];

// In production, we want to use just the middleware.
// In development, we want to use some store enhancers from redux-devtools.
// UglifyJS will eliminate the dead code depending on the build environment.
let finalCreateStore;

if (process.env.NODE_ENV === 'production') {
  finalCreateStore = applyMiddleware(middleware)(createStore);
} else {
  finalCreateStore = compose(
    applyMiddleware(middleware),
    require('redux-devtools').devTools(),
    require('redux-devtools').persistState(
      window.location.href.match(/[?&]debug_session=([^&]+)\b/)
    ),
    createStore
  );
}

// Build reducers.
const composedReducers = combineReducers([] /* reducers */);

// Build final store.
const state = {};
const store = finalCreateStore(composedReducers, state);


// The main application class.
export default class App extends React.Component {
  render() {
    const history = HashHistory;

    return (
      <Provider store={store}>
        {renderRoutes.bind(null, history)}
      </Provider>
    );
  }
}


import routes from './routes';
function renderRoutes(history) {
  return (
    <Router history={history}>
      {routes}
    </Router>
  );
}
