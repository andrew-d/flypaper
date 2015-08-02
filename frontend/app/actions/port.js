import * as constants from '../constants';
import PortAPI from '../api/port';


export function fetchPort(num) {
  return dispatch => {
    PortAPI.fetchPort(num)
           .then(res => dispatch({
             type: constants.FETCH_PORT,
             port: res,
           });
  };
}
