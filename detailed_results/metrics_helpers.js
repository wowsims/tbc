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
        else if (name == "Mind Flay") { // for now we can just check the name and use special tagging rules.
            if (actionMetric.actionId.tag == 1) {
                actionMetric.name = name + ' (1 Tick)';
            }
            else if (actionMetric.actionId.tag == 2) {
                actionMetric.name = name + ' (2 Tick)';
            }
            else if (actionMetric.actionId.tag == 3) {
                actionMetric.name = name + ' (3 Tick)';
            }
        }
        else {
            if (actionMetric.actionId.tag == 1) {
                actionMetric.name = name + ' (LO)';
            }
            else {
                actionMetric.name = name + ' (??)';
            }
        }
    })
        .then(() => getIconUrl(actionMetric.actionId.id))
        .then(iconUrl => actionMetric.iconUrl = iconUrl)))
        .then(() => actionMetrics);
}
export function parseAuraMetrics(auraMetricProtos) {
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
            uptimeSecondsAvg: auraMetric.uptimeSecondsAvg,
            uptimeSecondsStdev: auraMetric.uptimeSecondsStdev,
        };
    });
    return Promise.all(auraMetrics.map(auraMetric => getName(auraMetric.actionId.id)
        .then(name => {
        auraMetric.name = name;
    })
        .then(() => getIconUrl(auraMetric.actionId.id))
        .then(iconUrl => auraMetric.iconUrl = iconUrl)))
        .then(() => auraMetrics);
}
