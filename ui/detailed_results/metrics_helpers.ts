import { ActionMetrics as ActionMetricsProto } from '/tbc/core/proto/api.js';
import { ActionId } from '/tbc/core/resources.js';
import { getIconUrl } from '/tbc/core/resources.js';
import { getName } from '/tbc/core/resources.js';

export type ActionMetrics = {
	actionId: ActionId,
	name: string,
	iconUrl: string,
	casts: number,
	hits: number,
	crits: number,
	misses: number,
	totalDmg: number,
};

export function getActionId(actionMetric: ActionMetricsProto): ActionId {
	if (actionMetric.id!.rawId.oneofKind == 'spellId') {
		return {
			id: {
				spellId: actionMetric.id!.rawId.spellId,
			},
			tag: actionMetric.id!.tag,
		};
	} else if (actionMetric.id!.rawId.oneofKind == 'itemId') {
		return {
			id: {
				itemId: actionMetric.id!.rawId.itemId,
			},
			tag: actionMetric.id!.tag,
		};
	} else if (actionMetric.id!.rawId.oneofKind == 'otherId') {
		return {
			id: {
				otherId: actionMetric.id!.rawId.otherId,
			},
			tag: actionMetric.id!.tag,
		};
	} else {
		throw new Error('Invalid action metric with no ID');
	}
}

export function parseActionMetrics(actionMetricProtos: Array<ActionMetricsProto>): Promise<Array<ActionMetrics>> {
	const actionMetrics = actionMetricProtos.map(actionMetric => {
		return {
			actionId: getActionId(actionMetric),
			name: '',
			iconUrl: '',
			casts: actionMetric.casts,
			hits: actionMetric.hits,
			crits: actionMetric.crits,
			misses: actionMetric.misses,
			totalDmg: actionMetric.damage,
		};
	});

	return Promise.all(actionMetrics.map(actionMetric => 
		getName(actionMetric.actionId.id)
		.then(name => {
			if (actionMetric.actionId.tag == 0) {
				actionMetric.name = name;
			} else if (actionMetric.actionId.tag == 1) {
				actionMetric.name = name + ' (LO)';
			} else if (actionMetric.actionId.tag == 2) {
				actionMetric.name = name + ' (2 Tick)';
			} else if (actionMetric.actionId.tag == 3) {
				actionMetric.name = name + ' (3 Tick)';
			} else {
				actionMetric.name = name + ' (??)';
			}
		})
		.then(() => getIconUrl(actionMetric.actionId.id))
		.then(iconUrl => actionMetric.iconUrl = iconUrl)
	))
	.then(() => actionMetrics);
}
