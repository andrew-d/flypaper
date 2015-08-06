import React from 'react';
import { Router } from 'react-router';
import {root} from 'baobab-react/decorators';

import routes from './routes';
import tree from './tree.js';


@root(tree)
export default class Root extends React.Component {
  static propTypes = {
    history: React.PropTypes.object.isRequired,
  }

  render() {
    return (
      <Router history={this.props.history} key='router'>
        {routes}
      </Router>
    );
  }
}
