
import React, { Component } from 'react'
// Lychee refs


class RecursiveChildComponent extends React.Component{
    render() {
      return <div>
        {this.recursiveCloneChildren(this.props.children)}
      </div>
    }

    recursiveCloneChildren(children) {
      return React.Children.map(children, child => {
        if(!React.isValidElement(child)) return child;
        var childProps = {
            id: "propToAdd"
        };
        childProps.children = this.recursiveCloneChildren(child.props.children);
        return <span>{ React.cloneElement(child, childProps)}</span>
      })
    }
  }

  const Coucou = () => {
      return(
          <div>
              <RecursiveChildComponent>
                  <div>Coucou 1</div>
                  <div>
                      <p>
                          Coucou nested
                      </p>
                  </div>
              </RecursiveChildComponent>
          </div>
      )
  }

  export default Coucou