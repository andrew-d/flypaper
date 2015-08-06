import Baobab from 'baobab';

import RegionsList from './facets/regionsList';


const tree = new Baobab({
  // The top-level `data` key contains all models.
  data: {
    ports: {},
    regions: {},
  },

  // The top-level `ui` key contains presentation logic - i.e. displaying the
  // models from `data` in a particular way.
  ui: {
    $regionsList: RegionsList,
  },
});

export default tree;
