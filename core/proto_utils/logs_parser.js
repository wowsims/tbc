export class Entity {
    constructor(name, ownerName, index, isTarget, isPet) {
        this.name = name;
        this.ownerName = ownerName;
        this.index = index;
        this.isTarget = isTarget;
        this.isPet = isPet;
    }
    static parseAll(str) {
        return Array.from(str.matchAll(Entity.parseRegex)).map(match => {
            if (match[1]) {
                return new Entity(match[1], '', parseInt(match[2]) - 1, true, false);
            }
            else if (match[3]) {
                return new Entity(match[6], match[4], parseInt(match[5]) - 1, false, true);
            }
            else if (match[7]) {
                return new Entity(match[8], '', parseInt(match[9]) - 1, false, false);
            }
            else {
                throw new Error('Invalid Entity match');
            }
        });
    }
}
// Parses one or more Entities from a string.
// Each entity label should be one of:
//   'Target 1' if a target,
//   'PlayerName (#1)' if a player, or
//   'PlayerName (#1) - PetName' if a pet.
Entity.parseRegex = /(Target (\d+))|(([a-zA-Z0-9]+) \(#(\d+)\) - ([a-zA-Z0-9]+))|(([a-zA-Z0-9]+) \(#(\d+)\))/g;
export class SimLog {
    constructor(params) {
        this.raw = params.raw;
        this.timestamp = params.timestamp;
        this.source = params.source;
        this.target = params.target;
    }
    static parseAll(result) {
        const lines = result.logs.split('\n');
        return lines.map(line => {
            const params = {
                raw: line,
                timestamp: 0,
                source: null,
                target: null,
            };
            let match = line.match(/\[([0-9]+\.[0-9]+)\]\w*(.*)/);
            if (!match || !match[1]) {
                return new SimLog(params);
            }
            params.timestamp = parseFloat(match[1]);
            let remainder = match[2];
            const entities = Entity.parseAll(remainder);
            params.source = entities[0] || null;
            params.target = entities[1] || null;
            return AuraGainedLog.parse(params)
                || AuraFadedLog.parse(params)
                || ManaChangedLog.parse(params)
                || new SimLog(params);
        });
    }
}
export class AuraGainedLog extends SimLog {
    constructor(params, auraName) {
        super(params);
        this.auraName = auraName;
    }
    static parse(params) {
        const match = params.raw.match(/Aura gained: \[(.*)\]/);
        if (match && match[1]) {
            return new AuraGainedLog(params, match[1]);
        }
        else {
            return null;
        }
    }
}
export class AuraFadedLog extends SimLog {
    constructor(params, auraName) {
        super(params);
        this.auraName = auraName;
    }
    static parse(params) {
        const match = params.raw.match(/Aura faded: \[(.*)\]/);
        if (match && match[1]) {
            return new AuraFadedLog(params, match[1]);
        }
        else {
            return null;
        }
    }
}
export class ManaChangedLog extends SimLog {
    constructor(params, manaBefore, manaAfter, cause) {
        super(params);
        this.manaBefore = manaBefore;
        this.manaAfter = manaAfter;
        this.cause = cause;
    }
    static parse(params) {
        const match = params.raw.match(/[Gained|Spent] \d+\.\d+ mana from (.*) \((\d+\.\d+) --> (\d+\.\d+)\)/);
        if (match && match[1]) {
            let cause = match[1];
            //if (cause.endsWith('s Regen')) {
            //	cause = 'Regen';
            //}
            return new ManaChangedLog(params, parseFloat(match[2]), parseFloat(match[3]), cause);
        }
        else {
            return null;
        }
    }
}
