import React from 'react';
import { connect } from 'react-redux';

import { fetchRegion } from '../actions/region';


class Region extends React.Component {
  componentWillMount() {
    this.props.dispatch(fetchRegion(this.props.routeParams.id));
  }

  render() {
    const id = this.props.routeParams.id;
    const region = this.props.regions[id];

    if( !region ) {
      return <h3>No Such Region</h3>;
    }

    return (
      <div>
        <h3>{`Region '${region.name}'`}</h3>

        <p>TODO: test start/end, hosts, edit, etc.</p>
      </div>
    );
  }
}


export default connect(state => ({
  regions: state.region,
}))(Region);
