import React from "react";
import Relay, { createFragmentContainer, graphql } from "react-relay";
import BlogPostPreview from "./BlogPostPreview";
import Router from "next/router";
import _ from "lodash";

class BlogPosts extends React.Component {
	render() {
		const { props } = this;
		if (typeof window != "undefined") console.log("props ---", _.get(props, "data.allBlogPosts"));

		// return "BlogPosts";

		let afterParam = _.get(props, "data.allBlogPosts.pageInfo.endCursor");
		afterParam = afterParam ? `&after=${afterParam}` : "";

		let hasNextPage = _.get(props, "data.allBlogPosts.pageInfo.hasNextPage");
		hasNextPage = hasNextPage || props.relayVariables.before;

		let hasPrevPage = _.get(props, "data.allBlogPosts.pageInfo.hasPreviousPage");
		hasPrevPage = hasPrevPage || props.relayVariables.after;

		let beforeParam = _.get(props, "data.allBlogPosts.pageInfo.startCursor");
		beforeParam = beforeParam ? `&before=${beforeParam}` : "";

		const nextOnClick = () => Router.push(`/?first=2${afterParam}`);
		const prevOnClick = () => Router.push(`/?last=2${beforeParam}`);
		if (!props.data.allBlogPosts) return "Loading...";
		return (
			<div>
				<h1>Blog posts</h1>
				{props.data.allBlogPosts.edges.map(({ node }) => (
					<BlogPostPreview key={node.id} post={node} />
				))}
				<br />
				<p>Next {afterParam}</p>
				<p>Previous {beforeParam}</p>
				<button disabled={!hasPrevPage} onClick={prevOnClick}>
					Previous Page
				</button>
				&nbsp;
				<button disabled={!hasNextPage} onClick={nextOnClick}>
					Next Page
				</button>
			</div>
		);
	}
}

export default Relay.createFragmentContainer(BlogPosts, {
	data: graphql`
		fragment BlogPosts_data on Root
			@argumentDefinitions(after: { type: "String" }, before: { type: "String" }, first: { type: "Int", defaultValue: 2 }, last: { type: "Int" }) {
			allBlogPosts(after: $after, before: $before, first: $first, last: $last) {
				pageInfo {
					hasNextPage
					hasPreviousPage
					startCursor
					endCursor
				}
				edges {
					node {
						...BlogPostPreview_post
						title
						id
					}
				}
			}
		}
	`
});
