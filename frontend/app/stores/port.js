import * as constants from '../constants';


const initialState = {
  ports: {},
};


const actionsMap = {
  [constants.FETCH_PORT]: (state, action) => ({
	...state,
	[action.port.id]: action.port,
  }),
};


export default function portStore(state = initialState, action) {
  const reduceFunc = actionsMap[action.type];
  if( !reduceFunc ) return state;

  return Object.assign({}, state, reduceFunc(state, action));
}
