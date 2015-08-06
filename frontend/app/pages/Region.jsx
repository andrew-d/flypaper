import React from 'react';
import {branch} from 'baobab-react/decorators';

@branch((props) => ({
  region: ['data', 'regions', props.routeParams.id],
}))
export default class Region extends React.Component {
  static propTypes = {
    region: React.PropTypes.object.isRequired,
  }

  render() {
    if( !this.props.region ) {
      return <h3>No Such Region</h3>;
    }

    return (
      <div>
        <h3>{`Region '${this.props.region.name}'`}</h3>

        <p>TODO: test start/end, hosts, edit, etc.</p>
      </div>
    );
  }
}
