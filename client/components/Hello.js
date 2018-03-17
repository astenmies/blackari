import React from 'react'
import { QueryRenderer, graphql } from 'react-relay'

const HelloQuery = graphql`
         query HelloQuery{
          post(slug: "second-post") {
            title
          }
         }
     `

class Hello extends React.Component {
  render() {
    return (
      <div>
        <QueryRenderer
          environment={this.props.environment}
          query={HelloQuery}
          operationName="post"
          variables=""
          render={({ error, props }) => {
              if (error) {
                  return <div>{error.message}</div>;
                } else if (props) {
            return (
              <div>This comes from Blackari server:
                <span style={{color: 'red'}}> {props.post.title}</span>
              </div>
              );
            }
            return <div>...</div>;
          }}
        />
      </div>
    )
  }
}
export default Hello