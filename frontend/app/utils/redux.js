/**
 * Note: taken from: https://gist.github.com/iNikNik/3c1b870f63dc0de67c38
 */

import map from 'lodash-node/modern/collection/map';
import mapValues from 'lodash-node/modern/object/mapValues';
import zipObject from 'lodash-node/modern/array/zipObject';
import uniqueId from 'uniqueid';


// Create actions that don't need constants
export const createActions = (actionObj) => {
  const baseId = uniqueId();

  return zipObject(map(actionObj, (actionCreator, key) => {
    const actionId = `${baseId}-${key}`;
    const asyncTypes = ['BEGIN', 'SUCCESS', 'FAILURE'].map( (state) => `${actionId}-${state}`);
    const method = (...args) => {
      const result = actionCreator(...args);

      if (result instanceof Promise) {
        // Promise (async)
        return {
          types: asyncTypes,
          promise: result,
        };

      } else if (typeof result === 'function') {
        // Function (async)
        return (...args) => { // eslint-disable-line no-shadow
          return {
            type: actionId,
            ...(result(...args) || {}),
          };
        };

      } else { // eslint-disable-line no-else-return
        // Object (synchronous)
        return {
          type: actionId,
          ...(result || {}),
        };
      }
    };

    if (actionCreator._async === true) {
      const [ begin, success, failure ] = asyncTypes;
      method._id = {
        begin,
        success,
        failure,
      };
    } else {
      method._id = actionId;
    }

    return [key, method];
  }));
};


// Get action ids from actions created with `createActions`
export const getActionIds = (actionCreators) => {
  return mapValues(actionCreators, (value, key) => { // eslint-disable-line no-unused-vars
    return value._id;
  });
};


// Replace switch statements in stores (taken from the Redux README)
export const createStore = (initialState, handlers) => {
  return (state = initialState, action = {}) =>
    handlers[action.type] ?
      handlers[action.type](state, action) :
      state;
};


export function promiseMiddleware() {
  return (next) => (action) => {
    const { promise, types, ...rest } = action;
    if (!promise) {
      return next(action);
    }

    const [REQUEST, SUCCESS, FAILURE] = types;
    next({ ...rest, type: REQUEST });

    return promise.then(
      (result) => next({ ...rest, result, type: SUCCESS }),
      (error) => next({ ...rest, error, type: FAILURE })
    );
  };
}


export function asyncAction() {
  return (target, name, descriptor) => {
    descriptor.value._async = true;
    return descriptor;
  };
}
