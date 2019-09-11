import React, { Component } from "react";
import { graphql } from "react-relay";
import withData from "../lib/withData";
import BlogPosts from "../components/BlogPosts";
import { withRouter } from "next/router";

class Index extends Component {
	static displayName = `Index`;

	static async getInitialProps(context) {
		let { after, before, first, last } = context.query;

		if (last === undefined) {
			first = first || 2;
		}

		return {
			relayVariables: {
				after,
				before,
				first: first ? parseInt(first, 10) : first,
				last: last ? parseInt(last, 10) : last
			}
		};
	}

	render() {
		const { __fragments, __id, __fragmentOwner, relayVariables, ...nonQueryProps } = this.props;
		return (
			<div>
				<BlogPosts {...nonQueryProps} data={{ __fragments, __id, __fragmentOwner }} relayVariables={relayVariables} />
			</div>
		);
	}
}

export default withData(Index, {
	query: graphql`
		query pages_indexQuery($after: String, $before: String, $first: Int, $last: Int) {
			...BlogPosts_data @arguments(after: $after, before: $before, first: $first, last: $last)
		}
	`
});
