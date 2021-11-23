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
	hits: number,
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
			itemId: actionMetric.actionId.itemId,
		};
	} else if (actionMetric.actionId.oneofKind == 'otherId') {
		return {
			otherId: actionMetric.actionId.otherId,
		};
	} else {
		throw new Error('Invalid action metric with no ID');
	}
}

export function parseActionMetrics(actionMetricProtos: Array<ActionMetricProto>): Promise<Array<ActionMetric>> {
	const actionMetrics = actionMetricProtos.map(actionMetric => {
		return {
			actionId: getActionId(actionMetric),
			name: '',
			iconUrl: '',
			tagIndex: actionMetric.tag,
			casts: actionMetric.casts,
			hits: actionMetric.hits,
			crits: actionMetric.crits,
			misses: actionMetric.misses,
			totalDmg: actionMetric.damage,
		};
	});

	return Promise.all(actionMetrics.map(actionMetric => 
		getName(actionMetric.actionId)
		.then(name => {
			// TODO: We should map these in each class specifically since they are not currently shared.
			//   The other option is switch on spell (only certain spells need tags)
			if (actionMetric.tagIndex == 0) {
				actionMetric.name = name;
			} else if (actionMetric.tagIndex == 1) {
				actionMetric.name = name + ' (LO)';
			} else if (actionMetric.tagIndex == 2) {
				actionMetric.name = name + ' (2 Tick)';
			} else if (actionMetric.tagIndex == 3) {
				actionMetric.name = name + ' (3 Tick)';
			} else {
				actionMetric.name = name + ' (??)';
			}
		})
		.then(() => getIconUrl(actionMetric.actionId))
		.then(iconUrl => actionMetric.iconUrl = iconUrl)
	))
	.then(() => actionMetrics);
}
