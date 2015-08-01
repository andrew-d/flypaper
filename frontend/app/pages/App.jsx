import React from 'react';
/*
import { Nav, Navbar } from 'react-bootstrap';
import NavItemLink from '../utils/NavItemLink';
*/


export default class App extends React.Component {
  render() {
    return (
      <div className='page-wrapper'>
        {/*
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
        */}

        <div className='container-fluid'>
			{this.props.children}
        </div>
      </div>
    );
  }
}
