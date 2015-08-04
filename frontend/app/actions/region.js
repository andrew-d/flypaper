import * as constants from '../constants';
import RegionAPI from '../api/region';


export function fetchRegions(id) {
  return dispatch => {
    RegionAPI.fetchRegions(id)
             .then(res => dispatch({
               type: constants.FETCH_REGIONS,
               regions: res,
             }));
  };
}

export function fetchRegion(id) {
  return dispatch => {
    RegionAPI.fetchRegion(id)
             .then(res => dispatch({
               type: constants.FETCH_REGION,
               region: res,
             }));
  };
}
