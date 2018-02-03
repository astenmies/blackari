import React from 'react'
import { QueryRenderer, graphql } from 'react-relay'

const HelloQuery = graphql`
         query HelloQuery{
           hello
         }
     `

class Hello extends React.Component {
  render() {
    return (
      <div>
        <QueryRenderer
          environment={this.props.environment}
          query={HelloQuery}
          operationName="HelloQuery"
          variables=""
          render={({ error, props }) => {
              if (error) {
                  return <div>{error.message}</div>;
                } else if (props) {
            return <div>Hello {props.hello}</div>;
            }
            return <div>...</div>;
          }}
        />
      </div>
    )
  }
}
export default Hello