import React from 'react';
import { connect } from 'react-redux';
import map from 'lodash-node/modern/collection/map';
import moment from 'moment';

import { fetchRegions } from '../actions/region';
import TagLink from '../components/TagLink';


class Regions extends React.Component {
  componentDidMount() {
    let { dispatch } = this.props;
    dispatch(fetchRegions());
  }

  render() {
    return (
      <div>
        <h2>Regions</h2>

        <table className='table table-striped' style={{minWidth: '100%'}}>
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
      </div>
    );
  }

  renderRegions() {
    return map(this.props.regions, (region) => {
      const test_start = region.test_start ?
        moment(region.test_start).format("dddd, MMMM Do YYYY, h:mm:ss a") :
        '';
      const test_end = region.test_end ?
        moment(region.test_end).format("dddd, MMMM Do YYYY, h:mm:ss a") :
        '';

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


export default connect(state => ({
  regions: state.region.regions,
}))(Regions);
