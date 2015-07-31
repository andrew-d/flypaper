import React from 'react';
import { RouteHandler } from 'react-router';
import { Nav, Navbar } from 'react-bootstrap';
import { NavItemLink } from 'react-router-bootstrap';


export default class App extends React.Component {
  render() {
    return (
      <div className='page-wrapper'>
        <Navbar fluid={true} staticTop={true} brand='Template'>
          <Nav>
            <NavItemLink to='index'>
              Home
            </NavItemLink>
            <NavItemLink to='about'>
              About
            </NavItemLink>
          </Nav>
        </Navbar>

        <div className='container-fluid'>
          <RouteHandler />
        </div>
      </div>
    );
  }
}
