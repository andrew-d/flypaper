import regionsLoader from '../loaders/regions';


let regionsList = {
  cursors: {
    regions: ['data', 'regions'],
  },
  get(state) {
    // If we have any regions, then we don't need to fetch from the server.
    if( Object.keys(state.regions).length ) {
      return state.regions;
    }

    // Use the project loader to get the set of regions (and also load more
    // from the server).
    return regionsLoader(this);
  },
};


export default regionsList;
