// This must come first, in this order
import 'ie8';
import 'html5shiv/dist/html5shiv';
import 'html5shiv/dist/html5shiv-printshiv';
import 'babel/polyfill';

import React from 'react';
import ReactDOM from 'react-dom';

// Import vendor styles here.
import 'bootstrap/dist/css/bootstrap.css';
import 'font-awesome/css/font-awesome.css';

// Render the app.
import App from './app';
ReactDOM.render(<App />, document.getElementById('app'));
