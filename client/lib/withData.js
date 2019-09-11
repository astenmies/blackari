import React from "react";
import initEnvironment from "./createRelayEnvironment";
import { fetchQuery, ReactRelayContext } from "react-relay";

export default (ComposedComponent, options = {}) => {
	return class WithData extends React.Component {
		static displayName = `WithData(${ComposedComponent.displayName})`;

		static async getInitialProps(ctx) {
			// Evaluate the composed component's getInitialProps()
			let composedInitialProps = {};
			if (ComposedComponent.getInitialProps) {
				composedInitialProps = await ComposedComponent.getInitialProps(ctx);
			}

			let queryProps = {};
			let queryRecords = {};
			const environment = initEnvironment();

			let variables = {};
			if (options.query) {
				variables = composedInitialProps.relayVariables || {};
				// TODO: Consider RelayQueryResponseCache
				// https://github.com/facebook/relay/issues/1687#issuecomment-302931855
				queryProps = await fetchQuery(environment, options.query, variables);
				queryRecords = environment
					.getStore()
					.getSource()
					.toJSON();
			}

			return {
				...composedInitialProps,
				...queryProps,
				queryRecords,
				variables
			};
		}

		constructor(props) {
			super(props);
			this.variables = props.variables;
			this.environment = initEnvironment({
				records: props.queryRecords
			});
		}

		render() {
			return (
				<ReactRelayContext.Provider value={{ ...this.props, environment: this.environment, variables: {} }}>
					<ComposedComponent {...this.props} />
				</ReactRelayContext.Provider>
			);
		}
	};
};
