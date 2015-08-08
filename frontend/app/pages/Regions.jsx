import React from 'react';
import map from 'lodash-node/modern/collection/map';
import moment from 'moment';
import {branch} from 'baobab-react/decorators';

import TagLink from '../components/TagLink';


@branch({
  regions: ['ui', '$regionsList'],
})
export default class Regions extends React.Component {
  static propTypes = {
    regions: React.PropTypes.object.isRequired,
  }

  render() {
    let regionsTable = [];

    if( this.props.regions.$error ) {
      regionsTable = [<p>Could not fetch regions: {this.props.regions.$error}</p>];
    } else {
      if( this.props.regions.$isLoading ) {
        regionsTable.push(<h4 key='ll'>Loading...</h4>);
      }

      regionsTable.push(
        <table key='rt' className='table table-striped' style={{minWidth: '100%'}}>
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Test Start (UTC)</th>
              <th>Test End (UTC)</th>
            </tr>
          </thead>

          <tbody>
            {this.renderRegions()}
          </tbody>
        </table>
      );
    }

    return (
      <div>
        <h2>Regions</h2>

        {regionsTable}
      </div>
    );
  }

  renderRegions() {
    return map(this.props.regions, (region, key) => {
      if( key[0] === '$' ) return null;

      const test_start = region.test_start ?
        moment(region.test_start).format('dddd, MMMM Do YYYY, h:mm:ss a') :
        <i>none</i>;
      const test_end = region.test_end ?
        moment(region.test_end).format('dddd, MMMM Do YYYY, h:mm:ss a') :
        <i>none</i>;

      return (
        <TagLink
          tagName='tr'
          to={`/regions/${region.id}`}
          key={`region-${region.id}`}
          >
          <td>{region.id}</td>
          <td>{region.name}</td>
          <td>{test_start}</td>
          <td>{test_end}</td>
          <td>
            <button
              className='btn btn-xs btn-default'
              onClick={this.handleEdit.bind(this, region.id)}
              >
              Edit
            </button>
            <button
              className='btn btn-xs btn-danger'
              onClick={this.handleDelete.bind(this, region.id)}
              >
              Delete
            </button>
          </td>
        </TagLink>
      );
    });
  }

  handleEdit(id, e) {
    e.preventDefault();
    e.stopPropagation();

    console.log(`TODO: Would edit region ${id}`);
  }

  handleDelete(id, e) {
    e.preventDefault();
    e.stopPropagation();

    console.log(`TODO: Would delete region ${id}`);
  }
}
