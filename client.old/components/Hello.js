import React from "react";
import { QueryRenderer, graphql } from "react-relay";

const HelloQuery = graphql`
  query HelloQuery {
    post(slug: "second") {
      title
    }
  }
`;

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
            console.log("PROPS --", props);
            if (error) {
              return <div>{error.message}</div>;
            } else if (props) {
              let post = props.post || {};

              return (
                <div>
                  This comes from Lychee server:
                  <span style={{ color: "red" }}> {post.title}</span>
                </div>
              );
            }
            return <div>...</div>;
          }}
        />
      </div>
    );
  }
}
export default Hello;
