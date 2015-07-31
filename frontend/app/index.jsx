// This must come first, in this order
import 'ie8';
import 'html5shiv/dist/html5shiv';
import 'html5shiv/dist/html5shiv-printshiv';
import 'babel/polyfill';

import React from 'react';
import Router from 'react-router';

// Import vendor styles here.
import 'bootstrap/dist/css/bootstrap.css';

import routes from './routes';
Router.run(routes, Router.HistoryLocation, function(Handler) {
  React.render(<Handler />, document.body);
});
