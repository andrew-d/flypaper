import Icon from 'react-fa';
import React from 'react';


export default class About extends React.Component {
  render() {
    return (
      <div>
        <p>This is the about page, demonstrating an inline style.</p>

        <p>
          This is an icon from the react-fa project: <Icon name='spinner' spin={true} />
        </p>
      </div>
    );
  }
}
