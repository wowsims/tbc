import { ActionMetrics as ActionMetricsProto } from '/tbc/core/proto/api.js';
import { AuraMetrics as AuraMetricsProto } from '/tbc/core/proto/api.js';
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

export type AuraMetrics = {
	actionId: ActionId,
	name: string,
	iconUrl: string,
	uptimeSeconds: number,
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
			} else {
				actionMetric.name = name + ' (LO)';
			}
		})
		.then(() => getIconUrl(actionMetric.actionId.id))
		.then(iconUrl => actionMetric.iconUrl = iconUrl)
	))
	.then(() => actionMetrics);
}

export function parseAuraMetrics(auraMetricProtos: Array<AuraMetricsProto>): Promise<Array<AuraMetrics>> {
	const auraMetrics = auraMetricProtos.map(auraMetric => {
		return {
			actionId: {
				id: {
					spellId: auraMetric.id,
				},
				tag: 0,
			},
			name: '',
			iconUrl: '',
			uptimeSeconds: auraMetric.uptimeSeconds,
		};
	});

	return Promise.all(auraMetrics.map(auraMetric => 
		getName(auraMetric.actionId.id)
		.then(name => {
			auraMetric.name = name;
		})
		.then(() => getIconUrl(auraMetric.actionId.id))
		.then(iconUrl => auraMetric.iconUrl = iconUrl)
	))
	.then(() => auraMetrics);
}

