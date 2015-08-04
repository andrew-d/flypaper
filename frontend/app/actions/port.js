import * as constants from '../constants';
import PortAPI from '../api/port';


export function fetchPort(id) {
  return dispatch => {
    PortAPI.fetchPort(id)
           .then(res => dispatch({
             type: constants.FETCH_PORT,
             port: res,
           });
  };
}
