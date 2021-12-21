import { Component } from '/tbc/core/components/component.js';
import { IconEnumPicker } from '/tbc/core/components/icon_enum_picker.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Blessings } from '/tbc/core/proto/ui.js';
import { BlessingsAssignment } from '/tbc/core/proto/ui.js';
import { BlessingsAssignments } from '/tbc/core/proto/ui.js';
import { classColors, specIconsLarge, specNames, } from '/tbc/core/proto_utils/utils.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { implementedSpecs } from './presets.js';
const MAX_PALADINS = 4;
const NUM_SPECS = getEnumValues(Spec).length;
const specOrder = [
    Spec.SpecBalanceDruid,
    Spec.SpecHunter,
    Spec.SpecMage,
    Spec.SpecRetributionPaladin,
    Spec.SpecShadowPriest,
    Spec.SpecRogue,
    Spec.SpecElementalShaman,
    Spec.SpecWarlock,
    Spec.SpecWarrior,
];
// Makes a new set of assignments with everything 0'd out.
function makeBlankBlessingsAssignments() {
    const assignments = BlessingsAssignments.create();
    for (let i = 0; i < MAX_PALADINS; i++) {
        assignments.paladins.push(BlessingsAssignment.create({
            blessings: new Array(NUM_SPECS).fill(Blessings.BlessingUnknown),
        }));
    }
    return assignments;
}
function makeBlessingsAssignments(data) {
    const assignments = makeBlankBlessingsAssignments();
    for (let i = 0; i < data.length; i++) {
        const spec = data[i].spec;
        const blessings = data[i].blessings;
        for (let j = 0; j < blessings.length; j++) {
            assignments.paladins[j].blessings[spec] = blessings[j];
        }
    }
    return assignments;
}
const defaultBlessings = makeBlessingsAssignments([
    { spec: Spec.SpecBalanceDruid, blessings: [Blessings.BlessingOfKings, Blessings.BlessingOfSalvation, Blessings.BlessingOfWisdom] },
    { spec: Spec.SpecHunter, blessings: [Blessings.BlessingOfKings, Blessings.BlessingOfSalvation, Blessings.BlessingOfMight, Blessings.BlessingOfWisdom] },
    { spec: Spec.SpecMage, blessings: [Blessings.BlessingOfKings, Blessings.BlessingOfSalvation, Blessings.BlessingOfWisdom] },
    { spec: Spec.SpecRetributionPaladin, blessings: [Blessings.BlessingOfKings, Blessings.BlessingOfMight, Blessings.BlessingOfSalvation, Blessings.BlessingOfWisdom] },
    { spec: Spec.SpecShadowPriest, blessings: [Blessings.BlessingOfKings, Blessings.BlessingOfSalvation, Blessings.BlessingOfWisdom] },
    { spec: Spec.SpecRogue, blessings: [Blessings.BlessingOfKings, Blessings.BlessingOfSalvation, Blessings.BlessingOfMight] },
    { spec: Spec.SpecElementalShaman, blessings: [Blessings.BlessingOfKings, Blessings.BlessingOfSalvation, Blessings.BlessingOfWisdom] },
    { spec: Spec.SpecWarlock, blessings: [Blessings.BlessingOfKings, Blessings.BlessingOfSalvation, Blessings.BlessingOfWisdom] },
    { spec: Spec.SpecWarrior, blessings: [Blessings.BlessingOfKings, Blessings.BlessingOfSalvation, Blessings.BlessingOfMight] },
]);
export class BlessingsPicker extends Component {
    constructor(parentElem, raidSimUI) {
        super(parentElem, 'blessings-picker-root');
        this.changeEmitter = new TypedEvent();
        this.raidSimUI = raidSimUI;
        this.assignments = BlessingsAssignments.clone(defaultBlessings);
        this.rootElem.innerHTML = `
		<table class="blessings-table">
			<thead class="blessings-table-header">
				<tr class="blessings-table-header-row">
					<th class="blessings-table-header-cell"></th>
				</tr>
			</thead>
			<tbody class="blessings-table-body">
			</tbody>
		</table>
		`;
        const headerRow = this.rootElem.getElementsByClassName('blessings-table-header-row')[0];
        const bodyElem = this.rootElem.getElementsByClassName('blessings-table-body')[0];
        specOrder.forEach(spec => {
            if (!implementedSpecs.includes(spec)) {
                return;
            }
            const cell = document.createElement('th');
            cell.classList.add('blessings-table-header-cell');
            headerRow.appendChild(cell);
            const icon = document.createElement('img');
            icon.src = specIconsLarge[spec];
            cell.appendChild(icon);
            tippy(icon, {
                'content': specNames[spec],
                'allowHTML': true,
            });
        });
        this.rows = [...Array(MAX_PALADINS).keys()].map(rowIndex => {
            const row = document.createElement('tr');
            row.classList.add('blessings-table-row');
            bodyElem.appendChild(row);
            const cell = document.createElement('td');
            cell.classList.add('blessings-table-cell', 'blessings-table-label-cell');
            cell.textContent = 'Paladin ' + (rowIndex + 1);
            row.appendChild(cell);
            specOrder.forEach(spec => {
                if (!implementedSpecs.includes(spec)) {
                    return;
                }
                const cell = document.createElement('td');
                cell.classList.add('blessings-table-cell');
                row.appendChild(cell);
                const blessingPicker = new IconEnumPicker(cell, this, {
                    extraCssClasses: [
                        'blessing-picker',
                    ],
                    values: [
                        { color: classColors[Class.ClassPaladin], value: Blessings.BlessingUnknown },
                        { id: { spellId: 25898 }, value: Blessings.BlessingOfKings },
                        { id: { spellId: 25895 }, value: Blessings.BlessingOfSalvation },
                        { id: { spellId: 27141 }, value: Blessings.BlessingOfMight },
                        { id: { spellId: 27143 }, value: Blessings.BlessingOfWisdom },
                    ],
                    changedEvent: (picker) => picker.changeEmitter,
                    getValue: (picker) => picker.assignments.paladins[rowIndex]?.blessings[spec] || Blessings.BlessingUnknown,
                    setValue: (eventID, picker, newValue) => {
                        const currentValue = picker.assignments.paladins[rowIndex].blessings[spec];
                        if (currentValue != newValue) {
                            picker.assignments.paladins[rowIndex].blessings[spec] = newValue;
                            this.changeEmitter.emit(eventID);
                        }
                    },
                });
            });
            return row;
        });
        this.setNumPaladins(raidSimUI.getClassCount(Class.ClassPaladin));
        raidSimUI.compChangeEmitter.on(eventID => {
            this.setNumPaladins(raidSimUI.getClassCount(Class.ClassPaladin));
        });
    }
    setNumPaladins(numPaladins) {
        numPaladins = Math.min(numPaladins, MAX_PALADINS);
        for (let i = 0; i < numPaladins; i++) {
            this.rows[i].classList.add('paladin-active');
        }
        for (let i = numPaladins; i < this.rows.length; i++) {
            this.rows[i].classList.remove('paladin-active');
        }
    }
    getAssignments() {
        // Defensive copy.
        return BlessingsAssignments.clone(this.assignments);
    }
}
