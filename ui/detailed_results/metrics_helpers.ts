import { ActionMetric as ActionMetricProto } from '/tbc/core/proto/api.js';
import { ActionId } from '/tbc/core/resources.js';
import { getIconUrl } from '/tbc/core/resources.js';
import { getName } from '/tbc/core/resources.js';

export type ActionMetric = {
	actionId: ActionId,
	name: string,
	tagIndex: number,
	iconUrl: string,
	casts: number,
	crits: number,
	misses: number,
	totalDmg: number,
};

export function getActionId(actionMetric: ActionMetricProto): ActionId {
	if (actionMetric.actionId.oneofKind == 'spellId') {
		return {
			spellId: actionMetric.actionId.spellId,
		};
	} else if (actionMetric.actionId.oneofKind == 'itemId') {
		return {
			spellId: actionMetric.actionId.itemId,
		};
	} else if (actionMetric.actionId.oneofKind == 'otherId') {
		return {
			spellId: actionMetric.actionId.otherId,
		};
	} else {
		throw new Error('Invalid action metric with no ID');
	}
}

export function parseActionMetrics(actionMetricProtos: Array<ActionMetricProto>): Promise<Array<ActionMetric>> {
	const actionMetrics = actionMetricProtos.map(actionMetric => actionMetric.casts.map((_, i) => {
		return {
			actionId: getActionId(actionMetric),
			name: '',
			iconUrl: '',
			tagIndex: i,
			casts: actionMetric.casts[i],
			crits: actionMetric.crits[i],
			misses: actionMetric.misses[i],
			totalDmg: actionMetric.dmgs[i],
		};
	})).flat();

	return Promise.all(actionMetrics.map(actionMetric => 
		getName(actionMetric.actionId)
		.then(name => actionMetric.name = name)
		.then(() => getIconUrl(actionMetric.actionId))
		.then(iconUrl => actionMetric.iconUrl = iconUrl)
	))
	.then(() => actionMetrics);
}
