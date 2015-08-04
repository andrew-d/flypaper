/*global fetch */
import { checkStatus } from './util';

export default {
  fetchPort(id) {
    return fetch(`/api/ports/${port}`)
           .then(checkStatus)
           .then(res => res.json());
  },

  fetchPorts() {
    return fetch(`/api/ports`)
           .then(checkStatus)
           .then(res => res.json());
  },

  fetchPortsForHost(host) {
    return fetch(`/api/hosts/${host}/ports`)
           .then(checkStatus)
           .then(res => res.json());
  },
}
