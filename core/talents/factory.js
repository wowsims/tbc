import { Spec } from '/tbc/core/proto/common.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { ShamanTalentsPicker } from './shaman.js';
export function newTalentsPicker(spec, parent, sim) {
    switch (spec) {
        case Spec.SpecElementalShaman:
            return new ShamanTalentsPicker(parent, sim);
            break;
        default:
            const playerClass = specToClass[spec];
            throw new Error('Unimplemented class talents: ' + playerClass);
    }
}
