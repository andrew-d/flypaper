import * as constants from '../constants';


const initialState = {
  regions: {},
};


const actionsMap = {
  [constants.FETCH_REGION]: (state, action) => ({
    ...state,
    [action.region.id]: action.region,
  }),

  [constants.FETCH_REGIONS]: (state, action) => {
    let newState = state;

    for( let region of action.regions ) {
      newState.regions[region.id] = region;
    }

    return newState;
  },
};


export default function regionStore(state = initialState, action) {
  const reduceFunc = actionsMap[action.type];
  if( !reduceFunc ) return state;

  return Object.assign({}, state, reduceFunc(state, action));
}
