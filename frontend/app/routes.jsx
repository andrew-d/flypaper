import React from 'react';
import { Redirect, Route } from 'react-router';

// Require routes
import About from './pages/About';
import App from './pages/App';
import Home from './pages/Home';
import Region from './pages/Region';
import Regions from './pages/Regions';


const routes = (
  <Route component={App}>
    {/* Home page */}
    <Route path='home' component={Home} />

    {/* Regions */}
    <Route path='regions' component={Regions} />
    <Route path='regions/:id' component={Region} />

    {/* About page */}
    <Route path='about' component={About} />

    <Redirect from='/' to='/home' />
  </Route>
);


export default routes;
