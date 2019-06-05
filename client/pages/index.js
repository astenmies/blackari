
import React, { Component } from 'react'
import { graphql } from 'react-relay'

// Lychee refs
import Hello from '../components/Hello.js'
import withData from '../lib/withData'

class IndexPage extends Component {
  static displayName = `IndexPage`

  render(props) {
    return (
      <div>
        <Hello {...this.props} />
      </div>
    )
  }
}

export default withData(IndexPage)