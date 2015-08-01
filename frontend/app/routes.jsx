import React from 'react';
import { Redirect, Route } from 'react-router';

// Require routes
import About from './pages/About';
import App from './pages/App';
import Home from './pages/Home';


const routes = (
  <Route component={App}>
    {/* Introduction page */}
    <Route path='/' component={Home} />

    {/* About page */}
    <Route path='/about' component={About} />
  </Route>
);


export default routes;
