import { getIconUrl } from '/tbc/core/resources.js';
import { getName } from '/tbc/core/resources.js';
export function getActionId(actionMetric) {
    if (actionMetric.id.rawId.oneofKind == 'spellId') {
        return {
            id: {
                spellId: actionMetric.id.rawId.spellId,
            },
            tag: actionMetric.id.tag,
        };
    }
    else if (actionMetric.id.rawId.oneofKind == 'itemId') {
        return {
            id: {
                itemId: actionMetric.id.rawId.itemId,
            },
            tag: actionMetric.id.tag,
        };
    }
    else if (actionMetric.id.rawId.oneofKind == 'otherId') {
        return {
            id: {
                otherId: actionMetric.id.rawId.otherId,
            },
            tag: actionMetric.id.tag,
        };
    }
    else {
        throw new Error('Invalid action metric with no ID');
    }
}
export function parseActionMetrics(actionMetricProtos) {
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
    return Promise.all(actionMetrics.map(actionMetric => getName(actionMetric.actionId.id)
        .then(name => {
        if (actionMetric.actionId.tag == 0) {
            actionMetric.name = name;
        }
        else {
            actionMetric.name = name + ' (LO)';
        }
    })
        .then(() => getIconUrl(actionMetric.actionId.id))
        .then(iconUrl => actionMetric.iconUrl = iconUrl)))
        .then(() => actionMetrics);
}
