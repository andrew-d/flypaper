import React from 'react';
import { Redirect, Route } from 'react-router';

// Require routes
import About from './pages/About';
import App from './pages/App';
import Home from './pages/Home';
import Regions from './pages/Regions';


const routes = (
  <Route component={App}>
    {/* Home page */}
    <Route path='home' component={Home} />

    {/* Regions page */}
    <Route path='regions' component={Regions} />

    {/* About page */}
    <Route path='about' component={About} />

    <Redirect from='/' to='/home' />
  </Route>
);


export default routes;
