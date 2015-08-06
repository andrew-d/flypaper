/*global fetch */
import { checkStatus } from './util';


export default function(tree) {
  // Kick off a request to load regions from the server.
  fetch(`/api/regions`)
    .then(checkStatus)
    .then(res => res.json())
    .then(function (regions) {
      // Update all regions in the list.
      regions.forEach(function (region) {
        tree.set(['data', 'regions', region.id], region);
      });
    })
    .catch(function(ex) {
      tree.set(['data', 'regions', '$error'], ex);
    });

  // Fetch the set of regions.
  const regions = tree.get(['data', 'regions']);

  // Create a clone of the regions that we can set as 'is loading'.
  let loadingRegions = Object.create(regions);
  loadingRegions.$isLoading = true;

  return loadingRegions;
}
