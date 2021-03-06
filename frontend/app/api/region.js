/*global fetch */
import { checkStatus } from './util';

export default {
  fetchRegion(id) {
    return fetch(`/api/regions/${id}`)
           .then(checkStatus)
           .then(res => res.json());
  },

  fetchRegions() {
    return fetch(`/api/regions`)
           .then(checkStatus)
           .then(res => res.json());
  },
}
