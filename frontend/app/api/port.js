/*global fetch */

export default {
  fetchPort(id) {
    return fetch(`/api/ports/${port}`)
            .then(res => res.json());
  },

  fetchPortForHost(host) {
    return fetch(`/api/hosts/${host}/ports`)
            .then(res => res.json());
  },
}
